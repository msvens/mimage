package metadata

import (
	"bytes"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"github.com/msvens/mimage/photoshop"
	"time"
)

const EnvelopRecordVersion = uint16(4)
const JpegFileFormat = uint16(11)
const ApplicationRecordVersion = uint16(4)

//var ps30Prefix = []byte("Photoshop 3.0\000")

func iptcDateTimeStr(t time.Time) (string, string) {
	return t.Format(IptcShortDate), t.Format(IptcTime)
}

type IptcEditor struct {
	raw        map[IptcRecordTag]IptcRecordDataset
	resources  map[uint16]photoshop.PhotoshopImageResource
	segmentIdx int
	dirty      bool
}

//func TryPhotoshopInfo(f string) error {
//	jmp := jpegstructure.NewJpegMediaParser()
//	intfc, err := jmp.ParseFile(f)
//	if err != nil {
//		return err
//	}
//	segments := intfc.(*jpegstructure.SegmentList)
//	_, s, err := segments.FindIptc()
//	if err != nil {
//		return err
//	}
//
//	//save the data to files:
//	if err = os.WriteFile("assets/leicaPhotoshopResourcesWithPrefix.data", s.Data, 0644); err != nil {
//		fmt.Println("Could not save data file")
//	}
//	l := len(ps30Prefix)
//
//	if len(s.Data) < l {
//		return fmt.Errorf("no photoshop data")
//	}
//
//	if bytes.Equal(s.Data[:l], ps30Prefix) == false {
//		return fmt.Errorf("no photoshop data")
//	}
//
//	fmt.Println("len buff: ",len(s.Data))
//	data := s.Data[l:]
//	if err = os.WriteFile("assets/leicaPhotoshopResources.data", data, 0644); err != nil {
//		fmt.Println("Could not save data file")
//	}
//	b := bytes.NewBuffer(data)
//
//	//index, err := photoshopinfo.ReadPhotoshop30Info(b)
//	//fmt.Println(len(s.Data))
//	//b := bytes.NewBuffer(s.Data)
//	index, err := photoshop.Decode(b, false)
//	if err != nil {
//		return fmt.Errorf("Could not read photoshop30Info: %v", err)
//	}
//	for _,v := range index {
//		fmt.Println(v)
//	}
//	/*
//	for k,v := range index {
//		if k == 0x0425 {
//			fmt.Println(string(v.Data))
//		}
//		fmt.Printf("key: %#04x name %s type: %s datalen %v\n",k, v.Name, v.RecordType, len(v.Data))
//	}*/
//	return nil
//}

//Todo: Error Handling
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
	if iptcData, ok := ret.resources[photoshop.IptcId]; !ok { //simply return
		return &ret, nil
	} else {
		ret.raw, err = DecodeIptc(bytes.NewReader(iptcData.Data))
		return &ret, err
	}
}

func NewIptcEditorEmpty(dirty bool) *IptcEditor {
	return &IptcEditor{raw: map[IptcRecordTag]IptcRecordDataset{}, segmentIdx: -1,
		resources: map[uint16]photoshop.PhotoshopImageResource{}, dirty: dirty}
}

func (ie *IptcEditor) Clear(dirty bool) {
	ie.raw = map[IptcRecordTag]IptcRecordDataset{}
	ie.dirty = dirty
}

func (ie IptcEditor) IsDirty() bool {
	return ie.dirty
}

func (ie IptcEditor) IsEmpty() bool {
	return len(ie.raw) == 0
}

//generate Photoshop Image Resource block including IPTC information
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

/*
func (ie *IptcEditor) Bytes(prefix bool) ([]byte, error){
	if ie.IsDirty() {
		ie.setMandatoryTags()
	}
	out := bytes.Buffer{}
	if prefix {
		out.Write(ps30Prefix)
	}
	keys := []IptcRecordTag{}
	for k,_ := range ie.raw {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i,j int) bool {
		if keys[i].Record == keys[j].Record {
			return keys[i].Tag < keys[j].Tag
		} else {
			return keys[i].Record < keys[j].Record
		}
	})
	for _,k := range keys {
		if err := encodeIptcRecordTag(&out, ie.raw[k]); err != nil {
			return nil, err
		}
	}
	ie.dirty = false
	return out.Bytes(), nil
}
*/

/*
func encodeData(out *bytes.Buffer, record IptcRecord, tag IptcTag, val interface{}) error {
	//tag marker, record and tag/dataset:
	fmt.Printf("encode (%v,%v):%v buff length: %v\n",record,tag,val,out.Len())
	binary.Write(out, defaultEncoding, TagMarker)
	binary.Write(out, defaultEncoding, record)
	binary.Write(out, defaultEncoding, tag)
	fmt.Println("buff len: ",out.Len())
	switch vtype := val.(type) {
	case string:
		if len(vtype) > 32767 {
			return fmt.Errorf("Could not encode value larger thatn 32767 bytes...")
		}
		binary.Write(out, defaultEncoding, uint16(len(vtype)))
		out.WriteString(vtype)
	case uint8:
		binary.Write(out, defaultEncoding, uint16(1))
		binary.Write(out, defaultEncoding, vtype)
	case uint16:
		fmt.Println("writing uint16")
		binary.Write(out, defaultEncoding, uint16(2))
		binary.Write(out, defaultEncoding, vtype)
	case uint32:
		binary.Write(out, defaultEncoding, uint16(4))
		binary.Write(out, defaultEncoding, vtype)
	case []byte:
		if len(vtype) > 32767 {
			return fmt.Errorf("Could not encode value larger than 32767 bytes...")
		}
		binary.Write(out, defaultEncoding, uint16(len(vtype)))
		out.Write(vtype)
	default:
		return IptcUndefinedTypeErr
	}
	fmt.Printf("encode done (%v,%v), buff length: %v\n",record,tag, out.Len())
	return nil
}
*/
/*
func encodeIptcRecordTag(out *bytes.Buffer, rd IptcRecordDataset) error {
	if rd.Repeatable {
		switch vtype := rd.Data.(type) {
		case []string:
			for _,val := range vtype {
				if err := encodeData(out, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		case []uint8:
			for _,val := range vtype {
				if err := encodeData(out, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		case []uint16:
			for _,val := range vtype {
				if err := encodeData(out, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		case []uint32:
			for _,val := range vtype {
				if err := encodeData(out, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		case [][]byte:
			for _,val := range vtype {
				if err := encodeData(out, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		default:
			return IptcUndefinedTypeErr
		}
	} else {
		return encodeData(out, rd.Record, rd.Tag, rd.Data)
	}
	return nil
}
*/

func (ie IptcEditor) SegmentIndex() int {
	return ie.segmentIdx
}

func (ie *IptcEditor) SetApplication(tag IptcTag, value interface{}) error {
	return ie.Set(IPTCApplication, tag, value)
}

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
		if err = ie.SetApplication(IPTCApplication_ReleaseDate, dateStr); err != nil {
			return err
		}
		if err = ie.SetApplication(IPTCApplication_ReleaseTime, timeStr); err != nil {
			return err
		}
	case ExpirationDate:
		if err = ie.SetApplication(IPTCApplication_ExpirationDate, dateStr); err != nil {
			return err
		}
		if err = ie.SetApplication(IPTCApplication_ExpirationTime, timeStr); err != nil {
			return err
		}
	case DateCreated:
		if err = ie.SetApplication(IPTCApplication_DateCreated, dateStr); err != nil {
			return err
		}
		if err = ie.SetApplication(IPTCApplication_TimeCreated, timeStr); err != nil {
			return err
		}
	case DigitalCreationDate:
		if err = ie.SetApplication(IPTCApplication_DigitalCreationDate, dateStr); err != nil {
			return err
		}
		if err = ie.SetApplication(IPTCApplication_DigitalCreationTime, timeStr); err != nil {
			return err
		}
	default:
		return IptcTagNotFoundErr
	}
	return nil
}

func (ie *IptcEditor) SetDirty() {
	ie.dirty = true
}

func (ie *IptcEditor) SetEnvelope(tag IptcTag, value interface{}) error {
	return ie.Set(IPTCEnvelope, tag, value)
}

func (ie *IptcEditor) setMandatoryTags() error {
	_ = ie.SetEnvelope(IPTCEnvelope_EnvelopeRecordVersion, EnvelopRecordVersion)
	_ = ie.SetApplication(IPTCApplication_ApplicationRecordVersion, ApplicationRecordVersion)
	_ = ie.SetEnvelope(IPTCEnvelope_FileFormat, JpegFileFormat)
	return ie.SetEnvelope(IPTCEnvelope_CodedCharacterSet, iptcUtfCharSet)
}

func (ie *IptcEditor) SetKeywords(keywords []string) error {
	return ie.Set(IPTCApplication, IPTCApplication_Keywords, keywords)
}

//Todo: Check that value is Ok
func (ie *IptcEditor) Set(record IptcRecord, tag IptcTag, value interface{}) error {
	rt := IptcRecordTag{record, tag}
	tagDesc, found := IptcTagDescriptions[rt]
	if !found {
		return IptcTagNotFoundErr
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
	} else {
		return IptcTagValueErr
	}
}

func (ie *IptcEditor) SetTitle(title string) error {
	return ie.Set(IPTCApplication, IPTCApplication_ObjectName, title)
}
