package metadata

import "strings"

//Detailed information about various Exif/Ifd types and their values
//Extracted from: https://exiftool.org/TagNames/EXIF.html

var colorSpace = map[uint16]string{
	0x1:    "sRGB",
	0x2:    "Adobe RGB",
	0xfffd: "Wide Gamut RGB",
	0xfffe: "ICC Profile",
	0xffff: "Uncalibrated",
}

func ColorSpace(cs uint16) string {
	if str, found := colorSpace[cs]; found {
		return str
	} else {
		return "Unknown ColorSpace"
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

func FlashMode(mode uint16) string {
	if str, found := flashModes[mode]; found {
		return str
	} else {
		return "Unknown flash mode"
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

func ExposureProgram(ep uint16) string {
	if str, found := exposureProgram[ep]; found {
		return str
	} else {
		return "Unknown exposure program"
	}
}

const (
	IfdByte      uint16 = 1
	IfdAscii     uint16 = 2
	IfdShort     uint16 = 3
	IfdLong      uint16 = 4
	IfdRational  uint16 = 5
	IfdUndefined uint16 = 7
	IfdSLong     uint16 = 9
	IfdSRational uint16 = 10
)

var IfdTypeId = map[string]uint16{
	"BYTE":      IfdByte,
	"ASCII":     IfdAscii,
	"SHORT":     IfdShort,
	"LONG":      IfdLong,
	"RATIONAL":  IfdRational,
	"UNDEFINED": IfdUndefined,
	"SLONG":     IfdSLong,
	"SRATIONAL": IfdSRational,
}

var IfdTypeName = map[uint16]string{
	IfdByte:      "BYTE",
	IfdAscii:     "ASCII",
	IfdShort:     "SHORT",
	IfdLong:      "LONG",
	IfdRational:  "RATIONAL",
	IfdUndefined: "UNDEFINED",
	IfdSLong:     "SLONG",
	IfdSRational: "SRATIONAL",
}

func IfdType(str string) uint16 {
	s := strings.ToUpper(strings.TrimSpace(str))
	if id, found := IfdTypeId[s]; found {
		return id
	} else {
		return IfdUndefined
	}
}
