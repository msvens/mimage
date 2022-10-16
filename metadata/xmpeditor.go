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

const xmpEditorSoftware = "github.com/msvens/mimage (go-XmpEditor)"

var xmpPrefix = []byte("http://ns.adobe.com/xap/1.0/\000")

// XmpEditor add a dirty field to the XmpData
type XmpEditor struct {
	XmpData
	dirty bool
}

// NewXmpEditor from a jpeg segment list
func NewXmpEditor(sl *jpegstructure.SegmentList) (*XmpEditor, error) {
	if sl == nil {
		return &XmpEditor{}, fmt.Errorf("nil segment list")
	}
	xmpData, err := NewXmpData(sl)
	if err != nil {
		ret := XmpEditor{}
		ret.Clear(false)
		return &ret, nil
	}
	return &XmpEditor{
		XmpData{
			xmpData.rawXmp,
		},
		false,
	}, nil

}

// NewXmpEditorFromDocument from an xmp.document
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

// NewXmpEditorFromBytes from a marshalled xmp.Document
func NewXmpEditorFromBytes(data []byte) (*XmpEditor, error) {
	xmpData, err := NewXmpDataFromBytes(data)
	if err != nil {
		return &XmpEditor{}, err
	}
	return &XmpEditor{
		XmpData{
			xmpData.rawXmp,
		},
		false,
	}, nil

}

// Bytes commits any changes and writes the xmpDocument to bytes. If prefix
// is adds "http://ns.adobe.com/xap/1.0/\000" so it can be added to a jpeg segment list
func (xe *XmpEditor) Bytes(prefix bool) ([]byte, error) {
	xe.setSoftware()
	if !prefix {
		return xmp.Marshal(xe.rawXmp)
	}
	buff := bytes.Buffer{}
	b, err := xmp.Marshal(xe.rawXmp)
	if err != nil {
		return b, err
	}
	buff.Write(xmpPrefix)
	buff.Write(b)
	return buff.Bytes(), nil
}

// Document commits any changes and returns the xmp.Document
func (xe *XmpEditor) Document() *xmp.Document {
	xe.setSoftware()
	return xe.rawXmp
}

// IsDirty true if any changes has been made to this editor
func (xe XmpEditor) IsDirty() bool {
	return xe.dirty
}

// Clear the underlying xmp.Document and sets the dirty flag
func (xe *XmpEditor) Clear(dirty bool) {
	xe.rawXmp = xmp.NewDocument()
	xe.dirty = dirty
}

// SetDirty force this editor to be marked as dirty
func (xe *XmpEditor) SetDirty() {
	xe.dirty = true
}

// SetDocument sets the xmp.Document of this editor
func (xe *XmpEditor) SetDocument(doc *xmp.Document, dirty bool) {
	xe.rawXmp = doc
	xe.dirty = dirty
}

// SetKeywords sets the DublicCore keywords
func (xe *XmpEditor) SetKeywords(keywords []string) {
	dcore := xe.dcOrCreate()
	dcore.Subject = xmp.NewStringArray(keywords...)
	xe.dirty = true
}

// SetRating sets the Base rating
func (xe *XmpEditor) SetRating(rating uint16) {
	base := xe.baseOrCreate()
	base.Rating = xmpbase.Rating(rating)
	xe.dirty = true
}

// SetTitle sets the Dublin Core title
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
		base.CreatorTool = xmpEditorSoftware
		base.ModifyDate = xmp.NewDate(t)
	}
	mm := xe.mmOrCreate()
	if mm != nil {
		re := xmpmm.ResourceEvent{}
		re.Action = xmpmm.ActionSaved
		re.SoftwareAgent = xmpEditorSoftware
		re.When = xmp.NewDate(t)
		mm.AddHistory(&re)
	}
	xe.dirty = false
}

func (xe *XmpEditor) dcOrCreate() *dc.DublinCore {
	if ret := xe.DublinCore(); ret != nil {
		return ret
	}
	m, err := xe.rawXmp.MakeModel(dc.NsDc)
	if err != nil {
		return nil
	}
	return m.(*dc.DublinCore)
}

func (xe *XmpEditor) mmOrCreate() *xmpmm.XmpMM {
	ret := xe.MM()
	if ret != nil {
		return ret
	}
	m, err := xe.rawXmp.MakeModel(xmpmm.NsXmpMM)
	if err != nil {
		return nil
	}
	return m.(*xmpmm.XmpMM)

}

func (xe *XmpEditor) baseOrCreate() *xmpbase.XmpBase {
	ret := xe.Base()
	if ret != nil {
		return ret
	}
	m, err := xe.rawXmp.MakeModel(xmpbase.NsXmp)
	if err != nil {
		return nil
	}
	return m.(*xmpbase.XmpBase)
}
