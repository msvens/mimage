package metadata

import (
	"fmt"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

type Rat struct {
	Numerator   int32
	Denominator int32
}

func NewRatFromSignedRational(v exifcommon.SignedRational) Rat {
	return Rat{v.Numerator, v.Denominator}
}

func (r Rat) Float64() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func (r Rat) Float32() float32 {
	return float32(r.Numerator) / float32(r.Denominator)
}

func (r Rat) toSignedRational() exifcommon.SignedRational {
	return exifcommon.SignedRational{Numerator: r.Numerator, Denominator: r.Denominator}
}

func (r Rat) IsZero() bool {
	return r.Numerator == 0 && r.Denominator == 0
}

func (r Rat) String() string {
	return fmt.Sprintf("%v/%v", r.Numerator, r.Denominator)
}

type URat struct {
	Numerator   uint32
	Denominator uint32
}

func NewURatFromRational(v exifcommon.Rational) URat {
	return URat{v.Numerator, v.Denominator}
}

func (r URat) Float64() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func (r URat) Float32() float32 {
	return float32(r.Numerator) / float32(r.Denominator)
}

func (r URat) IsZero() bool {
	return r.Numerator == 0 && r.Denominator == 0
}

func (r URat) toRational() exifcommon.Rational {
	return exifcommon.Rational{Numerator: r.Numerator, Denominator: r.Denominator}
}

func (r URat) String() string {
	return fmt.Sprintf("%v/%v", r.Numerator, r.Denominator)
}

type LensInfo struct {
	MinFocalLength           URat `json:"minFocalLength,omitempty"`
	MaxFocalLength           URat `json:"maxFocalLength,omitempty"`
	MinFNumberMinFocalLength URat `json:"minFNumberMinFocalLength,omitempty"`
	MinFNumberMaxFocalLength URat `json:"MinFNumberMaxFocalLength,omitempty"`
}

func NewLensInfoFromRational(vals []exifcommon.Rational) (LensInfo, error) {
	ret := LensInfo{}
	if len(vals) != 4 {
		return ret, fmt.Errorf("Expected 4 lensinfo values got %v", len(vals))
	}
	ret.MinFocalLength = NewURatFromRational(vals[0])
	ret.MaxFocalLength = NewURatFromRational(vals[1])
	ret.MinFNumberMinFocalLength = NewURatFromRational(vals[2])
	ret.MinFNumberMaxFocalLength = NewURatFromRational(vals[3])
	return ret, nil
}

func (li LensInfo) toExif() []exifcommon.Rational {
	return []exifcommon.Rational{li.MinFocalLength.toRational(), li.MaxFocalLength.toRational(),
		li.MinFNumberMinFocalLength.toRational(), li.MinFNumberMaxFocalLength.toRational()}
}

func (li LensInfo) String() string {
	return fmt.Sprintf("[%v,%v,%v,%v]", li.MinFocalLength, li.MaxFocalLength,
		li.MinFNumberMinFocalLength, li.MinFNumberMaxFocalLength)
}

/*
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
*/
