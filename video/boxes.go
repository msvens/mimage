package video

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// checkComplete walks the top level boxes of an mp4 or mov and returns an error
// wrapping ErrTruncated when one claims more bytes than the file holds.
//
// ffprobe cannot be relied on for this. It stops quietly at end of file, so a
// web export, whose index sits at the front, still probes as a whole video after
// losing its tail. Only box headers are read, never the media.
//
// Anything else odd about the layout is left for ffprobe to judge: this looks
// for one thing, a file shorter than it says it is
func checkComplete(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	info, err := f.Stat()
	if err != nil {
		return err
	}
	return checkBoxes(f, info.Size())
}

func checkBoxes(r io.ReaderAt, fileSize int64) error {
	header := make([]byte, 16)
	for offset := int64(0); fileSize-offset >= 8; {
		if _, err := r.ReadAt(header[:8], offset); err != nil {
			return err
		}
		size := int64(binary.BigEndian.Uint32(header[:4]))
		switch size {
		case 0:
			//the last box, running to the end of the file whatever its length
			return nil
		case 1:
			//a 64 bit size follows the type
			if fileSize-offset < 16 {
				return fmt.Errorf("%w: box header cut off at %d", ErrTruncated, offset)
			}
			if _, err := r.ReadAt(header[8:16], offset+8); err != nil {
				return err
			}
			size = int64(binary.BigEndian.Uint64(header[8:16]))
			if size < 16 {
				return nil
			}
		default:
			if size < 8 {
				return nil
			}
		}
		if size > fileSize-offset {
			return fmt.Errorf("%w: %q box at %d needs %d bytes, the file has %d",
				ErrTruncated, header[4:8], offset, size, fileSize-offset)
		}
		offset += size
	}
	//fewer than 8 bytes left over is padding, not a box
	return nil
}
