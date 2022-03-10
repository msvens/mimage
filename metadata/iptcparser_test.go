package metadata

import (
	"bytes"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"github.com/msvens/mimage/photoshop"
	"reflect"
	"testing"
)

func getIptcDataFromImage(fname string, t *testing.T) []byte {
	jpegParser := jpegstructure.NewJpegMediaParser()
	ec, err := jpegParser.ParseFile(fname)
	if err != nil {
		t.Fatalf("Could not parse jpeg file: %v", err)
		return nil
	}
	sl := ec.(*jpegstructure.SegmentList)
	_, res, err := photoshop.ParseJpeg(sl)
	if err != nil {
		t.Fatalf("Could not parse segmentlist: %v", err)
	}
	iptcData, ok := res[photoshop.IptcId]
	if ok {
		return iptcData.Data
	}
	t.Fatalf("Photoshop resources did not contain IPTC data")
	return nil

}

func checkIptcValues(actual map[IptcRecordTag]IptcRecordDataset, t *testing.T) {
	//checks against NikonImg

	contains := func(in []string, s string) bool {
		for _, ss := range in {
			if ss == s {
				return true
			}
		}
		return false
	}
	var ds IptcRecordDataset
	ds = actual[IptcRecordTag{Record: 2, Tag: 60}]
	if ds.Data != "104050" {
		t.Errorf("Expected 104050 got %v", ds.Data)
	}
	ds = actual[IptcRecordTag{Record: 2, Tag: 62}]
	if ds.Data != "20100811" {
		t.Errorf("Expected 20100811 got %v", ds.Data)
	}
	ds = actual[IptcRecordTag{Record: 2, Tag: 0}]
	if ds.Data != uint16(4) {
		t.Errorf("Expected 4 got %v", ds.Data)
	}
	ds = actual[IptcRecordTag{Record: 2, Tag: 25}]
	actKeywords := ds.Data.([]string)
	if len(actKeywords) != 4 {
		t.Errorf("Expected 4 keywords got %v keywords", len(actKeywords))
	}
	for _, kw := range []string{"martin", "portland", "semester", "vintage"} {
		if !contains(actKeywords, kw) {
			t.Errorf("Expected keyword %v", kw)
		}
	}
	ds = actual[IptcRecordTag{Record: 1, Tag: 90}]
	if !reflect.DeepEqual(ds.Data, iptcUtfCharSet) {
		t.Errorf("Expected UTF charset")
	}
}

func TestDecodeIptc(t *testing.T) {
	iptcData := getIptcDataFromImage(NikonImg, t)
	iptc, err := DecodeIptc(bytes.NewReader(iptcData))
	if err != nil {
		t.Fatalf("Could not decode iptc data: %v", err)
	}
	checkIptcValues(iptc, t)
}

func TestEncodeIptc(t *testing.T) {
	iptcData := getIptcDataFromImage(NikonImg, t)
	iptc, err := DecodeIptc(bytes.NewReader(iptcData))
	if err != nil {
		t.Fatalf("Could not decode iptc data: %v", err)
	}
	//change title:
	objectName := IptcRecordTag{Record: 2, Tag: 5}
	ds := iptc[objectName]
	ds.Data = "New Title"
	iptc[objectName] = ds
	out := &bytes.Buffer{}
	if err = EncodeIptc(out, iptc); err != nil {
		t.Fatalf("Could not encode iptc data: %v", err)
	}
	//now decode
	iptc, err = DecodeIptc(bytes.NewReader(out.Bytes()))
	if err != nil {
		t.Fatalf("Could not decode encoded data: %v", err)
	}
	checkIptcValues(iptc, t) //should still hold
	if iptc[objectName].Data != "New Title" {
		t.Errorf("Expected New Title got %v", iptc[objectName].Data)
	}

}
