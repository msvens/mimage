package generator

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type iptcRecord struct {
	Id      uint8        `json:"id"`
	Name    string       `json:"name"`
	Records []rawIptcTag `json:"records"`
}

type rawIptcTag struct {
	Id        uint8  `json:"id"`
	Name      string `json:"name"`
	TypeName  string `json:"typeName"`
	MinLength int    `json:"minLength"`
	MaxLength int    `json:"maxLength"`
}

var recordSections = map[string]string{
	"EnvelopeRecord":    "IPTCEnvelope",
	"ApplicationRecord": "IPTCApplication",
	"NewsPhoto":         "IPTCNewsPhoto",
	"PreObjectData":     "IPTCPreObjectData",
	"ObjectData":        "IPTCObjectData",
	"PostObjectData":    "IPTCPostObjectData",
	"FotoStation":       "IPTCFotoStation",
}

var recordNameMapping = map[string]uint8{
	"IPTCEnvelope":       1,
	"IPTCApplication":    2,
	"IPTCNewsPhoto":      3,
	"IPTCPreObjectData":  7,
	"IPTCObjectData":     8,
	"IPTCPostObjectData": 9,
	"IPTCFotoStation":    240,
}

//Digitis and strins can have length specifiers (min,max)
var typeRegExp, _ = regexp.Compile(`^\s*(\w+).*?\[?(\d*),?(\d*)\]?\s*$`)

//Generate json file and go source file with IPTC tag data from
//https://exiftool.org/TagNames/IPTC.html
func GenerateIptcTags() error {
	resp, err := http.Get("https://exiftool.org/TagNames/IPTC.html")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	heads := findRecordHeadings(doc)
	records := make(map[string]iptcRecord)

	for k, v := range heads {
		if tags, err := parseRecord(v); err != nil {
			return err
		} else {
			r := iptcRecord{recordNameMapping[k], k, tags}
			records[k] = r

		}
	}

	outBytes, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./metadata/iptctags.json", outBytes, 0644)
	if err != nil {
		return err
	}

	err = generateIptcTagSources(records)
	return err
}

func generateIptcTagSources(rawIptcs map[string]iptcRecord) error {
	var buff strings.Builder
	var err error
	buff.WriteString(`package metadata
//Do not edit! This is an automatically generated file (see generator.GenerateIptcTags()).
//This file was generated based on https://exiftool.org/TagNames/IPTC.html


`)

	buff.WriteString(`//IPTC Records
const(
  IPTCEnvelop uint8 = 1
  IPTCApplication uint8 = 2
  IPTCNewsPhoto uint8 = 3
  IPTCPreObjectData uint8 = 7
  IPTCObjectData uint8 = 8
  IPTCPostObjectData uint8 = 9
  IPTCFotoStation uint8 = 240
)

var IptcRecordName = map[uint8]string{
  IPTCEnvelop: "IPTCEnvelop",
  IPTCApplication: "IPTCApplication",
  IPTCNewsPhoto: "IPTCNewsPhoto",
  IPTCPreObjectData: "IPTCPreObjectData",
  IPTCObjectData: "IPTCObjectData",
  IPTCPostObjectData: "IPTCPostObjectData",
  IPTCFotoStation: "IPTCFotoStation",
}


`)
	for k, v := range rawIptcs {
		//first generate constants. A constant is prefixed by IfdName, e.g. IFD_Make or Exif_ExposureTime
		buff.WriteString(fmt.Sprintf("//%s Tag Ids\nconst(\n", k))
		for _, t := range v.Records {
			buff.WriteString((fmt.Sprintf("  Iptc%v_%s uint8 = %v\n", v.Id, t.Name, t.Id)))
		}
		buff.WriteString(")\n\n")
		//create id <-> name mapping
		buff.WriteString(fmt.Sprintf("var Iptc%vName = map[uint8]string{\n", v.Id))
		for _, t := range v.Records {
			buff.WriteString(fmt.Sprintf("  Iptc%v_%s: \"%s\",\n", v.Id, t.Name, t.Name))
		}
		buff.WriteString("}\n\n")
	}
	err = ioutil.WriteFile("./metadata/geniptctags.go", []byte(buff.String()), 0644)
	return err
	return err
}

func findRecordHeadings(node *html.Node) map[string]*html.Node {

	//find all <h2><a name='recordName'>IPTC EnvelopeRecord Tags</a></h2>
	ret := make(map[string]*html.Node)

	match := func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "name" {
					if v, found := recordSections[a.Val]; found {
						ret[v] = n.Parent
						return true
					}

				}
			}
		}
		return false
	}
	_ = findNodes(node, match) //we dont care about the result

	return ret
}

func parseRecord(node *html.Node) ([]rawIptcTag, error) {
	//Outline:
	//<blockquote>
	//  <table class=frame><tr><td>
	//    <table class=inner cellspacing=1>
	//      <tr></tr>
	//    </table>
	//  </table>
	//</blockquote>
	var ret []rawIptcTag

	matchTable := func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "table" {
			return true
		}
		return false
	}

	block, err := findSibling(node, func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "blockquote" {
			return true
		} else {
			return false
		}
	})
	if err != nil {
		return ret, err
	}
	outer, err := findNode(block, matchTable)
	if err != nil {
		return ret, err
	}
	inner, err := findNode(outer.FirstChild, matchTable)
	if err != nil {
		return ret, err
	}
	rows := findNodes(inner, func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "tr" {
			return true
		} else {
			return false
		}
	})

	//skip first row (header row)
	for i := 1; i < len(rows); i++ {
		if tag, err := parseIptcTableRow(rows[i]); err == nil {
			ret = append(ret, tag)
		} else {
			fmt.Println(err)
		}
	}
	return ret, nil
}

func parseIptcTableRow(tr *html.Node) (rawIptcTag, error) {
	var ret = rawIptcTag{}
	var td []string
	for c := tr.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "td" {
			td = append(td, c.FirstChild.Data)
		}
	}
	if len(td) != 4 {
		return ret, fmt.Errorf("Wrong number of cells in table row: %v", len(td))
	}
	if id, err := strconv.ParseUint(td[0], 10, 8); err != nil {
		return ret, err
	} else {
		ret.Id = uint8(id)
	}
	ret.Name = strings.Replace(td[1], "-", "", -1)
	if err := parseType(td[2], &ret); err != nil {
		return ret, fmt.Errorf("Could not parse type: ", td[2])
	}
	return ret, nil
}

func parseType(str string, td *rawIptcTag) error {
	lower := strings.ToLower(str)
	parts := typeRegExp.FindStringSubmatch(lower)
	if parts == nil {
		return fmt.Errorf("Could not parse type")
	}
	//set length indicators
	if parts[2] != "" {
		td.MinLength, _ = strconv.Atoi(parts[2])
	}
	if parts[3] != "" {
		td.MaxLength, _ = strconv.Atoi(parts[3])
	}
	//now check the type
	t := parts[1]
	if t == "string" || t == "digits" || t == "int32u" || t == "int16u" ||
		t == "int8u" || t == "undef" {
		td.TypeName = t
	} else if t == "no" {
		td.TypeName = "nonWriteable"
	} else {
		fmt.Println("could not determine type: ", t)
		td.TypeName = "undef"
	}
	return nil

	/*upper := strings.TrimSpace(strings.ToUpper(str))
	if strings.HasPrefix(upper, "STRING") {
		return "STRING"
	} else if strings.HasPrefix(upper, "DIGITS") {
		return "DIGITS"
	} else if strings.HasPrefix(upper, "INT32U") {
		return "UINT32"
	} else if strings.HasPrefix(upper, "INT16U") {
		return "UINT16"
	}  else if strings.HasPrefix(upper, "INT8U") {
		return "UINT8"
	} else if strings.HasPrefix(upper, "UNDEF") {
		return "UNDEFINED"
	} else if strings.HasPrefix(upper, "NO") {
		return "NON-WRITABLE"
	} else {
		fmt.Println("unrecognised type: ", upper)
		return "UNDEFINED"
	}*/

}
