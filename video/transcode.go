package video

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/msvens/mimage/img"
	"github.com/msvens/mimage/metadata"
)

// Quality is a named set of encoder settings for Transcode
type Quality int

const (
	//QualityStandard looks the same as the source to most eyes at a
	//reasonable size. The default
	QualityStandard Quality = iota
	//QualityHigh keeps more detail at roughly twice the size, and encodes
	//slower
	QualityHigh
	//QualitySmall trades visible detail for roughly half the size
	QualitySmall
)

var qualityNames = map[Quality]string{
	QualityStandard: "standard",
	QualityHigh:     "high",
	QualitySmall:    "small",
}

func (q Quality) String() string {
	if n, found := qualityNames[q]; found {
		return n
	}
	return "unknown"
}

// ParseQuality reads a Quality by its name, as String writes it
func ParseQuality(name string) (Quality, error) {
	for q, n := range qualityNames {
		if strings.EqualFold(n, name) {
			return q, nil
		}
	}
	return 0, fmt.Errorf("mimage/video: unknown quality %q, use standard, high or small", name)
}

// the encoder settings behind each Quality. crf is libx264's constant rate
// factor, lower is better and larger, 18 is often called visually lossless
type qualitySettings struct {
	crf          int
	preset       string
	audioBitrate int
}

var qualitySettingsFor = map[Quality]qualitySettings{
	QualityStandard: {crf: 21, preset: "medium", audioBitrate: 128},
	QualityHigh:     {crf: 18, preset: "slow", audioBitrate: 192},
	QualitySmall:    {crf: 26, preset: "medium", audioBitrate: 96},
}

// Transcode defaults
const (
	//DefaultMaxShortSide caps the shorter side of the output, so 1080p in
	//either orientation. A 4K phone clip comes out as 1080p
	DefaultMaxShortSide = 1080
	//DefaultMaxFrameRate caps the frame rate. Slow motion recorded at 120 or
	//240 fps is dropped to 60, which is also all a browser shows
	DefaultMaxFrameRate = 60
)

// TranscodeOptions controls Transcode. The zero value is the recommended
// setting: QualityStandard, at most 1080p and 60 fps. Any field set explicitly
// overrides what Quality would choose
type TranscodeOptions struct {
	Quality Quality
	//CRF is libx264's constant rate factor, 0-51 with lower meaning better.
	//Zero means the value for Quality
	CRF int
	//Preset is a libx264 preset, from "ultrafast" to "veryslow". Slower gives
	//a smaller file at the same quality. Empty means the value for Quality
	Preset string
	//AudioBitrate in kbit/s, used when the sound has to be re-encoded. Zero
	//means the value for Quality
	AudioBitrate int
	//MaxShortSide caps the shorter side of the output. Zero means
	//DefaultMaxShortSide, a negative value never downscales
	MaxShortSide int
	//MaxFrameRate caps the frame rate. Zero means DefaultMaxFrameRate, a
	//negative value never drops frames
	MaxFrameRate float64
	//ForceEncode re-encodes a video that could have been copied as is
	ForceEncode bool
}

// the options with every default filled in
type transcodeSettings struct {
	qualitySettings
	maxShortSide int
	maxFrameRate float64
	forceEncode  bool
}

func (o TranscodeOptions) settings() (transcodeSettings, error) {
	qs, found := qualitySettingsFor[o.Quality]
	if !found {
		return transcodeSettings{}, fmt.Errorf("mimage/video: unknown quality %d", o.Quality)
	}
	if o.CRF < 0 || o.CRF > 51 {
		return transcodeSettings{}, fmt.Errorf("mimage/video: crf %d outside 0-51", o.CRF)
	}
	s := transcodeSettings{qualitySettings: qs, maxShortSide: o.MaxShortSide,
		maxFrameRate: o.MaxFrameRate, forceEncode: o.ForceEncode}
	if o.CRF > 0 {
		s.crf = o.CRF
	}
	if o.Preset != "" {
		s.preset = o.Preset
	}
	if o.AudioBitrate > 0 {
		s.audioBitrate = o.AudioBitrate
	}
	if s.maxShortSide == 0 {
		s.maxShortSide = DefaultMaxShortSide
	}
	if s.maxFrameRate == 0 {
		s.maxFrameRate = DefaultMaxFrameRate
	}
	return s, nil
}

// TranscodeResult describes what Transcode wrote
type TranscodeResult struct {
	//Path is the mp4 written, destBase with ".mp4" appended
	Path string
	//Encoded is true when the picture was re-encoded, false when the source
	//was already web ready and only copied into a fresh mp4
	Encoded bool
	//Source is the probe of the original. Its metadata is what to keep, the
	//output carries the same date and tags but Source is the authority
	Source *Summary
	//Output is the probe of the file written
	Output *Summary
}

// Transcode errors. Check with errors.Is
var (
	//ErrHDRUnsupported is returned for an HDR source. Converting HDR to
	//standard range needs an ffmpeg built with zimg, which not every install
	//has, and encoding it without conversion gives a washed out picture. An
	//iPhone records HDR unless HDR Video is switched off in its settings
	ErrHDRUnsupported = errors.New("mimage/video: HDR video is not supported")
	//ErrOutputCheck is returned when the written file does not match the
	//source: a different length, orientation or date. Nothing is left behind
	ErrOutputCheck = errors.New("mimage/video: transcoded output failed its check")
)

// Transcode writes src as an mp4 that plays in every browser: H.264 video in
// 8 bit standard range, AAC sound, the index at the front so playback starts
// before the download ends, and the source's date, camera and location tags.
//
// A source that is already all of that, within the size and frame rate caps, is
// copied rather than re-encoded, which is fast and loses nothing. Otherwise it is
// encoded, fixing on the way whatever needs it: rotation is turned into upright
// pixels, interlacing removed, full range and non square pixels converted. HDR
// is refused with ErrHDRUnsupported.
//
// ".mp4" is appended to destBase, which must not carry an extension of its own.
// The result is checked against the source before it is kept, and on any error
// nothing is left at the destination
func Transcode(ctx context.Context, src, destBase string, opts TranscodeOptions) (*TranscodeResult, error) {
	if ext := filepath.Ext(destBase); metadata.IsImageExtension(ext) || isVideoExtension(ext) {
		return nil, fmt.Errorf("%w: %s, pass a name without one", img.ErrDestHasExtension, destBase)
	}
	settings, err := opts.settings()
	if err != nil {
		return nil, err
	}
	source, err := Probe(ctx, src)
	if err != nil {
		return nil, err
	}
	if source.HDR {
		return nil, fmt.Errorf("%w: %s", ErrHDRUnsupported, src)
	}

	dest := destBase + metadata.FormatMp4.Extension()
	//written beside the destination and renamed, so a failure never leaves a
	//half written file under the final name. Not a dot file: iCloud Drive sets
	//the hidden flag on dot files it sees, and a rename keeps that flag, so
	//the finished video would be invisible in Finder
	tmp, err := os.CreateTemp(filepath.Dir(dest), filepath.Base(destBase)+".transcoding-*.mp4")
	if err != nil {
		return nil, err
	}
	_ = tmp.Close()
	keep := false
	defer func() {
		if !keep {
			_ = os.Remove(tmp.Name())
		}
	}()

	args, encode := transcodeArgs(source, src, tmp.Name(), settings)
	if _, err = run(ctx, FFmpegPath, args...); err != nil {
		return nil, err
	}

	output, err := Probe(ctx, tmp.Name())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrOutputCheck, err)
	}
	if err = checkOutput(source, output); err != nil {
		return nil, err
	}
	//CreateTemp makes the file private to its owner, the result is an
	//ordinary file others may read
	if err = os.Chmod(tmp.Name(), 0644); err != nil {
		return nil, err
	}
	if err = os.Rename(tmp.Name(), dest); err != nil {
		return nil, err
	}
	keep = true
	return &TranscodeResult{Path: dest, Encoded: encode, Source: source, Output: output}, nil
}

// webReadyProfiles are the H.264 profiles every browser decodes
var webReadyProfiles = map[string]bool{
	"Baseline": true, "Constrained Baseline": true, "Main": true, "High": true,
}

// canCopyVideo reports whether the picture can go into the output untouched
func canCopyVideo(s *Summary, t transcodeSettings) bool {
	return !t.forceEncode &&
		s.VideoCodec == "h264" && webReadyProfiles[s.VideoProfile] && s.PixelFormat == "yuv420p" &&
		!s.HDR && !s.Interlaced && s.PixelAspect == 1 &&
		(t.maxShortSide < 0 || min(s.Width, s.Height) <= t.maxShortSide) &&
		(t.maxFrameRate < 0 || s.FrameRate <= t.maxFrameRate+0.01)
}

// transcodeArgs builds the ffmpeg command, and reports whether the picture is
// re-encoded. It is pure so the decisions can be tested without ffmpeg
func transcodeArgs(s *Summary, src, dest string, t transcodeSettings) ([]string, bool) {
	args := []string{"-hide_banner", "-v", "error", "-nostdin", "-y", "-i", src,
		//the video stream Probe chose, never cover art, and the first sound if
		//there is one
		"-map", "0:" + strconv.Itoa(s.streamIndex), "-map", "0:a:0?"}

	encode := !canCopyVideo(s, t)
	if encode {
		args = append(args, "-c:v", "libx264", "-crf", strconv.Itoa(t.crf), "-preset", t.preset,
			"-profile:v", "high", "-pix_fmt", "yuv420p")
		if vf := videoFilters(s, t); vf != "" {
			args = append(args, "-vf", vf)
		}
	} else {
		args = append(args, "-c:v", "copy")
	}

	if s.AudioCodec == "aac" {
		args = append(args, "-c:a", "copy")
	} else {
		args = append(args, "-c:a", "aac", "-b:a", strconv.Itoa(t.audioBitrate)+"k")
	}

	//use_metadata_tags keeps tags mp4 has no standard box for, Apple's and
	//Android's keys. It would also copy the source's own container
	//description, which the new file must not claim, so those are cleared
	args = append(args, "-map_metadata", "0", "-movflags", "+faststart+use_metadata_tags")
	for _, key := range []string{"major_brand", "minor_version", "compatible_brands", "encoder"} {
		args = append(args, "-metadata", key+"=")
	}
	if !s.CreationTime.IsZero() {
		//written explicitly so the date survives whatever the source container
		//kept it in
		args = append(args, "-metadata", "creation_time="+s.CreationTime.UTC().Format(time.RFC3339))
	}
	return append(args, "-f", "mp4", dest), encode
}

// videoFilters builds the -vf chain for an encode. Rotation is not in it:
// ffmpeg turns the frames upright by itself before these filters run
func videoFilters(s *Summary, t transcodeSettings) string {
	var vf []string
	if s.Interlaced {
		//one frame per frame: the default of one per field doubles the rate
		vf = append(vf, "bwdif=mode=send_frame")
	}

	//the frame as it enters the filters: already turned upright
	inW, inH := s.StoredWidth, s.StoredHeight
	if s.Rotation == 90 || s.Rotation == 270 {
		inW, inH = inH, inW
	}
	w, h := outputSize(s.Width, s.Height, t.maxShortSide)
	if w != inW || h != inH || s.PixelAspect != 1 {
		vf = append(vf, fmt.Sprintf("scale=%d:%d", w, h), "setsar=1")
	}

	if t.maxFrameRate > 0 && s.FrameRate > t.maxFrameRate+0.01 {
		vf = append(vf, "fps="+strconv.FormatFloat(t.maxFrameRate, 'f', -1, 64))
	}
	//motion jpeg, from old cameras, is full range. Left as is the output is
	//flagged full range too, which players show with the wrong contrast
	if strings.HasPrefix(s.PixelFormat, "yuvj") {
		vf = append(vf, "scale=out_range=tv")
	}
	vf = append(vf, "format=yuv420p")
	return strings.Join(vf, ",")
}

// outputSize scales a display size so its shorter side is at most maxShort,
// keeping the aspect ratio. Both sides come out even, which 4:2:0 video needs
func outputSize(w, h, maxShort int) (int, int) {
	if short := min(w, h); maxShort > 0 && short > maxShort {
		f := float64(maxShort) / float64(short)
		w, h = int(math.Round(float64(w)*f)), int(math.Round(float64(h)*f))
	}
	return w &^ 1, h &^ 1
}

// checkOutput compares what was written with its source
func checkOutput(src, out *Summary) error {
	if out.VideoCodec != "h264" {
		return fmt.Errorf("%w: video is %s, not h264", ErrOutputCheck, out.VideoCodec)
	}
	diff := (out.Duration - src.Duration).Abs()
	if allowed := max(500*time.Millisecond, src.Duration/50); src.Duration > 0 && diff > allowed {
		return fmt.Errorf("%w: duration %v, source %v", ErrOutputCheck, out.Duration, src.Duration)
	}
	if (out.Width > out.Height) != (src.Width > src.Height) {
		return fmt.Errorf("%w: %dx%d from a %dx%d source", ErrOutputCheck, out.Width, out.Height, src.Width, src.Height)
	}
	if !src.CreationTime.IsZero() && !out.CreationTime.Equal(src.CreationTime) {
		return fmt.Errorf("%w: creation time %v, source %v", ErrOutputCheck, out.CreationTime, src.CreationTime)
	}
	return nil
}

func isVideoExtension(ext string) bool {
	switch strings.ToLower(ext) {
	case ".mp4", ".m4v", ".mov", ".avi", ".3gp", ".3g2":
		return true
	}
	return false
}
