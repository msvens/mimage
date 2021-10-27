package metadata

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/msvens/mimage/metadata/ifdexif"
	"github.com/msvens/mimage/metadata/ifdroot"
	"time"
)

const ExifDateTime = "2006:01:02 15:04:05"
const ExifDateTimeOffset = "2006:01:02 15:04:05 -07:00"

//Extracted from: https://exiftool.org/TagNames/EXIF.html
var colorSpace = map[uint16]string{
	0x1:    "sRGB",
	0x2:    "Adobe RGB",
	0xfffd: "Wide Gamut RGB",
	0xfffe: "ICC Profile",
	0xffff: "Uncalibrated",
}

func ColorSpaceName(cs uint16) string {
	if str, found := colorSpace[cs]; found {
		return str
	} else {
		return "Unknown ColorSpace"
	}
}

var exposureProgram = map[uint16]string{
	0: "Not Defined",
	1: "Manual",
	2: "Program AE",
	3: "Aperture-priority AE",
	4: "Shutter speed priority AE",
	5: "Creative (Slow speed)",
	6: "Action (High speed)",
	7: "Portrait",
	8: "Landscape",
	9: "Bulb",
}

func ExposureProgamName(ep uint16) string {
	if str, found := exposureProgram[ep]; found {
		return str
	} else {
		return "Unknown exposure program"
	}
}

var flashModes = map[uint16]string{
	0x0:  "No Flash",
	0x1:  "Fired",
	0x5:  "Fired, Return not detected",
	0x7:  "Fired, Return detected",
	0x8:  "On, Did not fire",
	0x9:  "On, Fired",
	0xd:  "On, Return not detected",
	0xf:  "On, Return detected",
	0x10: "Off, Did not fire",
	0x14: "Off, Did not fire, Return not detected",
	0x18: "Auto, Did not fire",
	0x19: "Auto, Fired",
	0x1d: "Auto, Fired, Return not detected",
	0x1f: "Auto, Fired, Return detected",
	0x20: "No flash function",
	0x30: "Off, No flash function",
	0x41: "Fired, Red-eye reduction",
	0x45: "Fired, Red-eye reduction, Return not detected",
	0x47: "Fired, Red-eye reduction, Return detected",
	0x49: "On, Red-eye reduction",
	0x4d: "On, Red-eye reduction, Return not detected",
	0x4f: "On, Red-eye reduction, Return detected",
	0x50: "Off, Red-eye reduction",
	0x58: "Auto, Did not fire, Red-eye reduction",
	0x59: "Auto, Fired, Red-eye reduction",
	0x5d: "Auto, Fired, Red-eye reduction, Return not detected",
	0x5f: "Auto, Fired, Red-eye reduction, Return detected",
}

func FlashModeName(flashMode uint16) string {
	if str, found := flashModes[flashMode]; found {
		return str
	} else {
		return "Unknown flash mode"
	}
}

func ParseIfdDateTime(dt string, offset string) (time.Time, error) {
	if offset == "" {
		return time.Parse(ExifDateTime, dt)
	} else {
		return time.Parse(ExifDateTimeOffset, dt+" "+offset)
	}
}

func getIfdTagName(ifd *exif.Ifd, fieldId uint16) string {
	if ifd.IfdIdentity().UnindexedString() == "IFD" {
		return ifdroot.TagName(fieldId)
	} else if ifd.IfdIdentity().UnindexedString() == "IFD/Exif" {
		return ifdexif.TagName(fieldId)
	} else {
		return "Unknown ifd Identity"
	}
}

func scanIfdTag(ifd *exif.Ifd, tagId uint16, dest interface{}) error {
	entries, err := ifd.FindTagWithId(tagId)
	if err != nil {
		return IfdTagNotFoundErr
	}
	if len(entries) < 1 {
		return fmt.Errorf("No entry data for: %s", getIfdTagName(ifd, tagId))
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
	case *Rat:
		if entry.TagType() == exifcommon.TypeSignedRational {
			v := value.([]exifcommon.SignedRational)
			dtype.Numerator = int64(v[0].Numerator)
			dtype.Denominator = int64(v[0].Denominator)
		} else if entry.TagType() == exifcommon.TypeRational {
			v := value.([]exifcommon.Rational)
			dtype.Numerator = int64(v[0].Numerator)
			dtype.Denominator = int64(v[0].Denominator)
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
