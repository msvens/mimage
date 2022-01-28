use feature ':5.10';
use strict;
use JSON;

#Script to extract IPTC Json information from exiftool sources. This script is based on information in
#These sources are used to generate some go sources used by this lib (see internal/generate).
#https://github.com/exiftool/exiftool/blob/master/lib/Image/ExifTool/IPTC.pm

my %fileFormat = (
    0  => 'No ObjectData',
    1  => 'IPTC-NAA Digital Newsphoto Parameter Record',
    2  => 'IPTC7901 Recommended Message Format',
    3  => 'Tagged Image File Format (Adobe/Aldus Image data)',
    4  => 'Illustrator (Adobe Graphics data)',
    5  => 'AppleSingle (Apple Computer Inc)',
    6  => 'NAA 89-3 (ANPA 1312)',
    7  => 'MacBinary II',
    8  => 'IPTC Unstructured Character Oriented File Format (UCOFF)',
    9  => 'United Press International ANPA 1312 variant',
    10 => 'United Press International Down-Load Message',
    11 => 'JPEG File Interchange (JFIF)',
    12 => 'Photo-CD Image-Pac (Eastman Kodak)',
    13 => 'Bit Mapped Graphics File [.BMP] (Microsoft)',
    14 => 'Digital Audio File [.WAV] (Microsoft & Creative Labs)',
    15 => 'Audio plus Moving Video [.AVI] (Microsoft)',
    16 => 'PC DOS/Windows Executable Files [.COM][.EXE]',
    17 => 'Compressed Binary File [.ZIP] (PKWare Inc)',
    18 => 'Audio Interchange File Format AIFF (Apple Computer Inc)',
    19 => 'RIFF Wave (Microsoft Corporation)',
    20 => 'Freehand (Macromedia/Aldus)',
    21 => 'Hypertext Markup Language [.HTML] (The Internet Society)',
    22 => 'MPEG 2 Audio Layer 2 (Musicom), ISO/IEC',
    23 => 'MPEG 2 Audio Layer 3, ISO/IEC',
    24 => 'Portable Document File [.PDF] Adobe',
    25 => 'News Industry Text Format (NITF)',
    26 => 'Tape Archive [.TAR]',
    27 => 'Tidningarnas Telegrambyra NITF version (TTNITF DTD)',
    28 => 'Ritzaus Bureau NITF version (RBNITF DTD)',
    29 => 'Corel Draw [.CDR]',
);


# Record 1 -- EnvelopeRecord
my %EnvelopeRecord = (
    # GROUPS => { 2 => 'Other' },
    # WRITE_PROC => \&WriteIPTC,
    # CHECK_PROC => \&CheckIPTC,
    # WRITABLE => 1,
    0   => {
        Name      => 'EnvelopeRecordVersion',
        Format    => 'int16u',
        Mandatory => 1,
    },
    5   => {
        Name   => 'Destination',
        Flags  => 'List',
        Groups => { 2 => 'Location' },
        Format => 'string[0,1024]',
    },
    20  => {
        Name      => 'FileFormat',
        Groups    => { 2 => 'Image' },
        Format    => 'int16u',
        PrintConv => \%fileFormat,
    },
    22  => {
        Name   => 'FileVersion',
        Groups => { 2 => 'Image' },
        Format => 'int16u',
    },
    30  => {
        Name   => 'ServiceIdentifier',
        Format => 'string[0,10]',
    },
    40  => {
        Name   => 'EnvelopeNumber',
        Format => 'digits[8]',
    },
    50  => {
        Name   => 'ProductID',
        Flags  => 'List',
        Format => 'string[0,32]',
    },
    60  => {
        Name      => 'EnvelopePriority',
        Format    => 'digits[1]',
        PrintConv => {
            0 => '0 (reserved)',
            1 => '1 (most urgent)',
            2 => 2,
            3 => 3,
            4 => 4,
            5 => '5 (normal urgency)',
            6 => 6,
            7 => 7,
            8 => '8 (least urgent)',
            9 => '9 (user-defined priority)',
        },
    },
    70  => {
        Name         => 'DateSent',
        Groups       => { 2 => 'Time' },
        Format       => 'digits[8]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifDate($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcDate($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    80  => {
        Name         => 'TimeSent',
        Groups       => { 2 => 'Time' },
        Format       => 'string[11]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifTime($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcTime($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    90  => {
        Name         => 'CodedCharacterSet',
        Notes        => q{
            values are entered in the form "ESC X Y[, ...]".  The escape sequence for
            UTF-8 character coding is "ESC % G", but this is displayed as "UTF8" for
            convenience.  Either string may be used when writing.  The value of this tag
            affects the decoding of string values in the Application and NewsPhoto
            records.  This tag is marked as "unsafe" to prevent it from being copied by
            default in a group operation because existing tags in the destination image
            may use a different encoding.  When creating a new IPTC record from scratch,
            it is suggested that this be set to "UTF8" if special characters are a
            possibility
        },
        Protected    => 1,
        Format       => 'string[0,32]',
        ValueConvInv => '$val =~ /^UTF-?8$/i ? "\x1b%G" : $val',
        # convert ISO 2022 escape sequences to a more readable format
        PrintConv    => \&PrintCodedCharset,
        PrintConvInv => \&PrintInvCodedCharset,
    },
    100 => {
        Name   => 'UniqueObjectName',
        Format => 'string[14,80]',
    },
    120 => {
        Name   => 'ARMIdentifier',
        Format => 'int16u',
    },
    122 => {
        Name   => 'ARMVersion',
        Format => 'int16u',
    },
);

# Record 2 -- ApplicationRecord
my %ApplicationRecord = (
    # GROUPS => { 2 => 'Other' },
    # WRITE_PROC => \&WriteIPTC,
    # CHECK_PROC => \&CheckIPTC,
    # WRITABLE => 1,
    0   => {
        Name      => 'ApplicationRecordVersion',
        Format    => 'int16u',
        Mandatory => 1,
    },
    3   => {
        Name   => 'ObjectTypeReference',
        Format => 'string[3,67]',
    },
    4   => {
        Name   => 'ObjectAttributeReference',
        Flags  => 'List',
        Format => 'string[4,68]',
    },
    5   => {
        Name   => 'ObjectName',
        Format => 'string[0,64]',
    },
    7   => {
        Name   => 'EditStatus',
        Format => 'string[0,64]',
    },
    8   => {
        Name      => 'EditorialUpdate',
        Format    => 'digits[2]',
        PrintConv => {
            '01' => 'Additional language',
        },
    },
    10  => {
        Name      => 'Urgency',
        Format    => 'digits[1]',
        PrintConv => {
            0 => '0 (reserved)',
            1 => '1 (most urgent)',
            2 => 2,
            3 => 3,
            4 => 4,
            5 => '5 (normal urgency)',
            6 => 6,
            7 => 7,
            8 => '8 (least urgent)',
            9 => '9 (user-defined priority)',
        },
    },
    12  => {
        Name   => 'SubjectReference',
        Flags  => 'List',
        Format => 'string[13,236]',
    },
    15  => {
        Name   => 'Category',
        Format => 'string[0,3]',
    },
    20  => {
        Name   => 'SupplementalCategories',
        Flags  => 'List',
        Format => 'string[0,32]',
    },
    22  => {
        Name   => 'FixtureIdentifier',
        Format => 'string[0,32]',
    },
    25  => {
        Name   => 'Keywords',
        Flags  => 'List',
        Format => 'string[0,64]',
    },
    26  => {
        Name   => 'ContentLocationCode',
        Flags  => 'List',
        Groups => { 2 => 'Location' },
        Format => 'string[3]',
    },
    27  => {
        Name   => 'ContentLocationName',
        Flags  => 'List',
        Groups => { 2 => 'Location' },
        Format => 'string[0,64]',
    },
    30  => {
        Name         => 'ReleaseDate',
        Groups       => { 2 => 'Time' },
        Format       => 'digits[8]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifDate($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcDate($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    35  => {
        Name         => 'ReleaseTime',
        Groups       => { 2 => 'Time' },
        Format       => 'string[11]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifTime($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcTime($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    37  => {
        Name         => 'ExpirationDate',
        Groups       => { 2 => 'Time' },
        Format       => 'digits[8]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifDate($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcDate($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    38  => {
        Name         => 'ExpirationTime',
        Groups       => { 2 => 'Time' },
        Format       => 'string[11]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifTime($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcTime($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    40  => {
        Name   => 'SpecialInstructions',
        Format => 'string[0,256]',
    },
    42  => {
        Name      => 'ActionAdvised',
        Format    => 'digits[2]',
        PrintConv => {
            ''   => '',
            '01' => 'Object Kill',
            '02' => 'Object Replace',
            '03' => 'Object Append',
            '04' => 'Object Reference',
        },
    },
    45  => {
        Name   => 'ReferenceService',
        Flags  => 'List',
        Format => 'string[0,10]',
    },
    47  => {
        Name         => 'ReferenceDate',
        Groups       => { 2 => 'Time' },
        Flags        => 'List',
        Format       => 'digits[8]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifDate($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcDate($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    50  => {
        Name   => 'ReferenceNumber',
        Flags  => 'List',
        Format => 'digits[8]',
    },
    55  => {
        Name         => 'DateCreated',
        Groups       => { 2 => 'Time' },
        Format       => 'digits[8]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifDate($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcDate($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    60  => {
        Name         => 'TimeCreated',
        Groups       => { 2 => 'Time' },
        Format       => 'string[11]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifTime($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcTime($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    62  => {
        Name         => 'DigitalCreationDate',
        Groups       => { 2 => 'Time' },
        Format       => 'digits[8]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifDate($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcDate($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    63  => {
        Name         => 'DigitalCreationTime',
        Groups       => { 2 => 'Time' },
        Format       => 'string[11]',
        Shift        => 'Time',
        ValueConv    => 'Image::ExifTool::Exif::ExifTime($val)',
        ValueConvInv => 'Image::ExifTool::IPTC::IptcTime($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InverseDateOrTime($self,$val)',
    },
    65  => {
        Name   => 'OriginatingProgram',
        Format => 'string[0,32]',
    },
    70  => {
        Name   => 'ProgramVersion',
        Format => 'string[0,10]',
    },
    75  => {
        Name      => 'ObjectCycle',
        Format    => 'string[1]',
        PrintConv => {
            'a' => 'Morning',
            'p' => 'Evening',
            'b' => 'Both Morning and Evening',
        },
    },
    80  => {
        Name   => 'By-line',
        Flags  => 'List',
        Format => 'string[0,32]',
        Groups => { 2 => 'Author' },
    },
    85  => {
        Name   => 'By-lineTitle',
        Flags  => 'List',
        Format => 'string[0,32]',
        Groups => { 2 => 'Author' },
    },
    90  => {
        Name   => 'City',
        Format => 'string[0,32]',
        Groups => { 2 => 'Location' },
    },
    92  => {
        Name   => 'Sub-location',
        Format => 'string[0,32]',
        Groups => { 2 => 'Location' },
    },
    95  => {
        Name   => 'Province-State',
        Format => 'string[0,32]',
        Groups => { 2 => 'Location' },
    },
    100 => {
        Name   => 'Country-PrimaryLocationCode',
        Format => 'string[3]',
        Groups => { 2 => 'Location' },
    },
    101 => {
        Name   => 'Country-PrimaryLocationName',
        Format => 'string[0,64]',
        Groups => { 2 => 'Location' },
    },
    103 => {
        Name   => 'OriginalTransmissionReference',
        Format => 'string[0,32]',
        Notes  => 'now used as a job identifier',
    },
    105 => {
        Name   => 'Headline',
        Format => 'string[0,256]',
    },
    110 => {
        Name   => 'Credit',
        Groups => { 2 => 'Author' },
        Format => 'string[0,32]',
    },
    115 => {
        Name   => 'Source',
        Groups => { 2 => 'Author' },
        Format => 'string[0,32]',
    },
    116 => {
        Name   => 'CopyrightNotice',
        Groups => { 2 => 'Author' },
        Format => 'string[0,128]',
    },
    118 => {
        Name   => 'Contact',
        Flags  => 'List',
        Groups => { 2 => 'Author' },
        Format => 'string[0,128]',
    },
    120 => {
        Name   => 'Caption-Abstract',
        Format => 'string[0,2000]',
    },
    121 => {
        Name   => 'LocalCaption',
        Format => 'string[0,256]', # (guess)
        Notes  => q{
            I haven't found a reference for the format of tags 121, 184-188 and
            225-232, so I have just make them writable as strings with
            reasonable length.  Beware that if this is wrong, other utilities
            may not be able to read these tags as written by ExifTool
        },
    },
    122 => {
        Name   => 'Writer-Editor',
        Flags  => 'List',
        Groups => { 2 => 'Author' },
        Format => 'string[0,32]',
    },
    125 => {
        Name   => 'RasterizedCaption',
        Format => 'undef[7360]',
        Binary => 1,
    },
    130 => {
        Name   => 'ImageType',
        Groups => { 2 => 'Image' },
        Format => 'string[2]',
    },
    131 => {
        Name      => 'ImageOrientation',
        Groups    => { 2 => 'Image' },
        Format    => 'string[1]',
        PrintConv => {
            P => 'Portrait',
            L => 'Landscape',
            S => 'Square',
        },
    },
    135 => {
        Name   => 'LanguageIdentifier',
        Format => 'string[2,3]',
    },
    150 => {
        Name      => 'AudioType',
        Format    => 'string[2]',
        PrintConv => {
            '1A' => 'Mono Actuality',
            '2A' => 'Stereo Actuality',
            '1C' => 'Mono Question and Answer Session',
            '2C' => 'Stereo Question and Answer Session',
            '1M' => 'Mono Music',
            '2M' => 'Stereo Music',
            '1Q' => 'Mono Response to a Question',
            '2Q' => 'Stereo Response to a Question',
            '1R' => 'Mono Raw Sound',
            '2R' => 'Stereo Raw Sound',
            '1S' => 'Mono Scener',
            '2S' => 'Stereo Scener',
            '0T' => 'Text Only',
            '1V' => 'Mono Voicer',
            '2V' => 'Stereo Voicer',
            '1W' => 'Mono Wrap',
            '2W' => 'Stereo Wrap',
        },
    },
    151 => {
        Name   => 'AudioSamplingRate',
        Format => 'digits[6]',
    },
    152 => {
        Name   => 'AudioSamplingResolution',
        Format => 'digits[2]',
    },
    153 => {
        Name   => 'AudioDuration',
        Format => 'digits[6]',
    },
    154 => {
        Name   => 'AudioOutcue',
        Format => 'string[0,64]',
    },
    184 => {
        Name   => 'JobID',
        Format => 'string[0,64]', # (guess)
    },
    185 => {
        Name   => 'MasterDocumentID',
        Format => 'string[0,256]', # (guess)
    },
    186 => {
        Name   => 'ShortDocumentID',
        Format => 'string[0,64]', # (guess)
    },
    187 => {
        Name   => 'UniqueDocumentID',
        Format => 'string[0,128]', # (guess)
    },
    188 => {
        Name   => 'OwnerID',
        Format => 'string[0,128]', # (guess)
    },
    200 => {
        Name      => 'ObjectPreviewFileFormat',
        Groups    => { 2 => 'Image' },
        Format    => 'int16u',
        PrintConv => \%fileFormat,
    },
    201 => {
        Name   => 'ObjectPreviewFileVersion',
        Groups => { 2 => 'Image' },
        Format => 'int16u',
    },
    202 => {
        Name   => 'ObjectPreviewData',
        Groups => { 2 => 'Preview' },
        Format => 'undef[0,256000]',
        Binary => 1,
    },
    221 => {
        Name         => 'Prefs',
        Groups       => { 2 => 'Image' },
        Format       => 'string[0,64]',
        Notes        => 'PhotoMechanic preferences',
        PrintConv    => q{
            $val =~ s[\s*(\d+):\s*(\d+):\s*(\d+):\s*(\S*)]
                     [Tagged:$1, ColorClass:$2, Rating:$3, FrameNum:$4];
            return $val;
        },
        PrintConvInv => q{
            $val =~ s[Tagged:\s*(\d+).*ColorClass:\s*(\d+).*Rating:\s*(\d+).*FrameNum:\s*(\S*)]
                     [$1:$2:$3:$4]is;
            return $val;
        },
    },
    225 => {
        Name   => 'ClassifyState',
        Format => 'string[0,64]', # (guess)
    },
    228 => {
        Name   => 'SimilarityIndex',
        Format => 'string[0,32]', # (guess)
    },
    230 => {
        Name   => 'DocumentNotes',
        Format => 'string[0,1024]', # (guess)
    },
    231 => {
        Name         => 'DocumentHistory',
        Format       => 'string[0,256]',            # (guess)
        ValueConv    => '$val =~ s/\0+/\n/g; $val', # (have seen embedded nulls)
        ValueConvInv => '$val',
    },
    232 => {
        Name   => 'ExifCameraInfo',
        Format => 'string[0,4096]', # (guess)
    },
    255 => { #PH
        Name   => 'CatalogSets',
        List   => 1,
        Format => 'string[0,256]', # (guess)
        Notes  => 'written by iView MediaPro',
    },
);

# Record 3 -- News photo
my %NewsPhoto = (
    #GROUPS => { 2 => 'Image' },
    #WRITE_PROC => \&WriteIPTC,
    #CHECK_PROC => \&CheckIPTC,
    #WRITABLE => 1,
    0   => {
        Name      => 'NewsPhotoVersion',
        Format    => 'int16u',
        Mandatory => 1,
    },
    10  => {
        Name         => 'IPTCPictureNumber',
        Format       => 'string[16]',
        Notes        => '4 numbers: 1-Manufacturer ID, 2-Equipment ID, 3-Date, 4-Sequence',
        PrintConv    => 'Image::ExifTool::IPTC::ConvertPictureNumber($val)',
        PrintConvInv => 'Image::ExifTool::IPTC::InvConvertPictureNumber($val)',
    },
    20  => {
        Name   => 'IPTCImageWidth',
        Format => 'int16u',
    },
    30  => {
        Name   => 'IPTCImageHeight',
        Format => 'int16u',
    },
    40  => {
        Name   => 'IPTCPixelWidth',
        Format => 'int16u',
    },
    50  => {
        Name   => 'IPTCPixelHeight',
        Format => 'int16u',
    },
    55  => {
        Name      => 'SupplementalType',
        Format    => 'int8u',
        PrintConv => {
            0 => 'Main Image',
            1 => 'Reduced Resolution Image',
            2 => 'Logo',
            3 => 'Rasterized Caption',
        },
    },
    60  => {
        Name      => 'ColorRepresentation',
        Format    => 'int16u',
        PrintHex  => 1,
        PrintConv => {
            0x000 => 'No Image, Single Frame',
            0x100 => 'Monochrome, Single Frame',
            0x300 => '3 Components, Single Frame',
            0x301 => '3 Components, Frame Sequential in Multiple Objects',
            0x302 => '3 Components, Frame Sequential in One Object',
            0x303 => '3 Components, Line Sequential',
            0x304 => '3 Components, Pixel Sequential',
            0x305 => '3 Components, Special Interleaving',
            0x400 => '4 Components, Single Frame',
            0x401 => '4 Components, Frame Sequential in Multiple Objects',
            0x402 => '4 Components, Frame Sequential in One Object',
            0x403 => '4 Components, Line Sequential',
            0x404 => '4 Components, Pixel Sequential',
            0x405 => '4 Components, Special Interleaving',
        },
    },
    64  => {
        Name      => 'InterchangeColorSpace',
        Format    => 'int8u',
        PrintConv => {
            1 => 'X,Y,Z CIE',
            2 => 'RGB SMPTE',
            3 => 'Y,U,V (K) (D65)',
            4 => 'RGB Device Dependent',
            5 => 'CMY (K) Device Dependent',
            6 => 'Lab (K) CIE',
            7 => 'YCbCr',
            8 => 'sRGB',
        },
    },
    65  => {
        Name   => 'ColorSequence',
        Format => 'int8u',
    },
    66  => {
        Name     => 'ICC_Profile',
        # ...could add SubDirectory support to read into this (if anybody cares)
        Writable => 0,
        Binary   => 1,
    },
    70  => {
        Name     => 'ColorCalibrationMatrix',
        Writable => 0,
        Binary   => 1,
    },
    80  => {
        Name     => 'LookupTable',
        Writable => 0,
        Binary   => 1,
    },
    84  => {
        Name   => 'NumIndexEntries',
        Format => 'int16u',
    },
    85  => {
        Name     => 'ColorPalette',
        Writable => 0,
        Binary   => 1,
    },
    86  => {
        Name   => 'IPTCBitsPerSample',
        Format => 'int8u',
    },
    90  => {
        Name      => 'SampleStructure',
        Format    => 'int8u',
        PrintConv => {
            0 => 'OrthogonalConstangSampling',
            1 => 'Orthogonal4-2-2Sampling',
            2 => 'CompressionDependent',
        },
    },
    100 => {
        Name      => 'ScanningDirection',
        Format    => 'int8u',
        PrintConv => {
            0 => 'L-R, Top-Bottom',
            1 => 'R-L, Top-Bottom',
            2 => 'L-R, Bottom-Top',
            3 => 'R-L, Bottom-Top',
            4 => 'Top-Bottom, L-R',
            5 => 'Bottom-Top, L-R',
            6 => 'Top-Bottom, R-L',
            7 => 'Bottom-Top, R-L',
        },
    },
    102 => {
        Name      => 'IPTCImageRotation',
        Format    => 'int8u',
        PrintConv => {
            0 => 0,
            1 => 90,
            2 => 180,
            3 => 270,
        },
    },
    110 => {
        Name   => 'DataCompressionMethod',
        Format => 'int32u',
    },
    120 => {
        Name      => 'QuantizationMethod',
        Format    => 'int8u',
        PrintConv => {
            0 => 'Linear Reflectance/Transmittance',
            1 => 'Linear Density',
            2 => 'IPTC Ref B',
            3 => 'Linear Dot Percent',
            4 => 'AP Domestic Analogue',
            5 => 'Compression Method Specific',
            6 => 'Color Space Specific',
            7 => 'Gamma Compensated',
        },
    },
    125 => {
        Name     => 'EndPoints',
        Writable => 0,
        Binary   => 1,
    },
    130 => {
        Name      => 'ExcursionTolerance',
        Format    => 'int8u',
        PrintConv => {
            0 => 'Not Allowed',
            1 => 'Allowed',
        },
    },
    135 => {
        Name   => 'BitsPerComponent',
        Format => 'int8u',
    },
    140 => {
        Name   => 'MaximumDensityRange',
        Format => 'int16u',
    },
    145 => {
        Name   => 'GammaCompensatedValue',
        Format => 'int16u',
    },
);

# Record 7 -- Pre-object Data
my %PreObjectData = (
    # (not actually writable, but used in BuildTagLookup to recognize IPTC tables)
    # WRITE_PROC => \&WriteIPTC,
    10 => {
        Name      => 'SizeMode',
        Format    => 'int8u',
        PrintConv => {
            0 => 'Size Not Known',
            1 => 'Size Known',
        },
    },
    20 => {
        Name   => 'MaxSubfileSize',
        Format => 'int32u',
    },
    90 => {
        Name   => 'ObjectSizeAnnounced',
        Format => 'int32u',
    },
    95 => {
        Name   => 'MaximumObjectSize',
        Format => 'int32u',
    },
);

# Record 8 -- ObjectData
my %ObjectData = (
    #WRITE_PROC => \&WriteIPTC,
    10 => {
        Name   => 'SubFile',
        Flags  => 'List',
        Binary => 1,
    },
);

# Record 9 -- PostObjectData
my %PostObjectData = (
    #WRITE_PROC => \&WriteIPTC,
    10 => {
        Name   => 'ConfirmedObjectSize',
        Format => 'int32u',
    },
);

# Record 240 -- FotoStation proprietary data (ref PH)
my %FotoStation = (
    #GROUPS => { 2 => 'Other' },
    #WRITE_PROC => \&WriteIPTC,
    #CHECK_PROC => \&CheckIPTC,
    #WRITABLE => 1,
);

# # IPTC Composite tags
# my %Composite = (
#     GROUPS => { 2 => 'Image' },
#     DateTimeCreated => {
#         Description => 'Date/Time Created',
#         Groups => { 2 => 'Time' },
#         Require => {
#             0 => 'IPTC:DateCreated',
#             1 => 'IPTC:TimeCreated',
#         },
#         ValueConv => '"$val[0] $val[1]"',
#         PrintConv => '$self->ConvertDateTime($val)',
#     },
#     DigitalCreationDateTime => {
#         Description => 'Digital Creation Date/Time',
#         Groups => { 2 => 'Time' },
#         Require => {
#             0 => 'IPTC:DigitalCreationDate',
#             1 => 'IPTC:DigitalCreationTime',
#         },
#         ValueConv => '"$val[0] $val[1]"',
#         PrintConv => '$self->ConvertDateTime($val)',
#     },
# );


my %IptcTags = (
    1   => {
        Name => 'IPTCEnvelope',
        Tags => \%EnvelopeRecord,
        # SubDirectory => {
        #     TagTable => 'Image::ExifTool::IPTC::EnvelopeRecord',
        # },
    },
    2   => {
        Name => 'IPTCApplication',
        Tags => \%ApplicationRecord,
    },
    3   => {
        Name => 'IPTCNewsPhoto',
        Tags => \%NewsPhoto,
    },
    7   => {
        Name => 'IPTCPreObjectData',
        Tags => \%PreObjectData,
    },
    8   => {
        Name => 'IPTCObjectData',
        Tags => \%ObjectData,
    },
    9   => {
        Name => 'IPTCPostObjectData',
        Tags => \%PostObjectData,
    },
    240 => {
        Name => 'IPTCFotoStation',
        Tags => \%FotoStation,
    },
);

my %allTags;

sub generateTags {
    my %record = %{$_[0]};
    my %table;

    for my $k (keys %record) {
        my $id = $k + 0;
        my $n = $record{$k}->{'Name'};
        my $fmt = $record{$k}->{'Format'} . "";
        my $notes = $record{$k}->{'Notes'} . "";
        my $flags = $record{$k}->{'Flags'} . "";
        my $mandatory = \0;
        if (exists $record{$k}->{'Mandatory'}) {
            $mandatory = \$record{$k}->{'Mandatory'}
        }
        #my $mandatory = $record{$k}->{'Mandatory'};
        my $binary = \0;
        if (exists $record{$k}->{'Binary'}) {
            $binary = \$record{$k}->{'Binary'}
        }
        #my $binary = $record{$k}->{'Binary'};
        my $writable = \1;
        if (exists $record{$k}->{'Writable'}) {
            $writable = \$record{$k}->{'Writable'}
        }
        #my $writable = $record{$k}->{'Writable'}+0;
        my $values = {};
        if (exists $record{$k}->{'PrintConv'}) {
            my $pc = $record{$k}->{'PrintConv'};
            if (UNIVERSAL::isa($pc, 'HASH')) {
                for my $vk (keys %$pc) {
                    #say "$vk"
                    $values->{$vk} = $pc->{$vk}.""
                }
                #$values = $pc
            }
        }
        $table{$id} = {
            id        => $id,
            name      => $n,
            fmt       => $fmt,
            flags     => $flags,
            mandatory => $mandatory,
            notes     => $notes,
            values    => $values,
            binary    => $binary,
            writable  => $writable,
        };
    }
    my @recs = values %table;
    return \@recs
}

for my $recId (keys %IptcTags) {
    my $rec = $IptcTags{$recId};
    my $newRec = {
        id => $recId + 0
    };
    $newRec->{'name'} = $rec->{'Name'};

    $newRec->{'tags'} = generateTags($rec->{'Tags'});
    #say "$rec->{'Name'}";
    $allTags{$rec->{'Name'}} = $newRec
}

my $jsonText = to_json(\%allTags, { utf8 => 1, pretty => 1, canonical => 1 });
say "$jsonText";

# my $js = JSON->new;
# $js->canonical(1);
#
# my $output = $js->encode($h);
# print $output;


# for my $k (keys %envelopeRecord) {
#     my $rec = $envelopeRecord{$k};
#     for my $kr (keys %$rec) {
#         say "$kr"
#     }
# }

# my %table;
#
# for my $k (keys %envelopeRecord) {
#    my $n = $envelopeRecord{$k}->{'Name'};
#    my $fmt = $envelopeRecord{$k}->{'Format'};
#     my $notes = $envelopeRecord{$k}->{'Notes'};
#     my $flags = $envelopeRecord{$k}->{'Flags'};
#     my $mandatory = $envelopeRecord{$k}->{'Mandatory'} + 0;
#     my $printconv = {};
#     if (exists $envelopeRecord{$k}->{'PrintConv'}) {
#         my $pc = $envelopeRecord{$k}->{'PrintConv'};
#         if ( UNIVERSAL::isa($pc, 'HASH')) {
#             $printconv = $pc
#         }
#     }
#     $table{$n} = {
#         id        => $k + 0,
#         fmt       => $fmt,
#         flags     => $flags."",
#         mandatory => $mandatory,
#         notes     => $notes."",
#         printconv => $printconv
#     };
#     say "$flags"
# }
#
# my $jsonText = to_json(\%table, {utf8 => 1, pretty => 1});
# say "jsonText $jsonText";

#for my $k (keys %fileFormat) {
#    say "$fileFormat{$k}\n"
#}
