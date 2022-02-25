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
		je, err := metadata.NewJpegEditorFile(source)
		if err != nil {
			return err
		}
		changed := false
		if cmd.Flags().Lookup("keywords").Changed {
			newKeywords, _ := cmd.Flags().GetStringSlice("keywords")
			fmt.Println("setting new keywords: ", strings.Join(newKeywords, ","))
			if err = je.SetKeywords(newKeywords); err != nil {
				return err
			}
			changed = true
		}
		if cmd.Flags().Lookup("title").Changed {
			newTitle, _ := cmd.Flags().GetString("title")
			fmt.Println("setting newTitle", newTitle)
			if err = je.SetTitle(newTitle); err != nil {
				return err
			}
			changed = true
		}
		if cmd.Flags().Lookup("rating").Changed {
			newRating, err := cmd.Flags().GetUint16("rating")
			if err != nil {
				return err
			}
			fmt.Println("setting new rating: ", newRating)
			je.Xmp().SetRating(newRating)
			changed = true

		}
		if changed {
			fmt.Println("Writing changes to ", dest)
			err = je.WriteFile(dest)
			if err != nil {
				return err
			}
		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(editCommand)
	editCommand.Flags().StringP("dest", "d", "", "destination file. If not set the source image will be modified")
	editCommand.Flags().StringSliceP("keywords", "k", nil, "--keywords=\"k1,k2\"")
	editCommand.Flags().StringP("title", "t", "", "image title/description")
	editCommand.Flags().Uint16P("rating", "r", 0, "rating (1-5)")
}
