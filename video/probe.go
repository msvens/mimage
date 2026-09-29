package video

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/msvens/mimage/metadata"
)

// Location is a position in decimal degrees. Altitude is in meters, and zero
// when the source does not give one
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude,omitempty"`
}

// Summary holds the video metadata of interest. Anything the file does not
// carry is left at its zero value: a video with no camera tags or no location is
// normal, not an error
type Summary struct {
	//Duration is encoded in json as nanoseconds, like any time.Duration
	Duration time.Duration `json:"duration"`
	//Width and Height are the display dimensions, what a viewer sees, with
	//rotation and any non square pixels applied
	Width  int `json:"width"`
	Height int `json:"height"`
	//StoredWidth and StoredHeight are the dimensions of the encoded frames,
	//before rotation
	StoredWidth  int `json:"storedWidth"`
	StoredHeight int `json:"storedHeight"`
	//Rotation is how many degrees, clockwise, the stored frames are turned for
	//display: 0, 90, 180 or 270. A phone recording held upright is usually 90
	Rotation     int       `json:"rotation,omitempty"`
	CreationTime time.Time `json:"creationTime,omitzero"`
	CameraMake   string    `json:"cameraMake,omitempty"`
	CameraModel  string    `json:"cameraModel,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	Keywords     []string  `json:"keywords,omitempty"`
	Location     *Location `json:"location,omitempty"`
	VideoCodec   string    `json:"videoCodec,omitempty"`
	AudioCodec   string    `json:"audioCodec,omitempty"`
	FrameRate    float64   `json:"frameRate,omitempty"`

	//index of the video stream in the container, so ffmpeg can be pointed at
	//it rather than at cover art
	streamIndex int
}

// Probe reads the metadata of an mp4 or QuickTime video. It is the real check
// that a file is usable: metadata.Format.Supported only says the container
// looked right.
//
// The file must be a video by content, see metadata.DetectFormatFile, or
// ErrNotVideo is returned. A file cut short returns ErrTruncated, a container
// without a video stream returns ErrNoVideoStream, and one ffprobe cannot read
// returns a *ToolError
func Probe(ctx context.Context, fileName string) (*Summary, error) {
	format, err := metadata.DetectFormatFile(fileName)
	if err != nil {
		return nil, err
	}
	if !format.IsVideo() {
		return nil, fmt.Errorf("%w: %s is %v", ErrNotVideo, fileName, format)
	}
	if err = checkComplete(fileName); err != nil {
		return nil, err
	}
	out, err := run(ctx, FFprobePath, "-v", "error", "-print_format", "json",
		"-show_format", "-show_streams", "-export_xmp", "1", fileName)
	if err != nil {
		return nil, err
	}
	return parseProbe(out)
}

// the parts of ffprobe's json output that are read
type probeOutput struct {
	Streams []probeStream `json:"streams"`
	Format  struct {
		Duration string            `json:"duration"`
		Tags     map[string]string `json:"tags"`
	} `json:"format"`
}

type probeStream struct {
	Index             int               `json:"index"`
	CodecType         string            `json:"codec_type"`
	CodecName         string            `json:"codec_name"`
	Width             int               `json:"width"`
	Height            int               `json:"height"`
	SampleAspectRatio string            `json:"sample_aspect_ratio"`
	AvgFrameRate      string            `json:"avg_frame_rate"`
	RFrameRate        string            `json:"r_frame_rate"`
	Duration          string            `json:"duration"`
	Tags              map[string]string `json:"tags"`
	Disposition       struct {
		AttachedPic int `json:"attached_pic"`
	} `json:"disposition"`
	SideDataList []struct {
		SideDataType string  `json:"side_data_type"`
		Rotation     float64 `json:"rotation"`
	} `json:"side_data_list"`
}

// parseProbe turns ffprobe json into a Summary
func parseProbe(data []byte) (*Summary, error) {
	var probe probeOutput
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, fmt.Errorf("mimage/video: could not parse ffprobe output: %w", err)
	}

	var video *probeStream
	s := &Summary{}
	for i := range probe.Streams {
		st := &probe.Streams[i]
		switch {
		//cover art is a video stream of one picture, not the video
		case st.CodecType == "video" && st.Disposition.AttachedPic == 0 && video == nil:
			video = st
		case st.CodecType == "audio" && s.AudioCodec == "":
			s.AudioCodec = st.CodecName
		}
	}
	if video == nil {
		return nil, ErrNoVideoStream
	}

	s.streamIndex = video.Index
	s.VideoCodec = video.CodecName
	s.StoredWidth, s.StoredHeight = video.Width, video.Height
	s.Rotation = rotation(video)
	s.Width, s.Height = displaySize(video.Width, video.Height, video.SampleAspectRatio, s.Rotation)
	s.FrameRate = ratio(video.AvgFrameRate)
	if s.FrameRate == 0 {
		s.FrameRate = ratio(video.RFrameRate)
	}
	s.Duration = seconds(probe.Format.Duration)
	if s.Duration == 0 {
		s.Duration = seconds(video.Duration)
	}

	tags := probe.Format.Tags
	var xd metadata.XmpData
	if raw := tag(tags, "xmp"); raw != "" {
		//unparseable xmp is treated as absent, like any other missing tag
		xd, _ = metadata.NewXmpDataFromBytes([]byte(raw))
	}
	xmpMake, xmpModel := xd.GetCamera()

	s.CreationTime = creationTime(tags, video.Tags, xd)
	s.CameraMake = first(tag(tags, "com.apple.quicktime.make"), tag(tags, "com.android.manufacturer"),
		tag(tags, "make"), xmpMake)
	s.CameraModel = first(tag(tags, "com.apple.quicktime.model"), tag(tags, "com.android.model"),
		tag(tags, "model"), xmpModel)
	s.Title = first(tag(tags, "title"), xd.GetTitle())
	s.Description = first(tag(tags, "description"), xd.GetDescription())
	s.Keywords = xd.GetKeywords()
	if len(s.Keywords) == 0 {
		s.Keywords = nil
	}
	s.Location = location(tags, xd)
	return s, nil
}

// rotation returns the clockwise display rotation, normalised to 0, 90, 180 or
// 270. Current ffprobe reports a display matrix whose rotation is counter
// clockwise; ffmpeg 4 and earlier reported a clockwise "rotate" tag instead
func rotation(st *probeStream) int {
	deg := 0.0
	found := false
	for _, sd := range st.SideDataList {
		if sd.SideDataType == "Display Matrix" {
			deg, found = -sd.Rotation, true
			break
		}
	}
	if !found {
		if v, err := strconv.ParseFloat(tag(st.Tags, "rotate"), 64); err == nil {
			deg = v
		}
	}
	quarter := int(math.Round(deg/90)) % 4
	if quarter < 0 {
		quarter += 4
	}
	return quarter * 90
}

// displaySize applies non square pixels and rotation to the stored size
func displaySize(width, height int, sar string, rotation int) (int, int) {
	if r := ratio(sar); r > 0 && r != 1 {
		width = int(math.Round(float64(width) * r))
	}
	if rotation == 90 || rotation == 270 {
		return height, width
	}
	return width, height
}

// ratio parses ffprobe's "num/den" and "num:den" forms, 0 when there is none
func ratio(s string) float64 {
	num, den, found := strings.Cut(strings.ReplaceAll(s, ":", "/"), "/")
	if !found {
		return 0
	}
	n, err1 := strconv.ParseFloat(num, 64)
	d, err2 := strconv.ParseFloat(den, 64)
	if err1 != nil || err2 != nil || d == 0 {
		return 0
	}
	return n / d
}

func seconds(s string) time.Duration {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil || v <= 0 {
		return 0
	}
	return time.Duration(v * float64(time.Second))
}

// creationTime picks the most trustworthy capture time. Apple's creationdate
// keeps the local offset, xmp is what Lightroom writes, and the container
// creation_time is UTC and often reset by editing software, so it comes last
func creationTime(tags, streamTags map[string]string, xd metadata.XmpData) time.Time {
	if t, ok := parseTime(tag(tags, "com.apple.quicktime.creationdate")); ok {
		return t
	}
	if t := xd.GetDate(); !t.IsZero() && !isEpoch(t) {
		return t
	}
	for _, v := range []string{tag(tags, "creation_time"), tag(streamTags, "creation_time")} {
		if t, ok := parseTime(v); ok {
			return t
		}
	}
	return time.Time{}
}

var timeLayouts = []string{
	time.RFC3339Nano,
	"2006-01-02T15:04:05Z0700",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
}

// parseTime reads a tag timestamp. ffprobe joins repeated tags with ';', so
// only the first value is used. A container that never had its time set reads
// as the 1904 or 1970 epoch, which is treated as absent
func parseTime(s string) (time.Time, bool) {
	s, _, _ = strings.Cut(s, ";")
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false
	}
	for _, layout := range timeLayouts {
		if t, err := time.Parse(layout, s); err == nil && !isEpoch(t) {
			return t, true
		}
	}
	return time.Time{}, false
}

func isEpoch(t time.Time) bool {
	return t.Year() <= 1904 || t.Unix() == 0
}

func location(tags map[string]string, xd metadata.XmpData) *Location {
	for _, key := range []string{"com.apple.quicktime.location.ISO6709", "location", "location-eng"} {
		if l, ok := parseISO6709(tag(tags, key)); ok {
			return l
		}
	}
	if lat, long, ok := xd.GetLocation(); ok {
		return &Location{Latitude: lat, Longitude: long}
	}
	return nil
}

// parseISO6709 reads the decimal degree form phones write, such as
// "+59.3293+018.0686+012.345/": signed latitude, signed longitude and an
// optional signed altitude, with a trailing solidus
func parseISO6709(s string) (*Location, bool) {
	s = strings.TrimSuffix(strings.TrimSpace(s), "/")
	var parts []string
	for len(s) > 0 {
		if s[0] != '+' && s[0] != '-' {
			return nil, false
		}
		end := strings.IndexAny(s[1:], "+-")
		if end < 0 {
			parts = append(parts, s)
			break
		}
		parts = append(parts, s[:end+1])
		s = s[end+1:]
	}
	if len(parts) < 2 || len(parts) > 3 {
		return nil, false
	}
	var v [3]float64
	for i, p := range parts {
		f, err := strconv.ParseFloat(p, 64)
		if err != nil {
			return nil, false
		}
		v[i] = f
	}
	if math.Abs(v[0]) > 90 || math.Abs(v[1]) > 180 {
		return nil, false
	}
	return &Location{Latitude: v[0], Longitude: v[1], Altitude: v[2]}, true
}

// tag looks a key up case insensitively, since writers differ in case
func tag(tags map[string]string, key string) string {
	if v, ok := tags[key]; ok {
		return strings.TrimSpace(v)
	}
	for k, v := range tags {
		if strings.EqualFold(k, key) {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func first(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
