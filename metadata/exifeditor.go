package metadata

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	exifundefined "github.com/dsoprea/go-exif/v3/undefined"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"github.com/go-errors/errors"
	"time"
)

const exifEditorSoftware = "github.com/msvens/mimage (go-exif)"

// ExifEditor holds an IfdBuilder
type ExifEditor struct {
	rootIb *exif.IfdBuilder
	dirty  bool
}

func exifOffsetString(t time.Time) string {
	return t.Format("-0700")
}

// NewExifEditor from a jpeg segment list
func NewExifEditor(sl *jpegstructure.SegmentList) (*ExifEditor, error) {
	if sl == nil {
		return &ExifEditor{}, fmt.Errorf("nil segment list")
	}
	rootIfd, _, err := sl.Exif()
	if err != nil {
		if errors.Is(err, exif.ErrNoExif) {
			return NewExifEditorEmpty(false)
		}
		return &ExifEditor{}, err
	}
	rootIb := exif.NewIfdBuilderFromExistingChain(rootIfd)
	return &ExifEditor{rootIb, false}, nil
}

// NewExifEditorEmpty create a new empty editor and sets the dirty flag
func NewExifEditorEmpty(dirty bool) (*ExifEditor, error) {
	ret := ExifEditor{}
	err := ret.Clear(dirty)
	return &ret, err
}

// Clear this editor and sets the dirty flag
func (ee *ExifEditor) Clear(dirty bool) error {
	im := exifcommon.NewIfdMapping()

	if err := exifcommon.LoadStandardIfds(im); err != nil {
		return err
	}

	ti := exif.NewTagIndex()

	ee.rootIb = exif.NewIfdBuilder(im, ti,
		exifcommon.IfdStandardIfdIdentity,
		exifcommon.EncodeDefaultByteOrder)
	ee.dirty = dirty
	return nil

}

// IsDirty if this editor has made any edits
func (ee ExifEditor) IsDirty() bool {
	return ee.dirty
}

// IsEmpty returns true if the IfdBuilder contains no data
func (ee ExifEditor) IsEmpty() bool {
	next, _ := ee.rootIb.NextIb()
	return len(ee.rootIb.Tags()) == 0 && next == nil
}

// IfdBuilder returns the underlying IfdBuilder. If it was changed it will also set the IFD Software tag.
func (ee *ExifEditor) IfdBuilder() (*exif.IfdBuilder, bool) {
	changed := ee.dirty
	_ = ee.setSoftware()
	return ee.rootIb, changed
}

// SetDirty force this editor to be marked as dirty
func (ee *ExifEditor) SetDirty() {
	ee.dirty = true
}

// SetDate sets the specified dateTag to time. Both the DateTime and corresponding offest
// will be set
func (ee *ExifEditor) SetDate(dateTag ExifDate, time time.Time) error {
	var err error
	offset := exifOffsetString(time)
	switch dateTag {
	case OriginalDate:
		if err = ee.SetIfdExifTag(ExifIFD_DateTimeOriginal, time); err != nil {
			return err
		}
		if err = ee.SetIfdExifTag(ExifIFD_OffsetTimeOriginal, offset); err != nil {
			return err
		}
	case ModifyDate:
		if err = ee.SetIfdRootTag(IFD_ModifyDate, time); err != nil {
			return err
		}
		if err = ee.SetIfdExifTag(ExifIFD_OffsetTime, offset); err != nil {
			return err
		}
	case DigitizedDate:
		if err = ee.SetIfdExifTag(ExifIFD_CreateDate, time); err != nil {
			return err
		}
		if err = ee.SetIfdExifTag(ExifIFD_OffsetTimeDigitized, offset); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown date to set: %v", dateTag)
	}
	return nil
}

// SetImageDescription sets IFD_ImageDescription to description
func (ee *ExifEditor) SetImageDescription(description string) error {
	return ee.SetIfdRootTag(IFD_ImageDescription, description)
}

// SetIfdExifTag sets the ExifIFD tag id to value
func (ee *ExifEditor) SetIfdExifTag(id ExifTag, value interface{}) error {
	exifIb, err := exif.GetOrCreateIbFromRootIb(ee.rootIb, IFDPaths[ExifIFD])
	if err != nil {
		return err
	}
	if err = exifIb.SetStandard(uint16(id), toGoExifValue(value)); err != nil {
		return err
	}
	ee.dirty = true
	return nil
}

// SetIfdRootTag set RootIFD tag id to value
func (ee *ExifEditor) SetIfdRootTag(id ExifTag, value interface{}) error {
	if err := ee.rootIb.SetStandard(uint16(id), toGoExifValue(value)); err != nil {
		return err
	}
	ee.dirty = true
	return nil
}

func (ee *ExifEditor) setSoftware() error {
	if !ee.dirty {
		return nil
	}
	if err := ee.SetIfdRootTag(IFD_Software, exifEditorSoftware); err != nil {
		return err
	}
	ee.dirty = false
	return nil
}

// SetUserComment sets ExifIFD_UserComment to comment using exifundefined.Tag9286UserComment. The
// comment will be Unicode encoded
func (ee *ExifEditor) SetUserComment(comment string) error {
	uc := exifundefined.Tag9286UserComment{
		EncodingType:  exifundefined.TagUndefinedType_9286_UserComment_Encoding_UNICODE,
		EncodingBytes: []byte(comment),
	}
	return ee.SetIfdExifTag(ExifIFD_UserComment, uc)
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
	case LensInfo:
		return t.toRational()
	default:
		return t
	}
}
