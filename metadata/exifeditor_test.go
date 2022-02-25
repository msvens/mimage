package metadata

import (
	"testing"
	"time"
)

func TestNewExifEditor(t *testing.T) {
	tests := []string{LeicaImg, NoExifImg}
	for _, fname := range tests {
		b := getAssetBytes(fname, t)
		if sl, err := parseJpegBytes(b); err != nil {
			t.Errorf("Could not parse jpeg file %s", fname)
		} else {
			je, e1 := NewExifEditor(sl)
			if e1 != nil {
				t.Errorf("Could not open exif editor got %v", e1)
			}
			if je.dirty {
				t.Errorf("expected dirty to be false")
			}
			if je.rootIb == nil {
				t.Errorf("expected rootIb to be non nil")
			}
		}
	}
	if _, err := NewExifEditor(nil); err == nil {
		t.Errorf("Expected error when creating a NewExifEditor with a nil segment list")
	}

}

func TestNewExifEditorEmpty(t *testing.T) {
	tests := []bool{false, true}
	for _, flag := range tests {
		ee, err := NewExifEditorEmpty(flag)
		if err != nil {
			t.Errorf("Could not create empty exif editor with flag %v got error: %v", flag, err)
		}
		if ee.dirty != flag {
			t.Errorf("Expected dirty to be %v got %v", flag, ee.dirty)
		}
		if ee.rootIb == nil {
			t.Errorf("Expected rootIb to not be nil")
		}
	}
}

func TestExifEditor_EmptyBuilder(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	err := je.ee.Clear(true)
	if err != nil {
		t.Fatalf("Could not set an empty builder: %v", err)
	}
	b, err := je.Bytes()
	if err != nil {
		t.Fatalf("Could not write bytes: %v", err)
	}
	md, err := NewMetaData(b)
	if err != nil {
		t.Fatalf("Could not open MetaData")
	}
	if md.Summary().Software != ExifEditorSoftware {
		t.Errorf("Expected %s got %s", ExifEditorSoftware, md.Summary().Software)
	}
	if md.Summary().CameraMake != "" {
		t.Errorf("Expected empty camera make got %s", md.Summary().CameraMake)
	}

}

func TestExifEditor_IfdBuilder(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	je.Exif().SetDirty() //force write of software tag
	builder, changed := je.Exif().IfdBuilder()
	if builder == nil {
		t.Fatalf("IfdBuilder is nil")
	}
	if !changed {
		t.Fatalf("Expected builder to have changed")
	}
	je.Exif().SetDirty() //force jpegEditor to use the IfdBuilder
	b, err := je.Bytes()
	if err != nil {
		t.Fatalf("Error getting bytes")
	}
	md, err := NewMetaData(b)
	if err != nil {
		t.Fatalf("Could not open metadata")
	}
	softwareTag := ""
	if err = md.exifData.ScanIfdRoot(IFD_Software, &softwareTag); err != nil {
		t.Fatalf("Expected IFD_Software Tag got error: %v", err)
	}
	if softwareTag != ExifEditorSoftware {
		t.Errorf("Expected %s got %s", ExifEditorSoftware, softwareTag)
	}
}

func TestExifEditor_IsDirty(t *testing.T) {
	je, _ := NewExifEditorEmpty(true)
	if !je.IsDirty() {
		t.Errorf("Expected editor to not be dirty")
	}
	je, _ = NewExifEditorEmpty(false)
	if je.IsDirty() {
		t.Errorf("Expected editor to not be dirty")
	}
	//edit a field which should set the editor to dirty
	err := je.SetImageDescription("some description")
	if err != nil {
		t.Fatalf("Could not set tag: %v", err)
	}
	if !je.IsDirty() {
		t.Errorf("Expected editor to be dirty after setting tag")
	}

}

func TestExifEditor_SetDate(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	tnow := time.Now()
	tnow = tnow.Truncate(1 * time.Second)
	dates := []ExifDate{OriginalDate, ModifyDate, DigitizedDate}
	for _, dt := range dates {
		if err := je.Exif().SetDate(dt, tnow); err != nil {
			t.Errorf("Could not set date: %v got error %v", dt, err)
		}
	}
	md := jpegEditorMD(je, t)
	if cmpDates(tnow, md.Summary().OriginalDate) {
		t.Errorf("Expected %v got %v", tnow, md.Summary().OriginalDate)
	}
	if cmpDates(tnow, md.Summary().ModifyDate) {
		t.Errorf("Expected %v got %v", tnow, md.Summary().ModifyDate)
	}
	//TODO: Verifiy digitized date
}

func TestExifEditor_SetDirty(t *testing.T) {
	je, _ := NewExifEditorEmpty(false)
	if je.IsDirty() {
		t.Errorf("Expected editor to not be dirty")
	}
	je.SetDirty()
	if !je.IsDirty() {
		t.Errorf("Expected editor to be dirty")
	}

}

func TestExifEditor_SetIfdExifTag(t *testing.T) {
	//TODO: make sure to cover all various types as well non-happy paths
	expectedFocalLength := URat{350, 10}
	je := getJpegEditor(LeicaImg, t)
	if err := je.Exif().SetIfdExifTag(ExifIFD_FocalLength, expectedFocalLength); err != nil {
		t.Fatalf("Could not set exif tag: %v", err)
	}
	md := jpegEditorMD(je, t)
	focalLength := URat{}
	if err := md.exifData.ScanIfdExif(ExifIFD_FocalLength, &focalLength); err != nil {
		t.Errorf("Could not scan ifdexif: %v", err)
	}
	if focalLength != expectedFocalLength {
		t.Errorf("expected focalLength %v got %v", expectedFocalLength, focalLength)
	}
}

func TestExifEditor_SetIfdRootTag(t *testing.T) {
	//TODO: make sure to cover all various types as well non-happy paths
	expectedLensInfo := LensInfo{
		MinFocalLength:           URat{2800, 100},
		MaxFocalLength:           URat{2800, 100},
		MinFNumberMinFocalLength: URat{392, 256},
		MinFNumberMaxFocalLength: URat{500, 256},
	}
	je := getJpegEditor(LeicaImg, t)

	if err := je.Exif().SetIfdRootTag(IFD_DNGLensInfo, expectedLensInfo); err != nil {
		t.Fatalf("Could not set ifd tag: %v", err)
	}
	md := jpegEditorMD(je, t)
	lensInfo := LensInfo{}
	if err := md.exifData.ScanIfdRoot(IFD_DNGLensInfo, &lensInfo); err != nil {
		t.Errorf("Could not scan DNGLensInfo: %v", err)
	}
	if lensInfo != expectedLensInfo {
		t.Errorf("expected %v got %v", expectedLensInfo, lensInfo)
	}
}

func TestExifEditor_SetImageDescription(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	expImageDescription := "A new Image Description"
	if err := je.Exif().SetImageDescription(expImageDescription); err != nil {
		t.Fatalf("Could not set Image Description: %v", err)
	}
	md := jpegEditorMD(je, t)
	ret := md.exifData.GetIfdImageDescription()
	if ret != expImageDescription {
		t.Fatalf("Expected %s got %s", expImageDescription, ret)
	}
}

func TestExifEditor_SetUserComment(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	expUserComment := "A new User Comment"
	if err := je.Exif().SetUserComment(expUserComment); err != nil {
		t.Fatalf("Could not set User Comment: %v", err)
	}
	md := jpegEditorMD(je, t)
	ret := md.exifData.GetIfdUserComment()
	if ret != expUserComment {
		t.Fatalf("Expected %s go %s", expUserComment, ret)
	}

}
