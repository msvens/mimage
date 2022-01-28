package cmd

import (
	"github.com/msvens/mimage/internal/generator"
	"github.com/spf13/cobra"
)

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generate Sources",
	Long:  `Re(generate) sources from external information (exif-tool, etc)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generator.GenerateIptcTagsFromExifTool()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCommand)

	//flags
	generateCommand.Flags().BoolP("iptc", "i", false, "Generate IPTC sources")
	generateCommand.Flags().BoolP("exif", "e", false, "Generate Exif sources")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
