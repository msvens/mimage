package metadata

import (
	exifundefined "github.com/dsoprea/go-exif/v3/undefined"
	"strings"
	"testing"
	"time"
)

func getExifData(fname string, t *testing.T) *ExifData {
	bytes := getAssetBytes(fname, t)
	segments, err := parseJpegBytes(bytes)
	if err != nil {
		t.Fatalf("Could not retrieve segments for file: %v", err)
	}
	exifData, _ := NewExifData(segments)
	return exifData
}

func TestExifValueIsAllowed(t *testing.T) {
	expExifFlashValues := map[uint16]bool{
		0x0:   true,
		0x41:  true,
		0x100: false,
	}
	expIopInteroperabilityIndexValues := map[string]bool{
		"R03":      true,
		"THM":      true,
		"Whatever": false,
	}
	for k, exp := range expExifFlashValues {
		if v := ExifValueIsAllowed(ExifIFD, ExifIFD_Flash, k); v != exp {
			t.Errorf("Expected allowed value: %v got %v", exp, v)
		}
	}
	for k, exp := range expIopInteroperabilityIndexValues {
		if v := ExifValueIsAllowed(InteropIFD, InteropIFD_InteropIndex, k); v != exp {
			t.Errorf("Expected allowed value: %v got %v", exp, v)
		}
	}
}

func TestExifValueStringNoErr(t *testing.T) {
	expExifFlashValues := map[uint16]string{
		0x0:   "No Flash",
		0x41:  "Fired, Red-eye reduction",
		0x100: "Undefined",
	}
	expIopInteroperabilityIndexValues := map[string]string{
		"R03":      "R03 - DCF option file (Adobe RGB)",
		"THM":      "THM - DCF thumbnail file",
		"Whatever": "Undefined",
	}
	for k, exp := range expExifFlashValues {
		if v := ExifValueString(ExifIFD, ExifIFD_Flash, k); !strings.EqualFold(exp, v) {
			t.Errorf("Expected value string: %v got %v", exp, v)
		}
	}
	for k, exp := range expIopInteroperabilityIndexValues {
		if v := ExifValueString(InteropIFD, InteropIFD_InteropIndex, k); !strings.EqualFold(exp, v) {
			t.Errorf("Expected value string: %v got %v", exp, v)
		}
	}
}

func TestExifValueString(t *testing.T) {
	//exif flash mode
	expExifFlashValues := map[uint16]string{
		0x0:  "No Flash",
		0x41: "Fired, Red-eye reduction",
	}
	var noExifFlashValue uint16 = 0x100

	//iop
	expIopInteroperabilityIndexValues := map[string]string{
		"R03": "R03 - DCF option file (Adobe RGB)",
		"THM": "THM - DCF thumbnail file",
	}
	noIopInteroperabilityIndexValue := "Whatever"

	for k, exp := range expExifFlashValues {
		v, _ := ExifValueStringErr(ExifIFD, ExifIFD_Flash, k)
		if !strings.EqualFold(exp, v) {
			t.Errorf("Expected valuestring: %s got %s", exp, v)
		}
	}
	if _, err := ExifValueStringErr(ExifIFD, ExifIFD_Flash, noExifFlashValue); err == nil {
		t.Errorf("Expcted error on undefined flash value")
	}

	for k, exp := range expIopInteroperabilityIndexValues {
		v, _ := ExifValueStringErr(InteropIFD, InteropIFD_InteropIndex, k)
		if !strings.EqualFold(exp, v) {
			t.Errorf("Expected valuestring: %s got %s", exp, v)
		}
	}
	if _, err := ExifValueStringErr(InteropIFD, InteropIFD_InteropIndex, noIopInteroperabilityIndexValue); err == nil {
		t.Errorf("Expcted error on undefined interoperatiblityindex")
	}

}

//Todo: Create an asset image that has ImageDesc and UserComment already stored
func TestExifData_GetIfdImageDescription(t *testing.T) {
	ed := getExifData(LeicaImg, t)
	if desc, err := ed.GetIfdImageDescription(); err == nil {
		t.Errorf("Did not expect any image description got %s", desc)
	} else if err != IfdTagNotFoundErr {
		t.Errorf("Expected err %v got %v", IfdTagNotFoundErr, err)
	}
	je := getJpegEditor(LeicaImg, t)
	if err := je.Exif().SetImageDescription("some description"); err != nil {
		t.Errorf("Could not set image description: %v", err)
	}
	ed = jpegEditorMD(je, t).exifData
	if _, err := ed.GetIfdImageDescription(); err != nil {
		t.Errorf("Expected description got error %v", err)
	}
}

//Todo: Create an asset image that has ImageDesc and UserComment already stored
func TestExifData_GetIfdUserComment(t *testing.T) {
	ed := getExifData(LeicaImg, t)
	if desc, err := ed.GetIfdUserComment(); err == nil {
		t.Errorf("Did not expect any user comment got %s", desc)
	} else if err != IfdTagNotFoundErr {
		t.Errorf("Expected err %v got %v", IfdTagNotFoundErr, err)
	}
	je := getJpegEditor(LeicaImg, t)
	if err := je.Exif().SetUserComment("some comment"); err != nil {
		t.Errorf("Could not set user comment: %v", err)
	}
	ed = jpegEditorMD(je, t).exifData
	if _, err := ed.GetIfdUserComment(); err != nil {
		t.Errorf("Expected user comment got error %v", err)
	}
}

func TestExifData_GpsInfo(t *testing.T) {
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
	ed := getExifData(GPSImg, t)
	gps, err := ed.GpsInfo()
	if err != nil {
		t.Fatalf("Expected IfdGps information")
	}
	//check only latitude, timestamp and altitude
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
	ed = getExifData(LeicaImg, t)
	_, err = ed.GpsInfo()
	if err == nil {
		t.Errorf("Expected no GpsInfo")
	}

}

func TestExifData_HasIfd(t *testing.T) {
	//Todo: test also IFD1 and IFD/Iop
	edGps := getExifData(GPSImg, t)
	ed := getExifData(LeicaImg, t)
	edNo := getExifData(NoExifImg, t)

	if !ed.HasIfd(RootIFD) {
		t.Errorf("Expected IfdRoot")
	}
	if !ed.HasIfd(ExifIFD) {
		t.Errorf("Expected IfdExif")
	}
	if !edGps.HasIfd(GpsIFD) {
		t.Errorf("Expected IfdGpsInfo")
	}
	//test all paths for image with no Exif
	for k, v := range IFDPaths {
		if edNo.HasIfd(k) {
			t.Errorf("Expected No Ifd got %s", v)
		}
	}
}

func TestExifData_Ifd(t *testing.T) {
	edGps := getExifData(GPSImg, t)
	ed := getExifData(LeicaImg, t)
	edNo := getExifData(NoExifImg, t)

	if edGps.Ifd(GpsIFD) == nil {
		t.Errorf("Expected Exif to contain IFD/IfdGpsInfo")
	}
	if ed.Ifd(RootIFD) == nil {
		t.Errorf("Expected Exif to contain Root IFD")
	}
	if ed.Ifd(ExifIFD) == nil {
		t.Errorf("Expected Exif to contin IFD/Exif")
	}
	if edNo.Ifd(RootIFD) != nil {
		t.Errorf("Expected No Exif got root")
	}
}

func TestExifData_IsEmpty(t *testing.T) {
	ed := getExifData(LeicaImg, t)
	edNo := getExifData(NoExifImg, t)
	if ed.IsEmpty() {
		t.Errorf("Expected exif data")
	}
	if !edNo.IsEmpty() {
		t.Errorf("Expected no exif data")
	}

}

func TestExifData_Scan(t *testing.T) {

}

func TestExifData_ScanExifDate(t *testing.T) {
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
	ed := getExifData(LeicaImg, t)
	dates := map[ExifDate]string{
		OriginalDate:  expOrigDate,
		ModifyDate:    expModifyDate,
		DigitizedDate: expDigDate,
	}
	ret := time.Time{}
	for k, v := range dates {
		if err := ed.ScanExifDate(k, &ret); err != nil {
			t.Errorf("Expected ExifDate %v got error: %v", k, err)
		} else if ret.String() != v {
			t.Errorf("Expected %s got %s", v, ret)
		}
	}
}

func TestExifData_ScanIfdExif(t *testing.T) {
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
	ed := getExifData(LeicaImg, t)
	expExposureTime := URat{Numerator: 1, Denominator: 250}
	ret := URat{}
	if err := ed.ScanIfdExif(ExifIFD_ExposureTime, &ret); err != nil {
		t.Errorf("Expected IFD_ExposureTime got error: %v", err)
	} else if ret != expExposureTime {
		t.Errorf("Expected %s got %s", expExposureTime, ret)
	}
	//now try a tag that is non existent
	ret1 := exifundefined.Tag9286UserComment{}
	if err := ed.ScanIfdExif(ExifIFD_UserComment, &ret1); err == nil {
		t.Errorf("Expected error when reading non existent root tag")
	} else if err != IfdTagNotFoundErr {
		t.Errorf("Expected error %v got %v", IfdTagNotFoundErr, err)
	}
}

func TestExifData_ScanIfdRoot(t *testing.T) {
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
	ed := getExifData(LeicaImg, t)
	expSoftware := "Adobe Photoshop Lightroom Classic 10.1 (Macintosh)"
	ret := ""
	if err := ed.ScanIfdRoot(IFD_Software, &ret); err != nil {
		t.Errorf("Expected IFD_Software got error: %v", err)
	} else if ret != expSoftware {
		t.Errorf("Expected %s got %s", expSoftware, ret)
	}
	//now try a tag that is non existent
	if err := ed.ScanIfdRoot(IFD_ImageDescription, &ret); err == nil {
		t.Errorf("Expected error when reading non existent root tag")
	} else if err != IfdTagNotFoundErr {
		t.Errorf("Expected error %v got %v", IfdTagNotFoundErr, err)
	}
}

func TestExifData_String(t *testing.T) {

}

func TestExifTagName(t *testing.T) {
	/*extracted from gps img with exiftool:
	"ExifIFD": {
	    "ExposureTime": "1/13", //0x829a
	"IFD0": {
	    "Make": "samsung", //0x010f
	"IFD1": {
	    "Compression": "JPEG (old-style)", //0x0103
	"InteropIFD": {
	    "InteropIndex": "R98 - DCF basic file (sRGB)", //0x0001
	"GPS": {
	    "GPSVersionID": "2.2.0.0", //0x0000
	*/
	errS := "Unknown Ifd"
	if tn := ExifTagName(RootIFD, ExifTag(0x010f)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected Make got %s", tn)
	}
	if tn := ExifTagName(ExifIFD, ExifTag(0x829a)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected ExposureTime got %s", tn)
	}
	if tn := ExifTagName(GpsIFD, ExifTag(0x0000)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected GPSVersion got %s", tn)
	}
	if tn := ExifTagName(RootIFD, ExifTag(0x0103)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected Compression got %s", tn)
	}
	if tn := ExifTagName(InteropIFD, ExifTag(0x0001)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected InteropIndex got %s", tn)
	}
	//now try some bad cases
	var na ExifIndex = 200
	if tn := ExifTagName(na, ExifTag(0x0001)); !strings.HasPrefix(tn, "Unknown Ifd Tag: ") {
		t.Errorf("Expected Unknown Ifd Nil got %s", tn)
	}
}

//Todo:
func TestParseIfdDateTime(t *testing.T) {

}

//Todo:
func TestTimeOffsetString(t *testing.T) {

}
