package main

import (
	"fmt"
	"github.com/msvens/mimage/metadata"
	"regexp"
	"strings"
)

func main() {

	fname := "./assets/Leica.jpg"

	md, err := metadata.NewMetaDataFromJpegFile(fname)
	if err != nil {
		fmt.Println(err.Error())
	}
	if md != nil {
		fmt.Println(md.Summary)
		fmt.Println(md.PrintIfd())
	}

	//testRegExp()

	/*
		if err := generator.GenerateIptcTags(); err != nil {
			fmt.Println(err)
		}*/

	/*if err := generator.GenerateExifTags(); err != nil {
		fmt.Println(err)
	}*/
	//parseExiv2()
	//collectStandardTags()
	//fmt.Println(len(ifd.IfdSchema.Ifds))
}

func testRegExp() {
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
