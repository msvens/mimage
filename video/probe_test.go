package video

import (
	"errors"
	"math"
	"os"
	"slices"
	"testing"
	"time"
)

const assetPath = "../assets/video/"

// what the fixtures hold, see assets/video/gen.sh. Shared by the parsing tests,
// which read captured ffprobe output, and the integration tests, which run the
// installed ffprobe
type expectation struct {
	file     string
	err      error
	duration time.Duration
	width    int
	height   int
	stored   [2]int
	rotation int
	created  time.Time
	make     string
	model    string
	title    string
	desc     string
	keywords []string
	location *Location
	codecs   [2]string
	//zero means the 10 fps every synthetic fixture uses
	fps float64
}

var cest = time.FixedZone("", 2*3600)

var expectations = []expectation{
	{
		file: "plain.mp4", duration: 2 * time.Second, width: 320, height: 240, stored: [2]int{320, 240},
		created: time.Date(2024, 6, 1, 12, 15, 30, 0, cest), make: "Apple", model: "iPhone 15 Pro",
		title: "Plain fixture", location: &Location{59.3293, 18.0686, 12.345}, codecs: [2]string{"h264", "aac"},
	},
	{
		file: "plain.mov", duration: 2 * time.Second, width: 320, height: 240, stored: [2]int{320, 240},
		codecs: [2]string{"h264", "aac"},
	},
	{
		file: "rotated.mp4", duration: 2 * time.Second, width: 240, height: 320, stored: [2]int{320, 240},
		rotation: 90, codecs: [2]string{"h264", ""},
	},
	{
		file: "blackstart.mp4", duration: 2500 * time.Millisecond, width: 320, height: 240,
		stored: [2]int{320, 240}, codecs: [2]string{"h264", ""},
	},
	{
		file: "short.mp4", duration: 400 * time.Millisecond, width: 320, height: 240,
		stored: [2]int{320, 240}, codecs: [2]string{"h264", ""},
	},
	{
		//the xmp date wins over the container's UTC creation_time: the same
		//instant, but it keeps the local offset
		file: "xmp.mp4", duration: 2 * time.Second, width: 320, height: 240, stored: [2]int{320, 240},
		created: time.Date(2023, 8, 12, 18, 40, 0, 0, cest), make: "SONY", model: "ILCE-7M4",
		title: "Evening swim", desc: "Jumping off the pier", keywords: []string{"summer", "archipelago"},
		location: &Location{Latitude: 59.8586, Longitude: 17.6389}, codecs: [2]string{"h264", ""},
	},
	{file: "audioonly.mp4", err: ErrNoVideoStream},
	{file: "coverart.mp4", err: ErrNoVideoStream},
}

// real recordings, of which only the ffprobe output is committed. They are
// checked by the parsing tests alone
var realExpectations = []expectation{
	{
		//a Galaxy S23 original: HEVC, recorded upright, so landscape frames
		//turned for display. Location was off, which Samsung writes as 0,0
		file: "samsung-s23.mp4", duration: 25698944 * time.Microsecond, width: 1080, height: 1920,
		stored: [2]int{1920, 1080}, rotation: 90, created: time.Date(2024, 7, 14, 12, 4, 6, 0, time.UTC),
		codecs: [2]string{"hevc", "aac"}, fps: 30.007,
	},
	{
		//named .mov but an mp4 by content, already H.264 and with no date
		file: "landscape.mp4", duration: 30100 * time.Millisecond, width: 1280, height: 720,
		stored: [2]int{1280, 720}, codecs: [2]string{"h264", "aac"}, fps: 30,
	},
}

func checkSummary(t *testing.T, e expectation, s *Summary, err error) {
	t.Helper()
	if e.err != nil {
		if !errors.Is(err, e.err) {
			t.Errorf("%s: expected %v, got %v", e.file, e.err, err)
		}
		return
	}
	if err != nil {
		t.Fatalf("%s: unexpected error %v", e.file, err)
	}
	//container timestamps are not exact to the millisecond
	if d := s.Duration - e.duration; d < -50*time.Millisecond || d > 50*time.Millisecond {
		t.Errorf("%s: duration = %v, want %v", e.file, s.Duration, e.duration)
	}
	if s.Width != e.width || s.Height != e.height {
		t.Errorf("%s: display size = %dx%d, want %dx%d", e.file, s.Width, s.Height, e.width, e.height)
	}
	if s.StoredWidth != e.stored[0] || s.StoredHeight != e.stored[1] {
		t.Errorf("%s: stored size = %dx%d, want %dx%d", e.file, s.StoredWidth, s.StoredHeight, e.stored[0], e.stored[1])
	}
	if s.Rotation != e.rotation {
		t.Errorf("%s: rotation = %d, want %d", e.file, s.Rotation, e.rotation)
	}
	if !s.CreationTime.Equal(e.created) {
		t.Errorf("%s: creation time = %v, want %v", e.file, s.CreationTime, e.created)
	}
	if !e.created.IsZero() {
		_, gotOffset := s.CreationTime.Zone()
		_, wantOffset := e.created.Zone()
		if gotOffset != wantOffset {
			t.Errorf("%s: creation time offset = %d, want %d", e.file, gotOffset, wantOffset)
		}
	}
	if s.CameraMake != e.make || s.CameraModel != e.model {
		t.Errorf("%s: camera = %q %q, want %q %q", e.file, s.CameraMake, s.CameraModel, e.make, e.model)
	}
	if s.Title != e.title || s.Description != e.desc {
		t.Errorf("%s: title/description = %q %q, want %q %q", e.file, s.Title, s.Description, e.title, e.desc)
	}
	if !slices.Equal(s.Keywords, e.keywords) {
		t.Errorf("%s: keywords = %v, want %v", e.file, s.Keywords, e.keywords)
	}
	if !sameLocation(s.Location, e.location) {
		t.Errorf("%s: location = %v, want %v", e.file, s.Location, e.location)
	}
	if s.VideoCodec != e.codecs[0] || s.AudioCodec != e.codecs[1] {
		t.Errorf("%s: codecs = %q %q, want %q %q", e.file, s.VideoCodec, s.AudioCodec, e.codecs[0], e.codecs[1])
	}
	wantFps := e.fps
	if wantFps == 0 {
		wantFps = 10
	}
	if math.Abs(s.FrameRate-wantFps) > 0.01 {
		t.Errorf("%s: frame rate = %v, want %v", e.file, s.FrameRate, wantFps)
	}
}

func sameLocation(a, b *Location) bool {
	if a == nil || b == nil {
		return a == b
	}
	near := func(x, y float64) bool { return math.Abs(x-y) < 1e-4 }
	return near(a.Latitude, b.Latitude) && near(a.Longitude, b.Longitude) && near(a.Altitude, b.Altitude)
}

func readProbe(t *testing.T, file string) []byte {
	t.Helper()
	data, err := os.ReadFile(assetPath + "probe/" + file + ".json")
	if err != nil {
		t.Fatalf("could not read fixture: %v", err)
	}
	return data
}

func TestParseProbe(t *testing.T) {
	for _, e := range append(expectations, realExpectations...) {
		s, err := parseProbe(readProbe(t, e.file))
		checkSummary(t, e, s, err)
	}
}

// ffmpeg 4 and earlier gave rotation as a tag rather than a display matrix
func TestParseProbeLegacyRotation(t *testing.T) {
	s, err := parseProbe(readProbe(t, "rotated-legacy.mp4"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if s.Rotation != 90 || s.Width != 240 || s.Height != 320 {
		t.Errorf("legacy rotation: got %d at %dx%d, want 90 at 240x320", s.Rotation, s.Width, s.Height)
	}
}

// the cases no fixture produces, as minimal ffprobe output
func TestParseProbeEdgeCases(t *testing.T) {
	stream := func(extra string) []byte {
		return []byte(`{"streams":[{"index":0,"codec_type":"video","codec_name":"h264","width":720,"height":480` +
			extra + `}],"format":{"duration":"1.5"}}`)
	}

	s, err := parseProbe(stream(`,"sample_aspect_ratio":"32:27"`))
	if err != nil || s.Width != 853 || s.Height != 480 {
		t.Errorf("anamorphic: got %dx%d, %v, want 853x480", s.Width, s.Height, err)
	}
	s, err = parseProbe(stream(`,"side_data_list":[{"side_data_type":"Display Matrix","rotation":180}]`))
	if err != nil || s.Rotation != 180 || s.Width != 720 {
		t.Errorf("upside down: got %d at %dx%d, %v", s.Rotation, s.Width, s.Height, err)
	}
	s, err = parseProbe(stream(`,"side_data_list":[{"side_data_type":"Display Matrix","rotation":90}]`))
	if err != nil || s.Rotation != 270 || s.Width != 480 {
		t.Errorf("counter clockwise: got %d at %dx%d, %v", s.Rotation, s.Width, s.Height, err)
	}
	//an unset container time is the 1904 epoch, not a real date
	s, err = parseProbe(stream(`,"tags":{"creation_time":"1904-01-01T00:00:00.000000Z"}`))
	if err != nil || !s.CreationTime.IsZero() {
		t.Errorf("epoch creation time: got %v, %v, want zero", s.CreationTime, err)
	}
	s, err = parseProbe(stream(`,"tags":{"creation_time":"2021-03-04T05:06:07.000000Z"}`))
	if want := time.Date(2021, 3, 4, 5, 6, 7, 0, time.UTC); err != nil || !s.CreationTime.Equal(want) {
		t.Errorf("stream creation time: got %v, %v, want %v", s.CreationTime, err, want)
	}
	//missing everything optional is fine
	s, err = parseProbe([]byte(`{"streams":[{"index":3,"codec_type":"video"}]}`))
	if err != nil || s.streamIndex != 3 || s.Duration != 0 || s.Location != nil {
		t.Errorf("bare stream: got %+v, %v", s, err)
	}
	if _, err = parseProbe([]byte(`{"streams":`)); err == nil {
		t.Errorf("expected an error for truncated json")
	}
	if _, err = parseProbe([]byte(`{}`)); !errors.Is(err, ErrNoVideoStream) {
		t.Errorf("no streams: expected ErrNoVideoStream, got %v", err)
	}
}

func TestParseISO6709(t *testing.T) {
	tests := []struct {
		in   string
		want *Location
	}{
		{"+59.3293+018.0686+012.345/", &Location{59.3293, 18.0686, 12.345}},
		{"+59.3293+018.0686/", &Location{Latitude: 59.3293, Longitude: 18.0686}},
		{"-33.8688+151.2093-002.000/", &Location{-33.8688, 151.2093, -2}},
		{"+40.7128-074.0060", &Location{Latitude: 40.7128, Longitude: -74.006}},
		{"+91.0000+018.0000/", nil},
		{"59.3293+018.0686/", nil},
		{"+59.3293/", nil},
		{"", nil},
	}
	for _, tc := range tests {
		got, ok := parseISO6709(tc.in)
		if ok != (tc.want != nil) || !sameLocation(got, tc.want) {
			t.Errorf("parseISO6709(%q) = %v, %v, want %v", tc.in, got, ok, tc.want)
		}
	}
}

func TestRatio(t *testing.T) {
	tests := map[string]float64{"10/1": 10, "30000/1001": 29.97003, "1:1": 1, "0/0": 0, "": 0, "abc": 0}
	for in, want := range tests {
		if got := ratio(in); math.Abs(got-want) > 1e-4 {
			t.Errorf("ratio(%q) = %v, want %v", in, got, want)
		}
	}
}
