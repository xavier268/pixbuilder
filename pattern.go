package pixbuilder

import (
	"image"
	"image/color"
	"math/rand"
)

// Pattern provide a color for each point, and repeats itself horizontally and vertically.
// A Pattern is also an image.Image.
type Pattern struct {
	xperiod, yperiod int // Horizontal and vertical periods
	image.Image          // The image that is repeated
}

// At provides color pixel for any coordinates, extending underlying At, and using provided period.
func (p *Pattern) At(x, y int) color.Color {
	return p.Image.At(x%p.xperiod, y%p.yperiod)
}

// NewPatternImage generate a Pattern from the provided image.
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

// NewPatternRandom1 generates a random Pattern of given period.
func NewPatternRandom1(x, y int) *Pattern {
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

// NewPatternRandom2 generate another random pattern of given period.
func NewPatternRandom2(x, y int) *Pattern {
	p := new(Pattern)
	p.xperiod, p.yperiod = x, y
	rgba := image.NewRGBA(image.Rect(0, 0, x, y))
	c := color.RGBA{
		R: 0xFF,
		G: 0xFF,
		B: 0xFF,
		A: 0xFF,
	}
	// Fill with white
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			rgba.Set(i, j, c)
		}
	}

	// draw "random walk" of random colors.
	// The number of "walks" are proportionnate to the surface, ensuring efficent coverage.
	for k := 0; k < p.xperiod*p.yperiod/10; k++ {
		c.R = uint8(rand.Intn(0x100))
		c.G = uint8(rand.Intn(0x100))
		c.B = uint8(rand.Intn(0x100))
		x, y := rand.Intn(p.xperiod), rand.Intn(p.yperiod)
		for l := 0; l < 300; l++ {
			rgba.Set(x, y, c)
			dx, dy := rand.Intn(5)-2, rand.Intn(5)-2
			x += p.xperiod
			y += p.yperiod
			x, y = (x+dx)%p.xperiod, (y+dy)%p.yperiod
		}
	}
	p.Image = rgba
	return p
}
