package metadata

import (
	"errors"
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	exifundefined "github.com/dsoprea/go-exif/v3/undefined"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"strings"
	"time"
)

var ErrExifNoData = errors.New("No Exif data")
var ErrExifTagNotFound = errors.New("Ifd tag not found")
var ErrExifValueNotFound = errors.New("Ifd value not found")
var ErrExifParseTag = errors.New("Exif tag could not be parsed")
var ErrExifUndefinedType = errors.New("Tag type undefined")

const ExifDateTime = "2006:01:02 15:04:05"
const ExifDateTimeOffset = "2006:01:02 15:04:05 -07:00"

func exifDateTime(dt string, offset string) (time.Time, error) {
	if offset == "" {
		return time.Parse(ExifDateTime, dt)
	}
	return time.Parse(ExifDateTimeOffset, dt+" "+offset)
}

type ExifDate int

const (
	OriginalDate ExifDate = iota
	ModifyDate
	DigitizedDate
)

func ExifTagName(index ExifIndex, tag ExifTag) string {
	tagDesc, found := ExifTagDescriptions[ExifIndexTag{Index: index, Tag: tag}]
	if found {
		return tagDesc.Name
	}
	return fmt.Sprintf("Unknown Ifd Tag: %v", tag)
}

func ExifValueIsAllowed(index ExifIndex, tag ExifTag, value interface{}) bool {
	tagDesc, found := ExifTagDescriptions[ExifIndexTag{index, tag}]
	if !found {
		return false
	} else if tagDesc.Values == nil {
		return true
	}
	switch tagDesc.Type {
	case ExifString:
		v, ok := value.(string)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[string]string)
		_, found = vals[v]
		return found

	case ExifUint8:
		v, ok := value.(uint8)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[uint8]string)
		_, found = vals[v]
		return found
	case ExifUint16:
		v, ok := value.(uint16)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[uint16]string)
		_, found = vals[v]
		return found
	case ExifUint32:
		v, ok := value.(uint32)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[uint32]string)
		_, found = vals[v]
		return found
	case ExifInt16:
		v, ok := value.(int16)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[int16]string)
		_, found = vals[v]
		return found

	case ExifInt32:
		v, ok := value.(int32)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[int32]string)
		_, found = vals[v]
		return found

	case ExifRational:
		v, ok := value.(Rat)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[Rat]string)
		_, found = vals[v]
		return found

	case ExifUrational:
		v, ok := value.(URat)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[URat]string)
		_, found = vals[v]
		return found

	case ExifFloat:
		v, ok := value.(float32)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[float32]string)
		_, found = vals[v]
		return found

	case ExifDouble:
		v, ok := value.(float64)
		if !ok {
			return false
		}
		vals := tagDesc.Values.(map[float64]string)
		_, found = vals[v]
		return found
	case ExifUndef:
		_, ok := value.([]byte)
		return ok
	default:
		return false
	}
	//return false
}

func ExifValueString(index ExifIndex, tag ExifTag, value interface{}) string {
	if v, err := ExifValueStringErr(index, tag, value); err == nil {
		return v
	}
	return "undefined"
}

func ExifValueStringErr(index ExifIndex, tag ExifTag, value interface{}) (string, error) {
	tagDesc, found := ExifTagDescriptions[ExifIndexTag{index, tag}]
	if !found {
		return "", ErrExifTagNotFound
	}
	if tagDesc.Values == nil {
		return "", ErrExifValueNotFound
	}
	retErr := func(ok bool) error {
		if ok {
			return nil
		}
		return ErrExifValueNotFound
	}
	ret := ""
	switch tagDesc.Type {
	case ExifString:
		v, ok := value.(string)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[string]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifUint8:
		v, ok := value.(uint8)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[uint8]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifUint16:
		v, ok := value.(uint16)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[uint16]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifUint32:
		v, ok := value.(uint32)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[uint32]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifInt16:
		v, ok := value.(int16)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[int16]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifInt32:
		v, ok := value.(int32)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[int32]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifRational:
		v, ok := value.(Rat)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[Rat]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifUrational:
		v, ok := value.(URat)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[URat]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifFloat:
		v, ok := value.(float32)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[float32]string)
		ret, found = vals[v]
		return ret, retErr(found)
	case ExifDouble:
		v, ok := value.(float64)
		if !ok {
			return "", ErrExifUndefinedType
		}
		vals := tagDesc.Values.(map[float64]string)
		ret, found = vals[v]
		return ret, retErr(found)
	default:
		return "", ErrExifUndefinedType
	}

}

type ExifData struct {
	rawExif *exif.IfdIndex
}

func NewExifData(segments *jpegstructure.SegmentList) (*ExifData, error) {
	var rawExif []byte
	var ifdMapping *exifcommon.IfdMapping
	var err error
	if _, rawExif, err = segments.Exif(); err != nil {
		return &ExifData{}, ErrExifNoData
	}
	if ifdMapping, err = exifcommon.NewIfdMappingWithStandard(); err != nil {
		return nil, err
	}
	ti := exif.NewTagIndex()
	if err = exif.LoadStandardTags(ti); err != nil {
		return nil, err
	}

	_, index, err := exif.Collect(ifdMapping, ti, rawExif)
	if err != nil {
		return &ExifData{}, err
	}
	return &ExifData{&index}, nil
}

func (ed *ExifData) IsEmpty() bool {
	return ed.rawExif == nil
}

func (ed *ExifData) GpsInfo() (*exif.GpsInfo, error) {
	gpsIfd := ed.Ifd(GpsIFD)
	if gpsIfd == nil {
		return nil, ErrExifTagNotFound
	}
	return gpsIfd.GpsInfo()
}

func (ed *ExifData) GetIfdImageDescription() string {
	ret := ""
	if err := ed.ScanIfdRoot(IFD_ImageDescription, &ret); err != nil {
		return ""
	}
	return ret
}

func (ed *ExifData) GetIfdUserComment() string {
	ret := exifundefined.Tag9286UserComment{}
	if err := ed.ScanIfdExif(ExifIFD_UserComment, &ret); err != nil {
		return ""
	}
	return string(ret.EncodingBytes)
}

func (ed *ExifData) HasIfd(index ExifIndex) bool {
	if ed.IsEmpty() {
		return false
	}
	_, found := ed.rawExif.Lookup[IFDPaths[index]]
	return found
}

func (ed *ExifData) Ifd(index ExifIndex) *exif.Ifd {
	if ed.IsEmpty() {
		return nil
	}
	return ed.rawExif.Lookup[IFDPaths[index]]
}

func scanExifValue(tagDesc ExifTagDesc, source interface{}, dest interface{}) error {
	switch dtype := dest.(type) {
	case *string:
		v, ok := source.(string)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *uint8:
		v, ok := source.([]uint8)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v[0]
	case *uint16:
		v, ok := source.([]uint16)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v[0]
	case *uint32:
		v, ok := source.([]uint32)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v[0]
	case *int16:
		v, ok := source.([]int16)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v[0]
	case *int32:
		v, ok := source.([]int32)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v[0]
	case *float32:
		v, ok := source.([]float32)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v[0]
	case *float64:
		v, ok := source.([]float64)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v[0]
	case *URat:
		v, ok := source.([]exifcommon.Rational)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = NewURatFromRational(v[0])
	case *Rat:
		v, ok := source.([]exifcommon.SignedRational)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = NewRatFromSignedRational(v[0])
	case *exifundefined.Tag9286UserComment:
		v, ok := source.(exifundefined.Tag9286UserComment)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *[]byte:
		v, ok := source.([]byte)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	default:
		return ErrExifUndefinedType
	}
	return nil
}

func scanMultipleExifValue(tagDesc ExifTagDesc, source interface{}, dest interface{}) error {
	switch dtype := dest.(type) {
	case *[]uint8:
		v, ok := source.([]uint8)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *[]uint16:
		v, ok := source.([]uint16)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *[]uint32:
		v, ok := source.([]uint32)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *[]int16:
		v, ok := source.([]int16)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *[]int32:
		v, ok := source.([]int32)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *[]float32:
		v, ok := source.([]float32)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *[]float64:
		v, ok := source.([]float64)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = v
	case *[]URat:
		vals, ok := source.([]exifcommon.Rational)
		if !ok {
			return ErrExifParseTag
		}
		for _, v := range vals {
			*dtype = append(*dtype, NewURatFromRational(v))
		}

	case *[]Rat:
		vals, ok := source.([]exifcommon.SignedRational)
		if !ok {
			return ErrExifParseTag
		}
		for _, v := range vals {
			*dtype = append(*dtype, NewRatFromSignedRational(v))
		}
	case *LensInfo:
		vals, ok := source.([]exifcommon.Rational)
		if !ok || len(vals) < 4 {
			return ErrExifParseTag
		}
		ret, e := NewLensInfoFromRational(vals)
		if e != nil {
			return e
		}
		*dtype = ret
	default:
		return ErrExifUndefinedType
	}
	return nil
}

func (ed *ExifData) Scan(ifdIndex ExifIndex, tagId ExifTag, dest interface{}) error {
	if ed.IsEmpty() {
		return ErrExifNoData
	}
	tagDesc, found := ExifTagDescriptions[ExifIndexTag{ifdIndex, tagId}]
	if !found {
		return ErrExifTagNotFound
	}
	ifd, found := ed.rawExif.Lookup[IFDPaths[ifdIndex]]
	if !found {
		return ErrExifTagNotFound
	}
	entries, err := ifd.FindTagWithId(uint16(tagId))
	if err != nil {
		return ErrExifTagNotFound
	}
	if len(entries) < 1 {
		return ErrExifValueNotFound
	}

	value, err := entries[0].Value()
	if err != nil {
		return ErrExifValueNotFound
	}

	if tagDesc.Count == 1 || tagDesc.Type == ExifString || tagDesc.Type == ExifUndef {
		return scanExifValue(tagDesc, value, dest)
	}

	return scanMultipleExifValue(tagDesc, value, dest)
}

/*
func (ed *ExifData) Scan2(ifdIndex ExifIndex, tagId ExifTag, dest interface{}) error {
	if ed.IsEmpty() {
		return ErrExifNoData
	}
	ifd, found := ed.rawExif.Lookup[IFDPaths[ifdIndex]]
	if !found {
		return ErrExifTagNotFound
	}

	entries, err := ifd.FindTagWithId(uint16(tagId))
	if err != nil {
		return ErrExifTagNotFound
	}
	if len(entries) < 1 {
		return fmt.Errorf("No entry data for: %s", ExifTagName(ifdIndex, tagId))
	}

	entry := entries[0]

	value, err := entry.Value()

	if err != nil {
		return err
	}

	wrongTagType := false

	switch dtype := dest.(type) {

	case *string:
		//Todo: do a proper convert instead of using format
		*dtype, _ = entry.Format()
	case *float32:
		if entry.TagType() == exifcommon.TypeFloat {
			v := value.([]float32)
			*dtype = float32(v[0])
		} else {
			wrongTagType = true
		}
	case *float64:
		if entry.TagType() == exifcommon.TypeDouble {
			v := value.([]float64)
			*dtype = float64(v[0])
		} else if entry.TagType() == exifcommon.TypeFloat {
			v := value.([]float32)
			*dtype = float64(v[0])
		} else if entry.TagType() == exifcommon.TypeRational {
			v := value.([]exifcommon.Rational)
			*dtype = float64(v[0].Numerator) / float64(v[0].Denominator)
		} else if entry.TagType() == exifcommon.TypeSignedRational {
			v := value.([]exifcommon.SignedRational)
			*dtype = float64(v[0].Numerator) / float64(v[0].Denominator)
		} else {
			wrongTagType = true
		}
	case *int64:
		if entry.TagType() == exifcommon.TypeShort {
			v := value.([]uint16)
			*dtype = int64(v[0])
		} else if entry.TagType() == exifcommon.TypeLong {
			v := value.([]uint32)
			*dtype = int64(v[0])
		} else if entry.TagType() == exifcommon.TypeSignedLong {
			v := value.([]int32)
			*dtype = int64(v[0])
		} else {
			wrongTagType = true
		}
	case *uint64:
		if entry.TagType() == exifcommon.TypeShort {
			v := value.([]uint16)
			*dtype = uint64(v[0])
		} else if entry.TagType() == exifcommon.TypeLong {
			v := value.([]uint32)
			*dtype = uint64(v[0])
		} else {
			wrongTagType = true
		}
	case *int32:
		if entry.TagType() == exifcommon.TypeShort {
			v := value.([]uint16)
			*dtype = int32(v[0])
		} else if entry.TagType() == exifcommon.TypeSignedLong {
			v := value.([]int32)
			*dtype = v[0]
		} else {
			wrongTagType = true
		}
	case *uint32:
		if entry.TagType() == exifcommon.TypeShort {
			v := value.([]uint16)
			*dtype = uint32(v[0])
		} else if entry.TagType() == exifcommon.TypeLong {
			v := value.([]uint32)
			*dtype = v[0]
		} else {
			wrongTagType = true
		}
	case *uint16:
		if entry.TagType() == exifcommon.TypeShort {
			v := value.([]uint16)
			*dtype = v[0]
			return nil
		}
		wrongTagType = true
	case *URat:
		if entry.TagType() == exifcommon.TypeRational {
			v := value.([]exifcommon.Rational)
			*dtype = NewURatFromRational(v[0])
		} else {
			wrongTagType = true
		}
	case *Rat:
		if entry.TagType() == exifcommon.TypeSignedRational {
			v := value.([]exifcommon.SignedRational)
			*dtype = NewRatFromSignedRational(v[0])
		} else {
			wrongTagType = true
		}
	case *LensInfo:
		if entry.TagType() == exifcommon.TypeRational {
			v := value.([]exifcommon.Rational)
			if len(v) != 4 {
				return fmt.Errorf("Expected 4 values for LensInfo got %v", len(v))
			}
			ret, e := NewLensInfoFromRational(v)
			if e != nil {
				return e
			}
			*dtype = ret
		} else {
			wrongTagType = true
		}
	case *exifundefined.Tag9286UserComment:
		if entry.TagType() == exifcommon.TypeUndefined {
			v, ok := value.(exifundefined.Tag9286UserComment)
			if !ok {
				return fmt.Errorf("Cannot recognise User Comment Type")
			}
			*dtype = v
		} else {
			wrongTagType = true
		}
	default:
		return fmt.Errorf("Cannot handle destination type %T", dest)
	}

	if wrongTagType {
		n := exifcommon.TypeNames[entry.TagType()]
		return fmt.Errorf("Wrong TagType: %s for destination", n)
	}
	return nil
}
*/

func (ed *ExifData) ScanExifDate(dateTag ExifDate, dest *time.Time) error {
	if ed.IsEmpty() {
		return ErrExifNoData
	}
	var t, o string
	var err error
	switch dateTag {
	case OriginalDate:
		if err = ed.ScanIfdExif(ExifIFD_DateTimeOriginal, &t); err != nil {
			return err
		}
		_ = ed.ScanIfdExif(ExifIFD_OffsetTimeOriginal, &o) //dont care about offset errors
		*dest, err = exifDateTime(t, o)
		if err != nil {
			return err
		}
	case ModifyDate:
		if err = ed.ScanIfdRoot(IFD_ModifyDate, &t); err != nil {
			return err
		}
		_ = ed.ScanIfdExif(ExifIFD_OffsetTime, &o) //dont care about offset errors
		*dest, err = exifDateTime(t, o)
		if err != nil {
			return err
		}
	case DigitizedDate:
		if err = ed.ScanIfdExif(ExifIFD_CreateDate, &t); err != nil {
			return err
		}
		_ = ed.ScanIfdExif(ExifIFD_OffsetTimeDigitized, &o) //dont care about offset errors
		*dest, err = exifDateTime(t, o)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown date to scan: %v", dateTag)
	}
	return nil
}

func (ed *ExifData) ScanIfdExif(tag ExifTag, dest interface{}) error {
	return ed.Scan(ExifIFD, tag, dest)
}

func (ed *ExifData) ScanIfdIop(tagId ExifTag, dest interface{}) error {
	return ed.Scan(InteropIFD, tagId, dest)
}

func (ed *ExifData) ScanIfdRoot(tagId ExifTag, dest interface{}) error {
	return ed.Scan(RootIFD, tagId, dest)
}

func (ed *ExifData) ScanIfdThumbnail(tagId ExifTag, dest interface{}) error {
	return ed.Scan(ThumbnailIFD, tagId, dest)
}

func (ed *ExifData) String() string {
	if ed.IsEmpty() {
		return "No Exif Defined/Empty"
	}
	sb := strings.Builder{}
	sb.WriteString("exifdata:{\n")
	for _, ifd := range ed.rawExif.Ifds {
		sb.WriteString("  " + ifd.IfdIdentity().String() + ":{\n")
		for _, e := range ifd.Entries() {
			var f string
			if e.TagType() == exifcommon.TypeByte || e.TagType() == exifcommon.TypeUndefined {
				f = "unprintable value"
			} else {
				f, _ = e.Format()
			}
			val := fmt.Sprintf("    %#04x Name:%s Type:%v Length:%v Value:%s\n", e.TagId(), e.TagName(), e.TagType().String(), e.UnitCount(), f)
			sb.WriteString(val)
		}
		sb.WriteString("  }\n")
	}
	sb.WriteString("}")
	return sb.String()
}
