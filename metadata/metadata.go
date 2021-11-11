package metadata

import (
	"bytes"
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"image/jpeg"
	"io/ioutil"
	_ "trimmer.io/go-xmp/models"
	"trimmer.io/go-xmp/xmp"
)

func parseJpegBytes(data []byte) (*jpegstructure.SegmentList, error) {
	jmp := jpegstructure.NewJpegMediaParser()

	intfc, err := jmp.ParseBytes(data)
	if err != nil {
		return nil, err
	}
	segments := intfc.(*jpegstructure.SegmentList)
	return segments, nil
}

func NewMetaDataFromJpegFile(filename string) (*MetaData, error) {
	if data, err := ioutil.ReadFile(filename); err != nil {
		return nil, err
	} else {
		return NewMetaDataJpeg(data)
	}
}

func NewMetaDataJpeg(data []byte) (*MetaData, error) {
	ret := MetaData{}
	segments, err := parseJpegBytes(data)
	if err != nil {
		return nil, err
	}

	var exifErr, iptcErr, xmpErr error

	ret.ifd, exifErr = loadExif(segments)
	ret.iptc, iptcErr = segments.Iptc()
	ret.xmp, xmpErr = loadXmp(segments)

	//Extract Summary:
	summaryErr := ret.extractSummary()

	//Extract ImageWidth/Height
	//ImageWidth/Height
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return &ret, err
	}

	ret.ImageWidth = uint(img.Bounds().Dx())
	ret.ImageHeight = uint(img.Bounds().Dy())

	if exifErr != nil && iptcErr != nil && xmpErr != nil {
		return &MetaData{}, fmt.Errorf("Could not read any metadata (exif, iptc, xmp)")
	} else if summaryErr != nil {
		return &ret, summaryErr
	} else {
		return &ret, nil
	}
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

func (md *MetaData) IfdRoot() *exif.Ifd {
	if md.HasExif() {
		return md.ifd.RootIfd
	} else {
		return nil
	}
}

func (md *MetaData) IfdExif() *exif.Ifd {
	if md.HasExif() {
		return md.ifd.Lookup["IFD/Exif"]
	} else {
		return nil
	}
}

func (md *MetaData) IfdGPS() *exif.Ifd {
	if md.HasExif() {
		return md.ifd.Lookup["IFD/GPSInfo"]
	} else {
		return nil
	}
}

func (md *MetaData) PrintIfd() string {
	if md.HasExif() {
		return PrintExif(md.ifd)
	} else {
		return "No Exif defined"
	}
}

func (md *MetaData) WriteExif() error {
	return nil
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

func (md *MetaData) ScanIptc1Tag(dataset uint8, dest interface{}) error {
	if !md.HasIptc() {
		return NoIptcErr
	}
	return scanIptcTag(md.iptc, IPTCEnvelop, dataset, dest)
}

func (md *MetaData) ScanIptc2Tag(dataset uint8, dest interface{}) error {
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
		return nil, err
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
		return nil, err
	}
	str, err := s.FormattedXmp()
	if err != nil {
		return nil, err
	}
	model := &xmp.Document{}

	err = xmp.Unmarshal([]byte(str), model)
	if err != nil {
		return nil, err
	} else {
		return model, err
	}
}
