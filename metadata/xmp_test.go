package metadata

import (
	"testing"
)

func getXmpData(fname string, t *testing.T) XmpData {
	bytes := getAssetBytes(fname, t)
	segments, err := parseJpegBytes(bytes)
	if err != nil {
		t.Fatalf("Could not retrieve segments for file: %v", err)
	}
	xmpData, _ := NewXmpData(segments)
	return xmpData

}

func TestXmpData_IsEmpty(t *testing.T) {
	xmp := getXmpData(LeicaImg, t)
	noXmp := getXmpData(NoExifImg, t)
	if xmp.IsEmpty() {
		t.Errorf("Expected XmpEditor data")
	}
	if !noXmp.IsEmpty() {
		t.Errorf("Expected no XmpEditor data")
	}

}

func TestXmpData_String(t *testing.T) {

}
