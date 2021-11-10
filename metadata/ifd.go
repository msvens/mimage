package metadata

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"strings"
	"time"
)

const ExifDateTime = "2006:01:02 15:04:05"
const ExifDateTimeOffset = "2006:01:02 15:04:05 -07:00"

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

func ExifTagName(ifd *exif.Ifd, fieldId uint16) string {
	var name string
	var found bool
	if ifd.IfdIdentity().Name() == "IFD" {
		name, found = IFDName[fieldId]
	} else if ifd.IfdIdentity().Name() == "Exif" {
		name, found = ExifName[fieldId]
	}
	if found {
		return name
	} else {
		return "Unknown Ifd Identity"
	}
}

func PrintExif(ifdIndex *exif.IfdIndex) string {
	sb := strings.Builder{}
	sb.WriteString("exifdata:{\n")
	for _, ifd := range ifdIndex.Ifds {
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

func ScanIfdTag(ifd *exif.Ifd, tagId uint16, dest interface{}) error {
	entries, err := ifd.FindTagWithId(tagId)
	if err != nil {
		return IfdTagNotFoundErr
	}
	if len(entries) < 1 {
		return fmt.Errorf("No entry data for: %s", ExifTagName(ifd, tagId))
	}

	entry := entries[0]

	if entry.TagType() == exifcommon.TypeUndefined {
		return IfdUndefinedTypeErr
	}

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
			dtype.Numerator = v[0].Numerator
			dtype.Denominator = v[0].Denominator
		} else {
			wrongTagType = true
		}
	case *Rat:
		if entry.TagType() == exifcommon.TypeSignedRational {
			v := value.([]exifcommon.SignedRational)
			dtype.Numerator = v[0].Numerator
			dtype.Denominator = v[0].Denominator
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

//Writes exif tags from this MetaDataSummary. Will skip GPSInfo for now
func WriteSummary(md *MetaDataSummary, dest []byte) ([]byte, error) {
	return nil, nil
}
