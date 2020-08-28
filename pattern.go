package pixbuilder

import (
	"image"
	"image/color"
	"math/rand"
)

// Pattern provide a color for each point, and repeats itself horizontally and vertically.
// A Pattern is also an image.Image.
type Pattern struct {
	xperiod, yperiod int // The spatial periods
	image.Image          // The image that is repeated
}

// At provides color pixel for any coordinates, extending underlying At, and using provided period.
func (p *Pattern) At(x, y int) color.Color {
	return p.Image.At(x%p.xperiod, y%p.yperiod)
}

// NewPatternImage generate a Pattern from the underlying image.
func NewPatternImage(im image.Image) *Pattern {
	if im == nil {
		panic("Image expected to create Pattern")
	}
	p := new(Pattern)
	p.Image = im
	b := im.Bounds()
	p.xperiod, p.yperiod = b.Dx(), b.Dy()
	return p
}

// NewPatternRandom generates a Pattern of given period.
func NewPatternRandom(x, y int) *Pattern {
	p := new(Pattern)
	p.xperiod, p.yperiod = x, y
	rgba := image.NewRGBA(image.Rect(0, 0, x, y))

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			c := color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 0xFF,
			}
			c.R = uint8(rand.Intn(0x10000))
			c.G = uint8(rand.Intn(0x10000))
			c.B = uint8(rand.Intn(0x10000))
			rgba.Set(i, j, c)
		}
	}
	p.Image = rgba
	return p
}
