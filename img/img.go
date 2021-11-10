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

type DestImage struct {
	Width     int
	Height    int
	Quality   int
	Anchor    CropAnchor
	Transform TransformType
	Strategy  ResampleStrategy
	CopyExif  bool
	FileName  string
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

func hasCopyExif(dest []DestImage) bool {
	for _, d := range dest {
		if d.CopyExif {
			return true
		}
	}
	return false
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

func saveWithExif(srcBytes []byte, dstImage image.Image, opt DestImage) error {
	dstBytes := new(bytes.Buffer)
	err := imaging.Encode(dstBytes, dstImage, imaging.JPEG, imaging.JPEGQuality(opt.Quality))
	if err != nil {
		return err
	}
	mde, err := metadata.NewMetaDataEditor(dstBytes.Bytes())
	if err != nil {
		return err
	}
	err = mde.CopyMetaData(srcBytes, metadata.CopyAll)
	if err != nil {
		return err
	}
	err = mde.SetExifDate(metadata.ModifyDate, time.Now())
	if err != nil {
		return err
	}
	return mde.WriteFile(opt.FileName)
	//return nil
	/*
		if out, err := metadata.CopyExif(srcBytes,dstBytes.Bytes(),true); err != nil {
			return err
		} else {
			err = ioutil.WriteFile(opt.FileName, out, 0644)
			if err != nil {
				return err
			}
		}
		return nil*/
}

func TransformFile(source string, dest ...DestImage) error {
	var srcImg image.Image
	var err error
	var srcBytes []byte
	if hasCopyExif(dest) {
		srcImg, srcBytes, err = openForExifCopy(source)
	} else {
		srcImg, err = imaging.Open(source)
	}
	if err != nil {
		return err
	}

	for _, d := range dest {
		dst := transform(srcImg, d)
		if d.CopyExif {
			err = saveWithExif(srcBytes, dst, d)
		} else {
			err = imaging.Save(dst, d.FileName, imaging.JPEGQuality(d.Quality))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func transform(src image.Image, dest DestImage) image.Image {
	var dstImage image.Image

	switch dest.Transform {
	case ResizeAndCrop:
		dstImage = imaging.Fill(src, dest.Width, dest.Height, anchor(dest.Anchor), resampleFiler(dest.Strategy))
	case Crop:
		dstImage = imaging.CropAnchor(src, dest.Width, dest.Height, anchor(dest.Anchor))
	case Resize:
		dstImage = imaging.Resize(src, dest.Width, dest.Height, resampleFiler(dest.Strategy))
	case ResizeAndFit:
		dstImage = imaging.Fit(src, dest.Width, dest.Height, resampleFiler(dest.Strategy))
	}
	return dstImage
}
