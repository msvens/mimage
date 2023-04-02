package metadata

import (
	"fmt"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

// Rat represents a signed Exif float
type Rat struct {
	Numerator   int32
	Denominator int32
}

// URat represents an unsigned Exif float
type URat struct {
	Numerator   uint32
	Denominator uint32
}

// LensInfo reprsents an Exif LensInfo object
type LensInfo struct {
	MinFocalLength           URat `json:"minFocalLength,omitempty"`
	MaxFocalLength           URat `json:"maxFocalLength,omitempty"`
	MinFNumberMinFocalLength URat `json:"minFNumberMinFocalLength,omitempty"`
	MinFNumberMaxFocalLength URat `json:"MinFNumberMaxFocalLength,omitempty"`
}

func newRatFromSignedRational(v exifcommon.SignedRational) Rat {
	return Rat{v.Numerator, v.Denominator}
}

// Float64 convert to float64. Return NaNs as 0
func (r Rat) Float64() float64 {
	if f := float64(r.Numerator) / float64(r.Denominator); f != f {
		return 0
	} else {
		return f
	}
}

// Float32 convert to float32. Return NaNs as 0
func (r Rat) Float32() float32 {
	if f := float32(r.Numerator) / float32(r.Denominator); f != f {
		return 0
	} else {
		return f
	}
}

// IsZero true if both Numerator and Denominator are 0
func (r Rat) IsZero() bool {
	return r.Numerator == 0 && r.Denominator == 0
}

func (r Rat) String() string {
	return fmt.Sprintf("%v/%v", r.Numerator, r.Denominator)
}

func newURatFromRational(v exifcommon.Rational) URat {
	return URat{v.Numerator, v.Denominator}
}

// Float64 convert to float64. It will return NaN as 0
func (r URat) Float64() float64 {
	if f := float64(r.Numerator) / float64(r.Denominator); f != f {
		return 0
	} else {
		return f
	}
}

// Float32 convert to float32. It will return NaN as 0
func (r URat) Float32() float32 {
	if f := float32(r.Numerator) / float32(r.Denominator); f != f {
		return 0
	} else {
		return f
	}
}

// IsZero true if both Numerator and Denominator are 0
func (r URat) IsZero() bool {
	return r.Numerator == 0 && r.Denominator == 0
}

func (r URat) toRational() exifcommon.Rational {
	return exifcommon.Rational{Numerator: r.Numerator, Denominator: r.Denominator}
}

func (r URat) String() string {
	return fmt.Sprintf("%v/%v", r.Numerator, r.Denominator)
}

func newLensInfoFromRational(vals []exifcommon.Rational) (LensInfo, error) {
	ret := LensInfo{}
	if len(vals) != 4 {
		return ret, fmt.Errorf("Expected 4 lensinfo values got %v", len(vals))
	}
	ret.MinFocalLength = newURatFromRational(vals[0])
	ret.MaxFocalLength = newURatFromRational(vals[1])
	ret.MinFNumberMinFocalLength = newURatFromRational(vals[2])
	ret.MinFNumberMaxFocalLength = newURatFromRational(vals[3])
	return ret, nil
}

func (li LensInfo) toRational() []exifcommon.Rational {
	return []exifcommon.Rational{li.MinFocalLength.toRational(), li.MaxFocalLength.toRational(),
		li.MinFNumberMinFocalLength.toRational(), li.MinFNumberMaxFocalLength.toRational()}
}

func (li LensInfo) String() string {
	return fmt.Sprintf("[%v,%v,%v,%v]", li.MinFocalLength, li.MaxFocalLength,
		li.MinFNumberMinFocalLength, li.MinFNumberMaxFocalLength)
}
