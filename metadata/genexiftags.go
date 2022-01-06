package metadata

//Do not edit! This is an automatically generated file (see generator.GenerateExifTags()).
//This file was generated based on https://www.exiv2.org/tags.html

//IFD Tag Ids
const (
	IFD_ProcessingSoftware            uint16 = 0x000b
	IFD_NewSubfileType                uint16 = 0x00fe
	IFD_SubfileType                   uint16 = 0x00ff
	IFD_ImageWidth                    uint16 = 0x0100
	IFD_ImageLength                   uint16 = 0x0101
	IFD_BitsPerSample                 uint16 = 0x0102
	IFD_Compression                   uint16 = 0x0103
	IFD_PhotometricInterpretation     uint16 = 0x0106
	IFD_Thresholding                  uint16 = 0x0107
	IFD_CellWidth                     uint16 = 0x0108
	IFD_CellLength                    uint16 = 0x0109
	IFD_FillOrder                     uint16 = 0x010a
	IFD_DocumentName                  uint16 = 0x010d
	IFD_ImageDescription              uint16 = 0x010e
	IFD_Make                          uint16 = 0x010f
	IFD_Model                         uint16 = 0x0110
	IFD_StripOffsets                  uint16 = 0x0111
	IFD_Orientation                   uint16 = 0x0112
	IFD_SamplesPerPixel               uint16 = 0x0115
	IFD_RowsPerStrip                  uint16 = 0x0116
	IFD_StripByteCounts               uint16 = 0x0117
	IFD_XResolution                   uint16 = 0x011a
	IFD_YResolution                   uint16 = 0x011b
	IFD_PlanarConfiguration           uint16 = 0x011c
	IFD_GrayResponseUnit              uint16 = 0x0122
	IFD_GrayResponseCurve             uint16 = 0x0123
	IFD_T4Options                     uint16 = 0x0124
	IFD_T6Options                     uint16 = 0x0125
	IFD_ResolutionUnit                uint16 = 0x0128
	IFD_PageNumber                    uint16 = 0x0129
	IFD_TransferFunction              uint16 = 0x012d
	IFD_Software                      uint16 = 0x0131
	IFD_DateTime                      uint16 = 0x0132
	IFD_Artist                        uint16 = 0x013b
	IFD_HostComputer                  uint16 = 0x013c
	IFD_Predictor                     uint16 = 0x013d
	IFD_WhitePoint                    uint16 = 0x013e
	IFD_PrimaryChromaticities         uint16 = 0x013f
	IFD_ColorMap                      uint16 = 0x0140
	IFD_HalftoneHints                 uint16 = 0x0141
	IFD_TileWidth                     uint16 = 0x0142
	IFD_TileLength                    uint16 = 0x0143
	IFD_TileOffsets                   uint16 = 0x0144
	IFD_TileByteCounts                uint16 = 0x0145
	IFD_SubIFDs                       uint16 = 0x014a
	IFD_InkSet                        uint16 = 0x014c
	IFD_InkNames                      uint16 = 0x014d
	IFD_NumberOfInks                  uint16 = 0x014e
	IFD_DotRange                      uint16 = 0x0150
	IFD_TargetPrinter                 uint16 = 0x0151
	IFD_ExtraSamples                  uint16 = 0x0152
	IFD_SampleFormat                  uint16 = 0x0153
	IFD_SMinSampleValue               uint16 = 0x0154
	IFD_SMaxSampleValue               uint16 = 0x0155
	IFD_TransferRange                 uint16 = 0x0156
	IFD_ClipPath                      uint16 = 0x0157
	IFD_XClipPathUnits                uint16 = 0x0158
	IFD_YClipPathUnits                uint16 = 0x0159
	IFD_Indexed                       uint16 = 0x015a
	IFD_JPEGTables                    uint16 = 0x015b
	IFD_OPIProxy                      uint16 = 0x015f
	IFD_JPEGProc                      uint16 = 0x0200
	IFD_JPEGInterchangeFormat         uint16 = 0x0201
	IFD_JPEGInterchangeFormatLength   uint16 = 0x0202
	IFD_JPEGRestartInterval           uint16 = 0x0203
	IFD_JPEGLosslessPredictors        uint16 = 0x0205
	IFD_JPEGPointTransforms           uint16 = 0x0206
	IFD_JPEGQTables                   uint16 = 0x0207
	IFD_JPEGDCTables                  uint16 = 0x0208
	IFD_JPEGACTables                  uint16 = 0x0209
	IFD_YCbCrCoefficients             uint16 = 0x0211
	IFD_YCbCrSubSampling              uint16 = 0x0212
	IFD_YCbCrPositioning              uint16 = 0x0213
	IFD_ReferenceBlackWhite           uint16 = 0x0214
	IFD_XMLPacket                     uint16 = 0x02bc
	IFD_Rating                        uint16 = 0x4746
	IFD_RatingPercent                 uint16 = 0x4749
	IFD_VignettingCorrParams          uint16 = 0x7032
	IFD_ChromaticAberrationCorrParams uint16 = 0x7035
	IFD_DistortionCorrParams          uint16 = 0x7037
	IFD_ImageID                       uint16 = 0x800d
	IFD_CFARepeatPatternDim           uint16 = 0x828d
	IFD_CFAPattern                    uint16 = 0x828e
	IFD_BatteryLevel                  uint16 = 0x828f
	IFD_Copyright                     uint16 = 0x8298
	IFD_ExposureTime                  uint16 = 0x829a
	IFD_FNumber                       uint16 = 0x829d
	IFD_IPTCNAA                       uint16 = 0x83bb
	IFD_ImageResources                uint16 = 0x8649
	IFD_ExifTag                       uint16 = 0x8769
	IFD_InterColorProfile             uint16 = 0x8773
	IFD_ExposureProgram               uint16 = 0x8822
	IFD_SpectralSensitivity           uint16 = 0x8824
	IFD_GPSTag                        uint16 = 0x8825
	IFD_ISOSpeedRatings               uint16 = 0x8827
	IFD_OECF                          uint16 = 0x8828
	IFD_Interlace                     uint16 = 0x8829
	IFD_TimeZoneOffset                uint16 = 0x882a
	IFD_SelfTimerMode                 uint16 = 0x882b
	IFD_DateTimeOriginal              uint16 = 0x9003
	IFD_CompressedBitsPerPixel        uint16 = 0x9102
	IFD_ShutterSpeedValue             uint16 = 0x9201
	IFD_ApertureValue                 uint16 = 0x9202
	IFD_BrightnessValue               uint16 = 0x9203
	IFD_ExposureBiasValue             uint16 = 0x9204
	IFD_MaxApertureValue              uint16 = 0x9205
	IFD_SubjectDistance               uint16 = 0x9206
	IFD_MeteringMode                  uint16 = 0x9207
	IFD_LightSource                   uint16 = 0x9208
	IFD_Flash                         uint16 = 0x9209
	IFD_FocalLength                   uint16 = 0x920a
	IFD_FlashEnergy                   uint16 = 0x920b
	IFD_SpatialFrequencyResponse      uint16 = 0x920c
	IFD_Noise                         uint16 = 0x920d
	IFD_FocalPlaneXResolution         uint16 = 0x920e
	IFD_FocalPlaneYResolution         uint16 = 0x920f
	IFD_FocalPlaneResolutionUnit      uint16 = 0x9210
	IFD_ImageNumber                   uint16 = 0x9211
	IFD_SecurityClassification        uint16 = 0x9212
	IFD_ImageHistory                  uint16 = 0x9213
	IFD_SubjectLocation               uint16 = 0x9214
	IFD_ExposureIndex                 uint16 = 0x9215
	IFD_TIFFEPStandardID              uint16 = 0x9216
	IFD_SensingMethod                 uint16 = 0x9217
	IFD_XPTitle                       uint16 = 0x9c9b
	IFD_XPComment                     uint16 = 0x9c9c
	IFD_XPAuthor                      uint16 = 0x9c9d
	IFD_XPKeywords                    uint16 = 0x9c9e
	IFD_XPSubject                     uint16 = 0x9c9f
	IFD_PrintImageMatching            uint16 = 0xc4a5
	IFD_DNGVersion                    uint16 = 0xc612
	IFD_DNGBackwardVersion            uint16 = 0xc613
	IFD_UniqueCameraModel             uint16 = 0xc614
	IFD_LocalizedCameraModel          uint16 = 0xc615
	IFD_CFAPlaneColor                 uint16 = 0xc616
	IFD_CFALayout                     uint16 = 0xc617
	IFD_LinearizationTable            uint16 = 0xc618
	IFD_BlackLevelRepeatDim           uint16 = 0xc619
	IFD_BlackLevel                    uint16 = 0xc61a
	IFD_BlackLevelDeltaH              uint16 = 0xc61b
	IFD_BlackLevelDeltaV              uint16 = 0xc61c
	IFD_WhiteLevel                    uint16 = 0xc61d
	IFD_DefaultScale                  uint16 = 0xc61e
	IFD_DefaultCropOrigin             uint16 = 0xc61f
	IFD_DefaultCropSize               uint16 = 0xc620
	IFD_ColorMatrix1                  uint16 = 0xc621
	IFD_ColorMatrix2                  uint16 = 0xc622
	IFD_CameraCalibration1            uint16 = 0xc623
	IFD_CameraCalibration2            uint16 = 0xc624
	IFD_ReductionMatrix1              uint16 = 0xc625
	IFD_ReductionMatrix2              uint16 = 0xc626
	IFD_AnalogBalance                 uint16 = 0xc627
	IFD_AsShotNeutral                 uint16 = 0xc628
	IFD_AsShotWhiteXY                 uint16 = 0xc629
	IFD_BaselineExposure              uint16 = 0xc62a
	IFD_BaselineNoise                 uint16 = 0xc62b
	IFD_BaselineSharpness             uint16 = 0xc62c
	IFD_BayerGreenSplit               uint16 = 0xc62d
	IFD_LinearResponseLimit           uint16 = 0xc62e
	IFD_CameraSerialNumber            uint16 = 0xc62f
	IFD_LensInfo                      uint16 = 0xc630
	IFD_ChromaBlurRadius              uint16 = 0xc631
	IFD_AntiAliasStrength             uint16 = 0xc632
	IFD_ShadowScale                   uint16 = 0xc633
	IFD_DNGPrivateData                uint16 = 0xc634
	IFD_MakerNoteSafety               uint16 = 0xc635
	IFD_CalibrationIlluminant1        uint16 = 0xc65a
	IFD_CalibrationIlluminant2        uint16 = 0xc65b
	IFD_BestQualityScale              uint16 = 0xc65c
	IFD_RawDataUniqueID               uint16 = 0xc65d
	IFD_OriginalRawFileName           uint16 = 0xc68b
	IFD_OriginalRawFileData           uint16 = 0xc68c
	IFD_ActiveArea                    uint16 = 0xc68d
	IFD_MaskedAreas                   uint16 = 0xc68e
	IFD_AsShotICCProfile              uint16 = 0xc68f
	IFD_AsShotPreProfileMatrix        uint16 = 0xc690
	IFD_CurrentICCProfile             uint16 = 0xc691
	IFD_CurrentPreProfileMatrix       uint16 = 0xc692
	IFD_ColorimetricReference         uint16 = 0xc6bf
	IFD_CameraCalibrationSignature    uint16 = 0xc6f3
	IFD_ProfileCalibrationSignature   uint16 = 0xc6f4
	IFD_ExtraCameraProfiles           uint16 = 0xc6f5
	IFD_AsShotProfileName             uint16 = 0xc6f6
	IFD_NoiseReductionApplied         uint16 = 0xc6f7
	IFD_ProfileName                   uint16 = 0xc6f8
	IFD_ProfileHueSatMapDims          uint16 = 0xc6f9
	IFD_ProfileHueSatMapData1         uint16 = 0xc6fa
	IFD_ProfileHueSatMapData2         uint16 = 0xc6fb
	IFD_ProfileToneCurve              uint16 = 0xc6fc
	IFD_ProfileEmbedPolicy            uint16 = 0xc6fd
	IFD_ProfileCopyright              uint16 = 0xc6fe
	IFD_ForwardMatrix1                uint16 = 0xc714
	IFD_ForwardMatrix2                uint16 = 0xc715
	IFD_PreviewApplicationName        uint16 = 0xc716
	IFD_PreviewApplicationVersion     uint16 = 0xc717
	IFD_PreviewSettingsName           uint16 = 0xc718
	IFD_PreviewSettingsDigest         uint16 = 0xc719
	IFD_PreviewColorSpace             uint16 = 0xc71a
	IFD_PreviewDateTime               uint16 = 0xc71b
	IFD_RawImageDigest                uint16 = 0xc71c
	IFD_OriginalRawFileDigest         uint16 = 0xc71d
	IFD_SubTileBlockSize              uint16 = 0xc71e
	IFD_RowInterleaveFactor           uint16 = 0xc71f
	IFD_ProfileLookTableDims          uint16 = 0xc725
	IFD_ProfileLookTableData          uint16 = 0xc726
	IFD_OpcodeList1                   uint16 = 0xc740
	IFD_OpcodeList2                   uint16 = 0xc741
	IFD_OpcodeList3                   uint16 = 0xc74e
	IFD_NoiseProfile                  uint16 = 0xc761
	IFD_TimeCodes                     uint16 = 0xc763
	IFD_FrameRate                     uint16 = 0xc764
	IFD_TStop                         uint16 = 0xc772
	IFD_ReelName                      uint16 = 0xc789
	IFD_CameraLabel                   uint16 = 0xc7a1
	IFD_OriginalDefaultFinalSize      uint16 = 0xc791
	IFD_OriginalBestQualityFinalSize  uint16 = 0xc792
	IFD_OriginalDefaultCropSize       uint16 = 0xc793
	IFD_ProfileHueSatMapEncoding      uint16 = 0xc7a3
	IFD_ProfileLookTableEncoding      uint16 = 0xc7a4
	IFD_BaselineExposureOffset        uint16 = 0xc7a5
	IFD_DefaultBlackRender            uint16 = 0xc7a6
	IFD_NewRawImageDigest             uint16 = 0xc7a7
	IFD_RawToPreviewGain              uint16 = 0xc7a8
	IFD_DefaultUserCrop               uint16 = 0xc7b5
	IFD_DepthFormat                   uint16 = 0xc7e9
	IFD_DepthNear                     uint16 = 0xc7ea
	IFD_DepthFar                      uint16 = 0xc7eb
	IFD_DepthUnits                    uint16 = 0xc7ec
	IFD_DepthMeasureType              uint16 = 0xc7ed
	IFD_EnhanceParams                 uint16 = 0xc7ee
	IFD_ProfileGainTableMap           uint16 = 0xcd2d
	IFD_SemanticName                  uint16 = 0xcd2e
	IFD_SemanticInstanceID            uint16 = 0xcd30
	IFD_CalibrationIlluminant3        uint16 = 0xcd31
	IFD_CameraCalibration3            uint16 = 0xcd32
	IFD_ColorMatrix3                  uint16 = 0xcd33
	IFD_ForwardMatrix3                uint16 = 0xcd34
	IFD_IlluminantData1               uint16 = 0xcd35
	IFD_IlluminantData2               uint16 = 0xcd36
	IFD_IlluminantData3               uint16 = 0xcd37
	IFD_ProfileHueSatMapData3         uint16 = 0xcd39
	IFD_ReductionMatrix3              uint16 = 0xcd3a
)

var IFDName = map[uint16]string{
	IFD_ProcessingSoftware:            "ProcessingSoftware",
	IFD_NewSubfileType:                "NewSubfileType",
	IFD_SubfileType:                   "SubfileType",
	IFD_ImageWidth:                    "ImageWidth",
	IFD_ImageLength:                   "ImageLength",
	IFD_BitsPerSample:                 "BitsPerSample",
	IFD_Compression:                   "Compression",
	IFD_PhotometricInterpretation:     "PhotometricInterpretation",
	IFD_Thresholding:                  "Thresholding",
	IFD_CellWidth:                     "CellWidth",
	IFD_CellLength:                    "CellLength",
	IFD_FillOrder:                     "FillOrder",
	IFD_DocumentName:                  "DocumentName",
	IFD_ImageDescription:              "ImageDescription",
	IFD_Make:                          "Make",
	IFD_Model:                         "Model",
	IFD_StripOffsets:                  "StripOffsets",
	IFD_Orientation:                   "Orientation",
	IFD_SamplesPerPixel:               "SamplesPerPixel",
	IFD_RowsPerStrip:                  "RowsPerStrip",
	IFD_StripByteCounts:               "StripByteCounts",
	IFD_XResolution:                   "XResolution",
	IFD_YResolution:                   "YResolution",
	IFD_PlanarConfiguration:           "PlanarConfiguration",
	IFD_GrayResponseUnit:              "GrayResponseUnit",
	IFD_GrayResponseCurve:             "GrayResponseCurve",
	IFD_T4Options:                     "T4Options",
	IFD_T6Options:                     "T6Options",
	IFD_ResolutionUnit:                "ResolutionUnit",
	IFD_PageNumber:                    "PageNumber",
	IFD_TransferFunction:              "TransferFunction",
	IFD_Software:                      "Software",
	IFD_DateTime:                      "DateTime",
	IFD_Artist:                        "Artist",
	IFD_HostComputer:                  "HostComputer",
	IFD_Predictor:                     "Predictor",
	IFD_WhitePoint:                    "WhitePoint",
	IFD_PrimaryChromaticities:         "PrimaryChromaticities",
	IFD_ColorMap:                      "ColorMap",
	IFD_HalftoneHints:                 "HalftoneHints",
	IFD_TileWidth:                     "TileWidth",
	IFD_TileLength:                    "TileLength",
	IFD_TileOffsets:                   "TileOffsets",
	IFD_TileByteCounts:                "TileByteCounts",
	IFD_SubIFDs:                       "SubIFDs",
	IFD_InkSet:                        "InkSet",
	IFD_InkNames:                      "InkNames",
	IFD_NumberOfInks:                  "NumberOfInks",
	IFD_DotRange:                      "DotRange",
	IFD_TargetPrinter:                 "TargetPrinter",
	IFD_ExtraSamples:                  "ExtraSamples",
	IFD_SampleFormat:                  "SampleFormat",
	IFD_SMinSampleValue:               "SMinSampleValue",
	IFD_SMaxSampleValue:               "SMaxSampleValue",
	IFD_TransferRange:                 "TransferRange",
	IFD_ClipPath:                      "ClipPath",
	IFD_XClipPathUnits:                "XClipPathUnits",
	IFD_YClipPathUnits:                "YClipPathUnits",
	IFD_Indexed:                       "Indexed",
	IFD_JPEGTables:                    "JPEGTables",
	IFD_OPIProxy:                      "OPIProxy",
	IFD_JPEGProc:                      "JPEGProc",
	IFD_JPEGInterchangeFormat:         "JPEGInterchangeFormat",
	IFD_JPEGInterchangeFormatLength:   "JPEGInterchangeFormatLength",
	IFD_JPEGRestartInterval:           "JPEGRestartInterval",
	IFD_JPEGLosslessPredictors:        "JPEGLosslessPredictors",
	IFD_JPEGPointTransforms:           "JPEGPointTransforms",
	IFD_JPEGQTables:                   "JPEGQTables",
	IFD_JPEGDCTables:                  "JPEGDCTables",
	IFD_JPEGACTables:                  "JPEGACTables",
	IFD_YCbCrCoefficients:             "YCbCrCoefficients",
	IFD_YCbCrSubSampling:              "YCbCrSubSampling",
	IFD_YCbCrPositioning:              "YCbCrPositioning",
	IFD_ReferenceBlackWhite:           "ReferenceBlackWhite",
	IFD_XMLPacket:                     "XMLPacket",
	IFD_Rating:                        "Rating",
	IFD_RatingPercent:                 "RatingPercent",
	IFD_VignettingCorrParams:          "VignettingCorrParams",
	IFD_ChromaticAberrationCorrParams: "ChromaticAberrationCorrParams",
	IFD_DistortionCorrParams:          "DistortionCorrParams",
	IFD_ImageID:                       "ImageID",
	IFD_CFARepeatPatternDim:           "CFARepeatPatternDim",
	IFD_CFAPattern:                    "CFAPattern",
	IFD_BatteryLevel:                  "BatteryLevel",
	IFD_Copyright:                     "Copyright",
	IFD_ExposureTime:                  "ExposureTime",
	IFD_FNumber:                       "FNumber",
	IFD_IPTCNAA:                       "IPTCNAA",
	IFD_ImageResources:                "ImageResources",
	IFD_ExifTag:                       "ExifTag",
	IFD_InterColorProfile:             "InterColorProfile",
	IFD_ExposureProgram:               "ExposureProgram",
	IFD_SpectralSensitivity:           "SpectralSensitivity",
	IFD_GPSTag:                        "GPSTag",
	IFD_ISOSpeedRatings:               "ISOSpeedRatings",
	IFD_OECF:                          "OECF",
	IFD_Interlace:                     "Interlace",
	IFD_TimeZoneOffset:                "TimeZoneOffset",
	IFD_SelfTimerMode:                 "SelfTimerMode",
	IFD_DateTimeOriginal:              "DateTimeOriginal",
	IFD_CompressedBitsPerPixel:        "CompressedBitsPerPixel",
	IFD_ShutterSpeedValue:             "ShutterSpeedValue",
	IFD_ApertureValue:                 "ApertureValue",
	IFD_BrightnessValue:               "BrightnessValue",
	IFD_ExposureBiasValue:             "ExposureBiasValue",
	IFD_MaxApertureValue:              "MaxApertureValue",
	IFD_SubjectDistance:               "SubjectDistance",
	IFD_MeteringMode:                  "MeteringMode",
	IFD_LightSource:                   "LightSource",
	IFD_Flash:                         "Flash",
	IFD_FocalLength:                   "FocalLength",
	IFD_FlashEnergy:                   "FlashEnergy",
	IFD_SpatialFrequencyResponse:      "SpatialFrequencyResponse",
	IFD_Noise:                         "Noise",
	IFD_FocalPlaneXResolution:         "FocalPlaneXResolution",
	IFD_FocalPlaneYResolution:         "FocalPlaneYResolution",
	IFD_FocalPlaneResolutionUnit:      "FocalPlaneResolutionUnit",
	IFD_ImageNumber:                   "ImageNumber",
	IFD_SecurityClassification:        "SecurityClassification",
	IFD_ImageHistory:                  "ImageHistory",
	IFD_SubjectLocation:               "SubjectLocation",
	IFD_ExposureIndex:                 "ExposureIndex",
	IFD_TIFFEPStandardID:              "TIFFEPStandardID",
	IFD_SensingMethod:                 "SensingMethod",
	IFD_XPTitle:                       "XPTitle",
	IFD_XPComment:                     "XPComment",
	IFD_XPAuthor:                      "XPAuthor",
	IFD_XPKeywords:                    "XPKeywords",
	IFD_XPSubject:                     "XPSubject",
	IFD_PrintImageMatching:            "PrintImageMatching",
	IFD_DNGVersion:                    "DNGVersion",
	IFD_DNGBackwardVersion:            "DNGBackwardVersion",
	IFD_UniqueCameraModel:             "UniqueCameraModel",
	IFD_LocalizedCameraModel:          "LocalizedCameraModel",
	IFD_CFAPlaneColor:                 "CFAPlaneColor",
	IFD_CFALayout:                     "CFALayout",
	IFD_LinearizationTable:            "LinearizationTable",
	IFD_BlackLevelRepeatDim:           "BlackLevelRepeatDim",
	IFD_BlackLevel:                    "BlackLevel",
	IFD_BlackLevelDeltaH:              "BlackLevelDeltaH",
	IFD_BlackLevelDeltaV:              "BlackLevelDeltaV",
	IFD_WhiteLevel:                    "WhiteLevel",
	IFD_DefaultScale:                  "DefaultScale",
	IFD_DefaultCropOrigin:             "DefaultCropOrigin",
	IFD_DefaultCropSize:               "DefaultCropSize",
	IFD_ColorMatrix1:                  "ColorMatrix1",
	IFD_ColorMatrix2:                  "ColorMatrix2",
	IFD_CameraCalibration1:            "CameraCalibration1",
	IFD_CameraCalibration2:            "CameraCalibration2",
	IFD_ReductionMatrix1:              "ReductionMatrix1",
	IFD_ReductionMatrix2:              "ReductionMatrix2",
	IFD_AnalogBalance:                 "AnalogBalance",
	IFD_AsShotNeutral:                 "AsShotNeutral",
	IFD_AsShotWhiteXY:                 "AsShotWhiteXY",
	IFD_BaselineExposure:              "BaselineExposure",
	IFD_BaselineNoise:                 "BaselineNoise",
	IFD_BaselineSharpness:             "BaselineSharpness",
	IFD_BayerGreenSplit:               "BayerGreenSplit",
	IFD_LinearResponseLimit:           "LinearResponseLimit",
	IFD_CameraSerialNumber:            "CameraSerialNumber",
	IFD_LensInfo:                      "LensInfo",
	IFD_ChromaBlurRadius:              "ChromaBlurRadius",
	IFD_AntiAliasStrength:             "AntiAliasStrength",
	IFD_ShadowScale:                   "ShadowScale",
	IFD_DNGPrivateData:                "DNGPrivateData",
	IFD_MakerNoteSafety:               "MakerNoteSafety",
	IFD_CalibrationIlluminant1:        "CalibrationIlluminant1",
	IFD_CalibrationIlluminant2:        "CalibrationIlluminant2",
	IFD_BestQualityScale:              "BestQualityScale",
	IFD_RawDataUniqueID:               "RawDataUniqueID",
	IFD_OriginalRawFileName:           "OriginalRawFileName",
	IFD_OriginalRawFileData:           "OriginalRawFileData",
	IFD_ActiveArea:                    "ActiveArea",
	IFD_MaskedAreas:                   "MaskedAreas",
	IFD_AsShotICCProfile:              "AsShotICCProfile",
	IFD_AsShotPreProfileMatrix:        "AsShotPreProfileMatrix",
	IFD_CurrentICCProfile:             "CurrentICCProfile",
	IFD_CurrentPreProfileMatrix:       "CurrentPreProfileMatrix",
	IFD_ColorimetricReference:         "ColorimetricReference",
	IFD_CameraCalibrationSignature:    "CameraCalibrationSignature",
	IFD_ProfileCalibrationSignature:   "ProfileCalibrationSignature",
	IFD_ExtraCameraProfiles:           "ExtraCameraProfiles",
	IFD_AsShotProfileName:             "AsShotProfileName",
	IFD_NoiseReductionApplied:         "NoiseReductionApplied",
	IFD_ProfileName:                   "ProfileName",
	IFD_ProfileHueSatMapDims:          "ProfileHueSatMapDims",
	IFD_ProfileHueSatMapData1:         "ProfileHueSatMapData1",
	IFD_ProfileHueSatMapData2:         "ProfileHueSatMapData2",
	IFD_ProfileToneCurve:              "ProfileToneCurve",
	IFD_ProfileEmbedPolicy:            "ProfileEmbedPolicy",
	IFD_ProfileCopyright:              "ProfileCopyright",
	IFD_ForwardMatrix1:                "ForwardMatrix1",
	IFD_ForwardMatrix2:                "ForwardMatrix2",
	IFD_PreviewApplicationName:        "PreviewApplicationName",
	IFD_PreviewApplicationVersion:     "PreviewApplicationVersion",
	IFD_PreviewSettingsName:           "PreviewSettingsName",
	IFD_PreviewSettingsDigest:         "PreviewSettingsDigest",
	IFD_PreviewColorSpace:             "PreviewColorSpace",
	IFD_PreviewDateTime:               "PreviewDateTime",
	IFD_RawImageDigest:                "RawImageDigest",
	IFD_OriginalRawFileDigest:         "OriginalRawFileDigest",
	IFD_SubTileBlockSize:              "SubTileBlockSize",
	IFD_RowInterleaveFactor:           "RowInterleaveFactor",
	IFD_ProfileLookTableDims:          "ProfileLookTableDims",
	IFD_ProfileLookTableData:          "ProfileLookTableData",
	IFD_OpcodeList1:                   "OpcodeList1",
	IFD_OpcodeList2:                   "OpcodeList2",
	IFD_OpcodeList3:                   "OpcodeList3",
	IFD_NoiseProfile:                  "NoiseProfile",
	IFD_TimeCodes:                     "TimeCodes",
	IFD_FrameRate:                     "FrameRate",
	IFD_TStop:                         "TStop",
	IFD_ReelName:                      "ReelName",
	IFD_CameraLabel:                   "CameraLabel",
	IFD_OriginalDefaultFinalSize:      "OriginalDefaultFinalSize",
	IFD_OriginalBestQualityFinalSize:  "OriginalBestQualityFinalSize",
	IFD_OriginalDefaultCropSize:       "OriginalDefaultCropSize",
	IFD_ProfileHueSatMapEncoding:      "ProfileHueSatMapEncoding",
	IFD_ProfileLookTableEncoding:      "ProfileLookTableEncoding",
	IFD_BaselineExposureOffset:        "BaselineExposureOffset",
	IFD_DefaultBlackRender:            "DefaultBlackRender",
	IFD_NewRawImageDigest:             "NewRawImageDigest",
	IFD_RawToPreviewGain:              "RawToPreviewGain",
	IFD_DefaultUserCrop:               "DefaultUserCrop",
	IFD_DepthFormat:                   "DepthFormat",
	IFD_DepthNear:                     "DepthNear",
	IFD_DepthFar:                      "DepthFar",
	IFD_DepthUnits:                    "DepthUnits",
	IFD_DepthMeasureType:              "DepthMeasureType",
	IFD_EnhanceParams:                 "EnhanceParams",
	IFD_ProfileGainTableMap:           "ProfileGainTableMap",
	IFD_SemanticName:                  "SemanticName",
	IFD_SemanticInstanceID:            "SemanticInstanceID",
	IFD_CalibrationIlluminant3:        "CalibrationIlluminant3",
	IFD_CameraCalibration3:            "CameraCalibration3",
	IFD_ColorMatrix3:                  "ColorMatrix3",
	IFD_ForwardMatrix3:                "ForwardMatrix3",
	IFD_IlluminantData1:               "IlluminantData1",
	IFD_IlluminantData2:               "IlluminantData2",
	IFD_IlluminantData3:               "IlluminantData3",
	IFD_ProfileHueSatMapData3:         "ProfileHueSatMapData3",
	IFD_ReductionMatrix3:              "ReductionMatrix3",
}

//Exif Tag Ids
const (
	Exif_ExposureTime                        uint16 = 0x829a
	Exif_FNumber                             uint16 = 0x829d
	Exif_ExposureProgram                     uint16 = 0x8822
	Exif_SpectralSensitivity                 uint16 = 0x8824
	Exif_ISOSpeedRatings                     uint16 = 0x8827
	Exif_OECF                                uint16 = 0x8828
	Exif_SensitivityType                     uint16 = 0x8830
	Exif_StandardOutputSensitivity           uint16 = 0x8831
	Exif_RecommendedExposureIndex            uint16 = 0x8832
	Exif_ISOSpeed                            uint16 = 0x8833
	Exif_ISOSpeedLatitudeyyy                 uint16 = 0x8834
	Exif_ISOSpeedLatitudezzz                 uint16 = 0x8835
	Exif_ExifVersion                         uint16 = 0x9000
	Exif_DateTimeOriginal                    uint16 = 0x9003
	Exif_DateTimeDigitized                   uint16 = 0x9004
	Exif_OffsetTime                          uint16 = 0x9010
	Exif_OffsetTimeOriginal                  uint16 = 0x9011
	Exif_OffsetTimeDigitized                 uint16 = 0x9012
	Exif_ComponentsConfiguration             uint16 = 0x9101
	Exif_CompressedBitsPerPixel              uint16 = 0x9102
	Exif_ShutterSpeedValue                   uint16 = 0x9201
	Exif_ApertureValue                       uint16 = 0x9202
	Exif_BrightnessValue                     uint16 = 0x9203
	Exif_ExposureBiasValue                   uint16 = 0x9204
	Exif_MaxApertureValue                    uint16 = 0x9205
	Exif_SubjectDistance                     uint16 = 0x9206
	Exif_MeteringMode                        uint16 = 0x9207
	Exif_LightSource                         uint16 = 0x9208
	Exif_Flash                               uint16 = 0x9209
	Exif_FocalLength                         uint16 = 0x920a
	Exif_SubjectArea                         uint16 = 0x9214
	Exif_MakerNote                           uint16 = 0x927c
	Exif_UserComment                         uint16 = 0x9286
	Exif_SubSecTime                          uint16 = 0x9290
	Exif_SubSecTimeOriginal                  uint16 = 0x9291
	Exif_SubSecTimeDigitized                 uint16 = 0x9292
	Exif_Temperature                         uint16 = 0x9400
	Exif_Humidity                            uint16 = 0x9401
	Exif_Pressure                            uint16 = 0x9402
	Exif_WaterDepth                          uint16 = 0x9403
	Exif_Acceleration                        uint16 = 0x9404
	Exif_CameraElevationAngle                uint16 = 0x9405
	Exif_FlashpixVersion                     uint16 = 0xa000
	Exif_ColorSpace                          uint16 = 0xa001
	Exif_PixelXDimension                     uint16 = 0xa002
	Exif_PixelYDimension                     uint16 = 0xa003
	Exif_RelatedSoundFile                    uint16 = 0xa004
	Exif_InteroperabilityTag                 uint16 = 0xa005
	Exif_FlashEnergy                         uint16 = 0xa20b
	Exif_SpatialFrequencyResponse            uint16 = 0xa20c
	Exif_FocalPlaneXResolution               uint16 = 0xa20e
	Exif_FocalPlaneYResolution               uint16 = 0xa20f
	Exif_FocalPlaneResolutionUnit            uint16 = 0xa210
	Exif_SubjectLocation                     uint16 = 0xa214
	Exif_ExposureIndex                       uint16 = 0xa215
	Exif_SensingMethod                       uint16 = 0xa217
	Exif_FileSource                          uint16 = 0xa300
	Exif_SceneType                           uint16 = 0xa301
	Exif_CFAPattern                          uint16 = 0xa302
	Exif_CustomRendered                      uint16 = 0xa401
	Exif_ExposureMode                        uint16 = 0xa402
	Exif_WhiteBalance                        uint16 = 0xa403
	Exif_DigitalZoomRatio                    uint16 = 0xa404
	Exif_FocalLengthIn35mmFilm               uint16 = 0xa405
	Exif_SceneCaptureType                    uint16 = 0xa406
	Exif_GainControl                         uint16 = 0xa407
	Exif_Contrast                            uint16 = 0xa408
	Exif_Saturation                          uint16 = 0xa409
	Exif_Sharpness                           uint16 = 0xa40a
	Exif_DeviceSettingDescription            uint16 = 0xa40b
	Exif_SubjectDistanceRange                uint16 = 0xa40c
	Exif_ImageUniqueID                       uint16 = 0xa420
	Exif_CameraOwnerName                     uint16 = 0xa430
	Exif_BodySerialNumber                    uint16 = 0xa431
	Exif_LensSpecification                   uint16 = 0xa432
	Exif_LensMake                            uint16 = 0xa433
	Exif_LensModel                           uint16 = 0xa434
	Exif_LensSerialNumber                    uint16 = 0xa435
	Exif_CompositeImage                      uint16 = 0xa460
	Exif_SourceImageNumberOfCompositeImage   uint16 = 0xa461
	Exif_SourceExposureTimesOfCompositeImage uint16 = 0xa462
	Exif_Gamma                               uint16 = 0xa500
)

var ExifName = map[uint16]string{
	Exif_ExposureTime:                        "ExposureTime",
	Exif_FNumber:                             "FNumber",
	Exif_ExposureProgram:                     "ExposureProgram",
	Exif_SpectralSensitivity:                 "SpectralSensitivity",
	Exif_ISOSpeedRatings:                     "ISOSpeedRatings",
	Exif_OECF:                                "OECF",
	Exif_SensitivityType:                     "SensitivityType",
	Exif_StandardOutputSensitivity:           "StandardOutputSensitivity",
	Exif_RecommendedExposureIndex:            "RecommendedExposureIndex",
	Exif_ISOSpeed:                            "ISOSpeed",
	Exif_ISOSpeedLatitudeyyy:                 "ISOSpeedLatitudeyyy",
	Exif_ISOSpeedLatitudezzz:                 "ISOSpeedLatitudezzz",
	Exif_ExifVersion:                         "ExifVersion",
	Exif_DateTimeOriginal:                    "DateTimeOriginal",
	Exif_DateTimeDigitized:                   "DateTimeDigitized",
	Exif_OffsetTime:                          "OffsetTime",
	Exif_OffsetTimeOriginal:                  "OffsetTimeOriginal",
	Exif_OffsetTimeDigitized:                 "OffsetTimeDigitized",
	Exif_ComponentsConfiguration:             "ComponentsConfiguration",
	Exif_CompressedBitsPerPixel:              "CompressedBitsPerPixel",
	Exif_ShutterSpeedValue:                   "ShutterSpeedValue",
	Exif_ApertureValue:                       "ApertureValue",
	Exif_BrightnessValue:                     "BrightnessValue",
	Exif_ExposureBiasValue:                   "ExposureBiasValue",
	Exif_MaxApertureValue:                    "MaxApertureValue",
	Exif_SubjectDistance:                     "SubjectDistance",
	Exif_MeteringMode:                        "MeteringMode",
	Exif_LightSource:                         "LightSource",
	Exif_Flash:                               "Flash",
	Exif_FocalLength:                         "FocalLength",
	Exif_SubjectArea:                         "SubjectArea",
	Exif_MakerNote:                           "MakerNote",
	Exif_UserComment:                         "UserComment",
	Exif_SubSecTime:                          "SubSecTime",
	Exif_SubSecTimeOriginal:                  "SubSecTimeOriginal",
	Exif_SubSecTimeDigitized:                 "SubSecTimeDigitized",
	Exif_Temperature:                         "Temperature",
	Exif_Humidity:                            "Humidity",
	Exif_Pressure:                            "Pressure",
	Exif_WaterDepth:                          "WaterDepth",
	Exif_Acceleration:                        "Acceleration",
	Exif_CameraElevationAngle:                "CameraElevationAngle",
	Exif_FlashpixVersion:                     "FlashpixVersion",
	Exif_ColorSpace:                          "ColorSpace",
	Exif_PixelXDimension:                     "PixelXDimension",
	Exif_PixelYDimension:                     "PixelYDimension",
	Exif_RelatedSoundFile:                    "RelatedSoundFile",
	Exif_InteroperabilityTag:                 "InteroperabilityTag",
	Exif_FlashEnergy:                         "FlashEnergy",
	Exif_SpatialFrequencyResponse:            "SpatialFrequencyResponse",
	Exif_FocalPlaneXResolution:               "FocalPlaneXResolution",
	Exif_FocalPlaneYResolution:               "FocalPlaneYResolution",
	Exif_FocalPlaneResolutionUnit:            "FocalPlaneResolutionUnit",
	Exif_SubjectLocation:                     "SubjectLocation",
	Exif_ExposureIndex:                       "ExposureIndex",
	Exif_SensingMethod:                       "SensingMethod",
	Exif_FileSource:                          "FileSource",
	Exif_SceneType:                           "SceneType",
	Exif_CFAPattern:                          "CFAPattern",
	Exif_CustomRendered:                      "CustomRendered",
	Exif_ExposureMode:                        "ExposureMode",
	Exif_WhiteBalance:                        "WhiteBalance",
	Exif_DigitalZoomRatio:                    "DigitalZoomRatio",
	Exif_FocalLengthIn35mmFilm:               "FocalLengthIn35mmFilm",
	Exif_SceneCaptureType:                    "SceneCaptureType",
	Exif_GainControl:                         "GainControl",
	Exif_Contrast:                            "Contrast",
	Exif_Saturation:                          "Saturation",
	Exif_Sharpness:                           "Sharpness",
	Exif_DeviceSettingDescription:            "DeviceSettingDescription",
	Exif_SubjectDistanceRange:                "SubjectDistanceRange",
	Exif_ImageUniqueID:                       "ImageUniqueID",
	Exif_CameraOwnerName:                     "CameraOwnerName",
	Exif_BodySerialNumber:                    "BodySerialNumber",
	Exif_LensSpecification:                   "LensSpecification",
	Exif_LensMake:                            "LensMake",
	Exif_LensModel:                           "LensModel",
	Exif_LensSerialNumber:                    "LensSerialNumber",
	Exif_CompositeImage:                      "CompositeImage",
	Exif_SourceImageNumberOfCompositeImage:   "SourceImageNumberOfCompositeImage",
	Exif_SourceExposureTimesOfCompositeImage: "SourceExposureTimesOfCompositeImage",
	Exif_Gamma:                               "Gamma",
}

//Iop Tag Ids
const (
	Iop_InteroperabilityIndex   uint16 = 0x0001
	Iop_InteroperabilityVersion uint16 = 0x0002
	Iop_RelatedImageFileFormat  uint16 = 0x1000
	Iop_RelatedImageWidth       uint16 = 0x1001
	Iop_RelatedImageLength      uint16 = 0x1002
)

var IopName = map[uint16]string{
	Iop_InteroperabilityIndex:   "InteroperabilityIndex",
	Iop_InteroperabilityVersion: "InteroperabilityVersion",
	Iop_RelatedImageFileFormat:  "RelatedImageFileFormat",
	Iop_RelatedImageWidth:       "RelatedImageWidth",
	Iop_RelatedImageLength:      "RelatedImageLength",
}

//GPSInfo Tag Ids
const (
	GPSInfo_GPSVersionID         uint16 = 0x0000
	GPSInfo_GPSLatitudeRef       uint16 = 0x0001
	GPSInfo_GPSLatitude          uint16 = 0x0002
	GPSInfo_GPSLongitudeRef      uint16 = 0x0003
	GPSInfo_GPSLongitude         uint16 = 0x0004
	GPSInfo_GPSAltitudeRef       uint16 = 0x0005
	GPSInfo_GPSAltitude          uint16 = 0x0006
	GPSInfo_GPSTimeStamp         uint16 = 0x0007
	GPSInfo_GPSSatellites        uint16 = 0x0008
	GPSInfo_GPSStatus            uint16 = 0x0009
	GPSInfo_GPSMeasureMode       uint16 = 0x000a
	GPSInfo_GPSDOP               uint16 = 0x000b
	GPSInfo_GPSSpeedRef          uint16 = 0x000c
	GPSInfo_GPSSpeed             uint16 = 0x000d
	GPSInfo_GPSTrackRef          uint16 = 0x000e
	GPSInfo_GPSTrack             uint16 = 0x000f
	GPSInfo_GPSImgDirectionRef   uint16 = 0x0010
	GPSInfo_GPSImgDirection      uint16 = 0x0011
	GPSInfo_GPSMapDatum          uint16 = 0x0012
	GPSInfo_GPSDestLatitudeRef   uint16 = 0x0013
	GPSInfo_GPSDestLatitude      uint16 = 0x0014
	GPSInfo_GPSDestLongitudeRef  uint16 = 0x0015
	GPSInfo_GPSDestLongitude     uint16 = 0x0016
	GPSInfo_GPSDestBearingRef    uint16 = 0x0017
	GPSInfo_GPSDestBearing       uint16 = 0x0018
	GPSInfo_GPSDestDistanceRef   uint16 = 0x0019
	GPSInfo_GPSDestDistance      uint16 = 0x001a
	GPSInfo_GPSProcessingMethod  uint16 = 0x001b
	GPSInfo_GPSAreaInformation   uint16 = 0x001c
	GPSInfo_GPSDateStamp         uint16 = 0x001d
	GPSInfo_GPSDifferential      uint16 = 0x001e
	GPSInfo_GPSHPositioningError uint16 = 0x001f
)

var GPSInfoName = map[uint16]string{
	GPSInfo_GPSVersionID:         "GPSVersionID",
	GPSInfo_GPSLatitudeRef:       "GPSLatitudeRef",
	GPSInfo_GPSLatitude:          "GPSLatitude",
	GPSInfo_GPSLongitudeRef:      "GPSLongitudeRef",
	GPSInfo_GPSLongitude:         "GPSLongitude",
	GPSInfo_GPSAltitudeRef:       "GPSAltitudeRef",
	GPSInfo_GPSAltitude:          "GPSAltitude",
	GPSInfo_GPSTimeStamp:         "GPSTimeStamp",
	GPSInfo_GPSSatellites:        "GPSSatellites",
	GPSInfo_GPSStatus:            "GPSStatus",
	GPSInfo_GPSMeasureMode:       "GPSMeasureMode",
	GPSInfo_GPSDOP:               "GPSDOP",
	GPSInfo_GPSSpeedRef:          "GPSSpeedRef",
	GPSInfo_GPSSpeed:             "GPSSpeed",
	GPSInfo_GPSTrackRef:          "GPSTrackRef",
	GPSInfo_GPSTrack:             "GPSTrack",
	GPSInfo_GPSImgDirectionRef:   "GPSImgDirectionRef",
	GPSInfo_GPSImgDirection:      "GPSImgDirection",
	GPSInfo_GPSMapDatum:          "GPSMapDatum",
	GPSInfo_GPSDestLatitudeRef:   "GPSDestLatitudeRef",
	GPSInfo_GPSDestLatitude:      "GPSDestLatitude",
	GPSInfo_GPSDestLongitudeRef:  "GPSDestLongitudeRef",
	GPSInfo_GPSDestLongitude:     "GPSDestLongitude",
	GPSInfo_GPSDestBearingRef:    "GPSDestBearingRef",
	GPSInfo_GPSDestBearing:       "GPSDestBearing",
	GPSInfo_GPSDestDistanceRef:   "GPSDestDistanceRef",
	GPSInfo_GPSDestDistance:      "GPSDestDistance",
	GPSInfo_GPSProcessingMethod:  "GPSProcessingMethod",
	GPSInfo_GPSAreaInformation:   "GPSAreaInformation",
	GPSInfo_GPSDateStamp:         "GPSDateStamp",
	GPSInfo_GPSDifferential:      "GPSDifferential",
	GPSInfo_GPSHPositioningError: "GPSHPositioningError",
}
