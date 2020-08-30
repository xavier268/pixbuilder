package pixbuilder

import "image"

// NewASG constructs a new AutoStereoGram from the provided pattern and the provided depth map.
// The size of the resulting image is the size of the depth map image.
func NewASG(pat *Pattern, dmap *image.Gray) image.Image {

	// l,L : distance to screen, distance between eyes (in pixels)
	l, L := 200, 200

	rgba := image.NewNRGBA(dmap.Bounds())

	for y := 0; y < rgba.Rect.Dy(); y++ {
		// y is vertical coordinate. No influence on stereo effect.

		for x := 0; x < rgba.Rect.Dx(); x++ {
			// first, fill result image with pattern
			rgba.Set(x, y, pat.At(x, y))
		}

		for xz := 0; xz < rgba.Rect.Dx(); xz++ {
			// xz is the depth map horizontal coordinate, z is the gray color, representing the depth.
			// larger, darker is closer ; smaller, whiter is far away.
			z := int(dmap.GrayAt(xz, y).Y)
			// when both eyes point to xz, the distance between the points where the rays intersect the image is 2*dx
			dx := (z * L) / (z + l)
			// this x-dx/2 is the left eye intersect
			x := xz - dx/2
			if x > 0 {
				// get pixel from the left eye intersect ...
				c := rgba.At(x, y)
				// ... and copy it in the right eye pixel intersect, x+dx
				rgba.Set(x+dx, y, c)
			}
		}
	}

	return rgba
}
