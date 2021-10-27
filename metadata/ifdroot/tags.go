package ifdroot

import "fmt"

var TagNames = map[uint16]string{
	ProcessingSoftware:          "ProcessingSoftware",
	NewSubfileType:              "NewSubfileType",
	SubfileType:                 "SubfileType",
	ImageWidth:                  "ImageWidth",
	ImageLength:                 "ImageLength",
	BitsPerSample:               "BitsPerSample",
	Compression:                 "Compression",
	PhotometricInterpretation:   "PhotometricInterpretation",
	Thresholding:                "Thresholding",
	CellWidth:                   "CellWidth",
	CellLength:                  "CellLength",
	FillOrder:                   "FillOrder",
	DocumentName:                "DocumentName",
	ImageDescription:            "ImageDescription",
	Make:                        "Make",
	Model:                       "Model",
	StripOffsets:                "StripOffsets",
	Orientation:                 "Orientation",
	SamplesPerPixel:             "SamplesPerPixel",
	RowsPerStrip:                "RowsPerStrip",
	StripByteCounts:             "StripByteCounts",
	XResolution:                 "XResolution",
	YResolution:                 "YResolution",
	PlanarConfiguration:         "PlanarConfiguration",
	GrayResponseUnit:            "GrayResponseUnit",
	GrayResponseCurve:           "GrayResponseCurve",
	T4Options:                   "T4Options",
	T6Options:                   "T6Options",
	ResolutionUnit:              "ResolutionUnit",
	PageNumber:                  "PageNumber",
	TransferFunction:            "TransferFunction",
	Software:                    "Software",
	DateTime:                    "DateTime",
	Artist:                      "Artist",
	HostComputer:                "HostComputer",
	Predictor:                   "Predictor",
	WhitePoint:                  "WhitePoint",
	PrimaryChromaticities:       "PrimaryChromaticities",
	ColorMap:                    "ColorMap",
	HalftoneHints:               "HalftoneHints",
	TileWidth:                   "TileWidth",
	TileLength:                  "TileLength",
	TileOffsets:                 "TileOffsets",
	TileByteCounts:              "TileByteCounts",
	SubIFDs:                     "SubIFDs",
	InkSet:                      "InkSet",
	InkNames:                    "InkNames",
	NumberOfInks:                "NumberOfInks",
	DotRange:                    "DotRange",
	TargetPrinter:               "TargetPrinter",
	ExtraSamples:                "ExtraSamples",
	SampleFormat:                "SampleFormat",
	SMinSampleValue:             "SMinSampleValue",
	SMaxSampleValue:             "SMaxSampleValue",
	TransferRange:               "TransferRange",
	ClipPath:                    "ClipPath",
	XClipPathUnits:              "XClipPathUnits",
	YClipPathUnits:              "YClipPathUnits",
	Indexed:                     "Indexed",
	JPEGTables:                  "JPEGTables",
	OPIProxy:                    "OPIProxy",
	JPEGProc:                    "JPEGProc",
	JPEGInterchangeFormat:       "JPEGInterchangeFormat",
	JPEGInterchangeFormatLength: "JPEGInterchangeFormatLength",
	JPEGRestartInterval:         "JPEGRestartInterval",
	JPEGLosslessPredictors:      "JPEGLosslessPredictors",
	JPEGPointTransforms:         "JPEGPointTransforms",
	JPEGQTables:                 "JPEGQTables",
	JPEGDCTables:                "JPEGDCTables",
	JPEGACTables:                "JPEGACTables",
	YCbCrCoefficients:           "YCbCrCoefficients",
	YCbCrSubSampling:            "YCbCrSubSampling",
	YCbCrPositioning:            "YCbCrPositioning",
	ReferenceBlackWhite:         "ReferenceBlackWhite",
	XMLPacket:                   "XMLPacket",
	Rating:                      "Rating",
	RatingPercent:               "RatingPercent",
	ImageID:                     "ImageID",
	CFARepeatPatternDim:         "CFARepeatPatternDim",
	CFAPattern:                  "CFAPattern",
	BatteryLevel:                "BatteryLevel",
	Copyright:                   "Copyright",
	ExposureTime:                "ExposureTime",
	FNumber:                     "FNumber",
	IPTCNAA:                     "IPTCNAA",
	ImageResources:              "ImageResources",
	ExifTag:                     "ExifTag",
	InterColorProfile:           "InterColorProfile",
	ExposureProgram:             "ExposureProgram",
	SpectralSensitivity:         "SpectralSensitivity",
	GPSTag:                      "GPSTag",
	ISOSpeedRatings:             "ISOSpeedRatings",
	OECF:                        "OECF",
	Interlace:                   "Interlace",
	SensitivityType:             "SensitivityType",
	TimeZoneOffset:              "TimeZoneOffset",
	SelfTimerMode:               "SelfTimerMode",
	RecommendedExposureIndex:    "RecommendedExposureIndex",
	DateTimeOriginal:            "DateTimeOriginal",
	DateTimeDigitized:           "DateTimeDigitized",
	CompressedBitsPerPixel:      "CompressedBitsPerPixel",
	ShutterSpeedValue:           "ShutterSpeedValue",
	ApertureValue:               "ApertureValue",
	BrightnessValue:             "BrightnessValue",
	ExposureBiasValue:           "ExposureBiasValue",
	MaxApertureValue:            "MaxApertureValue",
	SubjectDistance:             "SubjectDistance",
	MeteringMode:                "MeteringMode",
	LightSource:                 "LightSource",
	Flash:                       "Flash",
	FocalLength:                 "FocalLength",
	FlashEnergy:                 "FlashEnergy",
	SpatialFrequencyResponse:    "SpatialFrequencyResponse",
	Noise:                       "Noise",
	FocalPlaneXResolution:       "FocalPlaneXResolution",
	FocalPlaneYResolution:       "FocalPlaneYResolution",
	FocalPlaneResolutionUnit:    "FocalPlaneResolutionUnit",
	ImageNumber:                 "ImageNumber",
	SecurityClassification:      "SecurityClassification",
	ImageHistory:                "ImageHistory",
	SubjectLocation:             "SubjectLocation",
	ExposureIndex:               "ExposureIndex",
	TIFFEPStandardID:            "TIFFEPStandardID",
	SensingMethod:               "SensingMethod",
	XPTitle:                     "XPTitle",
	XPComment:                   "XPComment",
	XPAuthor:                    "XPAuthor",
	XPKeywords:                  "XPKeywords",
	XPSubject:                   "XPSubject",
	PrintImageMatching:          "PrintImageMatching",
	DNGVersion:                  "DNGVersion",
	DNGBackwardVersion:          "DNGBackwardVersion",
	UniqueCameraModel:           "UniqueCameraModel",
	LocalizedCameraModel:        "LocalizedCameraModel",
	CFAPlaneColor:               "CFAPlaneColor",
	CFALayout:                   "CFALayout",
	LinearizationTable:          "LinearizationTable",
	BlackLevelRepeatDim:         "BlackLevelRepeatDim",
	BlackLevel:                  "BlackLevel",
	BlackLevelDeltaH:            "BlackLevelDeltaH",
	BlackLevelDeltaV:            "BlackLevelDeltaV",
	WhiteLevel:                  "WhiteLevel",
	DefaultScale:                "DefaultScale",
	DefaultCropOrigin:           "DefaultCropOrigin",
	DefaultCropSize:             "DefaultCropSize",
	ColorMatrix1:                "ColorMatrix1",
	ColorMatrix2:                "ColorMatrix2",
	CameraCalibration1:          "CameraCalibration1",
	CameraCalibration2:          "CameraCalibration2",
	ReductionMatrix1:            "ReductionMatrix1",
	ReductionMatrix2:            "ReductionMatrix2",
	AnalogBalance:               "AnalogBalance",
	AsShotNeutral:               "AsShotNeutral",
	AsShotWhiteXY:               "AsShotWhiteXY",
	BaselineExposure:            "BaselineExposure",
	BaselineNoise:               "BaselineNoise",
	BaselineSharpness:           "BaselineSharpness",
	BayerGreenSplit:             "BayerGreenSplit",
	LinearResponseLimit:         "LinearResponseLimit",
	CameraSerialNumber:          "CameraSerialNumber",
	LensInfo:                    "LensInfo",
	ChromaBlurRadius:            "ChromaBlurRadius",
	AntiAliasStrength:           "AntiAliasStrength",
	ShadowScale:                 "ShadowScale",
	DNGPrivateData:              "DNGPrivateData",
	MakerNoteSafety:             "MakerNoteSafety",
	CalibrationIlluminant1:      "CalibrationIlluminant1",
	CalibrationIlluminant2:      "CalibrationIlluminant2",
	BestQualityScale:            "BestQualityScale",
	RawDataUniqueID:             "RawDataUniqueID",
	OriginalRawFileName:         "OriginalRawFileName",
	OriginalRawFileData:         "OriginalRawFileData",
	ActiveArea:                  "ActiveArea",
	MaskedAreas:                 "MaskedAreas",
	AsShotICCProfile:            "AsShotICCProfile",
	AsShotPreProfileMatrix:      "AsShotPreProfileMatrix",
	CurrentICCProfile:           "CurrentICCProfile",
	CurrentPreProfileMatrix:     "CurrentPreProfileMatrix",
	ColorimetricReference:       "ColorimetricReference",
	CameraCalibrationSignature:  "CameraCalibrationSignature",
	ProfileCalibrationSignature: "ProfileCalibrationSignature",
	AsShotProfileName:           "AsShotProfileName",
	NoiseReductionApplied:       "NoiseReductionApplied",
	ProfileName:                 "ProfileName",
	ProfileHueSatMapDims:        "ProfileHueSatMapDims",
	ProfileHueSatMapData1:       "ProfileHueSatMapData1",
	ProfileHueSatMapData2:       "ProfileHueSatMapData2",
	ProfileToneCurve:            "ProfileToneCurve",
	ProfileEmbedPolicy:          "ProfileEmbedPolicy",
	ProfileCopyright:            "ProfileCopyright",
	ForwardMatrix1:              "ForwardMatrix1",
	ForwardMatrix2:              "ForwardMatrix2",
	PreviewApplicationName:      "PreviewApplicationName",
	PreviewApplicationVersion:   "PreviewApplicationVersion",
	PreviewSettingsName:         "PreviewSettingsName",
	PreviewSettingsDigest:       "PreviewSettingsDigest",
	PreviewColorSpace:           "PreviewColorSpace",
	PreviewDateTime:             "PreviewDateTime",
	RawImageDigest:              "RawImageDigest",
	OriginalRawFileDigest:       "OriginalRawFileDigest",
	SubTileBlockSize:            "SubTileBlockSize",
	RowInterleaveFactor:         "RowInterleaveFactor",
	ProfileLookTableDims:        "ProfileLookTableDims",
	ProfileLookTableData:        "ProfileLookTableData",
	OpcodeList1:                 "OpcodeList1",
	OpcodeList2:                 "OpcodeList2",
	OpcodeList3:                 "OpcodeList3",
	NoiseProfile:                "NoiseProfile",
	CacheVersion:                "CacheVersion",
}

//root ifdroot tag ids
const (
	ProcessingSoftware          uint16 = 0x000b
	NewSubfileType              uint16 = 0x00fe
	SubfileType                 uint16 = 0x00ff
	ImageWidth                  uint16 = 0x0100
	ImageLength                 uint16 = 0x0101
	BitsPerSample               uint16 = 0x0102
	Compression                 uint16 = 0x0103
	PhotometricInterpretation   uint16 = 0x0106
	Thresholding                uint16 = 0x0107
	CellWidth                   uint16 = 0x0108
	CellLength                  uint16 = 0x0109
	FillOrder                   uint16 = 0x010a
	DocumentName                uint16 = 0x010d
	ImageDescription            uint16 = 0x010e
	Make                        uint16 = 0x010f
	Model                       uint16 = 0x0110
	StripOffsets                uint16 = 0x0111
	Orientation                 uint16 = 0x0112
	SamplesPerPixel             uint16 = 0x0115
	RowsPerStrip                uint16 = 0x0116
	StripByteCounts             uint16 = 0x0117
	XResolution                 uint16 = 0x011a
	YResolution                 uint16 = 0x011b
	PlanarConfiguration         uint16 = 0x011c
	GrayResponseUnit            uint16 = 0x0122
	GrayResponseCurve           uint16 = 0x0123
	T4Options                   uint16 = 0x0124
	T6Options                   uint16 = 0x0125
	ResolutionUnit              uint16 = 0x0128
	PageNumber                  uint16 = 0x0129
	TransferFunction            uint16 = 0x012d
	Software                    uint16 = 0x0131
	DateTime                    uint16 = 0x0132
	Artist                      uint16 = 0x013b
	HostComputer                uint16 = 0x013c
	Predictor                   uint16 = 0x013d
	WhitePoint                  uint16 = 0x013e
	PrimaryChromaticities       uint16 = 0x013f
	ColorMap                    uint16 = 0x0140
	HalftoneHints               uint16 = 0x0141
	TileWidth                   uint16 = 0x0142
	TileLength                  uint16 = 0x0143
	TileOffsets                 uint16 = 0x0144
	TileByteCounts              uint16 = 0x0145
	SubIFDs                     uint16 = 0x014a
	InkSet                      uint16 = 0x014c
	InkNames                    uint16 = 0x014d
	NumberOfInks                uint16 = 0x014e
	DotRange                    uint16 = 0x0150
	TargetPrinter               uint16 = 0x0151
	ExtraSamples                uint16 = 0x0152
	SampleFormat                uint16 = 0x0153
	SMinSampleValue             uint16 = 0x0154
	SMaxSampleValue             uint16 = 0x0155
	TransferRange               uint16 = 0x0156
	ClipPath                    uint16 = 0x0157
	XClipPathUnits              uint16 = 0x0158
	YClipPathUnits              uint16 = 0x0159
	Indexed                     uint16 = 0x015a
	JPEGTables                  uint16 = 0x015b
	OPIProxy                    uint16 = 0x015f
	JPEGProc                    uint16 = 0x0200
	JPEGInterchangeFormat       uint16 = 0x0201
	JPEGInterchangeFormatLength uint16 = 0x0202
	JPEGRestartInterval         uint16 = 0x0203
	JPEGLosslessPredictors      uint16 = 0x0205
	JPEGPointTransforms         uint16 = 0x0206
	JPEGQTables                 uint16 = 0x0207
	JPEGDCTables                uint16 = 0x0208
	JPEGACTables                uint16 = 0x0209
	YCbCrCoefficients           uint16 = 0x0211
	YCbCrSubSampling            uint16 = 0x0212
	YCbCrPositioning            uint16 = 0x0213
	ReferenceBlackWhite         uint16 = 0x0214
	XMLPacket                   uint16 = 0x02bc
	Rating                      uint16 = 0x4746
	RatingPercent               uint16 = 0x4749
	ImageID                     uint16 = 0x800d
	CFARepeatPatternDim         uint16 = 0x828d
	CFAPattern                  uint16 = 0x828e
	BatteryLevel                uint16 = 0x828f
	Copyright                   uint16 = 0x8298
	ExposureTime                uint16 = 0x829a
	FNumber                     uint16 = 0x829d
	IPTCNAA                     uint16 = 0x83bb
	ImageResources              uint16 = 0x8649
	ExifTag                     uint16 = 0x8769
	InterColorProfile           uint16 = 0x8773
	ExposureProgram             uint16 = 0x8822
	SpectralSensitivity         uint16 = 0x8824
	GPSTag                      uint16 = 0x8825
	ISOSpeedRatings             uint16 = 0x8827
	OECF                        uint16 = 0x8828
	Interlace                   uint16 = 0x8829
	SensitivityType             uint16 = 0x8830
	TimeZoneOffset              uint16 = 0x882a
	SelfTimerMode               uint16 = 0x882b
	RecommendedExposureIndex    uint16 = 0x8832
	DateTimeOriginal            uint16 = 0x9003
	DateTimeDigitized           uint16 = 0x9004
	CompressedBitsPerPixel      uint16 = 0x9102
	ShutterSpeedValue           uint16 = 0x9201
	ApertureValue               uint16 = 0x9202
	BrightnessValue             uint16 = 0x9203
	ExposureBiasValue           uint16 = 0x9204
	MaxApertureValue            uint16 = 0x9205
	SubjectDistance             uint16 = 0x9206
	MeteringMode                uint16 = 0x9207
	LightSource                 uint16 = 0x9208
	Flash                       uint16 = 0x9209
	FocalLength                 uint16 = 0x920a
	FlashEnergy                 uint16 = 0x920b
	SpatialFrequencyResponse    uint16 = 0x920c
	Noise                       uint16 = 0x920d
	FocalPlaneXResolution       uint16 = 0x920e
	FocalPlaneYResolution       uint16 = 0x920f
	FocalPlaneResolutionUnit    uint16 = 0x9210
	ImageNumber                 uint16 = 0x9211
	SecurityClassification      uint16 = 0x9212
	ImageHistory                uint16 = 0x9213
	SubjectLocation             uint16 = 0x9214
	ExposureIndex               uint16 = 0x9215
	TIFFEPStandardID            uint16 = 0x9216
	SensingMethod               uint16 = 0x9217
	XPTitle                     uint16 = 0x9c9b
	XPComment                   uint16 = 0x9c9c
	XPAuthor                    uint16 = 0x9c9d
	XPKeywords                  uint16 = 0x9c9e
	XPSubject                   uint16 = 0x9c9f
	PrintImageMatching          uint16 = 0xc4a5
	DNGVersion                  uint16 = 0xc612
	DNGBackwardVersion          uint16 = 0xc613
	UniqueCameraModel           uint16 = 0xc614
	LocalizedCameraModel        uint16 = 0xc615
	CFAPlaneColor               uint16 = 0xc616
	CFALayout                   uint16 = 0xc617
	LinearizationTable          uint16 = 0xc618
	BlackLevelRepeatDim         uint16 = 0xc619
	BlackLevel                  uint16 = 0xc61a
	BlackLevelDeltaH            uint16 = 0xc61b
	BlackLevelDeltaV            uint16 = 0xc61c
	WhiteLevel                  uint16 = 0xc61d
	DefaultScale                uint16 = 0xc61e
	DefaultCropOrigin           uint16 = 0xc61f
	DefaultCropSize             uint16 = 0xc620
	ColorMatrix1                uint16 = 0xc621
	ColorMatrix2                uint16 = 0xc622
	CameraCalibration1          uint16 = 0xc623
	CameraCalibration2          uint16 = 0xc624
	ReductionMatrix1            uint16 = 0xc625
	ReductionMatrix2            uint16 = 0xc626
	AnalogBalance               uint16 = 0xc627
	AsShotNeutral               uint16 = 0xc628
	AsShotWhiteXY               uint16 = 0xc629
	BaselineExposure            uint16 = 0xc62a
	BaselineNoise               uint16 = 0xc62b
	BaselineSharpness           uint16 = 0xc62c
	BayerGreenSplit             uint16 = 0xc62d
	LinearResponseLimit         uint16 = 0xc62e
	CameraSerialNumber          uint16 = 0xc62f
	LensInfo                    uint16 = 0xc630
	ChromaBlurRadius            uint16 = 0xc631
	AntiAliasStrength           uint16 = 0xc632
	ShadowScale                 uint16 = 0xc633
	DNGPrivateData              uint16 = 0xc634
	MakerNoteSafety             uint16 = 0xc635
	CalibrationIlluminant1      uint16 = 0xc65a
	CalibrationIlluminant2      uint16 = 0xc65b
	BestQualityScale            uint16 = 0xc65c
	RawDataUniqueID             uint16 = 0xc65d
	OriginalRawFileName         uint16 = 0xc68b
	OriginalRawFileData         uint16 = 0xc68c
	ActiveArea                  uint16 = 0xc68d
	MaskedAreas                 uint16 = 0xc68e
	AsShotICCProfile            uint16 = 0xc68f
	AsShotPreProfileMatrix      uint16 = 0xc690
	CurrentICCProfile           uint16 = 0xc691
	CurrentPreProfileMatrix     uint16 = 0xc692
	ColorimetricReference       uint16 = 0xc6bf
	CameraCalibrationSignature  uint16 = 0xc6f3
	ProfileCalibrationSignature uint16 = 0xc6f4
	AsShotProfileName           uint16 = 0xc6f6
	NoiseReductionApplied       uint16 = 0xc6f7
	ProfileName                 uint16 = 0xc6f8
	ProfileHueSatMapDims        uint16 = 0xc6f9
	ProfileHueSatMapData1       uint16 = 0xc6fa
	ProfileHueSatMapData2       uint16 = 0xc6fb
	ProfileToneCurve            uint16 = 0xc6fc
	ProfileEmbedPolicy          uint16 = 0xc6fd
	ProfileCopyright            uint16 = 0xc6fe
	ForwardMatrix1              uint16 = 0xc714
	ForwardMatrix2              uint16 = 0xc715
	PreviewApplicationName      uint16 = 0xc716
	PreviewApplicationVersion   uint16 = 0xc717
	PreviewSettingsName         uint16 = 0xc718
	PreviewSettingsDigest       uint16 = 0xc719
	PreviewColorSpace           uint16 = 0xc71a
	PreviewDateTime             uint16 = 0xc71b
	RawImageDigest              uint16 = 0xc71c
	OriginalRawFileDigest       uint16 = 0xc71d
	SubTileBlockSize            uint16 = 0xc71e
	RowInterleaveFactor         uint16 = 0xc71f
	ProfileLookTableDims        uint16 = 0xc725
	ProfileLookTableData        uint16 = 0xc726
	OpcodeList1                 uint16 = 0xc740
	OpcodeList2                 uint16 = 0xc741
	OpcodeList3                 uint16 = 0xc74e
	NoiseProfile                uint16 = 0xc761
	CacheVersion                uint16 = 0xc7aa
)

func TagName(tagId uint16) string {
	if tagName, found := TagNames[tagId]; found {
		return tagName
	} else {
		return fmt.Sprintf("Unknown IFD Root TagId: %x", tagId)
	}

}
