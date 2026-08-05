package metadata

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	exifundefined "github.com/dsoprea/go-exif/v3/undefined"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"time"
)

const exifEditorSoftware = "github.com/msvens/mimage (go-exif)"

// ExifEditor holds an IfdBuilder
type ExifEditor struct {
	rootIb *exif.IfdBuilder
	dirty  bool
}

func exifOffsetString(t time.Time) string {
	return t.Format("-0700")
}

// NewExifEditor from a jpeg segment list
func NewExifEditor(sl *jpegstructure.SegmentList) (*ExifEditor, error) {
	if sl == nil {
		return &ExifEditor{}, fmt.Errorf("nil segment list")
	}
	rootIfd, _, err := sl.Exif()
	if err != nil {
		if errors.Is(err, exif.ErrNoExif) {
			return NewExifEditorEmpty(false)
		}
		return &ExifEditor{}, err
	}
	rootIb := exif.NewIfdBuilderFromExistingChain(rootIfd)
	return &ExifEditor{rootIb, false}, nil
}

// NewExifEditorFromIfd creates an editor seeded from an already parsed ifd
// chain, for instance one read straight out of a tiff, skipping any tag in
// excludeTags. The editor is marked dirty since the chain did not come from the
// image being edited.
//
// Tags are excluded while rebuilding rather than deleted afterwards, so a tag
// go-exif chokes on is never read at all. Only the root ifd and its children
// are copied, not the nextIfd chain, so a tiff thumbnail ifd is left behind
// rather than carried into the jpeg
func NewExifEditorFromIfd(rootIfd *exif.Ifd, excludeTags []ExifTag) (*ExifEditor, error) {
	if rootIfd == nil {
		return NewExifEditorEmpty(false)
	}
	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, err
	}
	ti := exif.NewTagIndex()
	if err = exif.LoadStandardTags(ti); err != nil {
		return nil, err
	}
	exclude := make([]uint16, 0, len(excludeTags))
	for _, t := range excludeTags {
		exclude = append(exclude, uint16(t))
	}
	ib, err := buildIfdChain(rootIfd, im, ti, exclude)
	if err != nil {
		return nil, err
	}
	//Rebuilding drops what go-exif cannot decode, but it can also encode a tag
	//into something it then refuses to read: FileSource is one. Rather than
	//keep a list of which tags misbehave, verify the result and rebuild once
	//without whatever did not survive
	if bad := unreadableAfterEncode(ib, im, ti); len(bad) > 0 {
		if ib, err = buildIfdChain(rootIfd, im, ti, append(exclude, bad...)); err != nil {
			return nil, err
		}
	}
	return &ExifEditor{ib, true}, nil
}

// unreadableAfterEncode encodes a builder and reports any tag that cannot be
// parsed back out of the result, so it can be excluded and the chain rebuilt.
// This is the backstop for go-exif encoding something it will not read again,
// and it needs no knowledge of which tags those are
func unreadableAfterEncode(ib *exif.IfdBuilder, im *exifcommon.IfdMapping, ti *exif.TagIndex) []uint16 {
	encoded, err := exif.NewIfdByteEncoder().EncodeToExif(ib)
	if err != nil {
		return nil
	}
	_, index, err := exif.Collect(im, ti, encoded)
	if err != nil {
		return nil
	}
	var bad []uint16
	var walk func(*exif.Ifd)
	walk = func(ifd *exif.Ifd) {
		for _, ite := range ifd.Entries() {
			if _, e := ite.GetRawBytes(); e != nil {
				bad = append(bad, ite.TagId())
			}
		}
		for _, c := range ifd.Children() {
			walk(c)
		}
	}
	walk(index.RootIfd)
	return bad
}

// buildIfdChain walks an ifd and its children and rebuilds them tag by tag,
// skipping anything in exclude.
//
// Values are re-encoded from their decoded form rather than copied as raw
// bytes. go-exif's own AddTagsFromExisting does a byte copy, which forces the
// result to keep the source's byte order: re-encoding a big endian tiff as
// little endian would byte swap every number and turn ISO 100 into 25600. It
// also carries a value through even when go-exif cannot encode it back, which
// produces exif that will not parse. Decoding first means the byte order is
// ours to choose and anything go-exif does not understand is dropped here,
// deliberately and predictably, instead of corrupting the output later
func buildIfdChain(src *exif.Ifd, im *exifcommon.IfdMapping, ti *exif.TagIndex, exclude []uint16) (*exif.IfdBuilder, error) {
	ib := exif.NewIfdBuilder(im, ti, src.IfdIdentity(), binary.LittleEndian)

	excluded := func(id uint16) bool {
		for _, e := range exclude {
			if e == id {
				return true
			}
		}
		return false
	}

	for _, ite := range src.Entries() {
		//child ifds are rebuilt below, and the thumbnail offset and length are
		//recalculated by the encoder
		if ite.ChildIfdPath() != "" || ite.IsThumbnailOffset() || ite.IsThumbnailSize() {
			continue
		}
		if excluded(ite.TagId()) {
			continue
		}
		value, err := ite.Value()
		if err != nil {
			//go-exif cannot decode it, so it could not encode it back either
			continue
		}
		if err = ib.SetStandard(ite.TagId(), value); err != nil {
			//not a tag go-exif knows how to write in this ifd
			continue
		}
	}

	for _, child := range src.Children() {
		childIb, err := buildIfdChain(child, im, ti, exclude)
		if err != nil {
			return nil, err
		}
		if err = ib.AddChildIb(childIb); err != nil {
			return nil, err
		}
	}
	return ib, nil
}

// TiffDropTags are the tags that must not survive a transplant from a tiff into
// a jpeg: the ones describing how the tiff stored its pixels, plus the two that
// carry xmp and iptc, which become their own jpeg segments instead
func TiffDropTags() []ExifTag {
	return append(append([]ExifTag{}, tiffLayoutTags...), tiffSidecarTags...)
}

// NewExifEditorEmpty create a new empty editor and sets the dirty flag
func NewExifEditorEmpty(dirty bool) (*ExifEditor, error) {
	ret := ExifEditor{}
	err := ret.Clear(dirty)
	return &ret, err
}

// Clear this editor and sets the dirty flag
func (ee *ExifEditor) Clear(dirty bool) error {
	im := exifcommon.NewIfdMapping()

	if err := exifcommon.LoadStandardIfds(im); err != nil {
		return err
	}

	ti := exif.NewTagIndex()

	ee.rootIb = exif.NewIfdBuilder(im, ti,
		exifcommon.IfdStandardIfdIdentity,
		exifcommon.EncodeDefaultByteOrder)
	ee.dirty = dirty
	return nil

}

// DropMakerNote removes the MakerNote tag from the ExifIFD and returns how many
// entries were removed. The editor is marked dirty only when something was
// actually removed. Images without exif data, without an ExifIFD, or without a
// MakerNote are left untouched and return (0, nil)
func (ee *ExifEditor) DropMakerNote() (int, error) {
	if ee.rootIb == nil {
		return 0, nil
	}
	//deliberately not GetOrCreateIbFromRootIb: that would fabricate an ExifIFD
	//for images that have none and dirty an editor we never needed to touch
	exifIb, err := ee.rootIb.ChildWithTagId(uint16(IFD_ExifOffset))
	if err != nil || exifIb == nil {
		return 0, nil
	}
	n, err := exifIb.DeleteAll(uint16(ExifIFD_MakerNote))
	if err != nil {
		return 0, err
	}
	if n > 0 {
		ee.dirty = true
	}
	return n, nil
}

// tiffLayoutTags describe how a tiff stores its pixels. They are meaningless
// once the image has been re-encoded as a jpeg, and StripOffsets and friends
// are worse than meaningless: they would point at data that is no longer there
// All of these sit in ifd0 of a tiff. TileOffsets and TileByteCounts have no
// IFD_ constant because exiftool groups them under the exif ifd, but the tag
// ids are what matter here
var tiffLayoutTags = []ExifTag{
	IFD_ImageWidth, IFD_ImageHeight, IFD_BitsPerSample, IFD_Compression,
	IFD_PhotometricInterpretation, IFD_StripOffsets, IFD_SamplesPerPixel,
	IFD_RowsPerStrip, IFD_StripByteCounts, IFD_PlanarConfiguration,
	IFD_SampleFormat, IFD_Predictor, IFD_TileWidth, IFD_TileLength,
	0x0144, //TileOffsets
	0x0145, //TileByteCounts
}

// tiffSidecarTags hold payloads that a tiff keeps inside ifd0 but a jpeg keeps
// in its own segment: xmp in APP1, iptc in APP13 and an icc profile in APP2.
// Leaving them in the transplanted exif would either duplicate them or bury
// them somewhere nothing looks
var tiffSidecarTags = []ExifTag{tiffXmpTag, tiffIptcTag, tiffPhotoshopTag, tiffIccTag}

// HasMakerNote reports whether the ExifIFD currently holds a MakerNote tag
func (ee *ExifEditor) HasMakerNote() bool {
	if ee.rootIb == nil {
		return false
	}
	exifIb, err := ee.rootIb.ChildWithTagId(uint16(IFD_ExifOffset))
	if err != nil || exifIb == nil {
		return false
	}
	_, err = exifIb.FindTag(uint16(ExifIFD_MakerNote))
	return err == nil
}

// IsDirty if this editor has made any edits
func (ee ExifEditor) IsDirty() bool {
	return ee.dirty
}

// IsEmpty returns true if the IfdBuilder contains no data
func (ee ExifEditor) IsEmpty() bool {
	next, _ := ee.rootIb.NextIb()
	return len(ee.rootIb.Tags()) == 0 && next == nil
}

// IfdBuilder returns the underlying IfdBuilder. If it was changed it will also set the IFD Software tag.
func (ee *ExifEditor) IfdBuilder() (*exif.IfdBuilder, bool) {
	changed := ee.dirty
	_ = ee.setSoftware()
	return ee.rootIb, changed
}

// SetDirty force this editor to be marked as dirty
func (ee *ExifEditor) SetDirty() {
	ee.dirty = true
}

// SetDate sets the specified dateTag to time. Both the DateTime and corresponding offest
// will be set
func (ee *ExifEditor) SetDate(dateTag ExifDate, time time.Time) error {
	var err error
	offset := exifOffsetString(time)
	switch dateTag {
	case OriginalDate:
		if err = ee.SetIfdExifTag(ExifIFD_DateTimeOriginal, time); err != nil {
			return err
		}
		if err = ee.SetIfdExifTag(ExifIFD_OffsetTimeOriginal, offset); err != nil {
			return err
		}
	case ModifyDate:
		if err = ee.SetIfdRootTag(IFD_ModifyDate, time); err != nil {
			return err
		}
		if err = ee.SetIfdExifTag(ExifIFD_OffsetTime, offset); err != nil {
			return err
		}
	case DigitizedDate:
		if err = ee.SetIfdExifTag(ExifIFD_CreateDate, time); err != nil {
			return err
		}
		if err = ee.SetIfdExifTag(ExifIFD_OffsetTimeDigitized, offset); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown date to set: %v", dateTag)
	}
	return nil
}

// SetImageDescription sets IFD_ImageDescription to description
func (ee *ExifEditor) SetImageDescription(description string) error {
	return ee.SetIfdRootTag(IFD_ImageDescription, description)
}

// SetIfdExifTag sets the ExifIFD tag id to value
func (ee *ExifEditor) SetIfdExifTag(id ExifTag, value interface{}) error {
	exifIb, err := exif.GetOrCreateIbFromRootIb(ee.rootIb, IFDPaths[ExifIFD])
	if err != nil {
		return err
	}
	if err = exifIb.SetStandard(uint16(id), toGoExifValue(value)); err != nil {
		return err
	}
	ee.dirty = true
	return nil
}

// SetIfdRootTag set RootIFD tag id to value
func (ee *ExifEditor) SetIfdRootTag(id ExifTag, value interface{}) error {
	if err := ee.rootIb.SetStandard(uint16(id), toGoExifValue(value)); err != nil {
		return err
	}
	ee.dirty = true
	return nil
}

func (ee *ExifEditor) setSoftware() error {
	if !ee.dirty {
		return nil
	}
	if err := ee.SetIfdRootTag(IFD_Software, exifEditorSoftware); err != nil {
		return err
	}
	ee.dirty = false
	return nil
}

// SetUserComment sets ExifIFD_UserComment to comment using exifundefined.Tag9286UserComment. The
// comment will be Unicode encoded
func (ee *ExifEditor) SetUserComment(comment string) error {
	uc := exifundefined.Tag9286UserComment{
		EncodingType:  exifundefined.TagUndefinedType_9286_UserComment_Encoding_UNICODE,
		EncodingBytes: []byte(comment),
	}
	return ee.SetIfdExifTag(ExifIFD_UserComment, uc)
}

func toGoExifValue(value interface{}) interface{} {
	switch t := value.(type) {
	case uint16:
		return []uint16{t}
	case uint32:
		return []uint32{t}
	case float32:
		return []float32{t}
	case float64:
		return []float64{t}
	case int32:
		return []int32{t}
	case URat:
		return []exifcommon.Rational{{Denominator: t.Denominator, Numerator: t.Numerator}}
	case Rat:
		return []exifcommon.SignedRational{{Denominator: t.Denominator, Numerator: t.Numerator}}
	case []URat:
		var ret []exifcommon.Rational
		for _, v := range t {
			ret = append(ret, exifcommon.Rational{Denominator: v.Denominator, Numerator: v.Numerator})
		}
		return ret
	case []Rat:
		var ret []exifcommon.SignedRational
		for _, v := range t {
			ret = append(ret, exifcommon.SignedRational{Denominator: v.Denominator, Numerator: v.Numerator})
		}
		return ret
	case LensInfo:
		return t.toRational()
	default:
		return t
	}
}
