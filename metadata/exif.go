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

// Exif errors
var (
	ErrExifNoData        = errors.New("No Exif data")
	ErrExifTagNotFound   = errors.New("Ifd tag not found")
	ErrExifValueNotFound = errors.New("Ifd value not found")
	ErrExifParseTag      = errors.New("Exif tag could not be parsed")
	ErrExifUndefinedType = errors.New("Tag type undefined")
)

// Exif Time formats
const (
	ExifDateTime       = "2006:01:02 15:04:05"
	ExifDateTimeOffset = "2006:01:02 15:04:05 -07:00"
)

func exifDateTime(dt string, offset string) (time.Time, error) {
	if offset == "" {
		return time.Parse(ExifDateTime, dt)
	}
	return time.Parse(ExifDateTimeOffset, dt+" "+offset)
}

// ExifDate type
type ExifDate int

// ExifDates found in the specification
const (
	OriginalDate ExifDate = iota
	ModifyDate
	DigitizedDate
)

// ExifTagName returns the common name of specific ExifTag
func ExifTagName(index ExifIndex, tag ExifTag) string {
	tagDesc, found := ExifTagDescriptions[ExifIndexTag{Index: index, Tag: tag}]
	if found {
		return tagDesc.Name
	}
	return fmt.Sprintf("Unknown Ifd Tag: %v", tag)
}

// ExifValueIsAllowed checks if an ExifTag value is ok to use. Most tags
// dont have any specific values defined. see (https://exiftool.org/TagNames/EXIF.html)
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

// ExifValueString returns the common name of specific ExifTag value. E.g. "No Flash" for value
// 0 for the FlashTag. If no mapping is found returns "undefined"
func ExifValueString(index ExifIndex, tag ExifTag, value interface{}) string {
	if v, err := ExifValueStringErr(index, tag, value); err == nil {
		return v
	}
	return "undefined"
}

// ExifValueStringErr returns the common name of the ExifTag value or an error, either
// ErrExifValueNotFound or ErrExifUndefinedType
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

// ExifData holds the underlying IfdIndex
type ExifData struct {
	rawExif *exif.IfdIndex
}

// NewExifData creates an ExifData from a jpeg segment list. Returns
// ErrExifNoData if the segment list did not contain any exif data
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

// IsEmpty returns true if the underlying IfdIndex is nil
func (ed *ExifData) IsEmpty() bool {
	return ed.rawExif == nil
}

// GpsInfo returns this Exifs gpsinfo or an error
func (ed *ExifData) GpsInfo() (*exif.GpsInfo, error) {
	gpsIfd := ed.Ifd(GpsIFD)
	if gpsIfd == nil {
		return nil, ErrExifTagNotFound
	}
	return gpsIfd.GpsInfo()
}

// GetImageDescription returns IFD_ImaageDescription or "" if the tag was
// not set
func (ed *ExifData) GetImageDescription() string {
	ret := ""
	if err := ed.ScanIfdRoot(IFD_ImageDescription, &ret); err != nil {
		return ""
	}
	return ret
}

// GetUserComment returns ExifIFD_UserComment as a string or "" if there was an error
// getting the comment. The comment is assumed to be encoded using exifundefined.Tag9286UserComment
func (ed *ExifData) GetUserComment() string {
	ret := exifundefined.Tag9286UserComment{}
	if err := ed.ScanIfdExif(ExifIFD_UserComment, &ret); err != nil {
		return ""
	}
	return string(ret.EncodingBytes)
}

// HasIfd checks if the specified index exists in this Ifd
func (ed *ExifData) HasIfd(index ExifIndex) bool {
	if ed.IsEmpty() {
		return false
	}
	_, found := ed.rawExif.Lookup[IFDPaths[index]]
	return found
}

// Ifd retrieves the given index or nil if it does not exist
func (ed *ExifData) Ifd(index ExifIndex) *exif.Ifd {
	if ed.IsEmpty() {
		return nil
	}
	return ed.rawExif.Lookup[IFDPaths[index]]
}

func scanExifValue(_ ExifTagDesc, source interface{}, dest interface{}) error {
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
		*dtype = newURatFromRational(v[0])
	case *Rat:
		v, ok := source.([]exifcommon.SignedRational)
		if !ok {
			return ErrExifParseTag
		}
		*dtype = newRatFromSignedRational(v[0])
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

func scanMultipleExifValue(_ ExifTagDesc, source interface{}, dest interface{}) error {
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
			*dtype = append(*dtype, newURatFromRational(v))
		}

	case *[]Rat:
		vals, ok := source.([]exifcommon.SignedRational)
		if !ok {
			return ErrExifParseTag
		}
		for _, v := range vals {
			*dtype = append(*dtype, newRatFromSignedRational(v))
		}
	case *LensInfo:
		vals, ok := source.([]exifcommon.Rational)
		if !ok || len(vals) < 4 {
			return ErrExifParseTag
		}
		ret, e := newLensInfoFromRational(vals)
		if e != nil {
			return e
		}
		*dtype = ret
	default:
		return ErrExifUndefinedType
	}
	return nil
}

// Scan reads the specific tag from index into dest
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

// ScanExifDate reads the given dateTag into dest
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

// ScanIfdExif scans tag from ExifIFD into dest
func (ed *ExifData) ScanIfdExif(tag ExifTag, dest interface{}) error {
	return ed.Scan(ExifIFD, tag, dest)
}

// ScanIfdIop scans tag from InteropIFD into dest
func (ed *ExifData) ScanIfdIop(tagId ExifTag, dest interface{}) error {
	return ed.Scan(InteropIFD, tagId, dest)
}

// ScanIfdRoot scans tag from RootIFD into dest
func (ed *ExifData) ScanIfdRoot(tagId ExifTag, dest interface{}) error {
	return ed.Scan(RootIFD, tagId, dest)
}

// ScanIfdThumbnail scans tag from ThumbnailIFD into dest
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
