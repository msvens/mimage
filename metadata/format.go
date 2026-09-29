package metadata

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
)

// Format identifies an image or video format. Detection is by content rather
// than by filename, since extensions are frequently wrong
type Format int

// Formats mimage handles. Anything else, webp, heif, camera raw and so on, is
// FormatUnknown: there is nothing a caller could do differently for a format
// that is equally unsupported, and naming a few of them while omitting the rest
// would draw an arbitrary line
const (
	//FormatUnknown is anything mimage does not handle
	FormatUnknown Format = iota
	FormatJpeg
	FormatPng
	FormatGif
	FormatTiff
	FormatBmp
	//FormatMp4 is an iso base media file with an mp4 family brand. Only the
	//container is known from the header, not the codec inside
	FormatMp4
	//FormatMov is a QuickTime movie. Only the container is known from the
	//header, not the codec inside
	FormatMov
)

var formatNames = map[Format]string{
	FormatUnknown: "unknown",
	FormatJpeg:    "jpeg",
	FormatPng:     "png",
	FormatGif:     "gif",
	FormatTiff:    "tiff",
	FormatBmp:     "bmp",
	FormatMp4:     "mp4",
	FormatMov:     "mov",
}

// canonical extension per named format
var formatExtensions = map[Format]string{
	FormatJpeg: ".jpg",
	FormatPng:  ".png",
	FormatGif:  ".gif",
	FormatTiff: ".tiff",
	FormatBmp:  ".bmp",
	FormatMp4:  ".mp4",
	FormatMov:  ".mov",
}

// imageExtensions are the extensions that mark a filename as already naming an
// image. Used to reject a destination base name that carries one. Deliberately
// wider than the formats mimage handles: "photo.heic" is just as wrong a base
// name as "photo.jpg", even though mimage cannot read either as heif
var imageExtensions = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
	".tif": true, ".tiff": true, ".bmp": true, ".webp": true,
	".heic": true, ".heif": true, ".avif": true, ".jxl": true,
}

func (f Format) String() string {
	if n, found := formatNames[f]; found {
		return n
	}
	return "unknown"
}

// Supported reports whether mimage can handle this format, image or video.
// Every named format is supported, so this is true for anything but
// FormatUnknown. Use IsImage and IsVideo to tell which kind it is.
//
// For video this is a first check only: the header identifies the container,
// not the codec inside it. The video package's Probe is what establishes that a
// particular file can actually be used
func (f Format) Supported() bool {
	return f != FormatUnknown
}

// IsImage reports whether this is an image format, one the img package can
// decode and transform
func (f Format) IsImage() bool {
	switch f {
	case FormatJpeg, FormatPng, FormatGif, FormatTiff, FormatBmp:
		return true
	default:
		return false
	}
}

// IsVideo reports whether this is a video container format, handled by the
// video package rather than img
func (f Format) IsVideo() bool {
	return f == FormatMp4 || f == FormatMov
}

// CanReadMetaData reports whether mimage can read exif, iptc and xmp from this
// format
func (f Format) CanReadMetaData() bool {
	return f == FormatJpeg || f == FormatTiff
}

// CanEditMetaData reports whether mimage can write metadata back into this
// format
func (f Format) CanEditMetaData() bool {
	return f == FormatJpeg
}

// Extension is the canonical file extension for this format including the
// leading dot, or "" for FormatUnknown. Note that mimage can only write image
// formats: a video extension is for naming a file, not a conversion target
func (f Format) Extension() string {
	return formatExtensions[f]
}

// IsImageExtension reports whether ext, with or without its leading dot, names
// an image format. Comparison is case insensitive
func IsImageExtension(ext string) bool {
	if ext == "" {
		return false
	}
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return imageExtensions[strings.ToLower(ext)]
}

// DetectFormat identifies an image or video from its leading bytes. Only the
// header is examined, so a short prefix of the file is enough. Returns
// FormatUnknown for anything mimage does not handle, including a slice too short
// to tell
func DetectFormat(data []byte) Format {
	if f := detectVideo(data); f != FormatUnknown {
		return f
	}
	switch {
	case len(data) >= 2 && data[0] == 0xff && data[1] == 0xd8:
		return FormatJpeg
	case len(data) >= 8 && bytes.Equal(data[:8], []byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a}):
		return FormatPng
	case len(data) >= 6 && (bytes.Equal(data[:6], []byte("GIF87a")) || bytes.Equal(data[:6], []byte("GIF89a"))):
		return FormatGif
	case len(data) >= 4 && (bytes.Equal(data[:4], []byte{'I', 'I', 0x2a, 0x00}) ||
		bytes.Equal(data[:4], []byte{'M', 'M', 0x00, 0x2a})):
		return FormatTiff
	case len(data) >= 2 && data[0] == 'B' && data[1] == 'M':
		return FormatBmp
	default:
		return FormatUnknown
	}
}

// headerBytes is how much of a file DetectFormat reads. Image signatures need at
// most 8 bytes, but an ftyp box lists its compatible brands after the first 16
const headerBytes = 64

// DetectFormatFile identifies an image or video file by its content. Only the
// header is read. An error means the file could not be read, not that it is not
// media: an unrecognised file returns FormatUnknown with a nil error
func DetectFormatFile(fileName string) (Format, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return FormatUnknown, err
	}
	defer func() { _ = f.Close() }()

	header := make([]byte, headerBytes)
	n, err := io.ReadFull(f, header)
	if err != nil && n == 0 {
		//an empty or truncated file is not an image, but it is not an io
		//failure either
		if errors.Is(err, io.EOF) {
			return FormatUnknown, nil
		}
		return FormatUnknown, err
	}
	return DetectFormat(header[:n]), nil
}

// SupportedFormat reports whether mimage can handle fileName, and what it
// actually is. The format is detected from the file content, so a tiff named
// .jpg is reported as tiff. An error means the file could not be read. See
// Format.Supported for what that means for a video
func SupportedFormat(fileName string) (bool, Format, error) {
	format, err := DetectFormatFile(fileName)
	if err != nil {
		return false, FormatUnknown, err
	}
	return format.Supported(), format, nil
}
