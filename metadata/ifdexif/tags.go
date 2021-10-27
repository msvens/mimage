package ifdexif

import (
	"fmt"
)

var TagNames = map[uint16]string{
	ExposureTime:              "ExposureTime",
	FNumber:                   "FNumber",
	ExposureProgram:           "ExposureProgram",
	SpectralSensitivity:       "SpectralSensitivity",
	ISOSpeedRatings:           "ISOSpeedRatings",
	OECF:                      "OECF",
	SensitivityType:           "SensitivityType",
	StandardOutputSensitivity: "StandardOutputSensitivity",
	RecommendedExposureIndex:  "RecommendedExposureIndex",
	ISOSpeed:                  "ISOSpeed",
	ISOSpeedLatitudeyyy:       "ISOSpeedLatitudeyyy",
	ISOSpeedLatitudezzz:       "ISOSpeedLatitudezzz",
	ExifVersion:               "ExifVersion",
	DateTimeOriginal:          "DateTimeOriginal",
	DateTimeDigitized:         "DateTimeDigitized",
	OffsetTime: "OffsetTime",
	OffsetTimeOriginal: "OffsetTimeOriginal",
	OffsetTimeDigitized: "OffsetTimeDigitized",
	ComponentsConfiguration:   "ComponentsConfiguration",
	CompressedBitsPerPixel:    "CompressedBitsPerPixel",
	ShutterSpeedValue:         "ShutterSpeedValue",
	ApertureValue:             "ApertureValue",
	BrightnessValue:           "BrightnessValue",
	ExposureBiasValue:         "ExposureBiasValue",
	MaxApertureValue:          "MaxApertureValue",
	SubjectDistance:           "SubjectDistance",
	MeteringMode:              "MeteringMode",
	LightSource:               "LightSource",
	Flash:                     "Flash",
	FocalLength:               "FocalLength",
	SubjectArea:               "SubjectArea",
	MakerNote:                 "MakerNote",
	UserComment:               "UserComment",
	SubSecTime:                "SubSecTime",
	SubSecTimeOriginal:        "SubSecTimeOriginal",
	SubSecTimeDigitized:       "SubSecTimeDigitized",
	FlashpixVersion:           "FlashpixVersion",
	ColorSpace:                "ColorSpace",
	PixelXDimension:           "PixelXDimension",
	PixelYDimension:           "PixelYDimension",
	RelatedSoundFile:          "RelatedSoundFile",
	InteroperabilityTag:       "InteroperabilityTag",
	FlashEnergy:               "FlashEnergy",
	SpatialFrequencyResponse:  "SpatialFrequencyResponse",
	FocalPlaneXResolution:     "FocalPlaneXResolution",
	FocalPlaneYResolution:     "FocalPlaneYResolution",
	FocalPlaneResolutionUnit:  "FocalPlaneResolutionUnit",
	SubjectLocation:           "SubjectLocation",
	ExposureIndex:             "ExposureIndex",
	SensingMethod:             "SensingMethod",
	FileSource:                "FileSource",
	SceneType:                 "SceneType",
	CFAPattern:                "CFAPattern",
	CustomRendered:            "CustomRendered",
	ExposureMode:              "ExposureMode",
	WhiteBalance:              "WhiteBalance",
	DigitalZoomRatio:          "DigitalZoomRatio",
	FocalLengthIn35mmFilm:     "FocalLengthIn35mmFilm",
	SceneCaptureType:          "SceneCaptureType",
	GainControl:               "GainControl",
	Contrast:                  "Contrast",
	Saturation:                "Saturation",
	Sharpness:                 "Sharpness",
	DeviceSettingDescription:  "DeviceSettingDescription",
	SubjectDistanceRange:      "SubjectDistanceRange",
	ImageUniqueID:             "ImageUniqueID",
	CameraOwnerName:           "CameraOwnerName",
	BodySerialNumber:          "BodySerialNumber",
	LensSpecification:         "LensSpecification",
	LensMake:                  "LensMake",
	LensModel:                 "LensModel",
	LensSerialNumber:          "LensSerialNumber",
}

// ifdroot/exif tag ids
const (
	ExposureTime              uint16 = 0x829a
	FNumber                   uint16 = 0x829d
	ExposureProgram           uint16 = 0x8822
	SpectralSensitivity       uint16 = 0x8824
	ISOSpeedRatings           uint16 = 0x8827
	OECF                      uint16 = 0x8828
	SensitivityType           uint16 = 0x8830
	StandardOutputSensitivity uint16 = 0x8831
	RecommendedExposureIndex  uint16 = 0x8832
	ISOSpeed                  uint16 = 0x8833
	ISOSpeedLatitudeyyy       uint16 = 0x8834
	ISOSpeedLatitudezzz       uint16 = 0x8835
	ExifVersion               uint16 = 0x9000
	DateTimeOriginal          uint16 = 0x9003
	DateTimeDigitized         uint16 = 0x9004
	OffsetTime uint16 = 0x9010
	OffsetTimeOriginal uint16 = 0x9011
	OffsetTimeDigitized uint16 = 0x9012
	ComponentsConfiguration   uint16 = 0x9101
	CompressedBitsPerPixel    uint16 = 0x9102
	ShutterSpeedValue         uint16 = 0x9201
	ApertureValue             uint16 = 0x9202
	BrightnessValue           uint16 = 0x9203
	ExposureBiasValue         uint16 = 0x9204
	MaxApertureValue          uint16 = 0x9205
	SubjectDistance           uint16 = 0x9206
	MeteringMode              uint16 = 0x9207
	LightSource               uint16 = 0x9208
	Flash                     uint16 = 0x9209
	FocalLength               uint16 = 0x920a
	SubjectArea               uint16 = 0x9214
	MakerNote                 uint16 = 0x927c
	UserComment               uint16 = 0x9286
	SubSecTime                uint16 = 0x9290
	SubSecTimeOriginal        uint16 = 0x9291
	SubSecTimeDigitized       uint16 = 0x9292
	FlashpixVersion           uint16 = 0xa000
	ColorSpace                uint16 = 0xa001
	PixelXDimension           uint16 = 0xa002
	PixelYDimension           uint16 = 0xa003
	RelatedSoundFile          uint16 = 0xa004
	InteroperabilityTag       uint16 = 0xa005
	FlashEnergy               uint16 = 0xa20b
	SpatialFrequencyResponse  uint16 = 0xa20c
	FocalPlaneXResolution     uint16 = 0xa20e
	FocalPlaneYResolution     uint16 = 0xa20f
	FocalPlaneResolutionUnit  uint16 = 0xa210
	SubjectLocation           uint16 = 0xa214
	ExposureIndex             uint16 = 0xa215
	SensingMethod             uint16 = 0xa217
	FileSource                uint16 = 0xa300
	SceneType                 uint16 = 0xa301
	CFAPattern                uint16 = 0xa302
	CustomRendered            uint16 = 0xa401
	ExposureMode              uint16 = 0xa402
	WhiteBalance              uint16 = 0xa403
	DigitalZoomRatio          uint16 = 0xa404
	FocalLengthIn35mmFilm     uint16 = 0xa405
	SceneCaptureType          uint16 = 0xa406
	GainControl               uint16 = 0xa407
	Contrast                  uint16 = 0xa408
	Saturation                uint16 = 0xa409
	Sharpness                 uint16 = 0xa40a
	DeviceSettingDescription  uint16 = 0xa40b
	SubjectDistanceRange      uint16 = 0xa40c
	ImageUniqueID             uint16 = 0xa420
	CameraOwnerName           uint16 = 0xa430
	BodySerialNumber          uint16 = 0xa431
	LensSpecification         uint16 = 0xa432
	LensMake                  uint16 = 0xa433
	LensModel                 uint16 = 0xa434
	LensSerialNumber          uint16 = 0xa435
)

func TagName(tagId uint16) string {
	if tagName, found := TagNames[tagId]; found {
		return tagName
	} else {
		return fmt.Sprintf("Unknown IFD/EXIF TagId: %x", tagId)
	}

}
