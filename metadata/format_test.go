package metadata

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFormatBytes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want Format
	}{
		{"jpeg", []byte{0xff, 0xd8, 0xff, 0xe0}, FormatJpeg},
		{"png", []byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a}, FormatPng},
		{"gif87a", []byte("GIF87a...."), FormatGif},
		{"gif89a", []byte("GIF89a...."), FormatGif},
		{"tiff little endian", []byte{'I', 'I', 0x2a, 0x00}, FormatTiff},
		{"tiff big endian", []byte{'M', 'M', 0x00, 0x2a}, FormatTiff},
		{"bmp", []byte("BM....."), FormatBmp},
		//formats mimage does not handle are simply unknown
		{"webp", []byte("RIFF\x00\x00\x00\x00WEBPVP8 "), FormatUnknown},
		{"heic", []byte("\x00\x00\x00\x18ftypheic"), FormatUnknown},
		{"png prefix only", []byte{0x89, 'P', 'N', 'G'}, FormatUnknown},
		{"empty", []byte{}, FormatUnknown},
		{"one byte", []byte{0xff}, FormatUnknown},
		{"three bytes", []byte{'I', 'I', 0x2a}, FormatUnknown},
		{"text", []byte("<?xml version=\"1.0\"?>"), FormatUnknown},
	}
	for _, tc := range tests {
		if got := DetectFormat(tc.data); got != tc.want {
			t.Errorf("%s: DetectFormat = %v, want %v", tc.name, got, tc.want)
		}
	}
}

// Detection is by content. A file named .jpg that holds a tiff must report tiff
func TestDetectFormatFile(t *testing.T) {
	tests := []struct {
		file string
		want Format
	}{
		{LeicaImg, FormatJpeg},
		{CanonImg, FormatJpeg},
		{NoExifImg, FormatJpeg},
		{TiffImg, FormatTiff},
		{AssetPath + "leica.png", FormatPng},
		{NonImageFile, FormatUnknown},
		{XmpFile, FormatUnknown},
	}
	for _, tc := range tests {
		got, err := DetectFormatFile(tc.file)
		if err != nil {
			t.Errorf("%s: unexpected error %v", tc.file, err)
			continue
		}
		if got != tc.want {
			t.Errorf("%s: DetectFormatFile = %v, want %v", tc.file, got, tc.want)
		}
	}
}

func TestDetectFormatFileIgnoresName(t *testing.T) {
	src, err := os.ReadFile(TiffImg)
	if err != nil {
		t.Fatalf("could not read fixture: %v", err)
	}
	liar := filepath.Join(t.TempDir(), "definitely.jpg")
	if err = os.WriteFile(liar, src, 0644); err != nil {
		t.Fatalf("could not write: %v", err)
	}
	got, err := DetectFormatFile(liar)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != FormatTiff {
		t.Errorf("a tiff named .jpg detected as %v, want tiff", got)
	}
}

// A missing file is an error. A file that simply is not an image is not
func TestDetectFormatFileErrors(t *testing.T) {
	if _, err := DetectFormatFile(filepath.Join(t.TempDir(), "nope.jpg")); err == nil {
		t.Errorf("expected an error for a missing file")
	}
	empty := filepath.Join(t.TempDir(), "empty.jpg")
	if err := os.WriteFile(empty, []byte{}, 0644); err != nil {
		t.Fatalf("could not write: %v", err)
	}
	got, err := DetectFormatFile(empty)
	if err != nil {
		t.Errorf("an empty file should not be an io error, got %v", err)
	}
	if got != FormatUnknown {
		t.Errorf("empty file detected as %v, want unknown", got)
	}
}

func TestSupportedFormat(t *testing.T) {
	tests := []struct {
		file    string
		wantOk  bool
		wantFmt Format
	}{
		{LeicaImg, true, FormatJpeg},
		{TiffImg, true, FormatTiff},
		{AssetPath + "leica.png", true, FormatPng},
		{NonImageFile, false, FormatUnknown},
	}
	for _, tc := range tests {
		ok, format, err := SupportedFormat(tc.file)
		if err != nil {
			t.Errorf("%s: unexpected error %v", tc.file, err)
			continue
		}
		if ok != tc.wantOk || format != tc.wantFmt {
			t.Errorf("%s: SupportedFormat = (%v, %v), want (%v, %v)",
				tc.file, ok, format, tc.wantOk, tc.wantFmt)
		}
	}
}

// The capability matrix is part of the public contract
func TestFormatCapabilities(t *testing.T) {
	tests := []struct {
		format    Format
		supported bool
		readMeta  bool
		editMeta  bool
		ext       string
	}{
		{FormatJpeg, true, true, true, ".jpg"},
		{FormatTiff, true, true, false, ".tiff"},
		{FormatPng, true, false, false, ".png"},
		{FormatGif, true, false, false, ".gif"},
		{FormatBmp, true, false, false, ".bmp"},
		{FormatUnknown, false, false, false, ""},
	}
	for _, tc := range tests {
		if got := tc.format.Supported(); got != tc.supported {
			t.Errorf("%v.Supported() = %v, want %v", tc.format, got, tc.supported)
		}
		if got := tc.format.CanReadMetaData(); got != tc.readMeta {
			t.Errorf("%v.CanReadMetaData() = %v, want %v", tc.format, got, tc.readMeta)
		}
		if got := tc.format.CanEditMetaData(); got != tc.editMeta {
			t.Errorf("%v.CanEditMetaData() = %v, want %v", tc.format, got, tc.editMeta)
		}
		if got := tc.format.Extension(); got != tc.ext {
			t.Errorf("%v.Extension() = %q, want %q", tc.format, got, tc.ext)
		}
	}
}

func TestIsImageExtension(t *testing.T) {
	yes := []string{".jpg", ".JPG", "jpeg", ".png", ".TIFF", ".tif", ".bmp", ".gif", ".heic", ".webp"}
	no := []string{"", ".", ".txt", ".v2", ".03", "photo", ".jpgx"}
	for _, e := range yes {
		if !IsImageExtension(e) {
			t.Errorf("IsImageExtension(%q) = false, want true", e)
		}
	}
	for _, e := range no {
		if IsImageExtension(e) {
			t.Errorf("IsImageExtension(%q) = true, want false", e)
		}
	}
}
