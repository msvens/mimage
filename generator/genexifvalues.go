package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type valueString struct {
	Value string
	Name  string
}
type RawExifValues struct {
	Id     uint16
	Group  string
	Values []valueString
}

//val = string
var valueStrRegExp, _ = regexp.Compile(`^\s*(\w+)\s*=\s*(\w+.*)\s*$`)
var tagRegExp, _ = regexp.Compile(`^\s*(\w+)\s*,\s*(\w+)\s*$`)

//var valueStrRegExp, _ = regexp.Compile(`^\s*(\w+).*?\[?(\d*),?(\d*)\]?\s*$`)

func generateExifTagValueSources(values []RawExifValues, exifs map[string]rawExif) error {
	findVal := func(r rawExifTag, group string) (RawExifValues, bool) {
		for _, v := range values {
			if v.Group == group && v.Id == r.Id {
				return v, true
			}
		}
		return RawExifValues{}, false
	}

	b := strings.Builder{}
	b.WriteString(`package metadata
//Do not edit! This is an automatically generated file (see generator.GenerateExifTags()).
//This file was generated based on https://www.exiv2.org/tags.html


`)

	for k, exif := range exifs {
		b.WriteString(fmt.Sprintf("var %sValues = map[uint16]interface{}{\n", k))
		for _, re := range exif.Tags {
			if ev, found := findVal(re, k); found {
				s, e := generateValueMap(ev, re)
				if e != nil {
					return e
				} else {
					b.WriteString(fmt.Sprintf("  %s_%s: %s,\n", k, re.Name, s))
				}
			}
		}
		b.WriteString("}\n\n")
	}
	return ioutil.WriteFile("./metadata/genexifvalues.go", []byte(b.String()), 0644)
	//fmt.Println(b.String())
	//return nil
}

func parseInt(v string, unsigned bool) (int64, error) {
	if strings.HasPrefix(v, "0x") {
		if unsigned {
			r, err := strconv.ParseUint(v[2:], 16, 32)
			return int64(r), err
		} else {
			return strconv.ParseInt(v[2:], 16, 32)
		}
	} else {
		if unsigned {
			r, err := strconv.ParseUint(v, 10, 32)
			return int64(r), err
		} else {
			return strconv.ParseInt(v[2:], 10, 32)
		}
	}
}
func generateValueMap(vals RawExifValues, exif rawExifTag) (string, error) {

	b := strings.Builder{}
	if exif.TypeName == "BYTE" {
		b.WriteString("map[[]byte]string{\n")
		for _, v := range vals.Values {
			b.WriteString(fmt.Sprintf("    []byte(\"%s\"): \"%s\",\n", v.Value, v.Name))
		}
		b.WriteString("  }")

	} else if exif.TypeName == "ASCII" {
		b.WriteString("map[string]string{\n")
		for _, v := range vals.Values {
			b.WriteString(fmt.Sprintf("    \"%s\": \"%s\",\n", v.Value, v.Name))
		}
		b.WriteString("  }")

	} else if exif.TypeName == "SHORT" {
		b.WriteString("map[uint16]string{\n")
		for _, v := range vals.Values {
			if val, e := parseInt(v.Value, true); e == nil {
				b.WriteString(fmt.Sprintf("    %#0x: \"%s\",\n", val, v.Name))
			} else {
				return "", fmt.Errorf("Could not parse value: %v", e)
			}

		}
		b.WriteString("  }")
	} else if exif.TypeName == "SSHORT" {
		b.WriteString("map[int16]string{\n")
		for _, v := range vals.Values {
			if val, e := parseInt(v.Value, false); e == nil {
				b.WriteString(fmt.Sprintf("    %#04x: \"%s\",\n", val, v.Name))
			} else {
				return "", fmt.Errorf("Could not parse value: %v", e)
			}

		}
		b.WriteString(" }")
	} else if exif.TypeName == "LONG" {
		b.WriteString("map[uint32]string{\n")
		for _, v := range vals.Values {
			if val, e := parseInt(v.Value, true); e == nil {
				b.WriteString(fmt.Sprintf("    %v: \"%s\",\n", val, v.Name))
			} else {
				return "", fmt.Errorf("Could not parse value: %v", e)
			}
		}
		b.WriteString("  }")

	} else if exif.TypeName == "SLONG" {
		b.WriteString("map[int32]string{\n")
		for _, v := range vals.Values {
			if val, e := parseInt(v.Value, true); e == nil {
				b.WriteString(fmt.Sprintf("    %v: \"%s\",\n", val, v.Name))
			} else {
				return "", fmt.Errorf("Could not parse value: %v", e)
			}

		}
		b.WriteString("  }")
	} else {
		return "", fmt.Errorf("Could not handle type: %s", exif.TypeName)
	}
	return b.String(), nil
}

func ParseExifValues() ([]RawExifValues, error) {
	file, err := os.Open("assets/exiftagvalues.txt")
	ret := []RawExifValues{}

	if err != nil {
		log.Fatalf("failed to open: %v", err)

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		l := scanner.Text()
		if strings.HasPrefix(l, "0x") {
			v, e := parseExifValues(scanner)
			if e != nil {
				fmt.Println("Could not parse ", e.Error())
			}
			ret = append(ret, v)
		}
	}
	return ret, nil
}

func convertGroup(group string) string {
	if strings.EqualFold(group, "IFD0") || strings.EqualFold(group, "SubIfd") {
		return "IFD"
	}
	if strings.EqualFold(group, "ExifIFD") {
		return "Exif"
	}
	if strings.EqualFold(group, "InteropIFD") {
		return "Iop"
	}
	return "Unknown"
}

func parseExifValues(scanner *bufio.Scanner) (RawExifValues, error) {
	ret := RawExifValues{}
	parts := tagRegExp.FindStringSubmatch(scanner.Text())
	if parts == nil {
		return ret, fmt.Errorf("Could not parse exif value header %v", scanner.Text())
	}
	if id, err := strconv.ParseUint(parts[1][2:], 16, 16); err != nil {
		return ret, fmt.Errorf("Could not parse exif id")
	} else {
		ret.Id = uint16(id)
		ret.Group = convertGroup(parts[2])
	}
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(l, "END") {
			break
		}
		parts = valueStrRegExp.FindStringSubmatch(l)
		if parts == nil {
			fmt.Println("Could not parse line", l)
		} else {
			ret.Values = append(ret.Values, valueString{Value: parts[1], Name: parts[2]})
		}
	}
	if len(ret.Values) == 0 {
		return ret, fmt.Errorf("error values is empty")
	}
	return ret, nil
}
