package img

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/msvens/mimage/metadata"
	"image"
	"image/jpeg"
	"io/ioutil"
	"time"
)

type CropAnchor int

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

type ResampleStrategy int

const (
	Lanczos ResampleStrategy = iota
	NearestNeighbor
	Box
	Linear
	Hermite
	MitchellNetravali
	CatmullRom
	BSpline
	Gaussian
	Bartlett
	Hann
	Hamming
	Blackman
	Welch
	Cosine
)

type TransformType int

const (
	//Scales the image to given dimension. To keep aspect ratio the image is cropped
	ResizeAndCrop TransformType = iota
	//Cuts a specified region of the image
	Crop
	//Scale the image to the specified dimensions. To keep aspect ratio set either width or height to 0
	Resize
	//Scale the image to fit the maximum specified dimensions.
	ResizeAndFit
)

type Options struct {
	Width     int
	Height    int
	Quality   int
	Anchor    CropAnchor
	Transform TransformType
	Strategy  ResampleStrategy
	CopyExif  bool
	//FileName  string
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

//Create Options that default to crop center strategy, image quality 90% and Resize strategy lanczos
func NewOptions(transform TransformType, width, height int, copyExif bool) Options {
	return Options{Width: width, Height: height, Quality: 90, Anchor: Center,
		Transform: transform, Strategy: Lanczos, CopyExif: copyExif}
}

func openForExifCopy(sourceFile string) (image.Image, []byte, error) {
	srcBytes, err := ioutil.ReadFile(sourceFile)
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

/*
func Transform(source string, dest string, opt Options) error {
	var srcImg image.Image
	var err error
	var srcBytes []byte
	if opt.CopyExif {
		srcImg, srcBytes, err = openForExifCopy(source)
	} else {
		srcImg, err = imaging.Open(source)
	}
	if err != nil {
		return err
	}
	dst := transform(srcImg, opt)
	if opt.CopyExif {
		err = saveWithExif(srcBytes, dst, opt, dest)
	} else {
		err = imaging.Save(dst, dest, imaging.JPEGQuality(opt.Quality))
	}
	return err
}
*/

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
