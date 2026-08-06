package metadata

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
)

// Format identifies an image format. Detection is by content rather than by
// filename, since extensions are frequently wrong
type Format int

// Image formats mimage handles. Anything else, webp, heif, camera raw and so
// on, is FormatUnknown: there is nothing a caller could do differently for a
// format that is equally unsupported, and naming a few of them while omitting
// the rest would draw an arbitrary line
const (
	//FormatUnknown is anything mimage does not handle
	FormatUnknown Format = iota
	FormatJpeg
	FormatPng
	FormatGif
	FormatTiff
	FormatBmp
)

var formatNames = map[Format]string{
	FormatUnknown: "unknown",
	FormatJpeg:    "jpeg",
	FormatPng:     "png",
	FormatGif:     "gif",
	FormatTiff:    "tiff",
	FormatBmp:     "bmp",
}

// canonical extension per writable format. Formats mimage cannot write are
// deliberately absent
var formatExtensions = map[Format]string{
	FormatJpeg: ".jpg",
	FormatPng:  ".png",
	FormatGif:  ".gif",
	FormatTiff: ".tiff",
	FormatBmp:  ".bmp",
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

// Supported reports whether mimage can decode and transform this format. Every
// named format is supported, so this is true for anything but FormatUnknown
func (f Format) Supported() bool {
	return f != FormatUnknown
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
// leading dot, or "" for a format mimage cannot write
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

// DetectFormat identifies an image from its leading bytes. Only the header is
// examined, so a short prefix of the file is enough. Returns FormatUnknown for
// anything mimage does not handle, including a slice too short to tell
func DetectFormat(data []byte) Format {
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

// headerBytes is how much of a file DetectFormat needs. The longest signature
// is png at 8 bytes
const headerBytes = 16

// DetectFormatFile identifies an image file by its content. Only the header is
// read. An error means the file could not be read, not that it is not an image:
// an unrecognised file returns FormatUnknown with a nil error
func DetectFormatFile(fileName string) (Format, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return FormatUnknown, err
	}
	defer func() { _ = f.Close() }()

	header := make([]byte, headerBytes)
	n, err := f.Read(header)
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
// .jpg is reported as tiff. An error means the file could not be read
func SupportedFormat(fileName string) (bool, Format, error) {
	format, err := DetectFormatFile(fileName)
	if err != nil {
		return false, FormatUnknown, err
	}
	return format.Supported(), format, nil
}
