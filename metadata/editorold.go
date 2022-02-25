package metadata

//import (
//	"bytes"
//	"fmt"
//	"github.com/dsoprea/go-exif/v3"
//	exifundefined "github.com/dsoprea/go-exif/v3/undefined"
//	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
//	"io/ioutil"
//	"path/filepath"
//	"time"
//	xmpbase "trimmer.io/go-XmpEditor/models/xmp_base"
//	xmpmm "trimmer.io/go-XmpEditor/models/xmp_mm"
//	"trimmer.io/go-XmpEditor/XmpEditor"
//)
//
//type CopyFlag uint8
//
//const (
//	EXIF CopyFlag = 1 << iota
//	XMP
//	IPTC
//)
//
//const CopyAll = EXIF | XMP | IPTC
//
//type MetaDataEditor struct {
//	sl        *jpegstructure.SegmentList
//	rootIb    *exif.IfdBuilder
//	dirtyExif bool
//}
//
//func NewMetaDataEditorFile(fileName string) (*MetaDataEditor, error) {
//	if b, err := ioutil.ReadFile(fileName); err != nil {
//		return nil, err
//	} else {
//		return NewMetaDataEditor(b)
//	}
//}
//func NewMetaDataEditor(imgSource []byte) (*MetaDataEditor, error) {
//	if segments, err := parseJpegBytes(imgSource); err != nil {
//		return nil, err
//	} else {
//		if rootIb, e1 := segments.ConstructExifBuilder(); e1 != nil {
//			return nil, e1
//		} else {
//			return &MetaDataEditor{segments, rootIb, false}, nil
//		}
//	}
//}
//
//func (mde *MetaDataEditor) Bytes() ([]byte, error) {
//	if err := mde.setExif(); err != nil {
//		return nil, err
//	}
//	out := new(bytes.Buffer)
//	if err := mde.sl.Write(out); err != nil {
//		return nil, err
//	} else {
//		return out.Bytes(), nil
//	}
//}
//
//func (mde *MetaDataEditor) SetXmp(doc *XmpEditor.Document) error {
//	//change creator tool:
//	t := time.Now()
//	base := xmpbase.FindModel(doc)
//	if base != nil {
//		base.CreatorTool = XmpEditorSoftware
//		base.ModifyDate = XmpEditor.NewDate(t)
//	}
//	mm := xmpmm.FindModel(doc)
//	if mm != nil {
//		re := xmpmm.ResourceEvent{}
//		re.Action = xmpmm.ActionSaved
//		re.SoftwareAgent = XmpEditorSoftware
//		re.When = XmpEditor.NewDate(t)
//		mm.AddHistory(&re)
//	}
//	buff := bytes.Buffer{}
//	buff.Write(xmpPrefix)
//	xmpBytes, err := XmpEditor.Marshal(doc)
//	if err != nil {
//		return err
//	}
//	buff.Write(xmpBytes)
//
//	_, s, err := mde.sl.FindXmp()
//	if err == nil { //replace existing XmpEditor data
//		s.Data = buff.Bytes()
//		return nil
//	} else if err == jpegstructure.ErrNoXmp { //add XmpEditor data
//		xmpS := &jpegstructure.Segment{MarkerId: jpegstructure.MARKER_APP1, Data: buff.Bytes()}
//		mde.appendSegment(1, xmpS)
//		return nil
//	} else {
//		return err
//	}
//}
//
//func (mde *MetaDataEditor) CopyMetaData(sourceImg []byte, sections CopyFlag) error {
//	sl, err := parseJpegBytes(sourceImg)
//	if err != nil {
//		return err
//	}
//	if sections&IPTC != 0 {
//		if err = mde.copyIptc(sl); err != nil {
//			return err
//		}
//	}
//	if sections&XMP != 0 {
//		if err = mde.copyXmp(sl); err != nil {
//			return err
//		}
//	}
//	if sections&EXIF != 0 {
//		if err = mde.copyExif(sl); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (mde *MetaDataEditor) CopyMetaDataFile(sourceImg string, sections CopyFlag) error {
//	if b, err := ioutil.ReadFile(sourceImg); err != nil {
//		return err
//	} else {
//		return mde.CopyMetaData(b, sections)
//	}
//}
//
//func (mde *MetaDataEditor) copyExif(sl *jpegstructure.SegmentList) error {
//	if mde.HasExif() {
//		if err := mde.DropExif(); err != nil {
//			return err
//		}
//	}
//	if _, s, e := sl.FindExif(); e == nil {
//		mde.appendSegment(1, s)
//		if rib, e1 := mde.sl.ConstructExifBuilder(); e1 != nil {
//			return e1
//		} else {
//			mde.rootIb = rib
//			mde.dirtyExif = false
//			return nil
//		}
//
//	} else if e == exif.ErrNoExif {
//		return nil
//	} else {
//		return e
//	}
//}
//
//func (mde *MetaDataEditor) copyIptc(sl *jpegstructure.SegmentList) error {
//	if err := mde.DropIptc(); err != nil {
//		return err
//	}
//	/*if mde.HasIptc() {
//		if err := mde.DropIptc(); err != nil {
//			return err
//		}
//	}*/
//	if _, s, e := sl.FindIptc(); e == nil {
//		mde.appendSegment(1, s)
//		return nil
//	} else if e == jpegstructure.ErrNoIptc {
//		return nil
//	} else {
//		return e
//	}
//
//}
//
//func (mde *MetaDataEditor) copyXmp(sl *jpegstructure.SegmentList) error {
//	if err := mde.DropXmp(); err != nil {
//		return err
//	}
//	/*
//	if mde.HasXmp() {
//		if err := mde.DropXmp(); err != nil {
//			return err
//		}
//	}*/
//
//	if _, s, e := sl.FindXmp(); e == nil {
//		mde.appendSegment(1, s)
//		return nil
//	} else if e == jpegstructure.ErrNoXmp {
//		return nil
//	} else {
//		return e
//	}
//}
//
//func (mde *MetaDataEditor) appendSegment(idx int, s *jpegstructure.Segment) {
//	newS := mde.sl.Segments()
//	newS = append(newS[:idx+1], newS[idx:]...)
//	newS[idx] = s
//	mde.sl = jpegstructure.NewSegmentList(newS)
//}
//
//func (mde *MetaDataEditor) DropAll() error {
//	if err := mde.DropExif(); err != nil {
//		return err
//	}
//	if err := mde.DropIptc(); err != nil {
//		return err
//	}
//	return mde.DropXmp()
//}
//
//func (mde *MetaDataEditor) DropExif() error {
//	_, err := mde.sl.DropExif()
//	if rbi, err := mde.sl.ConstructExifBuilder(); err != nil {
//		return err
//	} else {
//		mde.rootIb = rbi
//		mde.dirtyExif = false
//
//	}
//	return err
//}
//
//func (mde *MetaDataEditor) DropXmp() error {
//	i, _, err := mde.sl.FindXmp()
//	if err == nil {
//		segments := mde.sl.Segments()
//		segments = append(segments[:i], segments[i+1:]...)
//		mde.sl = jpegstructure.NewSegmentList(segments)
//		return nil
//	} else if err == jpegstructure.ErrNoXmp {
//		return nil
//	} else {
//		return err
//	}
//}
//
//func (mde *MetaDataEditor) DropIptc() error {
//	i, _, err := mde.sl.FindIptc()
//	if err == nil {
//		segments := mde.sl.Segments()
//		segments = append(segments[:i], segments[i+1:]...)
//		mde.sl = jpegstructure.NewSegmentList(segments)
//	} else if err == jpegstructure.ErrNoIptc {
//		return nil
//	}
//	return nil
//}
//
//func (mde *MetaDataEditor) PrintSegments() {
//	mde.sl.Print()
//}
//
//func (mde *MetaDataEditor) HasExif() bool {
//	if _, _, err := mde.sl.FindExif(); err != nil {
//		return false
//	} else {
//		return true
//	}
//}
//
//func (mde *MetaDataEditor) HasIptc() bool {
//	if _, _, err := mde.sl.FindIptc(); err != nil {
//		return false
//	} else {
//		return true
//	}
//}
//
//func (mde *MetaDataEditor) HasXmp() bool {
//	if _, _, err := mde.sl.FindXmp(); err != nil {
//		return false
//	} else {
//		return true
//	}
//}
//
//
////Retrives a metadata struct based on this editor. Will commit any changes first
//func (mde *MetaDataEditor) MetaData() (*MetaData, error) {
//	if b, e := mde.Bytes(); e != nil {
//		return nil, e
//	} else {
//		return NewMetaData(b)
//	}
//}
//
//func (mde *MetaDataEditor) SetExifTag(id ExifTag, value interface{}) error {
//
//	exifIb, err := exif.GetOrCreateIbFromRootIb(mde.rootIb, IFDPaths[ExifIFD])
//	if err != nil {
//		return err
//	} else {
//		if err := exifIb.SetStandard(uint16(id), toGoExifValue(value)); err != nil {
//			return err
//		}
//	}
//	mde.dirtyExif = true
//	return nil
//}
//
//func (mde *MetaDataEditor) SetIfdTag(id ExifTag, value interface{}) error {
//	if err := mde.rootIb.SetStandard(uint16(id), toGoExifValue(value)); err != nil {
//		return err
//	} else {
//		mde.dirtyExif = true
//		return nil
//	}
//}
//
//
//func (mde *MetaDataEditor) SetExifDate(dateTag ExifDate, time time.Time) error {
//	var err error
//	offset := exifOffsetString(time)
//	switch dateTag {
//	case OriginalDate:
//		if err = mde.SetExifTag(ExifIFD_DateTimeOriginal, time); err != nil {
//			return err
//		}
//		if err = mde.SetExifTag(ExifIFD_OffsetTimeOriginal, offset); err != nil {
//			return err
//		}
//	case ModifyDate:
//		if err = mde.SetIfdTag(IFD_ModifyDate, time); err != nil {
//			return err
//		}
//		if err = mde.SetExifTag(ExifIFD_OffsetTime, offset); err != nil {
//			return err
//		}
//	case DigitizedDate:
//		if err = mde.SetExifTag(ExifIFD_CreateDate, time); err != nil {
//			return err
//		}
//		if err = mde.SetExifTag(ExifIFD_OffsetTimeDigitized, offset); err != nil {
//			return err
//		}
//	default:
//		return fmt.Errorf("Unknown date to set: %v", dateTag)
//	}
//	return nil
//}
//
//func (mde *MetaDataEditor) SetImageDescription(description string) error {
//	return mde.SetIfdTag(IFD_ImageDescription, description)
//}
//
//func (mde *MetaDataEditor) SetUserComment(comment string) error {
//	uc := exifundefined.Tag9286UserComment{
//		EncodingType:  exifundefined.TagUndefinedType_9286_UserComment_Encoding_UNICODE,
//		EncodingBytes: []byte(comment),
//	}
//	return mde.SetExifTag(ExifIFD_UserComment, uc)
//}
//
//func (md *MetaDataEditor) CommitExifChanges() error {
//	return md.setExif()
//}
//
//func (mde *MetaDataEditor) setExif() error {
//	if mde.dirtyExif {
//		if err := mde.SetIfdTag(IFD_Software, ExifEditorSoftware); err != nil {
//			return err
//		}
//		if _, err := mde.sl.DropExif(); err != nil {
//			return err
//		}
//		if err := mde.sl.SetExif(mde.rootIb); err != nil {
//			return err
//		}
//		mde.dirtyExif = false
//		return nil
//	}
//	return nil
//}
//
//func (mde *MetaDataEditor) WriteFile(dest string) error {
//	//make sure dest has the right file extension
//	if filepath.Ext(dest) != ".jpg" && filepath.Ext(dest) != ".jpeg" {
//		return JpegWrongFileExtErr
//	}
//	if out, err := mde.Bytes(); err != nil {
//		return err
//	} else {
//		return ioutil.WriteFile(dest, out, 0644)
//	}
//}
