/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/msvens/mimage/command/cmd"

func main() {
	/*ts := TestStruct{}
	appendToTestStruct(&ts, "str1")
	appendToTestStruct(&ts, "str2")
	metadata.TryMsb()
	fmt.Println(ts.data)*/
	//generator.GenerateExifTagsFromExifTool()
	//generator.AlignExifToolJson()
	cmd.Execute()
	/*t := time.Now()
	fmt.Println(t.Format(metadata.IptcShortDate))
	fmt.Println(t.Format(metadata.IptcTime))
	fmt.Println(t.Format("-0700"))*/
	/*err := metadata.TryPhotoshopInfo("assets/leica.jpg")
	if err != nil {
		fmt.Println(err)
	}*/
	/*
		md, err := metadata.NewMetaDataFromFile("assets/leica.jpg")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(md.Iptc().String())
			fmt.Println(md.Summary())
		}*/

}
