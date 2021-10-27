package metadata

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	"github.com/dsoprea/go-iptc"
	"strings"
	"time"
	"trimmer.io/go-xmp/xmp"
)

type MetaData struct {
	xmp         *xmp.Document
	iptc        map[iptc.StreamTagKey][]iptc.TagData
	ifd         *exif.IfdIndex
	Summary     *ExifCompact
	ImageWidth  uint
	ImageHeight uint
}

type Rat struct {
	Numerator   int64
	Denominator int64
}

func (r Rat) Float64() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func (r Rat) Float32() float32 {
	return float32(r.Numerator) / float32(r.Denominator)
}

func (r Rat) String() string {
	return fmt.Sprintf("%v/%v", r.Numerator, r.Denominator)
}

type ExifCompact struct {
	Title    string   `json:"title,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
	Software string   `json:"software,omitempty"`
	Rating   uint16   `json:"rating,omitempty"`

	CameraMake              string `json:"cameraMake,omitempty"`
	CameraModel             string `json:"cameraModel,omitempty"`
	LensInfo                string `json:"lensInfo,omitempty"`
	LensModel               string `json:"lensModel,omitempty"`
	LensMake                string `json:"lensMake,omitempty"`
	FocalLength             Rat    `json:"focalLength,omitempty"`
	FocalLengthIn35mmFormat uint16 `json:"focalLengthIn35mmFormat,omitempty"`
	MaxApertureValue        Rat    `json:"maxApertureValue,omitempty"`
	Flash                   uint16 `json:"flash,omitempty"`

	ExposureTime         Rat    `json:"exposureTime,omitempty"`
	ExposureCompensation Rat    `json:"exposureCompensationn,omitempty"`
	ExposureProgram      uint16 `json:"exposoureProgram,omitempty"`
	FNumber              Rat    `json:"fNumber,omitempty"`
	ISO                  uint16 `json:"ISO,omitempty"`
	ColorSpace           uint16 `json:"colorSpace,omitempty"`
	XResolution          Rat    `json:"xResolution,omitempty"`
	YResolution          Rat    `json:"yResolution,omitempty"`

	DateTimeOriginal   string `json:"dateTimeOriginal,omitempty"`
	OffsetTimeOriginal string `json:"offsetTimeOriginal,omitempty"`
	DateTime           string `json:"dateTime,omitempty"`
	OffsetTime         string `json:"offsetTime,omitempty"`

	OriginalDate time.Time     `json:"originalDate,omitempty"`
	ModifyDate   time.Time     `json:"modifyDate,omitempty"`
	GPSInfo      *exif.GpsInfo `json:"gpsInfo,omitempty"`

	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}

func (ec ExifCompact) String() string {
	sb := &strings.Builder{}
	sb.WriteString("Summary:{\n")
	fmt.Fprintf(sb, "  Title: %v\n", ec.Title)
	fmt.Fprintf(sb, "  Keywords: %v\n", strings.Join(ec.Keywords, ", "))
	fmt.Fprintf(sb, "  Sofware: %v\n", ec.Software)
	fmt.Fprintf(sb, "  Rating: %v\n", ec.Rating)
	fmt.Fprintf(sb, "  Camera Make: %v\n", ec.CameraMake)
	fmt.Fprintf(sb, "  Camera Model: %v\n", ec.CameraModel)
	fmt.Fprintf(sb, "  Lens Info: %v\n", ec.LensInfo)
	fmt.Fprintf(sb, "  Lens Model: %v\n", ec.LensModel)
	fmt.Fprintf(sb, "  Lens Make: %v\n", ec.LensMake)
	fmt.Fprintf(sb, "  Focal Length: %v\n", ec.FocalLength.Float32())
	fmt.Fprintf(sb, "  Focal Length 35MM: %v\n", ec.FocalLengthIn35mmFormat)
	fmt.Fprintf(sb, "  Max Aperture Value: %v\n", ec.MaxApertureValue.Float32())
	fmt.Fprintf(sb, "  Flash Mode: %v\n", FlashModeName(ec.Flash))
	fmt.Fprintf(sb, "  Exposure Time: %v\n", ec.ExposureTime)
	fmt.Fprintf(sb, "  Exposure Compensation: %v\n", ec.ExposureCompensation.Float32())
	fmt.Fprintf(sb, "  Exposure Program: %v\n", ExposureProgamName(ec.ExposureProgram))
	fmt.Fprintf(sb, "  Fnumber: %v\n", ec.FNumber.Float32())
	fmt.Fprintf(sb, "  ISO: %v\n", ec.ISO)
	fmt.Fprintf(sb, "  ColorSpace: %v\n", ColorSpaceName(ec.ColorSpace))
	fmt.Fprintf(sb, "  XResolution: %v\n", ec.XResolution)
	fmt.Fprintf(sb, "  YResolution: %v\n", ec.YResolution)
	fmt.Fprintf(sb, "  OriginalDate: %v\n", ec.OriginalDate)
	fmt.Fprintf(sb, "  ModifyDate: %v\n", ec.ModifyDate)
	fmt.Fprintf(sb, "  GPSInfo:%v\n", ec.GPSInfo)
	fmt.Fprintf(sb, "  City: %v\n", ec.City)
	fmt.Fprintf(sb, "  State/Province: %v\n", ec.State)
	fmt.Fprintf(sb, "  Country:%v\n", ec.Country)
	sb.WriteString("}\n")
	/*
		GPSInfo *exif.GpsInfo `json:"gpsInfo,omitempty"`

		City         string `json:"city,omitempty"`
		Country      string  `json:"country,omitempty"`
		State        string  `json:"state,omitempty"`

	*/
	return sb.String()
}
