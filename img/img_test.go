package img

import (
	"errors"
	"fmt"
	"github.com/msvens/mimage/metadata"
	"os"
	"path"
	"testing"
)

func TestNewOptions(t *testing.T) {

}

func TestTransformFile(t *testing.T) {

}

// Cameras commonly emit .JPG. Those sources must still be recognised as jpeg
// so that CopyExif carries the metadata across, see isJpegFile
func TestTransformFileUppercaseSource(t *testing.T) {
	src, err := os.ReadFile("../assets/leica.jpg")
	if err != nil {
		t.Fatalf("Could not read source image: %v", err)
	}
	for _, ext := range []string{".JPG", ".JPEG", ".jpg"} {
		upperSrc := path.Join(t.TempDir(), "leica"+ext)
		if err = os.WriteFile(upperSrc, src, 0644); err != nil {
			t.Fatalf("Could not write %s: %v", ext, err)
		}
		dest := path.Join(t.TempDir(), "out")
		opts := NewOptions(Resize, 400, 0, true, FormatJpeg)
		if err = TransformFile(upperSrc, map[string]Options{dest: opts}); err != nil {
			t.Fatalf("Could not transform %s: %v", ext, err)
		}
		md, err := metadata.NewMetaDataFromFile(dest + FormatJpeg.Extension())
		if err != nil {
			t.Fatalf("Could not read metadata from %s output: %v", ext, err)
		}
		if make := md.Summary().CameraMake; make == "" {
			t.Errorf("%s source: expected exif to be copied, got empty CameraMake", ext)
		}
	}
}

func ExampleTransformFile() {
	sourceImg := "../assets/leica.jpg"
	homeDir, _ := os.UserHomeDir()
	sourceDir := path.Join(homeDir, "transform")
	_ = os.Mkdir(sourceDir, 0755)

	//for all but the thumb we are copying the original meta information
	thumb := NewOptions(ResizeAndCrop, 400, 400, false, FormatJpeg)
	landscape := NewOptions(ResizeAndCrop, 1200, 628, true, FormatJpeg)
	square := NewOptions(ResizeAndCrop, 1200, 1200, true, FormatJpeg)
	portrait := NewOptions(ResizeAndCrop, 1080, 1350, true, FormatJpeg)
	resize := NewOptions(Resize, 1200, 0, true, FormatJpeg)

	destImgs := map[string]Options{
		path.Join(sourceDir, "thumb"):     thumb,
		path.Join(sourceDir, "landscape"): landscape,
		path.Join(sourceDir, "square"):    square,
		path.Join(sourceDir, "portrait"):  portrait,
		path.Join(sourceDir, "resize"):    resize,
	}

	_ = TransformFile(sourceImg, destImgs)
	fmt.Println("Transformed leica.jpg")
	//Output: Transformed leica.jpg
}

// canon.jpg is the only asset carrying an exif MakerNote
const canonImg = "../assets/canon.jpg"

func TestTransformFileMakerNotePolicy(t *testing.T) {
	tests := []struct {
		name   string
		policy MakerNotePolicy
		want   bool
	}{
		{"preserve is the default", MakerNotePreserve, true},
		{"strip removes it", MakerNoteStrip, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := NewOptions(Resize, 400, 0, true, FormatJpeg)
			opts.MakerNote = tc.policy
			dest := path.Join(t.TempDir(), "out")
			if err := TransformFile(canonImg, map[string]Options{dest: opts}); err != nil {
				t.Fatalf("could not transform: %v", err)
			}
			md, err := metadata.NewMetaDataFromFile(dest + FormatJpeg.Extension())
			if err != nil {
				t.Fatalf("could not read metadata: %v", err)
			}
			if got := md.HasMakerNote(); got != tc.want {
				t.Errorf("HasMakerNote() = %v, want %v", got, tc.want)
			}
			//exif must survive either way
			if make := md.Summary().CameraMake; make == "" {
				t.Errorf("expected exif to be copied, got empty CameraMake")
			}
		})
	}
}

func TestTransformFileMakerNoteFail(t *testing.T) {
	opts := NewOptions(Resize, 400, 0, true, FormatJpeg)
	opts.MakerNote = MakerNoteFail
	dest := path.Join(t.TempDir(), "out")
	err := TransformFile(canonImg, map[string]Options{dest: opts})
	if !errors.Is(err, metadata.ErrMakerNotePresent) {
		t.Errorf("expected ErrMakerNotePresent, got %v", err)
	}
}

const tiffImg = "../assets/leica.tiff"

// A tiff source carries its metadata into a generated jpeg, which it silently
// failed to do before: CopyExif was only honoured for jpeg sources
func TestTransformFileFromTiff(t *testing.T) {
	dest := path.Join(t.TempDir(), "out")
	opts := NewOptions(Resize, 300, 0, true, FormatJpeg)
	if err := TransformFile(tiffImg, map[string]Options{dest: opts}); err != nil {
		t.Fatalf("could not transform tiff: %v", err)
	}
	out, err := metadata.NewMetaDataFromFile(dest + FormatJpeg.Extension())
	if err != nil {
		t.Fatalf("could not read output metadata: %v", err)
	}
	src, err := metadata.NewMetaDataFromFile(tiffImg)
	if err != nil {
		t.Fatalf("could not read source metadata: %v", err)
	}
	if got, want := out.Summary().CameraMake, src.Summary().CameraMake; got != want {
		t.Errorf("CameraMake = %q, want %q", got, want)
	}
	if got, want := out.Summary().ISO, src.Summary().ISO; got != want {
		t.Errorf("ISO = %d, want %d", got, want)
	}
	if got, want := out.Summary().Title, src.Summary().Title; got != want {
		t.Errorf("Title = %q, want %q", got, want)
	}
	//the output is the resized image, not the source
	if out.ImageWidth != 300 {
		t.Errorf("width = %d, want 300", out.ImageWidth)
	}
}

// Without CopyExif a tiff source must still produce a jpeg, just without the
// metadata
func TestTransformFileFromTiffNoCopyExif(t *testing.T) {
	dest := path.Join(t.TempDir(), "out")
	opts := NewOptions(Resize, 300, 0, false, FormatJpeg)
	if err := TransformFile(tiffImg, map[string]Options{dest: opts}); err != nil {
		t.Fatalf("could not transform tiff: %v", err)
	}
	md, err := metadata.NewMetaDataFromFile(dest + FormatJpeg.Extension())
	if err != nil {
		t.Fatalf("could not read output: %v", err)
	}
	if make := md.Summary().CameraMake; make != "" {
		t.Errorf("expected no exif without CopyExif, got CameraMake=%q", make)
	}
}

// Source formats that cannot carry metadata are still written, silently and
// without error, see the CopyExif field comment
func TestTransformFilePngSourceCopyExif(t *testing.T) {
	dest := path.Join(t.TempDir(), "out")
	opts := NewOptions(Resize, 200, 0, true, FormatJpeg)
	if err := TransformFile("../assets/leica.png", map[string]Options{dest: opts}); err != nil {
		t.Fatalf("a png source with CopyExif should still write: %v", err)
	}
	if _, err := os.Stat(dest + FormatJpeg.Extension()); err != nil {
		t.Errorf("expected an output file: %v", err)
	}
}

const pngImg = "../assets/leica.png"

// tiff -> jpeg at native size, metadata intact, no resize involved
func TestConvertFile(t *testing.T) {
	src, err := metadata.NewMetaDataFromFile(tiffImg)
	if err != nil {
		t.Fatalf("could not read source: %v", err)
	}
	base := path.Join(t.TempDir(), "converted")
	written, converted, err := ConvertFile(tiffImg, base, FormatJpeg, NewConvertOptions(true))
	if err != nil {
		t.Fatalf("could not convert: %v", err)
	}
	if !converted {
		t.Fatalf("expected a conversion to happen")
	}
	if written != base+".jpg" {
		t.Errorf("written = %q, want %q", written, base+".jpg")
	}
	out, err := metadata.NewMetaDataFromFile(written)
	if err != nil {
		t.Fatalf("could not read output: %v", err)
	}
	//dimensions are preserved, this is a convert and not a resize
	if out.ImageWidth != src.ImageWidth || out.ImageHeight != src.ImageHeight {
		t.Errorf("dimensions = %dx%d, want %dx%d",
			out.ImageWidth, out.ImageHeight, src.ImageWidth, src.ImageHeight)
	}
	for _, tc := range []struct{ name, got, want string }{
		{"CameraMake", out.Summary().CameraMake, src.Summary().CameraMake},
		{"CameraModel", out.Summary().CameraModel, src.Summary().CameraModel},
		{"Title", out.Summary().Title, src.Summary().Title},
	} {
		if tc.got != tc.want {
			t.Errorf("%s = %q, want %q", tc.name, tc.got, tc.want)
		}
	}
	if out.Summary().ISO != src.Summary().ISO {
		t.Errorf("ISO = %d, want %d", out.Summary().ISO, src.Summary().ISO)
	}
}

// A base name that already carries an image extension is rejected rather than
// silently producing out.jpg.jpg
func TestConvertFileDestExtension(t *testing.T) {
	dir := t.TempDir()
	rejected := []string{"out.jpg", "out.JPG", "out.tiff", "out.heic"}
	for _, name := range rejected {
		_, _, err := ConvertFile(tiffImg, path.Join(dir, name), FormatJpeg, NewConvertOptions(false))
		if !errors.Is(err, ErrDestHasExtension) {
			t.Errorf("%s: expected ErrDestHasExtension, got %v", name, err)
		}
	}
	//an ordinary dotted name is not an image extension and must be accepted
	for _, name := range []string{"backup.v2", "2020-10-27_09.34.03"} {
		written, _, err := ConvertFile(tiffImg, path.Join(dir, name), FormatJpeg, NewConvertOptions(false))
		if err != nil {
			t.Errorf("%s: unexpected error %v", name, err)
			continue
		}
		if written != path.Join(dir, name)+".jpg" {
			t.Errorf("%s: written = %q", name, written)
		}
	}
}

// Converting to the format the source already is writes nothing
func TestConvertFileAlreadyTargetFormat(t *testing.T) {
	base := path.Join(t.TempDir(), "same")
	written, converted, err := ConvertFile("../assets/leica.jpg", base, FormatJpeg, NewConvertOptions(true))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if converted || written != "" {
		t.Errorf("expected no conversion, got (%q, %v)", written, converted)
	}
	if _, err := os.Stat(base + ".jpg"); err == nil {
		t.Errorf("nothing should have been written")
	}
}

// The source format comes from content. A tiff named .jpg must still convert
func TestConvertFileDetectsSourceByContent(t *testing.T) {
	raw, err := os.ReadFile(tiffImg)
	if err != nil {
		t.Fatalf("could not read fixture: %v", err)
	}
	liar := path.Join(t.TempDir(), "liar.jpg")
	if err = os.WriteFile(liar, raw, 0644); err != nil {
		t.Fatalf("could not write: %v", err)
	}
	written, converted, err := ConvertFile(liar, path.Join(t.TempDir(), "out"), FormatJpeg, NewConvertOptions(true))
	if err != nil {
		t.Fatalf("a tiff named .jpg should still convert: %v", err)
	}
	if !converted {
		t.Errorf("expected a conversion, the source is a tiff by content")
	}
	md, err := metadata.NewMetaDataFromFile(written)
	if err != nil {
		t.Fatalf("could not read output: %v", err)
	}
	if md.Summary().CameraMake == "" {
		t.Errorf("expected metadata to be carried across")
	}
}

// Format problems are distinguishable from everything else
func TestConvertFileUnsupported(t *testing.T) {
	dir := t.TempDir()
	//a format mimage cannot write
	if _, _, err := ConvertFile(tiffImg, path.Join(dir, "a"), FormatUnknown, NewConvertOptions(false)); !errors.Is(err, ErrUnsupportedFormat) {
		t.Errorf("FormatUnknown: expected ErrUnsupportedFormat, got %v", err)
	}
	//a source that is not an image at all
	if _, _, err := ConvertFile("../assets/xmp.xml", path.Join(dir, "b"), FormatJpeg, NewConvertOptions(false)); !errors.Is(err, ErrUnsupportedFormat) {
		t.Errorf("non image source: expected ErrUnsupportedFormat, got %v", err)
	}
	//a missing file is an io problem, not a format problem
	_, _, err := ConvertFile(path.Join(dir, "nope.jpg"), path.Join(dir, "c"), FormatJpeg, NewConvertOptions(false))
	if err == nil {
		t.Errorf("expected an error for a missing source")
	} else if errors.Is(err, ErrUnsupportedFormat) {
		t.Errorf("a missing file should not report as a format problem: %v", err)
	}
}

// png cannot carry metadata, so CopyExif is ignored and the convert still works
func TestConvertFilePngSource(t *testing.T) {
	written, converted, err := ConvertFile(pngImg, path.Join(t.TempDir(), "out"), FormatJpeg, NewConvertOptions(true))
	if err != nil {
		t.Fatalf("could not convert png: %v", err)
	}
	if !converted {
		t.Errorf("expected a conversion")
	}
	md, err := metadata.NewMetaDataFromFile(written)
	if err != nil {
		t.Fatalf("could not read output: %v", err)
	}
	if make := md.Summary().CameraMake; make != "" {
		t.Errorf("a png source carries no exif, got CameraMake=%q", make)
	}
}

// TransformFile now requires a format and rejects destination keys that carry
// an extension
func TestTransformFileFormatRequired(t *testing.T) {
	dir := t.TempDir()
	noFormat := Options{Width: 200, Transform: Resize, Quality: 90}
	if err := TransformFile(tiffImg, map[string]Options{path.Join(dir, "a"): noFormat}); !errors.Is(err, ErrUnsupportedFormat) {
		t.Errorf("expected ErrUnsupportedFormat for a missing format, got %v", err)
	}
	withExt := NewOptions(Resize, 200, 0, false, FormatJpeg)
	if err := TransformFile(tiffImg, map[string]Options{path.Join(dir, "b.jpg"): withExt}); !errors.Is(err, ErrDestHasExtension) {
		t.Errorf("expected ErrDestHasExtension, got %v", err)
	}
	//nothing should have been written by either failure
	if entries, _ := os.ReadDir(dir); len(entries) != 0 {
		t.Errorf("expected no output files, found %d", len(entries))
	}
}

// The output name comes from Options.Format, not from the destination key
func TestTransformFileNamesFromFormat(t *testing.T) {
	dir := t.TempDir()
	jpg := NewOptions(Resize, 200, 0, true, FormatJpeg)
	png := NewOptions(Resize, 200, 0, false, FormatPng)
	err := TransformFile(tiffImg, map[string]Options{
		path.Join(dir, "asjpeg"): jpg,
		path.Join(dir, "aspng"):  png,
	})
	if err != nil {
		t.Fatalf("could not transform: %v", err)
	}
	for name, want := range map[string]Format{"asjpeg.jpg": FormatJpeg, "aspng.png": FormatPng} {
		got, e := metadata.DetectFormatFile(path.Join(dir, name))
		if e != nil {
			t.Errorf("%s: %v", name, e)
			continue
		}
		if got != want {
			t.Errorf("%s: detected %v, want %v", name, got, want)
		}
	}
}
