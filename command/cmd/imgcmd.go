package cmd

import (
	"fmt"
	"github.com/msvens/mimage/img"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

var imgCommand = &cobra.Command{
	Use:   "transform [flags] filename",
	Short: "Copy and resize image",
	Long:  `Resize,Crop and Copy image to a new location. Possibly saving any meta information `,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		copyExif, _ := cmd.Flags().GetBool("metadata")
		outputDir, _ := cmd.Flags().GetString("output")
		addDim, _ := cmd.Flags().GetBool("dimensions")
		cropAnchor, err := parseCropFlag(cmd)
		if err != nil {
			return err
		}
		transType, err := parseTypeFlag(cmd)
		if err != nil {
			return err
		}
		strategy, err := parseStrategyFlag(cmd)
		if err != nil {
			return err
		}

		width, _ := cmd.Flags().GetUint("xdim")
		height, _ := cmd.Flags().GetUint("ydim")
		quality, _ := cmd.Flags().GetUint("quality")
		if quality > 100 {
			return fmt.Errorf("Quality has to be between 0-100, %v", quality)
		}
		if width == 0 && height == 0 {
			return fmt.Errorf("xdim and ydim cannot both be 0")
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
		options := img.Options{Width: int(width), Height: int(height), Quality: int(quality), Anchor: cropAnchor,
			Transform: transType, Strategy: strategy, CopyExif: copyExif}

		fmt.Println(dest)
		fmt.Println(options)
		//Transform options
		transforms := map[string]img.Options{
			dest: options,
		}
		return img.TransformFile(source, transforms)

	},
}

func parseStrategyFlag(cmd *cobra.Command) (img.ResampleStrategy, error) {
	s, _ := cmd.Flags().GetString("strategy")
	switch s {
	case "Lanczos":
		return img.Lanczos, nil
	case "NearestNeighbor":
		return img.NearestNeighbor, nil
	default:
		return img.Lanczos, fmt.Errorf("Unknown resample strategy: %s", s)
	}
}

func parseTypeFlag(cmd *cobra.Command) (img.TransformType, error) {
	t, _ := cmd.Flags().GetString("type")
	switch t {
	case "ResizeAndCrop":
		return img.ResizeAndCrop, nil
	case "Crop":
		return img.Crop, nil
	case "Resize":
		return img.Resize, nil
	case "ResizeAndFit":
		return img.ResizeAndFit, nil
	default:
		return img.ResizeAndCrop, fmt.Errorf("Unknown Tranform Type: %s", t)
	}
}

func parseCropFlag(cmd *cobra.Command) (img.CropAnchor, error) {
	crop, _ := cmd.Flags().GetString("crop")
	switch crop {
	case "Center":
		return img.Center, nil
	case "TopLeft":
		return img.TopLeft, nil
	case "TopRight":
		return img.TopRight, nil
	case "Left":
		return img.Left, nil
	case "Right":
		return img.Right, nil
	case "BottomLeft":
		return img.BottomLeft, nil
	case "Bottom":
		return img.Bottom, nil
	case "BottomRight":
		return img.BottomRight, nil
	default:
		return img.Center, fmt.Errorf("Unrecognized Crop: %s", crop)
	}
}

func init() {
	rootCmd.AddCommand(imgCommand)

	//flags
	imgCommand.Flags().BoolP("metadata", "m", true, "Keep metadata information")
	imgCommand.Flags().StringP("output", "o", "", "output directory (defaults to current)")
	imgCommand.Flags().BoolP("dimensions", "d", true, "Append Dimensions to output file name")
	imgCommand.Flags().StringP("crop", "c", "Center", "Center, TopLeft, Top, TopRight, Left, Right, BottomLeft, Bottom, BottomRight")
	imgCommand.Flags().StringP("type", "t", "ResizeAndCrop", "ResizeAndCrop, Crop, Resize, ResizeAndFit")
	imgCommand.Flags().StringP("strategy", "s", "Lanczos", "Lanczos, NearestNeighbor")
	imgCommand.Flags().UintP("quality", "q", 90, "Output Quality 0-100")
	imgCommand.Flags().UintP("xdim", "x", 0, "Height of new image (If resize setting either height or width to 0 will keep aspect ratio)")
	imgCommand.Flags().UintP("ydim", "y", 0, "Width of new image (If resize setting either height or width to 0 will keep aspect ratio)")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
