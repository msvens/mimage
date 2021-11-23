package metadata

import (
	"strings"
	"testing"
)

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
	md := getMetaData(GPSImg, t)
	errS := "Unknown Ifd"
	if tn := ExifTagName(md.IfdRoot(), uint16(0x010f)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected Make got %s", tn)
	}
	if tn := ExifTagName(md.IfdExif(), uint16(0x829a)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected ExposureTime got %s", tn)
	}
	if tn := ExifTagName(md.IfdGPS(), uint16(0x0000)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected GPSVersion got %s", tn)
	}
	if tn := ExifTagName(md.GetIfd(ThumbNailPath), uint16(0x0103)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected Compression got %s", tn)
	}
	if tn := ExifTagName(md.GetIfd(IopPath), uint16(0x0001)); strings.HasPrefix(tn, errS) {
		t.Errorf("Expected InteropIndex got %s", tn)
	}
	//now try some bad cases
	if tn := ExifTagName(nil, uint16(0x0001)); !strings.HasPrefix(tn, "Unknown Ifd Nil") {
		t.Errorf("Expected Unknown Ifd Nil")
	}

}

func TestParseIfdDateTime(t *testing.T) {

}

func TestScanIfdTag(t *testing.T) {

}

func TestTimeOffsetString(t *testing.T) {

}
