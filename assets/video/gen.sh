#!/usr/bin/env bash
#
# Regenerates the synthetic video fixtures in this directory, and the ffprobe
# output captured from them in probe/. The fixtures are committed, so tests never
# run this: it documents how they were made and recreates them if they need to
# change. Real camera and Lightroom exports live next to them and are not
# produced here.
#
# Needs ffmpeg 7 or later (for -display_rotation) and exiftool (for xmp).
#
#   cd assets/video && ./gen.sh

set -euo pipefail
cd "$(dirname "$0")"

ff() { ffmpeg -hide_banner -loglevel error -y "$@"; }

# small, low frame rate h.264 with the flags a web export uses. Solid colours
# compress to almost nothing, which keeps the fixtures a few KB
h264=(-c:v libx264 -preset veryslow -crf 30 -pix_fmt yuv420p -r 10 -movflags +faststart)
aac=(-c:a aac -b:a 32k -ac 1)
silence=(-f lavfi -i "anullsrc=r=22050:cl=mono")

# 2s 320x240 with the tags an iPhone writes, as QuickTime mdta keys
ff -f lavfi -i "testsrc2=size=320x240:rate=10:duration=2" "${silence[@]}" -shortest \
  "${h264[@]}" "${aac[@]}" -movflags +faststart+use_metadata_tags \
  -metadata title="Plain fixture" \
  -metadata creation_time="2024-06-01T10:15:30Z" \
  -metadata com.apple.quicktime.creationdate="2024-06-01T12:15:30+0200" \
  -metadata com.apple.quicktime.make="Apple" \
  -metadata com.apple.quicktime.model="iPhone 15 Pro" \
  -metadata com.apple.quicktime.location.ISO6709="+59.3293+018.0686+012.345/" \
  plain.mp4

# the same as a QuickTime movie, with no metadata at all
ff -f lavfi -i "testsrc2=size=320x240:rate=10:duration=2" "${silence[@]}" -shortest \
  "${h264[@]}" "${aac[@]}" -map_metadata -1 -f mov plain.mov

# landscape pixels, top half red and bottom half blue, stored with a display
# matrix that turns them to portrait. This is how a phone records portrait
# video: -90 is what ffprobe reports for an iPhone held upright
ff -f lavfi -i "color=red:size=320x120:rate=10:duration=2[t];color=blue:size=320x120:rate=10:duration=2[b];[t][b]vstack" \
  "${h264[@]}" -an rotated.tmp.mp4
ff -display_rotation:v:0 -90 -i rotated.tmp.mp4 -c copy -movflags +faststart rotated.mp4
rm rotated.tmp.mp4

# half a second of black before any picture, the case the poster offset avoids
ff -f lavfi -i "color=black:size=320x240:rate=10:duration=0.5[a];color=0x33cc33:size=320x240:rate=10:duration=2[b];[a][b]concat=n=2:v=1:a=0" \
  "${h264[@]}" -an blackstart.mp4

# shorter than the default poster offset
ff -f lavfi -i "color=0x33cc33:size=320x240:rate=10:duration=0.4" "${h264[@]}" -an short.mp4

# an mp4 with sound and no picture: detected as mp4, rejected by Probe
ff "${silence[@]}" -t 1 "${aac[@]}" -f mp4 audioonly.mp4

# sound with cover art: the only picture is an attached_pic, which is not video
ff "${silence[@]}" -f lavfi -i "color=red:size=64x64:duration=0.1" -t 1 \
  -map 0 -map 1 "${aac[@]}" -c:v mjpeg -frames:v 1 -disposition:v:0 attached_pic coverart.mp4

# a Lightroom style export: no QuickTime keys, title, caption, keywords and
# capture details in an xmp packet
ff -f lavfi -i "testsrc2=size=320x240:rate=10:duration=2" "${h264[@]}" -an \
  -map_metadata -1 -metadata creation_time="2023-08-12T16:40:00Z" xmp.mp4
exiftool -q -overwrite_original \
  -XMP-dc:Title="Evening swim" \
  -XMP-dc:Description="Jumping off the pier" \
  -XMP-dc:Subject="summer" -XMP-dc:Subject="archipelago" \
  -XMP-xmp:CreateDate="2023-08-12T18:40:00+02:00" \
  -XMP-tiff:Make="SONY" -XMP-tiff:Model="ILCE-7M4" \
  -XMP-exif:GPSLatitude="59.8586 N" -XMP-exif:GPSLongitude="17.6389 E" \
  xmp.mp4

# --- sources for Transcode, one per thing it has to fix ---

# like an Android phone original: HEVC, recorded upright (landscape frames plus
# a display matrix), the media data before the index, and location switched off,
# which Samsung writes as 0,0. Red on top of the stored frame, so an upright
# result has red on the right
ff -f lavfi -i "color=red:size=320x120:rate=10:duration=2[t];color=blue:size=320x120:rate=10:duration=2[b];[t][b]vstack" \
  "${silence[@]}" -shortest -c:v libx265 -x265-params log-level=error -crf 30 -tag:v hvc1 -pix_fmt yuv420p \
  "${aac[@]}" samsung.tmp.mp4
ff -display_rotation:v:0 -90 -i samsung.tmp.mp4 -map 0 -c copy -map_metadata -1 \
  -movflags +use_metadata_tags \
  -metadata creation_time="2024-07-14T12:04:06Z" \
  -metadata location="+00.0000+000.0000/" \
  -metadata com.android.version="14" \
  samsung-like.mp4
rm samsung.tmp.mp4

# interlaced, as DV and early HD camcorders recorded
ff -f lavfi -i "testsrc2=size=320x240:rate=10:duration=2" -c:v libx264 -crf 30 -pix_fmt yuv420p \
  -flags +ildct+ilme -x264-params tff=1 -an interlaced.mp4

# non square pixels: 240x240 stored, 4:3 pixels, shown as 320x240
ff -f lavfi -i "testsrc2=size=240x240:rate=10:duration=2" -vf setsar=4/3 "${h264[@]}" -an anamorphic.mp4

# HDR as a recent iPhone records it: 10 bit HEVC, bt2020, HLG
ff -f lavfi -i "testsrc2=size=320x240:rate=10:duration=2" \
  -c:v libx265 -x265-params log-level=error:colorprim=bt2020:transfer=arib-std-b67:colormatrix=bt2020nc \
  -crf 30 -pix_fmt yuv420p10le -tag:v hvc1 \
  -color_primaries bt2020 -color_trc arib-std-b67 -colorspace bt2020nc -an hdr.mp4

# an old digicam movie: motion jpeg and uncompressed sound in avi
ff -f lavfi -i "testsrc2=size=160x120:rate=10:duration=1" -f lavfi -i "anullsrc=r=8000:cl=mono" -shortest \
  -c:v mjpeg -q:v 20 -c:a pcm_u8 -f avi old.avi

# what ffprobe reports for each, the input to the parsing tests
mkdir -p probe
for f in plain.mp4 plain.mov rotated.mp4 blackstart.mp4 short.mp4 audioonly.mp4 coverart.mp4 xmp.mp4 \
  samsung-like.mp4 interlaced.mp4 anamorphic.mp4 hdr.mp4 old.avi; do
  ffprobe -v error -print_format json -show_format -show_streams -export_xmp 1 "$f" \
    | sed "s|\"filename\": \".*\"|\"filename\": \"$f\"|" > "probe/${f%.*}.${f##*.}.json"
done

# ffmpeg 4 and earlier reported rotation as a clockwise "rotate" tag rather than
# display matrix side data. There is no old ffprobe to capture that from, so it
# is derived from the rotated capture
python3 - <<'PY'
import json
d = json.load(open("probe/rotated.mp4.json"))
v = d["streams"][0]
del v["side_data_list"]
v.setdefault("tags", {})["rotate"] = "90"
json.dump(d, open("probe/rotated-legacy.mp4.json", "w"), indent=4)
PY

ls -l *.mp4 *.mov *.avi
