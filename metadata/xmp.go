package metadata

import (
	"encoding/json"
	"errors"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"trimmer.io/go-xmp/models/dc"
	"trimmer.io/go-xmp/models/ps"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
	xmpmm "trimmer.io/go-xmp/models/xmp_mm"
	"trimmer.io/go-xmp/xmp"
)

type XmpData struct {
	rawXmp *xmp.Document
}

var ErrNoXmp = errors.New("No XMP data")

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

func NewXmpDataFromBytes(data []byte) (XmpData, error) {
	model := &xmp.Document{}
	err := xmp.Unmarshal(data, model)
	if err != nil {
		return XmpData{}, ErrNoXmp
	} else {
		return XmpData{model}, nil
	}
}

func (xd XmpData) Base() *xmpbase.XmpBase {
	if !xd.IsEmpty() {
		return xmpbase.FindModel(xd.rawXmp)
	}
	return nil
}
func (xd XmpData) DublinCore() *dc.DublinCore {
	if !xd.IsEmpty() {
		return dc.FindModel(xd.rawXmp)
	}
	return nil
}

func (xd XmpData) GetKeywords() []string {
	if dcore := xd.DublinCore(); dcore != nil {
		return dcore.Subject
	} else {
		return []string{}
	}
}

func (xd XmpData) GetRating() uint16 {
	if base := xd.Base(); base != nil {
		return uint16(base.Rating)
	}
	return 0
}

func (xd XmpData) GetTitle() string {
	if dcore := xd.DublinCore(); dcore != nil {
		return dcore.Title.Default()
	} else {
		return ""
	}
}

func (xd XmpData) IsEmpty() bool {
	return xd.rawXmp == nil || len(xd.rawXmp.Nodes()) == 0
}

func (xd XmpData) PhotoShop() *ps.PhotoshopInfo {
	if !xd.IsEmpty() {
		return ps.FindModel(xd.rawXmp)
	}
	return nil
}

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
	if bytes, err := json.MarshalIndent(xd.rawXmp, "", "  "); err != nil {
		return fmt.Sprintf("Could not marshal XmpEditor document: %v", err)
	} else {
		return string(bytes)
	}
}
