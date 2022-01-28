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

func exifDateTime(dt string, offset string) (time.Time, error) {
	if offset == "" {
		return time.Parse(ExifDateTime, dt)
	} else {
		return time.Parse(ExifDateTimeOffset, dt+" "+offset)
	}
}

const (
	OriginalDate ExifDate = iota
	ModifyDate
	DigitizedDate
)

/*
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
*/

/*
var ifdValueMap = map[IfdIndex]map[ExifTag]interface{}{
	IfdRoot:      IFDValues,
	IfdExif:      ExifValues,
	IfdGpsInfo:   GPSInfoValues,
	IfdIop:       IopValues,
	IfdThumbnail: IFDValues,
}
*/

func ExifTagName(index ExifIndex, tag ExifTag) string {
	tagDesc, found := ExifTagDescriptions[ExifIndexTag{Index: index, Tag: tag}]
	if found {
		return tagDesc.Name
	} else {
		return fmt.Sprintf("Unknown Ifd Tag: %v", tag)
	}
}

func ExifValueIsAllowed(index ExifIndex, tag ExifTag, value interface{}) bool {
	tagDesc, found := ExifTagDescriptions[ExifIndexTag{index, tag}]
	if !found {
		return false
	}
	switch tagDesc.Type {
	case ExifString:
		if v, ok := value.(string); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[string]string)
			_, found = vals[v]
			return found
		}
	case ExifUint8:
		if v, ok := value.(uint8); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[uint8]string)
			_, found = vals[v]
			return found
		}
	case ExifUint16:
		if v, ok := value.(uint16); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[uint16]string)
			_, found = vals[v]
			return found
		}
	case ExifUint32:
		if v, ok := value.(uint32); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[uint32]string)
			_, found = vals[v]
			return found
		}
	case ExifInt16:
		if v, ok := value.(int16); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[int16]string)
			_, found = vals[v]
			return found
		}
	case ExifInt32:
		if v, ok := value.(int32); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[int32]string)
			_, found = vals[v]
			return found
		}
	case ExifRational:
		if v, ok := value.(Rat); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[Rat]string)
			_, found = vals[v]
			return found
		}
	case ExifUrational:
		if v, ok := value.(URat); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[URat]string)
			_, found = vals[v]
			return found
		}
	case ExifFloat:
		if v, ok := value.(float32); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[float32]string)
			_, found = vals[v]
			return found
		}
	case ExifDouble:
		if v, ok := value.(float64); !ok {
			return false
		} else {
			if tagDesc.Values == nil {
				return true
			}
			vals := tagDesc.Values.(map[float64]string)
			_, found = vals[v]
			return found
		}
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
	} else {
		return "undefined"
	}
}

func ExifValueStringErr(index ExifIndex, tag ExifTag, value interface{}) (string, error) {
	tagDesc, found := ExifTagDescriptions[ExifIndexTag{index, tag}]
	if !found {
		return "", IfdTagNotFoundErr
	}
	var ret = ""
	switch tagDesc.Type {
	case ExifString:
		if v, ok := value.(string); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[string]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifUint8:
		if v, ok := value.(uint8); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[uint8]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifUint16:
		if v, ok := value.(uint16); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[uint16]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifUint32:
		if v, ok := value.(uint32); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[uint32]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifInt16:
		if v, ok := value.(int16); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[int16]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifInt32:
		if v, ok := value.(int32); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[int32]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifRational:
		if v, ok := value.(Rat); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[Rat]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifUrational:
		if v, ok := value.(URat); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[URat]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifFloat:
		if v, ok := value.(float32); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[float32]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifDouble:
		if v, ok := value.(float64); !ok {
			return "", IfdUndefinedTypeErr
		} else {
			if tagDesc.Values == nil {
				return "", IfdValueNotFoundErr
			}
			vals := tagDesc.Values.(map[float64]string)
			if ret, found = vals[v]; found {
				return ret, nil
			} else {
				return "", IfdValueNotFoundErr
			}
		}
	case ExifUndef:
		if _, ok := value.([]byte); ok {
			return "", IfdValueNotFoundErr
		} else {
			return "", IfdUndefinedTypeErr
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
	gpsIfd := ed.Ifd(GpsIFD)
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
	err := ed.ScanIfdExif(ExifIFD_UserComment, &ret)
	return string(ret.EncodingBytes), err
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

func (ed *ExifData) Scan(ifdIndex ExifIndex, tagId ExifTag, dest interface{}) error {
	if ed.IsEmpty() {
		return NoExifErr
	}

	ifd, found := ed.rawExif.Lookup[IFDPaths[ifdIndex]]
	if !found {
		return IfdTagNotFoundErr
	}

	entries, err := ifd.FindTagWithId(uint16(tagId))
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
