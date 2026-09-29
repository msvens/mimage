package video

import (
	"context"
	"errors"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/msvens/mimage/img"
	"github.com/msvens/mimage/metadata"
)

// skipFFmpegEnv opts out of the tests that run ffprobe and ffmpeg. It exists
// for CI runners where installing ffmpeg is not worth the time; without it a
// missing ffmpeg fails the test run rather than quietly skipping
const skipFFmpegEnv = "MIMAGE_SKIP_FFMPEG"

func requireFFmpeg(t *testing.T) {
	t.Helper()
	if os.Getenv(skipFFmpegEnv) == "1" {
		t.Skipf("%s=1, not running ffmpeg", skipFFmpegEnv)
	}
	if err := Available(); err != nil {
		t.Fatalf("%v: install ffmpeg, or set %s=1 to skip the video integration tests", err, skipFFmpegEnv)
	}
}

func testContext(t *testing.T) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(cancel)
	return ctx
}

// Probe running the installed ffprobe must agree with the parsing of the
// captured output, whatever ffmpeg version is installed
func TestProbe(t *testing.T) {
	requireFFmpeg(t)
	for _, e := range expectations {
		s, err := Probe(testContext(t), assetPath+e.file)
		checkSummary(t, e, s, err)
	}
}

func TestProbeRejects(t *testing.T) {
	requireFFmpeg(t)
	ctx := testContext(t)
	if _, err := Probe(ctx, "../assets/leica.jpg"); !errors.Is(err, ErrNotVideo) {
		t.Errorf("an image: expected ErrNotVideo, got %v", err)
	}
	if _, err := Probe(ctx, assetPath+"missing.mp4"); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("a missing file: expected os.ErrNotExist, got %v", err)
	}

	//both of these pass detection as an mp4, which is why Probe is the real
	//check
	plain, err := os.ReadFile(assetPath + "plain.mp4")
	if err != nil {
		t.Fatalf("could not read fixture: %v", err)
	}
	ftypBox := plain[:32]
	dir := t.TempDir()

	//cut off inside the movie header. ffprobe would report no streams, the
	//box check catches it first
	truncated := filepath.Join(dir, "truncated.mp4")
	if err = os.WriteFile(truncated, plain[:64], 0644); err != nil {
		t.Fatalf("could not write: %v", err)
	}
	if _, err = Probe(ctx, truncated); !errors.Is(err, ErrTruncated) {
		t.Errorf("a truncated mp4: expected ErrTruncated, got %v", err)
	}

	//media data and no movie header at all: ffprobe itself fails
	noMoov := filepath.Join(dir, "nomoov.mp4")
	junk := append([]byte{0, 0, 1, 0, 'm', 'd', 'a', 't'}, make([]byte, 248)...)
	if err = os.WriteFile(noMoov, append(ftypBox, junk...), 0644); err != nil {
		t.Fatalf("could not write: %v", err)
	}
	var toolErr *ToolError
	if _, err = Probe(ctx, noMoov); !errors.As(err, &toolErr) || toolErr.Stderr == "" {
		t.Errorf("an mp4 without moov: expected a ToolError with stderr, got %v", err)
	}
}

func TestMissingTool(t *testing.T) {
	saved := FFprobePath
	FFprobePath = "mimage-no-such-ffprobe"
	t.Cleanup(func() { FFprobePath = saved })

	if err := Available(); !errors.Is(err, ErrFFmpegNotFound) {
		t.Errorf("Available: expected ErrFFmpegNotFound, got %v", err)
	}
	if _, err := Probe(context.Background(), assetPath+"plain.mp4"); !errors.Is(err, ErrFFmpegNotFound) {
		t.Errorf("Probe: expected ErrFFmpegNotFound, got %v", err)
	}
}

func decode(t *testing.T, fileName string) image.Image {
	t.Helper()
	f, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("could not open %s: %v", fileName, err)
	}
	defer func() { _ = f.Close() }()
	im, _, err := image.Decode(f)
	if err != nil {
		t.Fatalf("could not decode %s: %v", fileName, err)
	}
	return im
}

// dominant names the strongest channel of a pixel, or "black" when it is dark
func dominant(c color.Color) string {
	r, g, b, _ := c.RGBA()
	switch {
	case max(r, g, b) < 0x4000:
		return "black"
	case r > g && r > b:
		return "red"
	case g > r && g > b:
		return "green"
	default:
		return "blue"
	}
}

func extract(t *testing.T, file string, opts PosterOptions) image.Image {
	t.Helper()
	out, err := ExtractPoster(testContext(t), assetPath+file, filepath.Join(t.TempDir(), "poster"), opts)
	if err != nil {
		t.Fatalf("%s: %v", file, err)
	}
	if filepath.Ext(out) != ".jpg" {
		t.Errorf("%s: poster written to %s, want a .jpg", file, out)
	}
	return decode(t, out)
}

func TestExtractPoster(t *testing.T) {
	requireFFmpeg(t)

	if b := extract(t, "plain.mp4", PosterOptions{}).Bounds(); b.Dx() != 320 || b.Dy() != 240 {
		t.Errorf("plain: poster is %v, want 320x240", b)
	}
	if b := extract(t, "plain.mov", PosterOptions{}).Bounds(); b.Dx() != 320 || b.Dy() != 240 {
		t.Errorf("mov: poster is %v, want 320x240", b)
	}

	//stored with red on top, turned 90 degrees clockwise for display: red
	//ends up on the right
	rotated := extract(t, "rotated.mp4", PosterOptions{})
	if b := rotated.Bounds(); b.Dx() != 240 || b.Dy() != 320 {
		t.Errorf("rotated: poster is %v, want 240x320", b)
	}
	if left, right := dominant(rotated.At(40, 160)), dominant(rotated.At(200, 160)); left != "blue" || right != "red" {
		t.Errorf("rotated: left is %s and right is %s, want blue and red", left, right)
	}

	//the default offset skips the black opening, an explicit one at the
	//start does not
	if c := dominant(extract(t, "blackstart.mp4", PosterOptions{}).At(160, 120)); c != "green" {
		t.Errorf("blackstart: poster is %s, want green", c)
	}
	if c := dominant(extract(t, "blackstart.mp4", PosterOptions{At: time.Millisecond}).At(160, 120)); c != "black" {
		t.Errorf("blackstart at the start: poster is %s, want black", c)
	}

	//shorter than the offset: clamped rather than seeking past the end
	if c := dominant(extract(t, "short.mp4", PosterOptions{}).At(160, 120)); c != "green" {
		t.Errorf("short: poster is %s, want green", c)
	}
	if c := dominant(extract(t, "short.mp4", PosterOptions{At: time.Hour}).At(160, 120)); c != "green" {
		t.Errorf("short with a huge offset: poster is %s, want green", c)
	}
}

func TestExtractPosterRejects(t *testing.T) {
	requireFFmpeg(t)
	ctx := testContext(t)
	dir := t.TempDir()

	if _, err := ExtractPoster(ctx, assetPath+"plain.mp4", filepath.Join(dir, "poster.jpg"), PosterOptions{}); !errors.Is(err, img.ErrDestHasExtension) {
		t.Errorf("destination with extension: expected ErrDestHasExtension, got %v", err)
	}
	if _, err := ExtractPoster(ctx, assetPath+"coverart.mp4", filepath.Join(dir, "a"), PosterOptions{}); !errors.Is(err, ErrNoVideoStream) {
		t.Errorf("cover art only: expected ErrNoVideoStream, got %v", err)
	}
	if _, err := ExtractPoster(ctx, assetPath+"plain.mp4", filepath.Join(dir, "missing", "b"), PosterOptions{}); err == nil {
		t.Errorf("unwritable destination: expected an error")
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 0 {
		t.Errorf("a failed extraction should leave nothing behind, found %d entries", len(entries))
	}
}

// The flow a photo service runs: detect, probe, extract a poster and make the
// same variants it makes for a photo. A portrait phone video must come out as
// portrait thumbnails
func TestPosterThroughImagePipeline(t *testing.T) {
	requireFFmpeg(t)
	ctx := testContext(t)
	src := assetPath + "rotated.mp4"

	format, err := metadata.DetectFormatFile(src)
	if err != nil || !format.Supported() || !format.IsVideo() {
		t.Fatalf("detect: got %v, %v, want a supported video", format, err)
	}
	s, err := Probe(ctx, src)
	if err != nil {
		t.Fatalf("probe: %v", err)
	}
	dir := t.TempDir()
	poster, err := ExtractPoster(ctx, src, filepath.Join(dir, "poster"), PosterOptions{})
	if err != nil {
		t.Fatalf("poster: %v", err)
	}

	variants := map[string]img.Options{
		filepath.Join(dir, "thumb"):     img.NewOptions(img.ResizeAndCrop, 400, 400, false, img.FormatJpeg),
		filepath.Join(dir, "landscape"): img.NewOptions(img.ResizeAndCrop, 1200, 628, true, img.FormatJpeg),
		filepath.Join(dir, "square"):    img.NewOptions(img.ResizeAndCrop, 1200, 1200, true, img.FormatJpeg),
		filepath.Join(dir, "portrait"):  img.NewOptions(img.ResizeAndCrop, 1080, 1350, true, img.FormatJpeg),
		filepath.Join(dir, "resize"):    img.NewOptions(img.Resize, 1200, 0, true, img.FormatJpeg),
	}
	if err = img.TransformFile(poster, variants); err != nil {
		t.Fatalf("transform: %v", err)
	}

	//resize keeps the aspect ratio, so it shows the poster was upright
	want := map[string][2]int{
		"thumb": {400, 400}, "landscape": {1200, 628}, "square": {1200, 1200},
		"portrait": {1080, 1350}, "resize": {1200, 1200 * s.Height / s.Width},
	}
	for name, dims := range want {
		b := decode(t, filepath.Join(dir, name+".jpg")).Bounds()
		if b.Dx() != dims[0] || b.Dy() != dims[1] {
			t.Errorf("%s: %dx%d, want %dx%d", name, b.Dx(), b.Dy(), dims[0], dims[1])
		}
	}
	if resize := decode(t, filepath.Join(dir, "resize.jpg")); dominant(resize.At(1100, 800)) != "red" {
		t.Errorf("resize: expected red on the right, the poster was not upright")
	}
}
