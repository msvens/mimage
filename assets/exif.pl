#!/usr/bin/perl
use feature ':5.10';
use strict;
use warnings FATAL => 'all';
use JSON;

#Changes from original source
#1. 0x927c MakerNote  is linked to another table (MakerNote.pm). Replaced by a single tag record
#2. 0x14a SubIFDs Array was replaced by a single tag record
#3. 0x117 StripByteCounts Array was replaced by a single tag record
#4. 0x201 ThumbnailOffset Array was replaced by a single tag record
#5. 0xc634 DNGPrivateData Array was replaced by a single tag record
#6. 0x202 ThumbnailLength Array was replaced by a single tag record


# EXIF LightSource PrintConv values
my %lightSource = (
    0 => 'Unknown',
    1 => 'Daylight',
    2 => 'Fluorescent',
    3 => 'Tungsten (Incandescent)',
    4 => 'Flash',
    9 => 'Fine Weather',
    10 => 'Cloudy',
    11 => 'Shade',
    12 => 'Daylight Fluorescent',   # (D 5700 - 7100K)
    13 => 'Day White Fluorescent',  # (N 4600 - 5500K)
    14 => 'Cool White Fluorescent', # (W 3800 - 4500K)
    15 => 'White Fluorescent',      # (WW 3250 - 3800K)
    16 => 'Warm White Fluorescent', # (L 2600 - 3250K)
    17 => 'Standard Light A',
    18 => 'Standard Light B',
    19 => 'Standard Light C',
    20 => 'D55',
    21 => 'D65',
    22 => 'D75',
    23 => 'D50',
    24 => 'ISO Studio Tungsten',
    255 => 'Other',
);

# EXIF Flash values
my %flash = (
    0x00 => 'No Flash',
    0x01 => 'Fired',
    0x05 => 'Fired, Return not detected',
    0x07 => 'Fired, Return detected',
    0x08 => 'On, Did not fire', # not charged up?
    0x09 => 'On, Fired',
    0x0d => 'On, Return not detected',
    0x0f => 'On, Return detected',
    0x10 => 'Off, Did not fire',
    0x14 => 'Off, Did not fire, Return not detected',
    0x18 => 'Auto, Did not fire',
    0x19 => 'Auto, Fired',
    0x1d => 'Auto, Fired, Return not detected',
    0x1f => 'Auto, Fired, Return detected',
    0x20 => 'No flash function',
    0x30 => 'Off, No flash function',
    0x41 => 'Fired, Red-eye reduction',
    0x45 => 'Fired, Red-eye reduction, Return not detected',
    0x47 => 'Fired, Red-eye reduction, Return detected',
    0x49 => 'On, Red-eye reduction',
    0x4d => 'On, Red-eye reduction, Return not detected',
    0x4f => 'On, Red-eye reduction, Return detected',
    0x50 => 'Off, Red-eye reduction',
    0x58 => 'Auto, Did not fire, Red-eye reduction',
    0x59 => 'Auto, Fired, Red-eye reduction',
    0x5d => 'Auto, Fired, Red-eye reduction, Return not detected',
    0x5f => 'Auto, Fired, Red-eye reduction, Return detected',
);

# TIFF Compression values
# (values with format "Xxxxx XXX Compressed" are used to identify RAW file types)
my %compression = (
    1 => 'Uncompressed',
    2 => 'CCITT 1D',
    3 => 'T4/Group 3 Fax',
    4 => 'T6/Group 4 Fax',
    5 => 'LZW',
    6 => 'JPEG (old-style)', #3
    7 => 'JPEG', #4
    8 => 'Adobe Deflate', #3
    9 => 'JBIG B&W', #3
    10 => 'JBIG Color', #3
    99 => 'JPEG', #16
    262 => 'Kodak 262', #16
    32766 => 'Next', #3
    32767 => 'Sony ARW Compressed', #16
    32769 => 'Packed RAW', #PH (used by Epson, Nikon, Samsung)
    32770 => 'Samsung SRW Compressed', #PH
    32771 => 'CCIRLEW', #3
    32772 => 'Samsung SRW Compressed 2', #PH (NX3000,NXmini)
    32773 => 'PackBits',
    32809 => 'Thunderscan', #3
    32867 => 'Kodak KDC Compressed', #PH
    32895 => 'IT8CTPAD', #3
    32896 => 'IT8LW', #3
    32897 => 'IT8MP', #3
    32898 => 'IT8BL', #3
    32908 => 'PixarFilm', #3
    32909 => 'PixarLog', #3
    # 32910,32911 - Pixar reserved
    32946 => 'Deflate', #3
    32947 => 'DCS', #3
    33003 => 'Aperio JPEG 2000 YCbCr', #https://openslide.org/formats/aperio/
    33005 => 'Aperio JPEG 2000 RGB', #https://openslide.org/formats/aperio/
    34661 => 'JBIG', #3
    34676 => 'SGILog', #3
    34677 => 'SGILog24', #3
    34712 => 'JPEG 2000', #3
    34713 => 'Nikon NEF Compressed', #PH
    34715 => 'JBIG2 TIFF FX', #20
    34718 => 'Microsoft Document Imaging (MDI) Binary Level Codec', #18
    34719 => 'Microsoft Document Imaging (MDI) Progressive Transform Codec', #18
    34720 => 'Microsoft Document Imaging (MDI) Vector', #18
    34887 => 'ESRI Lerc', #LibTiff
    # 34888,34889 - ESRI reserved
    34892 => 'Lossy JPEG', # (DNG 1.4)
    34925 => 'LZMA2', #LibTiff
    34926 => 'Zstd', #LibTiff
    34927 => 'WebP', #LibTiff
    34933 => 'PNG', # (TIFF mail list)
    34934 => 'JPEG XR', # (TIFF mail list)
    65000 => 'Kodak DCR Compressed', #PH
    65535 => 'Pentax PEF Compressed', #Jens
);

my %photometricInterpretation = (
    0 => 'WhiteIsZero',
    1 => 'BlackIsZero',
    2 => 'RGB',
    3 => 'RGB Palette',
    4 => 'Transparency Mask',
    5 => 'CMYK',
    6 => 'YCbCr',
    8 => 'CIELab',
    9 => 'ICCLab', #3
    10 => 'ITULab', #3
    32803 => 'Color Filter Array', #2
    32844 => 'Pixar LogL', #3
    32845 => 'Pixar LogLuv', #3
    32892 => 'Sequential Color Filter', #JR (Sony ARQ)
    34892 => 'Linear Raw', #2
    51177 => 'Depth Map', # (DNG 1.5)
    52527 => 'Semantic Mask', # (DNG 1.6)
);

my %orientation = (
    1 => 'Horizontal (normal)',
    2 => 'Mirror horizontal',
    3 => 'Rotate 180',
    4 => 'Mirror vertical',
    5 => 'Mirror horizontal and rotate 270 CW',
    6 => 'Rotate 90 CW',
    7 => 'Mirror horizontal and rotate 90 CW',
    8 => 'Rotate 270 CW',
);

my %subfileType = (
    0 => 'Full-resolution image',
    1 => 'Reduced-resolution image',
    2 => 'Single page of multi-page image',
    3 => 'Single page of multi-page reduced-resolution image',
    4 => 'Transparency mask',
    5 => 'Transparency mask of reduced-resolution image',
    6 => 'Transparency mask of multi-page image',
    7 => 'Transparency mask of reduced-resolution multi-page image',
    8 => 'Depth map', # (DNG 1.5)
    9 => 'Depth map of reduced-resolution image', # (DNG 1.5)
    16 => 'Enhanced image data', # (DNG 1.5)
    0x10001 => 'Alternate reduced-resolution image', # (DNG 1.2)
    0x10004 => 'Semantic Mask', # (DNG 1.6)
    0xffffffff => 'invalid', #(found in E5700 NEF's)
    BITMASK => {
        0 => 'Reduced resolution',
        1 => 'Single page',
        2 => 'Transparency mask',
        3 => 'TIFF/IT final page', #20 (repurposed as DepthMap repurposes by DNG 1.5)
        4 => 'TIFF-FX mixed raster content', #20 (repurposed as EnhancedImageData by DNG 1.5)
    },
);

# convert DNG UTF-8 string values (may be string or int8u format)
my %utf8StringConv = (
    Writable => 'string',
    Format => 'string',
    ValueConv => '$self->Decode($val, "UTF8")',
    ValueConvInv => '$self->Encode($val,"UTF8")',
);

# ValueConv that makes long values binary type
my %longBin = (
    ValueConv => 'length($val) > 64 ? \$val : $val',
    ValueConvInv => '$val',
    LongBinary => 1,        # flag to avoid decoding values of a large array
);

# PrintConv for SampleFormat (0x153)
my %sampleFormat = (
    1 => 'Unsigned',        # unsigned integer
    2 => 'Signed',          # two's complement signed integer
    3 => 'Float',           # IEEE floating point
    4 => 'Undefined',
    5 => 'Complex int',     # complex integer (ref 3)
    6 => 'Complex float',   # complex IEEE floating point (ref 3)
);

# save the values of these tags for additional validation checks
my %saveForValidate = (
    0x100 => 1, # ImageWidth
    0x101 => 1, # ImageHeight
    0x102 => 1, # BitsPerSample
    0x103 => 1, # Compression
    0x115 => 1, # SamplesPerPixel
);

# conversions for DNG OpcodeList tags
my %opcodeInfo = (
    Writable => 'undef',
    WriteGroup => 'SubIFD',
    Protected => 1,
    Binary => 1,
    ConvertBinary => 1, # needed because the ValueConv value is binary
    PrintConvColumns => 2,
    PrintConv => {
        OTHER => \&PrintOpcode,
        1 => 'WarpRectilinear',
        2 => 'WarpFisheye',
        3 => 'FixVignetteRadial',
        4 => 'FixBadPixelsConstant',
        5 => 'FixBadPixelsList',
        6 => 'TrimBounds',
        7 => 'MapTable',
        8 => 'MapPolynomial',
        9 => 'GainMap',
        10 => 'DeltaPerRow',
        11 => 'DeltaPerColumn',
        12 => 'ScalePerRow',
        13 => 'ScalePerColumn',
        14 => 'WarpRectilinear2', # (DNG 1.6)
    },
    PrintConvInv => undef,  # (so the inverse conversion is not performed)
);

# main EXIF tag table
my %ExifTable = (
    #GROUPS => { 0 => 'EXIF', 1 => 'IFD0', 2 => 'Image'},
    #WRITE_PROC => \&WriteExif,
    #CHECK_PROC => \&CheckExif,
    #WRITE_GROUP => 'ExifIFD',   # default write group
    #SET_GROUP1 => 1, # set group1 name to directory name for all tags in table
    0x1    => {
        Name        => 'InteropIndex',
        Description => 'Interoperability Index',
        Protected   => 1,
        Writable    => 'string',
        WriteGroup  => 'InteropIFD',
        PrintConv   => {
            R98 => 'R98 - DCF basic file (sRGB)',
            R03 => 'R03 - DCF option file (Adobe RGB)',
            THM => 'THM - DCF thumbnail file',
        },
    },
    0x2    => { #5
        Name        => 'InteropVersion',
        Description => 'Interoperability Version',
        Protected   => 1,
        Writable    => 'undef',
        Mandatory   => 1,
        WriteGroup  => 'InteropIFD',
        RawConv     => '$val=~s/\0+$//; $val', # (some idiots add null terminators)
    },
    0x0b   => { #PH
        Name       => 'ProcessingSoftware',
        Writable   => 'string',
        WriteGroup => 'IFD0',
        Notes      => 'used by ACD Systems Digital Imaging',
    },
    0xfe   => {
        Name       => 'SubfileType',
        Notes      => 'called NewSubfileType by the TIFF specification',
        Protected  => 1,
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
        # set priority directory if this is the full resolution image
        DataMember => 'SubfileType',
        RawConv    => '$self->SetPriorityDir() if $val eq "0"; $$self{SubfileType} = $val',
        PrintConv  => \%subfileType,
    },
    0xff   => {
        Name       => 'OldSubfileType',
        Notes      => 'called SubfileType by the TIFF specification',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        # set priority directory if this is the full resolution image
        RawConv    => '$self->SetPriorityDir() if $val eq "1"; $val',
        PrintConv  => {
            1 => 'Full-resolution image',
            2 => 'Reduced-resolution image',
            3 => 'Single page of multi-page image',
        },
    },
    0x100  => {
        Name       => 'ImageWidth',
        # even though Group 1 is set dynamically we need to register IFD1 once
        # so it will show up in the group lists
        Groups     => { 1 => 'IFD1' },
        Protected  => 1,
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
        # Note: priority 0 tags automatically have their priority increased for the
        # priority directory (the directory with a SubfileType of "Full-resolution image")
        Priority   => 0,
    },
    0x101  => {
        Name       => 'ImageHeight',
        Notes      => 'called ImageLength by the EXIF spec.',
        Protected  => 1,
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
        Priority   => 0,
    },
    0x102  => {
        Name       => 'BitsPerSample',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Count      => -1, # can be 1 or 3: -1 means 'variable'
        Priority   => 0,
    },
    0x103  => {
        Name          => 'Compression',
        Protected     => 1,
        Writable      => 'int16u',
        WriteGroup    => 'IFD0',
        Mandatory     => 1,
        DataMember    => 'Compression',
        SeparateTable => 'Compression',
        RawConv       => q{
            Image::ExifTool::Exif::IdentifyRawFile($self, $val);
            return $$self{Compression} = $val;
        },
        PrintConv     => \%compression,
        Priority      => 0,
    },
    0x106  => {
        Name       => 'PhotometricInterpretation',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv  => \%photometricInterpretation,
        Priority   => 0,
    },
    0x107  => {
        Name       => 'Thresholding',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv  => {
            1 => 'No dithering or halftoning',
            2 => 'Ordered dither or halftone',
            3 => 'Randomized dither',
        },
    },
    0x108  => {
        Name       => 'CellWidth',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
    },
    0x109  => {
        Name       => 'CellLength',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
    },
    0x10a  => {
        Name       => 'FillOrder',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv  => {
            1 => 'Normal',
            2 => 'Reversed',
        },
    },
    0x10d  => {
        Name       => 'DocumentName',
        Writable   => 'string',
        WriteGroup => 'IFD0',
    },
    0x10e  => {
        Name       => 'ImageDescription',
        Writable   => 'string',
        WriteGroup => 'IFD0',
        Priority   => 0,
    },
    0x10f  => {
        Name       => 'Make',
        Groups     => { 2 => 'Camera' },
        Writable   => 'string',
        WriteGroup => 'IFD0',
        DataMember => 'Make',
        # remove trailing blanks and save as an ExifTool member variable
        RawConv    => '$val =~ s/\s+$//; $$self{Make} = $val',
        # NOTE: trailing "blanks" (spaces) are removed from all EXIF tags which
        # may be "unknown" (filled with spaces) according to the EXIF spec.
        # This allows conditional replacement with "exiftool -TAG-= -TAG=VALUE".
        # - also removed are any other trailing whitespace characters
    },
    0x110  => {
        Name        => 'Model',
        Description => 'Camera Model Name',
        Groups      => { 2 => 'Camera' },
        Writable    => 'string',
        WriteGroup  => 'IFD0',
        DataMember  => 'Model',
        # remove trailing blanks and save as an ExifTool member variable
        RawConv     => '$val =~ s/\s+$//; $$self{Model} = $val',
    },
    0x111  => {
        # PreviewImageStart in IFD0 of CR2 images
        Condition  => '$$self{TIFF_TYPE} eq "CR2"',
        Name       => 'StripOffsets',
        IsOffset   => 1,
        OffsetPair => 0x117,
        Notes      => q{
                called StripOffsets in most locations, but it is PreviewImageStart in IFD0
                of CR2 images and various IFD's of DNG images except for SubIFD2 where it is
                JpgFromRawStart
            },
        DataTag    => 'PreviewImage',
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
        Protected  => 2,
        Permanent  => 1,
    },
    # 0x111  => [
    #     {
    #         Condition  => q[
    #             $$self{TIFF_TYPE} eq 'MRW' and $$self{DIR_NAME} eq 'IFD0' and
    #             $$self{Model} =~ /^DiMAGE A200/
    #         ],
    #         Name       => 'StripOffsets',
    #         IsOffset   => 1,
    #         OffsetPair => 0x117, # point to associated byte counts
    #         # A200 stores this information in the wrong byte order!!
    #         ValueConv  => '$val=join(" ",unpack("N*",pack("V*",split(" ",$val))));\$val',
    #         ByteOrder  => 'LittleEndian',
    #     },
    #     {
    #         # (APP1 IFD2 is for Leica JPEG preview)
    #         Condition  => q[
    #             not ($$self{TIFF_TYPE} eq 'CR2' and $$self{DIR_NAME} eq 'IFD0') and
    #             not ($$self{TIFF_TYPE} =~ /^(DNG|TIFF)$/ and $$self{Compression} eq '7' and $$self{SubfileType} ne '0') and
    #             not ($$self{TIFF_TYPE} eq 'APP1' and $$self{DIR_NAME} eq 'IFD2')
    #         ],
    #         Name       => 'StripOffsets',
    #         IsOffset   => 1,
    #         OffsetPair => 0x117, # point to associated byte counts
    #         ValueConv  => 'length($val) > 32 ? \$val : $val',
    #     },
    #     {
    #         # PreviewImageStart in IFD0 of CR2 images
    #         Condition  => '$$self{TIFF_TYPE} eq "CR2"',
    #         Name       => 'PreviewImageStart',
    #         IsOffset   => 1,
    #         OffsetPair => 0x117,
    #         Notes      => q{
    #             called StripOffsets in most locations, but it is PreviewImageStart in IFD0
    #             of CR2 images and various IFD's of DNG images except for SubIFD2 where it is
    #             JpgFromRawStart
    #         },
    #         DataTag    => 'PreviewImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'IFD0',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         # PreviewImageStart in various IFD's of DNG images except SubIFD2
    #         Condition  => '$$self{DIR_NAME} ne "SubIFD2"',
    #         Name       => 'PreviewImageStart',
    #         IsOffset   => 1,
    #         OffsetPair => 0x117,
    #         DataTag    => 'PreviewImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'All', # (writes to specific group of associated Composite tag)
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         # JpgFromRawStart in various IFD's of DNG images except SubIFD2
    #         Name       => 'JpgFromRawStart',
    #         IsOffset   => 1,
    #         OffsetPair => 0x117,
    #         DataTag    => 'JpgFromRaw',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD2',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    # ],
    0x112  => {
        Name       => 'Orientation',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv  => \%orientation,
        Priority   => 0, # so PRIORITY_DIR takes precedence
    },
    0x115  => {
        Name       => 'SamplesPerPixel',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Priority   => 0,
    },
    0x116  => {
        Name       => 'RowsPerStrip',
        Protected  => 1,
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
        Priority   => 0,
    },
    0x117  => {
        Name       => 'StripByteCounts',
        Notes      => 'The total number of bytes in each strip. With JPEG compressed data this designation is not needed and is omitted.',
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
        Protected  => 2,
        Permanent  => 1,
    },
    # 0x117  => [
    #     {
    #         Condition  => q[
    #             $$self{TIFF_TYPE} eq 'MRW' and $$self{DIR_NAME} eq 'IFD0' and
    #             $$self{Model} =~ /^DiMAGE A200/
    #         ],
    #         Name       => 'StripByteCounts',
    #         OffsetPair => 0x111, # point to associated offset
    #         # A200 stores this information in the wrong byte order!!
    #         ValueConv  => '$val=join(" ",unpack("N*",pack("V*",split(" ",$val))));\$val',
    #         ByteOrder  => 'LittleEndian',
    #     },
    #     {
    #         # (APP1 IFD2 is for Leica JPEG preview)
    #         Condition  => q[
    #             not ($$self{TIFF_TYPE} eq 'CR2' and $$self{DIR_NAME} eq 'IFD0') and
    #             not ($$self{TIFF_TYPE} =~ /^(DNG|TIFF)$/ and $$self{Compression} eq '7' and $$self{SubfileType} ne '0') and
    #             not ($$self{TIFF_TYPE} eq 'APP1' and $$self{DIR_NAME} eq 'IFD2')
    #         ],
    #         Name       => 'StripByteCounts',
    #         OffsetPair => 0x111, # point to associated offset
    #         ValueConv  => 'length($val) > 32 ? \$val : $val',
    #     },
    #     {
    #         # PreviewImageLength in IFD0 of CR2 images
    #         Condition  => '$$self{TIFF_TYPE} eq "CR2"',
    #         Name       => 'PreviewImageLength',
    #         OffsetPair => 0x111,
    #         Notes      => q{
    #             called StripByteCounts in most locations, but it is PreviewImageLength in
    #             IFD0 of CR2 images and various IFD's of DNG images except for SubIFD2 where
    #             it is JpgFromRawLength
    #         },
    #         DataTag    => 'PreviewImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'IFD0',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         # PreviewImageLength in various IFD's of DNG images except SubIFD2
    #         Condition  => '$$self{DIR_NAME} ne "SubIFD2"',
    #         Name       => 'PreviewImageLength',
    #         OffsetPair => 0x111,
    #         DataTag    => 'PreviewImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'All', # (writes to specific group of associated Composite tag)
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         # JpgFromRawLength in SubIFD2 of DNG images
    #         Name       => 'JpgFromRawLength',
    #         OffsetPair => 0x111,
    #         DataTag    => 'JpgFromRaw',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD2',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    # ],
    0x118  => {
        Name       => 'MinSampleValue',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
    },
    0x119  => {
        Name       => 'MaxSampleValue',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
    },
    0x11a  => {
        Name       => 'XResolution',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Mandatory  => 1,
        Priority   => 0, # so PRIORITY_DIR takes precedence
    },
    0x11b  => {
        Name       => 'YResolution',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Mandatory  => 1,
        Priority   => 0,
    },
    0x11c  => {
        Name       => 'PlanarConfiguration',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv  => {
            1 => 'Chunky',
            2 => 'Planar',
        },
        Priority   => 0,
    },
    0x11d  => {
        Name       => 'PageName',
        Writable   => 'string',
        WriteGroup => 'IFD0',
    },
    0x11e  => {
        Name       => 'XPosition',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
    },
    0x11f  => {
        Name       => 'YPosition',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
    },
    # FreeOffsets/FreeByteCounts are used by Ricoh for RMETA information
    # in TIFF images (not yet supported)
    0x120  => {
        Name       => 'FreeOffsets',
        IsOffset   => 1,
        OffsetPair => 0x121,
        ValueConv  => 'length($val) > 32 ? \$val : $val',
    },
    0x121  => {
        Name       => 'FreeByteCounts',
        OffsetPair => 0x120,
        ValueConv  => 'length($val) > 32 ? \$val : $val',
    },
    0x122  => {
        Name       => 'GrayResponseUnit',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv  => { #3
            1 => 0.1,
            2 => 0.001,
            3 => 0.0001,
            4 => 0.00001,
            5 => 0.000001,
        },
    },
    0x123  => {
        Name   => 'GrayResponseCurve',
        Binary => 1,
    },
    0x124  => {
        Name      => 'T4Options',
        PrintConv => { BITMASK => {
            0 => '2-Dimensional encoding',
            1 => 'Uncompressed',
            2 => 'Fill bits added',
        } }, #3
    },
    0x125  => {
        Name      => 'T6Options',
        PrintConv => { BITMASK => {
            1 => 'Uncompressed',
        } }, #3
    },
    0x128  => {
        Name       => 'ResolutionUnit',
        Notes      => 'the value 1 is not standard EXIF',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Mandatory  => 1,
        PrintConv  => {
            1 => 'None',
            2 => 'inches',
            3 => 'cm',
        },
        Priority   => 0,
    },
    0x129  => {
        Name       => 'PageNumber',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Count      => 2,
    },
    0x12c  => 'ColorResponseUnit', #9
    0x12d  => {
        Name       => 'TransferFunction',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Count      => 768,
        Binary     => 1,
    },
    0x131  => {
        Name       => 'Software',
        Writable   => 'string',
        WriteGroup => 'IFD0',
        DataMember => 'Software',
        RawConv    => '$val =~ s/\s+$//; $$self{Software} = $val', # trim trailing blanks
    },
    0x132  => {
        Name         => 'ModifyDate',
        Groups       => { 2 => 'Time' },
        Notes        => 'called DateTime by the EXIF spec.',
        Writable     => 'string',
        Shift        => 'Time',
        WriteGroup   => 'IFD0',
        Validate     => 'ValidateExifDate($val)',
        PrintConv    => '$self->ConvertDateTime($val)',
        PrintConvInv => '$self->InverseDateTime($val,0)',
    },
    0x13b  => {
        Name       => 'Artist',
        Groups     => { 2 => 'Author' },
        Notes      => 'becomes a list-type tag when the MWG module is loaded',
        Writable   => 'string',
        WriteGroup => 'IFD0',
        RawConv    => '$val =~ s/\s+$//; $val', # trim trailing blanks
    },
    0x13c  => {
        Name       => 'HostComputer',
        Writable   => 'string',
        WriteGroup => 'IFD0',
    },
    0x13d  => {
        Name       => 'Predictor',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv  => {
            1     => 'None',
            2     => 'Horizontal differencing',
            3     => 'Floating point',           # (DNG 1.5)
            34892 => 'Horizontal difference X2', # (DNG 1.5)
            34893 => 'Horizontal difference X4', # (DNG 1.5)
            34894 => 'Floating point X2',        # (DNG 1.5)
            34895 => 'Floating point X4',        # (DNG 1.5)
        },
    },
    0x13e  => {
        Name       => 'WhitePoint',
        Groups     => { 2 => 'Camera' },
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Count      => 2,
    },
    0x13f  => {
        Name       => 'PrimaryChromaticities',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Count      => 6,
        Priority   => 0,
    },
    0x140  => {
        Name   => 'ColorMap',
        Format => 'binary',
        Binary => 1,
    },
    0x141  => {
        Name       => 'HalftoneHints',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Count      => 2,
    },
    0x142  => {
        Name       => 'TileWidth',
        Protected  => 1,
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
    },
    0x143  => {
        Name       => 'TileLength',
        Protected  => 1,
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
    },
    0x144  => {
        Name       => 'TileOffsets',
        IsOffset   => 1,
        OffsetPair => 0x145,
        ValueConv  => 'length($val) > 32 ? \$val : $val',
    },
    0x145  => {
        Name       => 'TileByteCounts',
        OffsetPair => 0x144,
        ValueConv  => 'length($val) > 32 ? \$val : $val',
    },
    0x146  => 'BadFaxLines', #3
    0x147  => {              #3
        Name      => 'CleanFaxData',
        PrintConv => {
            0 => 'Clean',
            1 => 'Regenerated',
            2 => 'Unclean',
        },
    },
    0x148  => 'ConsecutiveBadFaxLines', #3
    0x14a  => {
        Name       => 'SubIFDs',
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
        Notes      => 'Defined by Adobe Corporation to enable TIFF Trees within a TIFF file.',
        IsOffset   => 1,

    },
    # 0x14a  => [
    #     {
    #         Name         => 'SubIFD',
    #         # use this opportunity to identify an ARW image, and if so we
    #         # must decide if this is a SubIFD or the A100 raw data
    #         # (use SubfileType, Compression and FILE_TYPE to identify ARW/SR2,
    #         # then call SetARW to finish the job)
    #         Condition    => q{
    #             $$self{DIR_NAME} ne 'IFD0' or $$self{FILE_TYPE} ne 'TIFF' or
    #             $$self{Make} !~ /^SONY/ or
    #             not $$self{SubfileType} or $$self{SubfileType} != 1 or
    #             not $$self{Compression} or $$self{Compression} != 6 or
    #             not require Image::ExifTool::Sony or
    #             Image::ExifTool::Sony::SetARW($self, $valPt)
    #         },
    #         Groups       => { 1 => 'SubIFD' },
    #         Flags        => 'SubIFD',
    #         SubDirectory => {
    #             Start      => '$val',
    #             MaxSubdirs => 10, # (have seen 5 in a DNG 1.4 image)
    #         },
    #     },
    #     { #16
    #         Name       => 'A100DataOffset',
    #         Notes      => 'the data offset in original Sony DSLR-A100 ARW images',
    #         DataMember => 'A100DataOffset',
    #         RawConv    => '$$self{A100DataOffset} = $val',
    #         WriteGroup => 'IFD0', # (only for Validate)
    #         IsOffset   => 1,
    #         Protected  => 2,
    #     },
    # ],
    0x14c  => {
        Name       => 'InkSet',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv  => { #3
            1 => 'CMYK',
            2 => 'Not CMYK',
        },
    },
    0x14d  => 'InkNames',     #3
    0x14e  => 'NumberofInks', #3
    0x150  => 'DotRange',
    0x151  => {
        Name       => 'TargetPrinter',
        Writable   => 'string',
        WriteGroup => 'IFD0',
    },
    0x152  => {
        Name      => 'ExtraSamples',
        PrintConv => { #20
            0 => 'Unspecified',
            1 => 'Associated Alpha',
            2 => 'Unassociated Alpha',
        },
    },
    0x153  => {
        Name             => 'SampleFormat',
        Notes            => 'SamplesPerPixel values',
        WriteGroup       => 'SubIFD', # (only for Validate)
        PrintConvColumns => 2,
        PrintConv        => [ \%sampleFormat, \%sampleFormat, \%sampleFormat, \%sampleFormat ],
    },
    0x154  => 'SMinSampleValue',
    0x155  => 'SMaxSampleValue',
    0x156  => 'TransferRange',
    0x157  => 'ClipPath',       #3
    0x158  => 'XClipPathUnits', #3
    0x159  => 'YClipPathUnits', #3
    0x15a  => {                 #3
        Name      => 'Indexed',
        PrintConv => { 0 => 'Not indexed', 1 => 'Indexed' },
    },
    0x15b  => {
        Name   => 'JPEGTables',
        Binary => 1,
    },
    0x15f  => { #10
        Name      => 'OPIProxy',
        PrintConv => {
            0 => 'Higher resolution image does not exist',
            1 => 'Higher resolution image exists',
        },
    },
    # 0x181 => 'Decode', #20 (typo! - should be 0x1b1, ref 21)
    # 0x182 => 'DefaultImageColor', #20 (typo! - should be 0x1b2, ref 21)
    0x190  => { #3
        Name         => 'GlobalParametersIFD',
        Groups       => { 1 => 'GlobParamIFD' },
        Flags        => 'SubIFD',
        SubDirectory => {
            DirName    => 'GlobParamIFD',
            Start      => '$val',
            MaxSubdirs => 1,
        },
    },
    0x191  => { #3
        Name      => 'ProfileType',
        PrintConv => { 0 => 'Unspecified', 1 => 'Group 3 FAX' },
    },
    0x192  => { #3
        Name      => 'FaxProfile',
        PrintConv => {
            0   => 'Unknown',
            1   => 'Minimal B&W lossless, S',
            2   => 'Extended B&W lossless, F',
            3   => 'Lossless JBIG B&W, J',
            4   => 'Lossy color and grayscale, C',
            5   => 'Lossless color and grayscale, L',
            6   => 'Mixed raster content, M',
            7   => 'Profile T',      #20
            255 => 'Multi Profiles', #20
        },
    },
    0x193  => { #3
        Name      => 'CodingMethods',
        PrintConv => { BITMASK => {
            0 => 'Unspecified compression',
            1 => 'Modified Huffman',
            2 => 'Modified Read',
            3 => 'Modified MR',
            4 => 'JBIG',
            5 => 'Baseline JPEG',
            6 => 'JBIG color',
        } },
    },
    0x194  => 'VersionYear',       #3
    0x195  => 'ModeNumber',        #3
    0x1b1  => 'Decode',            #3
    0x1b2  => 'DefaultImageColor', #3 (changed to ImageBaseColor, ref 21)
    0x1b3  => 'T82Options',        #20
    0x1b5  => {                    #19
        Name   => 'JPEGTables',
        Binary => 1,
    },
    0x200  => {
        Name      => 'JPEGProc',
        PrintConv => {
            1  => 'Baseline',
            14 => 'Lossless',
        },
    },
    0x201  => {
        Name       => 'ThumbnailOffset',
        Notes      => q{
                ThumbnailOffset in IFD1 of JPEG and some TIFF-based images, IFD0 of MRW
                images and AVI and MOV videos, and the SubIFD in IFD1 of SRW images;
                PreviewImageStart in MakerNotes and IFD0 of ARW and SR2 images;
                JpgFromRawStart in SubIFD of NEF images and IFD2 of PEF images; and
                OtherImageStart in everything else
            },
        IsOffset   => 1,
        Writable   => 'int32u',
        WriteGroup => 'IFD0',
        Protected  => 2,
        Permanent  => 1,
    },
    # 0x201  => [
    #     {
    #         Name           => 'ThumbnailOffset',
    #         Notes          => q{
    #             ThumbnailOffset in IFD1 of JPEG and some TIFF-based images, IFD0 of MRW
    #             images and AVI and MOV videos, and the SubIFD in IFD1 of SRW images;
    #             PreviewImageStart in MakerNotes and IFD0 of ARW and SR2 images;
    #             JpgFromRawStart in SubIFD of NEF images and IFD2 of PEF images; and
    #             OtherImageStart in everything else
    #         },
    #         # thumbnail is found in IFD1 of JPEG and TIFF images, and
    #         # IFD0 of EXIF information in FujiFilm AVI (RIFF) and MOV videos
    #         Condition      => q{
    #             # recognize NRW file from a JPEG-compressed thumbnail in IFD0
    #             if ($$self{TIFF_TYPE} eq 'NEF' and $$self{DIR_NAME} eq 'IFD0' and $$self{Compression} == 6) {
    #                 $self->OverrideFileType($$self{TIFF_TYPE} = 'NRW');
    #             }
    #             $$self{DIR_NAME} eq 'IFD1' or
    #             ($$self{DIR_NAME} eq 'IFD0' and $$self{FILE_TYPE} =~ /^(RIFF|MOV)$/)
    #         },
    #         IsOffset       => 1,
    #         OffsetPair     => 0x202,
    #         DataTag        => 'ThumbnailImage',
    #         Writable       => 'int32u',
    #         WriteGroup     => 'IFD1',
    #         # according to the EXIF spec. a JPEG-compressed thumbnail image may not
    #         # be stored in a TIFF file, but these TIFF-based RAW image formats
    #         # use IFD1 for a JPEG-compressed thumbnail:  CR2, ARW, SR2 and PEF.
    #         # (SRF also stores a JPEG image in IFD1, but it is actually a preview
    #         # and we don't yet write SRF anyway)
    #         WriteCondition => q{
    #             $$self{FILE_TYPE} ne "TIFF" or
    #             $$self{TIFF_TYPE} =~ /^(CR2|ARW|SR2|PEF)$/
    #         },
    #         Protected      => 2,
    #     },
    #     {
    #         Name       => 'ThumbnailOffset',
    #         # thumbnail in IFD0 of MRW images (Minolta A200)
    #         # and IFD0 of NRW images (Nikon Coolpix P6000,P7000,P7100)
    #         Condition  => '$$self{DIR_NAME} eq "IFD0" and $$self{TIFF_TYPE} =~ /^(MRW|NRW)$/',
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #         # A200 uses the wrong base offset for this pointer!!
    #         WrongBase  => '$$self{Model} =~ /^DiMAGE A200/ ? $$self{MRW_WrongBase} : undef',
    #         DataTag    => 'ThumbnailImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'IFD0',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'ThumbnailOffset',
    #         # in SubIFD of IFD1 in Samsung SRW images
    #         Condition  => q{
    #             $$self{TIFF_TYPE} eq 'SRW' and $$self{DIR_NAME} eq 'SubIFD' and
    #             $$self{PATH}[-2] eq 'IFD1'
    #         },
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #         DataTag    => 'ThumbnailImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'PreviewImageStart',
    #         Condition  => '$$self{DIR_NAME} eq "MakerNotes"',
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #         DataTag    => 'PreviewImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'MakerNotes',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'PreviewImageStart',
    #         # PreviewImage in IFD0 of ARW and SR2 files for all models
    #         Condition  => '$$self{DIR_NAME} eq "IFD0" and $$self{TIFF_TYPE} =~ /^(ARW|SR2)$/',
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #         DataTag    => 'PreviewImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'IFD0',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'JpgFromRawStart',
    #         Condition  => '$$self{DIR_NAME} eq "SubIFD"',
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #         DataTag    => 'JpgFromRaw',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD',
    #         # JpgFromRaw is in SubIFD of NEF, NRW and SRW files
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'JpgFromRawStart',
    #         Condition  => '$$self{DIR_NAME} eq "IFD2"',
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #         DataTag    => 'JpgFromRaw',
    #         Writable   => 'int32u',
    #         WriteGroup => 'IFD2',
    #         # JpgFromRaw is in IFD2 of PEF files
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'OtherImageStart',
    #         Condition  => '$$self{DIR_NAME} eq "SubIFD1"',
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #         DataTag    => 'OtherImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD1',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'OtherImageStart',
    #         Condition  => '$$self{DIR_NAME} eq "SubIFD2"',
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #         DataTag    => 'OtherImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD2',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'OtherImageStart',
    #         IsOffset   => 1,
    #         OffsetPair => 0x202,
    #     },
    # ],
    0x202  =>         {
        Name           => 'ThumbnailLength',
        Notes          => q{
                ThumbnailLength in IFD1 of JPEG and some TIFF-based images, IFD0 of MRW
                images and AVI and MOV videos, and the SubIFD in IFD1 of SRW images;
                PreviewImageLength in MakerNotes and IFD0 of ARW and SR2 images;
                JpgFromRawLength in SubIFD of NEF images, and IFD2 of PEF images; and
                OtherImageLength in everything else
            },
        Condition      => q{
                $$self{DIR_NAME} eq 'IFD1' or
                ($$self{DIR_NAME} eq 'IFD0' and $$self{FILE_TYPE} =~ /^(RIFF|MOV)$/)
            },
        OffsetPair     => 0x201,
        DataTag        => 'ThumbnailImage',
        Writable       => 'int32u',
        WriteGroup     => 'IFD1',
        WriteCondition => q{
                $$self{FILE_TYPE} ne "TIFF" or
                $$self{TIFF_TYPE} =~ /^(CR2|ARW|SR2|PEF)$/
            },
        Protected      => 2,
    },
    # 0x202  => [
    #     {
    #         Name           => 'ThumbnailLength',
    #         Notes          => q{
    #             ThumbnailLength in IFD1 of JPEG and some TIFF-based images, IFD0 of MRW
    #             images and AVI and MOV videos, and the SubIFD in IFD1 of SRW images;
    #             PreviewImageLength in MakerNotes and IFD0 of ARW and SR2 images;
    #             JpgFromRawLength in SubIFD of NEF images, and IFD2 of PEF images; and
    #             OtherImageLength in everything else
    #         },
    #         Condition      => q{
    #             $$self{DIR_NAME} eq 'IFD1' or
    #             ($$self{DIR_NAME} eq 'IFD0' and $$self{FILE_TYPE} =~ /^(RIFF|MOV)$/)
    #         },
    #         OffsetPair     => 0x201,
    #         DataTag        => 'ThumbnailImage',
    #         Writable       => 'int32u',
    #         WriteGroup     => 'IFD1',
    #         WriteCondition => q{
    #             $$self{FILE_TYPE} ne "TIFF" or
    #             $$self{TIFF_TYPE} =~ /^(CR2|ARW|SR2|PEF)$/
    #         },
    #         Protected      => 2,
    #     },
    #     {
    #         Name       => 'ThumbnailLength',
    #         # thumbnail in IFD0 of MRW images (Minolta A200)
    #         # and IFD0 of NRW images (Nikon Coolpix P6000,P7000,P7100)
    #         Condition  => '$$self{DIR_NAME} eq "IFD0" and $$self{TIFF_TYPE} =~ /^(MRW|NRW)$/',
    #         OffsetPair => 0x201,
    #         DataTag    => 'ThumbnailImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'IFD0',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'ThumbnailLength',
    #         # in SubIFD of IFD1 in Samsung SRW images
    #         Condition  => q{
    #             $$self{TIFF_TYPE} eq 'SRW' and $$self{DIR_NAME} eq 'SubIFD' and
    #             $$self{PATH}[-2] eq 'IFD1'
    #         },
    #         OffsetPair => 0x201,
    #         DataTag    => 'ThumbnailImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'PreviewImageLength',
    #         Condition  => '$$self{DIR_NAME} eq "MakerNotes"',
    #         OffsetPair => 0x201,
    #         DataTag    => 'PreviewImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'MakerNotes',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'PreviewImageLength',
    #         # PreviewImage in IFD0 of ARW and SR2 files for all models
    #         Condition  => '$$self{DIR_NAME} eq "IFD0" and $$self{TIFF_TYPE} =~ /^(ARW|SR2)$/',
    #         OffsetPair => 0x201,
    #         DataTag    => 'PreviewImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'IFD0',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'JpgFromRawLength',
    #         Condition  => '$$self{DIR_NAME} eq "SubIFD"',
    #         OffsetPair => 0x201,
    #         DataTag    => 'JpgFromRaw',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'JpgFromRawLength',
    #         Condition  => '$$self{DIR_NAME} eq "IFD2"',
    #         OffsetPair => 0x201,
    #         DataTag    => 'JpgFromRaw',
    #         Writable   => 'int32u',
    #         WriteGroup => 'IFD2',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'OtherImageLength',
    #         Condition  => '$$self{DIR_NAME} eq "SubIFD1"',
    #         OffsetPair => 0x201,
    #         DataTag    => 'OtherImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD1',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'OtherImageLength',
    #         Condition  => '$$self{DIR_NAME} eq "SubIFD2"',
    #         OffsetPair => 0x201,
    #         DataTag    => 'OtherImage',
    #         Writable   => 'int32u',
    #         WriteGroup => 'SubIFD2',
    #         Protected  => 2,
    #         Permanent  => 1,
    #     },
    #     {
    #         Name       => 'OtherImageLength',
    #         OffsetPair => 0x201,
    #     },
    # ],
    0x203  => 'JPEGRestartInterval',
    0x205  => 'JPEGLosslessPredictors',
    0x206  => 'JPEGPointTransforms',
    0x207  => {
        Name       => 'JPEGQTables',
        IsOffset   => 1,
        # this tag is not supported for writing, so define an
        # invalid offset pair to cause a "No size tag" error to be
        # generated if we try to write a file containing this tag
        OffsetPair => -1,
    },
    0x208  => {
        Name       => 'JPEGDCTables',
        IsOffset   => 1,
        OffsetPair => -1, # (see comment for JPEGQTables)
    },
    0x209  => {
        Name       => 'JPEGACTables',
        IsOffset   => 1,
        OffsetPair => -1, # (see comment for JPEGQTables)
    },
    0x211  => {
        Name       => 'YCbCrCoefficients',
        Protected  => 1,
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Count      => 3,
        Priority   => 0,
    },
    0x212  => {
        Name             => 'YCbCrSubSampling',
        Protected        => 1,
        Writable         => 'int16u',
        WriteGroup       => 'IFD0',
        Count            => 2,
        PrintConvColumns => 2,
        PrintConv        => \%Image::ExifTool::JPEG::yCbCrSubSampling,
        Priority         => 0,
    },
    0x213  => {
        Name       => 'YCbCrPositioning',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Mandatory  => 1,
        PrintConv  => {
            1 => 'Centered',
            2 => 'Co-sited',
        },
        Priority   => 0,
    },
    0x214  => {
        Name       => 'ReferenceBlackWhite',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Count      => 6,
        Priority   => 0,
    },
    # 0x220 - int32u: 0 (IFD0, Xaiomi Redmi models)
    # 0x221 - int32u: 0 (IFD0, Xaiomi Redmi models)
    # 0x222 - int32u: 0 (IFD0, Xaiomi Redmi models)
    # 0x223 - int32u: 0 (IFD0, Xaiomi Redmi models)
    # 0x224 - int32u: 0,1 (IFD0, Xaiomi Redmi models)
    # 0x225 - string: "" (IFD0, Xaiomi Redmi models)
    0x22f  => 'StripRowCounts',
    0x2bc  => {
        Name         => 'ApplicationNotes', # (writable directory!)
        Format       => 'undef',
        Writable     => 'int8u',
        WriteGroup   => 'IFD0', # (only for Validate)
        Flags        => [ 'Binary', 'Protected' ],
        # this could be an XMP block
        SubDirectory => {
            DirName  => 'XMP',
            TagTable => 'Image::ExifTool::XMP::Main',
        },
    },
    0x3e7  => 'USPTOMiscellaneous', #20
    0x1000 => {                     #5
        Name       => 'RelatedImageFileFormat',
        Protected  => 1,
        Writable   => 'string',
        WriteGroup => 'InteropIFD',
    },
    0x1001 => { #5
        Name       => 'RelatedImageWidth',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'InteropIFD',
    },
    0x1002 => { #5
        Name       => 'RelatedImageHeight',
        Notes      => 'called RelatedImageLength by the DCF spec.',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'InteropIFD',
    },
    # (0x474x tags written by MicrosoftPhoto)
    0x4746 => { #PH
        Name       => 'Rating',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Avoid      => 1,
    },
    0x4747 => { # (written by Digital Image Pro)
        Name      => 'XP_DIP_XML',
        Format    => 'undef',
        # the following reference indicates this is Unicode:
        # http://social.msdn.microsoft.com/Forums/en-US/isvvba/thread/ce6edcbb-8fc2-40c6-ad98-85f5d835ddfb
        ValueConv => '$self->Decode($val,"UCS2","II")',
    },
    0x4748 => {
        Name         => 'StitchInfo',
        SubDirectory => {
            TagTable  => 'Image::ExifTool::Microsoft::Stitch',
            ByteOrder => 'LittleEndian', #PH (NC)
        },
    },
    0x4749 => { #PH
        Name       => 'RatingPercent',
        Writable   => 'int16u',
        WriteGroup => 'IFD0',
        Avoid      => 1,
    },
    0x7000 => { #JR
        Name      => 'SonyRawFileType',
        # (only valid if Sony:FileFormat >= ARW 2.0, ref IB)
        # Writable => 'int16u', (don't allow writes for now)
        PrintConv => {
            0 => 'Sony Uncompressed 14-bit RAW',
            1 => 'Sony Uncompressed 12-bit RAW',   #IB
            2 => 'Sony Compressed RAW',            # (lossy, ref IB)
            3 => 'Sony Lossless Compressed RAW',   #IB
            4 => 'Sony Lossless Compressed RAW 2', #JR (ILCE-1)
        },
    },
    # 0x7001 - int16u[1] (in SubIFD of Sony ARW images) - values: 0,1
    0x7010 => { #IB
        Name => 'SonyToneCurve',
        # int16u[4] (in SubIFD of Sony ARW images -- don't allow writes for now)
        # - only the middle 4 points are stored (lower comes from black level,
        #   and upper from data maximum)
    },
    # 0x7011 - int16u[4] (in SubIFD of Sony ARW images) - values: "0 4912 8212 12287","4000 7200 10050 12075"
    # 0x7020 - int32u[1] (in SubIFD of Sony ARW images) - values: 0,3
    0x7031 => {
        Name       => 'VignettingCorrection',
        Notes      => 'found in Sony ARW images',
        Protected  => 1,
        Writable   => 'int16s',
        WriteGroup => 'SubIFD',
        PrintConv  => {
            256 => 'Off',
            257 => 'Auto',
            272 => 'Auto (ILCE-1)', #JR
            511 => 'No correction params available',
        },
    },
    0x7032 => {
        Name       => 'VignettingCorrParams', #forum7640
        Notes      => 'found in Sony ARW images',
        Protected  => 1,
        Writable   => 'int16s',
        WriteGroup => 'SubIFD',
        Count      => 17,
    },
    0x7034 => {
        Name       => 'ChromaticAberrationCorrection',
        Notes      => 'found in Sony ARW images',
        Protected  => 1,
        Writable   => 'int16s',
        WriteGroup => 'SubIFD',
        PrintConv  => {
            0   => 'Off',
            1   => 'Auto',
            255 => 'No correction params available',
        },
    },
    0x7035 => {
        Name       => 'ChromaticAberrationCorrParams', #forum6509
        Notes      => 'found in Sony ARW images',
        Protected  => 1,
        Writable   => 'int16s',
        WriteGroup => 'SubIFD',
        Count      => 33,
    },
    0x7036 => {
        Name       => 'DistortionCorrection',
        Notes      => 'found in Sony ARW images',
        Protected  => 1,
        Writable   => 'int16s',
        WriteGroup => 'SubIFD',
        PrintConv  => {
            0   => 'Off',
            1   => 'Auto',
            17  => 'Auto fixed by lens',
            255 => 'No correction params available',
        },
    },
    0x7037 => {
        Name       => 'DistortionCorrParams', #forum6509
        Notes      => 'found in Sony ARW images',
        Protected  => 1,
        Writable   => 'int16s',
        WriteGroup => 'SubIFD',
        Count      => 17,
    },
    0x74c7 => { #IB (in ARW images from some Sony cameras)
        Name       => 'SonyCropTopLeft',
        Writable   => 'int32u',
        WriteGroup => 'SubIFD',
        Count      => 2,
        Permanent  => 1,
        Protected  => 1,
    },
    0x74c8 => { #IB (in ARW images from some Sony cameras)
        Name       => 'SonyCropSize',
        Writable   => 'int32u',
        WriteGroup => 'SubIFD',
        Count      => 2,
        Permanent  => 1,
        Protected  => 1,
    },
    0x800d => 'ImageID',                           #10
    0x80a3 => { Name => 'WangTag1', Binary => 1 }, #20
    0x80a4 => { Name => 'WangAnnotation', Binary => 1 },
    0x80a5 => { Name => 'WangTag3', Binary => 1 }, #20
    0x80a6 => {                                    #20
        Name      => 'WangTag4',
        PrintConv => 'length($val) <= 64 ? $val : \$val',
    },
    # tags 0x80b8-0x80bc are registered to Island Graphics
    0x80b9 => 'ImageReferencePoints', #29
    0x80ba => 'RegionXformTackPoint', #29
    0x80bb => 'WarpQuadrilateral',    #29
    0x80bc => 'AffineTransformMat',   #29
    0x80e3 => 'Matteing',             #9
    0x80e4 => 'DataType',             #9
    0x80e5 => 'ImageDepth',           #9
    0x80e6 => 'TileDepth',            #9
    # tags 0x8214-0x8219 are registered to Pixar
    0x8214 => 'ImageFullWidth',      #29
    0x8215 => 'ImageFullHeight',     #29
    0x8216 => 'TextureFormat',       #29
    0x8217 => 'WrapModes',           #29
    0x8218 => 'FovCot',              #29
    0x8219 => 'MatrixWorldToScreen', #29
    0x821a => 'MatrixWorldToCamera', #29
    0x827d => 'Model2',              #29 (Eastman Kodak)
    0x828d => {                      #12
        Name       => 'CFARepeatPatternDim',
        Protected  => 1,
        Writable   => 'int16u',
        WriteGroup => 'SubIFD',
        Count      => 2,
    },
    0x828e => {
        Name       => 'CFAPattern2', #12
        Format     => 'int8u',       # (written incorrectly as 'undef' in Nikon NRW images)
        Protected  => 1,
        Writable   => 'int8u',
        WriteGroup => 'SubIFD',
        Count      => -1,
    },
    0x828f => { #12
        Name   => 'BatteryLevel',
        Groups => { 2 => 'Camera' },
    },
    0x8290 => {
        Name         => 'KodakIFD',
        Groups       => { 1 => 'KodakIFD' },
        Flags        => 'SubIFD',
        Notes        => 'used in various types of Kodak images',
        SubDirectory => {
            TagTable   => 'Image::ExifTool::Kodak::IFD',
            DirName    => 'KodakIFD',
            Start      => '$val',
            MaxSubdirs => 1,
        },
    },
    0x8298 => {
        Name         => 'Copyright',
        Groups       => { 2 => 'Author' },
        Format       => 'undef',
        Writable     => 'string',
        WriteGroup   => 'IFD0',
        Notes        => q{
            may contain copyright notices for photographer and editor, separated by a
            newline.  As per the EXIF specification, the newline is replaced by a null
            byte when writing to file, but this may be avoided by disabling the print
            conversion
        },
        # internally the strings are separated by a null character in this format:
        # Photographer only: photographer + NULL
        # Both:              photographer + NULL + editor + NULL
        # Editor only:       SPACE + NULL + editor + NULL
        # (this is done as a RawConv so conditional replaces will work properly)
        RawConv      => sub {
            my ($val, $self) = @_;
            $val =~ s/ *\0/\n/;  # translate first NULL to a newline, removing trailing blanks
            $val =~ s/ *\0.*//s; # truncate at second NULL and remove trailing blanks
            $val =~ s/\n$//;     # remove trailing newline if it exists
            # decode if necessary (note: this is the only non-'string' EXIF value like this)
            my $enc = $self->Options('CharsetEXIF');
            $val = $self->Decode($val, $enc) if $enc;
            return $val;
        },
        RawConvInv   => '$val . "\0"',
        PrintConvInv => sub {
            my ($val, $self) = @_;
            # encode if necessary (not automatic because Format is 'undef')
            my $enc = $self->Options('CharsetEXIF');
            $val = $self->Encode($val, $enc) if $enc and $val !~ /\0/;
            if ($val =~ /(.*?)\s*[\n\r]+\s*(.*)/s) {
                return $1 unless length $2;
                # photographer copyright set to ' ' if it doesn't exist, according to spec.
                return ((length($1) ? $1 : ' ') . "\0" . $2);
            }
            return $val;
        },
    },
    0x829a => {
        Name         => 'ExposureTime',
        Writable     => 'rational64u',
        PrintConv    => 'Image::ExifTool::Exif::PrintExposureTime($val)',
        PrintConvInv => '$val',
    },
    0x829d => {
        Name         => 'FNumber',
        Writable     => 'rational64u',
        PrintConv    => 'Image::ExifTool::Exif::PrintFNumber($val)',
        PrintConvInv => '$val',
    },
    0x82a5 => { #3
        Name  => 'MDFileTag',
        Notes => 'tags 0x82a5-0x82ac are used in Molecular Dynamics GEL files',
    },
    0x82a6 => 'MDScalePixel', #3
    0x82a7 => 'MDColorTable', #3
    0x82a8 => 'MDLabName',    #3
    0x82a9 => 'MDSampleInfo', #3
    0x82aa => 'MDPrepDate',   #3
    0x82ab => 'MDPrepTime',   #3
    0x82ac => 'MDFileUnits',  #3
    0x830e => {               #30 (GeoTiff)
        Name       => 'PixelScale',
        Writable   => 'double',
        WriteGroup => 'IFD0',
        Count      => 3,
    },
    0x8335 => 'AdventScale',        #20
    0x8336 => 'AdventRevision',     #20
    0x835c => 'UIC1Tag',            #23
    0x835d => 'UIC2Tag',            #23
    0x835e => 'UIC3Tag',            #23
    0x835f => 'UIC4Tag',            #23
    0x83bb => {                     #12
        Name         => 'IPTC-NAA', # (writable directory! -- but see note below)
        # this should actually be written as 'undef' (see
        # http://www.awaresystems.be/imaging/tiff/tifftags/iptc.html),
        # but Photoshop writes it as int32u and Nikon Capture won't read
        # anything else, so we do the same thing here...  Doh!
        Format       => 'undef',  # convert binary values as undef
        Writable     => 'int32u', # but write int32u format code in IFD
        WriteGroup   => 'IFD0',
        Flags        => [ 'Binary', 'Protected' ],
        SubDirectory => {
            DirName  => 'IPTC',
            TagTable => 'Image::ExifTool::IPTC::Main',
        },
        # Note: This directory may be written as a block via the IPTC-NAA tag,
        # but this technique is not recommended.  Instead, it is better to
        # write the Extra IPTC tag and let ExifTool decide where it should go.
    },
    0x847e => 'IntergraphPacketData',    #3
    0x847f => 'IntergraphFlagRegisters', #3
    0x8480 => {                          #30 (GeoTiff, obsolete)
        Name       => 'IntergraphMatrix',
        Writable   => 'double',
        WriteGroup => 'IFD0',
        Count      => -1,
    },
    0x8481 => 'INGRReserved', #20
    0x8482 => {               #30 (GeoTiff)
        Name       => 'ModelTiePoint',
        Groups     => { 2 => 'Location' },
        Writable   => 'double',
        WriteGroup => 'IFD0',
        Count      => -1,
    },
    0x84e0 => 'Site',          #9
    0x84e1 => 'ColorSequence', #9
    0x84e2 => 'IT8Header',     #9
    0x84e3 => {                #9
        Name      => 'RasterPadding',
        PrintConv => { #20
            0  => 'Byte',
            1  => 'Word',
            2  => 'Long Word',
            9  => 'Sector',
            10 => 'Long Sector',
        },
    },
    0x84e4 => 'BitsPerRunLength',         #9
    0x84e5 => 'BitsPerExtendedRunLength', #9
    0x84e6 => 'ColorTable',               #9
    0x84e7 => {                           #9
        Name      => 'ImageColorIndicator',
        PrintConv => { #20
            0 => 'Unspecified Image Color',
            1 => 'Specified Image Color',
        },
    },
    0x84e8 => { #9
        Name      => 'BackgroundColorIndicator',
        PrintConv => { #20
            0 => 'Unspecified Background Color',
            1 => 'Specified Background Color',
        },
    },
    0x84e9 => 'ImageColorValue',       #9
    0x84ea => 'BackgroundColorValue',  #9
    0x84eb => 'PixelIntensityRange',   #9
    0x84ec => 'TransparencyIndicator', #9
    0x84ed => 'ColorCharacterization', #9
    0x84ee => {                        #9
        Name      => 'HCUsage',
        PrintConv => { #20
            0 => 'CT',
            1 => 'Line Art',
            2 => 'Trap',
        },
    },
    0x84ef => 'TrapIndicator',  #17
    0x84f0 => 'CMYKEquivalent', #17
    0x8546 => {                 #11
        Name       => 'SEMInfo',
        Notes      => 'found in some scanning electron microscope images',
        Writable   => 'string',
        WriteGroup => 'IFD0',
    },
    0x8568 => {
        Name         => 'AFCP_IPTC',
        SubDirectory => {
            # must change directory name so we don't create this directory
            DirName  => 'AFCP_IPTC',
            TagTable => 'Image::ExifTool::IPTC::Main',
        },
    },
    0x85b8 => 'PixelMagicJBIGOptions', #20
    0x85d7 => 'JPLCartoIFD',           #exifprobe (NC)
    0x85d8 => {                        #30 (GeoTiff)
        Name       => 'ModelTransform',
        Groups     => { 2 => 'Location' },
        Writable   => 'double',
        WriteGroup => 'IFD0',
        Count      => 16,
    },
    0x8602 => { #16
        Name  => 'WB_GRGBLevels',
        Notes => 'found in IFD0 of Leaf MOS images',
    },
    # 0x8603 - Leaf CatchLight color matrix (ref 16)
    0x8606 => {
        Name         => 'LeafData',
        Format       => 'undef', # avoid converting huge block to string of int8u's!
        SubDirectory => {
            DirName  => 'LeafIFD',
            TagTable => 'Image::ExifTool::Leaf::Main',
        },
    },
    0x8649 => { #19
        Name         => 'PhotoshopSettings',
        Format       => 'binary',
        WriteGroup   => 'IFD0', # (only for Validate)
        SubDirectory => {
            DirName  => 'Photoshop',
            TagTable => 'Image::ExifTool::Photoshop::Main',
        },
    },
    0x8769 => {
        Name         => 'ExifOffset',
        Groups       => { 1 => 'ExifIFD' },
        WriteGroup   => 'IFD0', # (only for Validate)
        SubIFD       => 2,
        SubDirectory => {
            DirName => 'ExifIFD',
            Start   => '$val',
        },
    },
    0x8773 => {
        Name         => 'ICC_Profile',
        WriteGroup   => 'IFD0', # (only for Validate)
        SubDirectory => {
            TagTable => 'Image::ExifTool::ICC_Profile::Main',
        },
    },
    0x877f => { #20
        Name      => 'TIFF_FXExtensions',
        PrintConv => { BITMASK => {
            0 => 'Resolution/Image Width',
            1 => 'N Layer Profile M',
            2 => 'Shared Data',
            3 => 'B&W JBIG2',
            4 => 'JBIG2 Profile M',
        } },
    },
    0x8780 => { #20
        Name      => 'MultiProfiles',
        PrintConv => { BITMASK => {
            0  => 'Profile S',
            1  => 'Profile F',
            2  => 'Profile J',
            3  => 'Profile C',
            4  => 'Profile L',
            5  => 'Profile M',
            6  => 'Profile T',
            7  => 'Resolution/Image Width',
            8  => 'N Layer Profile M',
            9  => 'Shared Data',
            10 => 'JBIG2 Profile M',
        } },
    },
    0x8781 => { #22
        Name       => 'SharedData',
        IsOffset   => 1,
        # this tag is not supported for writing, so define an
        # invalid offset pair to cause a "No size tag" error to be
        # generated if we try to write a file containing this tag
        OffsetPair => -1,
    },
    0x8782 => 'T88Options', #20
    0x87ac => 'ImageLayer',
    0x87af => { #30
        Name       => 'GeoTiffDirectory',
        Format     => 'undef',
        Writable   => 'int16u',
        Notes      => q{
            these "GeoTiff" tags may read and written as a block, but they aren't
            extracted unless specifically requested.  Byte order changes are handled
            automatically when copying between TIFF images with different byte order
        },
        WriteGroup => 'IFD0',
        Binary     => 1,
        RawConv    => '$val . GetByteOrder()', # save byte order
        # swap byte order if necessary
        RawConvInv => q{
            return $val if length $val < 2;
            my $order = substr($val, -2);
            return $val unless $order eq 'II' or $order eq 'MM';
            $val = substr($val, 0, -2);
            return $val if $order eq GetByteOrder();
            return pack('v*',unpack('n*',$val));
        },
    },
    0x87b0 => { #30
        Name       => 'GeoTiffDoubleParams',
        Format     => 'undef',
        Writable   => 'double',
        WriteGroup => 'IFD0',
        Binary     => 1,
        RawConv    => '$val . GetByteOrder()', # save byte order
        # swap byte order if necessary
        RawConvInv => q{
            return $val if length $val < 2;
            my $order = substr($val, -2);
            return $val unless $order eq 'II' or $order eq 'MM';
            $val = substr($val, 0, -2);
            return $val if $order eq GetByteOrder();
            $val =~ s/(.{4})(.{4})/$2$1/sg; # swap words
            return pack('V*',unpack('N*',$val));
        },
    },
    0x87b1 => { #30
        Name       => 'GeoTiffAsciiParams',
        Format     => 'undef',
        Writable   => 'string',
        WriteGroup => 'IFD0',
        Binary     => 1,
    },
    0x87be => 'JBIGOptions', #29
    0x8822 => {
        Name      => 'ExposureProgram',
        Groups    => { 2 => 'Camera' },
        Notes     => 'the value of 9 is not standard EXIF, but is used by the Canon EOS 7D',
        Writable  => 'int16u',
        PrintConv => {
            0 => 'Not Defined',
            1 => 'Manual',
            2 => 'Program AE',
            3 => 'Aperture-priority AE',
            4 => 'Shutter speed priority AE',
            5 => 'Creative (Slow speed)',
            6 => 'Action (High speed)',
            7 => 'Portrait',
            8 => 'Landscape',
            9 => 'Bulb', #25
        },
    },
    0x8824 => {
        Name     => 'SpectralSensitivity',
        Groups   => { 2 => 'Camera' },
        Writable => 'string',
    },
    0x8825 => {
        Name         => 'GPSInfo',
        Groups       => { 1 => 'GPS' },
        WriteGroup   => 'IFD0', # (only for Validate)
        Flags        => 'SubIFD',
        SubDirectory => {
            DirName    => 'GPS',
            TagTable   => 'Image::ExifTool::GPS::Main',
            Start      => '$val',
            MaxSubdirs => 1,
        },
    },
    0x8827 => {
        Name         => 'ISO',
        Notes        => q{
            called ISOSpeedRatings by EXIF 2.2, then PhotographicSensitivity by the EXIF
            2.3 spec.
        },
        Writable     => 'int16u',
        Count        => -1,
        PrintConv    => '$val=~s/\s+/, /g; $val',
        PrintConvInv => '$val=~tr/,//d; $val',
    },
    0x8828 => {
        Name   => 'Opto-ElectricConvFactor',
        Notes  => 'called OECF by the EXIF spec.',
        Binary => 1,
    },
    0x8829 => 'Interlace', #12
    0x882a => {            #12
        Name     => 'TimeZoneOffset',
        Writable => 'int16s',
        Count    => -1, # can be 1 or 2
        Notes    => q{
            1 or 2 values: 1. The time zone offset of DateTimeOriginal from GMT in
            hours, 2. If present, the time zone offset of ModifyDate
        },
    },
    0x882b => { #12
        Name     => 'SelfTimerMode',
        Writable => 'int16u',
    },
    0x8830 => { #24
        Name      => 'SensitivityType',
        Notes     => 'applies to EXIF:ISO tag',
        Writable  => 'int16u',
        PrintConv => {
            0 => 'Unknown',
            1 => 'Standard Output Sensitivity',
            2 => 'Recommended Exposure Index',
            3 => 'ISO Speed',
            4 => 'Standard Output Sensitivity and Recommended Exposure Index',
            5 => 'Standard Output Sensitivity and ISO Speed',
            6 => 'Recommended Exposure Index and ISO Speed',
            7 => 'Standard Output Sensitivity, Recommended Exposure Index and ISO Speed',
        },
    },
    0x8831 => { #24
        Name     => 'StandardOutputSensitivity',
        Writable => 'int32u',
    },
    0x8832 => { #24
        Name     => 'RecommendedExposureIndex',
        Writable => 'int32u',
    },
    0x8833 => { #24
        Name     => 'ISOSpeed',
        Writable => 'int32u',
    },
    0x8834 => { #24
        Name        => 'ISOSpeedLatitudeyyy',
        Description => 'ISO Speed Latitude yyy',
        Writable    => 'int32u',
    },
    0x8835 => { #24
        Name        => 'ISOSpeedLatitudezzz',
        Description => 'ISO Speed Latitude zzz',
        Writable    => 'int32u',
    },
    0x885c => 'FaxRecvParams', #9
    0x885d => 'FaxSubAddress', #9
    0x885e => 'FaxRecvTime',   #9
    0x8871 => 'FedexEDR',      #exifprobe (NC)
    # 0x8889 - string: "portrait" (ExifIFD, Xiaomi POCO F1)
    0x888a => { #PH
        Name         => 'LeafSubIFD',
        Format       => 'int32u', # Leaf incorrectly uses 'undef' format!
        Groups       => { 1 => 'LeafSubIFD' },
        Flags        => 'SubIFD',
        SubDirectory => {
            TagTable => 'Image::ExifTool::Leaf::SubIFD',
            Start    => '$val',
        },
    },
    # 0x8891 - int16u: 35 (ExifIFD, Xiaomi POCO F1)
    # 0x8894 - int16u: 0 (ExifIFD, Xiaomi POCO F1)
    # 0x8895 - int16u: 0 (ExifIFD, Xiaomi POCO F1)
    # 0x889a - int16u: 0 (ExifIFD, Xiaomi POCO F1)
    # 0x89ab - seen "11 100 130 16 0 0 0 0" in IFD0 of TIFF image from IR scanner (forum8470)
    0x9000 => {
        Name         => 'ExifVersion',
        Writable     => 'undef',
        Mandatory    => 1,
        RawConv      => '$val=~s/\0+$//; $val', # (some idiots add null terminators)
        # (allow strings like "2.31" when writing)
        PrintConvInv => '$val=~tr/.//d; $val=~/^\d{4}$/ ? $val : $val =~ /^\d{3}$/ ? "0$val" : undef',
    },
    0x9003 => {
        Name         => 'DateTimeOriginal',
        Description  => 'Date/Time Original',
        Groups       => { 2 => 'Time' },
        Notes        => 'date/time when original image was taken',
        Writable     => 'string',
        Shift        => 'Time',
        Validate     => 'ValidateExifDate($val)',
        PrintConv    => '$self->ConvertDateTime($val)',
        PrintConvInv => '$self->InverseDateTime($val,0)',
    },
    0x9004 => {
        Name         => 'CreateDate',
        Groups       => { 2 => 'Time' },
        Notes        => 'called DateTimeDigitized by the EXIF spec.',
        Writable     => 'string',
        Shift        => 'Time',
        Validate     => 'ValidateExifDate($val)',
        PrintConv    => '$self->ConvertDateTime($val)',
        PrintConvInv => '$self->InverseDateTime($val,0)',
    },
    0x9009 => { # undef[44] (or undef[11]) written by Google Plus uploader - PH
        Name     => 'GooglePlusUploadCode',
        Format   => 'int8u',
        Writable => 'undef',
        Count    => -1,
    },
    0x9010 => {
        Name         => 'OffsetTime',
        Groups       => { 2 => 'Time' },
        Notes        => 'time zone for ModifyDate',
        Writable     => 'string',
        PrintConvInv => q{
            return "+00:00" if $val =~ /\d{2}Z$/;
            return sprintf("%s%.2d:%.2d",$1,$2,$3) if $val =~ /([-+])(\d{1,2}):(\d{2})/;
            return undef;
        },
    },
    0x9011 => {
        Name         => 'OffsetTimeOriginal',
        Groups       => { 2 => 'Time' },
        Notes        => 'time zone for DateTimeOriginal',
        Writable     => 'string',
        PrintConvInv => q{
            return "+00:00" if $val =~ /\d{2}Z$/;
            return sprintf("%s%.2d:%.2d",$1,$2,$3) if $val =~ /([-+])(\d{1,2}):(\d{2})/;
            return undef;
        },
    },
    0x9012 => {
        Name         => 'OffsetTimeDigitized',
        Groups       => { 2 => 'Time' },
        Notes        => 'time zone for CreateDate',
        Writable     => 'string',
        PrintConvInv => q{
            return "+00:00" if $val =~ /\d{2}Z$/;
            return sprintf("%s%.2d:%.2d",$1,$2,$3) if $val =~ /([-+])(\d{1,2}):(\d{2})/;
            return undef;
        },
    },
    0x9101 => {
        Name             => 'ComponentsConfiguration',
        Format           => 'int8u',
        Protected        => 1,
        Writable         => 'undef',
        Count            => 4,
        Mandatory        => 1,
        ValueConvInv     => '$val=~tr/,//d; $val', # (so we can copy from XMP with -n)
        PrintConvColumns => 2,
        PrintConv        => {
            0     => '-',
            1     => 'Y',
            2     => 'Cb',
            3     => 'Cr',
            4     => 'R',
            5     => 'G',
            6     => 'B',
            OTHER => sub {
                my ($val, $inv, $conv) = @_;
                my @a = split /,?\s+/, $val;
                if ($inv) {
                    my %invConv;
                    $invConv{lc $$conv{$_}} = $_ foreach keys %$conv;
                    # strings like "YCbCr" and "RGB" still work for writing
                    @a = $a[0] =~ /(Y|Cb|Cr|R|G|B)/g if @a == 1;
                    foreach (@a) {
                        $_ = $invConv{lc $_};
                        return undef unless defined $_;
                    }
                    push @a, 0 while @a < 4;
                }
                else {
                    foreach (@a) {
                        $_ = $$conv{$_} || "Err ($_)";
                    }
                }
                return join ', ', @a;
            },
        },
    },
    0x9102 => {
        Name      => 'CompressedBitsPerPixel',
        Protected => 1,
        Writable  => 'rational64u',
    },
    # 0x9103 - int16u: 1 (found in Pentax XG-1 samples)
    0x9201 => {
        Name         => 'ShutterSpeedValue',
        Notes        => 'displayed in seconds, but stored as an APEX value',
        Format       => 'rational64s', # Leica M8 patch (incorrectly written as rational64u)
        Writable     => 'rational64s',
        ValueConv    => 'abs($val)<100 ? 2**(-$val) : 0',
        ValueConvInv => '$val>0 ? -log($val)/log(2) : -100',
        PrintConv    => 'Image::ExifTool::Exif::PrintExposureTime($val)',
        PrintConvInv => 'Image::ExifTool::Exif::ConvertFraction($val)',
    },
    0x9202 => {
        Name         => 'ApertureValue',
        Notes        => 'displayed as an F number, but stored as an APEX value',
        Writable     => 'rational64u',
        ValueConv    => '2 ** ($val / 2)',
        ValueConvInv => '$val>0 ? 2*log($val)/log(2) : 0',
        PrintConv    => 'sprintf("%.1f",$val)',
        PrintConvInv => '$val',
    },
    # Wikipedia: BrightnessValue = Bv = Av + Tv - Sv
    # ExifTool:  LightValue = LV = Av + Tv - Sv + 5 (5 is the Sv for ISO 100 in Exif usage)
    0x9203 => {
        Name     => 'BrightnessValue',
        Writable => 'rational64s',
    },
    0x9204 => {
        Name         => 'ExposureCompensation',
        Format       => 'rational64s', # Leica M8 patch (incorrectly written as rational64u)
        Notes        => 'called ExposureBiasValue by the EXIF spec.',
        Writable     => 'rational64s',
        PrintConv    => 'Image::ExifTool::Exif::PrintFraction($val)',
        PrintConvInv => '$val',
    },
    0x9205 => {
        Name         => 'MaxApertureValue',
        Notes        => 'displayed as an F number, but stored as an APEX value',
        Groups       => { 2 => 'Camera' },
        Writable     => 'rational64u',
        ValueConv    => '2 ** ($val / 2)',
        ValueConvInv => '$val>0 ? 2*log($val)/log(2) : 0',
        PrintConv    => 'sprintf("%.1f",$val)',
        PrintConvInv => '$val',
    },
    0x9206 => {
        Name         => 'SubjectDistance',
        Groups       => { 2 => 'Camera' },
        Writable     => 'rational64u',
        PrintConv    => '$val =~ /^(inf|undef)$/ ? $val : "${val} m"',
        PrintConvInv => '$val=~s/\s*m$//;$val',
    },
    0x9207 => {
        Name      => 'MeteringMode',
        Groups    => { 2 => 'Camera' },
        Writable  => 'int16u',
        PrintConv => {
            0   => 'Unknown',
            1   => 'Average',
            2   => 'Center-weighted average',
            3   => 'Spot',
            4   => 'Multi-spot',
            5   => 'Multi-segment',
            6   => 'Partial',
            255 => 'Other',
        },
    },
    0x9208 => {
        Name          => 'LightSource',
        Groups        => { 2 => 'Camera' },
        Writable      => 'int16u',
        SeparateTable => 'LightSource',
        PrintConv     => \%lightSource,
    },
    0x9209 => {
        Name          => 'Flash',
        Groups        => { 2 => 'Camera' },
        Writable      => 'int16u',
        Flags         => 'PrintHex',
        SeparateTable => 'Flash',
        PrintConv     => \%flash,
    },
    0x920a => {
        Name         => 'FocalLength',
        Groups       => { 2 => 'Camera' },
        Writable     => 'rational64u',
        PrintConv    => 'sprintf("%.1f mm",$val)',
        PrintConvInv => '$val=~s/\s*mm$//;$val',
    },
    # Note: tags 0x920b-0x9217 are duplicates of 0xa20b-0xa217
    # (The EXIF standard uses 0xa2xx, but you'll find both in images)
    0x920b => { #12
        Name   => 'FlashEnergy',
        Groups => { 2 => 'Camera' },
    },
    0x920c => 'SpatialFrequencyResponse', #12 (not in Fuji images - PH)
    0x920d => 'Noise',                    #12
    0x920e => 'FocalPlaneXResolution',    #12
    0x920f => 'FocalPlaneYResolution',    #12
    0x9210 => {                           #12
        Name      => 'FocalPlaneResolutionUnit',
        Groups    => { 2 => 'Camera' },
        PrintConv => {
            1 => 'None',
            2 => 'inches',
            3 => 'cm',
            4 => 'mm',
            5 => 'um',
        },
    },
    0x9211 => { #12
        Name     => 'ImageNumber',
        Writable => 'int32u',
    },
    0x9212 => { #12
        Name      => 'SecurityClassification',
        Writable  => 'string',
        PrintConv => {
            T => 'Top Secret',
            S => 'Secret',
            C => 'Confidential',
            R => 'Restricted',
            U => 'Unclassified',
        },
    },
    0x9213 => { #12
        Name     => 'ImageHistory',
        Writable => 'string',
    },
    0x9214 => {
        Name     => 'SubjectArea',
        Groups   => { 2 => 'Camera' },
        Writable => 'int16u',
        Count    => -1, # 2, 3 or 4 values
    },
    0x9215 => 'ExposureIndex',     #12
    0x9216 => 'TIFF-EPStandardID', #12
    0x9217 => {                    #12
        Name      => 'SensingMethod',
        Groups    => { 2 => 'Camera' },
        PrintConv => {
            # (values 1 and 6 are not used by corresponding EXIF tag 0xa217)
            1 => 'Monochrome area',
            2 => 'One-chip color area',
            3 => 'Two-chip color area',
            4 => 'Three-chip color area',
            5 => 'Color sequential area',
            6 => 'Monochrome linear',
            7 => 'Trilinear',
            8 => 'Color sequential linear',
        },
    },
    0x923a => 'CIP3DataFile', #20
    0x923b => 'CIP3Sheet',    #20
    0x923c => 'CIP3Side',     #20
    0x923f => 'StoNits',      #9
    # handle maker notes as a conditional list
    #0x927c => \@Image::ExifTool::MakerNotes::Main,
    0x927c => {
        Name        => 'MakerNote',
        Description => 'A tag for manufacturers of Exif writers to record any desired information. The contents are up to the manufacturer.',
        Writable    => 'undef',
    },
    0x9286 => {
        Name       => 'UserComment',
        # I have seen other applications write it incorrectly as 'string' or 'int8u'
        Format     => 'undef',
        Writable   => 'undef',
        RawConv    => 'Image::ExifTool::Exif::ConvertExifText($self,$val,1,$tag)',
        #  (starts with "ASCII\0\0\0", "UNICODE\0", "JIS\0\0\0\0\0" or "\0\0\0\0\0\0\0\0")
        RawConvInv => 'Image::ExifTool::Exif::EncodeExifText($self,$val)',
        # SHOULD ADD SPECIAL LOGIC TO ALLOW CONDITIONAL OVERWRITE OF
        # "UNKNOWN" VALUES FILLED WITH SPACES
    },
    0x9290 => {
        Name         => 'SubSecTime',
        Groups       => { 2 => 'Time' },
        Notes        => 'fractional seconds for ModifyDate',
        Writable     => 'string',
        ValueConv    => '$val=~s/ +$//; $val', # trim trailing blanks
        # extract fractional seconds from a full date/time value
        ValueConvInv => '$val=~/^(\d+)\s*$/ ? $1 : ($val=~/\.(\d+)/ ? $1 : undef)',
    },
    0x9291 => {
        Name         => 'SubSecTimeOriginal',
        Groups       => { 2 => 'Time' },
        Notes        => 'fractional seconds for DateTimeOriginal',
        Writable     => 'string',
        ValueConv    => '$val=~s/ +$//; $val', # trim trailing blanks
        ValueConvInv => '$val=~/^(\d+)\s*$/ ? $1 : ($val=~/\.(\d+)/ ? $1 : undef)',
    },
    0x9292 => {
        Name         => 'SubSecTimeDigitized',
        Groups       => { 2 => 'Time' },
        Notes        => 'fractional seconds for CreateDate',
        Writable     => 'string',
        ValueConv    => '$val=~s/ +$//; $val', # trim trailing blanks
        ValueConvInv => '$val=~/^(\d+)\s*$/ ? $1 : ($val=~/\.(\d+)/ ? $1 : undef)',
    },
    # The following 3 tags are found in MSOffice TIFF images
    # References:
    # http://social.msdn.microsoft.com/Forums/en-US/os_standocs/thread/03086d55-294a-49d5-967a-5303d34c40f8/
    # http://blogs.msdn.com/openspecification/archive/2009/12/08/details-of-three-tiff-tag-extensions-that-microsoft-office-document-imaging-modi-software-may-write-into-the-tiff-files-it-generates.aspx
    # http://www.microsoft.com/downloads/details.aspx?FamilyID=0dbc435d-3544-4f4b-9092-2f2643d64a39&displaylang=en#filelist
    0x932f => 'MSDocumentText',
    0x9330 => {
        Name   => 'MSPropertySetStorage',
        Binary => 1,
    },
    0x9331 => {
        Name   => 'MSDocumentTextPosition',
        Binary => 1, # (just in case -- don't know what format this is)
    },
    0x935c => {                            #3/19
        Name         => 'ImageSourceData', # (writable directory!)
        Writable     => 'undef',
        WriteGroup   => 'IFD0',
        SubDirectory => { TagTable => 'Image::ExifTool::Photoshop::DocumentData' },
        Binary       => 1,
        Protected    => 1, # (because this can be hundreds of megabytes)
        ReadFromRAF  => 1, # don't load into memory when reading
    },
    0x9400 => {
        Name         => 'AmbientTemperature',
        Notes        => 'ambient temperature in degrees C, called Temperature by the EXIF spec.',
        Writable     => 'rational64s',
        PrintConv    => '"$val C"',
        PrintConvInv => '$val=~s/ ?C//; $val',
    },
    0x9401 => {
        Name     => 'Humidity',
        Notes    => 'ambient relative humidity in percent',
        Writable => 'rational64u',
    },
    0x9402 => {
        Name     => 'Pressure',
        Notes    => 'air pressure in hPa or mbar',
        Writable => 'rational64u',
    },
    0x9403 => {
        Name     => 'WaterDepth',
        Notes    => 'depth under water in metres, negative for above water',
        Writable => 'rational64s',
    },
    0x9404 => {
        Name     => 'Acceleration',
        Notes    => 'directionless camera acceleration in units of mGal, or 10-5 m/s2',
        Writable => 'rational64u',
    },
    0x9405 => {
        Name     => 'CameraElevationAngle',
        Writable => 'rational64s',
    },
    # 0x9999 - string: camera settings (ExifIFD, Xiaomi POCO F1)
    # 0x9aaa - int8u[2176]: ? (ExifIFD, Xiaomi POCO F1)
    0x9c9b => {
        Name         => 'XPTitle',
        Format       => 'undef',
        Writable     => 'int8u',
        WriteGroup   => 'IFD0',
        Notes        => q{
            tags 0x9c9b-0x9c9f are used by Windows Explorer; special characters
            in these values are converted to UTF-8 by default, or Windows Latin1
            with the -L option.  XPTitle is ignored by Windows Explorer if
            ImageDescription exists
        },
        ValueConv    => '$self->Decode($val,"UCS2","II")',
        ValueConvInv => '$self->Encode($val,"UCS2","II") . "\0\0"',
    },
    0x9c9c => {
        Name         => 'XPComment',
        Format       => 'undef',
        Writable     => 'int8u',
        WriteGroup   => 'IFD0',
        ValueConv    => '$self->Decode($val,"UCS2","II")',
        ValueConvInv => '$self->Encode($val,"UCS2","II") . "\0\0"',
    },
    0x9c9d => {
        Name         => 'XPAuthor',
        Groups       => { 2 => 'Author' },
        Format       => 'undef',
        Writable     => 'int8u',
        WriteGroup   => 'IFD0',
        Notes        => 'ignored by Windows Explorer if Artist exists',
        ValueConv    => '$self->Decode($val,"UCS2","II")',
        ValueConvInv => '$self->Encode($val,"UCS2","II") . "\0\0"',
    },
    0x9c9e => {
        Name         => 'XPKeywords',
        Format       => 'undef',
        Writable     => 'int8u',
        WriteGroup   => 'IFD0',
        ValueConv    => '$self->Decode($val,"UCS2","II")',
        ValueConvInv => '$self->Encode($val,"UCS2","II") . "\0\0"',
    },
    0x9c9f => {
        Name         => 'XPSubject',
        Format       => 'undef',
        Writable     => 'int8u',
        WriteGroup   => 'IFD0',
        ValueConv    => '$self->Decode($val,"UCS2","II")',
        ValueConvInv => '$self->Encode($val,"UCS2","II") . "\0\0"',
    },
    0xa000 => {
        Name         => 'FlashpixVersion',
        Writable     => 'undef',
        Mandatory    => 1,
        RawConv      => '$val=~s/\0+$//; $val', # (some idiots add null terminators)
        PrintConvInv => '$val=~tr/.//d; $val=~/^\d{4}$/ ? $val : undef',
    },
    0xa001 => {
        Name      => 'ColorSpace',
        Notes     => q{
            the value of 0x2 is not standard EXIF.  Instead, an Adobe RGB image is
            indicated by "Uncalibrated" with an InteropIndex of "R03".  The values
            0xfffd and 0xfffe are also non-standard, and are used by some Sony cameras
        },
        Writable  => 'int16u',
        Mandatory => 1,
        PrintHex  => 1,
        PrintConv => {
            1      => 'sRGB',
            2      => 'Adobe RGB',
            0xffff => 'Uncalibrated',
            # Sony uses these definitions: (ref JD)
            # 0xffff => 'Adobe RGB', (conflicts with Uncalibrated)
            0xfffe => 'ICC Profile',
            0xfffd => 'Wide Gamut RGB',
        },
    },
    0xa002 => {
        Name      => 'ExifImageWidth',
        Notes     => 'called PixelXDimension by the EXIF spec.',
        Writable  => 'int16u',
        Mandatory => 1,
    },
    0xa003 => {
        Name      => 'ExifImageHeight',
        Notes     => 'called PixelYDimension by the EXIF spec.',
        Writable  => 'int16u',
        Mandatory => 1,
    },
    0xa004 => {
        Name     => 'RelatedSoundFile',
        Writable => 'string',
    },
    0xa005 => {
        Name         => 'InteropOffset',
        Groups       => { 1 => 'InteropIFD' },
        Flags        => 'SubIFD',
        Description  => 'Interoperability Offset',
        SubDirectory => {
            DirName    => 'InteropIFD',
            Start      => '$val',
            MaxSubdirs => 1,
        },
    },
    # the following 4 tags found in SubIFD1 of some Samsung SRW images
    0xa010 => {
        Name       => 'SamsungRawPointersOffset',
        IsOffset   => 1,
        OffsetPair => 0xa011, # point to associated byte count
    },
    0xa011 => {
        Name       => 'SamsungRawPointersLength',
        OffsetPair => 0xa010, # point to associated offset
    },
    0xa101 => {
        Name      => 'SamsungRawByteOrder',
        Format    => 'undef',
        # this is written incorrectly as string[1], but is "\0\0MM" or "II\0\0"
        FixedSize => 4,
        Count     => 1,
    },
    0xa102 => {
        Name    => 'SamsungRawUnknown',
        Unknown => 1,
    },
    0xa20b => {
        Name     => 'FlashEnergy',
        Groups   => { 2 => 'Camera' },
        Writable => 'rational64u',
    },
    0xa20c => {
        Name      => 'SpatialFrequencyResponse',
        PrintConv => 'Image::ExifTool::Exif::PrintSFR($val)',
    },
    0xa20d => 'Noise',
    0xa20e => {
        Name     => 'FocalPlaneXResolution',
        Groups   => { 2 => 'Camera' },
        Writable => 'rational64u',
    },
    0xa20f => {
        Name     => 'FocalPlaneYResolution',
        Groups   => { 2 => 'Camera' },
        Writable => 'rational64u',
    },
    0xa210 => {
        Name      => 'FocalPlaneResolutionUnit',
        Groups    => { 2 => 'Camera' },
        Notes     => 'values 1, 4 and 5 are not standard EXIF',
        Writable  => 'int16u',
        PrintConv => {
            1 => 'None', # (not standard EXIF)
            2 => 'inches',
            3 => 'cm',
            4 => 'mm', # (not standard EXIF)
            5 => 'um', # (not standard EXIF)
        },
    },
    0xa211 => 'ImageNumber',
    0xa212 => 'SecurityClassification',
    0xa213 => 'ImageHistory',
    0xa214 => {
        Name     => 'SubjectLocation',
        Groups   => { 2 => 'Camera' },
        Writable => 'int16u',
        Count    => 2,
    },
    0xa215 => { Name => 'ExposureIndex', Writable => 'rational64u' },
    0xa216 => 'TIFF-EPStandardID',
    0xa217 => {
        Name      => 'SensingMethod',
        Groups    => { 2 => 'Camera' },
        Writable  => 'int16u',
        PrintConv => {
            1 => 'Not defined',
            2 => 'One-chip color area',
            3 => 'Two-chip color area',
            4 => 'Three-chip color area',
            5 => 'Color sequential area',
            7 => 'Trilinear',
            8 => 'Color sequential linear',
            # 15 - used by DJI XT2
        },
    },
    0xa300 => {
        Name         => 'FileSource',
        Writable     => 'undef',
        ValueConvInv => '($val=~/^\d+$/ and $val < 256) ? chr($val) : $val',
        PrintConv    => {
            1          => 'Film Scanner',
            2          => 'Reflection Print Scanner',
            3          => 'Digital Camera',
            # handle the case where Sigma incorrectly gives this tag a count of 4
            "\3\0\0\0" => 'Sigma Digital Camera',
        },
    },
    0xa301 => {
        Name         => 'SceneType',
        Writable     => 'undef',
        ValueConvInv => 'chr($val & 0xff)',
        PrintConv    => {
            1 => 'Directly photographed',
        },
    },
    0xa302 => {
        Name         => 'CFAPattern',
        Writable     => 'undef',
        RawConv      => 'Image::ExifTool::Exif::DecodeCFAPattern($self, $val)',
        RawConvInv   => q{
            my @a = split ' ', $val;
            return $val if @a <= 2; # also accept binary data for backward compatibility
            return pack(GetByteOrder() eq 'II' ? 'v2C*' : 'n2C*', @a);
        },
        PrintConv    => 'Image::ExifTool::Exif::PrintCFAPattern($val)',
        PrintConvInv => 'Image::ExifTool::Exif::GetCFAPattern($val)',
    },
    0xa401 => {
        Name      => 'CustomRendered',
        Writable  => 'int16u',
        Notes     => q{
            only 0 and 1 are standard EXIF, but other values are used by Apple iOS
            devices
        },
        PrintConv => {
            0 => 'Normal',
            1 => 'Custom',
            2 => 'HDR (no original saved)', #32 non-standard (Apple iOS)
            3 => 'HDR (original saved)',    #32 non-standard (Apple iOS)
            4 => 'Original (for HDR)',      #32 non-standard (Apple iOS)
            6 => 'Panorama',                # non-standard (Apple iOS, horizontal or vertical)
            7 => 'Portrait HDR',            #32 non-standard (Apple iOS)
            8 => 'Portrait',                # non-standard (Apple iOS, blurred background)
            # 9 - also seen (Apple iOS) (HDR Portrait?)
        },
    },
    0xa402 => {
        Name      => 'ExposureMode',
        Groups    => { 2 => 'Camera' },
        Writable  => 'int16u',
        PrintConv => {
            0 => 'Auto',
            1 => 'Manual',
            2 => 'Auto bracket',
            # have seen 3 from Samsung EX1, NX30, NX200 - PH
        },
    },
    0xa403 => {
        Name      => 'WhiteBalance',
        Groups    => { 2 => 'Camera' },
        Writable  => 'int16u',
        # set Priority to zero to keep this WhiteBalance from overriding the
        # MakerNotes WhiteBalance, since the MakerNotes WhiteBalance and is more
        # accurate and contains more information (if it exists)
        Priority  => 0,
        PrintConv => {
            0 => 'Auto',
            1 => 'Manual',
        },
    },
    0xa404 => {
        Name     => 'DigitalZoomRatio',
        Groups   => { 2 => 'Camera' },
        Writable => 'rational64u',
    },
    0xa405 => {
        Name         => 'FocalLengthIn35mmFormat',
        Notes        => 'called FocalLengthIn35mmFilm by the EXIF spec.',
        Groups       => { 2 => 'Camera' },
        Writable     => 'int16u',
        PrintConv    => '"$val mm"',
        PrintConvInv => '$val=~s/\s*mm$//;$val',
    },
    0xa406 => {
        Name      => 'SceneCaptureType',
        Groups    => { 2 => 'Camera' },
        Writable  => 'int16u',
        Notes     => 'the value of 4 is non-standard, and used by some Samsung models',
        PrintConv => {
            0 => 'Standard',
            1 => 'Landscape',
            2 => 'Portrait',
            3 => 'Night',
            4 => 'Other', # (non-standard Samsung, ref forum 5724)
        },
    },
    0xa407 => {
        Name      => 'GainControl',
        Groups    => { 2 => 'Camera' },
        Writable  => 'int16u',
        PrintConv => {
            0 => 'None',
            1 => 'Low gain up',
            2 => 'High gain up',
            3 => 'Low gain down',
            4 => 'High gain down',
        },
    },
    0xa408 => {
        Name         => 'Contrast',
        Groups       => { 2 => 'Camera' },
        Writable     => 'int16u',
        PrintConv    => {
            0 => 'Normal',
            1 => 'Low',
            2 => 'High',
        },
        PrintConvInv => 'Image::ExifTool::Exif::ConvertParameter($val)',
    },
    0xa409 => {
        Name         => 'Saturation',
        Groups       => { 2 => 'Camera' },
        Writable     => 'int16u',
        PrintConv    => {
            0 => 'Normal',
            1 => 'Low',
            2 => 'High',
        },
        PrintConvInv => 'Image::ExifTool::Exif::ConvertParameter($val)',
    },
    0xa40a => {
        Name         => 'Sharpness',
        Groups       => { 2 => 'Camera' },
        Writable     => 'int16u',
        PrintConv    => {
            0 => 'Normal',
            1 => 'Soft',
            2 => 'Hard',
        },
        PrintConvInv => 'Image::ExifTool::Exif::ConvertParameter($val)',
    },
    0xa40b => {
        Name   => 'DeviceSettingDescription',
        Groups => { 2 => 'Camera' },
        Binary => 1,
    },
    0xa40c => {
        Name      => 'SubjectDistanceRange',
        Groups    => { 2 => 'Camera' },
        Writable  => 'int16u',
        PrintConv => {
            0 => 'Unknown',
            1 => 'Macro',
            2 => 'Close',
            3 => 'Distant',
        },
    },
    # 0xa40d - int16u: 0 (GE E1486 TW)
    # 0xa40e - int16u: 1 (GE E1486 TW)
    0xa420 => { Name => 'ImageUniqueID', Writable => 'string' },
    0xa430 => { #24
        Name     => 'OwnerName',
        Notes    => 'called CameraOwnerName by the EXIF spec.',
        Writable => 'string',
    },
    0xa431 => { #24
        Name     => 'SerialNumber',
        Notes    => 'called BodySerialNumber by the EXIF spec.',
        Writable => 'string',
    },
    0xa432 => { #24
        Name         => 'LensInfo',
        Notes        => q{
            4 rational values giving focal and aperture ranges, called LensSpecification
            by the EXIF spec.
        },
        Writable     => 'rational64u',
        Count        => 4,
        # convert to the form "12-20mm f/3.8-4.5" or "50mm f/1.4"
        PrintConv    => \&PrintLensInfo,
        PrintConvInv => \&ConvertLensInfo,
    },
    0xa433 => { Name => 'LensMake', Writable => 'string' },         #24
    0xa434 => { Name => 'LensModel', Writable => 'string' },        #24
    0xa435 => { Name => 'LensSerialNumber', Writable => 'string' }, #24
    0xa460 => {                                                     #Exif2.32
        Name      => 'CompositeImage',
        Writable  => 'int16u',
        PrintConv => {
            0 => 'Unknown',
            1 => 'Not a Composite Image',
            2 => 'General Composite Image',
            3 => 'Composite Image Captured While Shooting',
        },
    },
    0xa461 => { #Exif2.32
        Name     => 'CompositeImageCount',
        Notes    => q{
            2 values: 1. Number of source images, 2. Number of images used.  Called
            SourceImageNumberOfCompositeImage by the EXIF spec.
        },
        Writable => 'int16u',
        Count    => 2,
    },
    0xa462 => { #Exif2.32
        Name         => 'CompositeImageExposureTimes',
        Notes        => q{
            11 or more values: 1. Total exposure time period, 2. Total exposure of all
            source images, 3. Total exposure of all used images, 4. Max exposure time of
            source images, 5. Max exposure time of used images, 6. Min exposure time of
            source images, 7. Min exposure of used images, 8. Number of sequences, 9.
            Number of source images in sequence. 10-N. Exposure times of each source
            image. Called SourceExposureTimesOfCompositeImage by the EXIF spec.
        },
        Writable     => 'undef',
        RawConv      => sub {
            my $val = shift;
            my @v;
            my $i = 0;
            for (;;) {
                if ($i == 56 or $i == 58) {
                    last if $i + 2 > length $val;
                    push @v, Get16u(\$val, $i);
                    $i += 2;
                }
                else {
                    last if $i + 8 > length $val;
                    push @v, Image::ExifTool::GetRational64u(\$val, $i);
                    $i += 8;
                }
            }
            return join ' ', @v;
        },
        RawConvInv   => sub {
            my $val = shift;
            my @v = split ' ', $val;
            my $i;
            for ($i = 0;; ++$i) {
                last unless defined $v[$i];
                $v[$i] = ($i == 7 or $i == 8) ? Set16u($v[$i]) : Image::ExifTool::SetRational64u($v[$i]);
            }
            return join '', @v;
        },
        PrintConv    => sub {
            my $val = shift;
            my @v = split ' ', $val;
            my $i;
            for ($i = 0;; ++$i) {
                last unless defined $v[$i];
                $v[$i] = PrintExposureTime($v[$i]) unless $i == 7 or $i == 8;
            }
            return join ' ', @v;
        },
        PrintConvInv => '$val',
    },
    0xa480 => { Name => 'GDALMetadata', Writable => 'string', WriteGroup => 'IFD0' }, #3
    0xa481 => { Name => 'GDALNoData', Writable => 'string', WriteGroup => 'IFD0' },   #3
    0xa500 => { Name => 'Gamma', Writable => 'rational64u' },
    0xafc0 => 'ExpandSoftware',                                                 #JD (Opanda)
    0xafc1 => 'ExpandLens',                                                     #JD (Opanda)
    0xafc2 => 'ExpandFilm',                                                     #JD (Opanda)
    0xafc3 => 'ExpandFilterLens',                                               #JD (Opanda)
    0xafc4 => 'ExpandScanner',                                                  #JD (Opanda)
    0xafc5 => 'ExpandFlashLamp',                                                #JD (Opanda)
    0xb4c3 => { Name => 'HasselbladRawImage', Format => 'undef', Binary => 1 }, #IB
    #
    # Windows Media Photo / HD Photo (WDP/HDP) tags
    #
    0xbc01 => { #13
        Name      => 'PixelFormat',
        PrintHex  => 1,
        Format    => 'undef',
        Notes     => q{
            tags 0xbc** are used in Windows HD Photo (HDP and WDP) images. The actual
            PixelFormat values are 16-byte GUID's but the leading 15 bytes,
            '6fddc324-4e03-4bfe-b1853-d77768dc9', have been removed below to avoid
            unnecessary clutter
        },
        ValueConv => q{
            require Image::ExifTool::ASF;
            $val = Image::ExifTool::ASF::GetGUID($val);
            # GUID's are too long, so remove redundant information
            $val =~ s/^6fddc324-4e03-4bfe-b185-3d77768dc9//i and $val = hex($val);
            return $val;
        },
        PrintConv => {
            0x0d => '24-bit RGB',
            0x0c => '24-bit BGR',
            0x0e => '32-bit BGR',
            0x15 => '48-bit RGB',
            0x12 => '48-bit RGB Fixed Point',
            0x3b => '48-bit RGB Half',
            0x18 => '96-bit RGB Fixed Point',
            0x1b => '128-bit RGB Float',
            0x0f => '32-bit BGRA',
            0x16 => '64-bit RGBA',
            0x1d => '64-bit RGBA Fixed Point',
            0x3a => '64-bit RGBA Half',
            0x1e => '128-bit RGBA Fixed Point',
            0x19 => '128-bit RGBA Float',
            0x10 => '32-bit PBGRA',
            0x17 => '64-bit PRGBA',
            0x1a => '128-bit PRGBA Float',
            0x1c => '32-bit CMYK',
            0x2c => '40-bit CMYK Alpha',
            0x1f => '64-bit CMYK',
            0x2d => '80-bit CMYK Alpha',
            0x20 => '24-bit 3 Channels',
            0x21 => '32-bit 4 Channels',
            0x22 => '40-bit 5 Channels',
            0x23 => '48-bit 6 Channels',
            0x24 => '56-bit 7 Channels',
            0x25 => '64-bit 8 Channels',
            0x2e => '32-bit 3 Channels Alpha',
            0x2f => '40-bit 4 Channels Alpha',
            0x30 => '48-bit 5 Channels Alpha',
            0x31 => '56-bit 6 Channels Alpha',
            0x32 => '64-bit 7 Channels Alpha',
            0x33 => '72-bit 8 Channels Alpha',
            0x26 => '48-bit 3 Channels',
            0x27 => '64-bit 4 Channels',
            0x28 => '80-bit 5 Channels',
            0x29 => '96-bit 6 Channels',
            0x2a => '112-bit 7 Channels',
            0x2b => '128-bit 8 Channels',
            0x34 => '64-bit 3 Channels Alpha',
            0x35 => '80-bit 4 Channels Alpha',
            0x36 => '96-bit 5 Channels Alpha',
            0x37 => '112-bit 6 Channels Alpha',
            0x38 => '128-bit 7 Channels Alpha',
            0x39 => '144-bit 8 Channels Alpha',
            0x08 => '8-bit Gray',
            0x0b => '16-bit Gray',
            0x13 => '16-bit Gray Fixed Point',
            0x3e => '16-bit Gray Half',
            0x3f => '32-bit Gray Fixed Point',
            0x11 => '32-bit Gray Float',
            0x05 => 'Black & White',
            0x09 => '16-bit BGR555',
            0x0a => '16-bit BGR565',
            0x13 => '32-bit BGR101010',
            0x3d => '32-bit RGBE',
        },
    },
    0xbc02 => { #13
        Name      => 'Transformation',
        PrintConv => {
            0 => 'Horizontal (normal)',
            1 => 'Mirror vertical',
            2 => 'Mirror horizontal',
            3 => 'Rotate 180',
            4 => 'Rotate 90 CW',
            5 => 'Mirror horizontal and rotate 90 CW',
            6 => 'Mirror horizontal and rotate 270 CW',
            7 => 'Rotate 270 CW',
        },
    },
    0xbc03 => { #13
        Name      => 'Uncompressed',
        PrintConv => { 0 => 'No', 1 => 'Yes' },
    },
    0xbc04 => { #13
        Name      => 'ImageType',
        PrintConv => { BITMASK => {
            0 => 'Preview',
            1 => 'Page',
        } },
    },
    0xbc80 => 'ImageWidth',       #13
    0xbc81 => 'ImageHeight',      #13
    0xbc82 => 'WidthResolution',  #13
    0xbc83 => 'HeightResolution', #13
    0xbcc0 => {                   #13
        Name       => 'ImageOffset',
        IsOffset   => 1,
        OffsetPair => 0xbcc1, # point to associated byte count
    },
    0xbcc1 => { #13
        Name       => 'ImageByteCount',
        OffsetPair => 0xbcc0, # point to associated offset
    },
    0xbcc2 => { #13
        Name       => 'AlphaOffset',
        IsOffset   => 1,
        OffsetPair => 0xbcc3, # point to associated byte count
    },
    0xbcc3 => { #13
        Name       => 'AlphaByteCount',
        OffsetPair => 0xbcc2, # point to associated offset
    },
    0xbcc4 => { #13
        Name      => 'ImageDataDiscard',
        PrintConv => {
            0 => 'Full Resolution',
            1 => 'Flexbits Discarded',
            2 => 'HighPass Frequency Data Discarded',
            3 => 'Highpass and LowPass Frequency Data Discarded',
        },
    },
    0xbcc5 => { #13
        Name      => 'AlphaDataDiscard',
        PrintConv => {
            0 => 'Full Resolution',
            1 => 'Flexbits Discarded',
            2 => 'HighPass Frequency Data Discarded',
            3 => 'Highpass and LowPass Frequency Data Discarded',
        },
    },
    #
    0xc427 => 'OceScanjobDesc',                       #3
    0xc428 => 'OceApplicationSelector',               #3
    0xc429 => 'OceIDNumber',                          #3
    0xc42a => 'OceImageLogic',                        #3
    0xc44f => { Name => 'Annotations', Binary => 1 }, #7/19
    0xc4a5 => {
        Name         => 'PrintIM', # (writable directory!)
        # must set Writable here so this tag will be saved with MakerNotes option
        Writable     => 'undef',
        WriteGroup   => 'IFD0',
        Binary       => 1,
        # (don't make Binary/Protected because we can't copy individual PrintIM tags anyway)
        Description  => 'Print Image Matching',
        SubDirectory => {
            TagTable => 'Image::ExifTool::PrintIM::Main',
        },
        PrintConvInv => '$val =~ /^PrintIM/ ? $val : undef', # quick validation
    },
    0xc51b => { # (Hasselblad H3D)
        Name    => 'HasselbladExif',
        Format  => 'undef',
        RawConv => q{
            $$self{DOC_NUM} = ++$$self{DOC_COUNT};
            $self->ExtractInfo(\$val, { ReEntry => 1 });
            $$self{DOC_NUM} = 0;
            return undef;
        },
    },
    0xc573 => { #PH
        Name  => 'OriginalFileName',
        Notes => 'used by some obscure software', # (possibly Swizzy Photosmacker?)
        # (it is a 'string', but obscure, so don't make it writable)
    },
    0xc580 => { #20
        Name      => 'USPTOOriginalContentType',
        PrintConv => {
            0 => 'Text or Drawing',
            1 => 'Grayscale',
            2 => 'Color',
        },
    },
    # 0xc5d8 - found in CR2 images
    # 0xc5d9 - found in CR2 images
    0xc5e0 => { #forum8153 (CR2 images)
        Name      => 'CR2CFAPattern',
        ValueConv => {
            1 => '0 1 1 2',
            2 => '2 1 1 0',
            3 => '1 2 0 1',
            4 => '1 0 2 1',
        },
        PrintConv => {
            '0 1 1 2' => '[Red,Green][Green,Blue]',
            '2 1 1 0' => '[Blue,Green][Green,Red]',
            '1 2 0 1' => '[Green,Blue][Red,Green]',
            '1 0 2 1' => '[Green,Red][Blue,Green]',
        },
    },
    #
    # DNG tags 0xc6XX, 0xc7XX and 0xcdXX (ref 2 unless otherwise stated)
    #
    0xc612 => {
        Name         => 'DNGVersion',
        Notes        => q{
            tags 0xc612-0xcd3b are defined by the DNG specification unless otherwise
            noted.  See L<https://helpx.adobe.com/photoshop/digital-negative.html> for
            the specification
        },
        Writable     => 'int8u',
        WriteGroup   => 'IFD0',
        Count        => 4,
        Protected    => 1, # (confuses Apple Preview if written to a TIFF image)
        DataMember   => 'DNGVersion',
        RawConv      => '$$self{DNGVersion} = $val',
        PrintConv    => '$val =~ tr/ /./; $val',
        PrintConvInv => '$val =~ tr/./ /; $val',
    },
    0xc613 => {
        Name         => 'DNGBackwardVersion',
        Writable     => 'int8u',
        WriteGroup   => 'IFD0',
        Count        => 4,
        Protected    => 1,
        PrintConv    => '$val =~ tr/ /./; $val',
        PrintConvInv => '$val =~ tr/./ /; $val',
    },
    0xc614 => {
        Name       => 'UniqueCameraModel',
        Writable   => 'string',
        WriteGroup => 'IFD0',
    },
    0xc615 => {
        Name         => 'LocalizedCameraModel',
        WriteGroup   => 'IFD0',
        %utf8StringConv,
        PrintConv    => '$self->Printable($val, 0)',
        PrintConvInv => '$val',
    },
    0xc616 => {
        Name       => 'CFAPlaneColor',
        WriteGroup => 'SubIFD', # (only for Validate)
        PrintConv  => q{
            my @cols = qw(Red Green Blue Cyan Magenta Yellow White);
            my @vals = map { $cols[$_] || "Unknown($_)" } split(' ', $val);
            return join(',', @vals);
        },
    },
    0xc617 => {
        Name       => 'CFALayout',
        WriteGroup => 'SubIFD', # (only for Validate)
        PrintConv  => {
            1 => 'Rectangular',
            2 => 'Even columns offset down 1/2 row',
            3 => 'Even columns offset up 1/2 row',
            4 => 'Even rows offset right 1/2 column',
            5 => 'Even rows offset left 1/2 column',
            # the following are new for DNG 1.3:
            6 => 'Even rows offset up by 1/2 row, even columns offset left by 1/2 column',
            7 => 'Even rows offset up by 1/2 row, even columns offset right by 1/2 column',
            8 => 'Even rows offset down by 1/2 row, even columns offset left by 1/2 column',
            9 => 'Even rows offset down by 1/2 row, even columns offset right by 1/2 column',
        },
    },
    0xc618 => {
        Name       => 'LinearizationTable',
        Writable   => 'int16u',
        WriteGroup => 'SubIFD',
        Count      => -1,
        Protected  => 1,
        Binary     => 1,
    },
    0xc619 => {
        Name       => 'BlackLevelRepeatDim',
        Writable   => 'int16u',
        WriteGroup => 'SubIFD',
        Count      => 2,
        Protected  => 1,
    },
    0xc61a => {
        Name       => 'BlackLevel',
        Writable   => 'rational64u',
        WriteGroup => 'SubIFD',
        Count      => -1,
        Protected  => 1,
    },
    0xc61b => {
        Name       => 'BlackLevelDeltaH',
        %longBin,
        Writable   => 'rational64s',
        WriteGroup => 'SubIFD',
        Count      => -1,
        Protected  => 1,
    },
    0xc61c => {
        Name       => 'BlackLevelDeltaV',
        %longBin,
        Writable   => 'rational64s',
        WriteGroup => 'SubIFD',
        Count      => -1,
        Protected  => 1,
    },
    0xc61d => {
        Name       => 'WhiteLevel',
        Writable   => 'int32u',
        WriteGroup => 'SubIFD',
        Count      => -1,
        Protected  => 1,
    },
    0xc61e => {
        Name       => 'DefaultScale',
        Writable   => 'rational64u',
        WriteGroup => 'SubIFD',
        Count      => 2,
        Protected  => 1,
    },
    0xc61f => {
        Name       => 'DefaultCropOrigin',
        Writable   => 'int32u',
        WriteGroup => 'SubIFD',
        Count      => 2,
        Protected  => 1,
    },
    0xc620 => {
        Name       => 'DefaultCropSize',
        Writable   => 'int32u',
        WriteGroup => 'SubIFD',
        Count      => 2,
        Protected  => 1,
    },
    0xc621 => {
        Name       => 'ColorMatrix1',
        Writable   => 'rational64s',
        WriteGroup => 'IFD0',
        Count      => -1,
        Protected  => 1,
    },
    0xc622 => {
        Name       => 'ColorMatrix2',
        Writable   => 'rational64s',
        WriteGroup => 'IFD0',
        Count      => -1,
        Protected  => 1,
    },
    0xc623 => {
        Name       => 'CameraCalibration1',
        Writable   => 'rational64s',
        WriteGroup => 'IFD0',
        Count      => -1,
        Protected  => 1,
    },
    0xc624 => {
        Name       => 'CameraCalibration2',
        Writable   => 'rational64s',
        WriteGroup => 'IFD0',
        Count      => -1,
        Protected  => 1,
    },
    0xc625 => {
        Name       => 'ReductionMatrix1',
        Writable   => 'rational64s',
        WriteGroup => 'IFD0',
        Count      => -1,
        Protected  => 1,
    },
    0xc626 => {
        Name       => 'ReductionMatrix2',
        Writable   => 'rational64s',
        WriteGroup => 'IFD0',
        Count      => -1,
        Protected  => 1,
    },
    0xc627 => {
        Name       => 'AnalogBalance',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Count      => -1,
        Protected  => 1,
    },
    0xc628 => {
        Name       => 'AsShotNeutral',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Count      => -1,
        Protected  => 1,
    },
    0xc629 => {
        Name       => 'AsShotWhiteXY',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Count      => 2,
        Protected  => 1,
    },
    0xc62a => {
        Name       => 'BaselineExposure',
        Writable   => 'rational64s',
        WriteGroup => 'IFD0',
        Protected  => 1,
    },
    0xc62b => {
        Name       => 'BaselineNoise',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Protected  => 1,
    },
    0xc62c => {
        Name       => 'BaselineSharpness',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Protected  => 1,
    },
    0xc62d => {
        Name       => 'BayerGreenSplit',
        Writable   => 'int32u',
        WriteGroup => 'SubIFD',
        Protected  => 1,
    },
    0xc62e => {
        Name       => 'LinearResponseLimit',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Protected  => 1,
    },
    0xc62f => {
        Name       => 'CameraSerialNumber',
        Groups     => { 2 => 'Camera' },
        Writable   => 'string',
        WriteGroup => 'IFD0',
    },
    0xc630 => {
        Name         => 'DNGLensInfo',
        Groups       => { 2 => 'Camera' },
        Writable     => 'rational64u',
        WriteGroup   => 'IFD0',
        Count        => 4,
        PrintConv    => \&PrintLensInfo,
        PrintConvInv => \&ConvertLensInfo,
    },
    0xc631 => {
        Name       => 'ChromaBlurRadius',
        Writable   => 'rational64u',
        WriteGroup => 'SubIFD',
        Protected  => 1,
    },
    0xc632 => {
        Name       => 'AntiAliasStrength',
        Writable   => 'rational64u',
        WriteGroup => 'SubIFD',
        Protected  => 1,
    },
    0xc633 => {
        Name       => 'ShadowScale',
        Writable   => 'rational64u',
        WriteGroup => 'IFD0',
        Protected  => 1,
    },
    0xc634 => {
        Name       => 'DNGPrivateData',
        Flags      => [ 'Binary', 'Protected' ],
        Format     => 'undef',
        Writable   => 'int8u',
        WriteGroup => 'IFD0',
        Notes      => 'Provides a way for camera manufacturers to store private data in the DNG file for use by their own raw converters, and to have that data preserved by programs that edit DNG files.',
    },
    # 0xc634 => [
    #     {
    #         Condition => '$$self{TIFF_TYPE} =~ /^(ARW|SR2)$/',
    #         Name => 'SR2Private',
    #         Groups => { 1 => 'SR2' },
    #         Flags => 'SubIFD',
    #         Format => 'int32u',
    #         # some utilities have problems unless this is int8u format:
    #         # - Adobe Camera Raw 5.3 gives an error
    #         # - Apple Preview 10.5.8 gets the wrong white balance
    #         FixFormat => 'int8u', # (stupid Sony)
    #         WriteGroup => 'IFD0', # (for Validate)
    #         SubDirectory => {
    #             DirName => 'SR2Private',
    #             TagTable => 'Image::ExifTool::Sony::SR2Private',
    #             Start => '$val',
    #         },
    #     },
    #     {
    #         Condition => '$$valPt =~ /^Adobe\0/',
    #         Name => 'DNGAdobeData',
    #         Flags => [ 'Binary', 'Protected' ],
    #         Writable => 'undef', # (writable directory!) (to make it possible to delete this mess)
    #         WriteGroup => 'IFD0',
    #         NestedHtmlDump => 1,
    #         SubDirectory => { TagTable => 'Image::ExifTool::DNG::AdobeData' },
    #         Format => 'undef',  # but written as int8u (change to undef for speed)
    #     },
    #     {
    #         # Pentax/Samsung models that write AOC maker notes in JPG images:
    #         # K-5,K-7,K-m,K-x,K-r,K10D,K20D,K100D,K110D,K200D,K2000,GX10,GX20
    #         # (Note: the following expression also appears in WriteExif.pl)
    #         Condition => q{
    #             $$valPt =~ /^(PENTAX |SAMSUNG)\0/ and
    #             $$self{Model} =~ /\b(K(-[57mrx]|(10|20|100|110|200)D|2000)|GX(10|20))\b/
    #         },
    #         Name => 'MakerNotePentax',
    #         MakerNotes => 1,    # (causes "MakerNotes header" to be identified in HtmlDump output)
    #         Binary => 1,
    #         WriteGroup => 'IFD0', # (for Validate)
    #         # Note: Don't make this block-writable for a few reasons:
    #         # 1) It would be dangerous (possibly confusing Pentax software)
    #         # 2) It is a different format from the JPEG version of MakerNotePentax
    #         # 3) It is converted to JPEG format by RebuildMakerNotes() when copying
    #         SubDirectory => {
    #             TagTable => 'Image::ExifTool::Pentax::Main',
    #             Start => '$valuePtr + 10',
    #             Base => '$start - 10',
    #             ByteOrder => 'Unknown', # easier to do this than read byteorder word
    #         },
    #         Format => 'undef',  # but written as int8u (change to undef for speed)
    #     },
    #     {
    #         # must duplicate the above tag with a different name for more recent
    #         # Pentax models which use the "PENTAX" instead of the "AOC" maker notes
    #         # in JPG images (needed when copying maker notes from DNG to JPG)
    #         Condition => '$$valPt =~ /^(PENTAX |SAMSUNG)\0/',
    #         Name => 'MakerNotePentax5',
    #         MakerNotes => 1,
    #         Binary => 1,
    #         WriteGroup => 'IFD0', # (for Validate)
    #         SubDirectory => {
    #             TagTable => 'Image::ExifTool::Pentax::Main',
    #             Start => '$valuePtr + 10',
    #             Base => '$start - 10',
    #             ByteOrder => 'Unknown',
    #         },
    #         Format => 'undef',
    #     },
    #     {
    #         # Ricoh models such as the GR III
    #         Condition => '$$valPt =~ /^RICOH\0(II|MM)/',
    #         Name => 'MakerNoteRicohPentax',
    #         MakerNotes => 1,
    #         Binary => 1,
    #         WriteGroup => 'IFD0', # (for Validate)
    #         SubDirectory => {
    #             TagTable => 'Image::ExifTool::Pentax::Main',
    #             Start => '$valuePtr + 8',
    #             Base => '$start - 8',
    #             ByteOrder => 'Unknown',
    #         },
    #         Format => 'undef',
    #     },
    #     # the DJI FC2103 writes some interesting stuff here (with sections labelled
    #     # awb_dbg_info, ae_dbg_info, ae_histogram_info, af_dbg_info, hiso, xidiri) - PH
    #     {
    #         Name => 'DNGPrivateData',
    #         Flags => [ 'Binary', 'Protected' ],
    #         Format => 'undef',
    #         Writable => 'int8u',
    #         WriteGroup => 'IFD0',
    #     },
    # ],
    0xc635 => {
        Name => 'MakerNoteSafety',
        Writable => 'int16u',
        WriteGroup => 'IFD0',
        PrintConv => {
            0 => 'Unsafe',
            1 => 'Safe',
        },
    },
    0xc640 => { #15
        Name => 'RawImageSegmentation',
        # (int16u[3], not writable)
        Notes => q{
            used in segmented Canon CR2 images.  3 numbers: 1. Number of segments minus
            one; 2. Pixel width of segments except last; 3. Pixel width of last segment
        },
    },
    0xc65a => {
        Name => 'CalibrationIlluminant1',
        Writable => 'int16u',
        WriteGroup => 'IFD0',
        Protected => 1,
        SeparateTable => 'LightSource',
        PrintConv => \%lightSource,
    },
    0xc65b => {
        Name => 'CalibrationIlluminant2',
        Writable => 'int16u',
        WriteGroup => 'IFD0',
        Protected => 1,
        SeparateTable => 'LightSource',
        PrintConv => \%lightSource,
    },
    0xc65c => {
        Name => 'BestQualityScale',
        Writable => 'rational64u',
        WriteGroup => 'SubIFD',
        Protected => 1,
    },
    0xc65d => {
        Name => 'RawDataUniqueID',
        Format => 'undef',
        Writable => 'int8u',
        WriteGroup => 'IFD0',
        Count => 16,
        Protected => 1,
        ValueConv => 'uc(unpack("H*",$val))',
        ValueConvInv => 'pack("H*", $val)',
    },
    0xc660 => { #3
        Name => 'AliasLayerMetadata',
        Notes => 'used by Alias Sketchbook Pro',
    },
    0xc68b => {
        Name => 'OriginalRawFileName',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc68c => {
        Name => 'OriginalRawFileData', # (writable directory!)
        Writable => 'undef', # must be defined here so tag will be extracted if specified
        WriteGroup => 'IFD0',
        Flags => [ 'Binary', 'Protected' ],
        SubDirectory => {
            TagTable => 'Image::ExifTool::DNG::OriginalRaw',
        },
    },
    0xc68d => {
        Name => 'ActiveArea',
        Writable => 'int32u',
        WriteGroup => 'SubIFD',
        Count => 4,
        Protected => 1,
    },
    0xc68e => {
        Name => 'MaskedAreas',
        Writable => 'int32u',
        WriteGroup => 'SubIFD',
        Count => -1,
        Protected => 1,
    },
    0xc68f => {
        Name => 'AsShotICCProfile', # (writable directory)
        Binary => 1,
        Writable => 'undef', # must be defined here so tag will be extracted if specified
        WriteGroup => 'IFD0',
        Protected => 1,
        WriteCheck => q{
            require Image::ExifTool::ICC_Profile;
            return Image::ExifTool::ICC_Profile::ValidateICC(\$val);
        },
        SubDirectory => {
            DirName => 'AsShotICCProfile',
            TagTable => 'Image::ExifTool::ICC_Profile::Main',
        },
    },
    0xc690 => {
        Name => 'AsShotPreProfileMatrix',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xc691 => {
        Name => 'CurrentICCProfile', # (writable directory)
        Binary => 1,
        Writable => 'undef', # must be defined here so tag will be extracted if specified
        SubDirectory => {
            DirName => 'CurrentICCProfile',
            TagTable => 'Image::ExifTool::ICC_Profile::Main',
        },
        Writable => 'undef',
        WriteGroup => 'IFD0',
        Protected => 1,
        WriteCheck => q{
            require Image::ExifTool::ICC_Profile;
            return Image::ExifTool::ICC_Profile::ValidateICC(\$val);
        },
    },
    0xc692 => {
        Name => 'CurrentPreProfileMatrix',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xc6bf => {
        Name => 'ColorimetricReference',
        Writable => 'int16u',
        WriteGroup => 'IFD0',
        Protected => 1,
    },
    0xc6c5 => { Name => 'SRawType', Description => 'SRaw Type', WriteGroup => 'IFD0' }, #exifprobe (CR2 proprietary)
    0xc6d2 => { #JD (Panasonic DMC-TZ5)
        # this text is UTF-8 encoded (hooray!) - PH (TZ5)
        Name => 'PanasonicTitle',
        Format => 'string', # written incorrectly as 'undef'
        Notes => 'proprietary Panasonic tag used for baby/pet name, etc',
        Writable => 'undef',
        WriteGroup => 'IFD0',
        # panasonic always records this tag (64 zero bytes),
        # so ignore it unless it contains valid information
        RawConv => 'length($val) ? $val : undef',
        ValueConv => '$self->Decode($val, "UTF8")',
        ValueConvInv => '$self->Encode($val,"UTF8")',
    },
    0xc6d3 => { #PH (Panasonic DMC-FS7)
        Name => 'PanasonicTitle2',
        Format => 'string', # written incorrectly as 'undef'
        Notes => 'proprietary Panasonic tag used for baby/pet name with age',
        Writable => 'undef',
        WriteGroup => 'IFD0',
        # panasonic always records this tag (128 zero bytes),
        # so ignore it unless it contains valid information
        RawConv => 'length($val) ? $val : undef',
        ValueConv => '$self->Decode($val, "UTF8")',
        ValueConvInv => '$self->Encode($val,"UTF8")',
    },
    # 0xc6dc - int32u[4]: found in CR2 images (PH, 7DmkIII)
    # 0xc6dd - int16u[256]: found in CR2 images (PH, 5DmkIV)
    0xc6f3 => {
        Name => 'CameraCalibrationSig',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc6f4 => {
        Name => 'ProfileCalibrationSig',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc6f5 => {
        Name => 'ProfileIFD', # (ExtraCameraProfiles)
        Groups => { 1 => 'ProfileIFD' },
        Flags => 'SubIFD',
        WriteGroup => 'IFD0', # (only for Validate)
        SubDirectory => {
            ProcessProc => \&ProcessTiffIFD,
            WriteProc => \&ProcessTiffIFD,
            DirName => 'ProfileIFD',
            Start => '$val',
            Base => '$start',   # offsets relative to start of TIFF-like header
            MaxSubdirs => 10,
            Magic => 0x4352,    # magic number for TIFF-like header
        },
    },
    0xc6f6 => {
        Name => 'AsShotProfileName',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc6f7 => {
        Name => 'NoiseReductionApplied',
        Writable => 'rational64u',
        WriteGroup => 'SubIFD',
        Protected => 1,
    },
    0xc6f8 => {
        Name => 'ProfileName',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc6f9 => {
        Name => 'ProfileHueSatMapDims',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Count => 3,
        Protected => 1,
    },
    0xc6fa => {
        Name => 'ProfileHueSatMapData1',
        %longBin,
        Writable => 'float',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xc6fb => {
        Name => 'ProfileHueSatMapData2',
        %longBin,
        Writable => 'float',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xc6fc => {
        Name => 'ProfileToneCurve',
        %longBin,
        Writable => 'float',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xc6fd => {
        Name => 'ProfileEmbedPolicy',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Protected => 1,
        PrintConv => {
            0 => 'Allow Copying',
            1 => 'Embed if Used',
            2 => 'Never Embed',
            3 => 'No Restrictions',
        },
    },
    0xc6fe => {
        Name => 'ProfileCopyright',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc714 => {
        Name => 'ForwardMatrix1',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xc715 => {
        Name => 'ForwardMatrix2',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xc716 => {
        Name => 'PreviewApplicationName',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc717 => {
        Name => 'PreviewApplicationVersion',
        Writable => 'string',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc718 => {
        Name => 'PreviewSettingsName',
        Writable => 'string',
        WriteGroup => 'IFD0',
        Protected => 1,
        %utf8StringConv,
    },
    0xc719 => {
        Name => 'PreviewSettingsDigest',
        Format => 'undef',
        Writable => 'int8u',
        WriteGroup => 'IFD0',
        Protected => 1,
        ValueConv => 'unpack("H*", $val)',
        ValueConvInv => 'pack("H*", $val)',
    },
    0xc71a => {
        Name => 'PreviewColorSpace',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Protected => 1,
        PrintConv => {
            0 => 'Unknown',
            1 => 'Gray Gamma 2.2',
            2 => 'sRGB',
            3 => 'Adobe RGB',
            4 => 'ProPhoto RGB',
        },
    },
    0xc71b => {
        Name => 'PreviewDateTime',
        Groups => { 2 => 'Time' },
        Writable => 'string',
        Shift => 'Time',
        WriteGroup => 'IFD0',
        Protected => 1,
        ValueConv => q{
            require Image::ExifTool::XMP;
            return Image::ExifTool::XMP::ConvertXMPDate($val);
        },
        ValueConvInv => q{
            require Image::ExifTool::XMP;
            return Image::ExifTool::XMP::FormatXMPDate($val);
        },
        PrintConv => '$self->ConvertDateTime($val)',
        PrintConvInv => '$self->InverseDateTime($val,1,1)',
    },
    0xc71c => {
        Name => 'RawImageDigest',
        Format => 'undef',
        Writable => 'int8u',
        WriteGroup => 'IFD0',
        Count => 16,
        Protected => 1,
        ValueConv => 'unpack("H*", $val)',
        ValueConvInv => 'pack("H*", $val)',
    },
    0xc71d => {
        Name => 'OriginalRawFileDigest',
        Format => 'undef',
        Writable => 'int8u',
        WriteGroup => 'IFD0',
        Count => 16,
        Protected => 1,
        ValueConv => 'unpack("H*", $val)',
        ValueConvInv => 'pack("H*", $val)',
    },
    0xc71e => 'SubTileBlockSize',
    0xc71f => 'RowInterleaveFactor',
    0xc725 => {
        Name => 'ProfileLookTableDims',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Count => 3,
        Protected => 1,
    },
    0xc726 => {
        Name => 'ProfileLookTableData',
        %longBin,
        Writable => 'float',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xc740 => { Name => 'OpcodeList1', %opcodeInfo }, # DNG 1.3
    0xc741 => { Name => 'OpcodeList2', %opcodeInfo }, # DNG 1.3
    0xc74e => { Name => 'OpcodeList3', %opcodeInfo }, # DNG 1.3
    0xc761 => { # DNG 1.3
        Name => 'NoiseProfile',
        Writable => 'double',
        WriteGroup => 'SubIFD',
        Count => -1,
        Protected => 1,
    },
    0xc763 => { #28
        Name => 'TimeCodes',
        Writable => 'int8u',
        WriteGroup => 'IFD0',
        Count => -1, # (8 * number of time codes, max 10)
        ValueConv => q{
            my @a = split ' ', $val;
            my @v;
            push @v, join('.', map { sprintf('%.2x',$_) } splice(@a,0,8)) while @a >= 8;
            join ' ', @v;
        },
        ValueConvInv => q{
            my @a = map hex, split /[. ]+/, $val;
            join ' ', @a;
        },
        # Note: Currently ignore the flags:
        #   byte 0 0x80 - color frame
        #   byte 0 0x40 - drop frame
        #   byte 1 0x80 - field phase
        PrintConv => q{
            my @a = map hex, split /[. ]+/, $val;
            my @v;
            while (@a >= 8) {
                my $str = sprintf("%.2x:%.2x:%.2x.%.2x", $a[3]&0x3f,
                                 $a[2]&0x7f, $a[1]&0x7f, $a[0]&0x3f);
                if ($a[3] & 0x80) { # date+timezone exist if BGF2 is set
                    my $tz = $a[7] & 0x3f;
                    my $bz = sprintf('%.2x', $tz);
                    $bz = 100 if $bz =~ /[a-f]/i; # not BCD
                    if ($bz < 26) {
                        $tz = ($bz < 13 ? 0 : 26) - $bz;
                    } elsif ($bz == 32) {
                        $tz = 12.75;
                    } elsif ($bz >= 28 and $bz <= 31) {
                        $tz = 0;    # UTC
                    } elsif ($bz < 100) {
                        undef $tz;  # undefined or user-defined
                    } elsif ($tz < 0x20) {
                        $tz = (($tz < 0x10 ? 10 : 20) - $tz) - 0.5;
                    } else {
                        $tz = (($tz < 0x30 ? 53 : 63) - $tz) + 0.5;
                    }
                    if ($a[7] & 0x80) { # MJD format (/w UTC time)
                        my ($h,$m,$s,$f) = split /[:.]/, $str;
                        my $jday = sprintf('%x%.2x%.2x', reverse @a[4..6]);
                        $str = ConvertUnixTime(($jday - 40587) * 24 * 3600
                                 + ((($h+$tz) * 60) + $m) * 60 + $s) . ".$f";
                        $str =~ s/^(\d+):(\d+):(\d+) /$1-$2-${3}T/;
                    } else { # YYMMDD (Note: CinemaDNG 1.1 example seems wrong)
                        my $yr = sprintf('%.2x',$a[6]) + 1900;
                        $yr += 100 if $yr < 1970;
                        $str = sprintf('%d-%.2x-%.2xT%s',$yr,$a[5],$a[4],$str);
                    }
                    $str .= TimeZoneString($tz*60) if defined $tz;
                }
                push @v, $str;
                splice @a, 0, 8;
            }
            join ' ', @v;
        },
        PrintConvInv => q{
            my @a = split ' ', $val;
            my @v;
            foreach (@a) {
                my @td = reverse split /T/;
                my $tz = 0x39; # default to unknown timezone
                if ($td[0] =~ s/([-+])(\d+):(\d+)$//) {
                    if ($3 == 0) {
                        $tz = hex(($1 eq '-') ? $2 : 0x26 - $2);
                    } elsif ($3 == 30) {
                        if ($1 eq '-') {
                            $tz = $2 + 0x0a;
                            $tz += 0x0a if $tz > 0x0f;
                        } else {
                            $tz = 0x3f - $2;
                            $tz -= 0x0a if $tz < 0x3a;
                        }
                    } elsif ($3 == 45) {
                        $tz = 0x32 if $1 eq '+' and $2 == 12;
                    }
                }
                my @t = split /[:.]/, $td[0];
                push @t, '00' while @t < 4;
                my $bg;
                if ($td[1]) {
                    # date was specified: fill in date & timezone
                    my @d = split /[-]/, $td[1];
                    next if @d < 3;
                    $bg = sprintf('.%.2d.%.2d.%.2d.%.2x', $d[2], $d[1], $d[0]%100, $tz);
                    $t[0] = sprintf('%.2x', hex($t[0]) + 0xc0); # set BGF1+BGF2
                } else { # time only
                    $bg = '.00.00.00.00';
                }
                push @v, join('.', reverse(@t[0..3])) . $bg;
            }
            join ' ', @v;
        },
    },
    0xc764 => { #28
        Name => 'FrameRate',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        PrintConv => 'int($val * 1000 + 0.5) / 1000',
        PrintConvInv => '$val',
    },
    0xc772 => { #28
        Name => 'TStop',
        Writable => 'rational64u',
        WriteGroup => 'IFD0',
        Count => -1, # (1 or 2)
        PrintConv => 'join("-", map { sprintf("%.2f",$_) } split " ", $val)',
        PrintConvInv => '$val=~tr/-/ /; $val',
    },
    0xc789 => { #28
        Name => 'ReelName',
        Writable => 'string',
        WriteGroup => 'IFD0',
    },
    0xc791 => { # DNG 1.4
        Name => 'OriginalDefaultFinalSize',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Count => 2,
        Protected => 1,
    },
    0xc792 => { # DNG 1.4
        Name => 'OriginalBestQualitySize',
        Notes => 'called OriginalBestQualityFinalSize by the DNG spec',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Count => 2,
        Protected => 1,
    },
    0xc793 => { # DNG 1.4
        Name => 'OriginalDefaultCropSize',
        Writable => 'rational64u',
        WriteGroup => 'IFD0',
        Count => 2,
        Protected => 1,
    },
    0xc7a1 => {  #28
        Name => 'CameraLabel',
        Writable => 'string',
        WriteGroup => 'IFD0',
    },
    0xc7a3 => { # DNG 1.4
        Name => 'ProfileHueSatMapEncoding',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Protected => 1,
        PrintConv => {
            0 => 'Linear',
            1 => 'sRGB',
        },
    },
    0xc7a4 => { # DNG 1.4
        Name => 'ProfileLookTableEncoding',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Protected => 1,
        PrintConv => {
            0 => 'Linear',
            1 => 'sRGB',
        },
    },
    0xc7a5 => { # DNG 1.4
        Name => 'BaselineExposureOffset',
        Writable => 'rational64s', # (incorrectly "RATIONAL" in DNG 1.4 spec)
        WriteGroup => 'IFD0',
        Protected => 1,
    },
    0xc7a6 => { # DNG 1.4
        Name => 'DefaultBlackRender',
        Writable => 'int32u',
        WriteGroup => 'IFD0',
        Protected => 1,
        PrintConv => {
            0 => 'Auto',
            1 => 'None',
        },
    },
    0xc7a7 => { # DNG 1.4
        Name => 'NewRawImageDigest',
        Format => 'undef',
        Writable => 'int8u',
        WriteGroup => 'IFD0',
        Count => 16,
        Protected => 1,
        ValueConv => 'unpack("H*", $val)',
        ValueConvInv => 'pack("H*", $val)',
    },
    0xc7a8 => { # DNG 1.4
        Name => 'RawToPreviewGain',
        Writable => 'double',
        WriteGroup => 'IFD0',
        Protected => 1,
    },
    # 0xc7a9 - CacheBlob (ref 31)
    0xc7aa => { #31 undocumented DNG tag written by LR4 (val=256, related to fast load data?)
        Name => 'CacheVersion',
        Writable => 'int32u',
        WriteGroup => 'SubIFD2',
        Format => 'int8u',
        Count => 4,
        Protected => 1,
        PrintConv => '$val =~ tr/ /./; $val',
        PrintConvInv => '$val =~ tr/./ /; $val',
    },
    0xc7b5 => { # DNG 1.4
        Name => 'DefaultUserCrop',
        Writable => 'rational64u',
        WriteGroup => 'SubIFD',
        Count => 4,
        Protected => 1,
    },
    0xc7d5 => { #PH (in SubIFD1 of Nikon Z6/Z7 NEF images)
        Name => 'NikonNEFInfo',
        Condition => '$$valPt =~ /^Nikon\0/',
        SubDirectory => {
            TagTable => 'Image::ExifTool::Nikon::NEFInfo',
            Start => '$valuePtr + 18',
            Base => '$start - 8',
            ByteOrder => 'Unknown',
        },
    },
    # 0xc7d6 - int8u: 1 (SubIFD1 of Nikon Z6/Z7 NEF)
    0xc7e9 => { # DNG 1.5
        Name => 'DepthFormat',
        Writable => 'int16u',
        Notes => 'tags 0xc7e9-0xc7ee added by DNG 1.5.0.0',
        Protected => 1,
        WriteGroup => 'IFD0',
        PrintConv => {
            0 => 'Unknown',
            1 => 'Linear',
            2 => 'Inverse',
        },
    },
    0xc7ea => { # DNG 1.5
        Name => 'DepthNear',
        Writable => 'rational64u',
        Protected => 1,
        WriteGroup => 'IFD0',
    },
    0xc7eb => { # DNG 1.5
        Name => 'DepthFar',
        Writable => 'rational64u',
        Protected => 1,
        WriteGroup => 'IFD0',
    },
    0xc7ec => { # DNG 1.5
        Name => 'DepthUnits',
        Writable => 'int16u',
        Protected => 1,
        WriteGroup => 'IFD0',
        PrintConv => {
            0 => 'Unknown',
            1 => 'Meters',
        },
    },
    0xc7ed => { # DNG 1.5
        Name => 'DepthMeasureType',
        Writable => 'int16u',
        Protected => 1,
        WriteGroup => 'IFD0',
        PrintConv => {
            0 => 'Unknown',
            1 => 'Optical Axis',
            2 => 'Optical Ray',
        },
    },
    0xc7ee => { # DNG 1.5
        Name => 'EnhanceParams',
        Writable => 'string',
        Protected => 1,
        WriteGroup => 'IFD0',
    },
    0xcd2d => { # DNG 1.6
        Name => 'ProfileGainTableMap',
        Writable => 'undef',
        WriteGroup => 'SubIFD',
        Protected => 1,
        Binary => 1,
    },
    0xcd2e => { # DNG 1.6
        Name => 'SemanticName',
        # Writable => 'string',
        WriteGroup => 'SubIFD' #? (NC) Semantic Mask IFD (only for Validate)
    },
    0xcd30 => { # DNG 1.6
        Name => 'SemanticInstanceIFD',
        # Writable => 'string',
        WriteGroup => 'SubIFD' #? (NC) Semantic Mask IFD (only for Validate)
    },
    0xcd31 => { # DNG 1.6
        Name => 'CalibrationIlluminant3',
        Writable => 'int16u',
        WriteGroup => 'IFD0',
        Protected => 1,
        SeparateTable => 'LightSource',
        PrintConv => \%lightSource,
    },
    0xcd32 => { # DNG 1.6
        Name => 'CameraCalibration3',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xcd33 => { # DNG 1.6
        Name => 'ColorMatrix3',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xcd34 => { # DNG 1.6
        Name => 'ForwardMatrix3',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xcd35 => { # DNG 1.6
        Name => 'IlluminantData1',
        Writable => 'undef',
        WriteGroup => 'IFD0',
        Protected => 1,
    },
    0xcd36 => { # DNG 1.6
        Name => 'IlluminantData2',
        Writable => 'undef',
        WriteGroup => 'IFD0',
        Protected => 1,
    },
    0xcd37 => { # DNG 1.6
        Name => 'IlluminantData3',
        Writable => 'undef',
        WriteGroup => 'IFD0',
        Protected => 1,
    },
    0xcd38 => { # DNG 1.6
        Name => 'MaskSubArea',
        # Writable => 'int32u',
        WriteGroup => 'SubIFD', #? (NC) Semantic Mask IFD (only for Validate)
        Count => 4,
    },
    0xcd39 => { # DNG 1.6
        Name => 'ProfileHueSatMapData3',
        %longBin,
        Writable => 'float',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xcd3a => { # DNG 1.6
        Name => 'ReductionMatrix3',
        Writable => 'rational64s',
        WriteGroup => 'IFD0',
        Count => -1,
        Protected => 1,
    },
    0xcd3b => { # DNG 1.6
        Name => 'RGBTables',
        Writable => 'undef',
        WriteGroup => 'IFD0',
        Protected => 1,
    },
    0xea1c => { #13
        Name => 'Padding',
        Binary => 1,
        Protected => 1,
        Writable => 'undef',
        # must start with 0x1c 0xea by the WM Photo specification
        # (not sure what should happen if padding is only 1 byte)
        # (why does MicrosoftPhoto write "1c ea 00 00 00 08"?)
        RawConvInv => '$val=~s/^../\x1c\xea/s; $val',
    },
    0xea1d => {
        Name => 'OffsetSchema',
        Notes => "Microsoft's ill-conceived maker note offset difference",
        Protected => 1,
        Writable => 'int32s',
        # From the Microsoft documentation:
        #
        #     Any time the "Maker Note" is relocated by Windows, the Exif MakerNote
        #     tag (37500) is updated automatically to reference the new location. In
        #     addition, Windows records the offset (or difference) between the old and
        #     new locations in the Exif OffsetSchema tag (59933). If the "Maker Note"
        #     contains relative references, the developer can add the value in
        #     OffsetSchema to the original references to find the correct information.
        #
        # My recommendation is for other developers to ignore this tag because the
        # information it contains is unreliable. It will be wrong if the image has
        # been subsequently edited by another application that doesn't recognize the
        # new Microsoft tag.
        #
        # The new tag unfortunately only gives the difference between the new maker
        # note offset and the original offset. Instead, it should have been designed
        # to store the original offset. The new offset may change if the image is
        # edited, which will invalidate the tag as currently written. If instead the
        # original offset had been stored, the new difference could be easily
        # calculated because the new maker note offset is known.
        #
        # I exchanged emails with a Microsoft technical representative, pointing out
        # this problem shortly after they released the update (Feb 2007), but so far
        # they have taken no steps to address this.
    },
    # 0xefee - int16u: 0 - seen this from a WIC-scanned image

    # tags in the range 0xfde8-0xfe58 have been observed in PS7 files
    # generated from RAW images.  They are all strings with the
    # tag name at the start of the string.  To accommodate these types
    # of tags, all tags with values above 0xf000 are handled specially
    # by ProcessExif().
    0xfde8 => {
        Name => 'OwnerName',
        Condition => '$$self{TIFF_TYPE} ne "DCR"', # (used for another purpose in Kodak DCR images)
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Owner's Name: $val"},
        Notes => q{
            tags 0xfde8-0xfdea and 0xfe4c-0xfe58 are generated by Photoshop Camera RAW.
            Some names are the same as other EXIF tags, but ExifTool will avoid writing
            these unless they already exist in the file
        },
    },
    0xfde9 => {
        Name => 'SerialNumber',
        Condition => '$$self{TIFF_TYPE} ne "DCR"', # (used for another purpose in Kodak DCR SubIFD)
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Serial Number: $val"},
    },
    0xfdea => {
        Name => 'Lens',
        Condition => '$$self{TIFF_TYPE} ne "DCR"', # (used for another purpose in Kodak DCR SubIFD)
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Lens: $val"},
    },
    0xfe4c => {
        Name => 'RawFile',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Raw File: $val"},
    },
    0xfe4d => {
        Name => 'Converter',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Converter: $val"},
    },
    0xfe4e => {
        Name => 'WhiteBalance',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"White Balance: $val"},
    },
    0xfe51 => {
        Name => 'Exposure',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Exposure: $val"},
    },
    0xfe52 => {
        Name => 'Shadows',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Shadows: $val"},
    },
    0xfe53 => {
        Name => 'Brightness',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Brightness: $val"},
    },
    0xfe54 => {
        Name => 'Contrast',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Contrast: $val"},
    },
    0xfe55 => {
        Name => 'Saturation',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Saturation: $val"},
    },
    0xfe56 => {
        Name => 'Sharpness',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Sharpness: $val"},
    },
    0xfe57 => {
        Name => 'Smoothness',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Smoothness: $val"},
    },
    0xfe58 => {
        Name => 'MoireFilter',
        Avoid => 1,
        PSRaw => 1,
        Writable => 'string',
        ValueConv => '$val=~s/^.*: //;$val',
        ValueConvInv => q{"Moire Filter: $val"},
    },

    #-------------
    0xfe00 => {
        Name => 'KDC_IFD',
        Groups => { 1 => 'KDC_IFD' },
        Flags => 'SubIFD',
        Notes => 'used in some Kodak KDC images',
        SubDirectory => {
            TagTable => 'Image::ExifTool::Kodak::KDC_IFD',
            DirName => 'KDC_IFD',
            Start => '$val',
        },
    },
);

my %coordConv = (
    ValueConv    => 'Image::ExifTool::GPS::ToDegrees($val)',
    ValueConvInv => 'Image::ExifTool::GPS::ToDMS($self, $val)',
    PrintConv    => 'Image::ExifTool::GPS::ToDMS($self, $val, 1)',
);

my %GpsTable = (
    # GROUPS => { 0 => 'EXIF', 1 => 'GPS', 2 => 'Location' },
    # WRITE_PROC => \&Image::ExifTool::Exif::WriteExif,
    # CHECK_PROC => \&Image::ExifTool::Exif::CheckExif,
    # WRITABLE => 1,
    # WRITE_GROUP => 'GPS',
    0x0000 => {
        Name => 'GPSVersionID',
        Writable => 'int8u',
        Mandatory => 1,
        Count => 4,
        PrintConv => '$val =~ tr/ /./; $val',
        PrintConvInv => '$val =~ tr/./ /; $val',
    },
    0x0001 => {
        Name => 'GPSLatitudeRef',
        Writable => 'string',
        Notes => q{
            tags 0x0001-0x0006 used for camera location according to MWG 2.0. ExifTool
            will also accept a number when writing GPSLatitudeRef, positive for north
            latitudes or negative for south, or a string containing N, North, S or South
        },
        Count => 2,
        PrintConv => {
            # extract N/S if written from Composite:GPSLatitude
            # (also allow writing from a signed number)
            OTHER => sub {
                my ($val, $inv) = @_;
                return undef unless $inv;
                return uc $2 if $val =~ /(^|[^A-Z])([NS])(orth|outh)?\b/i;
                return $1 eq '-' ? 'S' : 'N' if $val =~ /([-+]?)\d+/;
                return undef;
            },
            N => 'North',
            S => 'South',
        },
    },
    0x0002 => {
        Name => 'GPSLatitude',
        Writable => 'rational64u',
        Count => 3,
        %coordConv,
        PrintConvInv => 'Image::ExifTool::GPS::ToDegrees($val,undef,"lat")',
    },
    0x0003 => {
        Name => 'GPSLongitudeRef',
        Writable => 'string',
        Count => 2,
        Notes => q{
            ExifTool will also accept a number when writing this tag, positive for east
            longitudes or negative for west, or a string containing E, East, W or West
        },
        PrintConv => {
            # extract E/W if written from Composite:GPSLongitude
            # (also allow writing from a signed number)
            OTHER => sub {
                my ($val, $inv) = @_;
                return undef unless $inv;
                return uc $2 if $val =~ /(^|[^A-Z])([EW])(ast|est)?\b/i;
                return $1 eq '-' ? 'W' : 'E' if $val =~ /([-+]?)\d+/;
                return undef;
            },
            E => 'East',
            W => 'West',
        },
    },
    0x0004 => {
        Name => 'GPSLongitude',
        Writable => 'rational64u',
        Count => 3,
        %coordConv,
        PrintConvInv => 'Image::ExifTool::GPS::ToDegrees($val,undef,"lon")',
    },
    0x0005 => {
        Name => 'GPSAltitudeRef',
        Writable => 'int8u',
        Notes => q{
            ExifTool will also accept number when writing this tag, with negative
            numbers indicating below sea level
        },
        PrintConv => {
            OTHER => sub {
                my ($val, $inv) = @_;
                return undef unless $inv and $val =~ /^([-+0-9])/;
                return($1 eq '-' ? 1 : 0);
            },
            0 => 'Above Sea Level',
            1 => 'Below Sea Level',
        },
    },
    0x0006 => {
        Name => 'GPSAltitude',
        Writable => 'rational64u',
        # extricate unsigned decimal number from string
        ValueConvInv => '$val=~/((?=\d|\.\d)\d*(?:\.\d*)?)/ ? $1 : undef',
        PrintConv => '$val =~ /^(inf|undef)$/ ? $val : "$val m"',
        PrintConvInv => '$val=~s/\s*m$//;$val',
    },
    0x0007 => {
        Name => 'GPSTimeStamp',
        Groups => { 2 => 'Time' },
        Writable => 'rational64u',
        Count => 3,
        Shift => 'Time',
        Notes => q{
            UTC time of GPS fix.  When writing, date is stripped off if present, and
            time is adjusted to UTC if it includes a timezone
        },
        ValueConv => 'Image::ExifTool::GPS::ConvertTimeStamp($val)',
        ValueConvInv => '$val=~tr/:/ /;$val',
        PrintConv => 'Image::ExifTool::GPS::PrintTimeStamp($val)',
        # pull time out of any format date/time string
        # (converting to UTC if a timezone is given)
        PrintConvInv => sub {
            my ($v, $et) = @_;
            $v = $et->TimeNow() if lc($v) eq 'now';
            my @tz;
            if ($v =~ s/([-+])(\d{1,2}):?(\d{2})\s*(DST)?$//i) {    # remove timezone
                my $s = $1 eq '-' ? 1 : -1; # opposite sign to convert back to UTC
                my $t = $2;
                @tz = ($s*$2, $s*$3);
            }
            # (note: we must allow '.' as a time separator, eg. '10.30.00', with is tricky due to decimal seconds)
            # YYYYmmddHHMMSS[.ss] format
            my @a = ($v =~ /^[^\d]*\d{4}[^\d]*\d{1,2}[^\d]*\d{1,2}[^\d]*(\d{1,2})[^\d]*(\d{2})[^\d]*(\d{2}(?:\.\d+)?)[^\d]*$/);
            # HHMMSS[.ss] format
            @a or @a = ($v =~ /^[^\d]*(\d{1,2})[^\d]*(\d{2})[^\d]*(\d{2}(?:\.\d+)?)[^\d]*$/);
            @a or warn('Invalid time (use HH:MM:SS[.ss][+/-HH:MM|Z])'), return undef;
            if (@tz) {
                # adjust to UTC
                $a[1] += $tz[1];
                $a[0] += $tz[0];
                while ($a[1] >= 60) { $a[1] -= 60; ++$a[0] }
                while ($a[1] < 0)   { $a[1] += 60; --$a[0] }
                $a[0] = ($a[0] + 24) % 24;
            }
            return join(':', @a);
        },
    },
    0x0008 => {
        Name => 'GPSSatellites',
        Writable => 'string',
    },
    0x0009 => {
        Name => 'GPSStatus',
        Writable => 'string',
        Count => 2,
        PrintConv => {
            A => 'Measurement Active', # Exif2.2 "Measurement in progress"
            V => 'Measurement Void',   # Exif2.2 "Measurement Interoperability" (WTF?)
            # (meaning for 'V' taken from status code in NMEA GLL and RMC sentences)
        },
    },
    0x000a => {
        Name => 'GPSMeasureMode',
        Writable => 'string',
        Count => 2,
        PrintConv => {
            2 => '2-Dimensional Measurement',
            3 => '3-Dimensional Measurement',
        },
    },
    0x000b => {
        Name => 'GPSDOP',
        Description => 'GPS Dilution Of Precision',
        Writable => 'rational64u',
    },
    0x000c => {
        Name => 'GPSSpeedRef',
        Writable => 'string',
        Count => 2,
        PrintConv => {
            K => 'km/h',
            M => 'mph',
            N => 'knots',
        },
    },
    0x000d => {
        Name => 'GPSSpeed',
        Writable => 'rational64u',
    },
    0x000e => {
        Name => 'GPSTrackRef',
        Writable => 'string',
        Count => 2,
        PrintConv => {
            M => 'Magnetic North',
            T => 'True North',
        },
    },
    0x000f => {
        Name => 'GPSTrack',
        Writable => 'rational64u',
    },
    0x0010 => {
        Name => 'GPSImgDirectionRef',
        Writable => 'string',
        Count => 2,
        PrintConv => {
            M => 'Magnetic North',
            T => 'True North',
        },
    },
    0x0011 => {
        Name => 'GPSImgDirection',
        Writable => 'rational64u',
    },
    0x0012 => {
        Name => 'GPSMapDatum',
        Writable => 'string',
    },
    0x0013 => {
        Name => 'GPSDestLatitudeRef',
        Writable => 'string',
        Notes => 'tags 0x0013-0x001a used for subject location according to MWG 2.0',
        Count => 2,
        PrintConv => { N => 'North', S => 'South' },
    },
    0x0014 => {
        Name => 'GPSDestLatitude',
        Writable => 'rational64u',
        Count => 3,
        %coordConv,
        PrintConvInv => 'Image::ExifTool::GPS::ToDegrees($val,undef,"lat")',
    },
    0x0015 => {
        Name => 'GPSDestLongitudeRef',
        Writable => 'string',
        Count => 2,
        PrintConv => { E => 'East', W => 'West' },
    },
    0x0016 => {
        Name => 'GPSDestLongitude',
        Writable => 'rational64u',
        Count => 3,
        %coordConv,
        PrintConvInv => 'Image::ExifTool::GPS::ToDegrees($val,undef,"lon")',
    },
    0x0017 => {
        Name => 'GPSDestBearingRef',
        Writable => 'string',
        Count => 2,
        PrintConv => {
            M => 'Magnetic North',
            T => 'True North',
        },
    },
    0x0018 => {
        Name => 'GPSDestBearing',
        Writable => 'rational64u',
    },
    0x0019 => {
        Name => 'GPSDestDistanceRef',
        Writable => 'string',
        Count => 2,
        PrintConv => {
            K => 'Kilometers',
            M => 'Miles',
            N => 'Nautical Miles',
        },
    },
    0x001a => {
        Name => 'GPSDestDistance',
        Writable => 'rational64u',
    },
    0x001b => {
        Name => 'GPSProcessingMethod',
        Writable => 'undef',
        Notes => 'values of "GPS", "CELLID", "WLAN" or "MANUAL" by the EXIF spec.',
        RawConv => 'Image::ExifTool::Exif::ConvertExifText($self,$val,1,$tag)',
        RawConvInv => 'Image::ExifTool::Exif::EncodeExifText($self,$val)',
    },
    0x001c => {
        Name => 'GPSAreaInformation',
        Writable => 'undef',
        RawConv => 'Image::ExifTool::Exif::ConvertExifText($self,$val,1,$tag)',
        RawConvInv => 'Image::ExifTool::Exif::EncodeExifText($self,$val)',
    },
    0x001d => {
        Name => 'GPSDateStamp',
        Groups => { 2 => 'Time' },
        Writable => 'string',
        Format => 'undef', # (Casio EX-H20G uses "\0" instead of ":" as a separator)
        Count => 11,
        Shift => 'Time',
        Notes => q{
            when writing, time is stripped off if present, after adjusting date/time to
            UTC if time includes a timezone.  Format is YYYY:mm:dd
        },
        RawConv => '$val =~ s/\0+$//; $val',
        ValueConv => 'Image::ExifTool::Exif::ExifDate($val)',
        ValueConvInv => '$val',
        # pull date out of any format date/time string
        # (and adjust to UTC if this is a full date/time/timezone value)
        PrintConvInv => q{
            my $secs;
            $val = $self->TimeNow() if lc($val) eq 'now';
            if ($val =~ /[-+]/ and ($secs = Image::ExifTool::GetUnixTime($val, 1))) {
                $val = Image::ExifTool::ConvertUnixTime($secs);
            }
            return $val =~ /(\d{4}).*?(\d{2}).*?(\d{2})/ ? "$1:$2:$3" : undef;
        },
    },
    0x001e => {
        Name => 'GPSDifferential',
        Writable => 'int16u',
        PrintConv => {
            0 => 'No Correction',
            1 => 'Differential Corrected',
        },
    },
    0x001f => {
        Name => 'GPSHPositioningError',
        Description => 'GPS Horizontal Positioning Error',
        PrintConv => '"$val m"',
        PrintConvInv => '$val=~s/\s*m$//; $val',
        Writable => 'rational64u',
    },
    # 0xea1c - Nokia Lumina 1020, Samsung GT-I8750, and other Windows 8
    #          phones write this (padding) in GPS IFD - PH
);


sub parseTag {
    my %tag = %{$_[0]};
    my $name = $tag{'Name'};
    my $ifd = exists $tag{'WriteGroup'} ? $tag{'WriteGroup'} : 'ExifIFD';
    my $writable = exists $tag{'Writable'} ? $tag{'Writable'} : 'undef';
    #if(!exists $tag{'Writable'} && !exists $tag{'WriteGroup'}) {
    #    $ifd = 'UndefIFD';
    #}
    my $fmt = exists $tag{'Format'} ? $tag{'Format'} : 'undef';
    my $permanent = exists $tag{'Permanent'} ? $tag{'Permanent'} + 0: 0;
    my $mandatory = exists $tag{'Mandatory'} ? $tag{'Mandatory'} + 0: 0;
    my $protected = exists $tag{'Protected'} ? 1 : 0;
    my $offset = exists $tag{'IsOffset'} ? $tag{'IsOffset'} + 0: 0;
    my $offsetPair = exists $tag{'OffsetPair'} ? $tag{'OffsetPair'} + 0: 0;
    if ($offsetPair < 0) {
        $offsetPair = 0
    }
    my $count = exists $tag{'Count'} ? $tag{'Count'} + 0 : 0;
    my $notes = exists $tag{'Notes'} ? $tag{'Notes'} : '';
    my $description = exists $tag{'Description'} ? $tag{'Description'} : '';
    my $values = {};
    if (exists $tag{'PrintConv'}) {
        my $pc = $tag{'PrintConv'};
        if (UNIVERSAL::isa($pc, 'HASH')) {
            for my $vk (keys %$pc) {
                #say "$vk"
                $values->{$vk} = $pc->{$vk}.""
            }
            #$values = $pc
        }
    }
    my $dirName = '';
    my $subDir = 0;
    if (exists $tag{'SubDirectory'}) {
        $subDir = 1;
        my $sd = $tag{'SubDirectory'};
        if (exists $sd->{'DirName'}) {
            $dirName = $sd->{'DirName'}
        }
    }
    my $rec = {
        name        => $name,
        ifd         => $ifd,
        writable    => $writable,
        fmt         => $fmt,
        mandatory   => \$mandatory,
        permanent   => \$permanent,
        protected   => \$protected,
        offset      => \$offset,
        offsetPair  => $offsetPair,
        count       => $count,
        subDir => \$subDir,
        dirName => $dirName,
        notes       => $notes,
        description => $description,
        values      => $values,
    };
    return $rec;
    #say "$name $ifd $writable $fmt $permanent $protected $offset $offsetPair $count"


}



#parse main (ifd, ifdexif, interop tags
sub parseMainTags {
    my %mainTags;
    #my $cnt = 0;
    for my $tagId (keys %ExifTable) {
        my $tag = $ExifTable{$tagId};
        if (UNIVERSAL::isa($tag, 'HASH')) {
            my $rec = parseTag($tag);
            $rec->{'id'} = $tagId + 0;
            $mainTags{$tagId + 0} = $rec;

            # if (exists($tag->{'Writable'}) || exists($tag->{'WriteGroup'})) {
            #
            # } else {
            #     say $tag->{'Name'}
            # }
            #say $rec->{'permanent'};
        } elsif (UNIVERSAL::isa($tag, 'ARRAY')) { #make sure no arrays are left
        } else { #some tags are just string names (no Ifd, etc. skip those
            #say "$tagId is $tag";
        }
    }
    my @vals = values %mainTags;
    return \@vals;
}

#parse gps subifd tags
sub parseGpsTags {
    my %gpsTags;
    for my $tagId (keys %GpsTable) {
        my $tag = $GpsTable{$tagId};
        my $newTag = parseTag($tag);
        $newTag->{'id'} = $tagId+0;
        $newTag->{'ifd'} = 'GpsIFD';
        $gpsTags{$tagId} = $newTag;
    }
    my @vals = values %gpsTags;
    return \@vals;

}

my %allTags;
$allTags{'main'} = parseMainTags();
$allTags{'gps'} = parseGpsTags();

my $jsonText = to_json(\%allTags, { utf8 => 1, pretty => 1, canonical => 1 });
say "$jsonText";






