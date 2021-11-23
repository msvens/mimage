package metadata

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/dsoprea/go-iptc"
	"strings"
	"time"
	"trimmer.io/go-xmp/xmp"
)

type ImageType int

const (
	JpegType ImageType = iota
	PngType
	TiffType
)

type MetaData struct {
	xmp         *xmp.Document
	iptc        map[iptc.StreamTagKey][]iptc.TagData
	ifd         *exif.IfdIndex
	Summary     *MetaDataSummary
	ImageWidth  uint
	ImageHeight uint
}

type URat struct {
	Numerator   uint32
	Denominator uint32
}

type Rat struct {
	Numerator   int32
	Denominator int32
}

type ItpcType uint16

func (r Rat) Float64() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func (r Rat) Float32() float32 {
	return float32(r.Numerator) / float32(r.Denominator)
}

func (r Rat) toExif() exifcommon.SignedRational {
	return exifcommon.SignedRational{Numerator: r.Numerator, Denominator: r.Denominator}
}

func (r Rat) IsZero() bool {
	return r.Numerator == 0 && r.Denominator == 0
}

func fromExifRat(v exifcommon.SignedRational) Rat {
	return Rat{v.Numerator, v.Denominator}
}

func (r Rat) String() string {
	return fmt.Sprintf("%v/%v", r.Numerator, r.Denominator)
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

func (r URat) toExif() exifcommon.Rational {
	return exifcommon.Rational{Numerator: r.Numerator, Denominator: r.Denominator}
}

func fromExifURat(v exifcommon.Rational) URat {
	return URat{v.Numerator, v.Denominator}
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

func (li LensInfo) toExif() []exifcommon.Rational {
	return []exifcommon.Rational{li.MinFocalLength.toExif(), li.MaxFocalLength.toExif(),
		li.MinFNumberMinFocalLength.toExif(), li.MinFNumberMaxFocalLength.toExif()}
}

func fromExifLensInfo(vals []exifcommon.Rational) (LensInfo, error) {
	ret := LensInfo{}
	if len(vals) != 4 {
		return ret, fmt.Errorf("Expected 4 lensinfo values got %v", len(vals))
	}
	ret.MinFocalLength = fromExifURat(vals[0])
	ret.MaxFocalLength = fromExifURat(vals[1])
	ret.MinFNumberMinFocalLength = fromExifURat(vals[2])
	ret.MinFNumberMaxFocalLength = fromExifURat(vals[3])
	return ret, nil

}

func (li LensInfo) String() string {
	return fmt.Sprintf("[%v,%v,%v,%v]", li.MinFocalLength, li.MaxFocalLength,
		li.MinFNumberMinFocalLength, li.MinFNumberMaxFocalLength)
}

type MetaDataSummary struct {
	Title                   string        `json:"title,omitempty"`
	Keywords                []string      `json:"keywords,omitempty"`
	Software                string        `json:"software,omitempty"`
	Rating                  uint16        `json:"rating,omitempty"`
	CameraMake              string        `json:"cameraMake,omitempty"`
	CameraModel             string        `json:"cameraModel,omitempty"`
	LensInfo                LensInfo      `json:"lensInfo,omitempty"`
	LensModel               string        `json:"lensModel,omitempty"`
	LensMake                string        `json:"lensMake,omitempty"`
	FocalLength             URat          `json:"focalLength,omitempty"`
	FocalLengthIn35mmFormat uint16        `json:"focalLengthIn35mmFormat,omitempty"`
	MaxApertureValue        URat          `json:"maxApertureValue,omitempty"`
	FlashMode               uint16        `json:"flashMode,omitempty"`
	ExposureTime            URat          `json:"exposureTime,omitempty"`
	ExposureCompensation    Rat           `json:"exposureCompensationn,omitempty"`
	ExposureProgram         uint16        `json:"exposoureProgram,omitempty"`
	FNumber                 URat          `json:"fNumber,omitempty"`
	ISO                     uint16        `json:"ISO,omitempty"`
	ColorSpace              uint16        `json:"colorSpace,omitempty"`
	XResolution             URat          `json:"xResolution,omitempty"`
	YResolution             URat          `json:"yResolution,omitempty"`
	OriginalDate            time.Time     `json:"originalDate,omitempty"`
	ModifyDate              time.Time     `json:"modifyDate,omitempty"`
	GPSInfo                 *exif.GpsInfo `json:"gpsInfo,omitempty"`
	City                    string        `json:"city,omitempty"`
	Country                 string        `json:"country,omitempty"`
	State                   string        `json:"state,omitempty"`
}

func (ec MetaDataSummary) String() string {
	sb := &strings.Builder{}
	sb.WriteString("Summary:{\n")
	sb.WriteString(fmt.Sprintf("  Title: %v\n", ec.Title))
	sb.WriteString(fmt.Sprintf("  Keywords: %v\n", strings.Join(ec.Keywords, ", ")))
	sb.WriteString(fmt.Sprintf("  Sofware: %v\n", ec.Software))
	sb.WriteString(fmt.Sprintf("  Rating: %v\n", ec.Rating))
	sb.WriteString(fmt.Sprintf("  Camera Make: %v\n", ec.CameraMake))
	sb.WriteString(fmt.Sprintf("  Camera Model: %v\n", ec.CameraModel))
	sb.WriteString(fmt.Sprintf("  Lens Info: %v\n", ec.LensInfo))
	sb.WriteString(fmt.Sprintf("  Lens Model: %v\n", ec.LensModel))
	sb.WriteString(fmt.Sprintf("  Lens Make: %v\n", ec.LensMake))
	sb.WriteString(fmt.Sprintf("  Focal Length: %v\n", ec.FocalLength.Float32()))
	sb.WriteString(fmt.Sprintf("  Focal Length 35MM: %v\n", ec.FocalLengthIn35mmFormat))
	sb.WriteString(fmt.Sprintf("  Max Aperture Value: %v\n", ec.MaxApertureValue.Float32()))
	sb.WriteString(fmt.Sprintf("  Flash Mode: %v\n", FlashMode(ec.FlashMode)))
	sb.WriteString(fmt.Sprintf("  Exposure Time: %v\n", ec.ExposureTime))
	sb.WriteString(fmt.Sprintf("  Exposure Compensation: %v\n", ec.ExposureCompensation.Float32()))
	sb.WriteString(fmt.Sprintf("  Exposure Program: %v\n", ExposureProgram(ec.ExposureProgram)))
	sb.WriteString(fmt.Sprintf("  Fnumber: %v\n", ec.FNumber.Float32()))
	sb.WriteString(fmt.Sprintf("  ISO: %v\n", ec.ISO))
	sb.WriteString(fmt.Sprintf("  ColorSpace: %v\n", ColorSpace(ec.ColorSpace)))
	sb.WriteString(fmt.Sprintf("  XResolution: %v\n", ec.XResolution))
	sb.WriteString(fmt.Sprintf("  YResolution: %v\n", ec.YResolution))
	sb.WriteString(fmt.Sprintf("  OriginalDate: %v\n", ec.OriginalDate))
	sb.WriteString(fmt.Sprintf("  ModifyDate: %v\n", ec.ModifyDate))
	sb.WriteString(fmt.Sprintf("  GPSInfo:%v\n", ec.GPSInfo))
	sb.WriteString(fmt.Sprintf("  City: %v\n", ec.City))
	sb.WriteString(fmt.Sprintf("  State/Province: %v\n", ec.State))
	sb.WriteString(fmt.Sprintf("  Country:%v\n", ec.Country))
	sb.WriteString("}\n")
	return sb.String()
}
