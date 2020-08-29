package pixbuilder

import "image"

// NewASG constructs a new AutoStereoGram from the provided pattern and the provided depth map.
func NewASG(pat *Pattern, dmap *image.Gray) image.Image {

	// l,L : distance to screen, distance between eyes (in pixels)
	l, L := 200, 200

	rgba := image.NewNRGBA(dmap.Bounds())
	for y := 0; y < rgba.Rect.Dy(); y++ {
		for x := 0; x < rgba.Rect.Dx(); x++ {
			// fill with pattern
			rgba.Set(x, y, pat.At(x, y))
		}
		for xz := 0; xz < rgba.Rect.Dx(); xz++ {
			// force right eye value to match left eye value
			z := int(dmap.GrayAt(xz, y).Y)
			dx := (z * L) / (z + l)
			x := xz - dx/2
			if x > 0 {
				c := rgba.At(x, y)
				rgba.Set(x+dx, y, c)
			}
		}
	}

	return rgba
}
