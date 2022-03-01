package metadata

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const AssetPath = "../assets/"
const LeicaImg = AssetPath + "leica.jpg"
const NoExifImg = AssetPath + "noexif.jpg"
const NikonImg = AssetPath + "nikon.jpg"
const GPSImg = AssetPath + "gps.jpg"
const NonImageFile = AssetPath + "exiftool-leica-g1.json"
const XmpFile = AssetPath + "xmp.xml"

func getAssetBytes(fname string, t *testing.T) []byte {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatalf("Could not read file: %v", err)
	}
	return b
}

func getMetaData(fname string, t *testing.T) *MetaData {
	md, err := NewMetaDataFromFile(fname)
	if err != nil {
		t.Fatalf("Could not retrieve metadata for file: %v", err)
	}
	return md
}

func TestParse(t *testing.T) {
	data := getAssetBytes(NikonImg, t)
	if _, err := NewMetaData(data); err != nil {
		t.Errorf("Could not parse image file: %v", err)
	}
	empty := []byte{}
	if _, err := NewMetaData(empty); err == nil {
		t.Errorf("Expected parse error")
	} else if err != ParseImageErr {
		t.Errorf("Expected error %v got error %v", ParseImageErr, err)
	}
	nonImageFile := getAssetBytes(NonImageFile, t)
	if _, err := NewMetaData(nonImageFile); err == nil {
		t.Errorf("Expected parse error")
	} else if err != ParseImageErr {
		t.Errorf("Expected error %v got error %v", ParseImageErr, err)
	}
}

func TestParseFile(t *testing.T) {
	if _, err := NewMetaDataFromFile(NikonImg); err != nil {
		t.Errorf("Could not parse Image file: %v", err)
	}
	if _, err := NewMetaDataFromFile(NonImageFile); err == nil {
		t.Errorf("Expected parse error")
	} else if err != ParseImageErr {
		t.Errorf("Expected error %v got error %v", ParseImageErr, err)
	}
	//non existent file
	if _, err := NewMetaDataFromFile("somefile.jpg"); err == nil {
		t.Errorf("Expected parse error")
	}
}

func TestMetaData_Xmp(t *testing.T) {

}

func TestMetaData_Exif(t *testing.T) {

}

func TestMetaData_Iptc(t *testing.T) {

}

func TestMetaData_SummaryErr(t *testing.T) {

}

func TestMetaData_Summary(t *testing.T) {

}

/*
func TestMetaDataSummary_String(t *testing.T) {
	expLeica := `Summary:{
  Title: Morning Fog
  Keywords:
  Sofware: Adobe Photoshop Lightroom Classic 10.1 (Macintosh)
  Rating: 2
  Camera Make: LEICA CAMERA AG
  Camera Model: LEICA Q2
  Lens Info: [2800/100,2800/100,392/256,761/256]
  Lens Model: SUMMILUX 1:1.7/28 ASPH.
  Lens Make:
  Focal Length: 28
  Focal Length 35MM: 28
  Max Aperture Value: 1.53125
  Flash Mode: Off, Did not fire
  Exposure Time: 1/250
  Exposure Compensation: -0.3
  Exposure Program: Manual
  Fnumber: 2.5
  ISO: 100
  ColorSpace: sRGB
  XResolution: 144/1
  YResolution: 144/1
  OriginalDate: 2020-10-27 09:34:03 +0200 +0200
  ModifyDate: 2021-01-10 17:36:45 +0100 CET
  GPSInfo:<nil>
  City:
  State/Province:
  Country:
}
`
	expGps := `Summary:{
  Title:
  Keywords:
  Sofware: GIMP 2.8.20
  Rating: 0
  Camera Make: samsung
  Camera Model: SM-N920T
  Lens Info: [0/0,0/0,0/0,0/0]
  Lens Model:
  Lens Make:
  Focal Length: 4.3
  Focal Length 35MM: 0
  Max Aperture Value: 1.85
  Flash Mode: No Flash
  Exposure Time: 1/13
  Exposure Compensation: 0
  Exposure Program: Program AE
  Fnumber: 1.9
  ISO: 200
  ColorSpace: sRGB
  XResolution: 72/1
  YResolution: 72/1
  OriginalDate: 2018-04-28 21:23:12 +0000 UTC
  ModifyDate: 2018-06-09 01:07:30 +0000 UTC
  GPSInfo:GpsInfo<LAT=(26.58667) LON=(-80.05361) ALT=(0) TIME=[2018-04-29 01:22:57 +0000 UTC]>
  City:
  State/Province:
  Country:
}
`
	expNoExif := `Summary:{
  Title:
  Keywords:
  Sofware:
  Rating: 0
  Camera Make:
  Camera Model:
  Lens Info: [0/0,0/0,0/0,0/0]
  Lens Model:
  Lens Make:
  Focal Length: NaN
  Focal Length 35MM: 0
  Max Aperture Value: NaN
  Flash Mode: No Flash
  Exposure Time: 0/0
  Exposure Compensation: NaN
  Exposure Program: Not Defined
  Fnumber: NaN
  ISO: 0
  ColorSpace: undefined
  XResolution: 0/0
  YResolution: 0/0
  OriginalDate: 0001-01-01 00:00:00 +0000 UTC
  ModifyDate: 0001-01-01 00:00:00 +0000 UTC
  GPSInfo:<nil>
  City:
  State/Province:
  Country:
}
`
	cmpMap := map[string]string{
		LeicaImg:  expLeica,
		GPSImg:    expGps,
		NoExifImg: expNoExif,
	}
	for k, v := range cmpMap {
		md := getMetaData(k, t)
		if md.Summary().String() != v {
			t.Errorf("Expected summary\n%v got\n%v", len(v), len(md.Summary().String()))
		}
	}
}
*/

func ExampleNewMetaDataFromFile() {
	md, err := NewMetaDataFromFile("../assets/leica.jpg")
	if err != nil {
		fmt.Printf("Could not open file: %v\n", err)
		return
	}
	fmt.Printf("Make: %s, Model: %s\n", md.Summary().CameraMake, md.Summary().CameraModel)

	//Output: Make: LEICA CAMERA AG, Model: LEICA Q2
}
