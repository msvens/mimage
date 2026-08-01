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
		dest := path.Join(t.TempDir(), "out.jpg")
		opts := NewOptions(Resize, 400, 0, true)
		if err = TransformFile(upperSrc, map[string]Options{dest: opts}); err != nil {
			t.Fatalf("Could not transform %s: %v", ext, err)
		}
		md, err := metadata.NewMetaDataFromFile(dest)
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
	thumb := NewOptions(ResizeAndCrop, 400, 400, false)
	landscape := NewOptions(ResizeAndCrop, 1200, 628, true)
	square := NewOptions(ResizeAndCrop, 1200, 1200, true)
	portrait := NewOptions(ResizeAndCrop, 1080, 1350, true)
	resize := NewOptions(Resize, 1200, 0, true)

	destImgs := map[string]Options{
		path.Join(sourceDir, "thumb.jpg"):     thumb,
		path.Join(sourceDir, "landscape.jpg"): landscape,
		path.Join(sourceDir, "square.jpg"):    square,
		path.Join(sourceDir, "portrait.jpg"):  portrait,
		path.Join(sourceDir, "resize.jpg"):    resize,
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
			opts := NewOptions(Resize, 400, 0, true)
			opts.MakerNote = tc.policy
			dest := path.Join(t.TempDir(), "out.jpg")
			if err := TransformFile(canonImg, map[string]Options{dest: opts}); err != nil {
				t.Fatalf("could not transform: %v", err)
			}
			md, err := metadata.NewMetaDataFromFile(dest)
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
	opts := NewOptions(Resize, 400, 0, true)
	opts.MakerNote = MakerNoteFail
	dest := path.Join(t.TempDir(), "out.jpg")
	err := TransformFile(canonImg, map[string]Options{dest: opts})
	if !errors.Is(err, metadata.ErrMakerNotePresent) {
		t.Errorf("expected ErrMakerNotePresent, got %v", err)
	}
}
