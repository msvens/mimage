package photoshop

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	"io"
	"sort"
)

/*
This implementation is based on https://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_pgfId-1037504
*/

const (
	IptcId   uint16 = 0x0404
	DigestId uint16 = 0x0425
)

const photoshopBlockPrefix = "Photoshop 3.0\000"
const photoshopResourceSignature = "8BIM"

var defaultEncoding = binary.BigEndian

var ErrNoPrefix = fmt.Errorf("Block does not contain photoshop prefix")
var ErrNoData = fmt.Errorf("Block contains no data")
var ErrNoPhotoshopBlock = fmt.Errorf("Image contained no photoshop data")

/*
func even(n uint) bool {
	return n%2 == 0
}
*/

func odd(n uint) bool {
	return n%2 == 1
}

type PhotoshopImageResource struct {
	Signature  string //4 bytes. Should be '8BIM'
	ResourceId uint16
	Name       string //pascal encoded, padded to make the size even. (null name consists of two bytes of 0)
	Data       []byte
}

func NewPhotoshopImageResource(resourceId uint16, data []byte) PhotoshopImageResource {
	return PhotoshopImageResource{Signature: photoshopResourceSignature, ResourceId: resourceId, Data: data}
}

func (ir PhotoshopImageResource) String() string {
	return fmt.Sprintf("Signature=[%s] ResourceId=[0x%04x] Name=[%s] DataSize=(%d)", ir.Signature, ir.ResourceId, ir.Name, len(ir.Data))
}

func decodeImageResource(br *bufio.Reader) (PhotoshopImageResource, error) {
	var err error
	ret := PhotoshopImageResource{}

	//signature
	signature := make([]byte, 4)
	if _, err = io.ReadFull(br, signature); err != nil {
		//this is the only place where we expected an io.EOF error so just return the error from the io.ReadFulll
		return ret, err
	}
	ret.Signature = string(signature) //should check that it is correct 8BIM
	if ret.Signature != photoshopResourceSignature {
		return ret, fmt.Errorf("Wrong signature expected %s got %v", photoshopResourceSignature, ret.Signature)
	}

	//resource id
	if err = binary.Read(br, defaultEncoding, &ret.ResourceId); err != nil {
		return ret, fmt.Errorf("Could not read resourceId: %v", err)
	}

	//name (padded to make it even)
	nSize := uint8(0)
	if err = binary.Read(br, defaultEncoding, &nSize); err != nil {
		return ret, fmt.Errorf("Could not name size: %v", err)
	}
	if nSize > 0 {
		name := make([]byte, nSize)
		if _, err = io.ReadFull(br, name); err != nil {
			return ret, fmt.Errorf("Could not name: %v", err)
		}
		ret.Name = string(name)
	}
	if nSize == 0 || odd(uint(nSize+1)) { //need to account for the size byte
		if _, err = br.ReadByte(); err != nil {
			return ret, fmt.Errorf("Could not read name padding byte: %v", err)
		}
	}

	//data
	dSize := uint32(0)
	if err = binary.Read(br, defaultEncoding, &dSize); err != nil {
		return ret, fmt.Errorf("Could not read daa size: %v", err)
	}
	data := make([]byte, dSize)
	if _, err = io.ReadFull(br, data); err != nil {
		return ret, fmt.Errorf("Could not read data: %v", err)
	}
	ret.Data = data

	//padded to make it even)
	if odd(uint(dSize)) {
		if _, err = br.ReadByte(); err != nil {
			return ret, fmt.Errorf("Could not read data padding byte: %v", err)
		}
	}
	return ret, nil
}

func encodeImageResource(bw *bufio.Writer, r PhotoshopImageResource) error {
	//first write signature
	if r.Signature != photoshopResourceSignature {
		return fmt.Errorf("Expected signature %s got %s", photoshopResourceSignature, r.Signature)
	}
	var err error
	//signature
	if _, err = bw.WriteString(r.Signature); err != nil {
		return fmt.Errorf("Could not write signature: %v", err)
	}
	//resourceId
	if err = binary.Write(bw, defaultEncoding, r.ResourceId); err != nil {
		return fmt.Errorf("Could not write resourceId: %v", err)
	}
	//name
	if r.Name == "" {
		b := []byte{0, 0}
		if _, err = bw.Write(b); err != nil {
			return fmt.Errorf("Could not write empty name: %v", err)
		}
	} else if len(r.Name) < 255 {
		size := uint8(len(r.Name))
		if err = binary.Write(bw, defaultEncoding, size); err != nil {
			return fmt.Errorf("Could not write name length: %v", err)
		}
		if _, err = bw.WriteString(r.Name); err != nil {
			return fmt.Errorf("Could not write name: %v", err)
		}
		if odd(uint(size + 1)) { //need to account for size byte
			if err = bw.WriteByte(0); err != nil {
				return fmt.Errorf("Could not write name padding: %v", err)
			}
		}
	}
	//data size:
	dSize := uint32(len(r.Data))
	if err = binary.Write(bw, defaultEncoding, dSize); err != nil {
		return fmt.Errorf("Could not write data size: %v", err)
	}
	if _, err = bw.Write(r.Data); err != nil {
		return fmt.Errorf("Could not write data: %v", err)
	}
	if odd(uint(dSize)) {
		if err = bw.WriteByte(0); err != nil {
			return fmt.Errorf("Could not write data padding byte: %v", err)
		}
	}
	return nil
}

func Decode(r io.Reader, checkPrefix bool) (map[uint16]PhotoshopImageResource, error) {
	ret := map[uint16]PhotoshopImageResource{}
	var err error
	br := bufio.NewReader(r)
	if checkPrefix {
		b := make([]byte, len(photoshopBlockPrefix))
		if _, err = io.ReadFull(br, b); err != nil {
			return ret, ErrNoPrefix
		}
		if photoshopBlockPrefix != string(b) {
			return ret, ErrNoPrefix
		}
	}
	for err == nil {
		if ph, e := decodeImageResource(br); e != nil {
			err = e
		} else {
			ret[ph.ResourceId] = ph
		}
	}
	if err != io.EOF {
		return ret, err
	} else if len(ret) == 0 {
		return ret, ErrNoData
	} else {
		return ret, nil
	}

}

func Encode(w io.Writer, source map[uint16]PhotoshopImageResource, addPrefix bool) error {
	bw := bufio.NewWriter(w)
	var err error

	if len(source) == 0 {
		return fmt.Errorf("No photoshop resources to write")
	}
	if addPrefix {
		if _, err = bw.WriteString(photoshopBlockPrefix); err != nil {
			return err
		}
	}
	//sort resources (likely not needed but nice anyway)
	resources := []PhotoshopImageResource{}
	for _, v := range source {
		resources = append(resources, v)
	}
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].ResourceId < resources[j].ResourceId
	})

	for _, r := range resources {
		if err = encodeImageResource(bw, r); err != nil {
			return err
		}
	}
	if err = bw.Flush(); err != nil {
		return err
	}
	return nil
}

func Marshal(source map[uint16]PhotoshopImageResource, addPrefix bool) ([]byte, error) {
	out := &bytes.Buffer{}
	if err := Encode(out, source, addPrefix); err != nil {
		return nil, err
	} else {
		return out.Bytes(), err
	}
}

func Unmarshal(data []byte, hasPrefix bool, dest *map[uint16]PhotoshopImageResource) error {
	in := bytes.NewReader(data)
	if ret, err := Decode(in, hasPrefix); err != nil {
		return err
	} else {
		*dest = ret
		return nil
	}
}

func ParseJpeg(sl *jpegstructure.SegmentList) (int, map[uint16]PhotoshopImageResource, error) {
	ret := map[uint16]PhotoshopImageResource{}
	for idx, segment := range sl.Segments() {
		//Try parse:
		if err := Unmarshal(segment.Data, true, &ret); err == nil {
			return idx, ret, nil
		} else if err == ErrNoData {
			return idx, ret, ErrNoData
		}
	}
	return -1, ret, ErrNoPhotoshopBlock
}
