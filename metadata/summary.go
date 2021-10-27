package metadata

import (
	"fmt"
	"github.com/msvens/mimage/metadata/ifdexif"
	"github.com/msvens/mimage/metadata/ifdroot"
	"github.com/msvens/mimage/metadata/iptc2"
	"trimmer.io/go-xmp/models/dc"
	"trimmer.io/go-xmp/models/ps"
	xmpbase "trimmer.io/go-xmp/models/xmp_base"
)

func (md *MetaData) extractSummary() error {
	var exifErr, iptcErr, xmpErr error

	md.Summary = &ExifCompact{}

	if md.HasExif() {
		exifErr = md.extractExifTags()
	} else {
		fmt.Println("md has no exif data")
	}
	if md.HasIptc() {
		iptcErr = md.extractIPTC()
	} else {
		fmt.Println("md has no iptc data")
	}
	if md.HasXmp() {
		xmpErr = md.extractXmp()
	} else {
		fmt.Println("md has no xmp data")
	}

	if exifErr != nil {
		return exifErr
	} else if iptcErr != nil {
		return iptcErr
	} else if xmpErr != nil {
		return xmpErr
	} else {
		return nil
	}

}

func (md *MetaData) extractIPTC() error {
	var err error

	if err = md.ScanIptc2Tag(iptc2.ObjectName, &md.Summary.Title); err != nil {
		fmt.Println(err)
	}

	if err = md.ScanIptc2Tag(iptc2.Keywords, &md.Summary.Keywords); err != nil {
		fmt.Println(err)
	}

	return nil
}

func (md *MetaData) extractExifTags() error {
	var err error

	scanR := func(tagId uint16, dest interface{}) {
		e := md.ScanIfdRootTag(tagId, dest)
		if e != nil && e != IfdTagNotFoundErr {
			err = e
		}
	}
	scanE := func(tagId uint16, dest interface{}) {
		e := md.ScanIfdExifTag(tagId, dest)
		if e != nil && e != IfdTagNotFoundErr {
			err = e
		}
	}

	scanR(ifdroot.Make, &md.Summary.CameraMake)
	scanR(ifdroot.Model, &md.Summary.CameraModel)
	scanR(ifdroot.LensInfo, &md.Summary.LensInfo)
	scanE(ifdexif.LensModel, &md.Summary.LensModel)
	scanE(ifdexif.LensMake, &md.Summary.LensMake)
	scanE(ifdexif.FocalLength, &md.Summary.FocalLength)
	scanE(ifdexif.FocalLengthIn35mmFilm, &md.Summary.FocalLengthIn35mmFormat)
	scanE(ifdexif.MaxApertureValue, &md.Summary.MaxApertureValue)
	scanE(ifdexif.Flash, &md.Summary.Flash)
	scanE(ifdexif.ExposureTime, &md.Summary.ExposureTime)
	scanE(ifdexif.ExposureBiasValue, &md.Summary.ExposureCompensation)
	scanE(ifdexif.ExposureProgram, &md.Summary.ExposureProgram)
	scanE(ifdexif.FNumber, &md.Summary.FNumber)
	scanE(ifdexif.ISOSpeedRatings, &md.Summary.ISO)
	scanE(ifdexif.ColorSpace, &md.Summary.ColorSpace)
	scanR(ifdroot.XResolution, &md.Summary.XResolution)
	scanR(ifdroot.YResolution, &md.Summary.YResolution)
	scanE(ifdexif.OffsetTime, &md.Summary.OffsetTime)
	scanE(ifdexif.OffsetTimeOriginal, &md.Summary.OffsetTimeOriginal)
	scanR(ifdroot.DateTime, &md.Summary.DateTime)
	scanE(ifdexif.DateTimeOriginal, &md.Summary.DateTimeOriginal)
	scanR(ifdroot.Software, &md.Summary.Software)

	md.Summary.OriginalDate, _ = ParseIfdDateTime(md.Summary.DateTimeOriginal, md.Summary.OffsetTimeOriginal)
	md.Summary.ModifyDate, _ = ParseIfdDateTime(md.Summary.DateTime, md.Summary.OffsetTime)

	//GPSInfo
	if gpsIfd := md.IfdGPS(); gpsIfd != nil {
		if gpsinfo, err := gpsIfd.GpsInfo(); err != nil {
			fmt.Println("error retrieving gps info: ", err.Error())
		} else {
			md.Summary.GPSInfo = gpsinfo
		}
	}
	return err
}

func (md *MetaData) extractXmp() error {

	var err error
	//Try to extract keywords and title from dublin core
	if dcmodel := dc.FindModel(md.xmp); dcmodel == nil {
		err = NoXmpModelErr
	} else {
		t := dcmodel.Title.Default()
		if t != "" {
			md.Summary.Title = t
		}
		if len(dcmodel.Subject) > 0 {
			md.Summary.Keywords = dcmodel.Subject
		}
	}
	//Try to extract rating and software from xmp
	if xmpmodel := xmpbase.FindModel(md.xmp); xmpmodel == nil {
		err = NoXmpModelErr
	} else {
		md.Summary.Rating = uint16(xmpmodel.Rating)
		if xmpmodel.CreatorTool != "" {
			md.Summary.Software = xmpmodel.CreatorTool.String()
		}
	}
	//Try to extract location information from photoshop
	if psmodel := ps.FindModel(md.xmp); psmodel == nil {
		err = NoXmpModelErr
	} else {
		if psmodel.City != "" {
			md.Summary.City = psmodel.City
		}
		if psmodel.Country != "" {
			md.Summary.Country = psmodel.Country
		}
		if psmodel.State != "" {
			md.Summary.State = psmodel.State
		}
	}
	return err
}
