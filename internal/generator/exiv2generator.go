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

var exivTypeMap = map[string]string{
	"BYTE":      "ExifUint8",
	"ASCII":     "ExifString",
	"SHORT":     "ExifUint16",
	"LONG":      "ExifUint32",
	"SSHORT":    "ExifInt16",
	"SLONG":     "ExifInt32",
	"SRATIONAL": "ExifRational",
	"RATIONAL":  "ExifUrational",
	"UNDEFINED": "ExifUndef",
}

//Generate ExifTags Json from https://www.exiv2.org/tags.html. This file is used
//to build our master-exiftags.json
func GenerateExiv2ExifJson() error {

	//read source:
	resp, err := http.Get("https://www.exiv2.org/tags.html")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//find table body:
	tbody, err := findNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "tbody" {
			return true
		} else {
			return false
		}
	})
	data := parseTableBody(tbody)
	out := map[string][]rawExifTag{}
	for k, v := range data {
		out[k] = v.Tags
	}
	outBytes, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("./assets/exiv2-exiftags.json", outBytes, 0644)

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
	switch rd[2] {
	case "Image":
		ifd = "RootIFD"
	case "Photo":
		ifd = "ExifIFD"
	case "GPSInfo":
		ifd = "GpsIFD"
	case "Iop":
		ifd = "InteropIFD"
	default:
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
	if v, found := exivTypeMap[upper]; found {
		ret.TypeName = v
	} else {
		ret.TypeName = "ExifUndef"
	}

	//Description:
	ret.Description = rd[5]
	return ret, ifd, nil
}

type Matcher = func(node *html.Node) bool

func findNodes(node *html.Node, matcher Matcher) []*html.Node {
	var ret []*html.Node

	var f func(*html.Node)
	f = func(n *html.Node) {
		if matcher(n) {
			ret = append(ret, n)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(node)
	return ret
}

func findNode(node *html.Node, matcher Matcher) (*html.Node, error) {
	var ret *html.Node

	var f func(*html.Node)
	f = func(n *html.Node) {
		if ret != nil {
			return
		}
		if matcher(n) {
			ret = n
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(node)
	if ret == nil {
		return ret, fmt.Errorf("Node not found")
	} else {
		return ret, nil
	}
}
