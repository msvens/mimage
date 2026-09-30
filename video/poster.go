package video

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/msvens/mimage/img"
	"github.com/msvens/mimage/metadata"
)

// DefaultPosterOffset is how far into a video the poster is taken when
// PosterOptions.At is zero. Far enough to skip the black or faded opening frame
// many clips start with
const DefaultPosterOffset = time.Second

// PosterOptions controls ExtractPoster
type PosterOptions struct {
	//At is the position of the frame to take. Zero means DefaultPosterOffset.
	//It is clamped to half the duration, so a short clip still yields a frame
	At time.Duration
}

// ExtractPoster writes a single frame of the video at src as a jpeg, rotated
// upright, and returns the path written. ".jpg" is appended to destBase, which
// must not already carry an image extension, see img.ErrDestHasExtension. The
// jpeg is full size and meant as the source for img.TransformFile.
//
// The video is probed first, so the errors of Probe apply here as well. An
// existing file at the destination is overwritten
func ExtractPoster(ctx context.Context, src, destBase string, opts PosterOptions) (string, error) {
	if metadata.IsImageExtension(filepath.Ext(destBase)) {
		return "", fmt.Errorf("%w: %s, pass a name without one", img.ErrDestHasExtension, destBase)
	}
	dest := destBase + metadata.FormatJpeg.Extension()

	s, err := Probe(ctx, src)
	if err != nil {
		return "", err
	}
	at := opts.At
	if at <= 0 {
		at = DefaultPosterOffset
	}
	if s.Duration > 0 && at > s.Duration/2 {
		at = s.Duration / 2
	}

	//seeking before -i is fast and, since ffmpeg decodes from the preceding
	//keyframe, still frame accurate. ffmpeg rotates the frame upright by
	//default, which is why no rotation happens here
	_, err = run(ctx, FFmpegPath, "-hide_banner", "-v", "error", "-nostdin",
		"-ss", strconv.FormatFloat(at.Seconds(), 'f', 3, 64), "-i", src,
		"-map", "0:"+strconv.Itoa(s.streamIndex), "-frames:v", "1", "-an",
		"-q:v", "2", "-f", "image2", "-y", dest)
	if err != nil {
		_ = os.Remove(dest)
		return "", err
	}
	//a seek past the last frame is not an error to ffmpeg, it just writes
	//nothing
	if info, statErr := os.Stat(dest); statErr != nil || info.Size() == 0 {
		_ = os.Remove(dest)
		return "", fmt.Errorf("%w: %s at %v", ErrNoFrame, src, at)
	}
	return dest, nil
}
