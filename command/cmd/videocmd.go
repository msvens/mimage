package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/msvens/mimage/metadata"
	"github.com/msvens/mimage/video"
	"github.com/spf13/cobra"
)

var videoCommand = &cobra.Command{
	Use:   "video",
	Short: "Probe videos and extract poster frames",
	Long:  `Read metadata from mp4 and QuickTime videos and extract poster frames. Needs ffmpeg and ffprobe`,
}

var videoProbeCommand = &cobra.Command{
	Use:          "probe filename...",
	Short:        "Print the video summary",
	Long:         `Detect the format of each file and print its video summary as json`,
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, f := range args {
			format, err := metadata.DetectFormatFile(f)
			if err != nil {
				return err
			}
			s, err := video.Probe(cmd.Context(), f)
			if err != nil {
				return fmt.Errorf("%s (%v): %w", f, format, err)
			}
			out, err := json.MarshalIndent(s, "", "  ")
			if err != nil {
				return err
			}
			fmt.Printf("%s (%v):\n%s\n", f, format, out)
		}
		return nil
	},
}

var videoPosterCommand = &cobra.Command{
	Use:          "poster [flags] filename...",
	Short:        "Extract a poster frame",
	Long:         `Write an upright jpeg poster frame for each video, named after the video, into the output directory`,
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		outputDir, _ := cmd.Flags().GetString("output")
		at, _ := cmd.Flags().GetDuration("at")
		for _, f := range args {
			base := strings.TrimSuffix(filepath.Base(f), filepath.Ext(f))
			out, err := video.ExtractPoster(cmd.Context(), f, filepath.Join(outputDir, base), video.PosterOptions{At: at})
			if err != nil {
				return fmt.Errorf("%s: %w", f, err)
			}
			fmt.Println(out)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(videoCommand)
	videoCommand.AddCommand(videoProbeCommand, videoPosterCommand)

	videoPosterCommand.Flags().StringP("output", "o", ".", "Output directory")
	videoPosterCommand.Flags().Duration("at", 0, "Position of the frame, 0 for the default of "+video.DefaultPosterOffset.String())
}
