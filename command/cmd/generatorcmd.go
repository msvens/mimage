package cmd

import (
	"bytes"
	"fmt"
	"github.com/msvens/mimage/internal/generator"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generate Sources",
	Long:  `Re(generate) sources from external information (exif-tool, etc)`,
	RunE: func(cmd *cobra.Command, args []string) error {

		iptc, _ := cmd.Flags().GetBool("iptc")
		exif, _ := cmd.Flags().GetBool("exif")
		json, _ := cmd.Flags().GetBool("json")

		if !iptc && !exif && !json {
			return fmt.Errorf("You need to specific either iptc, exif, or json")
		}
		if json {
			fmt.Println("Generating iptc json sources using assets/iptc.pl")
			cmd := exec.Command("perl", "assets/iptc.pl")
			var out bytes.Buffer
			var err error
			cmd.Stdout = &out
			if err = cmd.Run(); err != nil {
				return err
			}
			_ = os.WriteFile("assets/exiftool-iptctags.json", out.Bytes(), 0644)

			fmt.Println("Generating exif json sources using assets/exif.pl")
			cmd = exec.Command("perl", "assets/exif.pl")
			out.Reset()
			cmd.Stdout = &out
			if err = cmd.Run(); err != nil {
				return err
			}
			_ = os.WriteFile("assets/exiftool-exiftags.json", out.Bytes(), 0644)

			fmt.Println("Genereting exiv2 exif sources")
			if err = generator.GenerateExiv2ExifJSON(); err != nil {
				return err
			}
			fmt.Println("Generate master exif json")
			if err = generator.GenerateMasterExifJSON(); err != nil {
				return err
			}
		}
		if exif {
			fmt.Println("Generate Exif Tags")
			if err := generator.GenerateExifTagsFromMasterExifJSON(); err != nil {
				return err
			}
		}
		if iptc {
			fmt.Println("Generate IPTC Tags")
			if err := generator.GenerateIptcTagsFromExifTool(); err != nil {
				return err
			}
		}
		return nil
		/*err := generator.GenerateIptcTagsFromExifTool()
		if err != nil {
			return err
		}
		return nil*/
	},
}

func init() {
	rootCmd.AddCommand(generateCommand)

	//flags
	generateCommand.Flags().BoolP("iptc", "i", false, "Generate IPTC sources")
	generateCommand.Flags().BoolP("exif", "e", false, "Generate Exif sources")
	generateCommand.Flags().BoolP("json", "j", false, "Generate Raw JSON Sources (for IPTC and Exif generation)")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
