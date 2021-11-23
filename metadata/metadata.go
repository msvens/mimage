package metadata

import (
	"bytes"
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	exifundefined "github.com/dsoprea/go-exif/v3/undefined"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"image/jpeg"
	"io/ioutil"
	"time"
	_ "trimmer.io/go-xmp/models"
	"trimmer.io/go-xmp/xmp"
)

func parseJpegBytes(data []byte) (*jpegstructure.SegmentList, error) {
	jmp := jpegstructure.NewJpegMediaParser()

	if intfc, err := jmp.ParseBytes(data); err != nil {
		return nil, ParseImageErr
	} else {
		segments := intfc.(*jpegstructure.SegmentList)
		return segments, nil
	}
}

func ParseFile(filename string) (*MetaData, error) {
	if data, err := ioutil.ReadFile(filename); err != nil {
		return nil, err
	} else {
		return Parse(data)
	}
}

func Parse(data []byte) (*MetaData, error) {
	ret := MetaData{}
	segments, err := parseJpegBytes(data)
	if err != nil {
		return nil, err
	}

	var exifErr, xmpErr error

	ret.ifd, exifErr = loadExif(segments)
	if exifErr != nil && exifErr != NoExifErr {
		return nil, exifErr
	}
	ret.iptc, _ = segments.Iptc() //dont care about the error

	ret.xmp, xmpErr = loadXmp(segments)
	if xmpErr != nil && xmpErr != NoXmpErr {
		return nil, xmpErr
	}

	//Extract Summary:
	if err = ret.extractSummary(); err != nil {
		return &ret, err
	}
	//summaryErr := ret.extractSummary()

	//Extract ImageWidth/Height
	//ImageWidth/Height
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return &ret, err
	}

	ret.ImageWidth = uint(img.Bounds().Dx())
	ret.ImageHeight = uint(img.Bounds().Dy())

	return &ret, nil
}

func (md *MetaData) HasExif() bool {
	if md.ifd != nil {
		return true
	} else {
		return false
	}
}
func (md *MetaData) HasXmp() bool {
	if md.xmp != nil {
		return true
	} else {
		return false
	}
}
func (md *MetaData) HasIptc() bool {
	if md.iptc != nil {
		return true
	} else {
		return false
	}
}

//Convinince method to retrive IFD_ImageDescription. The MetaDataEditor has
//a corresponding method to set the IFD_ImageDescription
func (md *MetaData) GetIfdImageDescription() (string, error) {
	ret := ""
	if !md.HasExif() {
		return ret, NoExifErr
	}
	err := md.ScanIfdRootTag(IFD_ImageDescription, &ret)
	return ret, err
}

//Convinience method to retrieve Exif_UserComment. As the UserComment is
//an undefined field this method will assume it has been set by the corresponding
//MetaDataEditor method
func (md *MetaData) GetIfdUserComment() (string, error) {
	ret := exifundefined.Tag9286UserComment{}
	if !md.HasExif() {
		return "", NoExifErr
	}
	err := md.ScanIfdExifTag(Exif_UserComment, &ret)
	return string(ret.EncodingBytes), err
}

func (md *MetaData) IfdRoot() *exif.Ifd {
	if md.HasExif() {
		return md.ifd.RootIfd
	} else {
		return nil
	}
}

func (md *MetaData) GetIfd(ifdPath string) *exif.Ifd {
	if md.HasExif() {
		return md.ifd.Lookup[ifdPath]
	}
	return nil
}

func (md *MetaData) HasIfd(ifdPath string) bool {
	if md.HasExif() {
		_, found := md.ifd.Lookup[ifdPath]
		return found
	}
	return false
}

func (md *MetaData) IfdExif() *exif.Ifd {
	return md.GetIfd(ExifPath)
}

func (md *MetaData) IfdGPS() *exif.Ifd {
	return md.GetIfd(GPSInfoPath)
}

func (md *MetaData) PrintIfd() string {
	if md.HasExif() {
		return printExif(md.ifd)
	} else {
		return "No Exif defined"
	}
}

func (md *MetaData) PrintIptc() string {
	if md.HasIptc() {
		return PrintIptc(md.iptc)
	} else {
		return "No Iptc defined"
	}
}

func (md *MetaData) ScanIfdRootTag(tagId uint16, dest interface{}) error {
	if !md.HasExif() {
		return NoExifErr
	}
	return ScanIfdTag(md.IfdRoot(), tagId, dest)
}

func (md *MetaData) ScanIfdExifTag(tagId uint16, dest interface{}) error {
	if !md.HasExif() {
		return NoExifErr
	}
	return ScanIfdTag(md.IfdExif(), tagId, dest)
}

func (md *MetaData) ScanIfdDate(dateTag ExifDate, dest *time.Time) error {
	if !md.HasExif() {
		return NoExifErr
	}
	var t, o string
	var err error
	switch dateTag {
	case OriginalDate:
		if err = ScanIfdTag(md.IfdExif(), Exif_DateTimeOriginal, &t); err != nil {
			return err
		}
		_ = ScanIfdTag(md.IfdExif(), Exif_OffsetTimeOriginal, &o) //dont care about offset errors
		*dest, err = ParseIfdDateTime(t, o)
		if err != nil {
			return err
		}
	case ModifyDate:
		if err = ScanIfdTag(md.IfdRoot(), IFD_DateTime, &t); err != nil {
			return err
		}
		_ = ScanIfdTag(md.IfdExif(), Exif_OffsetTime, &o) //dont care about offset errors
		*dest, err = ParseIfdDateTime(t, o)
		if err != nil {
			return err
		}
	case DigitizedDate:
		if err = ScanIfdTag(md.IfdExif(), Exif_DateTimeDigitized, &t); err != nil {
			return err
		}
		_ = ScanIfdTag(md.IfdExif(), Exif_OffsetTimeDigitized, &o) //dont care about offset errors
		*dest, err = ParseIfdDateTime(t, o)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown date to scan: %v", dateTag)
	}
	return nil
}

func (md *MetaData) ScanIptcEnvelopTag(dataset uint8, dest interface{}) error {
	if !md.HasIptc() {
		return NoIptcErr
	}
	return scanIptcTag(md.iptc, IPTCEnvelop, dataset, dest)
}

func (md *MetaData) ScanApplicationTag(dataset uint8, dest interface{}) error {
	if !md.HasIptc() {
		return NoIptcErr
	}
	return scanIptcTag(md.iptc, IPTCApplication, dataset, dest)
}

func (md *MetaData) Xmp() *xmp.Document {
	return md.xmp
}

func loadExif(segments *jpegstructure.SegmentList) (*exif.IfdIndex, error) {
	var rawExif []byte
	var ifdMapping *exifcommon.IfdMapping

	var err error
	if _, rawExif, err = segments.Exif(); err != nil {
		return nil, NoExifErr
	}
	if ifdMapping, err = exifcommon.NewIfdMappingWithStandard(); err != nil {
		return nil, err
	}
	ti := exif.NewTagIndex()
	if err = exif.LoadStandardTags(ti); err != nil {
		return nil, err
	}

	if _, index, err := exif.Collect(ifdMapping, ti, rawExif); err != nil {
		return nil, err
	} else {
		return &index, nil
	}
}

func loadXmp(segments *jpegstructure.SegmentList) (*xmp.Document, error) {
	_, s, err := segments.FindXmp()
	if err != nil {
		return nil, NoXmpErr
	}
	str, err := s.FormattedXmp()
	if err != nil {
		//We should log errors
		return nil, NoXmpErr
	}
	model := &xmp.Document{}

	err = xmp.Unmarshal([]byte(str), model)
	if err != nil {
		return nil, NoXmpErr
	} else {
		return model, nil
	}
}
