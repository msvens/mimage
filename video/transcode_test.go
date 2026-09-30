package video

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/msvens/mimage/img"
	"github.com/msvens/mimage/metadata"
)

// a web ready 1080p source, the starting point the decision tests vary
func webReady() *Summary {
	return &Summary{Width: 1920, Height: 1080, StoredWidth: 1920, StoredHeight: 1080,
		VideoCodec: "h264", VideoProfile: "High", PixelFormat: "yuv420p", PixelAspect: 1,
		AudioCodec: "aac", FrameRate: 30, Duration: 10 * time.Second}
}

func defaults(t *testing.T, o TranscodeOptions) transcodeSettings {
	t.Helper()
	s, err := o.settings()
	if err != nil {
		t.Fatalf("settings: %v", err)
	}
	return s
}

// value returns the argument following flag, or "" when flag is absent
func value(args []string, flag string) string {
	if i := slices.Index(args, flag); i >= 0 && i+1 < len(args) {
		return args[i+1]
	}
	return ""
}

func TestTranscodeDecision(t *testing.T) {
	std := defaults(t, TranscodeOptions{})
	tests := []struct {
		name   string
		change func(s *Summary)
		opts   TranscodeOptions
		encode bool
	}{
		{"web ready", func(*Summary) {}, TranscodeOptions{}, false},
		{"web ready, rotated by flag", func(s *Summary) { s.Rotation, s.Width, s.Height = 90, 1080, 1920 }, TranscodeOptions{}, false},
		{"force", func(*Summary) {}, TranscodeOptions{ForceEncode: true}, true},
		{"hevc", func(s *Summary) { s.VideoCodec = "hevc" }, TranscodeOptions{}, true},
		{"10 bit h264", func(s *Summary) { s.VideoProfile, s.PixelFormat = "High 10", "yuv420p10le" }, TranscodeOptions{}, true},
		{"full range jpeg pixels", func(s *Summary) { s.PixelFormat = "yuvj420p" }, TranscodeOptions{}, true},
		{"hdr", func(s *Summary) { s.HDR = true }, TranscodeOptions{}, true},
		{"interlaced", func(s *Summary) { s.Interlaced = true }, TranscodeOptions{}, true},
		{"anamorphic", func(s *Summary) { s.PixelAspect = 4.0 / 3 }, TranscodeOptions{}, true},
		{"4k", func(s *Summary) { s.Width, s.Height = 3840, 2160 }, TranscodeOptions{}, true},
		{"4k, no size cap", func(s *Summary) { s.Width, s.Height = 3840, 2160 }, TranscodeOptions{MaxShortSide: -1}, false},
		{"slow motion", func(s *Summary) { s.FrameRate = 240 }, TranscodeOptions{}, true},
		{"slow motion, no rate cap", func(s *Summary) { s.FrameRate = 240 }, TranscodeOptions{MaxFrameRate: -1}, false},
		{"29.97 under the cap", func(s *Summary) { s.FrameRate = 29.97 }, TranscodeOptions{MaxFrameRate: 30}, false},
	}
	for _, tc := range tests {
		s := webReady()
		tc.change(s)
		o := std
		if tc.opts != (TranscodeOptions{}) {
			o = defaults(t, tc.opts)
		}
		args, encode := transcodeArgs(s, "in", "out", o)
		if encode != tc.encode {
			t.Errorf("%s: encode = %v, want %v", tc.name, encode, tc.encode)
		}
		if got, want := value(args, "-c:v"), map[bool]string{true: "libx264", false: "copy"}[tc.encode]; got != want {
			t.Errorf("%s: -c:v %s, want %s", tc.name, got, want)
		}
	}
}

func TestTranscodeArgs(t *testing.T) {
	std := defaults(t, TranscodeOptions{})

	//audio is copied when it is aac, encoded otherwise
	s := webReady()
	args, _ := transcodeArgs(s, "in", "out", std)
	if value(args, "-c:a") != "copy" {
		t.Errorf("aac should be copied, got %v", args)
	}
	s.AudioCodec = "pcm_s16le"
	args, _ = transcodeArgs(s, "in", "out", std)
	if value(args, "-c:a") != "aac" || value(args, "-b:a") != "128k" {
		t.Errorf("pcm should be encoded to aac at 128k, got %v", args)
	}

	//the date is written explicitly in UTC, and the index goes first
	s.CreationTime = time.Date(2024, 7, 14, 14, 4, 6, 0, time.FixedZone("", 2*3600))
	args, _ = transcodeArgs(s, "in", "out", std)
	if !slices.Contains(args, "creation_time=2024-07-14T12:04:06Z") {
		t.Errorf("expected the creation time in UTC, got %v", args)
	}
	if !strings.Contains(value(args, "-movflags"), "+faststart") {
		t.Errorf("expected faststart, got %v", args)
	}
	//the source's container description is not carried over
	if !slices.Contains(args, "major_brand=") {
		t.Errorf("expected major_brand to be cleared, got %v", args)
	}

	//the stream Probe chose is mapped, not whatever ffmpeg would pick
	s.streamIndex = 2
	args, _ = transcodeArgs(s, "in", "out", std)
	if value(args, "-map") != "0:2" {
		t.Errorf("expected the video stream mapped by index, got %v", args)
	}
}

func TestTranscodeQuality(t *testing.T) {
	s := webReady()
	s.VideoCodec = "hevc"
	tests := []struct {
		opts   TranscodeOptions
		crf    string
		preset string
	}{
		{TranscodeOptions{}, "21", "medium"},
		{TranscodeOptions{Quality: QualityHigh}, "18", "slow"},
		{TranscodeOptions{Quality: QualitySmall}, "26", "medium"},
		//explicit settings beat the preset
		{TranscodeOptions{Quality: QualityHigh, CRF: 23, Preset: "fast"}, "23", "fast"},
	}
	for _, tc := range tests {
		args, _ := transcodeArgs(s, "in", "out", defaults(t, tc.opts))
		if value(args, "-crf") != tc.crf || value(args, "-preset") != tc.preset {
			t.Errorf("%+v: crf %s preset %s, want %s %s", tc.opts, value(args, "-crf"), value(args, "-preset"), tc.crf, tc.preset)
		}
	}
	if _, err := (TranscodeOptions{Quality: Quality(9)}).settings(); err == nil {
		t.Errorf("expected an error for an unknown quality")
	}
	if _, err := (TranscodeOptions{CRF: 60}).settings(); err == nil {
		t.Errorf("expected an error for a crf above 51")
	}
	for _, name := range []string{"standard", "High", "SMALL"} {
		if _, err := ParseQuality(name); err != nil {
			t.Errorf("ParseQuality(%q): %v", name, err)
		}
	}
	if _, err := ParseQuality("best"); err == nil {
		t.Errorf("expected an error for an unknown quality name")
	}
}

func TestVideoFilters(t *testing.T) {
	std := defaults(t, TranscodeOptions{})
	tests := []struct {
		name   string
		change func(s *Summary)
		opts   *TranscodeOptions
		want   string
	}{
		{"nothing to fix", func(*Summary) {}, nil, "format=yuv420p"},
		//a phone original: turned upright by ffmpeg, so no scaling needed
		{"rotated 1080p", func(s *Summary) { s.Rotation, s.Width, s.Height = 90, 1080, 1920 }, nil, "format=yuv420p"},
		{"rotated 4k", func(s *Summary) {
			s.Rotation, s.StoredWidth, s.StoredHeight, s.Width, s.Height = 90, 3840, 2160, 2160, 3840
		}, nil, "scale=1080:1920,setsar=1,format=yuv420p"},
		{"interlaced", func(s *Summary) { s.Interlaced = true }, nil, "bwdif=mode=send_frame,format=yuv420p"},
		{"anamorphic", func(s *Summary) {
			s.StoredWidth, s.StoredHeight, s.Width, s.Height, s.PixelAspect = 720, 480, 853, 480, 32.0/27
		}, nil, "scale=852:480,setsar=1,format=yuv420p"},
		{"slow motion", func(s *Summary) { s.FrameRate = 120 }, nil, "fps=60,format=yuv420p"},
		{"full range motion jpeg", func(s *Summary) { s.VideoCodec, s.PixelFormat = "mjpeg", "yuvj422p" }, nil,
			"scale=out_range=tv,format=yuv420p"},
		{"custom caps", func(s *Summary) { s.FrameRate = 50 }, &TranscodeOptions{MaxShortSide: 720, MaxFrameRate: 25},
			"scale=1280:720,setsar=1,fps=25,format=yuv420p"},
	}
	for _, tc := range tests {
		s := webReady()
		tc.change(s)
		o := std
		if tc.opts != nil {
			o = defaults(t, *tc.opts)
		}
		if got := videoFilters(s, o); got != tc.want {
			t.Errorf("%s:\n got %s\nwant %s", tc.name, got, tc.want)
		}
	}
}

func TestOutputSize(t *testing.T) {
	tests := []struct{ w, h, max, wantW, wantH int }{
		{1920, 1080, 1080, 1920, 1080},
		{3840, 2160, 1080, 1920, 1080},
		{2160, 3840, 1080, 1080, 1920},
		{1080, 1920, 720, 720, 1280},
		//odd sizes come out even
		{853, 480, 1080, 852, 480},
		{1441, 1081, 1080, 1440, 1080},
		{3840, 2160, -1, 3840, 2160},
	}
	for _, tc := range tests {
		if w, h := outputSize(tc.w, tc.h, tc.max); w != tc.wantW || h != tc.wantH {
			t.Errorf("outputSize(%d, %d, %d) = %dx%d, want %dx%d", tc.w, tc.h, tc.max, w, h, tc.wantW, tc.wantH)
		}
	}
}

func TestCheckOutput(t *testing.T) {
	src := webReady()
	src.CreationTime = time.Date(2024, 7, 14, 12, 4, 6, 0, time.UTC)
	good := *src

	if err := checkOutput(src, &good); err != nil {
		t.Errorf("identical output: %v", err)
	}
	for name, change := range map[string]func(s *Summary){
		"not h264":        func(s *Summary) { s.VideoCodec = "hevc" },
		"too short":       func(s *Summary) { s.Duration = 5 * time.Second },
		"turned sideways": func(s *Summary) { s.Width, s.Height = 1080, 1920 },
		"date lost":       func(s *Summary) { s.CreationTime = time.Time{} },
	} {
		out := good
		change(&out)
		if err := checkOutput(src, &out); !errors.Is(err, ErrOutputCheck) {
			t.Errorf("%s: expected ErrOutputCheck, got %v", name, err)
		}
	}
	//a small difference in length is normal after re-encoding
	out := good
	out.Duration += 100 * time.Millisecond
	if err := checkOutput(src, &out); err != nil {
		t.Errorf("100ms longer: %v", err)
	}
}

// --- integration, running ffmpeg ---

func transcode(t *testing.T, file string, opts TranscodeOptions) *TranscodeResult {
	t.Helper()
	dir := t.TempDir()
	r, err := Transcode(testContext(t), assetPath+file, filepath.Join(dir, "out"), opts)
	if err != nil {
		t.Fatalf("%s: %v", file, err)
	}
	if r.Path != filepath.Join(dir, "out.mp4") {
		t.Errorf("%s: wrote %s, want out.mp4", file, r.Path)
	}
	//web ready means h264 in 8 bit, with the index before the media
	if r.Output.VideoCodec != "h264" || r.Output.PixelFormat != "yuv420p" || r.Output.HDR || r.Output.Interlaced {
		t.Errorf("%s: output is %+v", file, r.Output)
	}
	if !moovFirst(t, r.Path) {
		t.Errorf("%s: the index is not at the front", file)
	}
	if info, err := os.Stat(r.Path); err != nil || info.Mode().Perm() != 0644 {
		t.Errorf("%s: output mode %v, %v, want 0644", file, info.Mode().Perm(), err)
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 1 {
		t.Errorf("%s: expected only the output in the directory, found %d entries", file, len(entries))
	}
	return r
}

// moovFirst reports whether the moov box comes before mdat
func moovFirst(t *testing.T, fileName string) bool {
	t.Helper()
	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("could not read %s: %v", fileName, err)
	}
	moov, mdat := strings.Index(string(data), "moov"), strings.Index(string(data), "mdat")
	return moov >= 0 && moov < mdat
}

func TestTranscode(t *testing.T) {
	requireFFmpeg(t)

	//already web ready: copied, tags and date kept
	r := transcode(t, "plain.mp4", TranscodeOptions{})
	if r.Encoded {
		t.Errorf("plain.mp4 should have been copied")
	}
	checkSummary(t, expectations[0], r.Output, nil)

	//a web ready file keeps its rotation flag, which browsers honour
	r = transcode(t, "rotated.mp4", TranscodeOptions{})
	if r.Encoded || r.Output.Rotation != 90 || r.Output.Width != 240 {
		t.Errorf("rotated.mp4: encoded %v, rotation %d, %dx%d", r.Encoded, r.Output.Rotation, r.Output.Width, r.Output.Height)
	}

	//a phone original: HEVC, rotated, index at the end, location off
	r = transcode(t, "samsung-like.mp4", TranscodeOptions{})
	if !r.Encoded || r.Output.Rotation != 0 || r.Output.Width != 240 || r.Output.Height != 320 {
		t.Errorf("samsung-like: encoded %v, rotation %d, %dx%d", r.Encoded, r.Output.Rotation, r.Output.Width, r.Output.Height)
	}
	if want := time.Date(2024, 7, 14, 12, 4, 6, 0, time.UTC); !r.Output.CreationTime.Equal(want) || r.Output.Location != nil {
		t.Errorf("samsung-like: created %v, location %v", r.Output.CreationTime, r.Output.Location)
	}
	if r.Output.AudioCodec != "aac" {
		t.Errorf("samsung-like: audio %q, want aac", r.Output.AudioCodec)
	}
	//the pixels were turned, red on top of the stored frame ends up right
	posterPath, err := ExtractPoster(testContext(t), r.Path, filepath.Join(t.TempDir(), "poster"), PosterOptions{})
	if err != nil {
		t.Fatalf("samsung-like poster: %v", err)
	}
	poster := decode(t, posterPath)
	if left, right := dominant(poster.At(40, 160)), dominant(poster.At(200, 160)); left != "blue" || right != "red" {
		t.Errorf("samsung-like: left %s right %s, want blue and red", left, right)
	}

	r = transcode(t, "interlaced.mp4", TranscodeOptions{})
	if !r.Encoded || r.Output.FrameRate != 10 {
		t.Errorf("interlaced: encoded %v at %v fps, want 10", r.Encoded, r.Output.FrameRate)
	}

	r = transcode(t, "anamorphic.mp4", TranscodeOptions{})
	if r.Output.Width != 320 || r.Output.Height != 240 || r.Output.PixelAspect != 1 {
		t.Errorf("anamorphic: %dx%d with pixel aspect %v", r.Output.Width, r.Output.Height, r.Output.PixelAspect)
	}

	//old avi: detected by content, sound re-encoded to aac
	if f, err := metadata.DetectFormatFile(assetPath + "old.avi"); err != nil || f != metadata.FormatAvi {
		t.Errorf("old.avi detected as %v, %v", f, err)
	}
	r = transcode(t, "old.avi", TranscodeOptions{})
	if r.Output.AudioCodec != "aac" || r.Output.Width != 160 {
		t.Errorf("old.avi: audio %q, %dx%d", r.Output.AudioCodec, r.Output.Width, r.Output.Height)
	}

	//options: forced encode, and the caps
	r = transcode(t, "plain.mp4", TranscodeOptions{ForceEncode: true, MaxShortSide: 120, MaxFrameRate: 5})
	if !r.Encoded || r.Output.Width != 160 || r.Output.Height != 120 || r.Output.FrameRate != 5 {
		t.Errorf("capped: encoded %v, %dx%d at %v fps", r.Encoded, r.Output.Width, r.Output.Height, r.Output.FrameRate)
	}
	//and the metadata survives an encode as well as a copy
	if r.Output.CameraModel != "iPhone 15 Pro" || !r.Output.CreationTime.Equal(r.Source.CreationTime) {
		t.Errorf("capped: model %q, created %v", r.Output.CameraModel, r.Output.CreationTime)
	}
}

func TestTranscodeRejects(t *testing.T) {
	requireFFmpeg(t)
	ctx := testContext(t)
	dir := t.TempDir()

	if _, err := Transcode(ctx, assetPath+"plain.mp4", filepath.Join(dir, "a.mp4"), TranscodeOptions{}); !errors.Is(err, img.ErrDestHasExtension) {
		t.Errorf("destination with extension: expected ErrDestHasExtension, got %v", err)
	}
	if _, err := Transcode(ctx, assetPath+"audioonly.mp4", filepath.Join(dir, "b"), TranscodeOptions{}); !errors.Is(err, ErrNoVideoStream) {
		t.Errorf("audio only: expected ErrNoVideoStream, got %v", err)
	}
	//10 bit HLG, as an iPhone records by default
	if _, err := Transcode(ctx, assetPath+"hdr.mp4", filepath.Join(dir, "h"), TranscodeOptions{}); !errors.Is(err, ErrHDRUnsupported) {
		t.Errorf("hdr: expected ErrHDRUnsupported, got %v", err)
	}
	if _, err := Transcode(ctx, "../assets/leica.jpg", filepath.Join(dir, "c"), TranscodeOptions{}); !errors.Is(err, ErrNotVideo) {
		t.Errorf("an image: expected ErrNotVideo, got %v", err)
	}
	if _, err := Transcode(ctx, assetPath+"plain.mp4", filepath.Join(dir, "d"), TranscodeOptions{Quality: Quality(7)}); err == nil {
		t.Errorf("unknown quality: expected an error")
	}

	plain, err := os.ReadFile(assetPath + "plain.mp4")
	if err != nil {
		t.Fatalf("could not read fixture: %v", err)
	}
	cut := filepath.Join(t.TempDir(), "cut.mp4")
	if err = os.WriteFile(cut, plain[:len(plain)/2], 0644); err != nil {
		t.Fatalf("could not write: %v", err)
	}
	if _, err = Transcode(ctx, cut, filepath.Join(dir, "e"), TranscodeOptions{}); !errors.Is(err, ErrTruncated) {
		t.Errorf("truncated: expected ErrTruncated, got %v", err)
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 0 {
		t.Errorf("failed transcodes should leave nothing behind, found %d entries", len(entries))
	}
}
