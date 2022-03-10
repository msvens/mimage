package metadata

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"github.com/msvens/mimage/photoshop"
	"io"
	"sort"
)

var defaultEncoding = binary.BigEndian

const TagMarker = uint8(0x1c)

const msb = 32768

func setMsb(n uint16) uint16 {
	return n | msb
}

//math.MaxInt only supported from 1.17 and up
const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	maxInt  = 1<<(intSize-1) - 1
)

func clearMsb(n uint16) uint16 {
	return n &^ msb
}

func hasMsb(n uint16) bool {
	return !(n&msb == 0)
}

type IptcRecordDataset struct {
	Record     IptcRecord
	Tag        IptcTag
	Data       interface{}
	Type       IptcTagType
	Repeatable bool
}

//Todo: Add checks that the value is legit. The only thing we check now are digits

func decodeIptcDataToSlice(desc IptcTagDesc, data [][]byte) (interface{}, error) {
	//var err error
	switch desc.Type {
	case IptcString:
		ret := []string{}
		for _, d := range data {
			ret = append(ret, string(d))
		}
		return ret, nil
	case IptcDigits:
		ret := []string{}
		for _, d := range data {
			ret = append(ret, string(d))
		}
		if isDigitsSlice(ret) {
			return ret, nil
		}
		return ret, ErrIptcTagValue
	case IptcUint8:
		ret := []uint8{}
		for _, d := range data {
			if len(d) != 1 {
				return ret, ErrIptcTagValue
			}
			val := uint8(0)
			r := bytes.NewReader(d)
			if err := binary.Read(r, defaultEncoding, &val); err != nil {
				return ret, err
			}
			ret = append(ret, val)
		}
		return ret, nil
	case IptcUint16:
		ret := []uint16{}
		for _, d := range data {
			if len(d) != 2 {
				return ret, ErrIptcTagValue
			}
			val := uint16(0)
			r := bytes.NewReader(d)
			if err := binary.Read(r, defaultEncoding, &val); err != nil {
				return ret, err
			}
			ret = append(ret, val)
		}
		return ret, nil
	case IptcUint32:
		ret := []uint32{}
		for _, d := range data {
			if len(d) != 4 {
				return ret, ErrIptcTagValue
			}
			val := uint32(0)
			r := bytes.NewReader(d)
			if err := binary.Read(r, defaultEncoding, &val); err != nil {
				return ret, err
			}
			ret = append(ret, val)
		}
		return ret, nil
	case IptcUndef:
		return data, nil
	default:
		return nil, ErrIptcUndefinedType
	}
}

func decodeIptcData(desc IptcTagDesc, data []byte) (interface{}, error) {
	var err error
	r := bytes.NewReader(data)
	switch desc.Type {
	case IptcString:
		return string(data), nil
	case IptcDigits:
		ret := string(data)
		if isDigits(ret) {
			return ret, nil
		}
		return ret, ErrIptcTagValue
	case IptcUint8:
		ret := uint8(0)
		if len(data) != 1 {
			return ret, ErrIptcTagValue
		}
		if err = binary.Read(r, defaultEncoding, &ret); err == nil {
			return ret, nil
		}
	case IptcUint16:
		ret := uint16(0)
		if len(data) != 2 {
			return ret, ErrIptcTagValue
		}
		if err = binary.Read(r, defaultEncoding, &ret); err == nil {
			return ret, nil
		}
	case IptcUint32:
		ret := uint32(0)
		if len(data) != 4 {
			return ret, ErrIptcTagValue
		}
		if err = binary.Read(r, defaultEncoding, &ret); err == nil {
			return ret, nil
		}
	case IptcUndef:
		return data[0], nil
	default:
		err = ErrIptcUndefinedType
	}
	return nil, err
}

/*
func decodeIptcData2(desc IptcTagDesc, data [][]byte) (interface{}, error) {
	var err error
	if desc.Repeatable {
		switch desc.Type {
		case IptcString:
			ret := []string{}
			for _, d := range data {
				ret = append(ret, string(d))
			}
			return ret, nil
		case IptcDigits:
			ret := []string{}
			for _, d := range data {
				ret = append(ret, string(d))
			}
			if isDigitsSlice(ret) {
				return ret, nil
			}
			return ret, ErrIptcTagValue
		case IptcUint8:
			ret := []uint8{}
			for _, d := range data {
				if len(d) != 1 {
					return ret, ErrIptcTagValue
				}
				val := uint8(0)
				r := bytes.NewReader(d)
				if err = binary.Read(r, defaultEncoding, &val); err != nil {
					break
				}
				ret = append(ret, val)
			}
			if err == nil {
				return ret, nil
			}
		case IptcUint16:
			ret := []uint16{}
			for _, d := range data {
				if len(d) != 2 {
					return ret, ErrIptcTagValue
				}
				val := uint16(0)
				r := bytes.NewReader(d)
				if err = binary.Read(r, defaultEncoding, &val); err != nil {
					break
				}
				ret = append(ret, val)
			}
			if err == nil {
				return ret, nil
			}
		case IptcUint32:
			ret := []uint32{}
			for _, d := range data {
				if len(d) != 4 {
					return ret, ErrIptcTagValue
				}
				val := uint32(0)
				r := bytes.NewReader(d)
				if err = binary.Read(r, defaultEncoding, &val); err != nil {
					break
				}
				ret = append(ret, val)
			}
			if err == nil {
				return ret, nil
			}
		case IptcUndef:
			return data, nil
		default:
			err = ErrIptcUndefinedType
		}
		return nil, err
	}

	//non repeatable
	if len(data) != 1 {
		return nil, ErrIptcTagValue
	}
	switch desc.Type {
	case IptcString:
		return string(data[0]), nil
	case IptcDigits:
		ret := string(data[0])
		if isDigits(ret) {
			return ret, nil
		}
		return ret, ErrIptcTagValue
	case IptcUint8:
		ret := uint8(0)
		if len(data[0]) != 1 {
			return ret, ErrIptcTagValue
		}
		r := bytes.NewReader(data[0])
		if err = binary.Read(r, defaultEncoding, &ret); err == nil {
			return ret, nil
		}
	case IptcUint16:
		ret := uint16(0)
		if len(data[0]) != 2 {
			return ret, ErrIptcTagValue
		}
		r := bytes.NewReader(data[0])
		if err = binary.Read(r, defaultEncoding, &ret); err == nil {
			return ret, nil
		}
	case IptcUint32:
		ret := uint32(0)
		if len(data[0]) != 4 {
			return ret, ErrIptcTagValue
		}
		r := bytes.NewReader(data[0])
		if err = binary.Read(r, defaultEncoding, &ret); err == nil {
			return ret, nil
		}
	case IptcUndef:
		return data[0], nil
	default:
		err = ErrIptcUndefinedType
	}
	return nil, err
}
*/

func decodeIptcRecordData(r io.Reader) (IptcRecordTag, []byte, error) {
	var err error
	tuint8 := uint8(0)
	recTag := IptcRecordTag{}

	//marker
	if err = binary.Read(r, defaultEncoding, &tuint8); err != nil {
		return recTag, nil, err
	} else if tuint8 != TagMarker {
		return recTag, nil, fmt.Errorf("Invalid IPTC Tag Marker")
	}
	//record
	if err = binary.Read(r, defaultEncoding, &tuint8); err != nil {
		return recTag, nil, err
	}
	recTag.Record = IptcRecord(tuint8)
	//tag
	if err = binary.Read(r, defaultEncoding, &tuint8); err != nil {
		return recTag, nil, err
	}
	recTag.Tag = IptcTag(tuint8)

	//size
	size16 := uint16(0)
	size32 := uint32(0)
	size := uint64(0) //size can be up to 8 bytes (unrealistic but still
	if err = binary.Read(r, defaultEncoding, &size16); err != nil {
		return recTag, nil, err
	}
	if !hasMsb(size16) {
		size = uint64(size16)
	} else {
		sizeLen := clearMsb(size16)
		if sizeLen == 4 {
			if err = binary.Read(r, defaultEncoding, &size32); err != nil {
				return recTag, nil, err
			}
			size = uint64(size32)
		} else if sizeLen == 8 {
			return recTag, nil, fmt.Errorf("Cannot handle data larger than MaxInt")
		} else {
			return recTag, nil, fmt.Errorf("Unknown data size: %v", sizeLen)
		}
	}
	//read the data:
	if size > maxInt {
		return recTag, nil, fmt.Errorf("Data Size exceeds MaxInt")
	}
	data := make([]byte, size)
	if _, err = io.ReadFull(r, data); err != nil {
		return recTag, nil, err
	}
	return recTag, data, nil
}

func DecodeIptc(r io.Reader) (map[IptcRecordTag]IptcRecordDataset, error) {
	allTags := map[IptcRecordTag][][]byte{}
	var err error

	for err == nil {
		if rt, data, e := decodeIptcRecordData(r); e != nil {
			err = e
		} else {
			if dataSlice, ok := allTags[rt]; ok {
				allTags[rt] = append(dataSlice, data)
			} else {
				allTags[rt] = [][]byte{data}
			}
		}
	}
	ret := map[IptcRecordTag]IptcRecordDataset{}

	for rt, data := range allTags {
		desc, ok := IptcTagDescriptions[rt]
		if !ok {
			fmt.Println("Could not find tag so skipping it: ", rt)
			continue
		}
		ds := IptcRecordDataset{Record: rt.Record, Tag: rt.Tag, Repeatable: desc.Repeatable, Type: desc.Type}
		if desc.Repeatable {
			ds.Data, err = decodeIptcDataToSlice(desc, data)
		} else {
			ds.Data, err = decodeIptcData(desc, data[0])
		}
		if err != nil {
			return ret, err
		}
		ret[rt] = ds
	}
	return ret, nil
}

func encodeIptcRecordData(bw *bufio.Writer, record IptcRecord, tag IptcTag, val interface{}) error {
	var err error
	writeSize := func(length int) error {
		if length < 0 {
			return fmt.Errorf("Negative size")
		}
		if length < msb {
			return binary.Write(bw, defaultEncoding, uint16(length))
		}
		sizeLen := uint16(4)
		sizeLen = setMsb(sizeLen)
		if e := binary.Write(bw, defaultEncoding, sizeLen); err != nil {
			return e
		}
		return binary.Write(bw, defaultEncoding, uint32(length))
	}
	if err = binary.Write(bw, defaultEncoding, TagMarker); err != nil {
		return err
	}
	if err = binary.Write(bw, defaultEncoding, record); err != nil {
		return err
	}
	if err = binary.Write(bw, defaultEncoding, tag); err != nil {
		return err
	}

	switch vtype := val.(type) {
	case string:
		if err = writeSize(len(vtype)); err != nil {
			return err
		}
		_, err = bw.WriteString(vtype)
	case uint8:
		if err = binary.Write(bw, defaultEncoding, uint16(1)); err != nil {
			return err
		}
		err = binary.Write(bw, defaultEncoding, vtype)
	case uint16:
		if err = binary.Write(bw, defaultEncoding, uint16(2)); err != nil {
			return err
		}
		err = binary.Write(bw, defaultEncoding, vtype)
	case uint32:
		if err = binary.Write(bw, defaultEncoding, uint16(4)); err != nil {
			return err
		}
		err = binary.Write(bw, defaultEncoding, vtype)
	case []byte:
		if err = writeSize(len(vtype)); err != nil {
			return err
		}
		_, err = bw.Write(vtype)
	default:
		err = ErrIptcUndefinedType
	}
	return err
}

func encodeIptcRecord(bw *bufio.Writer, rd IptcRecordDataset) error {
	if rd.Repeatable {
		switch vtype := rd.Data.(type) {
		case []string:
			for _, val := range vtype {
				if err := encodeIptcRecordData(bw, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		case []uint8:
			for _, val := range vtype {
				if err := encodeIptcRecordData(bw, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		case []uint16:
			for _, val := range vtype {
				if err := encodeIptcRecordData(bw, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		case []uint32:
			for _, val := range vtype {
				if err := encodeIptcRecordData(bw, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		case [][]byte:
			for _, val := range vtype {
				if err := encodeIptcRecordData(bw, rd.Record, rd.Tag, val); err != nil {
					return err
				}
			}
		default:
			return ErrIptcUndefinedType
		}
		return nil
	}
	return encodeIptcRecordData(bw, rd.Record, rd.Tag, rd.Data)

}

func EncodeIptc(w io.Writer, recs map[IptcRecordTag]IptcRecordDataset) error {
	keys := []IptcRecordTag{}
	for k := range recs {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Record == keys[j].Record {
			return keys[i].Tag < keys[j].Tag
		}
		return keys[i].Record < keys[j].Record
	})
	bw := bufio.NewWriter(w)
	for _, k := range keys {
		if err := encodeIptcRecord(bw, recs[k]); err != nil {
			return err
		}
	}
	return bw.Flush()
	//return nil
}

func ParseIptcJpeg(sl *jpegstructure.SegmentList) (map[IptcRecordTag]IptcRecordDataset, error) {
	ret := map[IptcRecordTag]IptcRecordDataset{}
	_, res, err := photoshop.ParseJpeg(sl)
	if err != nil && err == photoshop.ErrNoPhotoshopBlock {
		return ret, ErrNoIptc
	} else if err != nil {
		return ret, err
	}
	if iptcData, ok := res[photoshop.IptcId]; ok {
		return DecodeIptc(bytes.NewReader(iptcData.Data))
	}
	return ret, ErrNoIptc
}
