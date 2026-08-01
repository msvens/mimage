package metadata

//Do not edit! This is an automatically generated file (see generator.GenerateExifTagsFromListx()).
//Generated from assets/exiftool-listx.xml, produced by: exiftool -listx -EXIF:all

type ExifTag uint16
type ExifTagType uint8
type ExifIndex int
type ExifIndexTag struct {
	Index ExifIndex
	Tag   ExifTag
}

const (
	RootIFD ExifIndex = iota
	ExifIFD
	GpsIFD
	InteropIFD
	ThumbnailIFD
)

var IFDPaths = map[ExifIndex]string{
	RootIFD:      "IFD",
	ExifIFD:      "IFD/Exif",
	GpsIFD:       "IFD/GPSInfo",
	InteropIFD:   "IFD/Exif/Iop",
	ThumbnailIFD: "IFD1",
}

const (
	ExifString ExifTagType = iota
	ExifUint8
	ExifUint16
	ExifUint32
	ExifInt16
	ExifInt32
	ExifRational
	ExifUrational
	ExifFloat
	ExifDouble
	ExifUndef
)

type ExifTagDesc struct {
	Id     ExifTag     `json:"id"`
	Name   string      `json:"name"`
	Type   ExifTagType `json:"type"`
	Ifd    ExifIndex   `json:"ifd"`
	Count  int         `json:"count"`
	Values interface{} `json:"values"`
}

// IFD Tag Ids (includes all IFD, IFD1, etc tags)
const (
	IFD_ProcessingSoftware            ExifTag = 0x000b
	IFD_SubfileType                   ExifTag = 0x00fe
	IFD_OldSubfileType                ExifTag = 0x00ff
	IFD_ImageWidth                    ExifTag = 0x0100
	IFD_ImageHeight                   ExifTag = 0x0101
	IFD_BitsPerSample                 ExifTag = 0x0102
	IFD_Compression                   ExifTag = 0x0103
	IFD_PhotometricInterpretation     ExifTag = 0x0106
	IFD_Thresholding                  ExifTag = 0x0107
	IFD_CellWidth                     ExifTag = 0x0108
	IFD_CellLength                    ExifTag = 0x0109
	IFD_FillOrder                     ExifTag = 0x010a
	IFD_DocumentName                  ExifTag = 0x010d
	IFD_ImageDescription              ExifTag = 0x010e
	IFD_Make                          ExifTag = 0x010f
	IFD_Model                         ExifTag = 0x0110
	IFD_StripOffsets                  ExifTag = 0x0111
	IFD_Orientation                   ExifTag = 0x0112
	IFD_SamplesPerPixel               ExifTag = 0x0115
	IFD_RowsPerStrip                  ExifTag = 0x0116
	IFD_StripByteCounts               ExifTag = 0x0117
	IFD_MinSampleValue                ExifTag = 0x0118
	IFD_MaxSampleValue                ExifTag = 0x0119
	IFD_XResolution                   ExifTag = 0x011a
	IFD_YResolution                   ExifTag = 0x011b
	IFD_PlanarConfiguration           ExifTag = 0x011c
	IFD_PageName                      ExifTag = 0x011d
	IFD_XPosition                     ExifTag = 0x011e
	IFD_YPosition                     ExifTag = 0x011f
	IFD_GrayResponseUnit              ExifTag = 0x0122
	IFD_ResolutionUnit                ExifTag = 0x0128
	IFD_PageNumber                    ExifTag = 0x0129
	IFD_TransferFunction              ExifTag = 0x012d
	IFD_Software                      ExifTag = 0x0131
	IFD_ModifyDate                    ExifTag = 0x0132
	IFD_Artist                        ExifTag = 0x013b
	IFD_HostComputer                  ExifTag = 0x013c
	IFD_Predictor                     ExifTag = 0x013d
	IFD_WhitePoint                    ExifTag = 0x013e
	IFD_PrimaryChromaticities         ExifTag = 0x013f
	IFD_HalftoneHints                 ExifTag = 0x0141
	IFD_TileWidth                     ExifTag = 0x0142
	IFD_TileLength                    ExifTag = 0x0143
	IFD_SubIFDs                       ExifTag = 0x014a
	IFD_InkSet                        ExifTag = 0x014c
	IFD_TargetPrinter                 ExifTag = 0x0151
	IFD_SampleFormat                  ExifTag = 0x0153
	IFD_ThumbnailOffset               ExifTag = 0x0201
	IFD_ThumbnailLength               ExifTag = 0x0202
	IFD_YCbCrCoefficients             ExifTag = 0x0211
	IFD_YCbCrSubSampling              ExifTag = 0x0212
	IFD_YCbCrPositioning              ExifTag = 0x0213
	IFD_ReferenceBlackWhite           ExifTag = 0x0214
	IFD_ApplicationNotes              ExifTag = 0x02bc
	IFD_Rating                        ExifTag = 0x4746
	IFD_RatingPercent                 ExifTag = 0x4749
	IFD_VignettingCorrection          ExifTag = 0x7031
	IFD_VignettingCorrParams          ExifTag = 0x7032
	IFD_ChromaticAberrationCorrection ExifTag = 0x7034
	IFD_ChromaticAberrationCorrParams ExifTag = 0x7035
	IFD_DistortionCorrection          ExifTag = 0x7036
	IFD_DistortionCorrParams          ExifTag = 0x7037
	IFD_SonyRawImageSize              ExifTag = 0x7038
	IFD_BlackLevel_0x7310             ExifTag = 0x7310
	IFD_WB_RGGBLevels                 ExifTag = 0x7313
	IFD_SonyCropTopLeft               ExifTag = 0x74c7
	IFD_SonyCropSize                  ExifTag = 0x74c8
	IFD_CFARepeatPatternDim           ExifTag = 0x828d
	IFD_CFAPattern2                   ExifTag = 0x828e
	IFD_Copyright                     ExifTag = 0x8298
	IFD_PixelScale                    ExifTag = 0x830e
	IFD_IPTCNAA                       ExifTag = 0x83bb
	IFD_IntergraphMatrix              ExifTag = 0x8480
	IFD_ModelTiePoint                 ExifTag = 0x8482
	IFD_SEMInfo                       ExifTag = 0x8546
	IFD_ModelTransform                ExifTag = 0x85d8
	IFD_PhotoshopSettings             ExifTag = 0x8649
	IFD_ExifOffset                    ExifTag = 0x8769
	IFD_ICC_Profile                   ExifTag = 0x8773
	IFD_GeoTiffDirectory              ExifTag = 0x87af
	IFD_GeoTiffDoubleParams           ExifTag = 0x87b0
	IFD_GeoTiffAsciiParams            ExifTag = 0x87b1
	IFD_GPSInfo                       ExifTag = 0x8825
	IFD_ImageSourceData               ExifTag = 0x935c
	IFD_XPTitle                       ExifTag = 0x9c9b
	IFD_XPComment                     ExifTag = 0x9c9c
	IFD_XPAuthor                      ExifTag = 0x9c9d
	IFD_XPKeywords                    ExifTag = 0x9c9e
	IFD_XPSubject                     ExifTag = 0x9c9f
	IFD_GDALMetadata                  ExifTag = 0xa480
	IFD_GDALNoData                    ExifTag = 0xa481
	IFD_PrintIM                       ExifTag = 0xc4a5
	IFD_DNGVersion                    ExifTag = 0xc612
	IFD_DNGBackwardVersion            ExifTag = 0xc613
	IFD_UniqueCameraModel             ExifTag = 0xc614
	IFD_LocalizedCameraModel          ExifTag = 0xc615
	IFD_CFAPlaneColor                 ExifTag = 0xc616
	IFD_CFALayout                     ExifTag = 0xc617
	IFD_LinearizationTable            ExifTag = 0xc618
	IFD_BlackLevelRepeatDim           ExifTag = 0xc619
	IFD_BlackLevel                    ExifTag = 0xc61a
	IFD_BlackLevelDeltaH              ExifTag = 0xc61b
	IFD_BlackLevelDeltaV              ExifTag = 0xc61c
	IFD_WhiteLevel                    ExifTag = 0xc61d
	IFD_DefaultScale                  ExifTag = 0xc61e
	IFD_DefaultCropOrigin             ExifTag = 0xc61f
	IFD_DefaultCropSize               ExifTag = 0xc620
	IFD_ColorMatrix1                  ExifTag = 0xc621
	IFD_ColorMatrix2                  ExifTag = 0xc622
	IFD_CameraCalibration1            ExifTag = 0xc623
	IFD_CameraCalibration2            ExifTag = 0xc624
	IFD_ReductionMatrix1              ExifTag = 0xc625
	IFD_ReductionMatrix2              ExifTag = 0xc626
	IFD_AnalogBalance                 ExifTag = 0xc627
	IFD_AsShotNeutral                 ExifTag = 0xc628
	IFD_AsShotWhiteXY                 ExifTag = 0xc629
	IFD_BaselineExposure              ExifTag = 0xc62a
	IFD_BaselineNoise                 ExifTag = 0xc62b
	IFD_BaselineSharpness             ExifTag = 0xc62c
	IFD_BayerGreenSplit               ExifTag = 0xc62d
	IFD_LinearResponseLimit           ExifTag = 0xc62e
	IFD_CameraSerialNumber            ExifTag = 0xc62f
	IFD_DNGLensInfo                   ExifTag = 0xc630
	IFD_ChromaBlurRadius              ExifTag = 0xc631
	IFD_AntiAliasStrength             ExifTag = 0xc632
	IFD_ShadowScale                   ExifTag = 0xc633
	IFD_DNGPrivateData                ExifTag = 0xc634
	IFD_MakerNoteSafety               ExifTag = 0xc635
	IFD_CalibrationIlluminant1        ExifTag = 0xc65a
	IFD_CalibrationIlluminant2        ExifTag = 0xc65b
	IFD_BestQualityScale              ExifTag = 0xc65c
	IFD_RawDataUniqueID               ExifTag = 0xc65d
	IFD_OriginalRawFileName           ExifTag = 0xc68b
	IFD_OriginalRawFileData           ExifTag = 0xc68c
	IFD_ActiveArea                    ExifTag = 0xc68d
	IFD_MaskedAreas                   ExifTag = 0xc68e
	IFD_AsShotICCProfile              ExifTag = 0xc68f
	IFD_AsShotPreProfileMatrix        ExifTag = 0xc690
	IFD_CurrentICCProfile             ExifTag = 0xc691
	IFD_CurrentPreProfileMatrix       ExifTag = 0xc692
	IFD_ColorimetricReference         ExifTag = 0xc6bf
	IFD_SRawType                      ExifTag = 0xc6c5
	IFD_PanasonicTitle                ExifTag = 0xc6d2
	IFD_PanasonicTitle2               ExifTag = 0xc6d3
	IFD_CameraCalibrationSig          ExifTag = 0xc6f3
	IFD_ProfileCalibrationSig         ExifTag = 0xc6f4
	IFD_ProfileIFD                    ExifTag = 0xc6f5
	IFD_AsShotProfileName             ExifTag = 0xc6f6
	IFD_NoiseReductionApplied         ExifTag = 0xc6f7
	IFD_ProfileName                   ExifTag = 0xc6f8
	IFD_ProfileHueSatMapDims          ExifTag = 0xc6f9
	IFD_ProfileHueSatMapData1         ExifTag = 0xc6fa
	IFD_ProfileHueSatMapData2         ExifTag = 0xc6fb
	IFD_ProfileToneCurve              ExifTag = 0xc6fc
	IFD_ProfileEmbedPolicy            ExifTag = 0xc6fd
	IFD_ProfileCopyright              ExifTag = 0xc6fe
	IFD_ForwardMatrix1                ExifTag = 0xc714
	IFD_ForwardMatrix2                ExifTag = 0xc715
	IFD_PreviewApplicationName        ExifTag = 0xc716
	IFD_PreviewApplicationVersion     ExifTag = 0xc717
	IFD_PreviewSettingsName           ExifTag = 0xc718
	IFD_PreviewSettingsDigest         ExifTag = 0xc719
	IFD_PreviewColorSpace             ExifTag = 0xc71a
	IFD_PreviewDateTime               ExifTag = 0xc71b
	IFD_RawImageDigest                ExifTag = 0xc71c
	IFD_OriginalRawFileDigest         ExifTag = 0xc71d
	IFD_ProfileLookTableDims          ExifTag = 0xc725
	IFD_ProfileLookTableData          ExifTag = 0xc726
	IFD_OpcodeList1                   ExifTag = 0xc740
	IFD_OpcodeList2                   ExifTag = 0xc741
	IFD_OpcodeList3                   ExifTag = 0xc74e
	IFD_NoiseProfile                  ExifTag = 0xc761
	IFD_TimeCodes                     ExifTag = 0xc763
	IFD_FrameRate                     ExifTag = 0xc764
	IFD_TStop                         ExifTag = 0xc772
	IFD_ReelName                      ExifTag = 0xc789
	IFD_OriginalDefaultFinalSize      ExifTag = 0xc791
	IFD_OriginalBestQualitySize       ExifTag = 0xc792
	IFD_OriginalDefaultCropSize       ExifTag = 0xc793
	IFD_CameraLabel                   ExifTag = 0xc7a1
	IFD_ProfileHueSatMapEncoding      ExifTag = 0xc7a3
	IFD_ProfileLookTableEncoding      ExifTag = 0xc7a4
	IFD_BaselineExposureOffset        ExifTag = 0xc7a5
	IFD_DefaultBlackRender            ExifTag = 0xc7a6
	IFD_NewRawImageDigest             ExifTag = 0xc7a7
	IFD_RawToPreviewGain              ExifTag = 0xc7a8
	IFD_CacheVersion                  ExifTag = 0xc7aa
	IFD_DefaultUserCrop               ExifTag = 0xc7b5
	IFD_DepthFormat                   ExifTag = 0xc7e9
	IFD_DepthNear                     ExifTag = 0xc7ea
	IFD_DepthFar                      ExifTag = 0xc7eb
	IFD_DepthUnits                    ExifTag = 0xc7ec
	IFD_DepthMeasureType              ExifTag = 0xc7ed
	IFD_EnhanceParams                 ExifTag = 0xc7ee
	IFD_ProfileGainTableMap           ExifTag = 0xcd2d
	IFD_SemanticName                  ExifTag = 0xcd2e
	IFD_SemanticInstanceIFD           ExifTag = 0xcd30
	IFD_CalibrationIlluminant3        ExifTag = 0xcd31
	IFD_CameraCalibration3            ExifTag = 0xcd32
	IFD_ColorMatrix3                  ExifTag = 0xcd33
	IFD_ForwardMatrix3                ExifTag = 0xcd34
	IFD_IlluminantData1               ExifTag = 0xcd35
	IFD_IlluminantData2               ExifTag = 0xcd36
	IFD_IlluminantData3               ExifTag = 0xcd37
	IFD_MaskSubArea                   ExifTag = 0xcd38
	IFD_ProfileHueSatMapData3         ExifTag = 0xcd39
	IFD_ReductionMatrix3              ExifTag = 0xcd3a
	IFD_RGBTables                     ExifTag = 0xcd3b
	IFD_RGBTables_0xcd3f              ExifTag = 0xcd3f
	IFD_ProfileGainTableMap2          ExifTag = 0xcd40
	IFD_ColumnInterleaveFactor        ExifTag = 0xcd43
	IFD_ImageSequenceInfo             ExifTag = 0xcd44
	IFD_ImageStats                    ExifTag = 0xcd46
	IFD_ProfileDynamicRange           ExifTag = 0xcd47
	IFD_ProfileGroupName              ExifTag = 0xcd48
	IFD_JXLDistance                   ExifTag = 0xcd49
	IFD_JXLEffort                     ExifTag = 0xcd4a
	IFD_JXLDecodeSpeed                ExifTag = 0xcd4b
	IFD_SEAL                          ExifTag = 0xcea1
)

// ExifIFD Tag Ids
const (
	ExifIFD_StripOffsets                    ExifTag = 0x0111
	ExifIFD_StripByteCounts                 ExifTag = 0x0117
	ExifIFD_FreeOffsets                     ExifTag = 0x0120
	ExifIFD_FreeByteCounts                  ExifTag = 0x0121
	ExifIFD_GrayResponseCurve               ExifTag = 0x0123
	ExifIFD_T4Options                       ExifTag = 0x0124
	ExifIFD_T6Options                       ExifTag = 0x0125
	ExifIFD_ColorResponseUnit               ExifTag = 0x012c
	ExifIFD_ColorMap                        ExifTag = 0x0140
	ExifIFD_TileOffsets                     ExifTag = 0x0144
	ExifIFD_TileByteCounts                  ExifTag = 0x0145
	ExifIFD_BadFaxLines                     ExifTag = 0x0146
	ExifIFD_CleanFaxData                    ExifTag = 0x0147
	ExifIFD_ConsecutiveBadFaxLines          ExifTag = 0x0148
	ExifIFD_InkNames                        ExifTag = 0x014d
	ExifIFD_NumberofInks                    ExifTag = 0x014e
	ExifIFD_DotRange                        ExifTag = 0x0150
	ExifIFD_ExtraSamples                    ExifTag = 0x0152
	ExifIFD_SMinSampleValue                 ExifTag = 0x0154
	ExifIFD_SMaxSampleValue                 ExifTag = 0x0155
	ExifIFD_TransferRange                   ExifTag = 0x0156
	ExifIFD_ClipPath                        ExifTag = 0x0157
	ExifIFD_XClipPathUnits                  ExifTag = 0x0158
	ExifIFD_YClipPathUnits                  ExifTag = 0x0159
	ExifIFD_Indexed                         ExifTag = 0x015a
	ExifIFD_JPEGTables                      ExifTag = 0x015b
	ExifIFD_OPIProxy                        ExifTag = 0x015f
	ExifIFD_GlobalParametersIFD             ExifTag = 0x0190
	ExifIFD_ProfileType                     ExifTag = 0x0191
	ExifIFD_FaxProfile                      ExifTag = 0x0192
	ExifIFD_CodingMethods                   ExifTag = 0x0193
	ExifIFD_VersionYear                     ExifTag = 0x0194
	ExifIFD_ModeNumber                      ExifTag = 0x0195
	ExifIFD_Decode                          ExifTag = 0x01b1
	ExifIFD_DefaultImageColor               ExifTag = 0x01b2
	ExifIFD_T82Options                      ExifTag = 0x01b3
	ExifIFD_JPEGTables_0x01b5               ExifTag = 0x01b5
	ExifIFD_JPEGProc                        ExifTag = 0x0200
	ExifIFD_JPEGRestartInterval             ExifTag = 0x0203
	ExifIFD_JPEGLosslessPredictors          ExifTag = 0x0205
	ExifIFD_JPEGPointTransforms             ExifTag = 0x0206
	ExifIFD_JPEGQTables                     ExifTag = 0x0207
	ExifIFD_JPEGDCTables                    ExifTag = 0x0208
	ExifIFD_JPEGACTables                    ExifTag = 0x0209
	ExifIFD_StripRowCounts                  ExifTag = 0x022f
	ExifIFD_RenderingIntent                 ExifTag = 0x0303
	ExifIFD_USPTOMiscellaneous              ExifTag = 0x03e7
	ExifIFD_XP_DIP_XML                      ExifTag = 0x4747
	ExifIFD_StitchInfo                      ExifTag = 0x4748
	ExifIFD_ResolutionXUnit                 ExifTag = 0x5001
	ExifIFD_ResolutionYUnit                 ExifTag = 0x5002
	ExifIFD_ResolutionXLengthUnit           ExifTag = 0x5003
	ExifIFD_ResolutionYLengthUnit           ExifTag = 0x5004
	ExifIFD_PrintFlags                      ExifTag = 0x5005
	ExifIFD_PrintFlagsVersion               ExifTag = 0x5006
	ExifIFD_PrintFlagsCrop                  ExifTag = 0x5007
	ExifIFD_PrintFlagsBleedWidth            ExifTag = 0x5008
	ExifIFD_PrintFlagsBleedWidthScale       ExifTag = 0x5009
	ExifIFD_HalftoneLPI                     ExifTag = 0x500a
	ExifIFD_HalftoneLPIUnit                 ExifTag = 0x500b
	ExifIFD_HalftoneDegree                  ExifTag = 0x500c
	ExifIFD_HalftoneShape                   ExifTag = 0x500d
	ExifIFD_HalftoneMisc                    ExifTag = 0x500e
	ExifIFD_HalftoneScreen                  ExifTag = 0x500f
	ExifIFD_JPEGQuality                     ExifTag = 0x5010
	ExifIFD_GridSize                        ExifTag = 0x5011
	ExifIFD_ThumbnailFormat                 ExifTag = 0x5012
	ExifIFD_ThumbnailWidth                  ExifTag = 0x5013
	ExifIFD_ThumbnailHeight                 ExifTag = 0x5014
	ExifIFD_ThumbnailColorDepth             ExifTag = 0x5015
	ExifIFD_ThumbnailPlanes                 ExifTag = 0x5016
	ExifIFD_ThumbnailRawBytes               ExifTag = 0x5017
	ExifIFD_ThumbnailLength                 ExifTag = 0x5018
	ExifIFD_ThumbnailCompressedSize         ExifTag = 0x5019
	ExifIFD_ColorTransferFunction           ExifTag = 0x501a
	ExifIFD_ThumbnailData                   ExifTag = 0x501b
	ExifIFD_ThumbnailImageWidth             ExifTag = 0x5020
	ExifIFD_ThumbnailImageHeight            ExifTag = 0x5021
	ExifIFD_ThumbnailBitsPerSample          ExifTag = 0x5022
	ExifIFD_ThumbnailCompression            ExifTag = 0x5023
	ExifIFD_ThumbnailPhotometricInterp      ExifTag = 0x5024
	ExifIFD_ThumbnailDescription            ExifTag = 0x5025
	ExifIFD_ThumbnailEquipMake              ExifTag = 0x5026
	ExifIFD_ThumbnailEquipModel             ExifTag = 0x5027
	ExifIFD_ThumbnailStripOffsets           ExifTag = 0x5028
	ExifIFD_ThumbnailOrientation            ExifTag = 0x5029
	ExifIFD_ThumbnailSamplesPerPixel        ExifTag = 0x502a
	ExifIFD_ThumbnailRowsPerStrip           ExifTag = 0x502b
	ExifIFD_ThumbnailStripByteCounts        ExifTag = 0x502c
	ExifIFD_ThumbnailResolutionX            ExifTag = 0x502d
	ExifIFD_ThumbnailResolutionY            ExifTag = 0x502e
	ExifIFD_ThumbnailPlanarConfig           ExifTag = 0x502f
	ExifIFD_ThumbnailResolutionUnit         ExifTag = 0x5030
	ExifIFD_ThumbnailTransferFunction       ExifTag = 0x5031
	ExifIFD_ThumbnailSoftware               ExifTag = 0x5032
	ExifIFD_ThumbnailDateTime               ExifTag = 0x5033
	ExifIFD_ThumbnailArtist                 ExifTag = 0x5034
	ExifIFD_ThumbnailWhitePoint             ExifTag = 0x5035
	ExifIFD_ThumbnailPrimaryChromaticities  ExifTag = 0x5036
	ExifIFD_ThumbnailYCbCrCoefficients      ExifTag = 0x5037
	ExifIFD_ThumbnailYCbCrSubsampling       ExifTag = 0x5038
	ExifIFD_ThumbnailYCbCrPositioning       ExifTag = 0x5039
	ExifIFD_ThumbnailRefBlackWhite          ExifTag = 0x503a
	ExifIFD_ThumbnailCopyright              ExifTag = 0x503b
	ExifIFD_LuminanceTable                  ExifTag = 0x5090
	ExifIFD_ChrominanceTable                ExifTag = 0x5091
	ExifIFD_FrameDelay                      ExifTag = 0x5100
	ExifIFD_LoopCount                       ExifTag = 0x5101
	ExifIFD_GlobalPalette                   ExifTag = 0x5102
	ExifIFD_IndexBackground                 ExifTag = 0x5103
	ExifIFD_IndexTransparent                ExifTag = 0x5104
	ExifIFD_PixelUnits                      ExifTag = 0x5110
	ExifIFD_PixelsPerUnitX                  ExifTag = 0x5111
	ExifIFD_PixelsPerUnitY                  ExifTag = 0x5112
	ExifIFD_PaletteHistogram                ExifTag = 0x5113
	ExifIFD_SonyRawFileType                 ExifTag = 0x7000
	ExifIFD_SonyToneCurve                   ExifTag = 0x7010
	ExifIFD_ImageID                         ExifTag = 0x800d
	ExifIFD_WangTag1                        ExifTag = 0x80a3
	ExifIFD_WangAnnotation                  ExifTag = 0x80a4
	ExifIFD_WangTag3                        ExifTag = 0x80a5
	ExifIFD_WangTag4                        ExifTag = 0x80a6
	ExifIFD_ImageReferencePoints            ExifTag = 0x80b9
	ExifIFD_RegionXformTackPoint            ExifTag = 0x80ba
	ExifIFD_WarpQuadrilateral               ExifTag = 0x80bb
	ExifIFD_AffineTransformMat              ExifTag = 0x80bc
	ExifIFD_Matteing                        ExifTag = 0x80e3
	ExifIFD_DataType                        ExifTag = 0x80e4
	ExifIFD_ImageDepth                      ExifTag = 0x80e5
	ExifIFD_TileDepth                       ExifTag = 0x80e6
	ExifIFD_ImageFullWidth                  ExifTag = 0x8214
	ExifIFD_ImageFullHeight                 ExifTag = 0x8215
	ExifIFD_TextureFormat                   ExifTag = 0x8216
	ExifIFD_WrapModes                       ExifTag = 0x8217
	ExifIFD_FovCot                          ExifTag = 0x8218
	ExifIFD_MatrixWorldToScreen             ExifTag = 0x8219
	ExifIFD_MatrixWorldToCamera             ExifTag = 0x821a
	ExifIFD_Model2                          ExifTag = 0x827d
	ExifIFD_BatteryLevel                    ExifTag = 0x828f
	ExifIFD_KodakIFD                        ExifTag = 0x8290
	ExifIFD_ExposureTime                    ExifTag = 0x829a
	ExifIFD_FNumber                         ExifTag = 0x829d
	ExifIFD_MDFileTag                       ExifTag = 0x82a5
	ExifIFD_MDScalePixel                    ExifTag = 0x82a6
	ExifIFD_MDColorTable                    ExifTag = 0x82a7
	ExifIFD_MDLabName                       ExifTag = 0x82a8
	ExifIFD_MDSampleInfo                    ExifTag = 0x82a9
	ExifIFD_MDPrepDate                      ExifTag = 0x82aa
	ExifIFD_MDPrepTime                      ExifTag = 0x82ab
	ExifIFD_MDFileUnits                     ExifTag = 0x82ac
	ExifIFD_AdventScale                     ExifTag = 0x8335
	ExifIFD_AdventRevision                  ExifTag = 0x8336
	ExifIFD_UIC1Tag                         ExifTag = 0x835c
	ExifIFD_UIC2Tag                         ExifTag = 0x835d
	ExifIFD_UIC3Tag                         ExifTag = 0x835e
	ExifIFD_UIC4Tag                         ExifTag = 0x835f
	ExifIFD_IntergraphPacketData            ExifTag = 0x847e
	ExifIFD_IntergraphFlagRegisters         ExifTag = 0x847f
	ExifIFD_INGRReserved                    ExifTag = 0x8481
	ExifIFD_Site                            ExifTag = 0x84e0
	ExifIFD_ColorSequence                   ExifTag = 0x84e1
	ExifIFD_IT8Header                       ExifTag = 0x84e2
	ExifIFD_RasterPadding                   ExifTag = 0x84e3
	ExifIFD_BitsPerRunLength                ExifTag = 0x84e4
	ExifIFD_BitsPerExtendedRunLength        ExifTag = 0x84e5
	ExifIFD_ColorTable                      ExifTag = 0x84e6
	ExifIFD_ImageColorIndicator             ExifTag = 0x84e7
	ExifIFD_BackgroundColorIndicator        ExifTag = 0x84e8
	ExifIFD_ImageColorValue                 ExifTag = 0x84e9
	ExifIFD_BackgroundColorValue            ExifTag = 0x84ea
	ExifIFD_PixelIntensityRange             ExifTag = 0x84eb
	ExifIFD_TransparencyIndicator           ExifTag = 0x84ec
	ExifIFD_ColorCharacterization           ExifTag = 0x84ed
	ExifIFD_HCUsage                         ExifTag = 0x84ee
	ExifIFD_TrapIndicator                   ExifTag = 0x84ef
	ExifIFD_CMYKEquivalent                  ExifTag = 0x84f0
	ExifIFD_AFCP_IPTC                       ExifTag = 0x8568
	ExifIFD_PixelMagicJBIGOptions           ExifTag = 0x85b8
	ExifIFD_JPLCartoIFD                     ExifTag = 0x85d7
	ExifIFD_WB_GRGBLevels                   ExifTag = 0x8602
	ExifIFD_LeafData                        ExifTag = 0x8606
	ExifIFD_TIFF_FXExtensions               ExifTag = 0x877f
	ExifIFD_MultiProfiles                   ExifTag = 0x8780
	ExifIFD_SharedData                      ExifTag = 0x8781
	ExifIFD_T88Options                      ExifTag = 0x8782
	ExifIFD_ImageLayer                      ExifTag = 0x87ac
	ExifIFD_JBIGOptions                     ExifTag = 0x87be
	ExifIFD_ExposureProgram                 ExifTag = 0x8822
	ExifIFD_SpectralSensitivity             ExifTag = 0x8824
	ExifIFD_ISO                             ExifTag = 0x8827
	ExifIFD_OptoElectricConvFactor          ExifTag = 0x8828
	ExifIFD_Interlace                       ExifTag = 0x8829
	ExifIFD_TimeZoneOffset                  ExifTag = 0x882a
	ExifIFD_SelfTimerMode                   ExifTag = 0x882b
	ExifIFD_SensitivityType                 ExifTag = 0x8830
	ExifIFD_StandardOutputSensitivity       ExifTag = 0x8831
	ExifIFD_RecommendedExposureIndex        ExifTag = 0x8832
	ExifIFD_ISOSpeed                        ExifTag = 0x8833
	ExifIFD_ISOSpeedLatitudeyyy             ExifTag = 0x8834
	ExifIFD_ISOSpeedLatitudezzz             ExifTag = 0x8835
	ExifIFD_FaxRecvParams                   ExifTag = 0x885c
	ExifIFD_FaxSubAddress                   ExifTag = 0x885d
	ExifIFD_FaxRecvTime                     ExifTag = 0x885e
	ExifIFD_FedexEDR                        ExifTag = 0x8871
	ExifIFD_LeafSubIFD                      ExifTag = 0x888a
	ExifIFD_ExifVersion                     ExifTag = 0x9000
	ExifIFD_DateTimeOriginal                ExifTag = 0x9003
	ExifIFD_CreateDate                      ExifTag = 0x9004
	ExifIFD_GooglePlusUploadCode            ExifTag = 0x9009
	ExifIFD_OffsetTime                      ExifTag = 0x9010
	ExifIFD_OffsetTimeOriginal              ExifTag = 0x9011
	ExifIFD_OffsetTimeDigitized             ExifTag = 0x9012
	ExifIFD_ComponentsConfiguration         ExifTag = 0x9101
	ExifIFD_CompressedBitsPerPixel          ExifTag = 0x9102
	ExifIFD_ShutterSpeedValue               ExifTag = 0x9201
	ExifIFD_ApertureValue                   ExifTag = 0x9202
	ExifIFD_BrightnessValue                 ExifTag = 0x9203
	ExifIFD_ExposureCompensation            ExifTag = 0x9204
	ExifIFD_MaxApertureValue                ExifTag = 0x9205
	ExifIFD_SubjectDistance                 ExifTag = 0x9206
	ExifIFD_MeteringMode                    ExifTag = 0x9207
	ExifIFD_LightSource                     ExifTag = 0x9208
	ExifIFD_Flash                           ExifTag = 0x9209
	ExifIFD_FocalLength                     ExifTag = 0x920a
	ExifIFD_FlashEnergy                     ExifTag = 0x920b
	ExifIFD_SpatialFrequencyResponse_0x920c ExifTag = 0x920c
	ExifIFD_Noise                           ExifTag = 0x920d
	ExifIFD_FocalPlaneXResolution_0x920e    ExifTag = 0x920e
	ExifIFD_FocalPlaneYResolution_0x920f    ExifTag = 0x920f
	ExifIFD_FocalPlaneResolutionUnit        ExifTag = 0x9210
	ExifIFD_ImageNumber                     ExifTag = 0x9211
	ExifIFD_SecurityClassification          ExifTag = 0x9212
	ExifIFD_ImageHistory                    ExifTag = 0x9213
	ExifIFD_SubjectArea                     ExifTag = 0x9214
	ExifIFD_ExposureIndex_0x9215            ExifTag = 0x9215
	ExifIFD_TIFFEPStandardID                ExifTag = 0x9216
	ExifIFD_SensingMethod                   ExifTag = 0x9217
	ExifIFD_CIP3DataFile                    ExifTag = 0x923a
	ExifIFD_CIP3Sheet                       ExifTag = 0x923b
	ExifIFD_CIP3Side                        ExifTag = 0x923c
	ExifIFD_StoNits                         ExifTag = 0x923f
	ExifIFD_MakerNote                       ExifTag = 0x927c
	ExifIFD_UserComment                     ExifTag = 0x9286
	ExifIFD_SubSecTime                      ExifTag = 0x9290
	ExifIFD_SubSecTimeOriginal              ExifTag = 0x9291
	ExifIFD_SubSecTimeDigitized             ExifTag = 0x9292
	ExifIFD_MSDocumentText                  ExifTag = 0x932f
	ExifIFD_MSPropertySetStorage            ExifTag = 0x9330
	ExifIFD_MSDocumentTextPosition          ExifTag = 0x9331
	ExifIFD_AmbientTemperature              ExifTag = 0x9400
	ExifIFD_Humidity                        ExifTag = 0x9401
	ExifIFD_Pressure                        ExifTag = 0x9402
	ExifIFD_WaterDepth                      ExifTag = 0x9403
	ExifIFD_Acceleration                    ExifTag = 0x9404
	ExifIFD_CameraElevationAngle            ExifTag = 0x9405
	ExifIFD_XiaomiSettings                  ExifTag = 0x9999
	ExifIFD_XiaomiModel                     ExifTag = 0x9a00
	ExifIFD_FlashpixVersion                 ExifTag = 0xa000
	ExifIFD_ColorSpace                      ExifTag = 0xa001
	ExifIFD_ExifImageWidth                  ExifTag = 0xa002
	ExifIFD_ExifImageHeight                 ExifTag = 0xa003
	ExifIFD_RelatedSoundFile                ExifTag = 0xa004
	ExifIFD_InteropOffset                   ExifTag = 0xa005
	ExifIFD_SamsungRawPointersOffset        ExifTag = 0xa010
	ExifIFD_SamsungRawPointersLength        ExifTag = 0xa011
	ExifIFD_SamsungRawByteOrder             ExifTag = 0xa101
	ExifIFD_SamsungRawUnknown               ExifTag = 0xa102
	ExifIFD_FlashEnergy_0xa20b              ExifTag = 0xa20b
	ExifIFD_SpatialFrequencyResponse        ExifTag = 0xa20c
	ExifIFD_Noise_0xa20d                    ExifTag = 0xa20d
	ExifIFD_FocalPlaneXResolution           ExifTag = 0xa20e
	ExifIFD_FocalPlaneYResolution           ExifTag = 0xa20f
	ExifIFD_FocalPlaneResolutionUnit_0xa210 ExifTag = 0xa210
	ExifIFD_ImageNumber_0xa211              ExifTag = 0xa211
	ExifIFD_SecurityClassification_0xa212   ExifTag = 0xa212
	ExifIFD_ImageHistory_0xa213             ExifTag = 0xa213
	ExifIFD_SubjectLocation                 ExifTag = 0xa214
	ExifIFD_ExposureIndex                   ExifTag = 0xa215
	ExifIFD_TIFFEPStandardID_0xa216         ExifTag = 0xa216
	ExifIFD_SensingMethod_0xa217            ExifTag = 0xa217
	ExifIFD_FileSource                      ExifTag = 0xa300
	ExifIFD_SceneType                       ExifTag = 0xa301
	ExifIFD_CFAPattern                      ExifTag = 0xa302
	ExifIFD_CustomRendered                  ExifTag = 0xa401
	ExifIFD_ExposureMode                    ExifTag = 0xa402
	ExifIFD_WhiteBalance                    ExifTag = 0xa403
	ExifIFD_DigitalZoomRatio                ExifTag = 0xa404
	ExifIFD_FocalLengthIn35mmFormat         ExifTag = 0xa405
	ExifIFD_SceneCaptureType                ExifTag = 0xa406
	ExifIFD_GainControl                     ExifTag = 0xa407
	ExifIFD_Contrast                        ExifTag = 0xa408
	ExifIFD_Saturation                      ExifTag = 0xa409
	ExifIFD_Sharpness                       ExifTag = 0xa40a
	ExifIFD_DeviceSettingDescription        ExifTag = 0xa40b
	ExifIFD_SubjectDistanceRange            ExifTag = 0xa40c
	ExifIFD_ImageUniqueID                   ExifTag = 0xa420
	ExifIFD_OwnerName                       ExifTag = 0xa430
	ExifIFD_SerialNumber                    ExifTag = 0xa431
	ExifIFD_LensInfo                        ExifTag = 0xa432
	ExifIFD_LensMake                        ExifTag = 0xa433
	ExifIFD_LensModel                       ExifTag = 0xa434
	ExifIFD_LensSerialNumber                ExifTag = 0xa435
	ExifIFD_ImageTitle                      ExifTag = 0xa436
	ExifIFD_Photographer                    ExifTag = 0xa437
	ExifIFD_ImageEditor                     ExifTag = 0xa438
	ExifIFD_CameraFirmware                  ExifTag = 0xa439
	ExifIFD_RAWDevelopingSoftware           ExifTag = 0xa43a
	ExifIFD_ImageEditingSoftware            ExifTag = 0xa43b
	ExifIFD_MetadataEditingSoftware         ExifTag = 0xa43c
	ExifIFD_CompositeImage                  ExifTag = 0xa460
	ExifIFD_CompositeImageCount             ExifTag = 0xa461
	ExifIFD_CompositeImageExposureTimes     ExifTag = 0xa462
	ExifIFD_Gamma                           ExifTag = 0xa500
	ExifIFD_ExpandSoftware                  ExifTag = 0xafc0
	ExifIFD_ExpandLens                      ExifTag = 0xafc1
	ExifIFD_ExpandFilm                      ExifTag = 0xafc2
	ExifIFD_ExpandFilterLens                ExifTag = 0xafc3
	ExifIFD_ExpandScanner                   ExifTag = 0xafc4
	ExifIFD_ExpandFlashLamp                 ExifTag = 0xafc5
	ExifIFD_HasselbladRawImage              ExifTag = 0xb4c3
	ExifIFD_PixelFormat                     ExifTag = 0xbc01
	ExifIFD_Transformation                  ExifTag = 0xbc02
	ExifIFD_Uncompressed                    ExifTag = 0xbc03
	ExifIFD_ImageType                       ExifTag = 0xbc04
	ExifIFD_ImageWidth                      ExifTag = 0xbc80
	ExifIFD_ImageHeight                     ExifTag = 0xbc81
	ExifIFD_WidthResolution                 ExifTag = 0xbc82
	ExifIFD_HeightResolution                ExifTag = 0xbc83
	ExifIFD_ImageOffset                     ExifTag = 0xbcc0
	ExifIFD_ImageByteCount                  ExifTag = 0xbcc1
	ExifIFD_AlphaOffset                     ExifTag = 0xbcc2
	ExifIFD_AlphaByteCount                  ExifTag = 0xbcc3
	ExifIFD_ImageDataDiscard                ExifTag = 0xbcc4
	ExifIFD_AlphaDataDiscard                ExifTag = 0xbcc5
	ExifIFD_OceScanjobDesc                  ExifTag = 0xc427
	ExifIFD_OceApplicationSelector          ExifTag = 0xc428
	ExifIFD_OceIDNumber                     ExifTag = 0xc429
	ExifIFD_OceImageLogic                   ExifTag = 0xc42a
	ExifIFD_Annotations                     ExifTag = 0xc44f
	ExifIFD_HasselbladExif                  ExifTag = 0xc51b
	ExifIFD_OriginalFileName                ExifTag = 0xc573
	ExifIFD_USPTOOriginalContentType        ExifTag = 0xc580
	ExifIFD_CR2CFAPattern                   ExifTag = 0xc5e0
	ExifIFD_RawImageSegmentation            ExifTag = 0xc640
	ExifIFD_AliasLayerMetadata              ExifTag = 0xc660
	ExifIFD_SubTileBlockSize                ExifTag = 0xc71e
	ExifIFD_RowInterleaveFactor             ExifTag = 0xc71f
	ExifIFD_NikonNEFInfo                    ExifTag = 0xc7d5
	ExifIFD_ZIFMetadata                     ExifTag = 0xc7d7
	ExifIFD_ZIFAnnotations                  ExifTag = 0xc7d8
	ExifIFD_Padding                         ExifTag = 0xea1c
	ExifIFD_OffsetSchema                    ExifTag = 0xea1d
	ExifIFD_OwnerName_0xfde8                ExifTag = 0xfde8
	ExifIFD_SerialNumber_0xfde9             ExifTag = 0xfde9
	ExifIFD_Lens                            ExifTag = 0xfdea
	ExifIFD_KDC_IFD                         ExifTag = 0xfe00
	ExifIFD_RawFile                         ExifTag = 0xfe4c
	ExifIFD_Converter                       ExifTag = 0xfe4d
	ExifIFD_WhiteBalance_0xfe4e             ExifTag = 0xfe4e
	ExifIFD_Exposure                        ExifTag = 0xfe51
	ExifIFD_Shadows                         ExifTag = 0xfe52
	ExifIFD_Brightness                      ExifTag = 0xfe53
	ExifIFD_Contrast_0xfe54                 ExifTag = 0xfe54
	ExifIFD_Saturation_0xfe55               ExifTag = 0xfe55
	ExifIFD_Sharpness_0xfe56                ExifTag = 0xfe56
	ExifIFD_Smoothness                      ExifTag = 0xfe57
	ExifIFD_MoireFilter                     ExifTag = 0xfe58
)

// InteropIFD Tag Ids
const (
	InteropIFD_InteropIndex           ExifTag = 0x0001
	InteropIFD_InteropVersion         ExifTag = 0x0002
	InteropIFD_RelatedImageFileFormat ExifTag = 0x1000
	InteropIFD_RelatedImageWidth      ExifTag = 0x1001
	InteropIFD_RelatedImageHeight     ExifTag = 0x1002
)

// GpsIFD Tag Ids
const (
	GpsIFD_GPSVersionID         ExifTag = 0x0000
	GpsIFD_GPSLatitudeRef       ExifTag = 0x0001
	GpsIFD_GPSLatitude          ExifTag = 0x0002
	GpsIFD_GPSLongitudeRef      ExifTag = 0x0003
	GpsIFD_GPSLongitude         ExifTag = 0x0004
	GpsIFD_GPSAltitudeRef       ExifTag = 0x0005
	GpsIFD_GPSAltitude          ExifTag = 0x0006
	GpsIFD_GPSTimeStamp         ExifTag = 0x0007
	GpsIFD_GPSSatellites        ExifTag = 0x0008
	GpsIFD_GPSStatus            ExifTag = 0x0009
	GpsIFD_GPSMeasureMode       ExifTag = 0x000a
	GpsIFD_GPSDOP               ExifTag = 0x000b
	GpsIFD_GPSSpeedRef          ExifTag = 0x000c
	GpsIFD_GPSSpeed             ExifTag = 0x000d
	GpsIFD_GPSTrackRef          ExifTag = 0x000e
	GpsIFD_GPSTrack             ExifTag = 0x000f
	GpsIFD_GPSImgDirectionRef   ExifTag = 0x0010
	GpsIFD_GPSImgDirection      ExifTag = 0x0011
	GpsIFD_GPSMapDatum          ExifTag = 0x0012
	GpsIFD_GPSDestLatitudeRef   ExifTag = 0x0013
	GpsIFD_GPSDestLatitude      ExifTag = 0x0014
	GpsIFD_GPSDestLongitudeRef  ExifTag = 0x0015
	GpsIFD_GPSDestLongitude     ExifTag = 0x0016
	GpsIFD_GPSDestBearingRef    ExifTag = 0x0017
	GpsIFD_GPSDestBearing       ExifTag = 0x0018
	GpsIFD_GPSDestDistanceRef   ExifTag = 0x0019
	GpsIFD_GPSDestDistance      ExifTag = 0x001a
	GpsIFD_GPSProcessingMethod  ExifTag = 0x001b
	GpsIFD_GPSAreaInformation   ExifTag = 0x001c
	GpsIFD_GPSDateStamp         ExifTag = 0x001d
	GpsIFD_GPSDifferential      ExifTag = 0x001e
	GpsIFD_GPSHPositioningError ExifTag = 0x001f
)

// Exif Tag Descriptions
var ExifTagDescriptions = map[ExifIndexTag]ExifTagDesc{
	ExifIndexTag{InteropIFD, 0x0001}: ExifTagDesc{
		Id:    0x0001,
		Name:  "InteropIndex",
		Type:  ExifString,
		Ifd:   InteropIFD,
		Count: -1,
		Values: map[string]string{
			"R03": "R03 - DCF option file (Adobe RGB)",
			"R98": "R98 - DCF basic file (sRGB)",
			"THM": "THM - DCF thumbnail file",
		},
	},
	ExifIndexTag{InteropIFD, 0x0002}: ExifTagDesc{
		Id:     0x0002,
		Name:   "InteropVersion",
		Type:   ExifUndef,
		Ifd:    InteropIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x000b}: ExifTagDesc{
		Id:     0x000b,
		Name:   "ProcessingSoftware",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x00fe}: ExifTagDesc{
		Id:    0x00fe,
		Name:  "SubfileType",
		Type:  ExifUint32,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint32]string{
			0:          "Full-resolution image",
			1:          "Reduced-resolution image",
			16:         "Enhanced image data",
			2:          "Single page of multi-page image",
			3:          "Single page of multi-page reduced-resolution image",
			4:          "Transparency mask",
			4294967295: "invalid",
			5:          "Transparency mask of reduced-resolution image",
			6:          "Transparency mask of multi-page image",
			65537:      "Alternate reduced-resolution image",
			65540:      "Semantic Mask",
			7:          "Transparency mask of reduced-resolution multi-page image",
			8:          "Depth map",
			9:          "Depth map of reduced-resolution image",
		},
	},
	ExifIndexTag{RootIFD, 0x00ff}: ExifTagDesc{
		Id:    0x00ff,
		Name:  "OldSubfileType",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "Full-resolution image",
			2: "Reduced-resolution image",
			3: "Single page of multi-page image",
		},
	},
	ExifIndexTag{RootIFD, 0x0100}: ExifTagDesc{
		Id:     0x0100,
		Name:   "ImageWidth",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0101}: ExifTagDesc{
		Id:     0x0101,
		Name:   "ImageHeight",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0102}: ExifTagDesc{
		Id:     0x0102,
		Name:   "BitsPerSample",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0103}: ExifTagDesc{
		Id:    0x0103,
		Name:  "Compression",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1:     "Uncompressed",
			10:    "JBIG Color",
			2:     "CCITT 1D",
			262:   "Kodak 262",
			3:     "T4/Group 3 Fax",
			32766: "NeXt or Sony ARW Compressed 2",
			32767: "Sony ARW Compressed",
			32769: "Packed RAW",
			32770: "Samsung SRW Compressed",
			32771: "CCIRLEW",
			32772: "Samsung SRW Compressed 2",
			32773: "PackBits",
			32809: "Thunderscan",
			32867: "Kodak KDC Compressed",
			32895: "IT8CTPAD",
			32896: "IT8LW",
			32897: "IT8MP",
			32898: "IT8BL",
			32908: "PixarFilm",
			32909: "PixarLog",
			32946: "Deflate",
			32947: "DCS",
			33003: "Aperio JPEG 2000 YCbCr",
			33005: "Aperio JPEG 2000 RGB",
			34661: "JBIG",
			34676: "SGILog",
			34677: "SGILog24",
			34712: "JPEG 2000",
			34713: "Nikon NEF Compressed",
			34715: "JBIG2 TIFF FX",
			34718: "Microsoft Document Imaging (MDI) Binary Level Codec",
			34719: "Microsoft Document Imaging (MDI) Progressive Transform Codec",
			34720: "Microsoft Document Imaging (MDI) Vector",
			34887: "ESRI Lerc",
			34892: "Lossy JPEG",
			34925: "LZMA2",
			34926: "Zstd (old)",
			34927: "WebP (old)",
			34933: "PNG",
			34934: "JPEG XR",
			4:     "T6/Group 4 Fax",
			5:     "LZW",
			50000: "Zstd",
			50001: "WebP",
			50002: "JPEG XL (old)",
			52546: "JPEG XL",
			6:     "JPEG (old-style)",
			65000: "Kodak DCR Compressed",
			65535: "Pentax PEF Compressed",
			7:     "JPEG",
			8:     "Adobe Deflate",
			9:     "JBIG B&W",
			99:    "JPEG",
		},
	},
	ExifIndexTag{RootIFD, 0x0106}: ExifTagDesc{
		Id:    0x0106,
		Name:  "PhotometricInterpretation",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0:     "WhiteIsZero",
			1:     "BlackIsZero",
			10:    "ITULab",
			2:     "RGB",
			3:     "RGB Palette",
			32803: "Color Filter Array",
			32844: "Pixar LogL",
			32845: "Pixar LogLuv",
			32892: "Sequential Color Filter",
			34892: "Linear Raw",
			4:     "Transparency Mask",
			5:     "CMYK",
			51177: "Depth Map",
			52527: "Semantic Mask",
			6:     "YCbCr",
			8:     "CIELab",
			9:     "ICCLab",
		},
	},
	ExifIndexTag{RootIFD, 0x0107}: ExifTagDesc{
		Id:    0x0107,
		Name:  "Thresholding",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "No dithering or halftoning",
			2: "Ordered dither or halftone",
			3: "Randomized dither",
		},
	},
	ExifIndexTag{RootIFD, 0x0108}: ExifTagDesc{
		Id:     0x0108,
		Name:   "CellWidth",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0109}: ExifTagDesc{
		Id:     0x0109,
		Name:   "CellLength",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x010a}: ExifTagDesc{
		Id:    0x010a,
		Name:  "FillOrder",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "Normal",
			2: "Reversed",
		},
	},
	ExifIndexTag{RootIFD, 0x010d}: ExifTagDesc{
		Id:     0x010d,
		Name:   "DocumentName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x010e}: ExifTagDesc{
		Id:     0x010e,
		Name:   "ImageDescription",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x010f}: ExifTagDesc{
		Id:     0x010f,
		Name:   "Make",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0110}: ExifTagDesc{
		Id:     0x0110,
		Name:   "Model",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0111}: ExifTagDesc{
		Id:     0x0111,
		Name:   "StripOffsets",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0111}: ExifTagDesc{
		Id:     0x0111,
		Name:   "StripOffsets",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0112}: ExifTagDesc{
		Id:    0x0112,
		Name:  "Orientation",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "Horizontal (normal)",
			2: "Mirror horizontal",
			3: "Rotate 180",
			4: "Mirror vertical",
			5: "Mirror horizontal and rotate 270 CW",
			6: "Rotate 90 CW",
			7: "Mirror horizontal and rotate 90 CW",
			8: "Rotate 270 CW",
		},
	},
	ExifIndexTag{RootIFD, 0x0115}: ExifTagDesc{
		Id:     0x0115,
		Name:   "SamplesPerPixel",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0116}: ExifTagDesc{
		Id:     0x0116,
		Name:   "RowsPerStrip",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0117}: ExifTagDesc{
		Id:     0x0117,
		Name:   "StripByteCounts",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0117}: ExifTagDesc{
		Id:     0x0117,
		Name:   "StripByteCounts",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0118}: ExifTagDesc{
		Id:     0x0118,
		Name:   "MinSampleValue",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0119}: ExifTagDesc{
		Id:     0x0119,
		Name:   "MaxSampleValue",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x011a}: ExifTagDesc{
		Id:     0x011a,
		Name:   "XResolution",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x011b}: ExifTagDesc{
		Id:     0x011b,
		Name:   "YResolution",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x011c}: ExifTagDesc{
		Id:    0x011c,
		Name:  "PlanarConfiguration",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "Chunky",
			2: "Planar",
		},
	},
	ExifIndexTag{RootIFD, 0x011d}: ExifTagDesc{
		Id:     0x011d,
		Name:   "PageName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x011e}: ExifTagDesc{
		Id:     0x011e,
		Name:   "XPosition",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x011f}: ExifTagDesc{
		Id:     0x011f,
		Name:   "YPosition",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0120}: ExifTagDesc{
		Id:     0x0120,
		Name:   "FreeOffsets",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0121}: ExifTagDesc{
		Id:     0x0121,
		Name:   "FreeByteCounts",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0122}: ExifTagDesc{
		Id:    0x0122,
		Name:  "GrayResponseUnit",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "0.1",
			2: "0.001",
			3: "0.0001",
			4: "1e-05",
			5: "1e-06",
		},
	},
	ExifIndexTag{ExifIFD, 0x0123}: ExifTagDesc{
		Id:     0x0123,
		Name:   "GrayResponseCurve",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0124}: ExifTagDesc{
		Id:     0x0124,
		Name:   "T4Options",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0125}: ExifTagDesc{
		Id:     0x0125,
		Name:   "T6Options",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0128}: ExifTagDesc{
		Id:    0x0128,
		Name:  "ResolutionUnit",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "None",
			2: "inches",
			3: "cm",
		},
	},
	ExifIndexTag{RootIFD, 0x0129}: ExifTagDesc{
		Id:     0x0129,
		Name:   "PageNumber",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x012c}: ExifTagDesc{
		Id:     0x012c,
		Name:   "ColorResponseUnit",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x012d}: ExifTagDesc{
		Id:     0x012d,
		Name:   "TransferFunction",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  768,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0131}: ExifTagDesc{
		Id:     0x0131,
		Name:   "Software",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0132}: ExifTagDesc{
		Id:     0x0132,
		Name:   "ModifyDate",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x013b}: ExifTagDesc{
		Id:     0x013b,
		Name:   "Artist",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x013c}: ExifTagDesc{
		Id:     0x013c,
		Name:   "HostComputer",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x013d}: ExifTagDesc{
		Id:    0x013d,
		Name:  "Predictor",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1:     "None",
			2:     "Horizontal differencing",
			3:     "Floating point",
			34892: "Horizontal difference X2",
			34893: "Horizontal difference X4",
			34894: "Floating point X2",
			34895: "Floating point X4",
		},
	},
	ExifIndexTag{RootIFD, 0x013e}: ExifTagDesc{
		Id:     0x013e,
		Name:   "WhitePoint",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x013f}: ExifTagDesc{
		Id:     0x013f,
		Name:   "PrimaryChromaticities",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  6,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0140}: ExifTagDesc{
		Id:     0x0140,
		Name:   "ColorMap",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0141}: ExifTagDesc{
		Id:     0x0141,
		Name:   "HalftoneHints",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0142}: ExifTagDesc{
		Id:     0x0142,
		Name:   "TileWidth",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0143}: ExifTagDesc{
		Id:     0x0143,
		Name:   "TileLength",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0144}: ExifTagDesc{
		Id:     0x0144,
		Name:   "TileOffsets",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0145}: ExifTagDesc{
		Id:     0x0145,
		Name:   "TileByteCounts",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0146}: ExifTagDesc{
		Id:     0x0146,
		Name:   "BadFaxLines",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0147}: ExifTagDesc{
		Id:     0x0147,
		Name:   "CleanFaxData",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0148}: ExifTagDesc{
		Id:     0x0148,
		Name:   "ConsecutiveBadFaxLines",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x014a}: ExifTagDesc{
		Id:     0x014a,
		Name:   "SubIFDs",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x014c}: ExifTagDesc{
		Id:    0x014c,
		Name:  "InkSet",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "CMYK",
			2: "Not CMYK",
		},
	},
	ExifIndexTag{ExifIFD, 0x014d}: ExifTagDesc{
		Id:     0x014d,
		Name:   "InkNames",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x014e}: ExifTagDesc{
		Id:     0x014e,
		Name:   "NumberofInks",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0150}: ExifTagDesc{
		Id:     0x0150,
		Name:   "DotRange",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0151}: ExifTagDesc{
		Id:     0x0151,
		Name:   "TargetPrinter",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0152}: ExifTagDesc{
		Id:     0x0152,
		Name:   "ExtraSamples",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0153}: ExifTagDesc{
		Id:    0x0153,
		Name:  "SampleFormat",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "Unsigned",
			2: "Signed",
			3: "Float",
			4: "Undefined",
			5: "Complex int",
			6: "Complex float",
		},
	},
	ExifIndexTag{ExifIFD, 0x0154}: ExifTagDesc{
		Id:     0x0154,
		Name:   "SMinSampleValue",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0155}: ExifTagDesc{
		Id:     0x0155,
		Name:   "SMaxSampleValue",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0156}: ExifTagDesc{
		Id:     0x0156,
		Name:   "TransferRange",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0157}: ExifTagDesc{
		Id:     0x0157,
		Name:   "ClipPath",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0158}: ExifTagDesc{
		Id:     0x0158,
		Name:   "XClipPathUnits",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0159}: ExifTagDesc{
		Id:     0x0159,
		Name:   "YClipPathUnits",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x015a}: ExifTagDesc{
		Id:     0x015a,
		Name:   "Indexed",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x015b}: ExifTagDesc{
		Id:     0x015b,
		Name:   "JPEGTables",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x015f}: ExifTagDesc{
		Id:     0x015f,
		Name:   "OPIProxy",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0190}: ExifTagDesc{
		Id:     0x0190,
		Name:   "GlobalParametersIFD",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0191}: ExifTagDesc{
		Id:     0x0191,
		Name:   "ProfileType",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0192}: ExifTagDesc{
		Id:     0x0192,
		Name:   "FaxProfile",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0193}: ExifTagDesc{
		Id:     0x0193,
		Name:   "CodingMethods",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0194}: ExifTagDesc{
		Id:     0x0194,
		Name:   "VersionYear",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0195}: ExifTagDesc{
		Id:     0x0195,
		Name:   "ModeNumber",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x01b1}: ExifTagDesc{
		Id:     0x01b1,
		Name:   "Decode",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x01b2}: ExifTagDesc{
		Id:     0x01b2,
		Name:   "DefaultImageColor",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x01b3}: ExifTagDesc{
		Id:     0x01b3,
		Name:   "T82Options",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x01b5}: ExifTagDesc{
		Id:     0x01b5,
		Name:   "JPEGTables_0x01b5",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0200}: ExifTagDesc{
		Id:     0x0200,
		Name:   "JPEGProc",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0201}: ExifTagDesc{
		Id:     0x0201,
		Name:   "ThumbnailOffset",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0202}: ExifTagDesc{
		Id:     0x0202,
		Name:   "ThumbnailLength",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0203}: ExifTagDesc{
		Id:     0x0203,
		Name:   "JPEGRestartInterval",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0205}: ExifTagDesc{
		Id:     0x0205,
		Name:   "JPEGLosslessPredictors",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0206}: ExifTagDesc{
		Id:     0x0206,
		Name:   "JPEGPointTransforms",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0207}: ExifTagDesc{
		Id:     0x0207,
		Name:   "JPEGQTables",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0208}: ExifTagDesc{
		Id:     0x0208,
		Name:   "JPEGDCTables",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0209}: ExifTagDesc{
		Id:     0x0209,
		Name:   "JPEGACTables",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0211}: ExifTagDesc{
		Id:     0x0211,
		Name:   "YCbCrCoefficients",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0212}: ExifTagDesc{
		Id:     0x0212,
		Name:   "YCbCrSubSampling",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x0213}: ExifTagDesc{
		Id:    0x0213,
		Name:  "YCbCrPositioning",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "Centered",
			2: "Co-sited",
		},
	},
	ExifIndexTag{RootIFD, 0x0214}: ExifTagDesc{
		Id:     0x0214,
		Name:   "ReferenceBlackWhite",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  6,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x022f}: ExifTagDesc{
		Id:     0x022f,
		Name:   "StripRowCounts",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x02bc}: ExifTagDesc{
		Id:     0x02bc,
		Name:   "ApplicationNotes",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x0303}: ExifTagDesc{
		Id:    0x0303,
		Name:  "RenderingIntent",
		Type:  ExifUint8,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint8]string{
			0: "Perceptual",
			1: "Relative Colorimetric",
			2: "Saturation",
			3: "Absolute colorimetric",
		},
	},
	ExifIndexTag{ExifIFD, 0x03e7}: ExifTagDesc{
		Id:     0x03e7,
		Name:   "USPTOMiscellaneous",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{InteropIFD, 0x1000}: ExifTagDesc{
		Id:     0x1000,
		Name:   "RelatedImageFileFormat",
		Type:   ExifString,
		Ifd:    InteropIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{InteropIFD, 0x1001}: ExifTagDesc{
		Id:     0x1001,
		Name:   "RelatedImageWidth",
		Type:   ExifUint16,
		Ifd:    InteropIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{InteropIFD, 0x1002}: ExifTagDesc{
		Id:     0x1002,
		Name:   "RelatedImageHeight",
		Type:   ExifUint16,
		Ifd:    InteropIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x4746}: ExifTagDesc{
		Id:     0x4746,
		Name:   "Rating",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x4747}: ExifTagDesc{
		Id:     0x4747,
		Name:   "XP_DIP_XML",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x4748}: ExifTagDesc{
		Id:     0x4748,
		Name:   "StitchInfo",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x4749}: ExifTagDesc{
		Id:     0x4749,
		Name:   "RatingPercent",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5001}: ExifTagDesc{
		Id:     0x5001,
		Name:   "ResolutionXUnit",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5002}: ExifTagDesc{
		Id:     0x5002,
		Name:   "ResolutionYUnit",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5003}: ExifTagDesc{
		Id:     0x5003,
		Name:   "ResolutionXLengthUnit",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5004}: ExifTagDesc{
		Id:     0x5004,
		Name:   "ResolutionYLengthUnit",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5005}: ExifTagDesc{
		Id:     0x5005,
		Name:   "PrintFlags",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5006}: ExifTagDesc{
		Id:     0x5006,
		Name:   "PrintFlagsVersion",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5007}: ExifTagDesc{
		Id:     0x5007,
		Name:   "PrintFlagsCrop",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5008}: ExifTagDesc{
		Id:     0x5008,
		Name:   "PrintFlagsBleedWidth",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5009}: ExifTagDesc{
		Id:     0x5009,
		Name:   "PrintFlagsBleedWidthScale",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x500a}: ExifTagDesc{
		Id:     0x500a,
		Name:   "HalftoneLPI",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x500b}: ExifTagDesc{
		Id:     0x500b,
		Name:   "HalftoneLPIUnit",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x500c}: ExifTagDesc{
		Id:     0x500c,
		Name:   "HalftoneDegree",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x500d}: ExifTagDesc{
		Id:     0x500d,
		Name:   "HalftoneShape",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x500e}: ExifTagDesc{
		Id:     0x500e,
		Name:   "HalftoneMisc",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x500f}: ExifTagDesc{
		Id:     0x500f,
		Name:   "HalftoneScreen",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5010}: ExifTagDesc{
		Id:     0x5010,
		Name:   "JPEGQuality",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5011}: ExifTagDesc{
		Id:     0x5011,
		Name:   "GridSize",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5012}: ExifTagDesc{
		Id:     0x5012,
		Name:   "ThumbnailFormat",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5013}: ExifTagDesc{
		Id:     0x5013,
		Name:   "ThumbnailWidth",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5014}: ExifTagDesc{
		Id:     0x5014,
		Name:   "ThumbnailHeight",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5015}: ExifTagDesc{
		Id:     0x5015,
		Name:   "ThumbnailColorDepth",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5016}: ExifTagDesc{
		Id:     0x5016,
		Name:   "ThumbnailPlanes",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5017}: ExifTagDesc{
		Id:     0x5017,
		Name:   "ThumbnailRawBytes",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5018}: ExifTagDesc{
		Id:     0x5018,
		Name:   "ThumbnailLength",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5019}: ExifTagDesc{
		Id:     0x5019,
		Name:   "ThumbnailCompressedSize",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x501a}: ExifTagDesc{
		Id:     0x501a,
		Name:   "ColorTransferFunction",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x501b}: ExifTagDesc{
		Id:     0x501b,
		Name:   "ThumbnailData",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5020}: ExifTagDesc{
		Id:     0x5020,
		Name:   "ThumbnailImageWidth",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5021}: ExifTagDesc{
		Id:     0x5021,
		Name:   "ThumbnailImageHeight",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5022}: ExifTagDesc{
		Id:     0x5022,
		Name:   "ThumbnailBitsPerSample",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5023}: ExifTagDesc{
		Id:     0x5023,
		Name:   "ThumbnailCompression",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5024}: ExifTagDesc{
		Id:     0x5024,
		Name:   "ThumbnailPhotometricInterp",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5025}: ExifTagDesc{
		Id:     0x5025,
		Name:   "ThumbnailDescription",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5026}: ExifTagDesc{
		Id:     0x5026,
		Name:   "ThumbnailEquipMake",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5027}: ExifTagDesc{
		Id:     0x5027,
		Name:   "ThumbnailEquipModel",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5028}: ExifTagDesc{
		Id:     0x5028,
		Name:   "ThumbnailStripOffsets",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5029}: ExifTagDesc{
		Id:     0x5029,
		Name:   "ThumbnailOrientation",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x502a}: ExifTagDesc{
		Id:     0x502a,
		Name:   "ThumbnailSamplesPerPixel",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x502b}: ExifTagDesc{
		Id:     0x502b,
		Name:   "ThumbnailRowsPerStrip",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x502c}: ExifTagDesc{
		Id:     0x502c,
		Name:   "ThumbnailStripByteCounts",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x502d}: ExifTagDesc{
		Id:     0x502d,
		Name:   "ThumbnailResolutionX",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x502e}: ExifTagDesc{
		Id:     0x502e,
		Name:   "ThumbnailResolutionY",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x502f}: ExifTagDesc{
		Id:     0x502f,
		Name:   "ThumbnailPlanarConfig",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5030}: ExifTagDesc{
		Id:     0x5030,
		Name:   "ThumbnailResolutionUnit",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5031}: ExifTagDesc{
		Id:     0x5031,
		Name:   "ThumbnailTransferFunction",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5032}: ExifTagDesc{
		Id:     0x5032,
		Name:   "ThumbnailSoftware",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5033}: ExifTagDesc{
		Id:     0x5033,
		Name:   "ThumbnailDateTime",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5034}: ExifTagDesc{
		Id:     0x5034,
		Name:   "ThumbnailArtist",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5035}: ExifTagDesc{
		Id:     0x5035,
		Name:   "ThumbnailWhitePoint",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5036}: ExifTagDesc{
		Id:     0x5036,
		Name:   "ThumbnailPrimaryChromaticities",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5037}: ExifTagDesc{
		Id:     0x5037,
		Name:   "ThumbnailYCbCrCoefficients",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5038}: ExifTagDesc{
		Id:     0x5038,
		Name:   "ThumbnailYCbCrSubsampling",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5039}: ExifTagDesc{
		Id:     0x5039,
		Name:   "ThumbnailYCbCrPositioning",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x503a}: ExifTagDesc{
		Id:     0x503a,
		Name:   "ThumbnailRefBlackWhite",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x503b}: ExifTagDesc{
		Id:     0x503b,
		Name:   "ThumbnailCopyright",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5090}: ExifTagDesc{
		Id:     0x5090,
		Name:   "LuminanceTable",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5091}: ExifTagDesc{
		Id:     0x5091,
		Name:   "ChrominanceTable",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5100}: ExifTagDesc{
		Id:     0x5100,
		Name:   "FrameDelay",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5101}: ExifTagDesc{
		Id:     0x5101,
		Name:   "LoopCount",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5102}: ExifTagDesc{
		Id:     0x5102,
		Name:   "GlobalPalette",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5103}: ExifTagDesc{
		Id:     0x5103,
		Name:   "IndexBackground",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5104}: ExifTagDesc{
		Id:     0x5104,
		Name:   "IndexTransparent",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5110}: ExifTagDesc{
		Id:     0x5110,
		Name:   "PixelUnits",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5111}: ExifTagDesc{
		Id:     0x5111,
		Name:   "PixelsPerUnitX",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5112}: ExifTagDesc{
		Id:     0x5112,
		Name:   "PixelsPerUnitY",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x5113}: ExifTagDesc{
		Id:     0x5113,
		Name:   "PaletteHistogram",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x7000}: ExifTagDesc{
		Id:     0x7000,
		Name:   "SonyRawFileType",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x7010}: ExifTagDesc{
		Id:     0x7010,
		Name:   "SonyToneCurve",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x7031}: ExifTagDesc{
		Id:    0x7031,
		Name:  "VignettingCorrection",
		Type:  ExifInt16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[int16]string{
			256: "Off",
			257: "Auto",
			272: "Auto (ILCE-1)",
			511: "No correction params available",
		},
	},
	ExifIndexTag{RootIFD, 0x7032}: ExifTagDesc{
		Id:     0x7032,
		Name:   "VignettingCorrParams",
		Type:   ExifInt16,
		Ifd:    RootIFD,
		Count:  17,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x7034}: ExifTagDesc{
		Id:    0x7034,
		Name:  "ChromaticAberrationCorrection",
		Type:  ExifInt16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[int16]string{
			0:   "Off",
			1:   "Auto",
			255: "No correction params available",
		},
	},
	ExifIndexTag{RootIFD, 0x7035}: ExifTagDesc{
		Id:     0x7035,
		Name:   "ChromaticAberrationCorrParams",
		Type:   ExifInt16,
		Ifd:    RootIFD,
		Count:  33,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x7036}: ExifTagDesc{
		Id:    0x7036,
		Name:  "DistortionCorrection",
		Type:  ExifInt16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[int16]string{
			0:   "Off",
			1:   "Auto",
			17:  "Auto fixed by lens",
			255: "No correction params available",
		},
	},
	ExifIndexTag{RootIFD, 0x7037}: ExifTagDesc{
		Id:     0x7037,
		Name:   "DistortionCorrParams",
		Type:   ExifInt16,
		Ifd:    RootIFD,
		Count:  17,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x7038}: ExifTagDesc{
		Id:     0x7038,
		Name:   "SonyRawImageSize",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x7310}: ExifTagDesc{
		Id:     0x7310,
		Name:   "BlackLevel_0x7310",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x7313}: ExifTagDesc{
		Id:     0x7313,
		Name:   "WB_RGGBLevels",
		Type:   ExifInt16,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x74c7}: ExifTagDesc{
		Id:     0x74c7,
		Name:   "SonyCropTopLeft",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x74c8}: ExifTagDesc{
		Id:     0x74c8,
		Name:   "SonyCropSize",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x800d}: ExifTagDesc{
		Id:     0x800d,
		Name:   "ImageID",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80a3}: ExifTagDesc{
		Id:     0x80a3,
		Name:   "WangTag1",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80a4}: ExifTagDesc{
		Id:     0x80a4,
		Name:   "WangAnnotation",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80a5}: ExifTagDesc{
		Id:     0x80a5,
		Name:   "WangTag3",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80a6}: ExifTagDesc{
		Id:     0x80a6,
		Name:   "WangTag4",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80b9}: ExifTagDesc{
		Id:     0x80b9,
		Name:   "ImageReferencePoints",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80ba}: ExifTagDesc{
		Id:     0x80ba,
		Name:   "RegionXformTackPoint",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80bb}: ExifTagDesc{
		Id:     0x80bb,
		Name:   "WarpQuadrilateral",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80bc}: ExifTagDesc{
		Id:     0x80bc,
		Name:   "AffineTransformMat",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80e3}: ExifTagDesc{
		Id:     0x80e3,
		Name:   "Matteing",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80e4}: ExifTagDesc{
		Id:     0x80e4,
		Name:   "DataType",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80e5}: ExifTagDesc{
		Id:     0x80e5,
		Name:   "ImageDepth",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x80e6}: ExifTagDesc{
		Id:     0x80e6,
		Name:   "TileDepth",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8214}: ExifTagDesc{
		Id:     0x8214,
		Name:   "ImageFullWidth",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8215}: ExifTagDesc{
		Id:     0x8215,
		Name:   "ImageFullHeight",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8216}: ExifTagDesc{
		Id:     0x8216,
		Name:   "TextureFormat",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8217}: ExifTagDesc{
		Id:     0x8217,
		Name:   "WrapModes",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8218}: ExifTagDesc{
		Id:     0x8218,
		Name:   "FovCot",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8219}: ExifTagDesc{
		Id:     0x8219,
		Name:   "MatrixWorldToScreen",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x821a}: ExifTagDesc{
		Id:     0x821a,
		Name:   "MatrixWorldToCamera",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x827d}: ExifTagDesc{
		Id:     0x827d,
		Name:   "Model2",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x828d}: ExifTagDesc{
		Id:     0x828d,
		Name:   "CFARepeatPatternDim",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x828e}: ExifTagDesc{
		Id:     0x828e,
		Name:   "CFAPattern2",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x828f}: ExifTagDesc{
		Id:     0x828f,
		Name:   "BatteryLevel",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8290}: ExifTagDesc{
		Id:     0x8290,
		Name:   "KodakIFD",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x8298}: ExifTagDesc{
		Id:     0x8298,
		Name:   "Copyright",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x829a}: ExifTagDesc{
		Id:     0x829a,
		Name:   "ExposureTime",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x829d}: ExifTagDesc{
		Id:     0x829d,
		Name:   "FNumber",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x82a5}: ExifTagDesc{
		Id:     0x82a5,
		Name:   "MDFileTag",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x82a6}: ExifTagDesc{
		Id:     0x82a6,
		Name:   "MDScalePixel",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x82a7}: ExifTagDesc{
		Id:     0x82a7,
		Name:   "MDColorTable",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x82a8}: ExifTagDesc{
		Id:     0x82a8,
		Name:   "MDLabName",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x82a9}: ExifTagDesc{
		Id:     0x82a9,
		Name:   "MDSampleInfo",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x82aa}: ExifTagDesc{
		Id:     0x82aa,
		Name:   "MDPrepDate",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x82ab}: ExifTagDesc{
		Id:     0x82ab,
		Name:   "MDPrepTime",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x82ac}: ExifTagDesc{
		Id:     0x82ac,
		Name:   "MDFileUnits",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x830e}: ExifTagDesc{
		Id:     0x830e,
		Name:   "PixelScale",
		Type:   ExifDouble,
		Ifd:    RootIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8335}: ExifTagDesc{
		Id:     0x8335,
		Name:   "AdventScale",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8336}: ExifTagDesc{
		Id:     0x8336,
		Name:   "AdventRevision",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x835c}: ExifTagDesc{
		Id:     0x835c,
		Name:   "UIC1Tag",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x835d}: ExifTagDesc{
		Id:     0x835d,
		Name:   "UIC2Tag",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x835e}: ExifTagDesc{
		Id:     0x835e,
		Name:   "UIC3Tag",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x835f}: ExifTagDesc{
		Id:     0x835f,
		Name:   "UIC4Tag",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x83bb}: ExifTagDesc{
		Id:     0x83bb,
		Name:   "IPTCNAA",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x847e}: ExifTagDesc{
		Id:     0x847e,
		Name:   "IntergraphPacketData",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x847f}: ExifTagDesc{
		Id:     0x847f,
		Name:   "IntergraphFlagRegisters",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x8480}: ExifTagDesc{
		Id:     0x8480,
		Name:   "IntergraphMatrix",
		Type:   ExifDouble,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8481}: ExifTagDesc{
		Id:     0x8481,
		Name:   "INGRReserved",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x8482}: ExifTagDesc{
		Id:     0x8482,
		Name:   "ModelTiePoint",
		Type:   ExifDouble,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e0}: ExifTagDesc{
		Id:     0x84e0,
		Name:   "Site",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e1}: ExifTagDesc{
		Id:     0x84e1,
		Name:   "ColorSequence",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e2}: ExifTagDesc{
		Id:     0x84e2,
		Name:   "IT8Header",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e3}: ExifTagDesc{
		Id:     0x84e3,
		Name:   "RasterPadding",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e4}: ExifTagDesc{
		Id:     0x84e4,
		Name:   "BitsPerRunLength",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e5}: ExifTagDesc{
		Id:     0x84e5,
		Name:   "BitsPerExtendedRunLength",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e6}: ExifTagDesc{
		Id:     0x84e6,
		Name:   "ColorTable",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e7}: ExifTagDesc{
		Id:     0x84e7,
		Name:   "ImageColorIndicator",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e8}: ExifTagDesc{
		Id:     0x84e8,
		Name:   "BackgroundColorIndicator",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84e9}: ExifTagDesc{
		Id:     0x84e9,
		Name:   "ImageColorValue",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84ea}: ExifTagDesc{
		Id:     0x84ea,
		Name:   "BackgroundColorValue",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84eb}: ExifTagDesc{
		Id:     0x84eb,
		Name:   "PixelIntensityRange",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84ec}: ExifTagDesc{
		Id:     0x84ec,
		Name:   "TransparencyIndicator",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84ed}: ExifTagDesc{
		Id:     0x84ed,
		Name:   "ColorCharacterization",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84ee}: ExifTagDesc{
		Id:     0x84ee,
		Name:   "HCUsage",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84ef}: ExifTagDesc{
		Id:     0x84ef,
		Name:   "TrapIndicator",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x84f0}: ExifTagDesc{
		Id:     0x84f0,
		Name:   "CMYKEquivalent",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x8546}: ExifTagDesc{
		Id:     0x8546,
		Name:   "SEMInfo",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8568}: ExifTagDesc{
		Id:     0x8568,
		Name:   "AFCP_IPTC",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x85b8}: ExifTagDesc{
		Id:     0x85b8,
		Name:   "PixelMagicJBIGOptions",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x85d7}: ExifTagDesc{
		Id:     0x85d7,
		Name:   "JPLCartoIFD",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x85d8}: ExifTagDesc{
		Id:     0x85d8,
		Name:   "ModelTransform",
		Type:   ExifDouble,
		Ifd:    RootIFD,
		Count:  16,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8602}: ExifTagDesc{
		Id:     0x8602,
		Name:   "WB_GRGBLevels",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8606}: ExifTagDesc{
		Id:     0x8606,
		Name:   "LeafData",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x8649}: ExifTagDesc{
		Id:     0x8649,
		Name:   "PhotoshopSettings",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x8769}: ExifTagDesc{
		Id:     0x8769,
		Name:   "ExifOffset",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x8773}: ExifTagDesc{
		Id:     0x8773,
		Name:   "ICC_Profile",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x877f}: ExifTagDesc{
		Id:     0x877f,
		Name:   "TIFF_FXExtensions",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8780}: ExifTagDesc{
		Id:     0x8780,
		Name:   "MultiProfiles",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8781}: ExifTagDesc{
		Id:     0x8781,
		Name:   "SharedData",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8782}: ExifTagDesc{
		Id:     0x8782,
		Name:   "T88Options",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x87ac}: ExifTagDesc{
		Id:     0x87ac,
		Name:   "ImageLayer",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x87af}: ExifTagDesc{
		Id:     0x87af,
		Name:   "GeoTiffDirectory",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x87b0}: ExifTagDesc{
		Id:     0x87b0,
		Name:   "GeoTiffDoubleParams",
		Type:   ExifDouble,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x87b1}: ExifTagDesc{
		Id:     0x87b1,
		Name:   "GeoTiffAsciiParams",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x87be}: ExifTagDesc{
		Id:     0x87be,
		Name:   "JBIGOptions",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8822}: ExifTagDesc{
		Id:    0x8822,
		Name:  "ExposureProgram",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Not Defined",
			1: "Manual",
			2: "Program AE",
			3: "Aperture-priority AE",
			4: "Shutter speed priority AE",
			5: "Creative (Slow speed)",
			6: "Action (High speed)",
			7: "Portrait",
			8: "Landscape",
			9: "Bulb",
		},
	},
	ExifIndexTag{ExifIFD, 0x8824}: ExifTagDesc{
		Id:     0x8824,
		Name:   "SpectralSensitivity",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x8825}: ExifTagDesc{
		Id:     0x8825,
		Name:   "GPSInfo",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8827}: ExifTagDesc{
		Id:     0x8827,
		Name:   "ISO",
		Type:   ExifUint16,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8828}: ExifTagDesc{
		Id:     0x8828,
		Name:   "OptoElectricConvFactor",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8829}: ExifTagDesc{
		Id:     0x8829,
		Name:   "Interlace",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x882a}: ExifTagDesc{
		Id:     0x882a,
		Name:   "TimeZoneOffset",
		Type:   ExifInt16,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x882b}: ExifTagDesc{
		Id:     0x882b,
		Name:   "SelfTimerMode",
		Type:   ExifUint16,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8830}: ExifTagDesc{
		Id:    0x8830,
		Name:  "SensitivityType",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Unknown",
			1: "Standard Output Sensitivity",
			2: "Recommended Exposure Index",
			3: "ISO Speed",
			4: "Standard Output Sensitivity and Recommended Exposure Index",
			5: "Standard Output Sensitivity and ISO Speed",
			6: "Recommended Exposure Index and ISO Speed",
			7: "Standard Output Sensitivity, Recommended Exposure Index and ISO Speed",
		},
	},
	ExifIndexTag{ExifIFD, 0x8831}: ExifTagDesc{
		Id:     0x8831,
		Name:   "StandardOutputSensitivity",
		Type:   ExifUint32,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8832}: ExifTagDesc{
		Id:     0x8832,
		Name:   "RecommendedExposureIndex",
		Type:   ExifUint32,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8833}: ExifTagDesc{
		Id:     0x8833,
		Name:   "ISOSpeed",
		Type:   ExifUint32,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8834}: ExifTagDesc{
		Id:     0x8834,
		Name:   "ISOSpeedLatitudeyyy",
		Type:   ExifUint32,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8835}: ExifTagDesc{
		Id:     0x8835,
		Name:   "ISOSpeedLatitudezzz",
		Type:   ExifUint32,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x885c}: ExifTagDesc{
		Id:     0x885c,
		Name:   "FaxRecvParams",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x885d}: ExifTagDesc{
		Id:     0x885d,
		Name:   "FaxSubAddress",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x885e}: ExifTagDesc{
		Id:     0x885e,
		Name:   "FaxRecvTime",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x8871}: ExifTagDesc{
		Id:     0x8871,
		Name:   "FedexEDR",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x888a}: ExifTagDesc{
		Id:     0x888a,
		Name:   "LeafSubIFD",
		Type:   ExifUint32,
		Ifd:    ExifIFD,
		Count:  1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9000}: ExifTagDesc{
		Id:     0x9000,
		Name:   "ExifVersion",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9003}: ExifTagDesc{
		Id:     0x9003,
		Name:   "DateTimeOriginal",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9004}: ExifTagDesc{
		Id:     0x9004,
		Name:   "CreateDate",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9009}: ExifTagDesc{
		Id:     0x9009,
		Name:   "GooglePlusUploadCode",
		Type:   ExifUint8,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9010}: ExifTagDesc{
		Id:     0x9010,
		Name:   "OffsetTime",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9011}: ExifTagDesc{
		Id:     0x9011,
		Name:   "OffsetTimeOriginal",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9012}: ExifTagDesc{
		Id:     0x9012,
		Name:   "OffsetTimeDigitized",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9101}: ExifTagDesc{
		Id:     0x9101,
		Name:   "ComponentsConfiguration",
		Type:   ExifUint8,
		Ifd:    ExifIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9102}: ExifTagDesc{
		Id:     0x9102,
		Name:   "CompressedBitsPerPixel",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9201}: ExifTagDesc{
		Id:     0x9201,
		Name:   "ShutterSpeedValue",
		Type:   ExifRational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9202}: ExifTagDesc{
		Id:     0x9202,
		Name:   "ApertureValue",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9203}: ExifTagDesc{
		Id:     0x9203,
		Name:   "BrightnessValue",
		Type:   ExifRational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9204}: ExifTagDesc{
		Id:     0x9204,
		Name:   "ExposureCompensation",
		Type:   ExifRational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9205}: ExifTagDesc{
		Id:     0x9205,
		Name:   "MaxApertureValue",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9206}: ExifTagDesc{
		Id:     0x9206,
		Name:   "SubjectDistance",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9207}: ExifTagDesc{
		Id:    0x9207,
		Name:  "MeteringMode",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0:   "Unknown",
			1:   "Average",
			2:   "Center-weighted average",
			255: "Other",
			3:   "Spot",
			4:   "Multi-spot",
			5:   "Multi-segment",
			6:   "Partial",
		},
	},
	ExifIndexTag{ExifIFD, 0x9208}: ExifTagDesc{
		Id:    0x9208,
		Name:  "LightSource",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0:   "Unknown",
			1:   "Daylight",
			10:  "Cloudy",
			11:  "Shade",
			12:  "Daylight Fluorescent",
			13:  "Day White Fluorescent",
			14:  "Cool White Fluorescent",
			15:  "White Fluorescent",
			16:  "Warm White Fluorescent",
			17:  "Standard Light A",
			18:  "Standard Light B",
			19:  "Standard Light C",
			2:   "Fluorescent",
			20:  "D55",
			21:  "D65",
			22:  "D75",
			23:  "D50",
			24:  "ISO Studio Tungsten",
			255: "Other",
			3:   "Tungsten (Incandescent)",
			4:   "Flash",
			9:   "Fine Weather",
		},
	},
	ExifIndexTag{ExifIFD, 0x9209}: ExifTagDesc{
		Id:    0x9209,
		Name:  "Flash",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0:  "No Flash",
			1:  "Fired",
			13: "On, Return not detected",
			15: "On, Return detected",
			16: "Off, Did not fire",
			20: "Off, Did not fire, Return not detected",
			24: "Auto, Did not fire",
			25: "Auto, Fired",
			29: "Auto, Fired, Return not detected",
			31: "Auto, Fired, Return detected",
			32: "No flash function",
			48: "Off, No flash function",
			5:  "Fired, Return not detected",
			65: "Fired, Red-eye reduction",
			69: "Fired, Red-eye reduction, Return not detected",
			7:  "Fired, Return detected",
			71: "Fired, Red-eye reduction, Return detected",
			73: "On, Red-eye reduction",
			77: "On, Red-eye reduction, Return not detected",
			79: "On, Red-eye reduction, Return detected",
			8:  "On, Did not fire",
			80: "Off, Red-eye reduction",
			88: "Auto, Did not fire, Red-eye reduction",
			89: "Auto, Fired, Red-eye reduction",
			9:  "On, Fired",
			93: "Auto, Fired, Red-eye reduction, Return not detected",
			95: "Auto, Fired, Red-eye reduction, Return detected",
		},
	},
	ExifIndexTag{ExifIFD, 0x920a}: ExifTagDesc{
		Id:     0x920a,
		Name:   "FocalLength",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x920b}: ExifTagDesc{
		Id:     0x920b,
		Name:   "FlashEnergy",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x920c}: ExifTagDesc{
		Id:     0x920c,
		Name:   "SpatialFrequencyResponse_0x920c",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x920d}: ExifTagDesc{
		Id:     0x920d,
		Name:   "Noise",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x920e}: ExifTagDesc{
		Id:     0x920e,
		Name:   "FocalPlaneXResolution_0x920e",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x920f}: ExifTagDesc{
		Id:     0x920f,
		Name:   "FocalPlaneYResolution_0x920f",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9210}: ExifTagDesc{
		Id:     0x9210,
		Name:   "FocalPlaneResolutionUnit",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9211}: ExifTagDesc{
		Id:     0x9211,
		Name:   "ImageNumber",
		Type:   ExifUint32,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9212}: ExifTagDesc{
		Id:    0x9212,
		Name:  "SecurityClassification",
		Type:  ExifString,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[string]string{
			"C": "Confidential",
			"R": "Restricted",
			"S": "Secret",
			"T": "Top Secret",
			"U": "Unclassified",
		},
	},
	ExifIndexTag{ExifIFD, 0x9213}: ExifTagDesc{
		Id:     0x9213,
		Name:   "ImageHistory",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9214}: ExifTagDesc{
		Id:     0x9214,
		Name:   "SubjectArea",
		Type:   ExifUint16,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9215}: ExifTagDesc{
		Id:     0x9215,
		Name:   "ExposureIndex_0x9215",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9216}: ExifTagDesc{
		Id:     0x9216,
		Name:   "TIFFEPStandardID",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9217}: ExifTagDesc{
		Id:     0x9217,
		Name:   "SensingMethod",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x923a}: ExifTagDesc{
		Id:     0x923a,
		Name:   "CIP3DataFile",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x923b}: ExifTagDesc{
		Id:     0x923b,
		Name:   "CIP3Sheet",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x923c}: ExifTagDesc{
		Id:     0x923c,
		Name:   "CIP3Side",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x923f}: ExifTagDesc{
		Id:     0x923f,
		Name:   "StoNits",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x927c}: ExifTagDesc{
		Id:     0x927c,
		Name:   "MakerNote",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9286}: ExifTagDesc{
		Id:     0x9286,
		Name:   "UserComment",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9290}: ExifTagDesc{
		Id:     0x9290,
		Name:   "SubSecTime",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9291}: ExifTagDesc{
		Id:     0x9291,
		Name:   "SubSecTimeOriginal",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9292}: ExifTagDesc{
		Id:     0x9292,
		Name:   "SubSecTimeDigitized",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x932f}: ExifTagDesc{
		Id:     0x932f,
		Name:   "MSDocumentText",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9330}: ExifTagDesc{
		Id:     0x9330,
		Name:   "MSPropertySetStorage",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9331}: ExifTagDesc{
		Id:     0x9331,
		Name:   "MSDocumentTextPosition",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x935c}: ExifTagDesc{
		Id:     0x935c,
		Name:   "ImageSourceData",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9400}: ExifTagDesc{
		Id:     0x9400,
		Name:   "AmbientTemperature",
		Type:   ExifRational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9401}: ExifTagDesc{
		Id:     0x9401,
		Name:   "Humidity",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9402}: ExifTagDesc{
		Id:     0x9402,
		Name:   "Pressure",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9403}: ExifTagDesc{
		Id:     0x9403,
		Name:   "WaterDepth",
		Type:   ExifRational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9404}: ExifTagDesc{
		Id:     0x9404,
		Name:   "Acceleration",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9405}: ExifTagDesc{
		Id:     0x9405,
		Name:   "CameraElevationAngle",
		Type:   ExifRational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9999}: ExifTagDesc{
		Id:     0x9999,
		Name:   "XiaomiSettings",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0x9a00}: ExifTagDesc{
		Id:     0x9a00,
		Name:   "XiaomiModel",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x9c9b}: ExifTagDesc{
		Id:     0x9c9b,
		Name:   "XPTitle",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x9c9c}: ExifTagDesc{
		Id:     0x9c9c,
		Name:   "XPComment",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x9c9d}: ExifTagDesc{
		Id:     0x9c9d,
		Name:   "XPAuthor",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x9c9e}: ExifTagDesc{
		Id:     0x9c9e,
		Name:   "XPKeywords",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0x9c9f}: ExifTagDesc{
		Id:     0x9c9f,
		Name:   "XPSubject",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa000}: ExifTagDesc{
		Id:     0xa000,
		Name:   "FlashpixVersion",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa001}: ExifTagDesc{
		Id:    0xa001,
		Name:  "ColorSpace",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			1:     "sRGB",
			2:     "Adobe RGB",
			65533: "Wide Gamut RGB",
			65534: "ICC Profile",
			65535: "Uncalibrated",
		},
	},
	ExifIndexTag{ExifIFD, 0xa002}: ExifTagDesc{
		Id:     0xa002,
		Name:   "ExifImageWidth",
		Type:   ExifUint16,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa003}: ExifTagDesc{
		Id:     0xa003,
		Name:   "ExifImageHeight",
		Type:   ExifUint16,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa004}: ExifTagDesc{
		Id:     0xa004,
		Name:   "RelatedSoundFile",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa005}: ExifTagDesc{
		Id:     0xa005,
		Name:   "InteropOffset",
		Type:   ExifUint32,
		Ifd:    ExifIFD,
		Count:  1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa010}: ExifTagDesc{
		Id:     0xa010,
		Name:   "SamsungRawPointersOffset",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa011}: ExifTagDesc{
		Id:     0xa011,
		Name:   "SamsungRawPointersLength",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa101}: ExifTagDesc{
		Id:     0xa101,
		Name:   "SamsungRawByteOrder",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa102}: ExifTagDesc{
		Id:     0xa102,
		Name:   "SamsungRawUnknown",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa20b}: ExifTagDesc{
		Id:     0xa20b,
		Name:   "FlashEnergy_0xa20b",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa20c}: ExifTagDesc{
		Id:     0xa20c,
		Name:   "SpatialFrequencyResponse",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa20d}: ExifTagDesc{
		Id:     0xa20d,
		Name:   "Noise_0xa20d",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa20e}: ExifTagDesc{
		Id:     0xa20e,
		Name:   "FocalPlaneXResolution",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa20f}: ExifTagDesc{
		Id:     0xa20f,
		Name:   "FocalPlaneYResolution",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa210}: ExifTagDesc{
		Id:    0xa210,
		Name:  "FocalPlaneResolutionUnit_0xa210",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "None",
			2: "inches",
			3: "cm",
			4: "mm",
			5: "um",
		},
	},
	ExifIndexTag{ExifIFD, 0xa211}: ExifTagDesc{
		Id:     0xa211,
		Name:   "ImageNumber_0xa211",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa212}: ExifTagDesc{
		Id:     0xa212,
		Name:   "SecurityClassification_0xa212",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa213}: ExifTagDesc{
		Id:     0xa213,
		Name:   "ImageHistory_0xa213",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa214}: ExifTagDesc{
		Id:     0xa214,
		Name:   "SubjectLocation",
		Type:   ExifUint16,
		Ifd:    ExifIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa215}: ExifTagDesc{
		Id:     0xa215,
		Name:   "ExposureIndex",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa216}: ExifTagDesc{
		Id:     0xa216,
		Name:   "TIFFEPStandardID_0xa216",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa217}: ExifTagDesc{
		Id:    0xa217,
		Name:  "SensingMethod_0xa217",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "Not defined",
			2: "One-chip color area",
			3: "Two-chip color area",
			4: "Three-chip color area",
			5: "Color sequential area",
			7: "Trilinear",
			8: "Color sequential linear",
		},
	},
	ExifIndexTag{ExifIFD, 0xa300}: ExifTagDesc{
		Id:     0xa300,
		Name:   "FileSource",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa301}: ExifTagDesc{
		Id:     0xa301,
		Name:   "SceneType",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa302}: ExifTagDesc{
		Id:     0xa302,
		Name:   "CFAPattern",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa401}: ExifTagDesc{
		Id:    0xa401,
		Name:  "CustomRendered",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Normal",
			1: "Custom",
			2: "HDR (no original saved)",
			3: "HDR (original saved)",
			4: "Original (for HDR)",
			6: "Panorama",
			7: "Portrait HDR",
			8: "Portrait",
		},
	},
	ExifIndexTag{ExifIFD, 0xa402}: ExifTagDesc{
		Id:    0xa402,
		Name:  "ExposureMode",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Auto",
			1: "Manual",
			2: "Auto bracket",
		},
	},
	ExifIndexTag{ExifIFD, 0xa403}: ExifTagDesc{
		Id:    0xa403,
		Name:  "WhiteBalance",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Auto",
			1: "Manual",
		},
	},
	ExifIndexTag{ExifIFD, 0xa404}: ExifTagDesc{
		Id:     0xa404,
		Name:   "DigitalZoomRatio",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa405}: ExifTagDesc{
		Id:     0xa405,
		Name:   "FocalLengthIn35mmFormat",
		Type:   ExifUint16,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa406}: ExifTagDesc{
		Id:    0xa406,
		Name:  "SceneCaptureType",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Standard",
			1: "Landscape",
			2: "Portrait",
			3: "Night",
			4: "Other",
		},
	},
	ExifIndexTag{ExifIFD, 0xa407}: ExifTagDesc{
		Id:    0xa407,
		Name:  "GainControl",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "None",
			1: "Low gain up",
			2: "High gain up",
			3: "Low gain down",
			4: "High gain down",
		},
	},
	ExifIndexTag{ExifIFD, 0xa408}: ExifTagDesc{
		Id:    0xa408,
		Name:  "Contrast",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Normal",
			1: "Low",
			2: "High",
		},
	},
	ExifIndexTag{ExifIFD, 0xa409}: ExifTagDesc{
		Id:    0xa409,
		Name:  "Saturation",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Normal",
			1: "Low",
			2: "High",
		},
	},
	ExifIndexTag{ExifIFD, 0xa40a}: ExifTagDesc{
		Id:    0xa40a,
		Name:  "Sharpness",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Normal",
			1: "Soft",
			2: "Hard",
		},
	},
	ExifIndexTag{ExifIFD, 0xa40b}: ExifTagDesc{
		Id:     0xa40b,
		Name:   "DeviceSettingDescription",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa40c}: ExifTagDesc{
		Id:    0xa40c,
		Name:  "SubjectDistanceRange",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Unknown",
			1: "Macro",
			2: "Close",
			3: "Distant",
		},
	},
	ExifIndexTag{ExifIFD, 0xa420}: ExifTagDesc{
		Id:     0xa420,
		Name:   "ImageUniqueID",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa430}: ExifTagDesc{
		Id:     0xa430,
		Name:   "OwnerName",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa431}: ExifTagDesc{
		Id:     0xa431,
		Name:   "SerialNumber",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa432}: ExifTagDesc{
		Id:     0xa432,
		Name:   "LensInfo",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa433}: ExifTagDesc{
		Id:     0xa433,
		Name:   "LensMake",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa434}: ExifTagDesc{
		Id:     0xa434,
		Name:   "LensModel",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa435}: ExifTagDesc{
		Id:     0xa435,
		Name:   "LensSerialNumber",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa436}: ExifTagDesc{
		Id:     0xa436,
		Name:   "ImageTitle",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa437}: ExifTagDesc{
		Id:     0xa437,
		Name:   "Photographer",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa438}: ExifTagDesc{
		Id:     0xa438,
		Name:   "ImageEditor",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa439}: ExifTagDesc{
		Id:     0xa439,
		Name:   "CameraFirmware",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa43a}: ExifTagDesc{
		Id:     0xa43a,
		Name:   "RAWDevelopingSoftware",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa43b}: ExifTagDesc{
		Id:     0xa43b,
		Name:   "ImageEditingSoftware",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa43c}: ExifTagDesc{
		Id:     0xa43c,
		Name:   "MetadataEditingSoftware",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa460}: ExifTagDesc{
		Id:    0xa460,
		Name:  "CompositeImage",
		Type:  ExifUint16,
		Ifd:   ExifIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Unknown",
			1: "Not a Composite Image",
			2: "General Composite Image",
			3: "Composite Image Captured While Shooting",
		},
	},
	ExifIndexTag{ExifIFD, 0xa461}: ExifTagDesc{
		Id:     0xa461,
		Name:   "CompositeImageCount",
		Type:   ExifUint16,
		Ifd:    ExifIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa462}: ExifTagDesc{
		Id:     0xa462,
		Name:   "CompositeImageExposureTimes",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xa480}: ExifTagDesc{
		Id:     0xa480,
		Name:   "GDALMetadata",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xa481}: ExifTagDesc{
		Id:     0xa481,
		Name:   "GDALNoData",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xa500}: ExifTagDesc{
		Id:     0xa500,
		Name:   "Gamma",
		Type:   ExifUrational,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xafc0}: ExifTagDesc{
		Id:     0xafc0,
		Name:   "ExpandSoftware",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xafc1}: ExifTagDesc{
		Id:     0xafc1,
		Name:   "ExpandLens",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xafc2}: ExifTagDesc{
		Id:     0xafc2,
		Name:   "ExpandFilm",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xafc3}: ExifTagDesc{
		Id:     0xafc3,
		Name:   "ExpandFilterLens",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xafc4}: ExifTagDesc{
		Id:     0xafc4,
		Name:   "ExpandScanner",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xafc5}: ExifTagDesc{
		Id:     0xafc5,
		Name:   "ExpandFlashLamp",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xb4c3}: ExifTagDesc{
		Id:     0xb4c3,
		Name:   "HasselbladRawImage",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbc01}: ExifTagDesc{
		Id:     0xbc01,
		Name:   "PixelFormat",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbc02}: ExifTagDesc{
		Id:     0xbc02,
		Name:   "Transformation",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbc03}: ExifTagDesc{
		Id:     0xbc03,
		Name:   "Uncompressed",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbc04}: ExifTagDesc{
		Id:     0xbc04,
		Name:   "ImageType",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbc80}: ExifTagDesc{
		Id:     0xbc80,
		Name:   "ImageWidth",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbc81}: ExifTagDesc{
		Id:     0xbc81,
		Name:   "ImageHeight",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbc82}: ExifTagDesc{
		Id:     0xbc82,
		Name:   "WidthResolution",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbc83}: ExifTagDesc{
		Id:     0xbc83,
		Name:   "HeightResolution",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbcc0}: ExifTagDesc{
		Id:     0xbcc0,
		Name:   "ImageOffset",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbcc1}: ExifTagDesc{
		Id:     0xbcc1,
		Name:   "ImageByteCount",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbcc2}: ExifTagDesc{
		Id:     0xbcc2,
		Name:   "AlphaOffset",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbcc3}: ExifTagDesc{
		Id:     0xbcc3,
		Name:   "AlphaByteCount",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbcc4}: ExifTagDesc{
		Id:     0xbcc4,
		Name:   "ImageDataDiscard",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xbcc5}: ExifTagDesc{
		Id:     0xbcc5,
		Name:   "AlphaDataDiscard",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc427}: ExifTagDesc{
		Id:     0xc427,
		Name:   "OceScanjobDesc",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc428}: ExifTagDesc{
		Id:     0xc428,
		Name:   "OceApplicationSelector",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc429}: ExifTagDesc{
		Id:     0xc429,
		Name:   "OceIDNumber",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc42a}: ExifTagDesc{
		Id:     0xc42a,
		Name:   "OceImageLogic",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc44f}: ExifTagDesc{
		Id:     0xc44f,
		Name:   "Annotations",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc4a5}: ExifTagDesc{
		Id:     0xc4a5,
		Name:   "PrintIM",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc51b}: ExifTagDesc{
		Id:     0xc51b,
		Name:   "HasselbladExif",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc573}: ExifTagDesc{
		Id:     0xc573,
		Name:   "OriginalFileName",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc580}: ExifTagDesc{
		Id:     0xc580,
		Name:   "USPTOOriginalContentType",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc5e0}: ExifTagDesc{
		Id:     0xc5e0,
		Name:   "CR2CFAPattern",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc612}: ExifTagDesc{
		Id:     0xc612,
		Name:   "DNGVersion",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc613}: ExifTagDesc{
		Id:     0xc613,
		Name:   "DNGBackwardVersion",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc614}: ExifTagDesc{
		Id:     0xc614,
		Name:   "UniqueCameraModel",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc615}: ExifTagDesc{
		Id:     0xc615,
		Name:   "LocalizedCameraModel",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc616}: ExifTagDesc{
		Id:     0xc616,
		Name:   "CFAPlaneColor",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc617}: ExifTagDesc{
		Id:    0xc617,
		Name:  "CFALayout",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			1: "Rectangular",
			2: "Even columns offset down 1/2 row",
			3: "Even columns offset up 1/2 row",
			4: "Even rows offset right 1/2 column",
			5: "Even rows offset left 1/2 column",
			6: "Even rows offset up by 1/2 row, even columns offset left by 1/2 column",
			7: "Even rows offset up by 1/2 row, even columns offset right by 1/2 column",
			8: "Even rows offset down by 1/2 row, even columns offset left by 1/2 column",
			9: "Even rows offset down by 1/2 row, even columns offset right by 1/2 column",
		},
	},
	ExifIndexTag{RootIFD, 0xc618}: ExifTagDesc{
		Id:     0xc618,
		Name:   "LinearizationTable",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc619}: ExifTagDesc{
		Id:     0xc619,
		Name:   "BlackLevelRepeatDim",
		Type:   ExifUint16,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc61a}: ExifTagDesc{
		Id:     0xc61a,
		Name:   "BlackLevel",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc61b}: ExifTagDesc{
		Id:     0xc61b,
		Name:   "BlackLevelDeltaH",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc61c}: ExifTagDesc{
		Id:     0xc61c,
		Name:   "BlackLevelDeltaV",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc61d}: ExifTagDesc{
		Id:     0xc61d,
		Name:   "WhiteLevel",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc61e}: ExifTagDesc{
		Id:     0xc61e,
		Name:   "DefaultScale",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc61f}: ExifTagDesc{
		Id:     0xc61f,
		Name:   "DefaultCropOrigin",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc620}: ExifTagDesc{
		Id:     0xc620,
		Name:   "DefaultCropSize",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc621}: ExifTagDesc{
		Id:     0xc621,
		Name:   "ColorMatrix1",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc622}: ExifTagDesc{
		Id:     0xc622,
		Name:   "ColorMatrix2",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc623}: ExifTagDesc{
		Id:     0xc623,
		Name:   "CameraCalibration1",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc624}: ExifTagDesc{
		Id:     0xc624,
		Name:   "CameraCalibration2",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc625}: ExifTagDesc{
		Id:     0xc625,
		Name:   "ReductionMatrix1",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc626}: ExifTagDesc{
		Id:     0xc626,
		Name:   "ReductionMatrix2",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc627}: ExifTagDesc{
		Id:     0xc627,
		Name:   "AnalogBalance",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc628}: ExifTagDesc{
		Id:     0xc628,
		Name:   "AsShotNeutral",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc629}: ExifTagDesc{
		Id:     0xc629,
		Name:   "AsShotWhiteXY",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc62a}: ExifTagDesc{
		Id:     0xc62a,
		Name:   "BaselineExposure",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc62b}: ExifTagDesc{
		Id:     0xc62b,
		Name:   "BaselineNoise",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc62c}: ExifTagDesc{
		Id:     0xc62c,
		Name:   "BaselineSharpness",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc62d}: ExifTagDesc{
		Id:     0xc62d,
		Name:   "BayerGreenSplit",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc62e}: ExifTagDesc{
		Id:     0xc62e,
		Name:   "LinearResponseLimit",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc62f}: ExifTagDesc{
		Id:     0xc62f,
		Name:   "CameraSerialNumber",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc630}: ExifTagDesc{
		Id:     0xc630,
		Name:   "DNGLensInfo",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc631}: ExifTagDesc{
		Id:     0xc631,
		Name:   "ChromaBlurRadius",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc632}: ExifTagDesc{
		Id:     0xc632,
		Name:   "AntiAliasStrength",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc633}: ExifTagDesc{
		Id:     0xc633,
		Name:   "ShadowScale",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc634}: ExifTagDesc{
		Id:     0xc634,
		Name:   "DNGPrivateData",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc635}: ExifTagDesc{
		Id:    0xc635,
		Name:  "MakerNoteSafety",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Unsafe",
			1: "Safe",
		},
	},
	ExifIndexTag{ExifIFD, 0xc640}: ExifTagDesc{
		Id:     0xc640,
		Name:   "RawImageSegmentation",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc65a}: ExifTagDesc{
		Id:    0xc65a,
		Name:  "CalibrationIlluminant1",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0:   "Unknown",
			1:   "Daylight",
			10:  "Cloudy",
			11:  "Shade",
			12:  "Daylight Fluorescent",
			13:  "Day White Fluorescent",
			14:  "Cool White Fluorescent",
			15:  "White Fluorescent",
			16:  "Warm White Fluorescent",
			17:  "Standard Light A",
			18:  "Standard Light B",
			19:  "Standard Light C",
			2:   "Fluorescent",
			20:  "D55",
			21:  "D65",
			22:  "D75",
			23:  "D50",
			24:  "ISO Studio Tungsten",
			255: "Other",
			3:   "Tungsten (Incandescent)",
			4:   "Flash",
			9:   "Fine Weather",
		},
	},
	ExifIndexTag{RootIFD, 0xc65b}: ExifTagDesc{
		Id:    0xc65b,
		Name:  "CalibrationIlluminant2",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0:   "Unknown",
			1:   "Daylight",
			10:  "Cloudy",
			11:  "Shade",
			12:  "Daylight Fluorescent",
			13:  "Day White Fluorescent",
			14:  "Cool White Fluorescent",
			15:  "White Fluorescent",
			16:  "Warm White Fluorescent",
			17:  "Standard Light A",
			18:  "Standard Light B",
			19:  "Standard Light C",
			2:   "Fluorescent",
			20:  "D55",
			21:  "D65",
			22:  "D75",
			23:  "D50",
			24:  "ISO Studio Tungsten",
			255: "Other",
			3:   "Tungsten (Incandescent)",
			4:   "Flash",
			9:   "Fine Weather",
		},
	},
	ExifIndexTag{RootIFD, 0xc65c}: ExifTagDesc{
		Id:     0xc65c,
		Name:   "BestQualityScale",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc65d}: ExifTagDesc{
		Id:     0xc65d,
		Name:   "RawDataUniqueID",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  16,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc660}: ExifTagDesc{
		Id:     0xc660,
		Name:   "AliasLayerMetadata",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc68b}: ExifTagDesc{
		Id:     0xc68b,
		Name:   "OriginalRawFileName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc68c}: ExifTagDesc{
		Id:     0xc68c,
		Name:   "OriginalRawFileData",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc68d}: ExifTagDesc{
		Id:     0xc68d,
		Name:   "ActiveArea",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc68e}: ExifTagDesc{
		Id:     0xc68e,
		Name:   "MaskedAreas",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc68f}: ExifTagDesc{
		Id:     0xc68f,
		Name:   "AsShotICCProfile",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc690}: ExifTagDesc{
		Id:     0xc690,
		Name:   "AsShotPreProfileMatrix",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc691}: ExifTagDesc{
		Id:     0xc691,
		Name:   "CurrentICCProfile",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc692}: ExifTagDesc{
		Id:     0xc692,
		Name:   "CurrentPreProfileMatrix",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6bf}: ExifTagDesc{
		Id:    0xc6bf,
		Name:  "ColorimetricReference",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Scene-referred",
			1: "Output-referred (ICC Profile Dynamic Range)",
			2: "Output-referred (High Dyanmic Range)",
		},
	},
	ExifIndexTag{RootIFD, 0xc6c5}: ExifTagDesc{
		Id:     0xc6c5,
		Name:   "SRawType",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6d2}: ExifTagDesc{
		Id:     0xc6d2,
		Name:   "PanasonicTitle",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6d3}: ExifTagDesc{
		Id:     0xc6d3,
		Name:   "PanasonicTitle2",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6f3}: ExifTagDesc{
		Id:     0xc6f3,
		Name:   "CameraCalibrationSig",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6f4}: ExifTagDesc{
		Id:     0xc6f4,
		Name:   "ProfileCalibrationSig",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6f5}: ExifTagDesc{
		Id:     0xc6f5,
		Name:   "ProfileIFD",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6f6}: ExifTagDesc{
		Id:     0xc6f6,
		Name:   "AsShotProfileName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6f7}: ExifTagDesc{
		Id:     0xc6f7,
		Name:   "NoiseReductionApplied",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6f8}: ExifTagDesc{
		Id:     0xc6f8,
		Name:   "ProfileName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6f9}: ExifTagDesc{
		Id:     0xc6f9,
		Name:   "ProfileHueSatMapDims",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6fa}: ExifTagDesc{
		Id:     0xc6fa,
		Name:   "ProfileHueSatMapData1",
		Type:   ExifFloat,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6fb}: ExifTagDesc{
		Id:     0xc6fb,
		Name:   "ProfileHueSatMapData2",
		Type:   ExifFloat,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6fc}: ExifTagDesc{
		Id:     0xc6fc,
		Name:   "ProfileToneCurve",
		Type:   ExifFloat,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc6fd}: ExifTagDesc{
		Id:    0xc6fd,
		Name:  "ProfileEmbedPolicy",
		Type:  ExifUint32,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint32]string{
			0: "Allow Copying",
			1: "Embed if Used",
			2: "Never Embed",
			3: "No Restrictions",
		},
	},
	ExifIndexTag{RootIFD, 0xc6fe}: ExifTagDesc{
		Id:     0xc6fe,
		Name:   "ProfileCopyright",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc714}: ExifTagDesc{
		Id:     0xc714,
		Name:   "ForwardMatrix1",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc715}: ExifTagDesc{
		Id:     0xc715,
		Name:   "ForwardMatrix2",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc716}: ExifTagDesc{
		Id:     0xc716,
		Name:   "PreviewApplicationName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc717}: ExifTagDesc{
		Id:     0xc717,
		Name:   "PreviewApplicationVersion",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc718}: ExifTagDesc{
		Id:     0xc718,
		Name:   "PreviewSettingsName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc719}: ExifTagDesc{
		Id:     0xc719,
		Name:   "PreviewSettingsDigest",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc71a}: ExifTagDesc{
		Id:    0xc71a,
		Name:  "PreviewColorSpace",
		Type:  ExifUint32,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint32]string{
			0: "Unknown",
			1: "Gray Gamma 2.2",
			2: "sRGB",
			3: "Adobe RGB",
			4: "ProPhoto RGB",
		},
	},
	ExifIndexTag{RootIFD, 0xc71b}: ExifTagDesc{
		Id:     0xc71b,
		Name:   "PreviewDateTime",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc71c}: ExifTagDesc{
		Id:     0xc71c,
		Name:   "RawImageDigest",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  16,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc71d}: ExifTagDesc{
		Id:     0xc71d,
		Name:   "OriginalRawFileDigest",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  16,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc71e}: ExifTagDesc{
		Id:     0xc71e,
		Name:   "SubTileBlockSize",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc71f}: ExifTagDesc{
		Id:     0xc71f,
		Name:   "RowInterleaveFactor",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc725}: ExifTagDesc{
		Id:     0xc725,
		Name:   "ProfileLookTableDims",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc726}: ExifTagDesc{
		Id:     0xc726,
		Name:   "ProfileLookTableData",
		Type:   ExifFloat,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc740}: ExifTagDesc{
		Id:     0xc740,
		Name:   "OpcodeList1",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc741}: ExifTagDesc{
		Id:     0xc741,
		Name:   "OpcodeList2",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc74e}: ExifTagDesc{
		Id:     0xc74e,
		Name:   "OpcodeList3",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc761}: ExifTagDesc{
		Id:     0xc761,
		Name:   "NoiseProfile",
		Type:   ExifDouble,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc763}: ExifTagDesc{
		Id:     0xc763,
		Name:   "TimeCodes",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc764}: ExifTagDesc{
		Id:     0xc764,
		Name:   "FrameRate",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc772}: ExifTagDesc{
		Id:     0xc772,
		Name:   "TStop",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc789}: ExifTagDesc{
		Id:     0xc789,
		Name:   "ReelName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc791}: ExifTagDesc{
		Id:     0xc791,
		Name:   "OriginalDefaultFinalSize",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc792}: ExifTagDesc{
		Id:     0xc792,
		Name:   "OriginalBestQualitySize",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc793}: ExifTagDesc{
		Id:     0xc793,
		Name:   "OriginalDefaultCropSize",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7a1}: ExifTagDesc{
		Id:     0xc7a1,
		Name:   "CameraLabel",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7a3}: ExifTagDesc{
		Id:    0xc7a3,
		Name:  "ProfileHueSatMapEncoding",
		Type:  ExifUint32,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint32]string{
			0: "Linear",
			1: "sRGB",
		},
	},
	ExifIndexTag{RootIFD, 0xc7a4}: ExifTagDesc{
		Id:    0xc7a4,
		Name:  "ProfileLookTableEncoding",
		Type:  ExifUint32,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint32]string{
			0: "Linear",
			1: "sRGB",
		},
	},
	ExifIndexTag{RootIFD, 0xc7a5}: ExifTagDesc{
		Id:     0xc7a5,
		Name:   "BaselineExposureOffset",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7a6}: ExifTagDesc{
		Id:    0xc7a6,
		Name:  "DefaultBlackRender",
		Type:  ExifUint32,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint32]string{
			0: "Auto",
			1: "None",
		},
	},
	ExifIndexTag{RootIFD, 0xc7a7}: ExifTagDesc{
		Id:     0xc7a7,
		Name:   "NewRawImageDigest",
		Type:   ExifUint8,
		Ifd:    RootIFD,
		Count:  16,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7a8}: ExifTagDesc{
		Id:     0xc7a8,
		Name:   "RawToPreviewGain",
		Type:   ExifDouble,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7aa}: ExifTagDesc{
		Id:     0xc7aa,
		Name:   "CacheVersion",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7b5}: ExifTagDesc{
		Id:     0xc7b5,
		Name:   "DefaultUserCrop",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc7d5}: ExifTagDesc{
		Id:     0xc7d5,
		Name:   "NikonNEFInfo",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc7d7}: ExifTagDesc{
		Id:     0xc7d7,
		Name:   "ZIFMetadata",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xc7d8}: ExifTagDesc{
		Id:     0xc7d8,
		Name:   "ZIFAnnotations",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7e9}: ExifTagDesc{
		Id:    0xc7e9,
		Name:  "DepthFormat",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Unknown",
			1: "Linear",
			2: "Inverse",
		},
	},
	ExifIndexTag{RootIFD, 0xc7ea}: ExifTagDesc{
		Id:     0xc7ea,
		Name:   "DepthNear",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7eb}: ExifTagDesc{
		Id:     0xc7eb,
		Name:   "DepthFar",
		Type:   ExifUrational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xc7ec}: ExifTagDesc{
		Id:    0xc7ec,
		Name:  "DepthUnits",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Unknown",
			1: "Meters",
		},
	},
	ExifIndexTag{RootIFD, 0xc7ed}: ExifTagDesc{
		Id:    0xc7ed,
		Name:  "DepthMeasureType",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "Unknown",
			1: "Optical Axis",
			2: "Optical Ray",
		},
	},
	ExifIndexTag{RootIFD, 0xc7ee}: ExifTagDesc{
		Id:     0xc7ee,
		Name:   "EnhanceParams",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd2d}: ExifTagDesc{
		Id:     0xcd2d,
		Name:   "ProfileGainTableMap",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd2e}: ExifTagDesc{
		Id:     0xcd2e,
		Name:   "SemanticName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd30}: ExifTagDesc{
		Id:     0xcd30,
		Name:   "SemanticInstanceIFD",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd31}: ExifTagDesc{
		Id:    0xcd31,
		Name:  "CalibrationIlluminant3",
		Type:  ExifUint16,
		Ifd:   RootIFD,
		Count: -1,
		Values: map[uint16]string{
			0:   "Unknown",
			1:   "Daylight",
			10:  "Cloudy",
			11:  "Shade",
			12:  "Daylight Fluorescent",
			13:  "Day White Fluorescent",
			14:  "Cool White Fluorescent",
			15:  "White Fluorescent",
			16:  "Warm White Fluorescent",
			17:  "Standard Light A",
			18:  "Standard Light B",
			19:  "Standard Light C",
			2:   "Fluorescent",
			20:  "D55",
			21:  "D65",
			22:  "D75",
			23:  "D50",
			24:  "ISO Studio Tungsten",
			255: "Other",
			3:   "Tungsten (Incandescent)",
			4:   "Flash",
			9:   "Fine Weather",
		},
	},
	ExifIndexTag{RootIFD, 0xcd32}: ExifTagDesc{
		Id:     0xcd32,
		Name:   "CameraCalibration3",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd33}: ExifTagDesc{
		Id:     0xcd33,
		Name:   "ColorMatrix3",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd34}: ExifTagDesc{
		Id:     0xcd34,
		Name:   "ForwardMatrix3",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd35}: ExifTagDesc{
		Id:     0xcd35,
		Name:   "IlluminantData1",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd36}: ExifTagDesc{
		Id:     0xcd36,
		Name:   "IlluminantData2",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd37}: ExifTagDesc{
		Id:     0xcd37,
		Name:   "IlluminantData3",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd38}: ExifTagDesc{
		Id:     0xcd38,
		Name:   "MaskSubArea",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd39}: ExifTagDesc{
		Id:     0xcd39,
		Name:   "ProfileHueSatMapData3",
		Type:   ExifFloat,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd3a}: ExifTagDesc{
		Id:     0xcd3a,
		Name:   "ReductionMatrix3",
		Type:   ExifRational,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd3b}: ExifTagDesc{
		Id:     0xcd3b,
		Name:   "RGBTables",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd3f}: ExifTagDesc{
		Id:     0xcd3f,
		Name:   "RGBTables_0xcd3f",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd40}: ExifTagDesc{
		Id:     0xcd40,
		Name:   "ProfileGainTableMap2",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd43}: ExifTagDesc{
		Id:     0xcd43,
		Name:   "ColumnInterleaveFactor",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd44}: ExifTagDesc{
		Id:     0xcd44,
		Name:   "ImageSequenceInfo",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd46}: ExifTagDesc{
		Id:     0xcd46,
		Name:   "ImageStats",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd47}: ExifTagDesc{
		Id:     0xcd47,
		Name:   "ProfileDynamicRange",
		Type:   ExifUndef,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd48}: ExifTagDesc{
		Id:     0xcd48,
		Name:   "ProfileGroupName",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd49}: ExifTagDesc{
		Id:     0xcd49,
		Name:   "JXLDistance",
		Type:   ExifFloat,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd4a}: ExifTagDesc{
		Id:     0xcd4a,
		Name:   "JXLEffort",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcd4b}: ExifTagDesc{
		Id:     0xcd4b,
		Name:   "JXLDecodeSpeed",
		Type:   ExifUint32,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{RootIFD, 0xcea1}: ExifTagDesc{
		Id:     0xcea1,
		Name:   "SEAL",
		Type:   ExifString,
		Ifd:    RootIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xea1c}: ExifTagDesc{
		Id:     0xea1c,
		Name:   "Padding",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xea1d}: ExifTagDesc{
		Id:     0xea1d,
		Name:   "OffsetSchema",
		Type:   ExifInt32,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfde8}: ExifTagDesc{
		Id:     0xfde8,
		Name:   "OwnerName_0xfde8",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfde9}: ExifTagDesc{
		Id:     0xfde9,
		Name:   "SerialNumber_0xfde9",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfdea}: ExifTagDesc{
		Id:     0xfdea,
		Name:   "Lens",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe00}: ExifTagDesc{
		Id:     0xfe00,
		Name:   "KDC_IFD",
		Type:   ExifUndef,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe4c}: ExifTagDesc{
		Id:     0xfe4c,
		Name:   "RawFile",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe4d}: ExifTagDesc{
		Id:     0xfe4d,
		Name:   "Converter",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe4e}: ExifTagDesc{
		Id:     0xfe4e,
		Name:   "WhiteBalance_0xfe4e",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe51}: ExifTagDesc{
		Id:     0xfe51,
		Name:   "Exposure",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe52}: ExifTagDesc{
		Id:     0xfe52,
		Name:   "Shadows",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe53}: ExifTagDesc{
		Id:     0xfe53,
		Name:   "Brightness",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe54}: ExifTagDesc{
		Id:     0xfe54,
		Name:   "Contrast_0xfe54",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe55}: ExifTagDesc{
		Id:     0xfe55,
		Name:   "Saturation_0xfe55",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe56}: ExifTagDesc{
		Id:     0xfe56,
		Name:   "Sharpness_0xfe56",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe57}: ExifTagDesc{
		Id:     0xfe57,
		Name:   "Smoothness",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{ExifIFD, 0xfe58}: ExifTagDesc{
		Id:     0xfe58,
		Name:   "MoireFilter",
		Type:   ExifString,
		Ifd:    ExifIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0000}: ExifTagDesc{
		Id:     0x0000,
		Name:   "GPSVersionID",
		Type:   ExifUint8,
		Ifd:    GpsIFD,
		Count:  4,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0001}: ExifTagDesc{
		Id:     0x0001,
		Name:   "GPSLatitudeRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0002}: ExifTagDesc{
		Id:     0x0002,
		Name:   "GPSLatitude",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0003}: ExifTagDesc{
		Id:     0x0003,
		Name:   "GPSLongitudeRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0004}: ExifTagDesc{
		Id:     0x0004,
		Name:   "GPSLongitude",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0005}: ExifTagDesc{
		Id:    0x0005,
		Name:  "GPSAltitudeRef",
		Type:  ExifUint8,
		Ifd:   GpsIFD,
		Count: -1,
		Values: map[uint8]string{
			0: "Above Sea Level",
			1: "Below Sea Level",
			2: "Positive Sea Level (sea-level ref)",
			3: "Negative Sea Level (sea-level ref)",
		},
	},
	ExifIndexTag{GpsIFD, 0x0006}: ExifTagDesc{
		Id:     0x0006,
		Name:   "GPSAltitude",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0007}: ExifTagDesc{
		Id:     0x0007,
		Name:   "GPSTimeStamp",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0008}: ExifTagDesc{
		Id:     0x0008,
		Name:   "GPSSatellites",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0009}: ExifTagDesc{
		Id:     0x0009,
		Name:   "GPSStatus",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x000a}: ExifTagDesc{
		Id:     0x000a,
		Name:   "GPSMeasureMode",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x000b}: ExifTagDesc{
		Id:     0x000b,
		Name:   "GPSDOP",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x000c}: ExifTagDesc{
		Id:     0x000c,
		Name:   "GPSSpeedRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x000d}: ExifTagDesc{
		Id:     0x000d,
		Name:   "GPSSpeed",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x000e}: ExifTagDesc{
		Id:     0x000e,
		Name:   "GPSTrackRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x000f}: ExifTagDesc{
		Id:     0x000f,
		Name:   "GPSTrack",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0010}: ExifTagDesc{
		Id:     0x0010,
		Name:   "GPSImgDirectionRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0011}: ExifTagDesc{
		Id:     0x0011,
		Name:   "GPSImgDirection",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0012}: ExifTagDesc{
		Id:     0x0012,
		Name:   "GPSMapDatum",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0013}: ExifTagDesc{
		Id:     0x0013,
		Name:   "GPSDestLatitudeRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0014}: ExifTagDesc{
		Id:     0x0014,
		Name:   "GPSDestLatitude",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0015}: ExifTagDesc{
		Id:     0x0015,
		Name:   "GPSDestLongitudeRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0016}: ExifTagDesc{
		Id:     0x0016,
		Name:   "GPSDestLongitude",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  3,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0017}: ExifTagDesc{
		Id:     0x0017,
		Name:   "GPSDestBearingRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0018}: ExifTagDesc{
		Id:     0x0018,
		Name:   "GPSDestBearing",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x0019}: ExifTagDesc{
		Id:     0x0019,
		Name:   "GPSDestDistanceRef",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  2,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x001a}: ExifTagDesc{
		Id:     0x001a,
		Name:   "GPSDestDistance",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x001b}: ExifTagDesc{
		Id:     0x001b,
		Name:   "GPSProcessingMethod",
		Type:   ExifUndef,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x001c}: ExifTagDesc{
		Id:     0x001c,
		Name:   "GPSAreaInformation",
		Type:   ExifUndef,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x001d}: ExifTagDesc{
		Id:     0x001d,
		Name:   "GPSDateStamp",
		Type:   ExifString,
		Ifd:    GpsIFD,
		Count:  11,
		Values: nil,
	},
	ExifIndexTag{GpsIFD, 0x001e}: ExifTagDesc{
		Id:    0x001e,
		Name:  "GPSDifferential",
		Type:  ExifUint16,
		Ifd:   GpsIFD,
		Count: -1,
		Values: map[uint16]string{
			0: "No Correction",
			1: "Differential Corrected",
		},
	},
	ExifIndexTag{GpsIFD, 0x001f}: ExifTagDesc{
		Id:     0x001f,
		Name:   "GPSHPositioningError",
		Type:   ExifUrational,
		Ifd:    GpsIFD,
		Count:  -1,
		Values: nil,
	},
}
