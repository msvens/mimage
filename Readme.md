# Overiew

mimage is a native go package for handling 
image meta information (exif, iptc, xmp) as well as some basic
image manipulation (resizing, thumbnail creation, etc)

mimage is used by [mphotos](https://www.github.com/msvens/mphotos) and is not at
this time intended to be a general purpose package - it lacks proper testing and functionality.

## Why?

[mphotos](https://www.github.com/msvens/mphotos) has relied on two libraries/tools for
image manipulation and meta data extraction: [bimg](https://github.com/h2non/bimg) and 
[exiftool](https://exiftool.org/). Both of these are excellent tools but made the compile/deployment
process more complicated as bimg relies on libvips and exiftool is an external program that
needs to be installed. In effect making mphotos slighly less portable.

mimage seeks to remedy this by offering similar functionality using only go native code

# Usage

# Releases

# Todo
- IPTC editing

# Thanks

A big thanks to
- [exif-go](https://github.com/dsoprea/go-exif) for go native support of parsing and editing exif data
- [go-xmp](https://github.com/trimmer-io/go-xmp) for go native support of reading and editing XMP information
- [Imaging](https://github.com/disintegration/imaging) for go native support of image manipulation
- [exiftool](https://github.com/exiftool/exiftool) for providing the sources of all exif and iptc tag information



