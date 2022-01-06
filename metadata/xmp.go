package metadata

import (
	"errors"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"trimmer.io/go-xmp/models/dc"
	"trimmer.io/go-xmp/models/ps"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
	"trimmer.io/go-xmp/xmp"
)

type XmpData struct {
	rawXmp *xmp.Document
}

var NoXmpErr = errors.New("No XMP data")

func NewXmpData(segments *jpegstructure.SegmentList) (XmpData, error) {
	_, s, err := segments.FindXmp()
	if err != nil {
		return XmpData{}, NoXmpErr
	}
	str, err := s.FormattedXmp()
	if err != nil {
		//We should log errors
		return XmpData{}, NoXmpErr
	}
	model := &xmp.Document{}

	err = xmp.Unmarshal([]byte(str), model)
	if err != nil {
		return XmpData{}, NoXmpErr
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

func (xd XmpData) IsEmpty() bool {
	return xd.rawXmp == nil
}

func (xd XmpData) PhotoShop() *ps.PhotoshopInfo {
	if !xd.IsEmpty() {
		return ps.FindModel(xd.rawXmp)
	}
	return nil
}

func (xd XmpData) String() string {
	if xd.IsEmpty() {
		return "No XMP Data"
	}
	if bytes, err := xmp.MarshalIndent(xd.rawXmp, "", "  "); err != nil {
		return fmt.Sprintf("Could not marshal xmp document: %v", err)
	} else {
		return string(bytes)
	}
}
