[![](https://img.shields.io/github/workflow/status/msvens/mimage/Test?longCache=tru&label=Test&logo=github%20actions&logoColor=fff)](https://github.com/msvens/mimage/actions?query=workflow%3ATest)
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
The second function of image is to manipulate/transform images - typically to create thumbnails, portrait,
landscape and other scaled versions of your original image. These copy functions will also respect and copy
and original image metadata information - which the standard jpeg writers dont do.

The package *"github.com/msvens/mimage/img"* exposes one function
```go
func TransformFile(source string, destinations map[string]Options) error 
```

In the follwoing example we use TransformFile to create a number of versions of our sourceImage.

```go
sourceImg := "../assets/leica.jpg"
homeDir, _ := os.UserHomeDir()
sourceDir := path.Join(homeDir,"transform")
_ = os.Mkdir(sourceDir, 0755)

//for all but the thumb we are copying the original meta information
thumb := NewOptions(ResizeAndCrop, 400, 400, false)
landscape := NewOptions(ResizeAndCrop, 1200, 628, true)
square := NewOptions(ResizeAndCrop, 1200, 1200, true)
portrait := NewOptions(ResizeAndCrop, 1080, 1350, true)
resize := NewOptions(Resize, 1200, 0, true)

destImgs := map[string]Options{
	path.Join(sourceDir,"thumb.jpg"): thumb,
	path.Join(sourceDir, "landscape.jpg"): landscape,
	path.Join(sourceDir, "square.jpg"): square,
	path.Join(sourceDir, "portrait.jpg"): portrait,
	path.Join(sourceDir, "resize.jpg"): resize,
}

_ = TransformFile(sourceImg, destImgs)
```

Note that in a real situation you would handle the errors that we are now just skipping


## Motivation

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



