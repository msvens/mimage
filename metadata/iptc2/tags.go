package iptc2

const Record uint8 = 2

const (
	ApplicationRecordNumber       uint8 = 0
	ObjectTypeReference           uint8 = 3
	ObjectAttributeReference      uint8 = 4
	ObjectName                    uint8 = 5
	EditStatus                    uint8 = 7
	EditorialUpdate               uint8 = 8
	Urgency                       uint8 = 10
	SubjectReference              uint8 = 12
	Category                      uint8 = 15
	SupplementalCategories        uint8 = 20
	FixtureIdentifier             uint8 = 22
	Keywords                      uint8 = 25
	ContentLocationCode           uint8 = 26
	ContentLocationName           uint8 = 27
	ReleaseDate                   uint8 = 30
	ReleaseTime                   uint8 = 35
	ExpirationDate                uint8 = 37
	ExpirationTime                uint8 = 38
	SpecialInstructions           uint8 = 40
	ActionAdvised                 uint8 = 42
	ReferenceService              uint8 = 45
	ReferenceDate                 uint8 = 47
	ReferenceNumber               uint8 = 50
	DateCreated                   uint8 = 55
	TimeCreated                   uint8 = 60
	DigitalCreationDate           uint8 = 62
	DigitalCreationTime           uint8 = 63
	OriginatingProgram            uint8 = 65
	ProgramVersion                uint8 = 70
	ObjectCycle                   uint8 = 75
	Byline                        uint8 = 80
	BylineTitle                   uint8 = 85
	City                          uint8 = 90
	SubLocation                   uint8 = 92
	ProvinceState                 uint8 = 95
	CountryCode                   uint8 = 100
	CountryName                   uint8 = 101
	OriginalTransmissionReference uint8 = 103
	Headline                      uint8 = 105
	Credit                        uint8 = 110
	Source                        uint8 = 115
	CopyrightNotice               uint8 = 116
	Contact                       uint8 = 118
	CaptionAbstract               uint8 = 120
	LocalCaption                  uint8 = 121
	WriterEditor                  uint8 = 122
	RasterizedCaption             uint8 = 125
	ImageType                     uint8 = 130
	ImageOrientation              uint8 = 131
	LanguageIdentifier            uint8 = 135
	AudioType                     uint8 = 150
	AudioSamplingRate             uint8 = 151
	AudioSamplingResolution       uint8 = 152
	AudioDuration                 uint8 = 153
	AudioOutcue                   uint8 = 154
	JobId                         uint8 = 184
	MasterDocumentId              uint8 = 185
	ShortDocumentId               uint8 = 186
	UniqueDocumentId              uint8 = 187
	OwnerId                       uint8 = 188
	ObjectPreviewFileFormat       uint8 = 200
	ObjectPreviewFileVersion      uint8 = 201
	ObjectPreviewData             uint8 = 202
	Prefs                         uint8 = 221
	ClassifyState                 uint8 = 225
	SimilarityIndex               uint8 = 228
	DocumentNotes                 uint8 = 230
	DocumentHistory               uint8 = 231
	ExifCameraInfo                uint8 = 232
	CatalogSets                   uint8 = 255
)

var TagNames = map[uint8]string{
	ApplicationRecordNumber:       "ApplicationRecordNumber",
	ObjectTypeReference:           "ObjectTypeReference",
	ObjectAttributeReference:      "ObjectAttributeReference",
	ObjectName:                    "ObjectName",
	EditStatus:                    "EditStatus",
	EditorialUpdate:               "EditorialUpdate",
	Urgency:                       "Urgency",
	SubjectReference:              "SubjectReference",
	Category:                      "Category",
	SupplementalCategories:        "SupplementalCategories",
	FixtureIdentifier:             "FixtureIdentifier",
	Keywords:                      "Keywords",
	ContentLocationCode:           "ContentLocationCode",
	ContentLocationName:           "ContentLocationName",
	ReleaseDate:                   "ReleaseDate",
	ReleaseTime:                   "ReleaseTime",
	ExpirationDate:                "ExpirationDate",
	ExpirationTime:                "ExpirationTime",
	SpecialInstructions:           "SpecialInstructions",
	ActionAdvised:                 "ActionAdvised",
	ReferenceService:              "ReferenceService",
	ReferenceDate:                 "ReferenceDate",
	ReferenceNumber:               "ReferenceNumber",
	DateCreated:                   "DateCreated",
	TimeCreated:                   "TimeCreated",
	DigitalCreationDate:           "DigitalCreationDate",
	DigitalCreationTime:           "DigitalCreationTime",
	OriginatingProgram:            "OriginatingProgram",
	ProgramVersion:                "ProgramVersion",
	ObjectCycle:                   "ObjectCycle",
	Byline:                        "Byline",
	BylineTitle:                   "BylineTitle",
	City:                          "City",
	SubLocation:                   "SubLocation",
	ProvinceState:                 "ProvinceState",
	CountryCode:                   "CountryCode",
	CountryName:                   "CountryName",
	OriginalTransmissionReference: "OriginalTransmissionReference",
	Headline:                      "Headline",
	Credit:                        "Credit",
	Source:                        "Source",
	CopyrightNotice:               "CopyrightNotice",
	Contact:                       "Contact",
	CaptionAbstract:               "CaptionAbstract",
	LocalCaption:                  "LocalCaption",
	WriterEditor:                  "WriterEditor",
	RasterizedCaption:             "RasterizedCaption",
	ImageType:                     "ImageType",
	ImageOrientation:              "ImageOrientation",
	LanguageIdentifier:            "LanguageIdentifier",
	AudioType:                     "AudioType",
	AudioSamplingRate:             "AudioSamplingRate",
	AudioSamplingResolution:       "AudioSamplingResolution",
	AudioDuration:                 "AudioDuration",
	AudioOutcue:                   "AudioOutcue",
	JobId:                         "JobId",
	MasterDocumentId:              "MasterDocumentId",
	ShortDocumentId:               "ShortDocumentId",
	UniqueDocumentId:              "UniqueDocumentId",
	OwnerId:                       "OwnerId",
	ObjectPreviewFileFormat:       "ObjectPreviewFileFormat",
	ObjectPreviewFileVersion:      "ObjectPreviewFileVersion",
	ObjectPreviewData:             "ObjectPreviewData",
	Prefs:                         "Prefs",
	ClassifyState:                 "ClassifyState",
	SimilarityIndex:               "SimilarityIndex",
	DocumentNotes:                 "DocumentNotes",
	DocumentHistory:               "DocumentHistory",
	ExifCameraInfo:                "ExifCameraInfo",
	CatalogSets:                   "CatalogSets",
}
