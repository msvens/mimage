package main

import (
	"fmt"
	"github.com/msvens/mpimage/metadata"
)

func main() {

	fname := "/Users/mellowtech/go/src/github.com/msvens/mimage/assets/nikon.jpg"

	md, err := metadata.NewMetaDataFromJpegFile(fname);
	if err != nil {
		fmt.Println(err.Error())
	}
	if md != nil {
		fmt.Println(md.Summary)
	}

}
