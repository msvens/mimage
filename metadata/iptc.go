package metadata

import (
	"fmt"
	"github.com/dsoprea/go-iptc"
	"strings"
)

func DumpIPTCTags(tags map[iptc.StreamTagKey][]iptc.TagData) {
	for k, v := range tags {
		name := IptcTagName(k.RecordNumber, k.DatasetNumber)
		str := []string{}
		for _, data := range v {
			str = append(str, string(data))
		}
		fmt.Printf("%s (%v,%v): [%s]\n", name, k.RecordNumber, k.DatasetNumber, strings.Join(str, ", "))
	}
}

func IptcTagName(record, dataset uint8) string {
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

func scanIptcTag(tags map[iptc.StreamTagKey][]iptc.TagData, record, dataset uint8, dest interface{}) error {

	tagKey := iptc.StreamTagKey{RecordNumber: record, DatasetNumber: dataset}

	tagData, found := tags[tagKey]
	if !found {
		return IptcTagNotFoundErr
	}

	wrongType := false
	switch dtype := dest.(type) {
	case *string:
		if tagData[0].IsPrintable() {
			*dtype = string(tagData[0])
		} else {
			wrongType = true
		}
	case *[]string:
		for _, v := range tagData {
			if v.IsPrintable() {
				*dtype = append(*dtype, string(v))
			} else {
				wrongType = true
				break
			}

		}
	}
	if wrongType {
		return IptcUndefinedTypeErr
	}
	return nil
}
