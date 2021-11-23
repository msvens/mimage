package metadata

import (
	"errors"
	"fmt"
)

var ParseImageErr = errors.New("Could not parse image")
var NoExifErr = errors.New("No Exif data")
var NoIptcErr = errors.New("No IPTC data")
var NoXmpErr = errors.New("No XMP data")
var NoXmpModelErr = errors.New("Could not find XMP Model")

var IfdTagNotFoundErr = errors.New("Ifd tag not found")
var IfdUndefinedTypeErr = errors.New("Tag type undefined")
var IptcTagNotFoundErr = errors.New("Iptc tag not found")
var IptcUndefinedTypeErr = errors.New("Could not parse ")

var JpegWrongFileExtErr = fmt.Errorf("Only .jpeg or .jpg file extension allowed")

/*
type TagNotFoundError struct {
	errorStr string
}

func (tgf TagNotFoundError) Error() string {
	return tgf.errorStr
}

func NewTagNotFoundError(tag string) TagNotFoundError {
	return TagNotFoundError{fmt.Sprintf("Tag not found: %s", tag)}
}
*/
