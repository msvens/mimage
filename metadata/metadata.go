package metadata

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"image/jpeg"
	"io/ioutil"
	"strings"
	"time"
	//_ "trimmer.io/go-xmp/models"
)

var ErrParseImage = errors.New("Could not parse image")

//var ErrJpegWrongFileExt = fmt.Errorf("Only .jpeg or .jpg file extension allowed")

type Summary struct {
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
	ExposureCompensation    Rat           `json:"exposureCompensation,omitempty"`
	ExposureProgram         uint16        `json:"exposureProgram,omitempty"`
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

func (ec Summary) String() string {
	sb := &strings.Builder{}
	sb.WriteString("Summary:{\n")
	sb.WriteString(fmt.Sprintf("  Title: %v\n", ec.Title))
	sb.WriteString(fmt.Sprintf("  Keywords: %v\n", strings.Join(ec.Keywords, ", ")))
	sb.WriteString(fmt.Sprintf("  Software: %v\n", ec.Software))
	sb.WriteString(fmt.Sprintf("  Rating: %v\n", ec.Rating))
	sb.WriteString(fmt.Sprintf("  Camera Make: %v\n", ec.CameraMake))
	sb.WriteString(fmt.Sprintf("  Camera Model: %v\n", ec.CameraModel))
	sb.WriteString(fmt.Sprintf("  Lens Info: %v\n", ec.LensInfo))
	sb.WriteString(fmt.Sprintf("  Lens Model: %v\n", ec.LensModel))
	sb.WriteString(fmt.Sprintf("  Lens Make: %v\n", ec.LensMake))
	sb.WriteString(fmt.Sprintf("  Focal Length: %v\n", ec.FocalLength.Float32()))
	sb.WriteString(fmt.Sprintf("  Focal Length 35MM: %v\n", ec.FocalLengthIn35mmFormat))
	sb.WriteString(fmt.Sprintf("  Max Aperture Value: %v\n", ec.MaxApertureValue.Float32()))
	sb.WriteString(fmt.Sprintf("  Flash Mode: %v\n", ExifValueString(ExifIFD, ExifIFD_Flash, ec.FlashMode)))
	sb.WriteString(fmt.Sprintf("  Exposure Time: %v\n", ec.ExposureTime))
	sb.WriteString(fmt.Sprintf("  Exposure Compensation: %v\n", ec.ExposureCompensation.Float32()))
	sb.WriteString(fmt.Sprintf("  Exposure Program: %v\n", ExifValueString(ExifIFD, ExifIFD_ExposureProgram, ec.ExposureProgram)))
	sb.WriteString(fmt.Sprintf("  Fnumber: %v\n", ec.FNumber.Float32()))
	sb.WriteString(fmt.Sprintf("  ISO: %v\n", ec.ISO))
	sb.WriteString(fmt.Sprintf("  ColorSpace: %v\n", ExifValueString(ExifIFD, ExifIFD_ColorSpace, ec.ColorSpace)))
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

type MetaData struct {
	xmpData     XmpData
	iptcData    *IptcData
	exifData    *ExifData
	summary     *Summary
	summaryErr  error
	ImageWidth  uint
	ImageHeight uint
}

func parseJpegBytes(data []byte) (*jpegstructure.SegmentList, error) {
	jmp := jpegstructure.NewJpegMediaParser()
	intfc, err := jmp.ParseBytes(data)
	if err != nil {
		return nil, ErrParseImage
	}
	segments := intfc.(*jpegstructure.SegmentList)
	return segments, nil
}

func NewMetaDataFromFile(filename string) (*MetaData, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return NewMetaData(data)
}

func NewMetaData(data []byte) (*MetaData, error) {
	ret := MetaData{}
	segments, err := parseJpegBytes(data)
	if err != nil {
		return nil, err
	}

	var exifErr, xmpErr, iptcErr error

	ret.exifData, exifErr = NewExifData(segments)
	if exifErr != nil && exifErr != ErrExifNoData {
		return nil, exifErr
	}

	ret.iptcData, iptcErr = NewIptcData(segments)
	if iptcErr != nil && iptcErr != ErrNoIptc {
		return nil, iptcErr
	}

	ret.xmpData, xmpErr = NewXmpData(segments)
	if xmpErr != nil && xmpErr != ErrNoXmp {
		return nil, xmpErr
	}

	//Extract ImageWidth/Height
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return &ret, err
	}

	ret.ImageWidth = uint(img.Bounds().Dx())
	ret.ImageHeight = uint(img.Bounds().Dy())

	return &ret, nil
}

func (md *MetaData) Exif() *ExifData {
	return md.exifData
}

func (md *MetaData) Iptc() *IptcData {
	return md.iptcData
}

func (md *MetaData) Summary() *Summary {
	if md.summary != nil {
		return md.summary
	}
	md.summary = &Summary{}
	var exifErr, iptcErr, xmpErr error
	if !md.exifData.IsEmpty() {
		exifErr = md.extractExifTags()
	}
	if !md.xmpData.IsEmpty() {
		xmpErr = md.extractXmp()
	}
	if !md.iptcData.IsEmpty() {
		iptcErr = md.extractIPTC()
	}
	if exifErr != nil {
		md.summaryErr = exifErr
	} else if iptcErr != nil {
		md.summaryErr = iptcErr
	} else if xmpErr != nil {
		md.summaryErr = xmpErr
	}
	return md.summary
}

func (md *MetaData) SummaryErr() error {
	return md.summaryErr
}

func (md *MetaData) Xmp() XmpData {
	return md.xmpData
}

func (md *MetaData) String() string {
	sb := strings.Builder{}
	sb.WriteString("==MetaData==\n==Summary==\n")
	sb.WriteString(md.Summary().String())
	sb.WriteString("\n==Exif==\n")
	sb.WriteString(md.exifData.String())
	sb.WriteString("\n==IPTC==\n")
	sb.WriteString(md.iptcData.String())
	sb.WriteString("\n==XMP==\n")
	sb.WriteString(md.xmpData.String())
	return sb.String()
}

func (md *MetaData) extractIPTC() error {
	var err error
	md.summary.Title = md.iptcData.GetTitle()
	md.summary.Keywords = md.iptcData.GetKeywords()
	return err
}

func (md *MetaData) extractExifTags() error {
	var err error

	scanR := func(tagId ExifTag, dest interface{}) {
		e := md.exifData.ScanIfdRoot(tagId, dest)
		if e != nil && e != ErrExifTagNotFound {
			err = e
		}
	}
	scanE := func(tagId ExifTag, dest interface{}) {
		e := md.exifData.ScanIfdExif(tagId, dest)
		if e != nil && e != ErrExifTagNotFound {
			err = e
		}
	}

	scanR(IFD_Make, &md.summary.CameraMake)
	scanR(IFD_Model, &md.summary.CameraModel)
	scanE(ExifIFD_LensInfo, &md.summary.LensInfo)
	if md.summary.LensInfo == (LensInfo{}) {
		scanR(IFD_DNGLensInfo, &md.summary.LensInfo)
	}
	scanE(ExifIFD_LensModel, &md.summary.LensModel)
	scanE(ExifIFD_LensMake, &md.summary.LensMake)
	scanE(ExifIFD_FocalLength, &md.summary.FocalLength)
	scanE(ExifIFD_FocalLengthIn35mmFormat, &md.summary.FocalLengthIn35mmFormat)
	scanE(ExifIFD_MaxApertureValue, &md.summary.MaxApertureValue)
	scanE(ExifIFD_Flash, &md.summary.FlashMode)
	scanE(ExifIFD_ExposureTime, &md.summary.ExposureTime)
	scanE(ExifIFD_ExposureCompensation, &md.summary.ExposureCompensation)
	scanE(ExifIFD_ExposureProgram, &md.summary.ExposureProgram)
	scanE(ExifIFD_FNumber, &md.summary.FNumber)
	scanE(ExifIFD_ISO, &md.summary.ISO)
	scanE(ExifIFD_ColorSpace, &md.summary.ColorSpace)
	scanR(IFD_XResolution, &md.summary.XResolution)
	scanR(IFD_YResolution, &md.summary.YResolution)
	scanR(IFD_Software, &md.summary.Software)

	_ = md.exifData.ScanExifDate(OriginalDate, &md.summary.OriginalDate)
	_ = md.exifData.ScanExifDate(ModifyDate, &md.summary.ModifyDate)

	//GPSInfo
	if md.exifData.HasIfd(GpsIFD) {
		md.summary.GPSInfo, _ = md.exifData.GpsInfo()
	}
	return err
}

func (md *MetaData) extractXmp() error {

	if md.summary.Title == "" {
		md.summary.Title = md.xmpData.GetTitle()
	}
	if len(md.summary.Keywords) == 0 {
		md.summary.Keywords = md.xmpData.GetKeywords()
	}
	if md.summary.Rating == 0 {
		md.summary.Rating = md.xmpData.GetRating()
	}
	if psm := md.xmpData.PhotoShop(); psm != nil {
		md.summary.City = psm.City
		md.summary.Country = psm.Country
		md.summary.State = psm.State
	}
	return nil
}
