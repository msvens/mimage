package generator

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ListxFile is the exiftool -listx dump the exif tables are generated from
const ListxFile = "assets/exiftool-listx.xml"

// exiftool table names we care about. Everything else in the dump (Composite,
// Extra, PanasonicRaw, ...) describes tags that are not plain exif entries
const (
	listxExifTable = "Exif::Main"
	listxGpsTable  = "GPS::Main"
)

const rawMain = "main"
const rawGps = "gps"

// listxTagInfo mirrors the <taginfo> document produced by exiftool -listx
type listxTagInfo struct {
	Tables []listxTable `xml:"table"`
}

type listxTable struct {
	Name string     `xml:"name,attr"`
	Tags []listxTag `xml:"tag"`
}

type listxTag struct {
	ID     string       `xml:"id,attr"`
	Name   string       `xml:"name,attr"`
	Type   string       `xml:"type,attr"`
	Count  string       `xml:"count,attr"`
	G1     string       `xml:"g1,attr"`
	Values *listxValues `xml:"values"`
}

type listxValues struct {
	Keys []listxKey `xml:"key"`
}

type listxKey struct {
	ID   string     `xml:"id,attr"`
	Vals []listxVal `xml:"val"`
}

type listxVal struct {
	Lang string `xml:"lang,attr"`
	Text string `xml:",chardata"`
}

// exifTag is the normalised tag the generator emits from
type exifTag struct {
	Id     uint16
	Name   string
	Type   string
	Ifd    string
	Count  int
	Values map[string]string
}

// subDirTags are tags exiftool models as SubDirectory entries rather than plain
// tags, so they never appear in a -listx dump. They are all pointers to another
// ifd or data block. mimage still needs them: IFD_ExifOffset in particular is
// used by ExifEditor.DropMakerNote to find the ExifIFD without creating one
var subDirTags = []exifTag{
	{Id: 0x0190, Name: "GlobalParametersIFD", Type: "ExifUndef", Ifd: "ExifIFD", Count: -1},
	{Id: 0x4748, Name: "StitchInfo", Type: "ExifUndef", Ifd: "ExifIFD", Count: -1},
	{Id: 0x8290, Name: "KodakIFD", Type: "ExifUndef", Ifd: "ExifIFD", Count: -1},
	{Id: 0x8568, Name: "AFCP_IPTC", Type: "ExifUndef", Ifd: "ExifIFD", Count: -1},
	{Id: 0x8606, Name: "LeafData", Type: "ExifUndef", Ifd: "ExifIFD", Count: -1},
	{Id: 0x8649, Name: "PhotoshopSettings", Type: "ExifUndef", Ifd: "RootIFD", Count: -1},
	{Id: 0x8769, Name: "ExifOffset", Type: "ExifUint32", Ifd: "RootIFD", Count: 1},
	{Id: 0x8773, Name: "ICC_Profile", Type: "ExifUndef", Ifd: "RootIFD", Count: -1},
	{Id: 0x8825, Name: "GPSInfo", Type: "ExifUint32", Ifd: "RootIFD", Count: 1},
	{Id: 0x888a, Name: "LeafSubIFD", Type: "ExifUint32", Ifd: "ExifIFD", Count: 1},
	{Id: 0xa005, Name: "InteropOffset", Type: "ExifUint32", Ifd: "ExifIFD", Count: 1},
	{Id: 0xc51b, Name: "HasselbladExif", Type: "ExifUndef", Ifd: "ExifIFD", Count: -1},
	{Id: 0xc6f5, Name: "ProfileIFD", Type: "ExifUint32", Ifd: "RootIFD", Count: 1},
	{Id: 0xc7d5, Name: "NikonNEFInfo", Type: "ExifUndef", Ifd: "ExifIFD", Count: -1},
	{Id: 0xcd3b, Name: "RGBTables", Type: "ExifUndef", Ifd: "RootIFD", Count: -1},
	{Id: 0xfe00, Name: "KDC_IFD", Type: "ExifUndef", Ifd: "ExifIFD", Count: -1},
}

// ifdOverrides pins tags whose -listx group does not match where mimage reads
// them. ThumbnailOffset/ThumbnailLength are reported as ExifIFD but belong with
// the root ifd chain here
var ifdOverrides = map[uint16]string{
	0x0201: "RootIFD",
	0x0202: "RootIFD",
}

// nameOverrides keeps exported constant names stable where the name exiftool
// lists first differs from the one mimage has always used. Renaming any of
// these would silently break callers, so every entry here is deliberate.
//
// The first group is tags where -listx leads with a different variant name.
// The second is the older duplicate ids that exiftool also lists: they share a
// name with a tag mimage already exports, and being numerically lower they
// would otherwise claim the unsuffixed name and push the established tag to a
// suffixed one
var nameOverrides = map[uint16]string{
	0x014a: "SubIFDs",             //-listx lists only the A100DataOffset variant
	0x0111: "StripOffsets",        //-listx leads with PreviewImageStart
	0x0117: "StripByteCounts",     //-listx leads with PreviewImageLength
	0x927c: "MakerNote",           //-listx leads with MakerNoteApple
	0xc634: "DNGPrivateData",      //-listx leads with DNGAdobeData
	0xcd30: "SemanticInstanceIFD", //exiftool renamed this to SemanticInstanceID

	0x920c: "SpatialFrequencyResponse_0x920c", //duplicate of 0xa20c
	0x920e: "FocalPlaneXResolution_0x920e",    //duplicate of 0xa20e
	0x920f: "FocalPlaneYResolution_0x920f",    //duplicate of 0xa20f
	0x9215: "ExposureIndex_0x9215",            //duplicate of 0xa215
	0x7310: "BlackLevel_0x7310",               //duplicate of 0xc61a
}

// typeOverrides restores types that exiftool's EXIF.pm tables carry but that
// -listx reports as "?". Without these the affected tags would silently drop to
// ExifUndef and be scanned as raw bytes instead of their real type
var typeOverrides = map[uint16]string{
	0x014a: "ExifUint32", //SubIFDs
	0x0153: "ExifUint16", //SampleFormat
	0x8649: "ExifUint8",  //PhotoshopSettings
	0x9009: "ExifUint8",  //GooglePlusUploadCode
	0x9101: "ExifUint8",  //ComponentsConfiguration
	0xc616: "ExifUint8",  //CFAPlaneColor
	0xc617: "ExifUint16", //CFALayout
	0xc634: "ExifUint8",  //DNGPrivateData
	0xc6d2: "ExifString", //PanasonicTitle
	0xc6d3: "ExifString", //PanasonicTitle2
	0xcd2e: "ExifString", //SemanticName
	0xcd30: "ExifString", //SemanticInstanceIFD
	0xcd38: "ExifUint32", //MaskSubArea
}

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
  Ifd  ExifIndex ` + "`json:\"ifd\"`" + `
  Count int ` + "`json:\"count\"`" + `
  Values interface{} ` + "`json:\"values\"`" + `
}
`

// exifTypeMapping translates an exiftool -listx type to a mimage ExifTagType.
// "?", "binary" and "undef" all mean the value has no fixed scalar type
var exifTypeMapping = map[string]string{
	"string":      "ExifString",
	"int8u":       "ExifUint8",
	"int16u":      "ExifUint16",
	"int32u":      "ExifUint32",
	"int16s":      "ExifInt16",
	"int32s":      "ExifInt32",
	"rational64u": "ExifUrational",
	"rational64s": "ExifRational",
	"float":       "ExifFloat",
	"double":      "ExifDouble",
}

func exifType(listxType string) string {
	if t, found := exifTypeMapping[listxType]; found {
		return t
	}
	return "ExifUndef"
}

// exifIfd maps a -listx table and group to a mimage ExifIndex. Anything that is
// not explicitly the exif or interop ifd belongs with the root ifd, which is
// how the previous exiftool sourced generator behaved as well
func exifIfd(table, g1 string) string {
	if table == listxGpsTable {
		return "GpsIFD"
	}
	switch g1 {
	case "ExifIFD":
		return "ExifIFD"
	case "InteropIFD":
		return "InteropIFD"
	default:
		return "RootIFD"
	}
}

func alignName(rawName string) string {
	return strings.ReplaceAll(rawName, "-", "")
}

// englishValues picks the en translation out of a -listx <values> block
func englishValues(v *listxValues) map[string]string {
	if v == nil || len(v.Keys) == 0 {
		return nil
	}
	ret := map[string]string{}
	for _, k := range v.Keys {
		for _, val := range k.Vals {
			if val.Lang == "en" {
				ret[k.ID] = strings.TrimSpace(val.Text)
				break
			}
		}
	}
	if len(ret) == 0 {
		return nil
	}
	return ret
}

// ReadListxExifTags parses the committed exiftool -listx dump into the tag set
// the exif sources are generated from
func ReadListxExifTags() (map[string][]exifTag, error) {
	b, err := os.ReadFile(ListxFile)
	if err != nil {
		return nil, err
	}
	var doc listxTagInfo
	if err = xml.Unmarshal(b, &doc); err != nil {
		return nil, err
	}

	ret := map[string][]exifTag{rawMain: {}, rawGps: {}}
	for _, table := range doc.Tables {
		if table.Name != listxExifTable && table.Name != listxGpsTable {
			continue
		}
		group := rawMain
		if table.Name == listxGpsTable {
			group = rawGps
		}
		for _, t := range table.Tags {
			id, e := strconv.ParseUint(t.ID, 10, 16)
			if e != nil {
				//composite entries use non numeric ids such as Exif-JpgFromRaw
				continue
			}
			tag := exifTag{
				Id:     uint16(id),
				Name:   alignName(t.Name),
				Type:   exifType(t.Type),
				Ifd:    exifIfd(table.Name, t.G1),
				Count:  -1,
				Values: englishValues(t.Values),
			}
			if t.Count != "" {
				if c, e := strconv.Atoi(t.Count); e == nil {
					tag.Count = c
				}
			}
			if override, found := ifdOverrides[tag.Id]; found {
				tag.Ifd = override
			}
			ret[group] = append(ret[group], tag)
		}
	}

	ret[rawMain] = append(ret[rawMain], subDirTags...)
	collapseVariants(ret)
	dedupeNames(ret)
	return ret, nil
}

// collapseVariants keeps a single record per ifd and tag id. exiftool lists the
// same id several times when its meaning depends on context, for instance
// 0x0201 is ThumbnailOffset, PreviewImageStart, JpgFromRawStart and
// OtherImageStart. mimage keys tags on ifd and id so it can only hold one, and
// the first listed variant is the one it has always used
func collapseVariants(tags map[string][]exifTag) {
	for _, group := range []string{rawMain, rawGps} {
		seen := map[string]bool{}
		kept := make([]exifTag, 0, len(tags[group]))
		for _, t := range tags[group] {
			key := fmt.Sprintf("%s%d", t.Ifd, t.Id)
			if seen[key] {
				continue
			}
			seen[key] = true
			if name, found := nameOverrides[t.Id]; found {
				t.Name = name
			}
			if typ, found := typeOverrides[t.Id]; found {
				t.Type = typ
			}
			kept = append(kept, t)
		}
		sort.Slice(kept, func(i, j int) bool { return kept[i].Id < kept[j].Id })
		tags[group] = kept
	}
}

// dedupeNames suffixes a duplicate name within an ifd with its tag id, matching
// what the previous generator did for the few exif tags that share a name
func dedupeNames(tags map[string][]exifTag) {
	for _, group := range []string{rawMain, rawGps} {
		seen := map[string]bool{}
		for i := range tags[group] {
			t := &tags[group][i]
			key := t.Ifd + t.Name
			if seen[key] {
				t.Name = fmt.Sprintf("%s_%#04x", t.Name, t.Id)
			} else {
				seen[key] = true
			}
		}
	}
}

// GenerateExifTagsFromListx generates metadata/genexif.go from the exiftool
// -listx dump in assets
func GenerateExifTagsFromListx() error {
	raw, err := ReadListxExifTags()
	if err != nil {
		return err
	}
	sb := strings.Builder{}
	sb.WriteString(`package metadata
//Do not edit! This is an automatically generated file (see generator.GenerateExifTagsFromListx()).
//Generated from assets/exiftool-listx.xml, produced by: exiftool -listx -EXIF:all
`)
	sb.WriteString(exifTypesSrc)
	sb.WriteString(exifIndexSrc)
	sb.WriteString(exifTypeConstSrc)
	sb.WriteString(exifTagDescSrc)
	generateExifConstants(raw, &sb)
	generateExifTagDescriptions(raw, &sb)

	return os.WriteFile("./metadata/genexif.go", []byte(sb.String()), 0644)
}

func generateExifTagDescriptions(raw map[string][]exifTag, sb *strings.Builder) {
	sb.WriteString("//Exif Tag Descriptions\n")
	sb.WriteString("var ExifTagDescriptions = map[ExifIndexTag]ExifTagDesc{\n")

	for _, group := range []string{rawMain, rawGps} {
		for _, t := range raw[group] {
			valueMap := generateExifValueMap(t)
			descFmt := `ExifTagDesc{
  Id: %#04x,
  Name: "%s",
  Type: %v,
  Ifd: %s,
  Count: %v,
  Values: %s,
}`
			tagDesc := fmt.Sprintf(descFmt, t.Id, t.Name, t.Type, t.Ifd, t.Count, valueMap)
			fmt.Fprintf(sb, "ExifIndexTag{%s,%#04x}: %s,\n", t.Ifd, t.Id, tagDesc)
		}
	}
	sb.WriteString("}\n")
}

func generateExifConstants(raw map[string][]exifTag, sb *strings.Builder) {
	emit := func(header, prefix, ifd string, tags []exifTag) {
		fmt.Fprintf(sb, "//%s\nconst(\n", header)
		for _, t := range tags {
			if ifd != "" && t.Ifd != ifd {
				continue
			}
			fmt.Fprintf(sb, "  %s%s ExifTag = %#04x\n", prefix, t.Name, t.Id)
		}
		sb.WriteString(")\n")
	}
	emit("IFD Tag Ids (includes all IFD, IFD1, etc tags)", "IFD_", "RootIFD", raw[rawMain])
	emit("ExifIFD Tag Ids", "ExifIFD_", "ExifIFD", raw[rawMain])
	emit("InteropIFD Tag Ids", "InteropIFD_", "InteropIFD", raw[rawMain])
	emit("GpsIFD Tag Ids", "GpsIFD_", "", raw[rawGps])
}

// generateExifValueMap renders the enumerated values for a tag. Undefined tags,
// rationals and multi value tags do not get a lookup map
//
//gocyclo:ignore
func generateExifValueMap(tag exifTag) string {
	if len(tag.Values) == 0 || tag.Type == "ExifUndef" {
		return "nil"
	} else if tag.Count > 1 { //we dont generate a value map for complex values
		return "nil"
	} else if tag.Type == "ExifUrational" || tag.Type == "ExifRational" {
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

	//emit keys in a stable order so regeneration is reproducible
	keys := make([]string, 0, len(tag.Values))
	for k := range tag.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	numeric := func(goType string, bitSize int, float bool, unsigned bool) string {
		buff := strings.Builder{}
		fmt.Fprintf(&buff, "map[%s]string{\n", goType)
		for _, k := range keys {
			if err := checkNumber(k, bitSize, float, unsigned); err != nil {
				fmt.Println("could not parse value key:", k)
				continue
			}
			fmt.Fprintf(&buff, "    %s: \"%s\",\n", k, escape(tag.Values[k]))
		}
		buff.WriteString("  }")
		return buff.String()
	}

	switch tag.Type {
	case "ExifString":
		buff := strings.Builder{}
		buff.WriteString("map[string]string{\n")
		for _, k := range keys {
			fmt.Fprintf(&buff, "    \"%s\": \"%s\",\n", k, escape(tag.Values[k]))
		}
		buff.WriteString("  }")
		return buff.String()
	case "ExifFloat":
		return numeric("float", 32, true, false)
	case "ExifDouble":
		return numeric("double", 64, true, false)
	case "ExifUint8":
		return numeric("uint8", 8, false, true)
	case "ExifUint16":
		return numeric("uint16", 16, false, true)
	case "ExifUint32":
		return numeric("uint32", 32, false, true)
	case "ExifInt16":
		return numeric("int16", 16, false, false)
	case "ExifInt32":
		return numeric("int32", 32, false, false)
	default:
		fmt.Println("Type not found:", tag.Type)
		return "nil"
	}
}

// escape makes a value description safe to emit inside a go string literal
func escape(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}
