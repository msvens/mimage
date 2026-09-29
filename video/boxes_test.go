package video

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// boxOf builds a box with a 32 bit size and payload bytes of content
func boxOf(boxType string, payload int) []byte {
	b := binary.BigEndian.AppendUint32(nil, uint32(8+payload))
	b = append(b, boxType...)
	return append(b, make([]byte, payload)...)
}

// largeBox builds a box with a 64 bit size
func largeBox(boxType string, payload int) []byte {
	b := binary.BigEndian.AppendUint32(nil, 1)
	b = append(b, boxType...)
	b = binary.BigEndian.AppendUint64(b, uint64(16+payload))
	return append(b, make([]byte, payload)...)
}

func join(parts ...[]byte) []byte {
	return bytes.Join(parts, nil)
}

func TestCheckBoxes(t *testing.T) {
	ftyp := boxOf("ftyp", 16)
	moov := boxOf("moov", 100)
	mdat := boxOf("mdat", 1000)
	toEnd := append(binary.BigEndian.AppendUint32(nil, 0), "mdat"...)

	tests := []struct {
		name      string
		data      []byte
		truncated bool
	}{
		{"complete", join(ftyp, moov, mdat), false},
		{"64 bit mdat", join(ftyp, moov, largeBox("mdat", 1000)), false},
		{"mdat running to end of file", join(ftyp, moov, toEnd, make([]byte, 500)), false},
		{"padding after the last box", join(ftyp, moov, mdat, []byte{0, 0, 0}), false},
		{"media cut short", join(ftyp, moov, mdat)[:len(ftyp)+len(moov)+600], true},
		{"index cut short", join(ftyp, moov)[:len(ftyp)+50], true},
		{"64 bit mdat cut short", join(ftyp, moov, largeBox("mdat", 1000))[:len(ftyp)+len(moov)+500], true},
		{"64 bit size itself cut off", join(ftyp, largeBox("mdat", 10))[:len(ftyp)+12], true},
		//not truncation, so left for ffprobe to reject
		{"impossible box size", join(ftyp, binary.BigEndian.AppendUint32(nil, 3), []byte("junk"), mdat), false},
	}
	for _, tc := range tests {
		err := checkBoxes(bytes.NewReader(tc.data), int64(len(tc.data)))
		if tc.truncated != errors.Is(err, ErrTruncated) {
			t.Errorf("%s: got %v, want truncated %v", tc.name, err, tc.truncated)
		}
		if !tc.truncated && err != nil {
			t.Errorf("%s: unexpected error %v", tc.name, err)
		}
	}
}

// every fixture is a complete file, including the QuickTime one
func TestCheckCompleteFixtures(t *testing.T) {
	for _, e := range expectations {
		if err := checkComplete(assetPath + e.file); err != nil {
			t.Errorf("%s: %v", e.file, err)
		}
	}
}

// An export with its index at the front still probes as a whole video after
// losing its tail. ffprobe cannot tell, the box check can
func TestProbeTruncated(t *testing.T) {
	plain, err := os.ReadFile(assetPath + "plain.mp4")
	if err != nil {
		t.Fatalf("could not read fixture: %v", err)
	}
	cut := filepath.Join(t.TempDir(), "cut.mp4")
	if err = os.WriteFile(cut, plain[:len(plain)*2/3], 0644); err != nil {
		t.Fatalf("could not write: %v", err)
	}
	if err = checkComplete(cut); !errors.Is(err, ErrTruncated) {
		t.Fatalf("checkComplete: expected ErrTruncated, got %v", err)
	}

	requireFFmpeg(t)
	//show what the check guards against: ffprobe alone accepts the file
	out, err := run(testContext(t), FFprobePath, "-v", "error", "-print_format", "json",
		"-show_format", "-show_streams", cut)
	if err != nil {
		t.Fatalf("ffprobe rejected the cut file itself, the check may no longer be needed: %v", err)
	}
	if s, err := parseProbe(out); err != nil || s.Duration == 0 {
		t.Fatalf("ffprobe no longer reports a cut file as whole (%v, %v): revisit checkComplete", s, err)
	}
	if _, err = Probe(testContext(t), cut); !errors.Is(err, ErrTruncated) {
		t.Errorf("Probe: expected ErrTruncated, got %v", err)
	}
	if _, err = ExtractPoster(testContext(t), cut, filepath.Join(t.TempDir(), "p"), PosterOptions{}); !errors.Is(err, ErrTruncated) {
		t.Errorf("ExtractPoster: expected ErrTruncated, got %v", err)
	}
}
