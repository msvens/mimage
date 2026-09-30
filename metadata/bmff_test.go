package metadata

import (
	"encoding/binary"
	"testing"
)

// ftyp builds the start of an iso base media file: an ftyp box with the given
// major brand and compatible brands, followed by the start of a moov box
func ftyp(major string, compatible ...string) []byte {
	size := 16 + 4*len(compatible)
	b := make([]byte, 0, size+8)
	b = binary.BigEndian.AppendUint32(b, uint32(size))
	b = append(b, "ftyp"...)
	b = append(b, major...)
	b = append(b, 0, 0, 0, 0) //minor version
	for _, c := range compatible {
		b = append(b, c...)
	}
	b = binary.BigEndian.AppendUint32(b, 1024)
	return append(b, "moov"...)
}

// box builds a bare box header of the given size and type
func box(size uint32, boxType string) []byte {
	return append(binary.BigEndian.AppendUint32(nil, size), boxType...)
}

func TestDetectVideo(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want Format
	}{
		//what the common producers write
		{"lightroom/ffmpeg mp4", ftyp("isom", "isom", "iso2", "avc1", "mp41"), FormatMp4},
		{"mp42", ftyp("mp42", "isom", "mp42"), FormatMp4},
		{"iphone mov", ftyp("qt  ", "qt  "), FormatMov},
		{"itunes m4v", ftyp("M4V ", "M4V ", "M4A ", "mp42", "isom"), FormatUnknown},
		{"m4v without audio brand", ftyp("M4V ", "M4V ", "mp42", "isom"), FormatMp4},
		{"3gp", ftyp("3gp4", "isom", "3gp4"), FormatMp4},
		{"dash", ftyp("dash", "iso6", "mp41"), FormatMp4},
		//a vendor brand is accepted through its compatible brands
		{"vendor brand mp4", ftyp("XAVC", "XAVC", "mp42", "iso2"), FormatMp4},
		{"vendor brand mov", ftyp("XAVC", "qt  "), FormatMov},
		{"vendor brand alone", ftyp("XAVC"), FormatUnknown},
		//same container, not a video
		{"heic", ftyp("heic", "mif1", "heic"), FormatUnknown},
		{"heif", ftyp("mif1", "mif1", "heic"), FormatUnknown},
		{"avif", ftyp("avif", "avif", "mif1", "miaf"), FormatUnknown},
		{"avif sequence", ftyp("avis", "avis", "msf1", "iso8"), FormatUnknown},
		{"m4a", ftyp("M4A ", "M4A ", "mp42", "isom"), FormatUnknown},
		{"image brand among compatible", ftyp("isom", "isom", "mif1"), FormatUnknown},
		//QuickTime from before ftyp existed
		{"legacy moov", box(2048, "moov"), FormatMov},
		{"legacy wide", box(8, "wide"), FormatMov},
		{"legacy mdat to end of file", box(0, "mdat"), FormatMov},
		{"legacy mdat 64 bit size", box(1, "mdat"), FormatMov},
		{"legacy moov with impossible size", box(3, "moov"), FormatUnknown},
		{"free box first", box(64, "free"), FormatUnknown},
		//truncated and malformed
		{"ftyp too short to hold a brand", box(24, "ftyp"), FormatUnknown},
		{"ftyp size below minimum", append(box(12, "ftyp"), "isom"...), FormatUnknown},
		{"brand cut off at 12 bytes", ftyp("isom")[:12], FormatMp4},
		{"eleven bytes", ftyp("isom")[:11], FormatUnknown},
	}
	for _, tc := range tests {
		if got := DetectFormat(tc.data); got != tc.want {
			t.Errorf("%s: DetectFormat = %v, want %v", tc.name, got, tc.want)
		}
	}
}

// A compatible brand list longer than the header DetectFormatFile reads is
// scanned only as far as the bytes go
func TestDetectVideoTruncatedBrands(t *testing.T) {
	brands := make([]string, 20)
	for i := range brands {
		brands[i] = "XXXX"
	}
	brands[len(brands)-1] = "mp42"
	data := ftyp("XAVC", brands...)
	if got := DetectFormat(data[:headerBytes]); got != FormatUnknown {
		t.Errorf("a brand beyond the header should not be seen, got %v", got)
	}
	if got := DetectFormat(data); got != FormatMp4 {
		t.Errorf("with the whole box, got %v, want mp4", got)
	}
}
