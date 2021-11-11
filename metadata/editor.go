package metadata

import (
	"bytes"
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"io/ioutil"
	"time"
	"trimmer.io/go-xmp/xmp"
)

const ExifPath = "IFD/Exif"
const EditorSoftware = "github.com/msvens/mexif (go-exif)"

var xmpPrefix = []byte("http://ns.adobe.com/xap/1.0/\000")

type ExifDate int

type CopyFlag uint8

const (
	EXIF CopyFlag = 1 << iota
	XMP
	IPTC
)

const CopyAll = EXIF | XMP | IPTC

const (
	OriginalDate ExifDate = iota
	ModifyDate
	DigitizedDate
)

type MetaDataEditor struct {
	sl        *jpegstructure.SegmentList
	rootIb    *exif.IfdBuilder
	dirtyExif bool
}

func NewMetaDataEditorFile(fileName string) (*MetaDataEditor, error) {
	if b, err := ioutil.ReadFile(fileName); err != nil {
		return nil, err
	} else {
		return NewMetaDataEditor(b)
	}
}
func NewMetaDataEditor(imgSource []byte) (*MetaDataEditor, error) {
	if segments, err := parseJpegBytes(imgSource); err != nil {
		return nil, err
	} else {
		if rootIb, e1 := segments.ConstructExifBuilder(); e1 != nil {
			return nil, e1
		} else {
			return &MetaDataEditor{segments, rootIb, false}, nil
		}
	}
}

func (mde *MetaDataEditor) Bytes() ([]byte, error) {
	if mde.dirtyExif {
		if err := mde.SetIfdTag(IFD_Software, EditorSoftware); err != nil {
			return nil, err
		}
	}
	if err := mde.setExif(); err != nil {
		return nil, err
	}
	out := new(bytes.Buffer)
	if err := mde.sl.Write(out); err != nil {
		return nil, err
	} else {
		return out.Bytes(), nil
	}
}

func (mde *MetaDataEditor) SetXmp(doc *xmp.Document) error {

	buff := bytes.Buffer{}
	buff.Write(xmpPrefix)
	xmpBytes, err := xmp.Marshal(doc)
	if err != nil {
		return err
	}
	buff.Write(xmpBytes)

	_, s, err := mde.sl.FindXmp()
	if err == nil { //replace existing xmp data
		s.Data = buff.Bytes()
		return nil
	} else if err == jpegstructure.ErrNoXmp { //add xmp data
		xmpS := &jpegstructure.Segment{MarkerId: jpegstructure.MARKER_APP1, Data: buff.Bytes()}
		mde.appendSegment(1, xmpS)
		return nil
	} else {
		return err
	}
}

func (mde *MetaDataEditor) CopyMetaData(sourceImg []byte, sections CopyFlag) error {
	sl, err := parseJpegBytes(sourceImg)
	if err != nil {
		return err
	}
	if sections&IPTC != 0 {
		if err = mde.copyIptc(sl); err != nil {
			return err
		}
	}
	if sections&XMP != 0 {
		if err = mde.copyXmp(sl); err != nil {
			return err
		}
	}
	if sections&EXIF != 0 {
		if err = mde.copyExif(sl); err != nil {
			return err
		}
	}
	return nil
}

func (mde *MetaDataEditor) CopyMetaDataFile(sourceImg string, sections CopyFlag) error {
	if b, err := ioutil.ReadFile(sourceImg); err != nil {
		return err
	} else {
		return mde.CopyMetaData(b, sections)
	}
}

func (mde *MetaDataEditor) copyExif(sl *jpegstructure.SegmentList) error {
	if mde.HasExif() {
		if err := mde.DropExif(); err != nil {
			return err
		}
	}
	if _, s, e := sl.FindExif(); e == nil {
		mde.appendSegment(1, s)
		if rib, e1 := mde.sl.ConstructExifBuilder(); e1 != nil {
			return e1
		} else {
			mde.rootIb = rib
			mde.dirtyExif = false
			return nil
		}

	} else if e == exif.ErrNoExif {
		return nil
	} else {
		return e
	}
}

func (mde *MetaDataEditor) copyIptc(sl *jpegstructure.SegmentList) error {
	if mde.HasIptc() {
		if err := mde.DropIptc(); err != nil {
			return err
		}
	}
	if _, s, e := sl.FindIptc(); e == nil {
		mde.appendSegment(1, s)
		return nil
	} else if e == jpegstructure.ErrNoIptc {
		return nil
	} else {
		return e
	}

}

func (mde *MetaDataEditor) copyXmp(sl *jpegstructure.SegmentList) error {
	if mde.HasXmp() {
		if err := mde.DropXmp(); err != nil {
			return err
		}
	}
	if _, s, e := sl.FindXmp(); e == nil {
		mde.appendSegment(1, s)
		return nil
	} else if e == jpegstructure.ErrNoXmp {
		return nil
	} else {
		return e
	}
}

func (mde *MetaDataEditor) appendSegment(idx int, s *jpegstructure.Segment) {
	newS := mde.sl.Segments()
	newS = append(newS[:idx+1], newS[idx:]...)
	newS[idx] = s
	mde.sl = jpegstructure.NewSegmentList(newS)
}

func (mde *MetaDataEditor) DropAll() error {
	if err := mde.DropExif(); err != nil {
		return err
	}
	if err := mde.DropIptc(); err != nil {
		return err
	}
	return mde.DropXmp()
}

func (mde *MetaDataEditor) DropExif() error {
	_, err := mde.sl.DropExif()
	if rbi, err := mde.sl.ConstructExifBuilder(); err != nil {
		return err
	} else {
		mde.rootIb = rbi
		mde.dirtyExif = false

	}
	return err
}

func (mde *MetaDataEditor) DropXmp() error {
	i, _, err := mde.sl.FindXmp()
	if err == nil {
		segments := mde.sl.Segments()
		segments = append(segments[:i], segments[i+1:]...)
		mde.sl = jpegstructure.NewSegmentList(segments)
		return nil
	} else if err == jpegstructure.ErrNoXmp {
		return nil
	} else {
		return err
	}
}

func (mde *MetaDataEditor) DropIptc() error {
	i, _, err := mde.sl.FindIptc()
	if err == nil {
		segments := mde.sl.Segments()
		segments = append(segments[:i], segments[i+1:]...)
		mde.sl = jpegstructure.NewSegmentList(segments)
	} else if err == jpegstructure.ErrNoIptc {
		return nil
	}
	return nil
}

func (mde *MetaDataEditor) PrintSegments() {
	mde.sl.Print()
}

func (mde *MetaDataEditor) HasExif() bool {
	if _, _, err := mde.sl.FindExif(); err != nil {
		return false
	} else {
		return true
	}
}

func (mde *MetaDataEditor) HasIptc() bool {
	if _, _, err := mde.sl.FindIptc(); err != nil {
		return false
	} else {
		return true
	}
}

func (mde *MetaDataEditor) HasXmp() bool {
	if _, _, err := mde.sl.FindXmp(); err != nil {
		return false
	} else {
		return true
	}
}

func (mde *MetaDataEditor) MetaData() (*MetaData, error) {
	if b, e := mde.Bytes(); e != nil {
		return nil, e
	} else {
		return NewMetaDataJpeg(b)
	}
}

func (mde *MetaDataEditor) SetExifTag(id uint16, value interface{}) error {

	exifIb, err := exif.GetOrCreateIbFromRootIb(mde.rootIb, ExifPath)
	if err != nil {
		return err
	} else {
		if err := exifIb.SetStandard(id, toGoExifValue(value)); err != nil {
			return err
		}
	}
	mde.dirtyExif = true
	return nil
}

func (mde *MetaDataEditor) SetIfdTag(id uint16, value interface{}) error {
	if err := mde.rootIb.SetStandard(id, toGoExifValue(value)); err != nil {
		return err
	} else {
		mde.dirtyExif = true
		return nil
	}
}

func toGoExifValue(value interface{}) interface{} {
	switch t := value.(type) {
	case uint16:
		return []uint16{t}
	case uint32:
		return []uint32{t}
	case float32:
		return []float32{t}
	case float64:
		return []float64{t}
	case int32:
		return []int32{t}
	case URat:
		return []exifcommon.Rational{{Denominator: t.Denominator, Numerator: t.Numerator}}
	case Rat:
		return []exifcommon.SignedRational{{Denominator: t.Denominator, Numerator: t.Numerator}}
	case []URat:
		var ret []exifcommon.Rational
		for _, v := range t {
			ret = append(ret, exifcommon.Rational{Denominator: v.Denominator, Numerator: v.Numerator})
		}
		return ret
	case []Rat:
		var ret []exifcommon.SignedRational
		for _, v := range t {
			ret = append(ret, exifcommon.SignedRational{Denominator: v.Denominator, Numerator: v.Numerator})
		}
		return ret
	default:
		return t
	}
}

func (mde *MetaDataEditor) SetExifDate(dateTag ExifDate, time time.Time) error {
	var err error
	offset := TimeOffsetString(time)
	switch dateTag {
	case OriginalDate:
		if err = mde.SetExifTag(Exif_DateTimeOriginal, time); err != nil {
			return err
		}
		if err = mde.SetExifTag(Exif_OffsetTimeOriginal, offset); err != nil {
			return err
		}
	case ModifyDate:
		if err = mde.SetIfdTag(IFD_DateTime, time); err != nil {
			return err
		}
		if err = mde.SetExifTag(Exif_OffsetTime, offset); err != nil {
			return err
		}
	case DigitizedDate:
		if err = mde.SetExifTag(Exif_DateTimeDigitized, time); err != nil {
			return err
		}
		if err = mde.SetExifTag(Exif_OffsetTimeDigitized, offset); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown date to set: %v", dateTag)
	}
	return nil
}

func (mde *MetaDataEditor) SetImageDescription(description string) error {
	return mde.SetIfdTag(IFD_ImageDescription, description)
}

func (mde *MetaDataEditor) SetUserComment(comment string) error {
	return mde.SetExifTag(Exif_UserComment, comment)
}

func (mde *MetaDataEditor) setExif() error {
	if mde.dirtyExif {
		//drop existing exif first
		if _, err := mde.sl.DropExif(); err != nil {
			return err
		}
		if err := mde.sl.SetExif(mde.rootIb); err != nil {
			return err
		}
		mde.dirtyExif = false
		return nil
	}
	return nil
}

func (mde *MetaDataEditor) WriteFile(dest string) error {
	if out, err := mde.Bytes(); err != nil {
		return err
	} else {
		return ioutil.WriteFile(dest, out, 0644)
	}
}
