package cmd

import (
	"fmt"
	"github.com/msvens/mimage/metadata"
	"github.com/spf13/cobra"
)

var metadataCommand = &cobra.Command{
	Use:   "metadata",
	Short: "Extract Metadata",
	Long:  `Extract all or some of the metadata found in the provided jpeg image`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		md, err := metadata.NewMetaDataFromFile(args[0])
		if err != nil {
			return err
		}
		fmt.Println(md)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(metadataCommand)

	//flags
	metadataCommand.Flags().BoolP("all", "a", false, "Extract all metadata")
	metadataCommand.Flags().BoolP("summary", "s", true, "Extract metadata summary")
	metadataCommand.Flags().BoolP("exif", "e", false, "Extract Exif sata")
	metadataCommand.Flags().BoolP("xmp", "x", false, "Extract Xmp data")
	metadataCommand.Flags().BoolP("iptc", "i", false, "Extract Iptc data")
	metadataCommand.Flags().BoolP("json", "j", false, "Output as Json (xmp will always be outputted as xml)")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
