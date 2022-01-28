package metadata

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"
	"trimmer.io/go-xmp/xmp"
)

func getXmpModel(t *testing.T) *xmp.Document {
	bytes := getAssetBytes(XmpFile, t)
	model := &xmp.Document{}
	err := xmp.Unmarshal(bytes, model)
	if err != nil {
		t.Fatalf("Could not parse XMP Document %s ", err.Error())
	}
	return model
}

func TestNewMetaDataEditor(t *testing.T) {

}

func TestNewMetaDataEditorFile(t *testing.T) {

}

func TestMetaDataEditor_Bytes(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	var b []byte
	var err error

	//Test 1: Just write bytes and read them back again
	if b, err = mde.Bytes(); err != nil {
		t.Errorf(err.Error())
	} else {
		mde, err = NewMetaDataEditor(b)
		if err != nil {
			t.Errorf("Could not open editor from written bytes: %s", err.Error())
		}

	}
	//Test 2: Write bytes after an edit
	err = mde.SetExifDate(ModifyDate, time.Now())
	if b, err = mde.Bytes(); err != nil {
		t.Errorf(err.Error())
	} else {
		mde, err = NewMetaDataEditor(b)
		if err != nil {
			t.Errorf("Could not open editor from written bytes: %s", err.Error())
		}
	}

}

//Simpler test as TestMetaDataEditor_CopyMetaData tests this in more detail
func TestMetaDataEditor_CopyMetaDataFile(t *testing.T) {
	mde := getEditor(NoExifImg, t)
	if err := mde.CopyMetaDataFile(LeicaImg, CopyAll); err != nil {
		t.Fatalf("Could not copy metadata from file: %v", err)
	}
	md := editorMD(mde, t)
	if md.xmpData.IsEmpty() {
		t.Errorf("Expected XMP Section")
	}
	if md.iptcData.IsEmpty() {
		t.Errorf("Expected IPTC Section")
	}
	if md.exifData.IsEmpty() {
		t.Errorf("Expected Exif Section")
	}
	//test with a file that does not exist.
	if err := mde.CopyMetaDataFile("someDummyFile.jpg", CopyAll); err == nil {
		t.Errorf("Expected to fail when copying non existent file")
	}

}

func TestMetaDataEditor_CopyMetaData(t *testing.T) {
	sourceBytes := getAssetBytes(LeicaImg, t)
	destBytes := getAssetBytes(NoExifImg, t)

	destMde, err := NewMetaDataEditor(destBytes)
	if err != nil {
		t.Fatalf("could not open dest editor: %v", err)
	}
	err = destMde.CopyMetaData(sourceBytes, CopyAll)
	if err != nil {
		t.Fatalf("could not copy metadata: %v", err)
	}
	sourceMde, err := NewMetaDataEditor(sourceBytes)
	if err != nil {
		t.Fatalf("could not open source editor")
	}

	//check exif:
	_, sourceSeq, e1 := sourceMde.sl.FindExif()
	_, destSeq, e2 := destMde.sl.FindExif()

	if e1 != nil || e2 != nil {
		t.Errorf("could not find source exif segement or dest exif segment")
	} else if bytes.Compare(sourceSeq.Data, destSeq.Data) != 0 {
		t.Errorf("exif segments are not the same")
	}

	//check xmp:
	_, sourceSeq, e1 = sourceMde.sl.FindXmp()
	_, destSeq, e2 = destMde.sl.FindXmp()

	if e1 != nil || e2 != nil {
		t.Errorf("could not find source xmp segement or dest exif segment")
	} else if bytes.Compare(sourceSeq.Data, destSeq.Data) != 0 {
		t.Errorf("xmp segments are not the same")
	}

	//check xmp:
	_, sourceSeq, e1 = sourceMde.sl.FindIptc()
	_, destSeq, e2 = destMde.sl.FindIptc()

	if e1 != nil || e2 != nil {
		t.Errorf("could not find source iptc segement or dest exif segment")
	} else if bytes.Compare(sourceSeq.Data, destSeq.Data) != 0 {
		t.Errorf("iptc segments are not the same")
	}

}

func TestMetaDataEditor_DropAll(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	err := mde.DropAll()
	if err != nil {
		t.Fatalf("error droppging all metadata: %v", err)
	}
	//make sure segments are dropped
	if mde.HasIptc() {
		t.Errorf("Image still contains IPTC data")
	}
	if mde.HasExif() {
		t.Errorf("Image still contains EXIF data")
	}
	if mde.HasXmp() {
		t.Errorf("Image still contains XMP data")
	}
	//now write and reopen and make sure
	b, err := mde.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes after drop metadata: %v", err)
	}
	mde, err = NewMetaDataEditor(b)
	if err != nil {
		t.Fatalf("could not open editor afater dropping metadata: %v ", err)
	}
	if mde.HasIptc() {
		t.Errorf("Image still contains IPTC data after writing bytes")
	}
	if mde.HasExif() {
		t.Errorf("Image still contains EXIF data after writing bytes")
	}
	if mde.HasXmp() {
		t.Errorf("Image still contains XMP data after writing bytes")
	}
}

func TestMetaDataEditor_DropIptc(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	err := mde.DropIptc()
	if err != nil {
		t.Fatalf("error droppging iptc data: %v", err)
	}
	//make sure segments are dropped
	if mde.HasIptc() {
		t.Errorf("Image still contains IPTC data")
	}
	//now write and reopen and make sure
	b, err := mde.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes after drop iptc: %v", err)
	}
	mde, err = NewMetaDataEditor(b)
	if err != nil {
		t.Fatalf("could not open editor after dropping iptc: %v ", err)
	}
	if mde.HasIptc() {
		t.Errorf("Image still contains IPTC data after writing bytes")
	}
}

func TestMetaDataEditor_DropXmp(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	err := mde.DropXmp()
	if err != nil {
		t.Fatalf("error droppging xmp data: %v", err)
	}
	//make sure segments are dropped
	if mde.HasXmp() {
		t.Errorf("Image still contains IPTC data")
	}
	//now write and reopen and make sure
	b, err := mde.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes after drop xmp: %v", err)
	}
	mde, err = NewMetaDataEditor(b)
	if err != nil {
		t.Fatalf("could not open editor after dropping xmp: %v ", err)
	}
	if mde.HasXmp() {
		t.Errorf("Image still contains IPTC data after writing bytes")
	}
}

func TestMetaDataEditor_DropExif(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	err := mde.DropExif()
	if err != nil {
		t.Fatalf("error droppging exif data: %v", err)
	}
	//make sure segments are dropped
	if mde.HasExif() {
		t.Errorf("Image still contains Exif data")
	}
	//now write and reopen and make sure
	b, err := mde.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes after drop exif: %v", err)
	}
	mde, err = NewMetaDataEditor(b)
	if err != nil {
		t.Fatalf("could not open editor after dropping exif: %v ", err)
	}
	if mde.HasExif() {
		t.Errorf("Image still contains IPTC data after writing bytes")
	}
}

func TestMetaDataEditor_HasExif(t *testing.T) {
	withMetaData := getEditor(LeicaImg, t)
	noMetaData := getEditor(NoExifImg, t)
	if !withMetaData.HasExif() {
		t.Errorf("expected exif")
	}
	if noMetaData.HasExif() {
		t.Errorf("expected no exif")
	}
}

func TestMetaDataEditor_HasIptc(t *testing.T) {
	withMetaData := getEditor(LeicaImg, t)
	noMetaData := getEditor(NoExifImg, t)
	if !withMetaData.HasIptc() {
		t.Errorf("expected iptc")
	}
	if noMetaData.HasIptc() {
		t.Errorf("expected no iptc")
	}
}

func TestMetaDataEditor_HasXmp(t *testing.T) {
	withMetaData := getEditor(LeicaImg, t)
	noMetaData := getEditor(NoExifImg, t)
	if !withMetaData.HasXmp() {
		t.Errorf("expected xmp")
	}
	if noMetaData.HasXmp() {
		t.Errorf("expected no xmp")
	}
}

func TestMetaDataEditor_MetaData(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	md, err := mde.MetaData()
	expectedTitle := "Morning Fog"
	expectedCameraModel := "LEICA Q2"

	if err != nil {
		t.Fatalf("Could not retrive Metadata")
	}
	if md.iptcData.IsEmpty() {
		t.Errorf("Expected IPTC in Metada")
	}
	if md.xmpData.IsEmpty() {
		t.Errorf("Expected Xmp in Metada")
	}
	if md.exifData.IsEmpty() {
		t.Errorf("Expected Exif in Metada")
	}
	//check some fields in summary
	if md.Summary().Title != expectedTitle {
		t.Errorf("Expected %s got %s", expectedTitle, md.Summary().Title)
	}
	if md.Summary().CameraModel != expectedCameraModel {
		t.Errorf("Expected %s got %s", expectedCameraModel, md.Summary().CameraModel)
	}

}

func TestMetaDataEditor_SetExifDate(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	tnow := time.Now()
	tnow = tnow.Truncate(1 * time.Second)
	dates := []ExifDate{OriginalDate, ModifyDate, DigitizedDate}
	for _, dt := range dates {
		if err := mde.SetExifDate(dt, tnow); err != nil {
			t.Errorf("Could not set date: %v got error %v", dt, err)
		}
	}
	md := editorMD(mde, t)
	if cmpDates(tnow, md.Summary().OriginalDate) {
		t.Errorf("Expected %v got %v", tnow, md.Summary().OriginalDate)
	}
	if cmpDates(tnow, md.Summary().ModifyDate) {
		t.Errorf("Expected %v got %v", tnow, md.Summary().ModifyDate)
	}
}

func TestMetaDataEditor_CommitExifChanges(t *testing.T) {
	var err error
	mde := getEditor(LeicaImg, t)
	md := editorMD(mde, t)
	model := md.Summary().CameraModel
	newModel := "A new model"

	//change model
	mde.SetIfdTag(IFD_Model, newModel)

	//load meta data again. It should not have changed the underlying segments
	out := new(bytes.Buffer)
	if err = mde.sl.Write(out); err != nil {
		t.Fatalf("Could not write segments: %v", err)
	}
	if md, err = NewMetaData(out.Bytes()); err != nil {
		t.Fatalf("Could not read metadata: %v", err)
	}
	if model != md.Summary().CameraModel {
		t.Errorf("Expected %s got %s", model, md.Summary().CameraModel)
	}

	//now commit changes
	if err := mde.CommitExifChanges(); err != nil {
		t.Fatalf("Could not commit changes: %v", err)
	}

	//write bytes again
	out = new(bytes.Buffer)
	if err = mde.sl.Write(out); err != nil {
		t.Fatalf("Could not write segments: %v", err)
	}
	if md, err = NewMetaData(out.Bytes()); err != nil {
		t.Fatalf("Could not read metadata: %v", err)
	}
	if newModel != md.Summary().CameraModel {
		t.Errorf("Expected %s got %s", newModel, md.Summary().CameraModel)
	}
}

func TestMetaDataEditor_SetExifTag(t *testing.T) {

	//TODO: make sure to cover all various types as well non-happy paths

	expectedFocalLength := URat{350, 10}
	mde := getEditor(LeicaImg, t)
	if err := mde.SetExifTag(ExifIFD_FocalLength, expectedFocalLength); err != nil {
		t.Fatalf("Could not set exif tag: %v", err)
	}
	md := editorMD(mde, t)
	focalLength := URat{}
	md.exifData.ScanIfdExif(ExifIFD_FocalLength, &focalLength)
	if focalLength != expectedFocalLength {
		t.Errorf("expected focalLength %v got %v", expectedFocalLength, focalLength)
	}
}

func TestMetaDataEditor_SetIfdTag(t *testing.T) {

	//TODO: make sure to cover all various types as well non-happy paths
	//TODO: this this is broken. Fix so its pulling the actual lens info value and not the summary which
	//will default to Exif_LensSpecification
	expectedLensInfo := LensInfo{
		MinFocalLength:           URat{2800, 100},
		MaxFocalLength:           URat{2800, 100},
		MinFNumberMinFocalLength: URat{392, 256},
		MinFNumberMaxFocalLength: URat{500, 256},
	}
	mde := getEditor(LeicaImg, t)
	md := editorMD(mde, t)

	if expectedLensInfo.MinFocalLength.IsZero() {
		t.Fatalf("Expected Lens Info to not be zero")
	}

	if err := mde.SetIfdTag(IFD_DNGLensInfo, expectedLensInfo); err != nil {
		t.Fatalf("Could not set ifd tag: %v", err)
	}
	md = editorMD(mde, t)
	lensInfo := LensInfo{}
	md.exifData.ScanIfdRoot(IFD_DNGLensInfo, &lensInfo)
	if lensInfo != expectedLensInfo {
		t.Errorf("expected %v got %v", expectedLensInfo, lensInfo)
	}

}

func TestMetaDataEditor_SetImageDescription(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	expImageDescription := "A new Image Description"
	if err := mde.SetImageDescription(expImageDescription); err != nil {
		t.Fatalf("Could not set Image Description: %v", err)
	}
	md := editorMD(mde, t)
	ret, err := md.exifData.GetIfdImageDescription()
	if err != nil {
		t.Fatalf("Could not get Image Description from Metadata: %v", err)
	} else if ret != expImageDescription {
		t.Fatalf("Expected %s got %s", expImageDescription, ret)
	}
}

func TestMetaDataEditor_SetUserComment(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	expUserComment := "A new User Comment"
	if err := mde.SetUserComment(expUserComment); err != nil {
		t.Fatalf("Could not set User Comment: %v", err)
	}
	md := editorMD(mde, t)
	ret, err := md.exifData.GetIfdUserComment()
	if err != nil {
		t.Fatalf("Could not get Image Comment from Metadata: %v", err)
	} else if ret != expUserComment {
		t.Fatalf("Expected %s go %s", expUserComment, ret)
	}
}

func TestMetaDataEditor_SetExifAndXmp(t *testing.T) {
	expectedFocalLength := URat{350, 10}
	expectedLensInfo := LensInfo{
		MinFocalLength:           URat{2800, 100},
		MaxFocalLength:           URat{2800, 100},
		MinFNumberMinFocalLength: URat{392, 256},
		MinFNumberMaxFocalLength: URat{500, 256},
	}
	expTitle := "New Title"
	expKeywords := []string{"keyword 1", "keyword 2"}

	testSetExifXmp := func(assetName string, t *testing.T) {
		e := getEditor(assetName, t)
		model := getXmpModel(t)

		e.SetXmp(model)
		if err := e.SetIfdTag(IFD_DNGLensInfo, expectedLensInfo); err != nil {
			t.Errorf("Could not set ifd tag for %s: %v", assetName, err)
		}
		if err := e.SetExifTag(ExifIFD_FocalLength, expectedFocalLength); err != nil {
			t.Errorf("Could not set exif tag for %s: %v", assetName, err)
		}
		md := editorMD(e, t)

		//now verify
		if dcm := md.xmpData.DublinCore(); dcm != nil {
			if dcm.Title.Default() != expTitle {
				t.Errorf("Expected title for %s: %s got %s", assetName, expTitle, dcm.Title.Default())
			}
			if len(dcm.Subject) != len(expKeywords) {
				t.Errorf("Expected number of keywords for %s: %v got %v", assetName, len(expKeywords), len(dcm.Subject))
			}
			for _, kv := range dcm.Subject {
				if kv != expKeywords[0] && kv != expKeywords[1] {
					t.Errorf("Did not find expected keyword for %s: %s", assetName, kv)
				}
			}
		} else {
			t.Errorf("Expected Dublin Core Xmp data got nil")
		}
		focalLength := URat{}
		md.exifData.ScanIfdExif(ExifIFD_FocalLength, &focalLength)
		if focalLength != expectedFocalLength {
			t.Errorf("expected focalLength for %s: %v got %v", assetName, expectedFocalLength, focalLength)
		}
		lensInfo := LensInfo{}
		md.exifData.ScanIfdRoot(IFD_DNGLensInfo, &lensInfo)
		if lensInfo != expectedLensInfo {
			t.Errorf("expected lensInfo for %s: %v got %v", assetName, expectedLensInfo, lensInfo)
		}
	}
	testSetExifXmp(NoExifImg, t)
	testSetExifXmp(LeicaImg, t)

}

func TestMetaDataEditor_SetXmp(t *testing.T) {

	model := getXmpModel(t)
	expTitle := "New Title"
	expKeywords := []string{"keyword 1", "keyword 2"}

	mde := getEditor(NoExifImg, t)
	if err := mde.SetXmp(model); err != nil {
		t.Fatalf("Could not set xmp model")
	}
	md := editorMD(mde, t)
	if dcm := md.xmpData.DublinCore(); dcm != nil {
		if dcm.Title.Default() != expTitle {
			t.Errorf("Expected title %s got %s", expTitle, dcm.Title.Default())
		}
		if len(dcm.Subject) != len(expKeywords) {
			t.Errorf("Expected %v keywords got %v", len(expKeywords), len(dcm.Subject))
		}
		for _, kv := range dcm.Subject {
			if kv != expKeywords[0] && kv != expKeywords[1] {
				t.Errorf("Did not find expected keyword %s", kv)
			}
		}
	} else {
		t.Errorf("Expected Dublin Core Xmp data got nil")
	}
}

func TestMetaDataEditor_WriteFile(t *testing.T) {
	mde := getEditor(LeicaImg, t)
	//make a change
	expImageDesc := "This is a new description"
	mde.SetImageDescription(expImageDesc)
	out := filepath.Join(os.TempDir(), "TestWriteFile.jpg")
	if err := mde.WriteFile(out); err != nil {
		t.Fatalf("Could not write file: %v", err)
	}
	//now reopen
	md := getMetaData(out, t)
	if desc, err := md.exifData.GetIfdImageDescription(); err != nil {
		t.Errorf("Could not retrieve image description: %v", err)
	} else if desc != expImageDesc {
		t.Errorf("Expected %s got %s", expImageDesc, desc)
	}
	//finally delete temp file
	if err := os.Remove(out); err != nil {
		t.Errorf("Could not delete temp file: %v", err)
	}

	//test wrong file extension
	wrongOut := filepath.Join(os.TempDir(), "TestWriteFile.png")
	if err := mde.WriteFile(wrongOut); err == nil {
		t.Errorf("Write file should not accept a png extension")
	} else if err != JpegWrongFileExtErr {
		t.Errorf("Expecteded error %v got %v", JpegWrongFileExtErr, err)
	}
}

func cmpDates(t1, t2 time.Time) bool {
	if t1.Year() != t2.Year() {
		return false
	}
	if t1.Month() != t2.Month() {
		return false
	}
	if t1.Day() != t2.Day() {
		return false
	}
	if t1.Hour() != t2.Hour() {
		return false
	}
	if t1.Minute() != t2.Minute() {
		return false
	}
	return t1.Second() == t2.Second()
}
