package main

import (
	"fmt"
	"github.com/msvens/mimage/metadata"
)

func main() {

	fname := "/Users/mellowtech/go/src/github.com/msvens/mimage/assets/leica.jpg"

	md, err := metadata.NewMetaDataFromJpegFile(fname)
	if err != nil {
		fmt.Println(err.Error())
	}
	if md != nil {
		fmt.Println(md.Summary)
	}

}
