package cmd

import (
	"fmt"
	"github.com/msvens/mimage/metadata"
	"github.com/spf13/cobra"
	"strings"
)

var editCommand = &cobra.Command{
	Use:   "edit [flags] filename",
	Short: "edit image metadata",
	Long:  `Edit title, keywords and rating of your image`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		source := args[0]
		dest, _ := cmd.Flags().GetString("dest")
		if dest == "" {
			dest = source
		}
		mde, err := metadata.NewJpegEditorFile(source)
		if err != nil {
			return err
		}
		changed := false
		if cmd.Flags().Lookup("keywords").Changed {
			newKeywords, _ := cmd.Flags().GetStringSlice("keywords")

			fmt.Println("newKeywords: ", strings.Join(newKeywords, ","))
		}
		if cmd.Flags().Lookup("title").Changed {
			newTitle, _ := cmd.Flags().GetString("title")
			fmt.Println("newTitle", newTitle)
			/*err = mde.SetImageDescription(newTitle)
			if err != nil {
				return err
			}*/
			changed = false
		}
		if cmd.Flags().Lookup("rating").Changed {
			newRating, _ := cmd.Flags().GetUint("rating")
			fmt.Println("newRating: ", newRating)
		}
		if changed {
			fmt.Println("Writing changes to ", dest)
			_ = mde.WriteFile(dest)
		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(editCommand)
	editCommand.Flags().StringP("dest", "d", "", "destination file. If not set the source image will be modified")
	editCommand.Flags().StringSliceP("keywords", "k", nil, "--keywords=\"k1,k2\"")
	editCommand.Flags().StringP("title", "t", "", "image title/description")
	editCommand.Flags().UintP("rating", "r", 0, "rating (1-5)")
}
