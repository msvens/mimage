package metadata

import (
	"bytes"
	"testing"
)

func getIptcData(fname string, t *testing.T) *IptcData {
	bytes := getAssetBytes(fname, t)
	segments, err := parseJpegBytes(bytes)
	if err != nil {
		t.Fatalf("Could not retrieve segments for file: %v", err)
	}
	iptcData, _ := NewIptcData(segments)
	return iptcData

}

func TestIptcData_IsEmpty(t *testing.T) {
	iptc := getIptcData(LeicaImg, t)
	noIptc := getIptcData(NoExifImg, t)
	if iptc.IsEmpty() {
		t.Errorf("Expected iptc data")
	}
	if !noIptc.IsEmpty() {
		t.Errorf("Expected no iptc data")
	}
}

func TestIptcData_RawIptc(t *testing.T) {
	iptc := getIptcData(LeicaImg, t)
	noIptc := getIptcData(NoExifImg, t)
	if len(iptc.raw) != len(iptc.RawIptc()) {
		t.Errorf("Expected the same iptc length got: %v, %v", len(iptc.raw), len(iptc.RawIptc()))
	}
	if len(noIptc.raw) != len(noIptc.RawIptc()) {
		t.Errorf("Expected the same noiptc length got: %v, %v", len(noIptc.raw), len(noIptc.RawIptc()))
	}
}

func TestIptcData_Scan(t *testing.T) {

}

func TestIptcData_ScanApplicationTag(t *testing.T) {
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
	iptcData := getIptcData(LeicaImg, t)
	expObjectName := "Morning Fog"
	ret := ""
	if err := iptcData.ScanApplication(Application_ObjectName, &ret); err != nil {
		t.Errorf("Expected IPTC Application Object Name to be set got: %v", err)
	} else if ret != expObjectName {
		t.Errorf("Expected %s got %s", expObjectName, ret)
	}
	//non existent tag
	if err := iptcData.ScanApplication(Application_Keywords, &ret); err == nil {
		t.Errorf("Expected error got %s", ret)
	} else if err != IptcTagNotFoundErr {
		t.Errorf("Expected error %v got error %v", IptcTagNotFoundErr, err)
	}
	//md.PrintIptc()
}

func TestIptcData_ScanIptcEnvelopTag(t *testing.T) {
	/*exiftool extract IPTC application from Leica Img
	"IPTC": {
	    "CodedCharacterSet": "UTF8",
	  },
	*/
	iptcData := getIptcData(LeicaImg, t)
	//expObjectName := "Morning Fog"
	expectedCharSet := []byte{27, 37, 71} //ESC%G (which is UTF
	ret := []byte{}
	if err := iptcData.ScanEnvelope(Envelope_CodedCharacterSet, &ret); err != nil {
		t.Errorf("Expected CodedCharacterSet got error %v", err)
	} else if bytes.Compare(ret, expectedCharSet) != 0 {
		t.Errorf("Expected %v got %v", expectedCharSet, ret)
	}
	//non existent tag
	if err := iptcData.ScanEnvelope(Envelope_Destination, &ret); err == nil {
		t.Errorf("Expected error for non existing tag")
	} else if err != IptcTagNotFoundErr {
		t.Errorf("Expected error %v got error %v", IptcTagNotFoundErr, err)
	}

}
