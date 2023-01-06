package img

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/msvens/mimage/metadata"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"time"
)

// CropAnchor specifies where the anchor point should be when cropping an image
type CropAnchor int

// Various Anchors for your crop
const (
	Center CropAnchor = iota
	TopLeft
	Top
	TopRight
	Left
	Right
	BottomLeft
	Bottom
	BottomRight
)

// ResampleStrategy specifies what method to use when resampling an image
type ResampleStrategy int

const (
	//Lanczos sampling. Default
	Lanczos ResampleStrategy = iota
	//NearestNeighbor sampling
	NearestNeighbor
	//Box sampling
	Box
	//Linear sampling
	Linear
	//Hermite sampling
	Hermite
	//MitchellNetravali sampling
	MitchellNetravali
	//CatmullRom sampling
	CatmullRom
	//BSpline sampling
	BSpline
	//Gaussian sampling
	Gaussian
	//Bartlett sampling
	Bartlett
	//Hann sampling
	Hann
	//Hamming sampling
	Hamming
	//Blackman sampling
	Blackman
	//Welch sampling
	Welch
	//Cosine sampling
	Cosine
)

// TransformType specifies how to transform an image, crop, resize, etc...
type TransformType int

const (
	//ResizeAndCrop scales the image to given dimension. To keep aspect ratio the image is cropped
	ResizeAndCrop TransformType = iota
	//Crop cuts a specified region of the image
	Crop
	//Resize scale the image to the specified dimensions. To keep aspect ratio set either width or height to 0
	Resize
	//ResizeAndFit scale the image to fit the maximum specified dimensions.
	ResizeAndFit
)

// Options holds all options for a given image transformation job
type Options struct {
	Width     int
	Height    int
	Quality   int
	Anchor    CropAnchor
	Transform TransformType
	Strategy  ResampleStrategy
	X         int
	Y         int
	Angle     int
	CopyExif  bool
}

func resampleFiler(strategy ResampleStrategy) imaging.ResampleFilter {
	switch strategy {
	case Lanczos:
		return imaging.Lanczos
	case NearestNeighbor:
		return imaging.NearestNeighbor
	case Box:
		return imaging.Box
	case Linear:
		return imaging.Linear
	case Hermite:
		return imaging.Hermite
	case MitchellNetravali:
		return imaging.MitchellNetravali
	case CatmullRom:
		return imaging.CatmullRom
	case BSpline:
		return imaging.BSpline
	case Gaussian:
		return imaging.Gaussian
	case Bartlett:
		return imaging.Bartlett
	case Hann:
		return imaging.Hann
	case Hamming:
		return imaging.Hamming
	case Blackman:
		return imaging.Blackman
	case Welch:
		return imaging.Welch
	case Cosine:
		return imaging.Cosine
	default:
		return imaging.Lanczos
	}
}

func anchor(ca CropAnchor) imaging.Anchor {
	switch ca {
	case Center:
		return imaging.Center
	case TopLeft:
		return imaging.TopLeft
	case Top:
		return imaging.Top
	case TopRight:
		return imaging.TopRight
	case Left:
		return imaging.Left
	case Right:
		return imaging.Right
	case BottomLeft:
		return imaging.BottomLeft
	case Bottom:
		return imaging.Bottom
	case BottomRight:
		return imaging.BottomRight
	default:
		return imaging.Center
	}
}

// NewOptions creates Options that default to crop center strategy, image quality 90% and Resize strategy lanczos
func NewOptions(transform TransformType, width, height int, copyExif bool) Options {
	return Options{Width: width, Height: height, Quality: 90, Anchor: Center,
		Transform: transform, Strategy: Lanczos, CopyExif: copyExif}
}

func Open(fileName string) (image.Image, error) {
	img, _, err := OpenOpts(fileName, false, false)
	return img, err
}

func OpenOpts(fileName string, autoOrientation bool, srcBytes bool) (image.Image, []byte, error) {
	if !srcBytes {
		img, err := imaging.Open(fileName, imaging.AutoOrientation(autoOrientation))
		return img, nil, err
	}
	src, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}
	img, err := imaging.Decode(bytes.NewReader(src), imaging.AutoOrientation(autoOrientation))
	return img, src, err
}

func openForExifCopy(sourceFile string) (image.Image, []byte, error) {
	srcBytes, err := os.ReadFile(sourceFile)
	if err != nil {
		return nil, nil, err
	}
	srcImg, err := jpeg.Decode(bytes.NewReader(srcBytes))
	if err != nil {
		return nil, nil, err
	}
	return srcImg, srcBytes, nil
}

func saveWithExif(srcBytes []byte, dstImage image.Image, opt Options, fileName string) error {
	dstBytes := new(bytes.Buffer)
	err := imaging.Encode(dstBytes, dstImage, imaging.JPEG, imaging.JPEGQuality(opt.Quality))
	if err != nil {
		return err
	}
	mde, err := metadata.NewJpegEditor(dstBytes.Bytes())
	if err != nil {
		return err
	}
	err = mde.CopyMetaData(srcBytes)
	if err != nil {
		return err
	}
	err = mde.Exif().SetDate(metadata.ModifyDate, time.Now())
	if err != nil {
		return err
	}
	return mde.WriteFile(fileName)
}

func Save(image image.Image, fileName string) error {
	return SaveOpts(image, fileName, 90, nil)
}

func SaveOpts(image image.Image, fileName string, quality int, srcExif []byte) error {
	if quality < 1 || quality > 100 {
		quality = 90
	}
	if srcExif == nil {
		return imaging.Save(image, fileName, imaging.JPEGQuality(quality))
	}
	//we will add exif information
	dstBytes := new(bytes.Buffer)
	err := imaging.Encode(dstBytes, image, imaging.JPEG, imaging.JPEGQuality(quality))
	if err != nil {
		return err
	}
	mde, err := metadata.NewJpegEditor(dstBytes.Bytes())
	if err != nil {
		return err
	}
	err = mde.CopyMetaData(srcExif)
	if err != nil {
		return err
	}
	err = mde.Exif().SetDate(metadata.ModifyDate, time.Now())
	if err != nil {
		return err
	}
	return mde.WriteFile(fileName)
}

// negative angle indicates a counter clockwise rotation
func CropImage(img image.Image, crop image.Rectangle) image.Image {
	if crop.Empty() || !crop.In(img.Bounds()) {
		return img
	}
	return imaging.Crop(img, crop)
}

func RotateImage(img image.Image, angle int) image.Image {
	//fix angle first:
	angle = angle % 360
	if angle == 0 {
		return img
	}
	if angle < 0 {
		angle = -angle
	} else {
		angle = 360 - angle
	}
	return imaging.Rotate(img, float64(angle), color.Black)
}

func RotateAndCropFile(source string, dest string, opts Options) error {

	angle := opts.Angle
	crop := opts.Rectangle()
	if angle == 0 && crop.Empty() {
		return fmt.Errorf("Neither angle or crop was provided")
	}

	srcImg, srcBytes, err := OpenOpts(source, false, opts.CopyExif)
	if err != nil {
		return err
	}

	srcImg = RotateImage(srcImg, angle)
	srcImg = CropImage(srcImg, crop)

	return SaveOpts(srcImg, dest, opts.Quality, srcBytes)

	/*	angle := opts.Angle
		crop := opts.Rectangle()
		if angle == 0 && crop.Empty() {
			return fmt.Errorf("Neither angle or crop was provided")
		}

		srcImg, srcBytes, err := openForExifCopy(source)
		if err != nil {
			return err
		}

		var destImg *image.NRGBA
		//We should add tests for nil here
		if angle != 0 && !crop.Empty() {
			destImg = imaging.Rotate(srcImg, angle, color.Black)
			destImg = imaging.Crop(destImg, crop)
		} else if angle != 0 {
			destImg = imaging.Rotate(srcImg, angle, color.Black)
		} else {
			destImg = imaging.Crop(srcImg, crop)
		}
		if opts.CopyExif {
			err = saveWithExif(srcBytes, destImg, opts, dest)
		} else {
			err = imaging.Save(destImg, dest, imaging.JPEGQuality(opts.Quality))
		}
		if err != nil {
			return err
		}
		return nil*/
}

// TransformFile creates versions of source based on destinations
func TransformFile(source string, destinations map[string]Options) error {
	srcImg, srcBytes, err := openForExifCopy(source)
	if err != nil {
		return err
	}
	for dest, options := range destinations {
		destImg := transform(srcImg, options)
		if options.CopyExif {
			err = saveWithExif(srcBytes, destImg, options, dest)
		} else {
			err = imaging.Save(destImg, dest, imaging.JPEGQuality(options.Quality))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func transform(src image.Image, opt Options) image.Image {
	var dstImage image.Image

	a := anchor(opt.Anchor)
	rf := resampleFiler(opt.Strategy)
	switch opt.Transform {
	case ResizeAndCrop:
		dstImage = imaging.Fill(src, opt.Width, opt.Height, a, rf)
	case Crop:
		dstImage = imaging.CropAnchor(src, opt.Width, opt.Height, a)
	case Resize:
		dstImage = imaging.Resize(src, opt.Width, opt.Height, rf)
	case ResizeAndFit:
		dstImage = imaging.Fit(src, opt.Width, opt.Height, rf)
	}
	return dstImage
}
