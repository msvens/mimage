package metadata

import (
	"bytes"

	"github.com/dsoprea/go-exif/v3"
	"github.com/msvens/mimage/photoshop"
)

// Tiff carries its metadata as plain ifd0 tag values rather than in separate
// container segments the way jpeg does
const (
	//tiffXmpTag holds an XMP packet (XMLPacket)
	tiffXmpTag = ExifTag(0x02bc)
	//tiffIptcTag holds a raw IPTC-NAA datastream
	tiffIptcTag = ExifTag(0x83bb)
	//tiffPhotoshopTag holds a photoshop image resource block, which may in turn
	//contain IPTC, exactly as a jpeg APP13 segment does
	tiffPhotoshopTag = ExifTag(0x8649)
	//tiffIccTag holds an embedded icc colour profile (InterColorProfile). A
	//jpeg carries this in APP2 rather than in exif
	tiffIccTag = ExifTag(0x8773)
)

// unreadableRootTags returns the ifd0 tags whose value go-exif cannot decode.
// Copying one into a builder fails the whole transplant, and the static drop
// list cannot anticipate every tag a camera or editor might write, so scan for
// them instead of guessing
func unreadableRootTags(ed *ExifData) []ExifTag {
	ifd := ed.Ifd(RootIFD)
	if ifd == nil {
		return nil
	}
	var ret []ExifTag
	for _, ite := range ifd.Entries() {
		if _, err := ite.GetRawBytes(); err != nil {
			ret = append(ret, ExifTag(ite.TagId()))
		}
	}
	return ret
}

// newMetaDataFromTiff reads exif, xmp and iptc from a tiff image. A tiff is an
// ifd chain from its first byte, so the exif parser can read the file directly
func newMetaDataFromTiff(data []byte) (*MetaData, error) {
	ret := MetaData{}

	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		return nil, ErrParseImage
	}
	if ret.exifData, err = NewExifDataFromBytes(rawExif); err != nil {
		return nil, ErrParseImage
	}

	//xmp and iptc live inside ifd0, and neither is required
	if b := tiffTagBytes(ret.exifData, tiffXmpTag); b != nil {
		if xmpData, e := NewXmpDataFromBytes(b); e == nil {
			ret.xmpData = xmpData
		}
	}
	ret.iptcData = &IptcData{raw: tiffIptc(ret.exifData)}

	ret.ImageWidth, ret.ImageHeight = tiffDimensions(ret.exifData)
	return &ret, nil
}

// tiffTagBytes returns the raw bytes of an ifd0 tag, or nil if it is absent
func tiffTagBytes(ed *ExifData, tag ExifTag) []byte {
	ifd := ed.Ifd(RootIFD)
	if ifd == nil {
		return nil
	}
	entries, err := ifd.FindTagWithId(uint16(tag))
	if err != nil || len(entries) == 0 {
		return nil
	}
	b, err := entries[0].GetRawBytes()
	if err != nil || len(b) == 0 {
		return nil
	}
	return b
}

// tiffIptc reads iptc from a tiff, preferring the raw IPTC-NAA tag and falling
// back to a photoshop resource block, the same wrapping a jpeg uses
func tiffIptc(ed *ExifData) map[IptcRecordTag]IptcRecordDataset {
	empty := map[IptcRecordTag]IptcRecordDataset{}
	if b := tiffTagBytes(ed, tiffIptcTag); b != nil {
		if recs, err := DecodeIptc(bytes.NewReader(b)); err == nil {
			return recs
		}
	}
	b := tiffTagBytes(ed, tiffPhotoshopTag)
	if b == nil {
		return empty
	}
	res := map[uint16]photoshop.ImageResource{}
	if err := photoshop.Unmarshal(b, false, &res); err != nil {
		return empty
	}
	iptcRes, found := res[photoshop.IptcId]
	if !found {
		return empty
	}
	if recs, err := DecodeIptc(bytes.NewReader(iptcRes.Data)); err == nil {
		return recs
	}
	return empty
}

// tiffDimensions reads the image size from ifd0 rather than decoding pixels.
// The jpeg path decodes only to learn the dimensions, which would be wasteful
// for a multi megabyte tiff
func tiffDimensions(ed *ExifData) (uint, uint) {
	var w, h uint32
	if err := ed.ScanIfdRoot(IFD_ImageWidth, &w); err != nil {
		var w16 uint16
		if e := ed.ScanIfdRoot(IFD_ImageWidth, &w16); e == nil {
			w = uint32(w16)
		}
	}
	if err := ed.ScanIfdRoot(IFD_ImageHeight, &h); err != nil {
		var h16 uint16
		if e := ed.ScanIfdRoot(IFD_ImageHeight, &h16); e == nil {
			h = uint32(h16)
		}
	}
	return uint(w), uint(h)
}
