package metadata

import (
	"encoding/json"
	"math"
	"os"
	"testing"
	"time"
)

// xmp as Lightroom and exiftool write it into an mp4, taken from the fixture's
// ffprobe output
func videoXmp(t *testing.T) XmpData {
	t.Helper()
	raw, err := os.ReadFile(AssetPath + "video/probe/xmp.mp4.json")
	if err != nil {
		t.Fatalf("could not read fixture: %v", err)
	}
	var probe struct {
		Format struct {
			Tags map[string]string `json:"tags"`
		} `json:"format"`
	}
	if err = json.Unmarshal(raw, &probe); err != nil {
		t.Fatalf("could not parse fixture: %v", err)
	}
	xd, err := NewXmpDataFromBytes([]byte(probe.Format.Tags["xmp"]))
	if err != nil {
		t.Fatalf("could not parse xmp: %v", err)
	}
	return xd
}

func TestXmpVideoFields(t *testing.T) {
	xd := videoXmp(t)
	if got := xd.GetTitle(); got != "Evening swim" {
		t.Errorf("title = %q", got)
	}
	if got := xd.GetDescription(); got != "Jumping off the pier" {
		t.Errorf("description = %q", got)
	}
	if got := xd.GetKeywords(); len(got) != 2 || got[0] != "summer" || got[1] != "archipelago" {
		t.Errorf("keywords = %v", got)
	}
	if mk, model := xd.GetCamera(); mk != "SONY" || model != "ILCE-7M4" {
		t.Errorf("camera = %q %q", mk, model)
	}
	want := time.Date(2023, 8, 12, 16, 40, 0, 0, time.UTC)
	if got := xd.GetDate(); !got.Equal(want) {
		t.Errorf("date = %v, want %v", got, want)
	}
	lat, long, ok := xd.GetLocation()
	if !ok || math.Abs(lat-59.8586) > 1e-4 || math.Abs(long-17.6389) > 1e-4 {
		t.Errorf("location = %v %v %v", lat, long, ok)
	}
}

func TestXmpEmptyFields(t *testing.T) {
	var xd XmpData
	if mk, model := xd.GetCamera(); mk != "" || model != "" {
		t.Errorf("camera from empty xmp = %q %q", mk, model)
	}
	if !xd.GetDate().IsZero() || xd.GetDescription() != "" {
		t.Errorf("expected zero values from empty xmp")
	}
	if _, _, ok := xd.GetLocation(); ok {
		t.Errorf("expected no location from empty xmp")
	}
}

func TestParseXmpCoordinate(t *testing.T) {
	tests := []struct {
		in   string
		want float64
		ok   bool
	}{
		{"59,51.516N", 59.8586, true},
		{"17,38,20.04E", 17.6389, true},
		{"33,52.2S", -33.87, true},
		{"118,14.4W", -118.24, true},
		{"59,51.516", 0, false},
		{"59N", 0, false},
		{"a,b N", 0, false},
		{"", 0, false},
	}
	for _, tc := range tests {
		got, ok := parseXmpCoordinate(tc.in)
		if ok != tc.ok || math.Abs(got-tc.want) > 1e-4 {
			t.Errorf("parseXmpCoordinate(%q) = %v, %v, want %v, %v", tc.in, got, ok, tc.want, tc.ok)
		}
	}
}

func TestParseXmpDate(t *testing.T) {
	cest := time.FixedZone("", 2*3600)
	tests := []struct {
		in   string
		want time.Time
	}{
		{"2023-08-12T18:40:00+02:00", time.Date(2023, 8, 12, 18, 40, 0, 0, cest)},
		{"2023-08-12T18:40:00.25Z", time.Date(2023, 8, 12, 18, 40, 0, 250e6, time.UTC)},
		{"2023-08-12T18:40:00", time.Date(2023, 8, 12, 18, 40, 0, 0, time.UTC)},
		{"2023-08-12T18:40+02:00", time.Date(2023, 8, 12, 18, 40, 0, 0, cest)},
		{"2023-08-12", time.Date(2023, 8, 12, 0, 0, 0, 0, time.UTC)},
	}
	for _, tc := range tests {
		got, ok := parseXmpDate(tc.in)
		if !ok || !got.Equal(tc.want) {
			t.Errorf("parseXmpDate(%q) = %v, %v, want %v", tc.in, got, ok, tc.want)
		}
	}
	if _, ok := parseXmpDate("yesterday"); ok {
		t.Errorf("expected garbage to be rejected")
	}
}
