package metadata

import (
	"encoding/binary"
)

// Video containers are iso base media files (mp4) or their QuickTime ancestor
// (mov). Both are a sequence of boxes, each a 4 byte big endian size followed by
// a 4 byte type. A modern file opens with an ftyp box naming a major brand and a
// list of compatible brands, and the brand is what tells the formats apart.
//
// heif, avif and m4a share the container, so an ftyp box alone does not make a
// file a video: their brands are recognised and rejected

// major brands that identify an mp4 video
var mp4Brands = map[string]bool{
	"isom": true, "iso2": true, "iso3": true, "iso4": true, "iso5": true, "iso6": true,
	"mp41": true, "mp42": true, "avc1": true, "dash": true,
	"M4V ": true, "M4VH": true, "M4VP": true,
	"3gp4": true, "3gp5": true, "3gp6": true, "3g2a": true,
}

const quickTimeBrand = "qt  "

// brands of formats that share the container but are not video: heif and avif
// still images, and m4a/m4b audio. One of these anywhere in the ftyp box means
// the file is not treated as video
var notVideoBrands = map[string]bool{
	"heic": true, "heix": true, "hevc": true, "hevx": true, "heim": true, "heis": true,
	"mif1": true, "msf1": true, "avif": true, "avis": true,
	"M4A ": true, "M4B ": true, "M4P ": true,
}

// the first box of a QuickTime file that predates ftyp
var legacyQuickTimeBoxes = map[string]bool{
	"moov": true, "mdat": true, "wide": true, "pnot": true,
}

// detectVideo identifies an mp4 or mov from its leading bytes, or returns
// FormatUnknown
func detectVideo(data []byte) Format {
	if len(data) < 8 {
		return FormatUnknown
	}
	size := binary.BigEndian.Uint32(data[:4])
	boxType := string(data[4:8])

	if boxType != "ftyp" {
		//a box size of 0 (to end of file) or 1 (64 bit size follows) is only
		//plausible for mdat, the others always carry a real size
		if legacyQuickTimeBoxes[boxType] && (size >= 8 || (boxType == "mdat" && size <= 1)) {
			return FormatMov
		}
		return FormatUnknown
	}

	if size < 16 || len(data) < 12 {
		return FormatUnknown
	}
	major := string(data[8:12])
	if notVideoBrands[major] {
		return FormatUnknown
	}
	if major == quickTimeBrand {
		return FormatMov
	}

	//compatible brands follow the minor version, up to the end of the box or
	//of what we were given
	end := min(int(size), len(data))
	var compatible []string
	for i := 16; i+4 <= end; i += 4 {
		compatible = append(compatible, string(data[i:i+4]))
	}
	for _, b := range compatible {
		if notVideoBrands[b] {
			return FormatUnknown
		}
	}

	if mp4Brands[major] {
		return FormatMp4
	}
	//an unfamiliar major brand, camera vendors have their own, is still a
	//video when it declares compatibility with a known one
	for _, b := range compatible {
		if b == quickTimeBrand {
			return FormatMov
		}
		if mp4Brands[b] {
			return FormatMp4
		}
	}
	return FormatUnknown
}
