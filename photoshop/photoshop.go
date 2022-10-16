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
	//IptcId id for Iptc resource
	IptcId uint16 = 0x0404
	//DigestId id for Digest resource
	DigestId uint16 = 0x0425
)

const photoshopBlockPrefix = "Photoshop 3.0\000"
const photoshopResourceSignature = "8BIM"

var defaultEncoding = binary.BigEndian

// ErrNoPrefix expected photoshop block prefix
var ErrNoPrefix = fmt.Errorf("Block does not contain photoshop prefix")

// ErrNoData no data in block
var ErrNoData = fmt.Errorf("Block contains no data")

// ErrNoPhotoshopBlock could not find a photoshop segment/block in an image
var ErrNoPhotoshopBlock = fmt.Errorf("Image contained no photoshop data")

func odd(n uint) bool {
	return n%2 == 1
}

// ImageResource holds a photoshop image source
type ImageResource struct {
	//Signature is always 4 bytes and should be 8BIM
	Signature string
	//ResourceId. E.g. 0x0404 for Iptc Data
	ResourceId uint16
	//Name of the resource. Most often this ""
	Name string
	//Image resource data
	Data []byte
}

// NewPhotoshopImageResource from an id and data
func NewPhotoshopImageResource(resourceId uint16, data []byte) ImageResource {
	return ImageResource{Signature: photoshopResourceSignature, ResourceId: resourceId, Data: data}
}

func (ir ImageResource) String() string {
	return fmt.Sprintf("Signature=[%s] ResourceId=[0x%04x] Name=[%s] DataSize=(%d)", ir.Signature, ir.ResourceId, ir.Name, len(ir.Data))
}

func decodeImageResource(br *bufio.Reader) (ImageResource, error) {
	var err error
	ret := ImageResource{}

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

func encodeImageResource(bw *bufio.Writer, r ImageResource) error {
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

// Decode photoshop image resources according to https://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_pgfId-1037504
func Decode(r io.Reader, checkPrefix bool) (map[uint16]ImageResource, error) {
	ret := map[uint16]ImageResource{}
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

// Encode according to https://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_pgfId-1037504
func Encode(w io.Writer, source map[uint16]ImageResource, addPrefix bool) error {
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
	resources := []ImageResource{}
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

// Marshal writes photoshop images resources. If addPrefix adds the photoshop
// block prefix ("Photoshop 3.0\000")
func Marshal(source map[uint16]ImageResource, addPrefix bool) ([]byte, error) {
	out := &bytes.Buffer{}
	err := Encode(out, source, addPrefix)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), err

}

// Unmarshal reads photoshop imagesources into dest. If hasPrefix it will first scan
// for a photoshop block prefix ("Photoshop 3.0\000")
func Unmarshal(data []byte, hasPrefix bool, dest *map[uint16]ImageResource) error {
	in := bytes.NewReader(data)
	ret, err := Decode(in, hasPrefix)
	if err != nil {
		return err
	}
	*dest = ret
	return nil
}

// ParseJpeg reads a jpeg segement list and extract the photoshop resources (if they exist)
func ParseJpeg(sl *jpegstructure.SegmentList) (int, map[uint16]ImageResource, error) {
	ret := map[uint16]ImageResource{}
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
