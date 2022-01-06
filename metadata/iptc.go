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

func (ipd *IptcData) ScanIptcEnvelopTag(dataset uint8, dest interface{}) error {
	return ipd.Scan(IPTCEnvelop, dataset, dest)
}

func (ipd *IptcData) ScanApplicationTag(dataset uint8, dest interface{}) error {
	return ipd.Scan(IPTCApplication, dataset, dest)
}

func (ipd *IptcData) Scan(record, dataset uint8, dest interface{}) error {
	if ipd.IsEmpty() {
		return NoIptcErr
	}

	tagKey := iptc.StreamTagKey{RecordNumber: record, DatasetNumber: dataset}

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
		tagData[0].IsPrintable()
		/*if tagData[0].IsPrintable() {
			*dtype = string(tagData[0])
		} else {
			wrongType = true
		}*/
	case *[]string:
		for _, v := range tagData {
			*dtype = append(*dtype, string(v))
			/*if v.IsPrintable() {
				*dtype = append(*dtype, string(v))
			} else {
				wrongType = true
				break
			}*/

		}
	}
	if wrongType {
		return IptcUndefinedTypeErr
	}
	return nil
}

func (ipd *IptcData) String() string {
	if ipd.IsEmpty() {
		return "No IPTC data"
	}
	buff := strings.Builder{}
	for k, v := range ipd.raw {
		name := ipd.TagName(k.RecordNumber, k.DatasetNumber)
		str := []string{}
		for _, data := range v {
			str = append(str, string(data))
		}
		buff.WriteString(fmt.Sprintf("%s (%v,%v): [%s]\n", name, k.RecordNumber, k.DatasetNumber, strings.Join(str, ", ")))
	}
	return buff.String()
}

func (ipd *IptcData) TagName(record, dataset uint8) string {
	var str string
	var found bool
	if record == IPTCEnvelop {
		str, found = Iptc1Name[dataset]
	} else if record == IPTCApplication {
		str, found = Iptc2Name[dataset]
	} else {
		found = false //unneeded but nice for clarity
	}
	if found {
		return str
	} else {
		return fmt.Sprintf("Unknown Tag. Record: %v, Dataset: %v", record, dataset)
	}
}
