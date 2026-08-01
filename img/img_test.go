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
