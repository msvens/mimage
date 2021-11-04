package metadata

//Do not edit! This is an automatically generated file (see generator.GenerateIptcTags()).
//This file was generated based on https://exiftool.org/TagNames/IPTC.html

//IPTC Records
const (
	IPTCEnvelop        uint8 = 1
	IPTCApplication    uint8 = 2
	IPTCNewsPhoto      uint8 = 3
	IPTCPreObjectData  uint8 = 7
	IPTCObjectData     uint8 = 8
	IPTCPostObjectData uint8 = 9
	IPTCFotoStation    uint8 = 240
)

var IptcRecordName = map[uint8]string{
	IPTCEnvelop:        "IPTCEnvelop",
	IPTCApplication:    "IPTCApplication",
	IPTCNewsPhoto:      "IPTCNewsPhoto",
	IPTCPreObjectData:  "IPTCPreObjectData",
	IPTCObjectData:     "IPTCObjectData",
	IPTCPostObjectData: "IPTCPostObjectData",
	IPTCFotoStation:    "IPTCFotoStation",
}

//IPTCPreObjectData Tag Ids
const (
	Iptc7_SizeMode            uint8 = 10
	Iptc7_MaxSubfileSize      uint8 = 20
	Iptc7_ObjectSizeAnnounced uint8 = 90
	Iptc7_MaximumObjectSize   uint8 = 95
)

var Iptc7Name = map[uint8]string{
	Iptc7_SizeMode:            "SizeMode",
	Iptc7_MaxSubfileSize:      "MaxSubfileSize",
	Iptc7_ObjectSizeAnnounced: "ObjectSizeAnnounced",
	Iptc7_MaximumObjectSize:   "MaximumObjectSize",
}

//IPTCObjectData Tag Ids
const (
	Iptc8_SubFile uint8 = 10
)

var Iptc8Name = map[uint8]string{
	Iptc8_SubFile: "SubFile",
}

//IPTCPostObjectData Tag Ids
const (
	Iptc9_ConfirmedObjectSize uint8 = 10
)

var Iptc9Name = map[uint8]string{
	Iptc9_ConfirmedObjectSize: "ConfirmedObjectSize",
}

//IPTCFotoStation Tag Ids
const ()

var Iptc240Name = map[uint8]string{}

//IPTCEnvelope Tag Ids
const (
	Iptc1_EnvelopeRecordVersion uint8 = 0
	Iptc1_Destination           uint8 = 5
	Iptc1_FileFormat            uint8 = 20
	Iptc1_FileVersion           uint8 = 22
	Iptc1_ServiceIdentifier     uint8 = 30
	Iptc1_EnvelopeNumber        uint8 = 40
	Iptc1_ProductID             uint8 = 50
	Iptc1_EnvelopePriority      uint8 = 60
	Iptc1_DateSent              uint8 = 70
	Iptc1_TimeSent              uint8 = 80
	Iptc1_CodedCharacterSet     uint8 = 90
	Iptc1_UniqueObjectName      uint8 = 100
	Iptc1_ARMIdentifier         uint8 = 120
	Iptc1_ARMVersion            uint8 = 122
)

var Iptc1Name = map[uint8]string{
	Iptc1_EnvelopeRecordVersion: "EnvelopeRecordVersion",
	Iptc1_Destination:           "Destination",
	Iptc1_FileFormat:            "FileFormat",
	Iptc1_FileVersion:           "FileVersion",
	Iptc1_ServiceIdentifier:     "ServiceIdentifier",
	Iptc1_EnvelopeNumber:        "EnvelopeNumber",
	Iptc1_ProductID:             "ProductID",
	Iptc1_EnvelopePriority:      "EnvelopePriority",
	Iptc1_DateSent:              "DateSent",
	Iptc1_TimeSent:              "TimeSent",
	Iptc1_CodedCharacterSet:     "CodedCharacterSet",
	Iptc1_UniqueObjectName:      "UniqueObjectName",
	Iptc1_ARMIdentifier:         "ARMIdentifier",
	Iptc1_ARMVersion:            "ARMVersion",
}

//IPTCApplication Tag Ids
const (
	Iptc2_ApplicationRecordVersion      uint8 = 0
	Iptc2_ObjectTypeReference           uint8 = 3
	Iptc2_ObjectAttributeReference      uint8 = 4
	Iptc2_ObjectName                    uint8 = 5
	Iptc2_EditStatus                    uint8 = 7
	Iptc2_EditorialUpdate               uint8 = 8
	Iptc2_Urgency                       uint8 = 10
	Iptc2_SubjectReference              uint8 = 12
	Iptc2_Category                      uint8 = 15
	Iptc2_SupplementalCategories        uint8 = 20
	Iptc2_FixtureIdentifier             uint8 = 22
	Iptc2_Keywords                      uint8 = 25
	Iptc2_ContentLocationCode           uint8 = 26
	Iptc2_ContentLocationName           uint8 = 27
	Iptc2_ReleaseDate                   uint8 = 30
	Iptc2_ReleaseTime                   uint8 = 35
	Iptc2_ExpirationDate                uint8 = 37
	Iptc2_ExpirationTime                uint8 = 38
	Iptc2_SpecialInstructions           uint8 = 40
	Iptc2_ActionAdvised                 uint8 = 42
	Iptc2_ReferenceService              uint8 = 45
	Iptc2_ReferenceDate                 uint8 = 47
	Iptc2_ReferenceNumber               uint8 = 50
	Iptc2_DateCreated                   uint8 = 55
	Iptc2_TimeCreated                   uint8 = 60
	Iptc2_DigitalCreationDate           uint8 = 62
	Iptc2_DigitalCreationTime           uint8 = 63
	Iptc2_OriginatingProgram            uint8 = 65
	Iptc2_ProgramVersion                uint8 = 70
	Iptc2_ObjectCycle                   uint8 = 75
	Iptc2_Byline                        uint8 = 80
	Iptc2_BylineTitle                   uint8 = 85
	Iptc2_City                          uint8 = 90
	Iptc2_Sublocation                   uint8 = 92
	Iptc2_ProvinceState                 uint8 = 95
	Iptc2_CountryPrimaryLocationCode    uint8 = 100
	Iptc2_CountryPrimaryLocationName    uint8 = 101
	Iptc2_OriginalTransmissionReference uint8 = 103
	Iptc2_Headline                      uint8 = 105
	Iptc2_Credit                        uint8 = 110
	Iptc2_Source                        uint8 = 115
	Iptc2_CopyrightNotice               uint8 = 116
	Iptc2_Contact                       uint8 = 118
	Iptc2_CaptionAbstract               uint8 = 120
	Iptc2_LocalCaption                  uint8 = 121
	Iptc2_WriterEditor                  uint8 = 122
	Iptc2_RasterizedCaption             uint8 = 125
	Iptc2_ImageType                     uint8 = 130
	Iptc2_ImageOrientation              uint8 = 131
	Iptc2_LanguageIdentifier            uint8 = 135
	Iptc2_AudioType                     uint8 = 150
	Iptc2_AudioSamplingRate             uint8 = 151
	Iptc2_AudioSamplingResolution       uint8 = 152
	Iptc2_AudioDuration                 uint8 = 153
	Iptc2_AudioOutcue                   uint8 = 154
	Iptc2_JobID                         uint8 = 184
	Iptc2_MasterDocumentID              uint8 = 185
	Iptc2_ShortDocumentID               uint8 = 186
	Iptc2_UniqueDocumentID              uint8 = 187
	Iptc2_OwnerID                       uint8 = 188
	Iptc2_ObjectPreviewFileFormat       uint8 = 200
	Iptc2_ObjectPreviewFileVersion      uint8 = 201
	Iptc2_ObjectPreviewData             uint8 = 202
	Iptc2_Prefs                         uint8 = 221
	Iptc2_ClassifyState                 uint8 = 225
	Iptc2_SimilarityIndex               uint8 = 228
	Iptc2_DocumentNotes                 uint8 = 230
	Iptc2_DocumentHistory               uint8 = 231
	Iptc2_ExifCameraInfo                uint8 = 232
	Iptc2_CatalogSets                   uint8 = 255
)

var Iptc2Name = map[uint8]string{
	Iptc2_ApplicationRecordVersion:      "ApplicationRecordVersion",
	Iptc2_ObjectTypeReference:           "ObjectTypeReference",
	Iptc2_ObjectAttributeReference:      "ObjectAttributeReference",
	Iptc2_ObjectName:                    "ObjectName",
	Iptc2_EditStatus:                    "EditStatus",
	Iptc2_EditorialUpdate:               "EditorialUpdate",
	Iptc2_Urgency:                       "Urgency",
	Iptc2_SubjectReference:              "SubjectReference",
	Iptc2_Category:                      "Category",
	Iptc2_SupplementalCategories:        "SupplementalCategories",
	Iptc2_FixtureIdentifier:             "FixtureIdentifier",
	Iptc2_Keywords:                      "Keywords",
	Iptc2_ContentLocationCode:           "ContentLocationCode",
	Iptc2_ContentLocationName:           "ContentLocationName",
	Iptc2_ReleaseDate:                   "ReleaseDate",
	Iptc2_ReleaseTime:                   "ReleaseTime",
	Iptc2_ExpirationDate:                "ExpirationDate",
	Iptc2_ExpirationTime:                "ExpirationTime",
	Iptc2_SpecialInstructions:           "SpecialInstructions",
	Iptc2_ActionAdvised:                 "ActionAdvised",
	Iptc2_ReferenceService:              "ReferenceService",
	Iptc2_ReferenceDate:                 "ReferenceDate",
	Iptc2_ReferenceNumber:               "ReferenceNumber",
	Iptc2_DateCreated:                   "DateCreated",
	Iptc2_TimeCreated:                   "TimeCreated",
	Iptc2_DigitalCreationDate:           "DigitalCreationDate",
	Iptc2_DigitalCreationTime:           "DigitalCreationTime",
	Iptc2_OriginatingProgram:            "OriginatingProgram",
	Iptc2_ProgramVersion:                "ProgramVersion",
	Iptc2_ObjectCycle:                   "ObjectCycle",
	Iptc2_Byline:                        "Byline",
	Iptc2_BylineTitle:                   "BylineTitle",
	Iptc2_City:                          "City",
	Iptc2_Sublocation:                   "Sublocation",
	Iptc2_ProvinceState:                 "ProvinceState",
	Iptc2_CountryPrimaryLocationCode:    "CountryPrimaryLocationCode",
	Iptc2_CountryPrimaryLocationName:    "CountryPrimaryLocationName",
	Iptc2_OriginalTransmissionReference: "OriginalTransmissionReference",
	Iptc2_Headline:                      "Headline",
	Iptc2_Credit:                        "Credit",
	Iptc2_Source:                        "Source",
	Iptc2_CopyrightNotice:               "CopyrightNotice",
	Iptc2_Contact:                       "Contact",
	Iptc2_CaptionAbstract:               "CaptionAbstract",
	Iptc2_LocalCaption:                  "LocalCaption",
	Iptc2_WriterEditor:                  "WriterEditor",
	Iptc2_RasterizedCaption:             "RasterizedCaption",
	Iptc2_ImageType:                     "ImageType",
	Iptc2_ImageOrientation:              "ImageOrientation",
	Iptc2_LanguageIdentifier:            "LanguageIdentifier",
	Iptc2_AudioType:                     "AudioType",
	Iptc2_AudioSamplingRate:             "AudioSamplingRate",
	Iptc2_AudioSamplingResolution:       "AudioSamplingResolution",
	Iptc2_AudioDuration:                 "AudioDuration",
	Iptc2_AudioOutcue:                   "AudioOutcue",
	Iptc2_JobID:                         "JobID",
	Iptc2_MasterDocumentID:              "MasterDocumentID",
	Iptc2_ShortDocumentID:               "ShortDocumentID",
	Iptc2_UniqueDocumentID:              "UniqueDocumentID",
	Iptc2_OwnerID:                       "OwnerID",
	Iptc2_ObjectPreviewFileFormat:       "ObjectPreviewFileFormat",
	Iptc2_ObjectPreviewFileVersion:      "ObjectPreviewFileVersion",
	Iptc2_ObjectPreviewData:             "ObjectPreviewData",
	Iptc2_Prefs:                         "Prefs",
	Iptc2_ClassifyState:                 "ClassifyState",
	Iptc2_SimilarityIndex:               "SimilarityIndex",
	Iptc2_DocumentNotes:                 "DocumentNotes",
	Iptc2_DocumentHistory:               "DocumentHistory",
	Iptc2_ExifCameraInfo:                "ExifCameraInfo",
	Iptc2_CatalogSets:                   "CatalogSets",
}

//IPTCNewsPhoto Tag Ids
const (
	Iptc3_NewsPhotoVersion       uint8 = 0
	Iptc3_IPTCPictureNumber      uint8 = 10
	Iptc3_IPTCImageWidth         uint8 = 20
	Iptc3_IPTCImageHeight        uint8 = 30
	Iptc3_IPTCPixelWidth         uint8 = 40
	Iptc3_IPTCPixelHeight        uint8 = 50
	Iptc3_SupplementalType       uint8 = 55
	Iptc3_ColorRepresentation    uint8 = 60
	Iptc3_InterchangeColorSpace  uint8 = 64
	Iptc3_ColorSequence          uint8 = 65
	Iptc3_ICC_Profile            uint8 = 66
	Iptc3_ColorCalibrationMatrix uint8 = 70
	Iptc3_LookupTable            uint8 = 80
	Iptc3_NumIndexEntries        uint8 = 84
	Iptc3_ColorPalette           uint8 = 85
	Iptc3_IPTCBitsPerSample      uint8 = 86
	Iptc3_SampleStructure        uint8 = 90
	Iptc3_ScanningDirection      uint8 = 100
	Iptc3_IPTCImageRotation      uint8 = 102
	Iptc3_DataCompressionMethod  uint8 = 110
	Iptc3_QuantizationMethod     uint8 = 120
	Iptc3_EndPoints              uint8 = 125
	Iptc3_ExcursionTolerance     uint8 = 130
	Iptc3_BitsPerComponent       uint8 = 135
	Iptc3_MaximumDensityRange    uint8 = 140
	Iptc3_GammaCompensatedValue  uint8 = 145
)

var Iptc3Name = map[uint8]string{
	Iptc3_NewsPhotoVersion:       "NewsPhotoVersion",
	Iptc3_IPTCPictureNumber:      "IPTCPictureNumber",
	Iptc3_IPTCImageWidth:         "IPTCImageWidth",
	Iptc3_IPTCImageHeight:        "IPTCImageHeight",
	Iptc3_IPTCPixelWidth:         "IPTCPixelWidth",
	Iptc3_IPTCPixelHeight:        "IPTCPixelHeight",
	Iptc3_SupplementalType:       "SupplementalType",
	Iptc3_ColorRepresentation:    "ColorRepresentation",
	Iptc3_InterchangeColorSpace:  "InterchangeColorSpace",
	Iptc3_ColorSequence:          "ColorSequence",
	Iptc3_ICC_Profile:            "ICC_Profile",
	Iptc3_ColorCalibrationMatrix: "ColorCalibrationMatrix",
	Iptc3_LookupTable:            "LookupTable",
	Iptc3_NumIndexEntries:        "NumIndexEntries",
	Iptc3_ColorPalette:           "ColorPalette",
	Iptc3_IPTCBitsPerSample:      "IPTCBitsPerSample",
	Iptc3_SampleStructure:        "SampleStructure",
	Iptc3_ScanningDirection:      "ScanningDirection",
	Iptc3_IPTCImageRotation:      "IPTCImageRotation",
	Iptc3_DataCompressionMethod:  "DataCompressionMethod",
	Iptc3_QuantizationMethod:     "QuantizationMethod",
	Iptc3_EndPoints:              "EndPoints",
	Iptc3_ExcursionTolerance:     "ExcursionTolerance",
	Iptc3_BitsPerComponent:       "BitsPerComponent",
	Iptc3_MaximumDensityRange:    "MaximumDensityRange",
	Iptc3_GammaCompensatedValue:  "GammaCompensatedValue",
}
