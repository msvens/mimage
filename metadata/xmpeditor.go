package metadata

import (
	"bytes"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"time"
	"trimmer.io/go-xmp/models/dc"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
	xmpmm "trimmer.io/go-xmp/models/xmp_mm"
	"trimmer.io/go-xmp/xmp"
)

const XmpEditorSoftware = "github.com/msvens/mimage (go-XmpEditor)"

var xmpPrefix = []byte("http://ns.adobe.com/xap/1.0/\000")

type XmpEditor struct {
	XmpData
	dirty bool
}

func NewXmpEditor(sl *jpegstructure.SegmentList) (*XmpEditor, error) {
	if sl == nil {
		return &XmpEditor{}, fmt.Errorf("nil segment list")
	}
	if xmpData, err := NewXmpData(sl); err != nil {
		ret := XmpEditor{}
		ret.Clear(false)
		return &ret, nil
	} else {
		return &XmpEditor{
			XmpData{
				xmpData.rawXmp,
			},
			false,
		}, nil
	}
}

func NewXmpEditorFromDocument(doc *xmp.Document) (*XmpEditor, error) {
	if doc == nil {
		return nil, fmt.Errorf("Document is nil")
	}
	return &XmpEditor{
		XmpData{
			doc,
		},
		false,
	}, nil
}

func NewXmpEditorFromBytes(data []byte) (*XmpEditor, error) {
	if xmpData, err := NewXmpDataFromBytes(data); err != nil {
		return &XmpEditor{}, err
	} else {
		return &XmpEditor{
			XmpData{
				xmpData.rawXmp,
			},
			false,
		}, nil
	}
}

func (xe *XmpEditor) Bytes(prefix bool) ([]byte, error) {

	xe.setSoftware()

	if !prefix {
		return xmp.Marshal(xe.rawXmp)
	} else {
		buff := bytes.Buffer{}
		if b, err := xmp.Marshal(xe.rawXmp); err != nil {
			return b, err
		} else {
			buff.Write(xmpPrefix)
			buff.Write(b)
			return buff.Bytes(), nil
		}
	}

}

func (xe *XmpEditor) Document() *xmp.Document {
	xe.setSoftware()
	return xe.rawXmp
}

func (xe XmpEditor) IsDirty() bool {
	return xe.dirty
}

func (xe *XmpEditor) Clear(dirty bool) {
	xe.rawXmp = xmp.NewDocument()
	xe.dirty = dirty
}

func (xe *XmpEditor) SetDirty() {
	xe.dirty = true
}

func (xe *XmpEditor) SetDocument(doc *xmp.Document, dirty bool) {
	xe.rawXmp = doc
	xe.dirty = dirty
}

func (xe *XmpEditor) SetKeywords(keywords []string) {
	dcore := xe.dcOrCreate()
	dcore.Subject = xmp.NewStringArray(keywords...)
	xe.dirty = true
}

func (xe *XmpEditor) SetRating(rating uint16) {
	base := xe.baseOrCreate()
	base.Rating = xmpbase.Rating(rating)
	xe.dirty = true
}

func (xe *XmpEditor) SetTitle(title string) {
	dcore := xe.dcOrCreate()
	dcore.Title.Set("", title)
	xe.dirty = true
}

func (xe *XmpEditor) setSoftware() {
	if !xe.IsDirty() {
		return
	}
	t := time.Now()
	base := xe.baseOrCreate()
	if base != nil {
		base.CreatorTool = XmpEditorSoftware
		base.ModifyDate = xmp.NewDate(t)
	}
	mm := xe.mmOrCreate()
	if mm != nil {
		re := xmpmm.ResourceEvent{}
		re.Action = xmpmm.ActionSaved
		re.SoftwareAgent = XmpEditorSoftware
		re.When = xmp.NewDate(t)
		mm.AddHistory(&re)
	}
	xe.dirty = false
}

func (xe *XmpEditor) dcOrCreate() *dc.DublinCore {
	if ret := xe.DublinCore(); ret != nil {
		return ret
	}
	if m, err := xe.rawXmp.MakeModel(dc.NsDc); err != nil {
		return nil
	} else {
		return m.(*dc.DublinCore)
	}
}

func (xe *XmpEditor) mmOrCreate() *xmpmm.XmpMM {
	ret := xe.MM()
	if ret != nil {
		return ret
	}
	if m, err := xe.rawXmp.MakeModel(xmpmm.NsXmpMM); err != nil {
		return nil
	} else {
		return m.(*xmpmm.XmpMM)
	}
}

func (xe *XmpEditor) baseOrCreate() *xmpbase.XmpBase {
	ret := xe.Base()
	if ret != nil {
		return ret
	}
	if m, err := xe.rawXmp.MakeModel(xmpbase.NsXmp); err != nil {
		return nil
	} else {
		return m.(*xmpbase.XmpBase)
	}
}
