package metadata
//Do not edit! This is an automatically generated file (see generator.GenerateIptcTags()).
//This file was generated based on https://exiftool.org/TagNames/IPTC.html


//IPTC Records
type IptcRecord uint8
const(
  Envelope IptcRecord = 1
  Application IptcRecord = 2
  NewsPhoto IptcRecord = 3
  PreObjectData IptcRecord = 7
  ObjectData IptcRecord = 8
  PostObjectData IptcRecord = 9
  FotoStation IptcRecord = 240
)

var IptcRecordName = map[IptcRecord]string{
  Envelope: "IPTCEnvelop",
  Application: "IPTCApplication",
  NewsPhoto: "IPTCNewsPhoto",
  PreObjectData: "IPTCPreObjectData",
  ObjectData: "IPTCObjectData",
  PostObjectData: "IPTCPostObjectData",
  FotoStation: "IPTCFotoStation",
}

type IptcTag uint8


//Envelope Tag Ids
const(
  Envelope_EnvelopeRecordVersion IptcTag = 0
  Envelope_Destination IptcTag = 5
  Envelope_FileFormat IptcTag = 20
  Envelope_FileVersion IptcTag = 22
  Envelope_ServiceIdentifier IptcTag = 30
  Envelope_EnvelopeNumber IptcTag = 40
  Envelope_ProductID IptcTag = 50
  Envelope_EnvelopePriority IptcTag = 60
  Envelope_DateSent IptcTag = 70
  Envelope_TimeSent IptcTag = 80
  Envelope_CodedCharacterSet IptcTag = 90
  Envelope_UniqueObjectName IptcTag = 100
  Envelope_ARMIdentifier IptcTag = 120
  Envelope_ARMVersion IptcTag = 122
)

var IptcEnvelopeName = map[IptcTag]string{
  Envelope_EnvelopeRecordVersion: "EnvelopeRecordVersion",
  Envelope_Destination: "Destination",
  Envelope_FileFormat: "FileFormat",
  Envelope_FileVersion: "FileVersion",
  Envelope_ServiceIdentifier: "ServiceIdentifier",
  Envelope_EnvelopeNumber: "EnvelopeNumber",
  Envelope_ProductID: "ProductID",
  Envelope_EnvelopePriority: "EnvelopePriority",
  Envelope_DateSent: "DateSent",
  Envelope_TimeSent: "TimeSent",
  Envelope_CodedCharacterSet: "CodedCharacterSet",
  Envelope_UniqueObjectName: "UniqueObjectName",
  Envelope_ARMIdentifier: "ARMIdentifier",
  Envelope_ARMVersion: "ARMVersion",
}

//Application Tag Ids
const(
  Application_ApplicationRecordVersion IptcTag = 0
  Application_ObjectTypeReference IptcTag = 3
  Application_ObjectAttributeReference IptcTag = 4
  Application_ObjectName IptcTag = 5
  Application_EditStatus IptcTag = 7
  Application_EditorialUpdate IptcTag = 8
  Application_Urgency IptcTag = 10
  Application_SubjectReference IptcTag = 12
  Application_Category IptcTag = 15
  Application_SupplementalCategories IptcTag = 20
  Application_FixtureIdentifier IptcTag = 22
  Application_Keywords IptcTag = 25
  Application_ContentLocationCode IptcTag = 26
  Application_ContentLocationName IptcTag = 27
  Application_ReleaseDate IptcTag = 30
  Application_ReleaseTime IptcTag = 35
  Application_ExpirationDate IptcTag = 37
  Application_ExpirationTime IptcTag = 38
  Application_SpecialInstructions IptcTag = 40
  Application_ActionAdvised IptcTag = 42
  Application_ReferenceService IptcTag = 45
  Application_ReferenceDate IptcTag = 47
  Application_ReferenceNumber IptcTag = 50
  Application_DateCreated IptcTag = 55
  Application_TimeCreated IptcTag = 60
  Application_DigitalCreationDate IptcTag = 62
  Application_DigitalCreationTime IptcTag = 63
  Application_OriginatingProgram IptcTag = 65
  Application_ProgramVersion IptcTag = 70
  Application_ObjectCycle IptcTag = 75
  Application_Byline IptcTag = 80
  Application_BylineTitle IptcTag = 85
  Application_City IptcTag = 90
  Application_Sublocation IptcTag = 92
  Application_ProvinceState IptcTag = 95
  Application_CountryPrimaryLocationCode IptcTag = 100
  Application_CountryPrimaryLocationName IptcTag = 101
  Application_OriginalTransmissionReference IptcTag = 103
  Application_Headline IptcTag = 105
  Application_Credit IptcTag = 110
  Application_Source IptcTag = 115
  Application_CopyrightNotice IptcTag = 116
  Application_Contact IptcTag = 118
  Application_CaptionAbstract IptcTag = 120
  Application_LocalCaption IptcTag = 121
  Application_WriterEditor IptcTag = 122
  Application_RasterizedCaption IptcTag = 125
  Application_ImageType IptcTag = 130
  Application_ImageOrientation IptcTag = 131
  Application_LanguageIdentifier IptcTag = 135
  Application_AudioType IptcTag = 150
  Application_AudioSamplingRate IptcTag = 151
  Application_AudioSamplingResolution IptcTag = 152
  Application_AudioDuration IptcTag = 153
  Application_AudioOutcue IptcTag = 154
  Application_JobID IptcTag = 184
  Application_MasterDocumentID IptcTag = 185
  Application_ShortDocumentID IptcTag = 186
  Application_UniqueDocumentID IptcTag = 187
  Application_OwnerID IptcTag = 188
  Application_ObjectPreviewFileFormat IptcTag = 200
  Application_ObjectPreviewFileVersion IptcTag = 201
  Application_ObjectPreviewData IptcTag = 202
  Application_Prefs IptcTag = 221
  Application_ClassifyState IptcTag = 225
  Application_SimilarityIndex IptcTag = 228
  Application_DocumentNotes IptcTag = 230
  Application_DocumentHistory IptcTag = 231
  Application_ExifCameraInfo IptcTag = 232
  Application_CatalogSets IptcTag = 255
)

var IptcApplicationName = map[IptcTag]string{
  Application_ApplicationRecordVersion: "ApplicationRecordVersion",
  Application_ObjectTypeReference: "ObjectTypeReference",
  Application_ObjectAttributeReference: "ObjectAttributeReference",
  Application_ObjectName: "ObjectName",
  Application_EditStatus: "EditStatus",
  Application_EditorialUpdate: "EditorialUpdate",
  Application_Urgency: "Urgency",
  Application_SubjectReference: "SubjectReference",
  Application_Category: "Category",
  Application_SupplementalCategories: "SupplementalCategories",
  Application_FixtureIdentifier: "FixtureIdentifier",
  Application_Keywords: "Keywords",
  Application_ContentLocationCode: "ContentLocationCode",
  Application_ContentLocationName: "ContentLocationName",
  Application_ReleaseDate: "ReleaseDate",
  Application_ReleaseTime: "ReleaseTime",
  Application_ExpirationDate: "ExpirationDate",
  Application_ExpirationTime: "ExpirationTime",
  Application_SpecialInstructions: "SpecialInstructions",
  Application_ActionAdvised: "ActionAdvised",
  Application_ReferenceService: "ReferenceService",
  Application_ReferenceDate: "ReferenceDate",
  Application_ReferenceNumber: "ReferenceNumber",
  Application_DateCreated: "DateCreated",
  Application_TimeCreated: "TimeCreated",
  Application_DigitalCreationDate: "DigitalCreationDate",
  Application_DigitalCreationTime: "DigitalCreationTime",
  Application_OriginatingProgram: "OriginatingProgram",
  Application_ProgramVersion: "ProgramVersion",
  Application_ObjectCycle: "ObjectCycle",
  Application_Byline: "Byline",
  Application_BylineTitle: "BylineTitle",
  Application_City: "City",
  Application_Sublocation: "Sublocation",
  Application_ProvinceState: "ProvinceState",
  Application_CountryPrimaryLocationCode: "CountryPrimaryLocationCode",
  Application_CountryPrimaryLocationName: "CountryPrimaryLocationName",
  Application_OriginalTransmissionReference: "OriginalTransmissionReference",
  Application_Headline: "Headline",
  Application_Credit: "Credit",
  Application_Source: "Source",
  Application_CopyrightNotice: "CopyrightNotice",
  Application_Contact: "Contact",
  Application_CaptionAbstract: "CaptionAbstract",
  Application_LocalCaption: "LocalCaption",
  Application_WriterEditor: "WriterEditor",
  Application_RasterizedCaption: "RasterizedCaption",
  Application_ImageType: "ImageType",
  Application_ImageOrientation: "ImageOrientation",
  Application_LanguageIdentifier: "LanguageIdentifier",
  Application_AudioType: "AudioType",
  Application_AudioSamplingRate: "AudioSamplingRate",
  Application_AudioSamplingResolution: "AudioSamplingResolution",
  Application_AudioDuration: "AudioDuration",
  Application_AudioOutcue: "AudioOutcue",
  Application_JobID: "JobID",
  Application_MasterDocumentID: "MasterDocumentID",
  Application_ShortDocumentID: "ShortDocumentID",
  Application_UniqueDocumentID: "UniqueDocumentID",
  Application_OwnerID: "OwnerID",
  Application_ObjectPreviewFileFormat: "ObjectPreviewFileFormat",
  Application_ObjectPreviewFileVersion: "ObjectPreviewFileVersion",
  Application_ObjectPreviewData: "ObjectPreviewData",
  Application_Prefs: "Prefs",
  Application_ClassifyState: "ClassifyState",
  Application_SimilarityIndex: "SimilarityIndex",
  Application_DocumentNotes: "DocumentNotes",
  Application_DocumentHistory: "DocumentHistory",
  Application_ExifCameraInfo: "ExifCameraInfo",
  Application_CatalogSets: "CatalogSets",
}

//NewsPhoto Tag Ids
const(
  NewsPhoto_NewsPhotoVersion IptcTag = 0
  NewsPhoto_IPTCPictureNumber IptcTag = 10
  NewsPhoto_IPTCImageWidth IptcTag = 20
  NewsPhoto_IPTCImageHeight IptcTag = 30
  NewsPhoto_IPTCPixelWidth IptcTag = 40
  NewsPhoto_IPTCPixelHeight IptcTag = 50
  NewsPhoto_SupplementalType IptcTag = 55
  NewsPhoto_ColorRepresentation IptcTag = 60
  NewsPhoto_InterchangeColorSpace IptcTag = 64
  NewsPhoto_ColorSequence IptcTag = 65
  NewsPhoto_ICC_Profile IptcTag = 66
  NewsPhoto_ColorCalibrationMatrix IptcTag = 70
  NewsPhoto_LookupTable IptcTag = 80
  NewsPhoto_NumIndexEntries IptcTag = 84
  NewsPhoto_ColorPalette IptcTag = 85
  NewsPhoto_IPTCBitsPerSample IptcTag = 86
  NewsPhoto_SampleStructure IptcTag = 90
  NewsPhoto_ScanningDirection IptcTag = 100
  NewsPhoto_IPTCImageRotation IptcTag = 102
  NewsPhoto_DataCompressionMethod IptcTag = 110
  NewsPhoto_QuantizationMethod IptcTag = 120
  NewsPhoto_EndPoints IptcTag = 125
  NewsPhoto_ExcursionTolerance IptcTag = 130
  NewsPhoto_BitsPerComponent IptcTag = 135
  NewsPhoto_MaximumDensityRange IptcTag = 140
  NewsPhoto_GammaCompensatedValue IptcTag = 145
)

var IptcNewsPhotoName = map[IptcTag]string{
  NewsPhoto_NewsPhotoVersion: "NewsPhotoVersion",
  NewsPhoto_IPTCPictureNumber: "IPTCPictureNumber",
  NewsPhoto_IPTCImageWidth: "IPTCImageWidth",
  NewsPhoto_IPTCImageHeight: "IPTCImageHeight",
  NewsPhoto_IPTCPixelWidth: "IPTCPixelWidth",
  NewsPhoto_IPTCPixelHeight: "IPTCPixelHeight",
  NewsPhoto_SupplementalType: "SupplementalType",
  NewsPhoto_ColorRepresentation: "ColorRepresentation",
  NewsPhoto_InterchangeColorSpace: "InterchangeColorSpace",
  NewsPhoto_ColorSequence: "ColorSequence",
  NewsPhoto_ICC_Profile: "ICC_Profile",
  NewsPhoto_ColorCalibrationMatrix: "ColorCalibrationMatrix",
  NewsPhoto_LookupTable: "LookupTable",
  NewsPhoto_NumIndexEntries: "NumIndexEntries",
  NewsPhoto_ColorPalette: "ColorPalette",
  NewsPhoto_IPTCBitsPerSample: "IPTCBitsPerSample",
  NewsPhoto_SampleStructure: "SampleStructure",
  NewsPhoto_ScanningDirection: "ScanningDirection",
  NewsPhoto_IPTCImageRotation: "IPTCImageRotation",
  NewsPhoto_DataCompressionMethod: "DataCompressionMethod",
  NewsPhoto_QuantizationMethod: "QuantizationMethod",
  NewsPhoto_EndPoints: "EndPoints",
  NewsPhoto_ExcursionTolerance: "ExcursionTolerance",
  NewsPhoto_BitsPerComponent: "BitsPerComponent",
  NewsPhoto_MaximumDensityRange: "MaximumDensityRange",
  NewsPhoto_GammaCompensatedValue: "GammaCompensatedValue",
}

//PreObjectData Tag Ids
const(
  PreObjectData_SizeMode IptcTag = 10
  PreObjectData_MaxSubfileSize IptcTag = 20
  PreObjectData_ObjectSizeAnnounced IptcTag = 90
  PreObjectData_MaximumObjectSize IptcTag = 95
)

var IptcPreObjectDataName = map[IptcTag]string{
  PreObjectData_SizeMode: "SizeMode",
  PreObjectData_MaxSubfileSize: "MaxSubfileSize",
  PreObjectData_ObjectSizeAnnounced: "ObjectSizeAnnounced",
  PreObjectData_MaximumObjectSize: "MaximumObjectSize",
}

//ObjectData Tag Ids
const(
  ObjectData_SubFile IptcTag = 10
)

var IptcObjectDataName = map[IptcTag]string{
  ObjectData_SubFile: "SubFile",
}

//PostObjectData Tag Ids
const(
  PostObjectData_ConfirmedObjectSize IptcTag = 10
)

var IptcPostObjectDataName = map[IptcTag]string{
  PostObjectData_ConfirmedObjectSize: "ConfirmedObjectSize",
}

//FotoStation Tag Ids
const(
)

var IptcFotoStationName = map[IptcTag]string{
}

