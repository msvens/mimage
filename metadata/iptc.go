package metadata

import (
	"errors"
	"fmt"
	"github.com/dsoprea/go-iptc"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"strings"
)

var NoIptcErr = errors.New("No IPTC data")
var IptcTagNotFoundErr = errors.New("Iptc tag not found")
var IptcUndefinedTypeErr = errors.New("Could not parse ")

type IptcData struct {
	raw map[iptc.StreamTagKey][]iptc.TagData
}

func IptcTagName(record IptcRecord, tag IptcTag) string {
	var str string
	var found bool
	switch record {
	case Envelope:
		str, found = IptcEnvelopeName[tag]
	case Application:
		str, found = IptcApplicationName[tag]
	case NewsPhoto:
		str, found = IptcNewsPhotoName[tag]
	case PreObjectData:
		str, found = IptcPreObjectDataName[tag]
	case ObjectData:
		str, found = IptcObjectDataName[tag]
	case PostObjectData:
		str, found = IptcPostObjectDataName[tag]
	case FotoStation:
		str, found = IptcFotoStationName[tag]
	}
	if found {
		return str
	} else {
		return fmt.Sprintf("Unknown Tag. Record: %v, Dataset: %v", record, tag)
	}
}

func NewIptcData(segments *jpegstructure.SegmentList) (*IptcData, error) {
	raw, err := segments.Iptc() //dont care about the error
	return &IptcData{raw}, err
}

func (ipd *IptcData) IsEmpty() bool {
	return len(ipd.raw) == 0
}

func (ipd *IptcData) RawIptc() map[iptc.StreamTagKey][]iptc.TagData {
	return ipd.raw
}

func (ipd *IptcData) Scan(record IptcRecord, tag IptcTag, dest interface{}) error {
	if ipd.IsEmpty() {
		return NoIptcErr
	}

	tagKey := iptc.StreamTagKey{RecordNumber: uint8(record), DatasetNumber: uint8(tag)}

	tagData, found := ipd.raw[tagKey]
	if !found {
		return IptcTagNotFoundErr
	}

	wrongType := false
	switch dtype := dest.(type) {
	case *[]byte:
		*dtype = tagData[0]
	case *string:
		*dtype = string(tagData[0])
	case *[]string:
		for _, v := range tagData {
			*dtype = append(*dtype, string(v))
		}
	}
	if wrongType {
		return IptcUndefinedTypeErr
	}
	return nil
}

func (ipd *IptcData) ScanEnvelope(tag IptcTag, dest interface{}) error {
	return ipd.Scan(Envelope, tag, dest)
}

func (ipd *IptcData) ScanApplication(tag IptcTag, dest interface{}) error {
	return ipd.Scan(Application, tag, dest)
}

func (ipd *IptcData) String() string {
	if ipd.IsEmpty() {
		return "No IPTC data"
	}
	buff := strings.Builder{}
	for k, v := range ipd.raw {
		name := IptcTagName(IptcRecord(k.RecordNumber), IptcTag(k.DatasetNumber))
		str := []string{}
		for _, data := range v {
			str = append(str, string(data))
		}
		buff.WriteString(fmt.Sprintf("%s (%v,%v): [%s]\n", name, k.RecordNumber, k.DatasetNumber, strings.Join(str, ", ")))
	}
	return buff.String()
}
