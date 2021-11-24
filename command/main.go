package main

import (
	"fmt"
	"github.com/msvens/mimage/img"
	"github.com/msvens/mimage/metadata"
	"regexp"
	"strings"
	"time"
)

func main() {

	//tryEditTags()
	//tryCreateMetaData()
	//tryMetaDataCopy()
	//tryImageTransform()
	//tryImageDesc()
	//tryUserComment()
	//trySegments()
	tryMetaData()

	//play with errors:

}

func tryCreateMetaData() {
	old := "./assets/screendump.jpg"
	dest := "/Users/mellowtech/imgtest/screendumpNewExif.jpg"
	var mde *metadata.MetaDataEditor
	var err error
	if mde, err = metadata.NewMetaDataEditorFile(old); err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now()
	mde.SetExifDate(metadata.ModifyDate, t)
	mde.SetExifDate(metadata.OriginalDate, t)
	mde.SetExifDate(metadata.DigitizedDate, t)
	err = mde.SetExifTag(metadata.Exif_Flash, uint16(0x10))
	mde.SetImageDescription("Screendump")
	err = mde.WriteFile(dest)
	if err != nil {
		fmt.Println(err)
	}
}

func tryMetaData() {
	fname := "./assets/xe3.jpg"

	md, err := metadata.ParseFile(fname)
	if err != nil {
		fmt.Println(err.Error())
	}
	if md != nil {
		fmt.Println(md.Summary.Title)
		//fmt.Println(md.PrintIfd())
		//fmt.Println(md.PrintIptc())
	}
	/*if md.HasXmp() {
		b, _ := xmp.MarshalIndent(md.Xmp(), "", "  ")
		fmt.Println(string(b))
	}*/
}

func tryEditTags() {
	source := "./assets/Leica.jpg"
	dest := "/Users/mellowtech/imgtest/LeicaEdited.jpg"
	var mde *metadata.MetaDataEditor
	var err error
	if mde, err = metadata.NewMetaDataEditorFile(source); err != nil {
		fmt.Println(err)
		return
	}
	mde.SetUserComment("Some new user comment")
	mde.SetImageDescription("Some new image description")
	mde.SetExifDate(metadata.ModifyDate, time.Now())
	if err = mde.WriteFile(dest); err != nil {
		fmt.Println(err)
	}
}

func tryMetaDataCopy() {
	source := "./assets/Leica.jpg"
	old := "./assets/screendump.jpg"
	dest := "/Users/mellowtech/imgtest/screendumpExif.jpg"
	var mde *metadata.MetaDataEditor
	var err error
	if mde, err = metadata.NewMetaDataEditorFile(old); err != nil {
		fmt.Println(err)
		return
	}
	if err = mde.CopyMetaDataFile(source, metadata.CopyAll); err != nil {
		fmt.Println(err)
	}
	mde.PrintSegments()
	fmt.Println()
	bytes, err := mde.Bytes()
	mdeNew, _ := metadata.NewMetaDataEditor(bytes)
	mdeNew.PrintSegments()

	if err = mde.WriteFile(dest); err != nil {
		fmt.Println(err)
	}
}

func tryImageTransform() {

	fname := "./assets/Leica.jpg"

	oDir := "/Users/mellowtech/imgtest/"

	tests := map[string]img.Options{
		oDir + "400x400.jpg":   {Width: 400, Height: 400, Quality: 90, Transform: img.ResizeAndCrop},
		oDir + "1200x628.jpg":  {Width: 1200, Height: 628, Quality: 90, Transform: img.ResizeAndCrop, CopyExif: true},
		oDir + "1200x1200.jpg": {Width: 1200, Height: 1200, Quality: 90, Transform: img.ResizeAndCrop, CopyExif: true},
		oDir + "1080x1350.jpg": {Width: 1080, Height: 1350, Quality: 90, Transform: img.ResizeAndCrop, CopyExif: true},
		oDir + "1200.jpg":      {Width: 1200, Height: 0, Quality: 90, Transform: img.Resize, CopyExif: true},
	}
	err := img.TransformFile(fname, tests)

	if err != nil {
		fmt.Println(err)
	}
}

func tryRegExp() {
	digit1 := "   uint16![,3   "
	string1 := "strIng  [8]"
	strings.TrimSpace(digit1)

	r, e := regexp.Compile(`^\s*(\w+).*?\[?(\d*),?(\d*)\]?\s*$`)
	if e != nil {
		fmt.Println(e)
	}

	m := r.FindStringSubmatch(string1)
	for i, s := range m {
		fmt.Printf("%v:%s\n", i, s)
	}

}
