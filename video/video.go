// Package video reads metadata from mp4, QuickTime and avi video, transcodes it
// to an mp4 every browser plays, and extracts a poster frame.
//
// It shells out to ffprobe and ffmpeg, which must be installed. This is the only
// mimage package that does: the metadata and img packages never import it, so
// a program that only handles images does not need either tool.
//
// The intended flow mirrors the image one. metadata.DetectFormatFile says
// whether a file is a video, Probe reads its metadata, Transcode makes the
// playable copy and ExtractPoster writes an upright jpeg still of it that
// img.TransformFile turns into thumbnails
package video

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// Paths of the ffprobe and ffmpeg binaries. A bare name is looked up in PATH
var (
	FFprobePath = "ffprobe"
	FFmpegPath  = "ffmpeg"
)

// Video errors. Check with errors.Is
var (
	//ErrFFmpegNotFound is returned when ffprobe or ffmpeg cannot be found
	ErrFFmpegNotFound = errors.New("mimage/video: ffmpeg not found")
	//ErrNotVideo is returned for a file whose content is not an mp4,
	//QuickTime or avi container, whatever its name
	ErrNotVideo = errors.New("mimage/video: not a video")
	//ErrNoVideoStream is returned for a container holding no playable video,
	//such as audio only, or audio with cover art
	ErrNoVideoStream = errors.New("mimage/video: no video stream")
	//ErrTruncated is returned for a file shorter than its own structure says
	//it is, typically an interrupted copy or download
	ErrTruncated = errors.New("mimage/video: file is truncated")
	//ErrNoFrame is returned when ffmpeg ran but produced no poster
	ErrNoFrame = errors.New("mimage/video: no frame extracted")
)

// ToolError is a failed ffprobe or ffmpeg run. Stderr holds what the tool said,
// which is usually the useful part
type ToolError struct {
	Tool   string
	Args   []string
	Stderr string
	Err    error
}

func (e *ToolError) Error() string {
	msg := fmt.Sprintf("mimage/video: %s failed: %v", e.Tool, e.Err)
	if e.Stderr != "" {
		msg += ": " + e.Stderr
	}
	return msg
}

func (e *ToolError) Unwrap() error {
	return e.Err
}

// Available reports whether both ffprobe and ffmpeg can be found, returning an
// error wrapping ErrFFmpegNotFound if not. Useful as a startup check
func Available() error {
	for _, tool := range []string{FFprobePath, FFmpegPath} {
		if _, err := exec.LookPath(tool); err != nil {
			return fmt.Errorf("%w: %s: %v", ErrFFmpegNotFound, tool, err)
		}
	}
	return nil
}

// run executes a tool and returns its stdout
func run(ctx context.Context, tool string, args ...string) ([]byte, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, tool, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return nil, fmt.Errorf("%w: %s", ErrFFmpegNotFound, tool)
		}
		return nil, &ToolError{Tool: tool, Args: args, Stderr: strings.TrimSpace(stderr.String()), Err: err}
	}
	return stdout.Bytes(), nil
}
