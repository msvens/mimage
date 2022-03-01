package metadata

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func getJpegEditor(fname string, t *testing.T) *JpegEditor {
	je, err := NewJpegEditorFile(fname)
	if err != nil {
		t.Fatalf("Could not retrieve editor for file: %v", err)
	}
	return je
}

func reloadJpegEditor(je *JpegEditor, fatal bool, t *testing.T) *JpegEditor {
	b, err := je.Bytes()
	if err != nil {
		if fatal {
			t.Fatalf("Could not get bytes from jpegEditor: %v", err)
		} else {
			t.Errorf("Could not get bytes from jpegEditor: %v", err)
		}
	}
	ret, err := NewJpegEditor(b)
	if err != nil {
		if fatal {
			t.Fatalf("Could not get bytes from jpegEditor: %v", err)
		} else {
			t.Errorf("Could not get bytes from jpegEditor: %v", err)
		}
	}
	return ret
}

func jpegEditorMD(je *JpegEditor, t *testing.T) *MetaData {
	md, err := je.MetaData()
	if err != nil {
		t.Fatalf("Could not get metadata ")
	}
	return md
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

func TestNewJpegEditor(t *testing.T) {
	assets := []string{LeicaImg, NoExifImg}
	for _, asset := range assets {
		if b, err := ioutil.ReadFile(asset); err != nil {
			t.Errorf("Could not read file: %s error %v", asset, err)
		} else {
			if _, err1 := NewJpegEditor(b); err1 != nil {
				t.Errorf("Could not open editor for bytes: %s error %v", asset, err1)
			}
		}
	}
}

func TestNewJpegEditorFile(t *testing.T) {
	assets := []string{LeicaImg, NoExifImg}
	for _, asset := range assets {
		if _, err := NewJpegEditorFile(asset); err != nil {
			t.Errorf("Could not open editor for file: %s error %v", asset, err)
		}
	}
}

func TestJpegEditor_Bytes(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	var b []byte
	var err error

	//Test 1: Just write bytes and read them back again
	if b, err = je.Bytes(); err != nil {
		t.Errorf(err.Error())
	} else {
		je, err = NewJpegEditor(b)
		if err != nil {
			t.Errorf("Could not open editor from written bytes: %s", err.Error())
		}

	}
	//Test 2: Write bytes after an edit
	err = je.Exif().SetDate(ModifyDate, time.Now())
	if b, err = je.Bytes(); err != nil {
		t.Errorf(err.Error())
	} else {
		je, err = NewJpegEditor(b)
		if err != nil {
			t.Errorf("Could not open editor from written bytes: %s", err.Error())
		}
	}

}

func TestJpegEditor_CopyMetaData(t *testing.T) {
	sourceBytes := getAssetBytes(LeicaImg, t)
	destBytes := getAssetBytes(NoExifImg, t)

	destJe, err := NewJpegEditor(destBytes)
	if err != nil {
		t.Fatalf("could not open dest editor: %v", err)
	}
	err = destJe.CopyMetaData(sourceBytes)
	if err != nil {
		t.Fatalf("could not copy metadata: %v", err)
	}
	//reopen dest
	b, err := destJe.Bytes()
	if err != nil {
		t.Fatalf("Could not write dest after copy: %v", err)
	}
	destJe, err = NewJpegEditor(b)
	if err != nil {
		t.Fatalf("Could not reopen jpegeditor after copy: %v", err)
	}

	sourceJe, err := NewJpegEditor(sourceBytes)
	if err != nil {
		t.Fatalf("could not open source editor")
	}

	//check exif:
	_, sourceSeq, e1 := sourceJe.sl.FindExif()
	_, destSeq, e2 := destJe.sl.FindExif()

	if e1 != nil || e2 != nil {
		fmt.Println(e1)
		fmt.Println(e2)
		t.Errorf("could not find source exif segement or dest exif segment")
	} else if bytes.Compare(sourceSeq.Data, destSeq.Data) != 0 {
		t.Errorf("exif segments are not the same")
	}

	//check XmpEditor:
	_, sourceSeq, e1 = sourceJe.sl.FindXmp()
	_, destSeq, e2 = destJe.sl.FindXmp()

	if e1 != nil || e2 != nil {
		t.Errorf("could not find source XmpEditor segement or dest exif segment")
	} else if bytes.Compare(sourceSeq.Data, destSeq.Data) != 0 {
		t.Errorf("XmpEditor segments are not the same")
	}

	//check iptc:
	_, sourceSeq, e1 = sourceJe.sl.FindIptc()
	_, destSeq, e2 = destJe.sl.FindIptc()

	if e1 != nil || e2 != nil {
		t.Errorf("could not find source iptc segement or dest exif segment")
	} else if bytes.Compare(sourceSeq.Data, destSeq.Data) != 0 {
		t.Errorf("iptc segments are not the same")
	}

}

func TestJpegEditor_DropExif(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	err := je.DropExif()
	if err != nil {
		t.Fatalf("error droppging all metadata: %v", err)
	}

	//now write and reopen and make sure no metadata exists
	b, err := je.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes after dropping metadata: %v", err)
	}
	var md *MetaData
	if md, err = NewMetaData(b); err != nil {
		t.Fatalf("could not open metadata after dropping metadata: %v ", err)
	}
	if !md.exifData.IsEmpty() {
		t.Errorf("Image still contains exifData")
	}
	if md.xmpData.IsEmpty() {
		t.Errorf("Image should contain Xmp data")
	}
	if md.iptcData.IsEmpty() {
		t.Errorf("Image should contain IPTC data")
	}
}

func TestJpegEditor_DropIptc(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	err := je.DropIptc()
	if err != nil {
		t.Fatalf("error droppging all metadata: %v", err)
	}

	//now write and reopen and make sure no metadata exists
	b, err := je.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes after dropping metadata: %v", err)
	}
	var md *MetaData
	if md, err = NewMetaData(b); err != nil {
		t.Fatalf("could not open metadata after dropping metadata: %v ", err)
	}
	if md.exifData.IsEmpty() {
		t.Errorf("Image should contain exifData")
	}
	if md.xmpData.IsEmpty() {
		t.Errorf("Image should contain Xmp data")
	}
	if !md.iptcData.IsEmpty() {
		t.Errorf("Image still contains IPTC data")
	}

}

func TestJpegEditor_DropMetaData(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	err := je.DropMetaData()
	if err != nil {
		t.Fatalf("error droppging all metadata: %v", err)
	}

	//now write and reopen and make sure no metadata exists
	b, err := je.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes after dropping metadata: %v", err)
	}
	var md *MetaData
	if md, err = NewMetaData(b); err != nil {
		t.Fatalf("could not open metadata after dropping metadata: %v ", err)
	}
	if !md.exifData.IsEmpty() {
		t.Errorf("Image still contains Exif data after writing bytes")
	}
	if !md.xmpData.IsEmpty() {
		t.Errorf("Image still contains Xmp data after writing bytes")
	}
	if !md.iptcData.IsEmpty() {
		t.Errorf("Image still contains IPTC data after writing bytes")
	}

}

func TestJpegEditor_DropXmp(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	err := je.DropXmp()
	if err != nil {
		t.Fatalf("error droppging all metadata: %v", err)
	}

	//now write and reopen and make sure no metadata exists
	b, err := je.Bytes()
	if err != nil {
		t.Fatalf("could not write bytes after dropping metadata: %v", err)
	}
	var md *MetaData
	if md, err = NewMetaData(b); err != nil {
		t.Fatalf("could not open metadata after dropping metadata: %v ", err)
	}
	if md.exifData.IsEmpty() {
		t.Errorf("Image should contain exifData")
	}
	if !md.xmpData.IsEmpty() {
		t.Errorf("Image still contains Xmp data")
	}
	if md.iptcData.IsEmpty() {
		t.Errorf("Image should contain IPTC data")
	}

}

func TestJpegEditor_MetaData(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	md, err := je.MetaData()
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

func TestJpegEditor_SetKeywords(t *testing.T) {
	expKeywords := []string{"keyword 1", "keyword 2", "keyword 3"}
	for _, fname := range []string{LeicaImg, NoExifImg} {
		je := getJpegEditor(fname, t)
		err := je.SetKeywords(expKeywords)
		if err != nil {
			t.Fatalf("Error setting keywords for image %s got error: %v", fname, err)
		}
		md := jpegEditorMD(je, t)
		actKeywords := md.Iptc().GetKeywords()
		if !reflect.DeepEqual(expKeywords, actKeywords) {
			t.Errorf("Expected %v got %v", expKeywords, actKeywords)
		}
		actKeywords = md.Xmp().GetKeywords()
		if !reflect.DeepEqual(expKeywords, actKeywords) {
			t.Errorf("Expected %v got %v", expKeywords, actKeywords)
		}
	}

}

func TestJpegEditor_SetTitle(t *testing.T) {
	fnames := []string{LeicaImg, NoExifImg}
	for _, fname := range fnames {
		je := getJpegEditor(fname, t)
		expTitle := "New Title"
		if err := je.SetTitle(expTitle); err != nil {
			t.Fatalf("Could not set title for image %s got error %v", fname, err)
		}
		md := jpegEditorMD(je, t)
		if md.Iptc().GetTitle() != expTitle {
			t.Errorf("Expected title %s got %s", expTitle, md.Iptc().GetTitle())
		}
		if md.Xmp().GetTitle() != expTitle {
			t.Errorf("Expected title %s got %s", expTitle, md.Xmp().GetTitle())
		}
		if md.Exif().GetIfdImageDescription() != expTitle {
			t.Errorf("Expected title %s got %s", expTitle, md.Xmp().GetTitle())
		}
	}
}

func TestJpegEditor_WriteFile(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	//make a change
	expImageDesc := "This is a new description"
	if err := je.Exif().SetImageDescription(expImageDesc); err != nil {
		t.Fatalf("Coul not set image description: %v", err)
	}
	out := filepath.Join(os.TempDir(), "TestWriteFile.jpg")
	if err := je.WriteFile(out); err != nil {
		t.Fatalf("Could not write file: %v", err)
	}
	//now reopen
	md := getMetaData(out, t)
	if desc := md.exifData.GetIfdImageDescription(); desc != expImageDesc {
		t.Errorf("Expected %s got %s", expImageDesc, desc)
	}
	//finally delete temp file
	if err := os.Remove(out); err != nil {
		t.Errorf("Could not delete temp file: %v", err)
	}

	//test wrong file extension
	wrongOut := filepath.Join(os.TempDir(), "TestWriteFile.png")
	if err := je.WriteFile(wrongOut); err == nil {
		t.Errorf("Write file should not accept a png extension")
	} else if err != JpegWrongFileExtErr {
		t.Errorf("Expecteded error %v got %v", JpegWrongFileExtErr, err)
	}
}

func ExampleJpegEditor_SetTitle() {
	je, err := NewJpegEditorFile("../assets/leica.jpg")
	if err != nil {
		fmt.Printf("Could not retrieve editor for file: %v\n", err)
		return
	}
	err = je.SetTitle("some new title")
	if err != nil {
		fmt.Printf("Could not set title: %v\n", err)
		return
	}
	md, err := je.MetaData()
	if err != nil {
		fmt.Printf("Could not get metadata: %v\n", err)
	}
	fmt.Printf("New Title: %v\n", md.Summary().Title)

	//Output: New Title: some new title
}
