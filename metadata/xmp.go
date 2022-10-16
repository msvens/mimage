package metadata

import (
	"encoding/json"
	"errors"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"trimmer.io/go-xmp/models/dc"
	"trimmer.io/go-xmp/models/ps"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
	xmpmm "trimmer.io/go-xmp/models/xmp_mm"
	"trimmer.io/go-xmp/xmp"
)

// XmpData holds the underlying xmp document
type XmpData struct {
	rawXmp *xmp.Document
}

// ErrNoXmp when a jpeg image does not contain any xmp data
var ErrNoXmp = errors.New("No XMP data")

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
