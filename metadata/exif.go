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

var NoExifErr = errors.New("No Exif data")
var IfdTagNotFoundErr = errors.New("Ifd tag not found")
var IfdValueNotFoundErr = errors.New("Ifd value not found")
var IfdUndefinedTypeErr = errors.New("Tag type undefined")

const ExifDateTime = "2006:01:02 15:04:05"
const ExifDateTimeOffset = "2006:01:02 15:04:05 -07:00"

const (
	OriginalDate ExifDate = iota
	ModifyDate
	DigitizedDate
)

type IfdIndex int

const (
	IfdRoot IfdIndex = iota
	IfdExif
	IfdGpsInfo
	IfdIop
	IfdThumbnail
)

var IfdPaths = map[IfdIndex]string{
	IfdRoot:      "IFD",
	IfdExif:      "IFD/Exif",
	IfdGpsInfo:   "IFD/GPSInfo",
	IfdIop:       "IFD/Exif/Iop",
	IfdThumbnail: "IFD1",
}

var ifdValueMap = map[IfdIndex]map[uint16]interface{}{
	IfdRoot:      IFDValues,
	IfdExif:      ExifValues,
	IfdGpsInfo:   GPSInfoValues,
	IfdIop:       IopValues,
	IfdThumbnail: IFDValues,
}

func ExifValueIsAllowed(index IfdIndex, tagId uint16, value interface{}) bool {
	ifdValueMap, found := ifdValueMap[index]
	if !found {
		return false
	}
	tagValues, found := ifdValueMap[tagId]
	if !found {
		return false
	}
	switch dtype := value.(type) {
	case uint16:
		if m, ok := tagValues.(map[uint16]string); !ok {
			return false
		} else {
			_, f := m[dtype]
			return f
		}
	case int16:
		if m, ok := tagValues.(map[int16]string); !ok {
			return false
		} else {
			_, f := m[dtype]
			return f
		}
	case uint32:
		if m, ok := tagValues.(map[uint32]string); !ok {
			return false
		} else {
			_, f := m[dtype]
			return f
		}
	case int32:
		if m, ok := tagValues.(map[int32]string); !ok {
			return false
		} else {
			_, f := m[dtype]
			return f
		}
	case string:
		if m, ok := tagValues.(map[string]string); !ok {
			return false
		} else {
			_, f := m[dtype]
			return f
		}
	case Rat:
		if m, ok := tagValues.(map[Rat]string); !ok {
			return false
		} else {
			_, f := m[dtype]
			return f
		}
	case URat:
		if m, ok := tagValues.(map[URat]string); !ok {
			return false
		} else {
			_, f := m[dtype]
			return f
		}
	default:
		return false
	}
}

func ExifValueStringNoErr(index IfdIndex, tagId uint16, value interface{}) string {
	if v, err := ExifValueString(index, tagId, value); err == nil {
		return v
	} else {
		return "undefined"
	}
}

func ExifValueString(index IfdIndex, tagId uint16, value interface{}) (string, error) {
	ifdValueMap, found := ifdValueMap[index]
	if !found {
		return "", IfdTagNotFoundErr
	}
	tagValues, found := ifdValueMap[tagId]
	if !found {
		return "", IfdTagNotFoundErr
	}
	switch dtype := value.(type) {
	case uint16:
		if m, ok := tagValues.(map[uint16]string); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if v, f := m[dtype]; f {
				return v, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case int16:
		if m, ok := tagValues.(map[int16]string); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if v, f := m[dtype]; f {
				return v, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case uint32:
		if m, ok := tagValues.(map[uint32]string); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if v, f := m[dtype]; f {
				return v, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case int32:
		if m, ok := tagValues.(map[int32]string); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if v, f := m[dtype]; f {
				return v, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case string:
		if m, ok := tagValues.(map[string]string); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if v, f := m[dtype]; f {
				return v, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case Rat:
		if m, ok := tagValues.(map[Rat]string); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if v, f := m[dtype]; f {
				return v, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case URat:
		if m, ok := tagValues.(map[URat]string); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if v, f := m[dtype]; f {
				return v, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	default:
		return "", IfdUndefinedTypeErr
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
		return &ExifData{}, NoExifErr
	}
	if ifdMapping, err = exifcommon.NewIfdMappingWithStandard(); err != nil {
		return nil, err
	}
	ti := exif.NewTagIndex()
	if err = exif.LoadStandardTags(ti); err != nil {
		return nil, err
	}

	if _, index, err := exif.Collect(ifdMapping, ti, rawExif); err != nil {
		return &ExifData{}, err
	} else {
		return &ExifData{&index}, nil
	}
}

func (ed *ExifData) IsEmpty() bool {
	return ed.rawExif == nil
}

func (ed *ExifData) GpsInfo() (*exif.GpsInfo, error) {
	gpsIfd := ed.Ifd(IfdGpsInfo)
	if gpsIfd == nil {
		return nil, IfdTagNotFoundErr
	} else {
		return gpsIfd.GpsInfo()
	}
}

//Convinince method to retrive IFD_ImageDescription. The MetaDataEditor has
//a corresponding method to set the IFD_ImageDescription
func (ed *ExifData) GetIfdImageDescription() (string, error) {
	ret := ""
	err := ed.ScanIfdRoot(IFD_ImageDescription, &ret)
	return ret, err
}

//Convinience method to retrieve Exif_UserComment. As the UserComment is
//an undefined field this method will assume it has been set by the corresponding
//MetaDataEditor method
func (ed *ExifData) GetIfdUserComment() (string, error) {
	ret := exifundefined.Tag9286UserComment{}
	err := ed.ScanIfdExif(Exif_UserComment, &ret)
	return string(ret.EncodingBytes), err
}

func (ed *ExifData) HasIfd(index IfdIndex) bool {
	if ed.IsEmpty() {
		return false
	}
	_, found := ed.rawExif.Lookup[IfdPaths[index]]
	return found
}

func (ed *ExifData) Ifd(index IfdIndex) *exif.Ifd {
	if ed.IsEmpty() {
		return nil
	}
	return ed.rawExif.Lookup[IfdPaths[index]]
}

func (ed *ExifData) Scan(ifdIndex IfdIndex, tagId uint16, dest interface{}) error {
	if ed.IsEmpty() {
		return NoExifErr
	}

	ifd, found := ed.rawExif.Lookup[IfdPaths[ifdIndex]]
	if !found {
		return IfdTagNotFoundErr
	}

	entries, err := ifd.FindTagWithId(tagId)
	if err != nil {
		return IfdTagNotFoundErr
	}
	if len(entries) < 1 {
		return fmt.Errorf("No entry data for: %s", ExifTagName(ifdIndex, tagId))
	}

	entry := entries[0]

	value, err := entry.Value()

	if err != nil {
		return err
	}

	//Simplest case....dest is a string just use the entry.Format function
	wrongTagType := false

	switch dtype := dest.(type) {

	case *string:
		*dtype, err = entry.Format()
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
		} else {
			wrongTagType = true
		}
	case *URat:
		if entry.TagType() == exifcommon.TypeRational {
			v := value.([]exifcommon.Rational)
			*dtype = NewURatFromRational(v[0])
			//dtype.Numerator = v[0].Numerator
			//dtype.Denominator = v[0].Denominator
		} else {
			wrongTagType = true
		}
	case *Rat:
		if entry.TagType() == exifcommon.TypeSignedRational {
			v := value.([]exifcommon.SignedRational)
			*dtype = Rat{v[0].Numerator, v[0].Denominator}
			//dtype.Numerator = v[0].Numerator
			//dtype.Denominator = v[0].Denominator
		} else {
			wrongTagType = true
		}
	case *LensInfo:
		if entry.TagType() == exifcommon.TypeRational {
			v := value.([]exifcommon.Rational)
			if len(v) != 4 {
				return fmt.Errorf("Expected 4 values for LensInfo got %v", len(v))
			}
			if ret, e := NewLensInfoFromRational(v); e != nil {
				return e
			} else {
				*dtype = ret
			}
		} else {
			wrongTagType = true
		}
	case *exifundefined.Tag9286UserComment:
		if entry.TagType() == exifcommon.TypeUndefined {
			v, ok := value.(exifundefined.Tag9286UserComment)
			if !ok {
				fmt.Errorf("Cannot recognise User Comment Type")
			} else {
				*dtype = v
			}
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

func (ed *ExifData) ScanExifDate(dateTag ExifDate, dest *time.Time) error {
	if ed.IsEmpty() {
		return NoExifErr
	}
	var t, o string
	var err error
	switch dateTag {
	case OriginalDate:
		if err = ed.ScanIfdExif(Exif_DateTimeOriginal, &t); err != nil {
			return err
		}
		_ = ed.ScanIfdExif(Exif_OffsetTimeOriginal, &o) //dont care about offset errors
		*dest, err = ParseIfdDateTime(t, o)
		if err != nil {
			return err
		}
	case ModifyDate:
		if err = ed.ScanIfdRoot(IFD_DateTime, &t); err != nil {
			return err
		}
		_ = ed.ScanIfdExif(Exif_OffsetTime, &o) //dont care about offset errors
		*dest, err = ParseIfdDateTime(t, o)
		if err != nil {
			return err
		}
	case DigitizedDate:
		if err = ed.ScanIfdExif(Exif_DateTimeDigitized, &t); err != nil {
			return err
		}
		_ = ed.ScanIfdExif(Exif_OffsetTimeDigitized, &o) //dont care about offset errors
		*dest, err = ParseIfdDateTime(t, o)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown date to scan: %v", dateTag)
	}
	return nil
}

func (ed *ExifData) ScanIfdRoot(tagId uint16, dest interface{}) error {
	return ed.Scan(IfdRoot, tagId, dest)
}

func (ed *ExifData) ScanIfdExif(tagId uint16, dest interface{}) error {
	return ed.Scan(IfdExif, tagId, dest)
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

func TimeOffsetString(t time.Time) string {
	_, offset := t.Zone()
	sign := '+'
	if offset < 0 {
		sign = '-'
		offset = -offset
	}
	h := offset / 3600
	m := (offset % 3600) / 60
	return fmt.Sprintf("%c%02v:%02v", sign, h, m)
}

func ParseIfdDateTime(dt string, offset string) (time.Time, error) {
	if offset == "" {
		return time.Parse(ExifDateTime, dt)
	} else {
		return time.Parse(ExifDateTimeOffset, dt+" "+offset)
	}
}

func ExifTagName(index IfdIndex, fieldId uint16) string {
	var name string
	var found bool
	switch index {
	case IfdRoot, IfdThumbnail:
		name, found = IFDName[fieldId]
	case IfdExif:
		name, found = ExifName[fieldId]
	case IfdIop:
		name, found = IopName[fieldId]
	case IfdGpsInfo:
		name, found = GPSInfoName[fieldId]
	}
	if found {
		return name
	} else {
		return fmt.Sprintf("Unknown Ifd Tag: %v", fieldId)
	}
}

/*
func ExifTagName(ifd *exif.Ifd, fieldId uint16) string {
	var name string
	var found bool
	if ifd == nil {
		return fmt.Sprintf("Unknown Ifd Nil")
	}
	if ifd.IfdIdentity().String() == IFDPath || ifd.IfdIdentity().String() == ThumbNailPath {
		name, found = IFDName[fieldId]
	} else if ifd.IfdIdentity().String() == ExifPath {
		name, found = ExifName[fieldId]
	} else if ifd.IfdIdentity().String() == IopPath {
		name, found = IopName[fieldId]
	} else if ifd.IfdIdentity().String() == GPSInfoPath {
		name, found = GPSInfoName[fieldId]
	} else {
		return fmt.Sprintf("Unknown Ifd Path %s", ifd.IfdIdentity().String())
	}
	if found {
		return name
	} else {
		return fmt.Sprintf("Unknown Ifd Tag: %v", fieldId)
	}
}
*/

/*
func ScanIfdTag(ifd *exif.Ifd, tagId uint16, dest interface{}) error {
	entries, err := ifd.FindTagWithId(tagId)
	if err != nil {
		return IfdTagNotFoundErr
	}
	if len(entries) < 1 {
		return fmt.Errorf("No entry data for: %s", ExifTagName(ifd, tagId))
	}

	entry := entries[0]

	value, err := entry.Value()

	if err != nil {
		return err
	}
	wrongTagType := false

	switch dtype := dest.(type) {

	case *string:
		*dtype, err = entry.Format()
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
		} else {
			wrongTagType = true
		}
	case *URat:
		if entry.TagType() == exifcommon.TypeRational {
			v := value.([]exifcommon.Rational)
			*dtype = NewURatFromRational(v[0])
			//dtype.Numerator = v[0].Numerator
			//dtype.Denominator = v[0].Denominator
		} else {
			wrongTagType = true
		}
	case *Rat:
		if entry.TagType() == exifcommon.TypeSignedRational {
			v := value.([]exifcommon.SignedRational)
			*dtype = fromExifRat(v[0])
			//dtype.Numerator = v[0].Numerator
			//dtype.Denominator = v[0].Denominator
		} else {
			wrongTagType = true
		}
	case *LensInfo:
		if entry.TagType() == exifcommon.TypeRational {
			v := value.([]exifcommon.Rational)
			if len(v) != 4 {
				return fmt.Errorf("Expected 4 values for LensInfo got %v", len(v))
			}
			if ret, e := NewLensInfoFromRational(v); e != nil {
				return e
			} else {
				*dtype = ret
			}
		} else {
			wrongTagType = true
		}
	case *exifundefined.Tag9286UserComment:
		if entry.TagType() == exifcommon.TypeUndefined {
			v, ok := value.(exifundefined.Tag9286UserComment)
			if !ok {
				fmt.Errorf("Cannot recognise User Comment Type")
			} else {
				*dtype = v
			}
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
