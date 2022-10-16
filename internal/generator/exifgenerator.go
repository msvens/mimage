package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type etExifTag struct {
	Id          uint16
	Name        string
	Writable    string
	Mandatory   bool
	Fmt         string
	Ifd         string
	Count       int
	SubDir      bool
	DirName     string
	Offset      bool
	OffsetPair  uint16
	Permanent   bool
	Protected   bool
	Description string
	Notes       string
	Values      map[string]string
}

/*
var exifTypeMapping = map[string]string{
	"double":      "ExifDouble",
	"float":       "ExifFloat",
	"int16s":      "ExifInt16",
	"int16u":      "ExifUint16",
	"int32s":      "ExifInt32",
	"int32u":      "ExifUint32",
	"int8u":       "ExifUint8",
	"rational64s": "ExifRational",
	"rational64u": "ExifUrational",
	"string":      "ExifString",
	"undef":       "ExifUndef",
}*/

const rawMain = "main"
const rawGps = "gps"

const exifTypesSrc = `
type ExifTag uint16
type ExifTagType uint8
type ExifIndex int
type ExifIndexTag struct {
  Index ExifIndex
  Tag ExifTag
}
`

const exifIndexSrc = `
const (
	RootIFD ExifIndex = iota
	ExifIFD
	GpsIFD
	InteropIFD
	ThumbnailIFD
)

var IFDPaths = map[ExifIndex]string{
	RootIFD:      "IFD",
	ExifIFD:      "IFD/Exif",
	GpsIFD:   "IFD/GPSInfo",
	InteropIFD:       "IFD/Exif/Iop",
	ThumbnailIFD: "IFD1",
}
`

const exifTypeConstSrc = `
const(
  ExifString ExifTagType = iota
  ExifUint8
  ExifUint16
  ExifUint32
  ExifInt16
  ExifInt32
  ExifRational
  ExifUrational
  ExifFloat
  ExifDouble
  ExifUndef
)
`

const exifTagDescSrc = `
type ExifTagDesc struct {
  Id        ExifTag  ` + "`json:\"id\"`" + `
  Name      string ` + "`json:\"name\"`" + `
  Type  ExifTagType ` + "`json:\"type\"`" + `
  Mandatory  bool ` + "`json:\"mandatory\"`" + `
  Ifd  ExifIndex ` + "`json:\"ifd\"`" + `
  Count int ` + "`json:\"count\"`" + `
  Offset bool ` + "`json:\"offset\"`" + `
  OffsetPair ExifTag ` + "`json:\"offsetPair\"`" + `
  Permanent bool ` + "`json:\"permanent\"`" + `
  Protected bool ` + "`json:\"protected\"`" + `
  Values interface{} ` + "`json:\"values\"`" + `
}
`

// GenerateMasterExifJSON aligns types and adds type/description info to tags from exiv2 json. Sorts the file based on TagId
func GenerateMasterExifJSON() error {
	b, err := os.ReadFile("assets/exiftool-exiftags.json")
	if err != nil {
		return err
	}
	exifMap := make(map[string][]*etExifTag)
	err = json.Unmarshal(b, &exifMap)
	if err != nil {
		return err
	}
	b, err = os.ReadFile("assets/exiv2-exiftags.json")
	if err != nil {
		return err
	}
	exivMap := make(map[string][]rawExifTag)
	err = json.Unmarshal(b, &exivMap)
	if err != nil {
		return err
	}

	//sort
	sort.Slice(exifMap[rawMain], func(i, j int) bool {
		return exifMap[rawMain][i].Id < exifMap[rawMain][j].Id
	})
	sort.Slice(exifMap[rawGps], func(i, j int) bool {
		return exifMap[rawGps][i].Id < exifMap[rawGps][j].Id
	})

	//create exiv2 maps
	exivIdMaps := map[string]map[uint16]rawExifTag{}
	for k, v := range exivMap {
		m := map[uint16]rawExifTag{}
		for _, t := range v {
			m[t.Id] = t
		}
		exivIdMaps[k] = m
	}

	//align information in exiftool json and handle duplicate names (we know of one at least)
	for _, v := range exifMap {
		tagNames := map[string]bool{}
		for _, t := range v {
			t.Name = alignName(t.Name)
			t.Ifd = alignIFD(t.Ifd)
			t.Writable = alignType(*t, exivIdMaps)
			t.Description = alignDesc(*t, exivIdMaps)
			if _, found := tagNames[t.Ifd+t.Name]; found {
				//fmt.Println("found duplicate")
				t.Name = fmt.Sprintf("%s_%#04x", t.Name, t.Id)
				//fmt.Println("found duplicate: ", t.Name)
			} else {
				tagNames[t.Ifd+t.Name] = true
			}
		}
	}

	//write
	outBytes, err := json.MarshalIndent(exifMap, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("./assets/master-exiftags.json", outBytes, 0644)

}

func alignDesc(tag etExifTag, exivMaps map[string]map[uint16]rawExifTag) string {
	exivType, found := exivMaps[tag.Ifd][tag.Id]
	if found && tag.Description == "" {
		return exivType.Description
	}
	return tag.Description
}

func alignType(tag etExifTag, exivMaps map[string]map[uint16]rawExifTag) string {
	t := tag.Writable
	if t == "undef" && tag.Fmt != "undef" {
		t = tag.Fmt
	}
	ret := "ExifUndef"
	switch t {
	case "string":
		return "ExifString"
	case "int8u":
		return "ExifUint8"
	case "int16u":
		return "ExifUint16"
	case "int32u":
		return "ExifUint32"
	case "int16s":
		return "ExifInt16"
	case "int32s":
		return "ExifInt32"
	case "rational64u":
		return "ExifUrational"
	case "rational64s":
		return "ExifRational"
	case "float":
		return "ExifFloat"
	case "double":
		return "ExifDouble"
	}
	if exivType, found := exivMaps[tag.Ifd][tag.Id]; found {
		if ret != exivType.TypeName {
			ret = exivType.TypeName
		}
	}

	return ret
}

func alignIFD(rawIfd string) string {
	if rawIfd == "ExifIFD" {
		return rawIfd
	} else if rawIfd == "InteropIFD" {
		return rawIfd
	} else if rawIfd == "GpsIFD" {
		return rawIfd
	} else {
		return "RootIFD"
	}
}

func alignName(rawName string) string {
	return strings.Replace(rawName, "-", "", -1)
}

// GenerateExifTagsFromMasterExifJSON generate exif sources from exif json file
func GenerateExifTagsFromMasterExifJSON() error {
	raw, err := readMasterExifJSON()
	if err != nil {
		fmt.Println(err)
		return err
	}
	sb := strings.Builder{}
	sb.WriteString(`package metadata
//Do not edit! This is an automatically generated file (see generator.GenerateExifTagsFromExifTool()).
//This file was generated based on https://github.com/exiftool/exiftool/blob/master/lib/Image/ExifTool/EXIF.pm
`)
	sb.WriteString(exifTypesSrc)
	sb.WriteString(exifIndexSrc)
	sb.WriteString(exifTypeConstSrc)
	sb.WriteString(exifTagDescSrc)
	_ = generateExifConstants(raw, &sb)
	_ = generateExifTagDescriptions(raw, &sb)

	err = os.WriteFile("./metadata/genexif.go", []byte(sb.String()), 0644)
	return err
}

func generateExifTagDescriptions(raw map[string][]etExifTag, sb *strings.Builder) error {
	sb.WriteString("//Exif Tag Descriptions\n")
	sb.WriteString("var ExifTagDescriptions = map[ExifIndexTag]ExifTagDesc{\n")

	order := []string{rawMain, rawGps}
	for _, v := range order {
		for _, t := range raw[v] {
			valueMap := generateExifValueMap(t.Writable, t)
			indexTag := fmt.Sprintf("ExifIndexTag{%s,%#04x}", t.Ifd, t.Id)
			descFmt := `ExifTagDesc{
  Id: %#04x,
  Name: "%s",
  Type: %v,
  Mandatory: %v,
  Ifd: %s,
  Count: %v,
  Offset: %v,
  OffsetPair: %v,
  Permanent: %v,
  Protected: %v,
  Values: %s,
}`
			adjustedCnt := t.Count
			if adjustedCnt == 0 {
				if t.Writable == "ExifString" || t.Writable == "ExifUndef" {
					adjustedCnt = -1
				} else {
					adjustedCnt = 1
				}
			}
			tagDesc := fmt.Sprintf(descFmt, t.Id, fixTagName(t.Name), t.Writable, t.Mandatory, t.Ifd, adjustedCnt, t.Offset, t.OffsetPair, t.Permanent, t.Protected, valueMap)
			sb.WriteString(fmt.Sprintf("%s: %s,\n", indexTag, tagDesc))
		}
	}
	sb.WriteString("}\n")
	return nil
}

func generateExifConstants(raw map[string][]etExifTag, sb *strings.Builder) error {
	tags := raw[rawMain]
	//ifds := map[string]bool{}

	//generate ifd constants
	sb.WriteString("//IFD Tag Ids (includes all IFD, IFD1, etc tags)\nconst(\n")
	for _, t := range tags {
		if t.Ifd == "RootIFD" {
			sb.WriteString(fmt.Sprintf("  IFD_%s ExifTag = %#04x\n", fixTagName(t.Name), t.Id))
		}
	}
	sb.WriteString(")\n")

	//generate exififd constants
	sb.WriteString("//ExifIFD Tag Ids\nconst(\n")
	for _, t := range tags {
		if t.Ifd == "ExifIFD" {
			sb.WriteString(fmt.Sprintf("  ExifIFD_%s ExifTag = %#04x\n", t.Name, t.Id))
		}
	}
	sb.WriteString(")\n")

	//generate exifInterop constants
	sb.WriteString("//InteropIFD Tag Ids\nconst(\n")
	for _, t := range tags {
		if t.Ifd == "InteropIFD" {
			sb.WriteString(fmt.Sprintf("  InteropIFD_%s ExifTag = %#04x\n", t.Name, t.Id))
		}
	}
	sb.WriteString(")\n")

	//generate gps constants
	tags = raw[rawGps]
	sb.WriteString("//GpsIFD Tag Ids\nconst(\n")
	for _, t := range tags {
		sb.WriteString(fmt.Sprintf("  GpsIFD_%s ExifTag = %#04x\n", t.Name, t.Id))
	}
	sb.WriteString(")\n")

	return nil
}

//gocyclo:ignore
func generateExifValueMap(exifType string, tag etExifTag) string {
	if len(tag.Values) == 0 || exifType == "ExifUndef" {
		return "nil"
	} else if tag.Count != 0 { //we dont generate a value map for complex values
		return "nil"
	} else if exifType == "ExifUrational" || exifType == "ExifRational" { //we dont generate a value map for rationals
		return "nil"
	}

	checkNumber := func(s string, bitSize int, float bool, unsigned bool) error {
		base := 10
		if strings.HasPrefix(s, "0x") {
			s = strings.TrimPrefix(s, "0x")
			base = 16
		}
		var err error
		if float {
			_, err = strconv.ParseFloat(s, bitSize)
		} else if unsigned {
			_, err = strconv.ParseUint(s, base, bitSize)
		} else {
			_, err = strconv.ParseInt(s, base, bitSize)
		}
		return err
	}

	buff := strings.Builder{}
	switch exifType {
	case "ExifString":
		buff.WriteString("map[string]string{\n")
		for k, v := range tag.Values {
			buff.WriteString(fmt.Sprintf("    \"%s\": \"%s\",\n", k, v))
		}
		buff.WriteString("  }")
	case "ExifFloat":
		buff.WriteString("map[float]string{\n")
		for k, v := range tag.Values {
			if err := checkNumber(k, 32, true, false); err != nil {
				fmt.Println("could not parse value key:", k)
			} else {
				buff.WriteString(fmt.Sprintf("    %s: \"%s\",\n", k, v))
			}
		}
		buff.WriteString("  }")
	case "ExifDouble":
		buff.WriteString("map[double]string{\n")
		for k, v := range tag.Values {
			if err := checkNumber(k, 64, true, false); err != nil {
				fmt.Println("could not parse value key:", k)
			} else {
				buff.WriteString(fmt.Sprintf("    %s: \"%s\",\n", k, v))
			}
		}
		buff.WriteString("  }")
	case "ExifUint8":
		buff.WriteString("map[uint8]string{\n")
		for k, v := range tag.Values {
			if err := checkNumber(k, 8, false, true); err != nil {
				fmt.Println("could not parse value key:", k)
			} else {
				buff.WriteString(fmt.Sprintf("    %s: \"%s\",\n", k, v))
			}
		}
		buff.WriteString("  }")
	case "ExifUint16":
		buff.WriteString("map[uint16]string{\n")
		for k, v := range tag.Values {
			if err := checkNumber(k, 16, false, true); err != nil {
				fmt.Println("could not parse value key:", k)
			} else {
				buff.WriteString(fmt.Sprintf("    %s: \"%s\",\n", k, v))
			}
		}
		buff.WriteString("  }")
	case "ExifUint32":
		buff.WriteString("map[uint32]string{\n")
		for k, v := range tag.Values {
			if err := checkNumber(k, 32, false, true); err != nil {
				fmt.Println("could not parse value key:", k)
			} else {
				buff.WriteString(fmt.Sprintf("    %s: \"%s\",\n", k, v))
			}
		}
		buff.WriteString("  }")
	case "ExifInt16":
		buff.WriteString("map[int16]string{\n")
		for k, v := range tag.Values {
			if err := checkNumber(k, 16, false, false); err != nil {
				fmt.Println("could not parse value key:", k)
			} else {
				buff.WriteString(fmt.Sprintf("    %s: \"%s\",\n", k, v))
			}
		}
		buff.WriteString("  }")
	case "ExifInt32":
		buff.WriteString("map[int32]string{\n")
		for k, v := range tag.Values {
			if err := checkNumber(k, 32, false, false); err != nil {
				fmt.Println("could not parse value key:", k)
			} else {
				buff.WriteString(fmt.Sprintf("    %s: \"%s\",\n", k, v))
			}
		}
		buff.WriteString("  }")
	default:
		fmt.Println("Type not found:", exifType)
		return "nil"
	}
	return buff.String()
}

func readMasterExifJSON() (map[string][]etExifTag, error) {
	b, err := os.ReadFile("assets/master-exiftags.json")
	if err != nil {
		return nil, err
	}
	exifMap := make(map[string][]etExifTag)
	err = json.Unmarshal(b, &exifMap)
	if err != nil {
		return nil, err
	}
	//sort the slices on id
	sort.Slice(exifMap[rawMain], func(i, j int) bool {
		return exifMap[rawMain][i].Id < exifMap[rawMain][j].Id
	})
	sort.Slice(exifMap[rawGps], func(i, j int) bool {
		return exifMap[rawGps][i].Id < exifMap[rawGps][j].Id
	})
	return exifMap, nil

}
