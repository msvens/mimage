package iptc1

//Tags extracted from: https://exiftool.org/TagNames/IPTC.html
const (
	EnvelopRecordVersion uint8 = 0
	Destination uint8 = 5
	FileFormat uint8 = 20
	FileVersion uint8 = 22
	ServiceIdentifier uint8 = 30
	EnvelopNumber uint8 = 40
	ProductId uint8 = 50
	EnvelopPriority uint8 = 60
	DateSent uint8 = 70
	TimeSent uint8 = 80
	CodedCharacterSet = 90
	UniqueObjectName uint8 = 100
	ARMIdentifier uint8 = 120
	ARMVersion uint8 = 122
)

const Record uint8 = 1

var TagNames = map[uint8]string{
	EnvelopRecordVersion: "EnvelopRecordVersion",
	Destination: "Destination",
	FileFormat: "FileFormat",
	FileVersion: "FileVersion",
	ServiceIdentifier: "ServiceIdentifier",
	EnvelopNumber: "EnvelopNumber",
	ProductId: "ProductId",
	EnvelopPriority: "EnvelopPriority",
	DateSent: "DateSent",
	TimeSent: "TimeSent",
	CodedCharacterSet: "CodedCharacterSet",
	UniqueObjectName: "UniqueObjectName",
	ARMIdentifier: "ARMIdentifier",
	ARMVersion: "ARMVersion",
}


