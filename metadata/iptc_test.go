package metadata

import (
	"reflect"
	"strings"
	"testing"
	"time"
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

func TestIptcData_GetDate(t *testing.T) {
	tests := []*IptcData{getIptcData(LeicaImg, t), getIptcData(NikonImg, t)}
	for _, ipd := range tests {
		var dt time.Time
		if dt = ipd.GetDate(DateCreated); dt.IsZero() {
			t.Errorf("Expected non zero DateCreated")
		}
		if dt = ipd.GetDate(DigitalCreationDate); dt.IsZero() {
			t.Errorf("Expected non zero DigitalCreationDate")
		}
		if dt = ipd.GetDate(ExpirationDate); !dt.IsZero() {
			t.Errorf("Expected zero ExpirationDate")
		}
	}
	//no exif image
	ipd := getIptcData(NoExifImg, t)
	if dt := ipd.GetDate(DateCreated); !dt.IsZero() {
		t.Errorf("Expected zero Date Created")
	}

}

func TestIptcData_GetKeywords(t *testing.T) {
	expKeywords := []string{"martin", "portland", "semester", "vintage"}
	ipd := getIptcData(NikonImg, t)
	if !reflect.DeepEqual(ipd.GetKeywords(), expKeywords) {
		t.Errorf("Expected %v got %v", expKeywords, ipd.GetKeywords())
	}
	ipd = getIptcData(NoExifImg, t)
	if len(ipd.GetKeywords()) != 0 {
		t.Errorf("Expected empty keywords got %v", ipd.GetKeywords())
	}

}

func TestIptcData_GetTitle(t *testing.T) {
	expTitle := "Morning Fog"
	ipd := getIptcData(LeicaImg, t)
	if ipd.GetTitle() != expTitle {
		t.Errorf("Expected %v got %v", expTitle, ipd.GetTitle())
	}
	ipd = getIptcData(NoExifImg, t)
	if ipd.GetTitle() != "" {
		t.Errorf("Expected empty title got %v", ipd.GetTitle())
	}
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

func TestIptcData_ScanApplication(t *testing.T) {
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
	if err := iptcData.ScanApplication(IPTCApplication_ObjectName, &ret); err != nil {
		t.Errorf("Expected IPTC Application Object Name to be set got: %v", err)
	} else if ret != expObjectName {
		t.Errorf("Expected %s got %s", expObjectName, ret)
	}
	//non existent tag
	if err := iptcData.ScanApplication(IPTCApplication_Keywords, &ret); err == nil {
		t.Errorf("Expected error got %s", ret)
	} else if err != IptcTagNotFoundErr {
		t.Errorf("Expected error %v got error %v", IptcTagNotFoundErr, err)
	}
	//md.PrintIptc()
}

func TestIptcData_ScanDate(t *testing.T) {
	tests := []*IptcData{getIptcData(LeicaImg, t), getIptcData(NikonImg, t)}
	dt := time.Time{}
	var err error
	for _, ipd := range tests {
		if err = ipd.ScanDate(DateCreated, &dt); err != nil {
			t.Errorf("Got error: %v", err)
		} else if dt.IsZero() {
			t.Errorf("Expected non zero DateCreated")
		}
		if err = ipd.ScanDate(DigitalCreationDate, &dt); err != nil {
			t.Errorf("Got error: %v", err)
		} else if dt.IsZero() {
			t.Errorf("Expected non zero DigitalCreationDate")
		}
		if err = ipd.ScanDate(ExpirationDate, &dt); err == nil {
			t.Errorf("Expected error when scanning non existing date")
		}
	}
	//no exif image
	ipd := getIptcData(NoExifImg, t)
	if err = ipd.ScanDate(DateCreated, &dt); err == nil {
		t.Errorf("Expected error when scanning non existing date")
	}
}

func TestIptcData_ScanEnvelope(t *testing.T) {
	/*exiftool extract IPTC application from Leica Img
	"IPTC": {
	    "CodedCharacterSet": "UTF8",
	  },
	*/
	iptcData := getIptcData(LeicaImg, t)
	//expObjectName := "Morning Fog"
	ret := ""
	if err := iptcData.ScanEnvelope(IPTCEnvelope_CodedCharacterSet, &ret); err != nil {
		t.Errorf("Expected CodedCharacterSet got error %v", err)
	} else if strings.Compare(ret, iptcUtfCharSet) != 0 {
		t.Errorf("Expected %v got %v", iptcUtfCharSet, ret)
	}
	//non existent tag
	if err := iptcData.ScanEnvelope(IPTCEnvelope_Destination, &ret); err == nil {
		t.Errorf("Expected error for non existing tag")
	} else if err != IptcTagNotFoundErr {
		t.Errorf("Expected error %v got error %v", IptcTagNotFoundErr, err)
	}

}
