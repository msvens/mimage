package generator

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var ifdTypes = map[string]uint16{
	"BYTE":      1,
	"ASCII":     2,
	"SHORT":     3,
	"LONG":      4,
	"RATIONAL":  5,
	"SSHORT":    6,
	"UNDEFINED": 7,
	"SLONG":     9,
	"SRATIONAL": 10,
}

type rawExifTag struct {
	Id          uint16 `json:"id"`
	Name        string `json:"name"`
	TypeName    string `json:"typeName"`
	Description string `json:"description"`
}

type rawExif struct {
	Name string       `json:"name"`
	Tags []rawExifTag `json:"tags"`
}

//Generate json and go source files with Exif tag data based on
//https://www.exiv2.org/tags.html
func GenerateExifTags() error {

	resp, err := http.Get("https://www.exiv2.org/tags.html")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	tbody, err := findTableBody(doc)
	if err != nil {
		fmt.Println("could not find table body")
		return err
	}
	data := parseTableBody(tbody)
	outBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./assets/exiftags.json", outBytes, 0644)
	if err != nil {
		return err
	}

	//create source files:
	err = generateExifTagSources(data)
	if err != nil {
		return err
	}
	//create exif values source files:
	exifVals, err := ParseExifValues()
	if err != nil {
		return err
	}
	return generateExifTagValueSources(exifVals, data)
}

func findTableBody(doc *html.Node) (*html.Node, error) {
	return findNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "tbody" {
			return true
		} else {
			return false
		}

	})
}

func parseTableBody(tbody *html.Node) map[string]rawExif {
	ret := make(map[string]rawExif)
	rows := findNodes(tbody, func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "tr" {
			return true
		} else {
			return false
		}
	})
	for _, row := range rows {
		td := parseExifRow(row)
		tag, ifd, err := parseExifTags(td)
		if err != nil {
			fmt.Println(err)
		} else {
			exif := ret[ifd]
			exif.Name = ifd
			exif.Tags = append(exif.Tags, tag)
			ret[ifd] = exif
		}
	}
	return ret
}

func parseExifRow(tr *html.Node) []string {
	var td []string
	for c := tr.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "td" {
			td = append(td, c.FirstChild.Data)
		}
	}
	return td
}

func parseExifTags(rd []string) (rawExifTag, string, error) {
	ret := rawExifTag{}
	ifd := ""
	if len(rd) != 6 {
		return ret, ifd, fmt.Errorf("Wrong number of table row elements")
	}
	//id
	cleaned := strings.Replace(rd[0], "0x", "", -1)
	if id, err := strconv.ParseUint(cleaned, 16, 16); err != nil {
		return ret, ifd, err
	} else {
		ret.Id = uint16(id)
	}
	//Ifd
	if rd[2] == "Image" {
		ifd = "IFD"
	} else if rd[2] == "Photo" {
		ifd = "Exif"
	} else {
		ifd = rd[2]
	}
	//Key
	lk := rd[3]
	for i := len(lk) - 1; i >= 0; i-- {
		if lk[i] == '.' {
			ret.Name = lk[i+1:]
			break
		}
	}

	//Type:
	upper := strings.ToUpper(strings.TrimSpace(rd[4]))
	if _, found := ifdTypes[upper]; found {
		ret.TypeName = upper
	} else {
		ret.TypeName = "UNDEFINED"
	}

	//Description:
	ret.Description = rd[5]
	return ret, ifd, nil
}

func generateExifTagSources(rawIfds map[string]rawExif) error {
	var buff strings.Builder
	var err error
	buff.WriteString(`package metadata
//Do not edit! This is an automatically generated file (see generator.GenerateExifTags()).
//This file was generated based on https://www.exiv2.org/tags.html


`)

	for k, v := range rawIfds {
		//first generate constants. A constant is prefixed by IfdName, e.g. IFD_Make or Exif_ExposureTime
		buff.WriteString(fmt.Sprintf("//%s Tag Ids\nconst(\n", k))
		for _, t := range v.Tags {
			buff.WriteString((fmt.Sprintf("  %s_%s uint16 = %#04x\n", k, t.Name, t.Id)))
		}
		buff.WriteString(")\n\n")

		//generate a id<->name mapping table
		buff.WriteString(fmt.Sprintf("var %sName = map[uint16]string{\n", k))
		for _, t := range v.Tags {
			buff.WriteString(fmt.Sprintf("  %s_%s: \"%s\",\n", k, t.Name, t.Name))
		}
		buff.WriteString("}\n\n")
	}
	err = ioutil.WriteFile("./metadata/genexiftags.go", []byte(buff.String()), 0644)
	return err
}
