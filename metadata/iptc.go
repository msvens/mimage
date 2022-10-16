package metadata

import (
	"errors"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"strings"
	"time"
	"unicode"
)

// Iptc related errors
var (
	ErrNoIptc            = errors.New("No IPTC data")
	ErrIptcTagNotFound   = errors.New("Iptc tag not found")
	ErrIptcTagValue      = errors.New("Iptc tag value not corret")
	ErrIptcUndefinedType = errors.New("Could not parse ")
)

var iptcUtfCharSet = string([]byte{27, 37, 71})

// IptcData holds a map of iptc record tags
type IptcData struct {
	raw map[IptcRecordTag]IptcRecordDataset
}

// IptcDate specifies date/time tag
type IptcDate int

// The different Iptc Date/Times
const (
	DateSent IptcDate = iota
	ReleaseDate
	ExpirationDate
	DateCreated
	DigitalCreationDate
)

// Iptc date & time format
const (
	IptcShortDate      = "20060102"
	IptcLongDateOffset = "20060102150405-0700"
	IptcLongDate       = "20060102150405"
	IptcTime           = "150405-0700"
)

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
		return time.Time{}, ErrIptcTagValue
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

// IptcTagName returns the common name for a tag in record
func IptcTagName(record IptcRecord, tag IptcTag) string {
	if desc, found := IptcTagDescriptions[IptcRecordTag{Record: record, Tag: tag}]; found {
		return desc.Name
	}
	return fmt.Sprintf("Unknown Tag. Record: %v, Dataset: %v", record, tag)
}

// NewIptcData creates IptcData from a jpeg segment list
func NewIptcData(segments *jpegstructure.SegmentList) (*IptcData, error) {
	if segments == nil {
		return nil, fmt.Errorf("Segmentlist is nil")
	}
	raw, err := ParseIptcJpeg(segments)
	return &IptcData{raw}, err
}

// IsEmpty returns true if IptcData has no tags
func (ipd *IptcData) IsEmpty() bool {
	return len(ipd.raw) == 0
}

// RawIptc returns the raw iptc data
func (ipd *IptcData) RawIptc() map[IptcRecordTag]IptcRecordDataset {
	return ipd.raw
}

// GetDate retrieves the given date value (including time if exists). In case of an error/missing date
// returns zero time
func (ipd *IptcData) GetDate(dateTag IptcDate) time.Time {
	ret := time.Time{}
	_ = ipd.ScanDate(dateTag, &ret)
	return ret
}

// GetKeywords  retrieves IPTCApplication_Keywords. Returnes an empty slice in case of an error
func (ipd *IptcData) GetKeywords() []string {
	ret := []string{}
	if err := ipd.ScanApplication(IPTCApplication_Keywords, &ret); err != nil {
		return []string{}
	}
	return ret
}

// GetTitle retrieves the IPTCApplication_ObjectName. Returns the empty string in case of an error
func (ipd *IptcData) GetTitle() string {
	ret := ""
	if err := ipd.ScanApplication(IPTCApplication_ObjectName, &ret); err != nil {
		return ""
	}
	return ret
}

// Scan reads tag from record into dest
func (ipd *IptcData) Scan(record IptcRecord, tag IptcTag, dest interface{}) error {
	if ipd.IsEmpty() {
		return ErrNoIptc
	}

	tagKey := IptcRecordTag{record, tag}
	recdata, found := ipd.raw[tagKey]
	if !found {
		return ErrIptcTagNotFound
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
				*dtype = vals
				/*for _, v := range vals {
					*dtype = append(*dtype, v)
				}*/
				return nil
			}
		case *[]uint8:
			if recdata.Repeatable && recdata.Type == IptcUint8 {
				vals := recdata.Data.([]uint8)
				*dtype = vals
				//*dtype = append(*dtype, vals...)
				/*for _, v := range vals {
					*dtype = append(*dtype, v)
				}*/
				return nil
			}
		case *[]uint16:
			if recdata.Repeatable && recdata.Type == IptcUint16 {
				vals := recdata.Data.([]uint16)
				*dtype = vals
				/**dtype = append(*dtype, vals...)
				for _, v := range vals {
					*dtype = append(*dtype, v)
				}*/
				return nil
			}
		case *[]uint32:
			if recdata.Repeatable && recdata.Type == IptcUint32 {
				vals := recdata.Data.([]uint32)
				*dtype = vals
				/*for _, v := range vals {
					*dtype = append(*dtype, v)
				}*/
				return nil
			}
		case *[][]byte:
			if recdata.Type == IptcUndef {
				vals := recdata.Data.([][]byte)
				*dtype = vals
				/*for _, v := range vals {
					*dtype = append(*dtype, v)
				}*/
				return nil
			}
		}
	}
	return ErrIptcUndefinedType
}

// ScanDate reads the specified date (both date and time) into dest
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
		return ErrIptcTagNotFound
	}
	if err != nil {
		return err
	}
	*dest, err = iptcDateTime(dateStr, timeStr)
	return err
}

// ScanEnvelope reads tag from IPTCEnvelope into dest
func (ipd *IptcData) ScanEnvelope(tag IptcTag, dest interface{}) error {
	return ipd.Scan(IPTCEnvelope, tag, dest)
}

// ScanApplication reads tag from IPTCApplication into dest
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
