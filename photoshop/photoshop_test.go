package photoshop

import (
	"bytes"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"io/ioutil"
	"testing"
)

const AssetPath = "../assets/"
const LeicaData = AssetPath + "leicaPhotoshopResources.data"
const LeicaDataWithPrefix = AssetPath + "leicaPhotoshopResourcesWithPrefix.data"
const Leica = AssetPath + "leica.jpg"
const NoExif = AssetPath + "noexif.jpg"

/* Layout of the Resources we expect
RECORD-TYPE=[8BIM] IMAGE-RESOURCE-ID=[0x03ed] NAME=[] DATA-SIZE=(16)
RECORD-TYPE=[8BIM] IMAGE-RESOURCE-ID=[0x0404] NAME=[] DATA-SIZE=(89)
RECORD-TYPE=[8BIM] IMAGE-RESOURCE-ID=[0x040c] NAME=[] DATA-SIZE=(10618)
RECORD-TYPE=[8BIM] IMAGE-RESOURCE-ID=[0x0425] NAME=[] DATA-SIZE=(16)
*/

var expectedResources = map[uint16]int{
	0x03ed: 16,
	0x0404: 89,
	0x040c: 10618,
	0x0425: 16,
}

/*const LeicaImg = AssetPath + "leica.jpg"
const NoExifImg = AssetPath + "noexif.jpg"
const NikonImg = AssetPath + "nikon.jpg"
const CanonImg = AssetPath + "canon.jpg"
const GPSImg = AssetPath + "gps.jpg"
const NonImageFile = AssetPath + "exiftool-leica-g1.json"
const XmpFile = AssetPath + "xmp.xml"*/

func checkExpectedResources(actual map[uint16]ImageResource, t *testing.T) {
	for k, v := range expectedResources {
		if r, ok := actual[k]; !ok {
			t.Errorf("Could not find resource: %#04x", k)
		} else if len(r.Data) != v {
			t.Errorf("Expected %#04x to have data length %v got %v", k, v, len(r.Data))
		}
	}
}

func getAssetBytes(fname string, t *testing.T) []byte {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatalf("Could not read file: %v", err)
	}
	return b
}

func getSegments(fname string, t *testing.T) *jpegstructure.SegmentList {
	parser := jpegstructure.NewJpegMediaParser()
	ec, err := parser.ParseFile(fname)
	if err != nil {
		t.Fatalf("Could not parse jpeg file: %v", err)
		return nil
	}
	return ec.(*jpegstructure.SegmentList)
}

func getTestData(t *testing.T) map[bool][]byte {
	return map[bool][]byte{
		false: getAssetBytes(LeicaData, t),
		true:  getAssetBytes(LeicaDataWithPrefix, t),
	}
}

func TestDecode(t *testing.T) {
	//first decode without prefix
	tests := getTestData(t)
	for prefix, data := range tests {
		buff := bytes.NewReader(data)
		if resources, err := Decode(buff, prefix); err != nil {
			t.Fatalf("Could not decode data with prefix %v got error: %v", prefix, err)
		} else {
			checkExpectedResources(resources, t)
		}

	}

	//decoding a buffer with prefix as a non prefixed buffer
	buff := bytes.NewBuffer(tests[true])
	_, err := Decode(buff, false)
	if err == nil {
		t.Fatalf("Expected error got nil")
	}

	//decoding an empty buffer
	_, err = Decode(bytes.NewBuffer([]byte{}), false)
	if err == nil {
		t.Fatalf("Expected error got nil")
	}
}

func TestEncode(t *testing.T) {
	//first encode without prefix
	tests := getTestData(t)

	for prefix, data := range tests {
		in := bytes.NewReader(data)
		resources, _ := Decode(in, prefix)
		outBuff := &bytes.Buffer{}
		if err := Encode(outBuff, resources, prefix); err != nil {
			t.Fatalf("Could not encode image resoure with prefix: %v got error: %v", prefix, err)
		}
		out := outBuff.Bytes()
		if len(data) != len(out) {
			t.Fatalf("Expected length %v got length %v", len(data), len(out))
		}
		//now Decode
		in = bytes.NewReader(out)
		if newResources, err := Decode(in, prefix); err != nil {
			t.Fatalf("Could not decode encoded resources: %v", err)
		} else {
			checkExpectedResources(newResources, t)
		}
	}
	outBuff := &bytes.Buffer{}
	err := Encode(outBuff, map[uint16]ImageResource{}, false)
	if err == nil {
		t.Fatalf("Expected error got nil")
	}
}

func TestMarshal(t *testing.T) {
	tests := getTestData(t)
	for prefix, data := range tests {
		m := map[uint16]ImageResource{}
		if err := Unmarshal(data, prefix, &m); err != nil {
			t.Fatalf("Could not unmarshal prefixed %v data got error: %v", prefix, err)
		}
		if out, err := Marshal(m, prefix); err != nil {
			t.Fatalf("Could not marshal prefixed %v data got error: %v", prefix, err)
		} else if len(out) != len(data) {
			t.Errorf("Expected data length %v got %v", len(data), len(out))
		}
	}
	//now unmarshal empty data, expect error
	m := map[uint16]ImageResource{}
	if _, err := Marshal(m, false); err == nil {
		t.Errorf("Expected error when marshaling empty resources")
	}
	if _, err := Marshal(m, true); err == nil {
		t.Errorf("Expected error when marshaling with prefix empty resources")
	}

}

func TestParseJpeg(t *testing.T) {
	sl := getSegments(Leica, t)
	if _, ret, err := ParseJpeg(sl); err != nil {
		t.Errorf("Could not parse segments: %v", err)
	} else {
		checkExpectedResources(ret, t)
	}
	sl = getSegments(NoExif, t)
	if _, _, err := ParseJpeg(sl); err != ErrNoPhotoshopBlock {
		t.Errorf("Expected ErrNoPhotoshopBlock got %v", err)
	}
}

func TestUnmarshal(t *testing.T) {
	tests := getTestData(t)
	for prefix, data := range tests {
		m := map[uint16]ImageResource{}
		if err := Unmarshal(data, prefix, &m); err != nil {
			t.Fatalf("Could not unmarshal data with prefix %v got error: %v", prefix, err)
		} else {
			checkExpectedResources(m, t)
		}
		//check for error on corrupted data
		if err := Unmarshal(data[:len(data)/2], prefix, &m); err == nil {
			t.Fatalf("Expected error on unmarshal corrupted data with prefix %v", prefix)
		}
	}
	//check empty data errors
	m := map[uint16]ImageResource{}
	if err := Unmarshal([]byte{}, false, &m); err == nil {
		t.Errorf("Expected error when unmarshaling an empty byte slice")
	}
	if err := Unmarshal([]byte{}, true, &m); err == nil {
		t.Errorf("Expected error when unmarshaling with prefix an empty byte slice")
	}

}
