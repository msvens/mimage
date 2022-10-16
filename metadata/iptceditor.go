package metadata

import (
	"bytes"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"github.com/msvens/mimage/photoshop"
	"time"
)

const envelopRecordVersion = uint16(4)
const jpegFileFormat = uint16(11)
const applicationRecordVersion = uint16(4)

//var ps30Prefix = []byte("Photoshop 3.0\000")

func iptcDateTimeStr(t time.Time) (string, string) {
	return t.Format(IptcShortDate), t.Format(IptcTime)
}

// IptcEditor holds raw iptc data
type IptcEditor struct {
	raw        map[IptcRecordTag]IptcRecordDataset
	resources  map[uint16]photoshop.ImageResource
	segmentIdx int
	dirty      bool
}

// NewIptcEditor from a jpeg segment list
func NewIptcEditor(sl *jpegstructure.SegmentList) (*IptcEditor, error) {
	if sl == nil {
		return nil, fmt.Errorf("Segmentlist is nil")
	}
	ret := IptcEditor{}
	ret.raw = map[IptcRecordTag]IptcRecordDataset{}
	var err error
	ret.segmentIdx, ret.resources, err = photoshop.ParseJpeg(sl)
	if err != nil && err != photoshop.ErrNoPhotoshopBlock {
		return &ret, err
	}
	iptcData, ok := ret.resources[photoshop.IptcId]
	if !ok { //simply return
		return &ret, nil
	}
	ret.raw, err = DecodeIptc(bytes.NewReader(iptcData.Data))
	return &ret, err

}

// NewIptcEditorEmpty creates a new empty IptcEditor
func NewIptcEditorEmpty(dirty bool) *IptcEditor {
	return &IptcEditor{raw: map[IptcRecordTag]IptcRecordDataset{}, segmentIdx: -1,
		resources: map[uint16]photoshop.ImageResource{}, dirty: dirty}
}

// Clear this model and set the dirty property
func (ie *IptcEditor) Clear(dirty bool) {
	ie.raw = map[IptcRecordTag]IptcRecordDataset{}
	ie.dirty = dirty
}

// IsDirty returns true if the IptcData has been changed
func (ie IptcEditor) IsDirty() bool {
	return ie.dirty
}

// IsEmpty returns true if this editor contains no data
func (ie IptcEditor) IsEmpty() bool {
	return len(ie.raw) == 0
}

// Bytes generate Photoshop Image Resource block including IPTC information
func (ie *IptcEditor) Bytes() ([]byte, error) {
	if ie.IsDirty() {
		if err := ie.setMandatoryTags(); err != nil {
			return nil, err
		}
		delete(ie.resources, photoshop.DigestId) //we discard any digest information
	}
	//generate IptcData
	out := &bytes.Buffer{}
	err := EncodeIptc(out, ie.raw)
	if err != nil {
		return nil, err
	}
	ie.resources[photoshop.IptcId] = photoshop.NewPhotoshopImageResource(photoshop.IptcId, out.Bytes())
	ie.dirty = false
	return photoshop.Marshal(ie.resources, true)
}

func (ie *IptcEditor) setApplication(tag IptcTag, value interface{}) error {
	return ie.Set(IPTCApplication, tag, value)
}

// SetDate set the given dateTag to to T. This will set both the corresponding Date and Time tags
func (ie *IptcEditor) SetDate(dateTag IptcDate, t time.Time) error {
	dateStr, timeStr := iptcDateTimeStr(t)
	var err error
	switch dateTag {
	case DateSent:
		if err = ie.SetEnvelope(IPTCEnvelope_DateSent, dateStr); err != nil {
			return err
		}
		if err = ie.SetEnvelope(IPTCEnvelope_TimeSent, timeStr); err != nil {
			return err
		}
	case ReleaseDate:
		if err = ie.setApplication(IPTCApplication_ReleaseDate, dateStr); err != nil {
			return err
		}
		if err = ie.setApplication(IPTCApplication_ReleaseTime, timeStr); err != nil {
			return err
		}
	case ExpirationDate:
		if err = ie.setApplication(IPTCApplication_ExpirationDate, dateStr); err != nil {
			return err
		}
		if err = ie.setApplication(IPTCApplication_ExpirationTime, timeStr); err != nil {
			return err
		}
	case DateCreated:
		if err = ie.setApplication(IPTCApplication_DateCreated, dateStr); err != nil {
			return err
		}
		if err = ie.setApplication(IPTCApplication_TimeCreated, timeStr); err != nil {
			return err
		}
	case DigitalCreationDate:
		if err = ie.setApplication(IPTCApplication_DigitalCreationDate, dateStr); err != nil {
			return err
		}
		if err = ie.setApplication(IPTCApplication_DigitalCreationTime, timeStr); err != nil {
			return err
		}
	default:
		return ErrIptcTagNotFound
	}
	return nil
}

// SetDirty force this editor to dirty
func (ie *IptcEditor) SetDirty() {
	ie.dirty = true
}

// SetEnvelope set IPTCEnvelop tag to value
func (ie *IptcEditor) SetEnvelope(tag IptcTag, value interface{}) error {
	return ie.Set(IPTCEnvelope, tag, value)
}

func (ie *IptcEditor) setMandatoryTags() error {
	_ = ie.SetEnvelope(IPTCEnvelope_EnvelopeRecordVersion, envelopRecordVersion)
	_ = ie.setApplication(IPTCApplication_ApplicationRecordVersion, applicationRecordVersion)
	_ = ie.SetEnvelope(IPTCEnvelope_FileFormat, jpegFileFormat)
	return ie.SetEnvelope(IPTCEnvelope_CodedCharacterSet, iptcUtfCharSet)
}

// SetKeywords sets IPTCApplication_Keywords to keywords
func (ie *IptcEditor) SetKeywords(keywords []string) error {
	return ie.Set(IPTCApplication, IPTCApplication_Keywords, keywords)
}

// Set sets record with tag to value
func (ie *IptcEditor) Set(record IptcRecord, tag IptcTag, value interface{}) error {
	rt := IptcRecordTag{record, tag}
	tagDesc, found := IptcTagDescriptions[rt]
	if !found {
		return ErrIptcTagNotFound
	}
	newTag := IptcRecordDataset{Record: rt.Record, Tag: rt.Tag, Type: tagDesc.Type, Repeatable: tagDesc.Repeatable}
	valueOk := false
	if !tagDesc.Repeatable {
		switch vtype := value.(type) {
		case string:
			if tagDesc.Type == IptcString || (tagDesc.Type == IptcDigits && isDigits(vtype)) {
				newTag.Data = vtype
				valueOk = true
			}
		case uint8:
			if tagDesc.Type == IptcUint8 {
				newTag.Data = vtype
				valueOk = true
			}
		case uint16:
			if tagDesc.Type == IptcUint16 {
				newTag.Data = vtype
				valueOk = true
			}
		case uint32:
			if tagDesc.Type == IptcUint32 {
				newTag.Data = vtype
				valueOk = true
			}
		case []byte:
			if tagDesc.Type == IptcUndef {
				newTag.Data = vtype
				valueOk = true
			}
		}
	} else {
		switch vtype := value.(type) {
		case []string:
			if tagDesc.Type == IptcString || (tagDesc.Type == IptcDigits && isDigitsSlice(vtype)) {
				newTag.Data = vtype
				valueOk = true
			}
		case []uint8:
			if tagDesc.Type == IptcUint8 {
				newTag.Data = vtype
				valueOk = true
			}
		case []uint16:
			if tagDesc.Type == IptcUint16 {
				newTag.Data = vtype
				valueOk = true
			}
		case []uint32:
			if tagDesc.Type == IptcUint32 {
				newTag.Data = vtype
				valueOk = true
			}
		case [][]byte:
			if tagDesc.Type == IptcUndef {
				newTag.Data = vtype
				valueOk = true
			}
		}
	}
	if valueOk {
		ie.raw[rt] = newTag
		ie.dirty = true
		return nil
	}
	return ErrIptcTagValue
}

// SetTitle sets IPTCApplication_ObjectName to title
func (ie *IptcEditor) SetTitle(title string) error {
	return ie.Set(IPTCApplication, IPTCApplication_ObjectName, title)
}
