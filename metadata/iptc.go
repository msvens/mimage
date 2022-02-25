package metadata

import (
	"errors"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"strings"
	"time"
	"unicode"
)

var NoIptcErr = errors.New("No IPTC data")
var IptcTagNotFoundErr = errors.New("Iptc tag not found")
var IptcTagValueErr = errors.New("Iptc tag value not corret")
var IptcUndefinedTypeErr = errors.New("Could not parse ")
var iptcUtfCharSet = string([]byte{27, 37, 71})

type IptcData struct {
	raw map[IptcRecordTag]IptcRecordDataset
}

type IptcDate int

const (
	DateSent IptcDate = iota
	ReleaseDate
	ExpirationDate
	DateCreated
	DigitalCreationDate
)

//Time format constants
const IptcShortDate = "20060102"
const IptcLongDateOffset = "20060102150405-0700"
const IptcLongDate = "20060102150405"
const IptcTime = "150405-0700"

//HHMMSS±HHMM
func iptcDateTime(dateStr string, timeStr string) (time.Time, error) {
	tstr := dateStr + timeStr
	switch len(tstr) {
	case 8:
		return time.Parse(IptcShortDate, tstr)
	case 19:
		return time.Parse(IptcLongDateOffset, tstr)
	case 14:
		return time.Parse(IptcLongDate, tstr)
	default:
		return time.Time{}, IptcTagValueErr
	}
}

func isDigitsSlice(strs []string) bool {
	for _, s := range strs {
		if !isDigits(s) {
			return false
		}
	}
	return true
}

func isDigits(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

/*
func convertStreamKeyToRecord(key iptc.StreamTagKey, data []iptc.TagData) (IptcRecordDataset, error) {
	rt := IptcRecordTag{Record: IptcRecord(key.RecordNumber), Tag: IptcTag(key.DatasetNumber)}
	tagDesc, found := IptcTagDescriptions[rt]
	if !found || len(data) < 1 {
		return IptcRecordDataset{}, IptcTagNotFoundErr
	}
	ret := IptcRecordDataset{Record: rt.Record, Tag: rt.Tag, Type: tagDesc.Type, Repeatable: tagDesc.Repeatable}
	switch tagDesc.Type {
	case IptcString: //willingling ignore any non utf8 encoded strings
		if !tagDesc.Repeatable {
			ret.Data = string(data[0])
		} else {
			vals := []string{}
			for _, d := range data {
				vals = append(vals, string(d))
			}
			ret.Data = vals
		}
	case IptcDigits:
		if !tagDesc.Repeatable {
			if isDigits(string(data[0])) {
				ret.Data = string(data[0])
			} else {
				return ret, IptcTagValueErr
			}
		} else {
			vals := []string{}
			for _, d := range data {
				if isDigits(string(d)) {
					vals = append(vals, string(d))
				} else {
					return ret, IptcTagValueErr
				}
			}
			ret.Data = vals
		}
	case IptcUint8:
		if !tagDesc.Repeatable {
			r := bytes.NewReader(data[0])
			val := uint8(0)
			if err := binary.Read(r, binary.BigEndian, &val); err != nil {
				return ret, IptcTagValueErr
			} else {
				ret.Data = val
			}
		} else {
			vals := []uint8{}
			for _, d := range data {
				r := bytes.NewReader(d)
				val := uint8(0)
				if err := binary.Read(r, binary.BigEndian, &val); err != nil {
					return ret, IptcTagValueErr
				} else {
					vals = append(vals, val)
				}
			}
			ret.Data = vals
		}
	case IptcUint16:
		if !tagDesc.Repeatable {
			r := bytes.NewReader(data[0])
			val := uint16(0)
			if err := binary.Read(r, binary.BigEndian, &val); err != nil {
				return ret, IptcTagValueErr
			} else {
				ret.Data = val
			}
		} else {
			vals := []uint16{}
			for _, d := range data {
				r := bytes.NewReader(d)
				val := uint16(0)
				if err := binary.Read(r, binary.BigEndian, &val); err != nil {
					return ret, IptcTagValueErr
				} else {
					vals = append(vals, val)
				}
			}
			ret.Data = vals
		}
	case IptcUint32:
		if !tagDesc.Repeatable {
			r := bytes.NewReader(data[0])
			val := uint32(0)
			if err := binary.Read(r, binary.BigEndian, &val); err != nil {
				return ret, IptcTagValueErr
			} else {
				ret.Data = val
			}
		} else {
			vals := []uint32{}
			for _, d := range data {
				r := bytes.NewReader(d)
				val := uint32(0)
				if err := binary.Read(r, binary.BigEndian, &val); err != nil {
					return ret, IptcTagValueErr
				} else {
					vals = append(vals, val)
				}
			}
			ret.Data = vals
		}
	case IptcUndef: //undef, etc
		if !tagDesc.Repeatable {
			var d []byte
			d = data[0]
			ret.Data = d
		} else {
			vals := [][]byte{}
			for _, d := range data {
				vals = append(vals, d)
			}
			ret.Data = vals
		}
	default:
		return ret, IptcUndefinedTypeErr
	}
	return ret, nil
}

func convertStreamKeysToRecords(source map[iptc.StreamTagKey][]iptc.TagData) map[IptcRecordTag]IptcRecordDataset{
	ret := map[IptcRecordTag]IptcRecordDataset{}
	if len(source) == 0 {
		return ret
	}
	for k, v := range source {
		if rd,err := convertStreamKeyToRecord(k,v); err != nil {
			fmt.Println(err)
		} else {
			ret[IptcRecordTag{rd.Record, rd.Tag}] = rd
		}
	}
	return ret
}
*/

func IptcTagName(record IptcRecord, tag IptcTag) string {
	if desc, found := IptcTagDescriptions[IptcRecordTag{Record: record, Tag: tag}]; found {
		return desc.Name
	} else {
		return fmt.Sprintf("Unknown Tag. Record: %v, Dataset: %v", record, tag)
	}

}

func NewIptcData(segments *jpegstructure.SegmentList) (*IptcData, error) {
	if segments == nil {
		return nil, fmt.Errorf("Segmentlist is nil")
	}
	raw, err := ParseIptcJpeg(segments)
	return &IptcData{raw}, err
	/*
		raw, err := segments.Iptc() //dont care about the error
		return &IptcData{convertStreamKeysToRecords(raw)}, err
	*/
}

func (ipd *IptcData) IsEmpty() bool {
	return len(ipd.raw) == 0
}

func (ipd *IptcData) RawIptc() map[IptcRecordTag]IptcRecordDataset {
	return ipd.raw
}

//Retrieves the given date value (including time if exists). In case of an error/missing date
//returns zero time
func (ipd *IptcData) GetDate(dateTag IptcDate) time.Time {
	ret := time.Time{}
	_ = ipd.ScanDate(dateTag, &ret)
	return ret
}

//Retrives IPTCApplication_Keywords. Returnes an empty slice in case of an error
func (ipd *IptcData) GetKeywords() []string {
	ret := []string{}
	if err := ipd.ScanApplication(IPTCApplication_Keywords, &ret); err != nil {
		return []string{}
	}
	return ret
}

//Retrieves the IPTCApplication_ObjectName. Returns the emtpy string in case of an error
func (ipd *IptcData) GetTitle() string {
	ret := ""
	if err := ipd.ScanApplication(IPTCApplication_ObjectName, &ret); err != nil {
		return ""
	}
	return ret
}

func (ipd *IptcData) Scan(record IptcRecord, tag IptcTag, dest interface{}) error {
	if ipd.IsEmpty() {
		return NoIptcErr
	}

	tagKey := IptcRecordTag{record, tag}
	recdata, found := ipd.raw[tagKey]
	if !found {
		return IptcTagNotFoundErr
	}
	if !recdata.Repeatable {
		switch dtype := dest.(type) {
		case *string:
			if recdata.Type == IptcString || recdata.Type == IptcDigits {
				v := recdata.Data.(string)
				*dtype = v
				return nil
			}
		case *uint8:
			if recdata.Type == IptcUint8 {
				v := recdata.Data.(uint8)
				*dtype = v
				return nil
			}
		case *uint16:
			if recdata.Type == IptcUint16 {
				v := recdata.Data.(uint16)
				*dtype = v
				return nil
			}
		case *uint32:
			if recdata.Type == IptcUint32 {
				v := recdata.Data.(uint32)
				*dtype = v
				return nil
			}
		case *[]byte:
			v := recdata.Data.([]byte)
			*dtype = v
			return nil
		}
	} else {
		switch dtype := dest.(type) {
		case *[]string:
			if recdata.Type == IptcString || recdata.Type == IptcDigits {
				vals := recdata.Data.([]string)
				for _, v := range vals {
					*dtype = append(*dtype, v)
				}
				return nil
			}
		case *[]uint8:
			if recdata.Repeatable && recdata.Type == IptcUint8 {
				vals := recdata.Data.([]uint8)
				for _, v := range vals {
					*dtype = append(*dtype, v)
				}
				return nil
			}
		case *[]uint16:
			if recdata.Repeatable && recdata.Type == IptcUint16 {
				vals := recdata.Data.([]uint16)
				for _, v := range vals {
					*dtype = append(*dtype, v)
				}
				return nil
			}
		case *[]uint32:
			if recdata.Repeatable && recdata.Type == IptcUint32 {
				vals := recdata.Data.([]uint32)
				for _, v := range vals {
					*dtype = append(*dtype, v)
				}
				return nil
			}
		case *[][]byte:
			if recdata.Type == IptcUndef {
				vals := recdata.Data.([][]byte)
				for _, v := range vals {
					*dtype = append(*dtype, v)
				}
				return nil
			}
		}
	}
	return IptcUndefinedTypeErr
}

func (ipd *IptcData) ScanDate(dateTag IptcDate, dest *time.Time) error {
	dateStr := ""
	timeStr := ""
	var err error
	switch dateTag {
	case DateSent:
		err = ipd.ScanEnvelope(IPTCEnvelope_DateSent, &dateStr)
		_ = ipd.ScanEnvelope(IPTCEnvelope_TimeSent, &timeStr)
	case ReleaseDate:
		err = ipd.ScanApplication(IPTCApplication_ReleaseDate, &dateStr)
		_ = ipd.ScanApplication(IPTCApplication_ReleaseTime, &timeStr)
	case ExpirationDate:
		err = ipd.ScanApplication(IPTCApplication_ExpirationDate, &dateStr)
		_ = ipd.ScanApplication(IPTCApplication_ExpirationTime, &timeStr)
	case DateCreated:
		err = ipd.ScanApplication(IPTCApplication_DateCreated, &dateStr)
		_ = ipd.ScanApplication(IPTCApplication_TimeCreated, &timeStr)
	case DigitalCreationDate:
		err = ipd.ScanApplication(IPTCApplication_DigitalCreationDate, &dateStr)
		_ = ipd.ScanApplication(IPTCApplication_DigitalCreationTime, &timeStr)
	default:
		return IptcTagNotFoundErr
	}
	if err != nil {
		return err
	}
	*dest, err = iptcDateTime(dateStr, timeStr)
	return err
}

func (ipd *IptcData) ScanEnvelope(tag IptcTag, dest interface{}) error {
	return ipd.Scan(IPTCEnvelope, tag, dest)
}

func (ipd *IptcData) ScanApplication(tag IptcTag, dest interface{}) error {
	return ipd.Scan(IPTCApplication, tag, dest)
}

func (ipd *IptcData) String() string {
	if ipd.IsEmpty() {
		return "No IPTC data"
	}
	buff := strings.Builder{}
	for k, v := range ipd.raw {
		name := IptcTagName(k.Record, k.Tag)
		if k.Record == IPTCEnvelope && k.Tag == IPTCEnvelope_CodedCharacterSet && v.Data == iptcUtfCharSet {
			buff.WriteString(fmt.Sprintf("%s (%v,%v): %v\n", name, k.Record, k.Tag, "UTF"))
		} else {
			buff.WriteString(fmt.Sprintf("%s (%v,%v): %v\n", name, k.Record, k.Tag, v.Data))
		}

	}
	return buff.String()
}
