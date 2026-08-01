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
			return fmt.Errorf("you need to specific either iptc, exif, or json")
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

			fmt.Printf("Generating exif sources using exiftool into %s\n", generator.ListxFile)
			cmd = exec.Command("exiftool", "-listx", "-EXIF:all")
			out.Reset()
			cmd.Stdout = &out
			if err = cmd.Run(); err != nil {
				return fmt.Errorf("could not run exiftool, is it installed?: %w", err)
			}
			if err = os.WriteFile(generator.ListxFile, out.Bytes(), 0644); err != nil {
				return err
			}
		}
		if exif {
			fmt.Println("Generate Exif Tags")
			if err := generator.GenerateExifTagsFromListx(); err != nil {
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
