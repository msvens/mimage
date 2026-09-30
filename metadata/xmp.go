package metadata

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"trimmer.io/go-xmp/models/dc"
	//registers the exif namespace, read by path in GetLocation
	_ "trimmer.io/go-xmp/models/exif"
	"trimmer.io/go-xmp/models/ps"
	//registers the tiff namespace, read by path in GetCamera
	_ "trimmer.io/go-xmp/models/tiff"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
	xmpmm "trimmer.io/go-xmp/models/xmp_mm"
	"trimmer.io/go-xmp/xmp"
)

// XmpData holds the underlying xmp document
type XmpData struct {
	rawXmp *xmp.Document
}

// ErrNoXmp when a jpeg image does not contain any xmp data
var ErrNoXmp = errors.New("no XMP data")

// NewXmpData creates an XmpData struct from a jpeg segment list
func NewXmpData(segments *jpegstructure.SegmentList) (XmpData, error) {
	_, s, err := segments.FindXmp()
	if err != nil {
		return XmpData{}, ErrNoXmp
	}
	str, err := s.FormattedXmp()
	if err != nil {
		//We should log errors
		return XmpData{}, ErrNoXmp
	}
	return NewXmpDataFromBytes([]byte(str))
}

// NewXmpDataFromBytes creates an XmpData struct from marshalled xmp.Document
func NewXmpDataFromBytes(data []byte) (XmpData, error) {
	model := &xmp.Document{}
	err := xmp.Unmarshal(data, model)
	if err != nil {
		return XmpData{}, ErrNoXmp
	}
	return XmpData{model}, nil
}

// Base retrieves the base model
func (xd XmpData) Base() *xmpbase.XmpBase {
	if !xd.IsEmpty() {
		return xmpbase.FindModel(xd.rawXmp)
	}
	return nil
}

// DublinCore retrieves the DublinCore model
func (xd XmpData) DublinCore() *dc.DublinCore {
	if !xd.IsEmpty() {
		return dc.FindModel(xd.rawXmp)
	}
	return nil
}

// GetKeywords returns the keywords from DublinCore
func (xd XmpData) GetKeywords() []string {
	if dcore := xd.DublinCore(); dcore != nil {
		return dcore.Subject
	}
	return []string{}
}

// GetRating returns rating from Base
func (xd XmpData) GetRating() uint16 {
	if base := xd.Base(); base != nil {
		return uint16(base.Rating)
	}
	return 0
}

// GetTitle returns the DublinCore title if it exists
func (xd XmpData) GetTitle() string {
	if dcore := xd.DublinCore(); dcore != nil {
		return dcore.Title.Default()
	}
	return ""
}

// GetDescription returns the DublinCore description, the caption, if it exists
func (xd XmpData) GetDescription() string {
	if dcore := xd.DublinCore(); dcore != nil {
		return dcore.Description.Default()
	}
	return ""
}

// GetCamera returns the tiff make and model, or empty strings
func (xd XmpData) GetCamera() (string, string) {
	return xd.pathValue("tiff:Make"), xd.pathValue("tiff:Model")
}

// GetDate returns when the content was created, trying exif:DateTimeOriginal,
// photoshop:DateCreated and xmp:CreateDate in that order. A date without a
// timezone is returned in UTC. The zero time if there is none
func (xd XmpData) GetDate() time.Time {
	for _, p := range []string{"exif:DateTimeOriginal", "photoshop:DateCreated", "xmp:CreateDate"} {
		if t, ok := parseXmpDate(xd.pathValue(p)); ok {
			return t
		}
	}
	return time.Time{}
}

// GetLocation returns the exif gps position in decimal degrees, with ok false
// when there is none
func (xd XmpData) GetLocation() (lat float64, long float64, ok bool) {
	lat, okLat := parseXmpCoordinate(xd.pathValue("exif:GPSLatitude"))
	long, okLong := parseXmpCoordinate(xd.pathValue("exif:GPSLongitude"))
	if !okLat || !okLong {
		return 0, 0, false
	}
	return lat, long, true
}

// pathValue reads a simple property by its prefixed name, "" if it is absent
func (xd XmpData) pathValue(path string) string {
	if xd.IsEmpty() {
		return ""
	}
	v, err := xd.rawXmp.GetPath(xmp.Path(path))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(v)
}

// the date forms the xmp spec allows, most precise first
var xmpDateLayouts = []string{
	time.RFC3339Nano,
	"2006-01-02T15:04:05",
	"2006-01-02T15:04Z07:00",
	"2006-01-02T15:04",
	"2006-01-02",
}

func parseXmpDate(s string) (time.Time, bool) {
	if s == "" {
		return time.Time{}, false
	}
	for _, layout := range xmpDateLayouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

// parseXmpCoordinate reads an xmp GPSCoordinate, "DDD,MM,SSk" or "DDD,MM.mmk"
// where k is one of N, S, E or W
func parseXmpCoordinate(s string) (float64, bool) {
	if len(s) < 2 {
		return 0, false
	}
	sign := 1.0
	switch s[len(s)-1] {
	case 'N', 'E':
	case 'S', 'W':
		sign = -1
	default:
		return 0, false
	}
	parts := strings.Split(s[:len(s)-1], ",")
	if len(parts) < 2 || len(parts) > 3 {
		return 0, false
	}
	deg := 0.0
	for i, p := range parts {
		v, err := strconv.ParseFloat(p, 64)
		if err != nil {
			return 0, false
		}
		//degrees, minutes, seconds
		deg += v / [3]float64{1, 60, 3600}[i]
	}
	return sign * deg, true
}

// IsEmpty returns true if the xmp document is nil or has no nodes
func (xd XmpData) IsEmpty() bool {
	return xd.rawXmp == nil || len(xd.rawXmp.Nodes()) == 0
}

// PhotoShop retrieves the Photoshop Model
func (xd XmpData) PhotoShop() *ps.PhotoshopInfo {
	if !xd.IsEmpty() {
		return ps.FindModel(xd.rawXmp)
	}
	return nil
}

// MM retrieves the MediaManagement (MM) model
func (xd XmpData) MM() *xmpmm.XmpMM {
	if !xd.IsEmpty() {
		return xmpmm.FindModel(xd.rawXmp)
	}
	return nil
}

func (xd XmpData) String() string {
	if xd.IsEmpty() {
		return "No XMP Data"
	}
	if bytes, err := json.MarshalIndent(xd.rawXmp, "", "  "); err == nil {
		return string(bytes)
	}
	return "Could not marshal XmpEditor document"
}
