package metadata

import (
	"reflect"
	"testing"
	"trimmer.io/go-xmp/xmp"
)

func getXmpEditorXmpFile(fatal bool, t *testing.T) *XmpEditor {
	b := getAssetBytes(XmpFile, t)
	ret, err := NewXmpEditorFromBytes(b)
	if err != nil {
		if fatal {
			t.Fatalf("Could not open xmp editor from bytes: %v", err)
		} else {
			t.Errorf("Could not open xmp editor from bytes: %v", err)
		}
	}
	return ret
}

func TestNewXmpEditor(t *testing.T) {
	images := []string{LeicaImg, NoExifImg}
	for _, fname := range images {
		b := getAssetBytes(fname, t)
		if sl, err := parseJpegBytes(b); err != nil {
			t.Errorf("Could not parse jpeg file %s", fname)
		} else {
			xe, e1 := NewXmpEditor(sl)
			if e1 != nil {
				t.Errorf("Could not open exif editor got %v", e1)
			}
			if xe.dirty {
				t.Errorf("expected dirty to be false")
			}
			if xe.rawXmp == nil {
				t.Errorf("expected rawXmp to be non nil")
			}
		}
	}
	if _, err := NewXmpEditor(nil); err == nil {
		t.Errorf("Expected error when creating a NewXmpEditor with a nil segment list")
	}

}

func TestNewXmpEditorFromBytes(t *testing.T) {
	xe := getXmpEditorXmpFile(true, t)
	if xe.dirty {
		t.Errorf("expected dirty to be false")
	}
	if xe.rawXmp == nil {
		t.Errorf("expected rawXmp to be no nil")
	}

	if _, err := NewXmpEditorFromBytes([]byte{}); err == nil {
		t.Errorf("Expected error when creating NewXmpEditor from zero bytes")
	}

}

func TestNewXmpEditorFromDocument(t *testing.T) {
	if xe, err := NewXmpEditorFromDocument(xmp.NewDocument()); err != nil {
		t.Errorf("Could not open xmp editor from document: %v", err)
	} else {
		if xe.dirty {
			t.Errorf("expected dirty to be false")
		}
		if xe.rawXmp == nil {
			t.Errorf("expected rawXmp to be no nil")
		}
	}
	if _, err := NewXmpEditorFromDocument(nil); err == nil {
		t.Errorf("Expected error when creating a n new XmpEditor from nil document")
	}
}

func TestXmpEditor_Bytes(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	xmpE := je.Xmp()
	xmpE.SetDirty()
	//first write bytes and reopen without prefix
	b, err := xmpE.Bytes(false)
	if err != nil {
		t.Fatalf("Could not write xmp to bytes: %v", err)
	}
	xmpE, err = NewXmpEditorFromBytes(b)
	if err != nil {
		t.Fatalf("Could not reopen xmp after written to bytes: %v", err)
	}
	//make sure that software tag was set
	base := xmpE.Base()
	if base == nil {
		t.Fatalf("xmp document did not have a base model")
	}
	if base.CreatorTool.String() != xmpEditorSoftware {
		t.Fatalf("Expected xmp creator tool %s got %s", xmpEditorSoftware, base.CreatorTool.String())
	}
	//now write and reopen with prefix
	je.Xmp().SetDirty()
	je = reloadJpegEditor(je, true, t)
	base = je.Xmp().Base()
	if base == nil {
		t.Fatalf("xmp document did not have a base model")
	}
	if base.CreatorTool.String() != xmpEditorSoftware {
		t.Fatalf("Expected xmp creator tool %s got %s", xmpEditorSoftware, base.CreatorTool.String())
	}
}

func TestXmpEditor_Clear(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	je.Xmp().Clear(true)
	if !je.Xmp().IsEmpty() {
		t.Errorf("Expected empty xmp document")
	}
	//make wure we can write an empty document if we want to
	je = reloadJpegEditor(je, true, t)
	base := je.Xmp().Base()
	if base == nil {
		t.Fatalf("xmp document did not have a base model")
	}
	if base.CreatorTool.String() != xmpEditorSoftware {
		t.Fatalf("Expected xmp creator tool %s got %s", xmpEditorSoftware, base.CreatorTool.String())
	}

}

func TestXmpEditor_Document(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	if je.Xmp().Document() == nil {
		t.Errorf("Expected document got nil")
	}
	je.Xmp().Clear(false)
	if je.Xmp().Document() == nil {
		t.Errorf("Expected document got nil")
	}
}

func TestXmpEditor_IsDirty(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	if je.Xmp().IsDirty() {
		t.Errorf("Expected not dirty xmp editor")
	}
	je.Xmp().SetTitle("some new title")
	if !je.Xmp().IsDirty() {
		t.Errorf("Expected dirty xmp editor after setting title")
	}
	_, _ = je.Bytes()
	if je.Xmp().IsDirty() {
		t.Errorf("Expected not dirty after retrieving bytes")
	}
	je.Xmp().SetDirty()
	if !je.Xmp().IsDirty() {
		t.Errorf("Expected dirty after forcing dirty")
	}

}

func TestXmpEditor_SetDirty(t *testing.T) {
	je := getJpegEditor(LeicaImg, t)
	if je.Xmp().IsDirty() {
		t.Errorf("Expected not dirty")
	}
	je.Xmp().SetDirty()
	if !je.Xmp().IsDirty() {
		t.Errorf("Expected not dirty")
	}
}

func TestXmpEditor_SetDocument(t *testing.T) {
	b := getAssetBytes(XmpFile, t)
	doc := &xmp.Document{}
	if err := xmp.Unmarshal(b, doc); err != nil {
		t.Fatalf("Could not unmarshal xmp document")
	}
	xe, err := NewXmpEditorFromDocument(xmp.NewDocument())
	if err != nil {
		t.Errorf("Could not create xmpeditor")
	}
	xe.SetDocument(doc, true)
	if !xe.IsDirty() {
		t.Errorf("Expected dirty")
	}
	_ = xe.Document()
	//now check title
	expTitle := "New Title"
	if xe.DublinCore().Title.Default() != expTitle {
		t.Errorf("Expected %s got %s", expTitle, xe.DublinCore().Title.Default())
	}

}

func TestXmpEditor_SetKeywords(t *testing.T) {
	expKeywords := []string{"keyword 1", "keyword 2", "keyword 3"}
	xePop := getXmpEditorXmpFile(true, t)
	xeEmpty, _ := NewXmpEditorFromDocument(xmp.NewDocument())
	tests := []*XmpEditor{xePop, xeEmpty}
	for _, xe := range tests {
		xe.SetKeywords(expKeywords)
		if !xe.IsDirty() {
			t.Errorf("expected dirty xmp editor")
		}
		_ = xe.Document()
		if !reflect.DeepEqual(xe.GetKeywords(), expKeywords) {
			t.Errorf("expected %v got %v", expKeywords, xe.GetKeywords())
		}
	}
}

func TestXmpEditor_SetRating(t *testing.T) {
	expRating := uint16(4)
	xePop := getXmpEditorXmpFile(true, t)
	xeEmpty, _ := NewXmpEditorFromDocument(xmp.NewDocument())
	tests := []*XmpEditor{xePop, xeEmpty}
	for _, xe := range tests {
		xe.SetRating(expRating)
		if !xe.IsDirty() {
			t.Errorf("expected dirty xmp editor")
		}
		_ = xe.Document()
		if xe.GetRating() != expRating {
			t.Errorf("expected %v got %v", expRating, xe.GetRating())
		}
	}
}

func TestXmpEditor_SetTitle(t *testing.T) {
	expTitle := "some cool title"
	xePop := getXmpEditorXmpFile(true, t)
	xeEmpty, _ := NewXmpEditorFromDocument(xmp.NewDocument())
	tests := []*XmpEditor{xePop, xeEmpty}
	for _, xe := range tests {
		xe.SetTitle(expTitle)
		if !xe.IsDirty() {
			t.Errorf("expected dirty xmp editor")
		}
		_ = xe.Document()
		if xe.GetTitle() != expTitle {
			t.Errorf("expected %v got %v", expTitle, xe.GetTitle())
		}
	}
}
