package img

import (
	"fmt"
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
