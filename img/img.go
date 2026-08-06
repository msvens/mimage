package img

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/msvens/mimage/metadata"
	"image"
	"image/color"
	"os"
	"path"
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

// MakerNotePolicy controls what happens to an exif MakerNote when metadata is
// copied to a transformed image. See metadata.MakerNotePolicy
type MakerNotePolicy = metadata.MakerNotePolicy

// MakerNote policies, aliased from the metadata package
const (
	//MakerNotePreserve copies the MakerNote through to the new image. Default
	MakerNotePreserve = metadata.MakerNotePreserve
	//MakerNoteStrip removes the MakerNote from the new image
	MakerNoteStrip = metadata.MakerNoteStrip
	//MakerNoteFail aborts with metadata.ErrMakerNotePresent if a MakerNote is present
	MakerNoteFail = metadata.MakerNoteFail
)

// Format identifies an image format. See metadata.Format
type Format = metadata.Format

// Image formats, aliased from the metadata package
const (
	//FormatUnknown is anything mimage does not handle
	FormatUnknown = metadata.FormatUnknown
	FormatJpeg    = metadata.FormatJpeg
	FormatPng     = metadata.FormatPng
	FormatGif     = metadata.FormatGif
	FormatTiff    = metadata.FormatTiff
	FormatBmp     = metadata.FormatBmp
)

// Image errors
var (
	//ErrUnsupportedFormat is returned when a source cannot be decoded as a
	//known image format, or a requested output format is one mimage cannot
	//write. Check with errors.Is
	ErrUnsupportedFormat = errors.New("mimage: unsupported image format")
	//ErrDestHasExtension is returned when a destination base name already
	//carries an image extension. mimage appends the extension for the chosen
	//format, so passing "out.jpg" would produce "out.jpg.jpg". Check with
	//errors.Is
	ErrDestHasExtension = errors.New("mimage: destination already has an image extension")
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
	//CopyExif carries the source metadata into the destination. Only honoured
	//when the destination is a jpeg and the source is a jpeg or a tiff. For
	//other source formats it is ignored
	CopyExif bool
	//MakerNote controls what happens to an exif MakerNote when CopyExif is set.
	//The zero value preserves it, matching earlier releases
	MakerNote MakerNotePolicy
	//Format of the output. Required: a destination whose Format is
	//FormatUnknown, or a format mimage cannot write, is an error
	Format Format
}

// ConvertOptions holds the options for ConvertFile
type ConvertOptions struct {
	//Quality is the jpeg quality, 1-100. Anything outside that range means 90
	Quality int
	//CopyExif carries the source metadata across. Honoured only when the
	//destination is a jpeg and the source is a jpeg or a tiff
	CopyExif bool
	//MakerNote controls what happens to an exif MakerNote when CopyExif is set
	MakerNote MakerNotePolicy
}

// NewConvertOptions creates ConvertOptions defaulting to image quality 90
func NewConvertOptions(copyExif bool) ConvertOptions {
	return ConvertOptions{Quality: 90, CopyExif: copyExif}
}

// destPath appends the canonical extension for format to destBase. A base that
// already carries an image extension is rejected rather than silently fixed,
// since "out.jpg" would otherwise become "out.jpg.jpg"
func destPath(destBase string, format Format) (string, error) {
	ext := format.Extension()
	if ext == "" {
		return "", fmt.Errorf("%w: cannot write %v", ErrUnsupportedFormat, format)
	}
	if metadata.IsImageExtension(path.Ext(destBase)) {
		return "", fmt.Errorf("%w: %s, pass a name without one", ErrDestHasExtension, destBase)
	}
	return destBase + ext, nil
}

// ConvertFile writes source to destBase in the given format, keeping the
// original dimensions, and returns the path it actually wrote. The canonical
// extension for format is appended to destBase, which must not already carry an
// image extension.
//
// The source format is detected from its content, not its filename. Reports
// false when the source is already in format: nothing is written and the caller
// should keep using the original file.
//
// An animated gif converts to its first frame only, the rest are discarded, and
// transparency is lost when the destination is a jpeg
func ConvertFile(source, destBase string, format Format, opts ConvertOptions) (string, bool, error) {
	dest, err := destPath(destBase, format)
	if err != nil {
		return "", false, err
	}

	srcFormat, err := metadata.DetectFormatFile(source)
	if err != nil {
		return "", false, err
	}
	if !srcFormat.Supported() {
		return "", false, fmt.Errorf("%w: cannot read %v", ErrUnsupportedFormat, srcFormat)
	}
	if srcFormat == format {
		return "", false, nil
	}

	srcImg, srcBytes, err := OpenOpts(source, false, opts.CopyExif)
	if err != nil {
		return "", false, err
	}
	//no resize: the decoded image is written as is
	if err = writeImage(srcImg, srcBytes, dest, format, srcFormat, opts.Quality, opts.MakerNote, opts.CopyExif); err != nil {
		return "", false, err
	}
	return dest, true, nil
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

// NewOptions creates Options for an output in the given format, defaulting to
// crop center strategy, image quality 90% and Lanczos resampling. The format is
// a parameter rather than a field to set afterwards because it is required, and
// forgetting it would only surface at runtime
func NewOptions(transform TransformType, width, height int, copyExif bool, format Format) Options {
	return Options{Width: width, Height: height, Quality: 90, Anchor: Center,
		Transform: transform, Strategy: Lanczos, CopyExif: copyExif, Format: format}
}

// Open an image for editing
func Open(fileName string) (image.Image, error) {
	img, _, err := OpenOpts(fileName, false, false)
	return img, err
}

// OpenOpts opens an Image for editing. Possibly returning the src file as a byte slice
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

// writeImage writes dstImage to dest in the given format. Metadata is copied
// from srcBytes only when the destination is a jpeg and the source format can
// carry it, matching Options.CopyExif. Everything else is a plain encode
func writeImage(dstImage image.Image, srcBytes []byte, dest string, format, srcFormat Format,
	quality int, policy MakerNotePolicy, copyExif bool) error {
	if quality < 1 || quality > 100 {
		quality = 90
	}
	if format != FormatJpeg || !copyExif || !srcFormat.CanReadMetaData() || srcBytes == nil {
		return imaging.Save(dstImage, dest, imaging.JPEGQuality(quality))
	}
	return saveWithMetaData(dstImage, srcBytes, dest, srcFormat, quality, policy)
}

// saveWithMetaData encodes dstImage as a jpeg and copies the metadata of the
// source into it. srcFormat selects which source container to read from
func saveWithMetaData(dstImage image.Image, srcBytes []byte, dest string, srcFormat Format,
	quality int, policy MakerNotePolicy) error {
	dstBytes := new(bytes.Buffer)
	err := imaging.Encode(dstBytes, dstImage, imaging.JPEG, imaging.JPEGQuality(quality))
	if err != nil {
		return err
	}
	mde, err := metadata.NewJpegEditor(dstBytes.Bytes())
	if err != nil {
		return err
	}
	if srcFormat == FormatTiff {
		err = mde.CopyMetaDataFromTiff(srcBytes)
	} else {
		err = mde.CopyMetaData(srcBytes)
	}
	if err != nil {
		return err
	}
	mde.SetMakerNotePolicy(policy)
	err = mde.Exif().SetDate(metadata.ModifyDate, time.Now())
	if err != nil {
		return err
	}
	return mde.WriteFile(dest)
}

// Save an image (defaults to Jpeg Quality 90)
func Save(image image.Image, fileName string) error {
	return SaveOpts(image, fileName, 90, nil)
}

// SaveOpts saves an image. If srcExif is != nil it will try to extract and exif information from that
// and append it to the new file. Any MakerNote in srcExif is preserved, use
// SaveOptsMakerNote to control that
func SaveOpts(image image.Image, fileName string, quality int, srcExif []byte) error {
	return SaveOptsMakerNote(image, fileName, quality, srcExif, MakerNotePreserve)
}

// SaveOptsMakerNote is SaveOpts with control over what happens to an exif
// MakerNote carried in srcExif
func SaveOptsMakerNote(image image.Image, fileName string, quality int, srcExif []byte, policy MakerNotePolicy) error {
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
	mde.SetMakerNotePolicy(policy)
	err = mde.Exif().SetDate(metadata.ModifyDate, time.Now())
	if err != nil {
		return err
	}
	return mde.WriteFile(fileName)
}

// CropImage crops an Image.
func CropImage(img image.Image, crop image.Rectangle) image.Image {
	if crop.Empty() || !crop.In(img.Bounds()) {
		return img
	}
	return imaging.Crop(img, crop)
}

// RotateImage rotates an image. Negative angle indicates a counter clockwise rotation
func RotateImage(img image.Image, angle int) image.Image {
	//fix angle first:
	angle = angle % 360
	if angle == 0 {
		return img
	}
	//imaging rotate counter clockwise so need to adjust for that
	if angle < 0 {
		angle = -angle
	} else {
		angle = 360 - angle
	}
	return imaging.Rotate(img, float64(angle), color.Black)
}

// RotateAndCropFile rotates and crops an image file and saves the result
func RotateAndCropFile(source string, dest string, opts Options) error {

	angle := opts.Angle
	crop := opts.rectangle()
	if angle == 0 && crop.Empty() {
		return fmt.Errorf("neither angle or crop was provided")
	}

	srcImg, srcBytes, err := OpenOpts(source, false, opts.CopyExif)
	if err != nil {
		return err
	}

	srcImg = RotateImage(srcImg, angle)
	srcImg = CropImage(srcImg, crop)

	return SaveOptsMakerNote(srcImg, dest, opts.Quality, srcBytes, opts.MakerNote)

	/*	angle := opts.Angle
		crop := opts.Rectangle()
		if angle == 0 && crop.Empty() {
			return fmt.Errorf("neither angle or crop was provided")
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

// isJpegFile and isTiffFile are gone: formats are now detected from content via
// metadata.DetectFormatFile, and outputs are named from Format.Extension()

// TransformFile creates versions of source based on destinations, decoding the
// source once and writing every output from it.
//
// The keys of destinations are base names without an extension: the canonical
// extension for that destination's Options.Format is appended, so a key of
// "/out/thumb" with FormatJpeg writes "/out/thumb.jpg". A key that already
// carries an image extension is an error, see ErrDestHasExtension.
//
// Options.Format is required. The source format is detected from its content
// rather than its filename. Quality and CopyExif only apply when the
// destination is a jpeg, and metadata is only carried when the source is a
// jpeg or a tiff
func TransformFile(source string, destinations map[string]Options) error {
	//validate everything before writing anything, so a bad destination does not
	//leave a half finished set of outputs behind
	paths := make(map[string]string, len(destinations))
	for dest, options := range destinations {
		p, err := destPath(dest, options.Format)
		if err != nil {
			return err
		}
		paths[dest] = p
	}

	srcFormat, err := metadata.DetectFormatFile(source)
	if err != nil {
		return err
	}
	if !srcFormat.Supported() {
		return fmt.Errorf("%w: cannot read %v", ErrUnsupportedFormat, srcFormat)
	}

	var srcBytes []byte
	var srcImg image.Image
	if srcFormat.CanReadMetaData() {
		//keep the source bytes so metadata can be copied across
		srcImg, srcBytes, err = OpenOpts(source, false, true)
	} else {
		srcImg, err = Open(source)
	}
	if err != nil {
		return err
	}

	for dest, options := range destinations {
		destImg := transform(srcImg, options)
		err = writeImage(destImg, srcBytes, paths[dest], options.Format, srcFormat,
			options.Quality, options.MakerNote, options.CopyExif)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
func TranformFile(source string, destinations map[string]Options) error {
	ext := path.Ext(source)
	if ext == ".jpg" || ext == ".jpeg" {
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
	} else {
		return transformPngFile(source, destinations)

	}
}

func transformPngFile(source string, destinations map[string]Options) error {
	srcImg, err := Open(source)
	if err != nil {
		return err
	}
	for dest, options := range destinations {
		destImg := transform(srcImg, options)
		err = imaging.Save(destImg, dest)
		if err != nil {
			return err
		}
	}
	return nil
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
