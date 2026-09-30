package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/msvens/mimage/metadata"
	"github.com/msvens/mimage/video"
	"github.com/spf13/cobra"
)

var videoCommand = &cobra.Command{
	Use:   "video",
	Short: "Probe, transcode and extract poster frames from videos",
	Long:  `Read metadata from mp4, QuickTime and avi videos, transcode them to web ready mp4 and extract poster frames. Needs ffmpeg and ffprobe`,
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

var videoTranscodeCommand = &cobra.Command{
	Use:   "transcode [flags] filename...",
	Short: "Transcode to a web ready mp4",
	Long: `Write each video as an H.264/AAC mp4 that plays in every browser, named after the source, into the
output directory. A source that is already web ready is copied rather than re-encoded. Dates, camera
and location tags are kept`,
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		outputDir, _ := cmd.Flags().GetString("output")
		qualityName, _ := cmd.Flags().GetString("quality")
		quality, err := video.ParseQuality(qualityName)
		if err != nil {
			return err
		}
		opts := video.TranscodeOptions{Quality: quality}
		opts.CRF, _ = cmd.Flags().GetInt("crf")
		opts.Preset, _ = cmd.Flags().GetString("preset")
		opts.MaxShortSide, _ = cmd.Flags().GetInt("max-size")
		opts.MaxFrameRate, _ = cmd.Flags().GetFloat64("max-fps")
		opts.ForceEncode, _ = cmd.Flags().GetBool("force")

		for _, f := range args {
			base := strings.TrimSuffix(filepath.Base(f), filepath.Ext(f))
			start := time.Now()
			r, err := video.Transcode(cmd.Context(), f, filepath.Join(outputDir, base), opts)
			if err != nil {
				return fmt.Errorf("%s: %w", f, err)
			}
			how := "copied"
			if r.Encoded {
				how = "encoded"
			}
			fmt.Printf("%s -> %s: %s in %v\n  %s %dx%d %.0ffps -> %s %dx%d %.0ffps\n", f, r.Path, how,
				time.Since(start).Round(time.Millisecond),
				r.Source.VideoCodec, r.Source.Width, r.Source.Height, r.Source.FrameRate,
				r.Output.VideoCodec, r.Output.Width, r.Output.Height, r.Output.FrameRate)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(videoCommand)
	videoCommand.AddCommand(videoProbeCommand, videoPosterCommand, videoTranscodeCommand)

	videoPosterCommand.Flags().StringP("output", "o", ".", "Output directory")
	videoPosterCommand.Flags().Duration("at", 0, "Position of the frame, 0 for the default of "+video.DefaultPosterOffset.String())

	videoTranscodeCommand.Flags().StringP("output", "o", ".", "Output directory")
	videoTranscodeCommand.Flags().StringP("quality", "q", "standard", "standard, high or small")
	videoTranscodeCommand.Flags().Int("crf", 0, "libx264 crf, 0 for the value of --quality")
	videoTranscodeCommand.Flags().String("preset", "", "libx264 preset, empty for the value of --quality")
	videoTranscodeCommand.Flags().Int("max-size", 0, "Cap on the shorter side, 0 for 1080, -1 for none")
	videoTranscodeCommand.Flags().Float64("max-fps", 0, "Cap on the frame rate, 0 for 60, -1 for none")
	videoTranscodeCommand.Flags().Bool("force", false, "Re-encode even a web ready source")
}
