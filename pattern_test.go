package pixbuilder

import (
	"image"
	"testing"
)

// Check that Pattern is an Image ?
var _ image.Image = new(Pattern)

func TestConstructPatternRandom(t *testing.T) {
	p := NewPatternRandom(50, 200)
	SaveAs(p, "testpatternrandom")
}

func TestConstructPatternImage(t *testing.T) {
	im := LoadFrom("Symbol.png")
	p := NewPatternImage(im)
	SaveAs(p, "testpatternimage")
}
