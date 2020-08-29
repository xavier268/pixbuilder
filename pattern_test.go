package pixbuilder

import (
	"image"
	"testing"
)

// Check that Pattern is an Image ?
var _ image.Image = new(Pattern)

func TestConstructPatternRandom(t *testing.T) {
	p := NewPatternRandom1(50, 200)
	SaveAs(p, "testpatternrandom1")
}
func TestConstructPatternRandom2(t *testing.T) {
	p := NewPatternRandom2(50, 200)
	SaveAs(p, "testpatternrandom2")
}

func TestConstructPatternImage(t *testing.T) {
	im := LoadFrom("Symbol.png")
	p := NewPatternImage(im)
	SaveAs(p, "testpatternimage")
}
