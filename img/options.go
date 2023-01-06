package img

import "image"

func (o Options) Rectangle() image.Rectangle {
	return image.Rect(o.X, o.Y, o.X+o.Width, o.Y+o.Height)
}
