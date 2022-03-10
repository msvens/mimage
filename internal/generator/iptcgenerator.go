package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type etIptcTag struct {
	Id        uint8
	Name      string
	Fmt       string
	Binary    bool
	Mandatory bool
	Writable  bool
	Flags     string
	Notes     string
	Values    map[string]string
}

type etIptcRecord struct {
	Id   uint8
	Name string
	Tags []etIptcTag
}

//var fmtParser, _ = regexp.Compile(`^\s*(\w+).*?\[?(\d*),?(\d*)\]?.*$`)
var typeRegExp, _ = regexp.Compile(`^\s*(\w+).*?\[?(\d*),?(\d*)\]?.*$`)

func fixTagName(name string) string {
	return strings.Replace(name, "-", "", -1)
}

func GenerateIptcTagsFromExifTool() error {
	raw, err := readExifToolIptcJSON()
	if err != nil {
		return err
	}
	sb := strings.Builder{}
	sb.WriteString(`package metadata
//Do not edit! This is an automatically generated file (see generator.GenerateIptcTagsFromExifTool()).
//This file was generated based on https://github.com/exiftool/exiftool/blob/master/lib/Image/ExifTool/IPTC.pm

`)
	generateHeaderSource(raw, &sb)
	generateTagConstants(raw, &sb)
	generateTagMap(raw, &sb)

	err = ioutil.WriteFile("./metadata/geniptc.go", []byte(sb.String()), 0644)

	/*
		for _,v := range raw {
			for _,tag := range v.Tags {
				if tag.Fmt != "" {
					fmt.Println(tag.Fmt)
				}
			}
		}
	*/
	return err
}

func generateHeaderSource(raw map[string]etIptcRecord, sb *strings.Builder) {
	sb.WriteString(`
type IptcTagType uint8
type IptcRecord uint8
type IptcTag uint8

type IptcRecordTag struct {
  Record IptcRecord
  Tag IptcTag
}

type IptcTagDesc struct {
  Id        IptcTag  ` + "`json:\"id\"`" + `
  Name      string ` + "`json:\"name\"`" + `
  Type  IptcTagType ` + "`json:\"type\"`" + `
  MinLength int    ` + "`json:\"minLength\"`" + `
  MaxLength int   ` + "`json:\"maxLength\"`" + `
  Mandatory bool ` + "`json:\"mandatory\"`" + `
  Repeatable bool ` + "`json:\"repeatable\"`" + `
  Writable bool ` + "`json:\"writable\"`" + `
  Values interface{} ` + "`json:\"values\"`" + `
}

const(
  IptcString IptcTagType = iota
  IptcDigits
  IptcUint8
  IptcUint16
  IptcUint32
  IptcUndef
)

const(
`)
	for k, v := range raw {
		sb.WriteString(fmt.Sprintf("  %s IptcRecord = %v\n", k, v.Id))
	}
	sb.WriteString(")\n\n")
	sb.WriteString("var IptcRecordName = map[IptcRecord]string{\n")
	for k := range raw {
		sb.WriteString(fmt.Sprintf("  %s: \"%s\",\n", k, k))
	}
	sb.WriteString("}\n\n")

}

func generateTagConstants(raw map[string]etIptcRecord, sb *strings.Builder) {
	for _, v := range raw {
		sb.WriteString(fmt.Sprintf("//%s tag constants\nconst(\n", v.Name))
		for _, tag := range v.Tags {

			sb.WriteString(fmt.Sprintf("  %s_%s IptcTag = %v\n", v.Name, fixTagName(tag.Name), tag.Id))
		}
		sb.WriteString(")\n\n")
	}
}

func generateTagMap(raw map[string]etIptcRecord, sb *strings.Builder) {
	sb.WriteString("//IPTC Tag Descriptions\n")
	sb.WriteString("var IptcTagDescriptions = map[IptcRecordTag]IptcTagDesc{\n")
	for _, rec := range raw {
		for _, tag := range rec.Tags {
			min, max, iptcType := parseFmt(tag.Fmt)
			repeatable := parseFlags(tag.Flags)
			valueMap := generateIptcValueMap(iptcType, tag.Values)
			recTag := fmt.Sprintf("IptcRecordTag{%s,%v}", rec.Name, tag.Id)
			descFmt := `IptcTagDesc{
  Id: %v,
  Name: "%s",
  Type: %v,
  MinLength: %v,
  MaxLength: %v,
  Mandatory: %v,
  Repeatable: %v,
  Writable: %v,
  Values: %s,
}`
			tagDesc := fmt.Sprintf(descFmt, tag.Id, fixTagName(tag.Name), iptcType, min, max, tag.Mandatory, repeatable, tag.Writable, valueMap)
			sb.WriteString(fmt.Sprintf("%s: %s,\n", recTag, tagDesc))
		}
	}
	sb.WriteString("}\n\n")
}

func generateIptcValueMap(iptcType string, values map[string]string) string {

	buff := strings.Builder{}
	vals, err := parseValues(iptcType, values)
	if err != nil {
		if err.Error() != "no values" {
			fmt.Println(err)
		}
		return "nil"
	}
	switch iptcType {
	case "IptcString", "IptcDigits":
		m, ok := vals.(map[string]string)
		if !ok {
			return "nil"
		}
		buff.WriteString("map[string]string{\n")
		for k, v := range m {
			buff.WriteString(fmt.Sprintf("    \"%s\": \"%s\",\n", k, v))
		}
		buff.WriteString("  }")

	case "IptcUint8":
		m, ok := vals.(map[uint8]string)
		if !ok {
			return "nil"
		}
		buff.WriteString("map[uint8]string{\n")
		for k, v := range m {
			buff.WriteString(fmt.Sprintf("    %v: \"%s\",\n", k, v))
		}
		buff.WriteString("  }")
	case "IptcUint16":
		m, ok := vals.(map[uint16]string)
		if !ok {
			return "nil"
		}
		buff.WriteString("map[uint16]string{\n")
		for k, v := range m {
			buff.WriteString(fmt.Sprintf("    %v: \"%s\",\n", k, v))
		}
		buff.WriteString("  }")
	case "IptcUint32":
		m, ok := vals.(map[uint32]string)
		if !ok {
			return "nil"
		}
		buff.WriteString("map[uint32]string{\n")
		for k, v := range m {
			buff.WriteString(fmt.Sprintf("    %v: \"%s\",\n", k, v))
		}
		buff.WriteString("  }")
	default:
		return "nil"
	}
	return buff.String()
}

func parseFlags(flags string) bool {
	return strings.ToLower(flags) == "list"
}

func parseFmt(fmtStr string) (int, int, string) {
	if fmtStr == "" {
		return 0, 0, "IptcUndef"
	}
	parts := typeRegExp.FindStringSubmatch(fmtStr)
	var min, max int
	iptcType := ""
	if parts[2] != "" {
		min, _ = strconv.Atoi(parts[2])
	}
	if parts[3] != "" {
		max, _ = strconv.Atoi(parts[3])
	}
	if min > max {
		min = max
	}

	switch parts[1] {
	case "string":
		iptcType = "IptcString"
	case "digits":
		iptcType = "IptcDigits"
	case "int8u":
		iptcType = "IptcUint8"
	case "int16u":
		iptcType = "IptcUint16"
	case "int32u":
		iptcType = "IptcUint32"
	default:
		iptcType = "IptcUndef"
	}
	return min, max, iptcType
}

func parseValues(iptcType string, values map[string]string) (interface{}, error) {
	if len(values) == 0 {
		return nil, fmt.Errorf("no values")
	}
	parseInt := func(s string, t string) (uint64, error) {
		base := 10
		if strings.HasPrefix(s, "0x") {
			s = strings.TrimPrefix(s, "0x")
			base = 16
		}
		switch t {
		case "IptcUint8":
			return strconv.ParseUint(s, base, 8)
		case "IptcUint16":
			return strconv.ParseUint(s, base, 16)
		case "IptcUint32":
			return strconv.ParseUint(s, base, 32)
		default:
			return 0, fmt.Errorf("Unknown type: %s", t)
		}
	}

	switch iptcType {
	case "IptcString":
		return values, nil
	case "IptcDigits":
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		for k := range values {
			if b := strings.IndexFunc(k, isNotDigit) == -1; b {
				return values, nil
			}
		}
		return values, nil

	case "IptcUint8":
		ret := make(map[uint8]string)
		for k, v := range values {
			i, err := parseInt(k, iptcType)
			if err != nil {
				return ret, err
			}
			ret[uint8(i)] = v
		}
		return ret, nil
	case "IptcUint16":
		ret := make(map[uint16]string)
		for k, v := range values {
			i, err := parseInt(k, iptcType)
			if err != nil {
				return ret, err
			}
			ret[uint16(i)] = v
		}
		return ret, nil
	case "IptcUint32":
		ret := make(map[uint32]string)
		for k, v := range values {
			i, err := parseInt(k, iptcType)
			if err != nil {
				return ret, err
			}
			ret[uint32(i)] = v
		}
		return ret, nil
	default:
		if len(values) == 0 {
			return values, nil
		}
		return values, fmt.Errorf("Cannot parse values of type: %s", iptcType)
	}
}

func readExifToolIptcJSON() (map[string]etIptcRecord, error) {
	b, err := ioutil.ReadFile("assets/exiftool-iptctags.json")
	if err != nil {
		return nil, err
	}
	iptcMap := make(map[string]etIptcRecord)
	err = json.Unmarshal(b, &iptcMap)
	if err != nil {
		return nil, err
	}
	return iptcMap, nil
}
