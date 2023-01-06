package cmd

import (
	"fmt"
	"github.com/msvens/mimage/img"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

var rotateCommand = &cobra.Command{
	Use:   "rotate [flags] filename",
	Short: "Rotate and Crop images",
	Long:  `Rotate and/or crop images. Possibly saving any meta information `,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		copyExif, _ := cmd.Flags().GetBool("metadata")
		outputDir, _ := cmd.Flags().GetString("output")
		addDim, _ := cmd.Flags().GetBool("dimensions")
		width, _ := cmd.Flags().GetUint("width")
		height, _ := cmd.Flags().GetUint("height")
		x, _ := cmd.Flags().GetUint("xpos")
		y, _ := cmd.Flags().GetUint("ypos")
		angle, _ := cmd.Flags().GetInt("angle")
		quality, _ := cmd.Flags().GetUint("quality")
		if quality > 100 {
			return fmt.Errorf("Quality has to be between 0-100, %v", quality)
		}

		source := args[0]
		fname := filepath.Base(source)
		if addDim {
			ext := filepath.Ext(fname)
			noExt := strings.TrimSuffix(fname, ext)
			fname = fmt.Sprintf("%s-%vx%v%s", noExt, width, height, ext)
		}
		//destination file:
		dest := filepath.Join(outputDir, fname)

		//transform options
		options := img.Options{Width: int(width), Height: int(height), Quality: int(quality), X: int(x), Y: int(y), Angle: angle, CopyExif: copyExif}

		//Transform options
		return img.RotateAndCropFile(source, dest, options)

	},
}

func init() {
	rootCmd.AddCommand(rotateCommand)

	//flags
	rotateCommand.Flags().BoolP("metadata", "m", true, "Keep metadata information")
	rotateCommand.Flags().StringP("output", "o", "", "output directory (defaults to current)")
	rotateCommand.Flags().BoolP("dimensions", "d", true, "Append Dimensions to output file name")
	rotateCommand.Flags().UintP("quality", "q", 90, "Output Quality 0-100")
	rotateCommand.Flags().UintP("height", "l", 0, "Height of new image (If resize setting either height or width to 0 will keep aspect ratio)")
	rotateCommand.Flags().UintP("width", "w", 0, "Width of new image (If resize setting either height or width to 0 will keep aspect ratio)")
	rotateCommand.Flags().UintP("xpos", "x", 0, "Crop starts at x")
	rotateCommand.Flags().UintP("ypos", "y", 0, "Crop starts at y")
	rotateCommand.Flags().IntP("angle", "a", 0, "Rotation angle (in degrees). Can be negative")

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
