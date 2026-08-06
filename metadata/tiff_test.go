package metadata

import (
	"os"
	"testing"
)

// The tiff fixture is leica.jpg converted, so every summary field that
// describes the photograph rather than the file must match
func TestNewMetaDataFromTiff(t *testing.T) {
	tiff := getMetaData(TiffImg, t)
	jpg := getMetaData(LeicaImg, t)
	ts, js := tiff.Summary(), jpg.Summary()

	if err := tiff.SummaryErr(); err != nil {
		t.Errorf("unexpected summary error: %v", err)
	}
	if ts.CameraMake != js.CameraMake {
		t.Errorf("CameraMake = %q, want %q", ts.CameraMake, js.CameraMake)
	}
	if ts.CameraModel != js.CameraModel {
		t.Errorf("CameraModel = %q, want %q", ts.CameraModel, js.CameraModel)
	}
	if ts.ISO != js.ISO {
		t.Errorf("ISO = %d, want %d", ts.ISO, js.ISO)
	}
	//compare rationals by value: converting to tiff renormalised them, so
	//f/2.5 is stored as 5/2 in the tiff and 25/10 in the jpeg
	if ts.FNumber.Float32() != js.FNumber.Float32() {
		t.Errorf("FNumber = %v, want %v", ts.FNumber.Float32(), js.FNumber.Float32())
	}
	if ts.ExposureTime.Float32() != js.ExposureTime.Float32() {
		t.Errorf("ExposureTime = %v, want %v", ts.ExposureTime.Float32(), js.ExposureTime.Float32())
	}
	if !cmpDates(ts.OriginalDate, js.OriginalDate) {
		t.Errorf("OriginalDate = %v, want %v", ts.OriginalDate, js.OriginalDate)
	}
	//Title and Rating come from the XMP packet in ifd0 tag 0x02bc
	if ts.Title != js.Title {
		t.Errorf("Title = %q, want %q (xmp not read?)", ts.Title, js.Title)
	}
	if ts.Rating != js.Rating {
		t.Errorf("Rating = %d, want %d (xmp not read?)", ts.Rating, js.Rating)
	}
}

// Dimensions come from ifd0 tags rather than from decoding the image
func TestTiffDimensions(t *testing.T) {
	md := getMetaData(TiffImg, t)
	if md.ImageWidth == 0 || md.ImageHeight == 0 {
		t.Errorf("expected non zero dimensions, got %dx%d", md.ImageWidth, md.ImageHeight)
	}
	if md.ImageWidth <= md.ImageHeight {
		t.Errorf("leica.tiff is landscape, got %dx%d", md.ImageWidth, md.ImageHeight)
	}
}

// A tiff carries an icc profile that go-exif cannot decode. It must be found so
// it can be excluded from a transplant rather than failing the whole copy
func TestUnreadableRootTags(t *testing.T) {
	md := getMetaData(TiffImg, t)
	unreadable := unreadableRootTags(md.Exif())
	found := false
	for _, tag := range unreadable {
		if tag == tiffIccTag {
			found = true
		}
	}
	if !found {
		t.Errorf("expected the icc profile tag %#04x among unreadable tags, got %v", tiffIccTag, unreadable)
	}
}

// Anything that is neither jpeg nor tiff must still be rejected
func TestNewMetaDataUnknownContainer(t *testing.T) {
	for _, f := range []string{NonImageFile, XmpFile} {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("could not read %s: %v", f, err)
		}
		if _, err := NewMetaData(b); err != ErrParseImage {
			t.Errorf("%s: expected ErrParseImage, got %v", f, err)
		}
	}
}

func TestJpegEditor_CopyMetaDataFromTiff(t *testing.T) {
	tiffBytes := getAssetBytes(TiffImg, t)
	//start from an image with no metadata at all
	je := getJpegEditor(NoExifImg, t)
	if err := je.CopyMetaDataFromTiff(tiffBytes); err != nil {
		t.Fatalf("could not copy metadata from tiff: %v", err)
	}
	md := jpegEditorMD(je, t)
	src := getMetaData(TiffImg, t)

	if got, want := md.Summary().CameraMake, src.Summary().CameraMake; got != want {
		t.Errorf("CameraMake = %q, want %q", got, want)
	}
	if got, want := md.Summary().ISO, src.Summary().ISO; got != want {
		t.Errorf("ISO = %d, want %d", got, want)
	}
	if got, want := md.Summary().Title, src.Summary().Title; got != want {
		t.Errorf("Title = %q, want %q (xmp segment not written?)", got, want)
	}
	if got, want := md.Summary().Rating, src.Summary().Rating; got != want {
		t.Errorf("Rating = %d, want %d", got, want)
	}
}

// Tags describing how the tiff stored its pixels, and the payloads that belong
// in their own jpeg segment, must not end up in the transplanted exif
func TestJpegEditor_CopyMetaDataFromTiffDropsTiffTags(t *testing.T) {
	je := getJpegEditor(NoExifImg, t)
	if err := je.CopyMetaDataFromTiff(getAssetBytes(TiffImg, t)); err != nil {
		t.Fatalf("could not copy metadata from tiff: %v", err)
	}
	md := jpegEditorMD(je, t)
	ifd := md.Exif().Ifd(RootIFD)
	if ifd == nil {
		t.Fatal("expected a root ifd in the output")
	}
	for _, tag := range TiffDropTags() {
		if entries, err := ifd.FindTagWithId(uint16(tag)); err == nil && len(entries) > 0 {
			t.Errorf("tag %#04x (%s) leaked into the transplanted exif",
				tag, ExifTagName(RootIFD, tag))
		}
	}
}

// The xmp packet becomes an APP1 segment, so it must not also remain as an
// ifd0 tag, or the image would carry two copies
func TestJpegEditor_CopyMetaDataFromTiffNoDuplicateXmp(t *testing.T) {
	je := getJpegEditor(NoExifImg, t)
	if err := je.CopyMetaDataFromTiff(getAssetBytes(TiffImg, t)); err != nil {
		t.Fatalf("could not copy metadata from tiff: %v", err)
	}
	b, err := je.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes: %v", err)
	}
	reopened, err := NewJpegEditor(b)
	if err != nil {
		t.Fatalf("could not reopen: %v", err)
	}
	//exactly one xmp segment
	n := 0
	for _, s := range reopened.sl.Segments() {
		if s.MarkerId == 0xe1 && len(s.Data) > 4 && string(s.Data[:4]) == "http" {
			n++
		}
	}
	if n > 1 {
		t.Errorf("expected at most one xmp segment, found %d", n)
	}
	//and the xmp is readable
	md := jpegEditorMD(reopened, t)
	if md.Summary().Title == "" {
		t.Errorf("expected the xmp title to survive the round trip")
	}
}
