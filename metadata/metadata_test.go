package metadata

import (
	"bytes"
	exifundefined "github.com/dsoprea/go-exif/v3/undefined"
	"io/ioutil"
	"testing"
	"time"
)

const AssetPath = "../assets/"
const LeicaImg = AssetPath + "leica.jpg"
const NoExifImg = AssetPath + "noexif.jpg"
const NikonImg = AssetPath + "nikon.jpg"
const CanonImg = AssetPath + "canon.jpg"
const GPSImg = AssetPath + "gps.jpg"
const NonImageFile = AssetPath + "exiftool-leica-g1.json"

func getAssetBytes(fname string, t *testing.T) []byte {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatalf("Could not read file: %v", err)
	}
	return b
}

func getEditor(fname string, t *testing.T) *MetaDataEditor {
	mde, err := NewMetaDataEditorFile(fname)
	if err != nil {
		t.Fatalf("Could not retrieve editor for file: %v", err)
	}
	return mde
}

func getMetaData(fname string, t *testing.T) *MetaData {
	md, err := ParseFile(fname)
	if err != nil {
		t.Fatalf("Could not retrieve metadata for file: %v", err)
	}
	return md
}

func editorMD(mde *MetaDataEditor, t *testing.T) *MetaData {
	md, err := mde.MetaData()
	if err != nil {
		t.Fatalf("Could not get metadata ")
	}
	return md
}

func TestParse(t *testing.T) {
	data := getAssetBytes(NikonImg, t)
	if _, err := Parse(data); err != nil {
		t.Errorf("Could not parse image file: %v", err)
	}
	empty := []byte{}
	if _, err := Parse(empty); err == nil {
		t.Errorf("Expected parse error")
	} else if err != ParseImageErr {
		t.Errorf("Expected error %v got error %v", ParseImageErr, err)
	}
	nonImageFile := getAssetBytes(NonImageFile, t)
	if _, err := Parse(nonImageFile); err == nil {
		t.Errorf("Expected parse error")
	} else if err != ParseImageErr {
		t.Errorf("Expected error %v got error %v", ParseImageErr, err)
	}
}

func TestParseFile(t *testing.T) {
	if _, err := ParseFile(NikonImg); err != nil {
		t.Errorf("Could not parse Image file: %v", err)
	}
	if _, err := ParseFile(NonImageFile); err == nil {
		t.Errorf("Expected parse error")
	} else if err != ParseImageErr {
		t.Errorf("Expected error %v got error %v", ParseImageErr, err)
	}
	//non existent file
	if _, err := ParseFile("somefile.jpg"); err == nil {
		t.Errorf("Expected parse error")
	}
}

func TestMetaData_GetIfdImageDescription(t *testing.T) {
	md := getMetaData(LeicaImg, t)
	if desc, err := md.GetIfdImageDescription(); err == nil {
		t.Errorf("Did not expect any image description got %s", desc)
	} else if err != IfdTagNotFoundErr {
		t.Errorf("Expected err %v got %v", IfdTagNotFoundErr, err)
	}
	mde := getEditor(LeicaImg, t)
	mde.SetImageDescription("some description")
	md = editorMD(mde, t)
	if _, err := md.GetIfdImageDescription(); err != nil {
		t.Errorf("Expected description got error %v", err)
	}
}

func TestMetaData_GetIfdUserComment(t *testing.T) {
	md := getMetaData(LeicaImg, t)
	if desc, err := md.GetIfdUserComment(); err == nil {
		t.Errorf("Did not expect any user comment got %s", desc)
	} else if err != IfdTagNotFoundErr {
		t.Errorf("Expected err %v got %v", IfdTagNotFoundErr, err)
	}
	mde := getEditor(LeicaImg, t)
	mde.SetUserComment("some comment")
	md = editorMD(mde, t)
	if _, err := md.GetIfdUserComment(); err != nil {
		t.Errorf("Expected user comment got error %v", err)
	}
}

func TestMetaData_HasExif(t *testing.T) {
	md := getMetaData(LeicaImg, t)
	mdNo := getMetaData(NoExifImg, t)
	if !md.HasExif() {
		t.Errorf("Expected exif data")
	}
	if mdNo.HasExif() {
		t.Errorf("Expected no exif data")
	}
}

func TestMetaData_HasIptc(t *testing.T) {
	md := getMetaData(LeicaImg, t)
	mdNo := getMetaData(NoExifImg, t)
	if !md.HasIptc() {
		t.Errorf("Expected iptc data")
	}
	if mdNo.HasIptc() {
		t.Errorf("Expected no iptc data")
	}
}

func TestMetaData_HasXmp(t *testing.T) {
	md := getMetaData(LeicaImg, t)
	mdNo := getMetaData(NoExifImg, t)
	if !md.HasXmp() {
		t.Errorf("Expected xmp data")
	}
	if mdNo.HasXmp() {
		t.Errorf("Expected no xmp data")
	}
}

func TestMetaData_IfdGPS(t *testing.T) {
	/*Exiftool GPSInfo extraction:
	"GPS": {
	    "GPSVersionID": "2.2.0.0",
	    "GPSLatitudeRef": "North",
	    "GPSLatitude": "26 deg 35' 12.00\"",
	    "GPSLongitudeRef": "West",
	    "GPSLongitude": "80 deg 3' 13.00\"",
	    "GPSAltitudeRef": "Below Sea Level",
	    "GPSAltitude": "0 m",
	    "GPSTimeStamp": "01:22:57",
	    "GPSDateStamp": "2018:04:29"
	  }
	*/
	md := getMetaData(GPSImg, t)
	if md.IfdGPS() == nil {
		t.Fatalf("Expected IfdGps information")
	}
	if md.Summary.GPSInfo == nil {
		t.Fatalf("Expected GPS information in Summary")
	}
	//check only latitude, timestamp and altitude
	gps := md.Summary.GPSInfo
	if gps.Altitude != 0 {
		t.Errorf("Expected 0 got %v", gps.Altitude)
	}
	if gps.Latitude.Degrees != 26 {
		t.Errorf("Expected 26 degrees got %v", gps.Latitude.Degrees)
	}
	if gps.Latitude.Minutes != 35 {
		t.Errorf("Expected 35 minutes got %v", gps.Latitude.Minutes)
	}
	if gps.Latitude.Seconds != 12 {
		t.Errorf("Expected 12 seconds got %v", gps.Latitude.Seconds)
	}
	//for the timestamp simplify and check year and seconds
	if gps.Timestamp.Year() != 2018 {
		t.Errorf("Expected Year 2018 got %v", gps.Timestamp.Year())
	}
	if gps.Timestamp.Second() != 57 {
		t.Errorf("Expected Second 57 got %v", gps.Timestamp.Second())
	}
	//make sure GPSInfo is nil when no such metainformation is available
	md = getMetaData(LeicaImg, t)
	if md.IfdGPS() != nil {
		t.Errorf("Expected no IdfGps")
	}
	if md.Summary.GPSInfo != nil {
		t.Errorf("Expected GPSInfo to be nil in summary")
	}

}

func TestMetaData_IfdExif(t *testing.T) {
	md := getMetaData(LeicaImg, t)
	mdNo := getMetaData(NoExifImg, t)
	if md.IfdExif() == nil {
		t.Errorf("Expected ifdExif got nil")
	}
	if mdNo.IfdExif() != nil {
		t.Errorf("Expected nil ifdExif")
	}
}

func TestMetaData_IfdRoot(t *testing.T) {
	md := getMetaData(LeicaImg, t)
	mdNo := getMetaData(NoExifImg, t)
	if md.IfdRoot() == nil {
		t.Errorf("Expected ifdRoot got nil")
	}
	if mdNo.IfdRoot() != nil {
		t.Errorf("Expected nil ifdRoot")
	}
}

func TestMetaData_ScanIfdDate(t *testing.T) {
	/*exif tool extracted dates:
	"ModifyDate": "2021:01:10 17:36:45"
	"DateTimeOriginal": "2020:10:27 09:34:03",
	"CreateDate": "2020:10:27 09:34:03",
	"OffsetTime": "+01:00",
	"OffsetTimeOriginal": "+02:00",
	*/
	expModifyDate := "2021-01-10 17:36:45 +0100 CET"
	expOrigDate := "2020-10-27 09:34:03 +0200 +0200"
	expDigDate := "2020-10-27 09:34:03 +0200 +0200"
	md := getMetaData(LeicaImg, t)
	dates := map[ExifDate]string{
		OriginalDate:  expOrigDate,
		ModifyDate:    expModifyDate,
		DigitizedDate: expDigDate,
	}
	ret := time.Time{}
	for k, v := range dates {
		if err := md.ScanIfdDate(k, &ret); err != nil {
			t.Errorf("Expected ExifDate %v got error: %v", k, err)
		} else if ret.String() != v {
			t.Errorf("Expected %s got %s", v, ret)
		}
	}
}

func TestMetaData_ScanIfdExifTag(t *testing.T) {
	/*Ifd exif extracted by exiftool (from lecia img):
	"ExifIFD": {
	    "ExposureTime": "1/250",
	    "FNumber": 2.5,
	    "ExposureProgram": "Manual",
	    "ISO": 100,
	    "SensitivityType": "Standard Output Sensitivity",
	    "StandardOutputSensitivity": 100,
	    "ShutterSpeedValue": "1/250",
	    "ApertureValue": 2.5,
	    "ExposureCompensation": -0.3,
		...
	*/
	md := getMetaData(LeicaImg, t)
	expExposureTime := URat{Numerator: 1, Denominator: 250}
	ret := URat{}
	if err := md.ScanIfdExifTag(Exif_ExposureTime, &ret); err != nil {
		t.Errorf("Expected IFD_ExposureTime got error: %v", err)
	} else if ret != expExposureTime {
		t.Errorf("Expected %s got %s", expExposureTime, ret)
	}
	//now try a tag that is non existent
	ret1 := exifundefined.Tag9286UserComment{}
	if err := md.ScanIfdRootTag(Exif_UserComment, &ret1); err == nil {
		t.Errorf("Expected error when reading non existent root tag")
	} else if err != IfdTagNotFoundErr {
		t.Errorf("Expected error %v got %v", IfdTagNotFoundErr, err)
	}

}

func TestMetaData_ScanIfdRootTag(t *testing.T) {
	/*Ifd root extracted by exiftool:
	  "IFD0": {
	    "Make": "LEICA CAMERA AG",
	    "Model": "LEICA Q2",
	    "XResolution": 144,
	    "YResolution": 144,
	    "ResolutionUnit": "inches",
	    "Software": "Adobe Photoshop Lightroom Classic 10.1 (Macintosh)",
	    "ModifyDate": "2021:01:10 17:36:45"
	  },
	*/
	md := getMetaData(LeicaImg, t)
	expSoftware := "Adobe Photoshop Lightroom Classic 10.1 (Macintosh)"
	ret := ""
	if err := md.ScanIfdRootTag(IFD_Software, &ret); err != nil {
		t.Errorf("Expected IFD_Software got error: %v", err)
	} else if ret != expSoftware {
		t.Errorf("Expected %s got %s", expSoftware, ret)
	}
	//now try a tag that is non existent
	if err := md.ScanIfdRootTag(IFD_ImageDescription, &ret); err == nil {
		t.Errorf("Expected error when reading non existent root tag")
	} else if err != IfdTagNotFoundErr {
		t.Errorf("Expected error %v got %v", IfdTagNotFoundErr, err)
	}

}

func TestMetaData_ScanIptcApplicationTag(t *testing.T) {
	/*exiftool extract IPTC application from Leica Img
	"IPTC": {
	    "ApplicationRecordVersion": 4,
	    "ObjectName": "Morning Fog",
	    "DateCreated": "2020:10:27",
	    "TimeCreated": "09:34:03+02:00",
	    "DigitalCreationDate": "2020:10:27",
	    "DigitalCreationTime": "09:34:03+02:00"
	  },
	*/
	md := getMetaData(LeicaImg, t)
	expObjectName := "Morning Fog"
	ret := ""
	if err := md.ScanApplicationTag(Iptc2_ObjectName, &ret); err != nil {
		t.Errorf("Expected IPTC Application Object Name to be set got: %v", err)
	} else if ret != expObjectName {
		t.Errorf("Expected %s got %s", expObjectName, ret)
	}
	//non existent tag
	if err := md.ScanApplicationTag(Iptc2_Keywords, &ret); err == nil {
		t.Errorf("Expected error got %s", ret)
	} else if err != IptcTagNotFoundErr {
		t.Errorf("Expected error %v got error %v", IptcTagNotFoundErr, err)
	}
	md.PrintIptc()
}

func TestMetaData_ScanIptcEnvelopTag(t *testing.T) {
	/*exiftool extract IPTC application from Leica Img
	"IPTC": {
	    "CodedCharacterSet": "UTF8",
	  },
	*/
	md := getMetaData(LeicaImg, t)
	//expObjectName := "Morning Fog"
	expectedCharSet := []byte{27, 37, 71} //ESC%G (which is UTF
	ret := []byte{}
	if err := md.ScanIptcEnvelopTag(Iptc1_CodedCharacterSet, &ret); err != nil {
		t.Errorf("Expected CodedCharacterSet got error %v", err)
	} else if bytes.Compare(ret, expectedCharSet) != 0 {
		t.Errorf("Expected %v got %v", expectedCharSet, ret)
	}
	//non existent tag
	if err := md.ScanIptcEnvelopTag(Iptc1_Destination, &ret); err == nil {
		t.Errorf("Expected error for non existing tag")
	} else if err != IptcTagNotFoundErr {
		t.Errorf("Expected error %v got error %v", IptcTagNotFoundErr, err)
	}

}

func TestMetaData_Xmp(t *testing.T) {
	md := getMetaData(LeicaImg, t)
	mdNo := getMetaData(NoExifImg, t)
	if md.Xmp() == nil {
		t.Errorf("Expected XMP got nil")
	}
	if mdNo.Xmp() != nil {
		t.Errorf("Expected nil XMP")
	}
}
