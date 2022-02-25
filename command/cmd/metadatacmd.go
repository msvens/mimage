package cmd

import (
	"fmt"
	"github.com/msvens/mimage/metadata"
	"github.com/spf13/cobra"
)

var metadataCommand = &cobra.Command{
	Use:   "metadata [flags] filename",
	Short: "Extract Metadata",
	Long:  `Extract all or some of the metadata found in the provided jpeg image`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("No file specified")
		}
		md, err := metadata.NewMetaDataFromFile(args[0])
		if err != nil {
			return err
		}
		summary, _ := cmd.Flags().GetBool("summary")
		exif, _ := cmd.Flags().GetBool("exif")
		xmp, _ := cmd.Flags().GetBool("xmp")
		iptc, _ := cmd.Flags().GetBool("iptc")
		json, _ := cmd.Flags().GetBool("json")
		if json {
			fmt.Println("output as json...not yet implemented")
		} else {
			if summary {
				fmt.Println(md.Summary().String())
			}
			if exif {
				fmt.Println(md.Exif().String())
			}
			if iptc {
				fmt.Println(md.Iptc().String())
			}
			if xmp {
				fmt.Println(md.Xmp().String())
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(metadataCommand)

	//flags
	metadataCommand.Flags().BoolP("summary", "s", true, "Extract metadata summary")
	metadataCommand.Flags().BoolP("exif", "e", false, "Extract Exif data")
	metadataCommand.Flags().BoolP("xmp", "x", false, "Extract Xmp data")
	metadataCommand.Flags().BoolP("iptc", "i", false, "Extract Iptc data")
	metadataCommand.Flags().BoolP("json", "j", false, "Output as Json")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
