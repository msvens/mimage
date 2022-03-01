[![Build Status](https://travis-ci.com/msvens/mimage.svg?branch=master)](https://travis-ci.com/msvens/mimage)
[![codecov](https://codecov.io/gh/dsoprea/go-exif/branch/master/graph/badge.svg)](https://codecov.io/gh/msvens/mimage)
[![Go Report Card](https://goreportcard.com/badge/github.com/msvens/mimage)](https://goreportcard.com/report/github.com/msvens/mimage)
[![GoDoc](https://godoc.org/github.com/msvens/mimage/v3?status.svg)](https://godoc.org/github.com/msvens/mimage)

# Overiew

mimage is a native go package for handling 
image meta information (exif, iptc, xmp) as well as some basic
image manipulation (resizing, thumbnail creation, etc)

mimage is used by [mphotos](https://www.github.com/msvens/mphotos) and we welcome others to try it out.

# Installation
Installing mimage is easy as it only has go native dependencies.

    go get -u github.com/msvens/mimage

Next include

```go
import "github.com/msvens/mimage/metadata"
```
if you want to access and edit image metadata, or

```go
import "github.com/msvens/mimage/img"
```
if you want to manipulate the actual image (crop,resize,thumbnails)

# Usage

mimage has 3 main use cases: reading and writing metadata (exif, iptc, xmp) and "resizing" images.

**Note**: For now this package assumes jpeg images and will not work with other formats

## Accessing Metadata

In its simplest form you instantiate a new MetaData struct from an image file or byte slice. This will
parse the file and creat the corresponding IPTC, XMP and Exif Metadata as well as a summary with
the most "frequently" used fields

```go
md, err := NewMetaDataFromFile("../assets/leica.jpg")
if err != nil {
	fmt.Printf("Could not open file: %v\n", err)
	return
}
fmt.Printf("Make: %s, Model: %s\n", md.Summary().CameraMake, md.Summary().CameraModel)
```
You can always retrieve any metadata by scan that specific field from either IPTC, Xmp or Exif respectively.

```go
cameraMake := ""
err = md.Exif().ScanIfdRoot(IFD_Make, &cameraMake)
if err != nil {
	fmt.Printf("Could not scan ifd tag: %v\n", err)
	return
}
fmt.Printf("Make: %s\n", cameraMake)
```

## Editing Metadata

One can also edit (or add) metadata to an image. All known IPTC,Exif,and Xmp tags can be edited - 
much in the same why as you read those fields. There are conveniance methods to edit image titles
and keywords (will set the fields in all relevant metadata sections) as those are typically 
not generated by your camera/image software

```go
je, err := NewJpegEditorFile("../assets/leica.jpg")
if err != nil {
	fmt.Printf("Could not retrieve editor for file: %v\n", err)
	return
}
err = je.SetTitle("some new title")
if err != nil {
	fmt.Printf("Could not set title: %v\n", err)
	return
}
md, err := je.MetaData()
if err != nil {
	fmt.Printf("Could not get metadata: %v\n", err)
}
fmt.Printf("New Title: %v\n",md.Summary().Title)
```

In order to **persist your editing** you need to call any of these methods on your editor
```go
je.MetaData() //Calls je.Bytes() then instantiates a new MetaData struct
je.Bytes() //Get []byte of the entire image
je.WriteFile(someFile) //Calls je.Bytes() then writes to someFile
```

## Copy and Manipulating Images



## Why?

[mphotos](https://www.github.com/msvens/mphotos) has relied on two libraries/tools for
image manipulation and meta data extraction: [bimg](https://github.com/h2non/bimg) and 
[exiftool](https://exiftool.org/). Both of these are excellent tools but made the compile/deployment
process more complicated as bimg relies on libvips and exiftool is an external program that
needs to be installed. In effect making mphotos slighly less portable.

mimage seeks to remedy this by offering similar functionality using only go native code



# Releases

# Todo
- Add finer grained control over editing (only accept allowed values, etc)
- Other image formats

# Thanks

A big thanks to
- [exif-go](https://github.com/dsoprea/go-exif) for go native support of parsing and editing exif data
- [go-xmp](https://github.com/trimmer-io/go-xmp) for go native support of reading and editing XMP information
- [Imaging](https://github.com/disintegration/imaging) for go native support of image manipulation
- [exiftool](https://github.com/exiftool/exiftool) for providing the sources of all exif and iptc tag information



