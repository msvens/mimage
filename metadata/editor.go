package metadata

import (
	"bytes"
	"errors"
	"github.com/dsoprea/go-exif/v3"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"os"
	"path/filepath"
)

var errJpegWrongFileExt = errors.New("File does not end with .jpg or .jpeg")

// JpegEditor holds the exif, xmp and iptc editors as well as the jpeg segment list
type JpegEditor struct {
	sl *jpegstructure.SegmentList
	xe *XmpEditor
	ee *ExifEditor
	ie *IptcEditor
}

// NewJpegEditorFile from a jpeg image file
func NewJpegEditorFile(fileName string) (*JpegEditor, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return NewJpegEditor(b)
}

// NewJpegEditor from a jpeg image byte slice
func NewJpegEditor(data []byte) (*JpegEditor, error) {
	var err error
	ret := JpegEditor{}
	if ret.sl, err = parseJpegBytes(data); err != nil {
		return &ret, err
	}
	if ret.xe, err = NewXmpEditor(ret.sl); err != nil && err != ErrNoXmp {
		return &ret, err
	}
	if ret.ee, err = NewExifEditor(ret.sl); err != nil {
		return &ret, err
	}
	if ret.ie, err = NewIptcEditor(ret.sl); err != nil {
		return &ret, err
	}
	return &ret, nil
}

func (je *JpegEditor) appendSegment(idx int, s *jpegstructure.Segment) {
	newS := je.sl.Segments()
	newS = append(newS[:idx+1], newS[idx:]...)
	newS[idx] = s
	je.sl = jpegstructure.NewSegmentList(newS)
}

// Bytes return jpeg image bytes from this editor. Any edits will be committed
func (je *JpegEditor) Bytes() ([]byte, error) {
	if je.ie.IsDirty() {
		if err := je.setIptc(); err != nil {
			return nil, err
		}
	}
	if je.xe.IsDirty() {
		if err := je.setXmp(); err != nil {
			return nil, err
		}
	}
	if je.ee.IsDirty() {
		if err := je.setExif(); err != nil {
			return nil, err
		}
	}
	out := new(bytes.Buffer)
	if err := je.sl.Write(out); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// CopyMetaData copies and replaces metadata (xmp,iptc,exif) from sourcImg
func (je *JpegEditor) CopyMetaData(sourceImg []byte) error {
	sl, err := parseJpegBytes(sourceImg)
	if err != nil {
		return err
	}
	//copy iptc
	if err = je.DropIptc(); err != nil {
		return err
	}
	if _, s, e := sl.FindIptc(); e == nil {
		je.appendSegment(1, s)
	} else if e != jpegstructure.ErrNoIptc {
		return e
	}

	//copy XmpEditor
	if err = je.DropXmp(); err != nil {
		return err
	}
	if _, s, e := sl.FindXmp(); e == nil {
		je.appendSegment(1, s)
		if je.xe, err = NewXmpEditor(sl); err != nil {
			return err
		}
	} else if e != jpegstructure.ErrNoXmp {
		return e
	}
	//copy exif
	if err = je.DropExif(); err != nil {
		return err
	}
	if _, s, e := sl.FindExif(); e == nil {
		je.appendSegment(1, s)
		if je.ee, err = NewExifEditor(sl); err != nil {
			return err
		}
	} else if !errors.Is(e, exif.ErrNoExif) {
		return e
	}
	return nil
}

// DropMetaData removes xmp, exif, iptc data from this editor
func (je *JpegEditor) DropMetaData() error {
	if err := je.DropExif(); err != nil {
		return err
	}
	if err := je.DropIptc(); err != nil {
		return err
	}
	return je.DropXmp()
}

// DropExif removes exif data from this editor
func (je *JpegEditor) DropExif() error {
	if _, err := je.sl.DropExif(); err != nil {
		return err
	}
	return je.ee.Clear(false)
}

// DropXmp removes xmp data from this editor
func (je *JpegEditor) DropXmp() error {
	i, _, err := je.sl.FindXmp()
	if err == nil {
		segments := je.sl.Segments()
		segments = append(segments[:i], segments[i+1:]...)
		je.sl = jpegstructure.NewSegmentList(segments)
		je.xe.Clear(false)
		return nil
	} else if err == jpegstructure.ErrNoXmp {
		je.xe.Clear(false)
		return nil
	} else {
		return err
	}
}

// DropIptc removes iptc data from this editor
func (je *JpegEditor) DropIptc() error {
	i, _, err := je.sl.FindIptc()
	if err == nil {
		segments := je.sl.Segments()
		segments = append(segments[:i], segments[i+1:]...)
		je.sl = jpegstructure.NewSegmentList(segments)
	} else if err == jpegstructure.ErrNoIptc {
		return nil
	}
	return nil
}

// Exif returns the exifEditor
func (je JpegEditor) Exif() *ExifEditor {
	return je.ee
}

// Iptc returns the iptcEditor
func (je JpegEditor) Iptc() *IptcEditor {
	return je.ie
}

// MetaData returns a metadata struct based on this editor. Will commit any changes first
func (je *JpegEditor) MetaData() (*MetaData, error) {
	b, e := je.Bytes()
	if e != nil {
		return nil, e
	}
	return NewMetaData(b)
}

func (je *JpegEditor) setExif() error {
	if _, err := je.sl.DropExif(); err != nil {
		return err
	}
	builder, _ := je.ee.IfdBuilder()
	if err := je.sl.SetExif(builder); err != nil {
		return err
	}
	return nil
}

func (je *JpegEditor) setIptc() error {
	//MarkerId: MARKER_APP13
	iptcBytes, err := je.ie.Bytes()
	if err != nil {
		return err
	}
	if je.ie.segmentIdx != -1 {
		s := je.sl.Segments()[je.ie.segmentIdx]
		s.Data = iptcBytes
		return nil
	}
	iptcS := &jpegstructure.Segment{MarkerId: jpegstructure.MARKER_APP13, Data: iptcBytes}
	je.appendSegment(1, iptcS)
	je.ie.segmentIdx = 1
	return nil

}

// SetKeywords sets the keywords in Xmp and Iptc
func (je *JpegEditor) SetKeywords(keywords []string) error {
	je.Xmp().SetKeywords(keywords)
	return je.ie.SetKeywords(keywords)
}

// SetTitle sets title in Xmp and Iptc and ImageDescription in Exif
func (je *JpegEditor) SetTitle(title string) error {
	je.Xmp().SetTitle(title)
	if err := je.ee.SetImageDescription(title); err != nil {
		return err
	}
	return je.Iptc().SetTitle(title)

}

func (je *JpegEditor) setXmp() error {
	xmpBytes, err := je.xe.Bytes(true)
	if err != nil {
		return err
	}
	_, s, err := je.sl.FindXmp()
	if err == nil { //replace existing XmpEditor data
		s.Data = xmpBytes
		return nil
	} else if err == jpegstructure.ErrNoXmp { //add XmpEditor data
		xmpS := &jpegstructure.Segment{MarkerId: jpegstructure.MARKER_APP1, Data: xmpBytes}
		je.appendSegment(1, xmpS)
		return nil
	} else {
		return err
	}
}

// WriteFile writes this editor to file by first committing any edits. Any existing
// file will be truncated. Destination needs to have jpg or jpeg extension
func (je *JpegEditor) WriteFile(dest string) error {
	//make sure dest has the right file extension
	if filepath.Ext(dest) != ".jpg" && filepath.Ext(dest) != ".jpeg" {
		return errJpegWrongFileExt
	}
	out, err := je.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(dest, out, 0644)

}

// Xmp returns the xmp editor
func (je JpegEditor) Xmp() *XmpEditor {
	return je.xe
}
