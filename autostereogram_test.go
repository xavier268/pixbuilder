package pixbuilder

import "testing"

func TestASG(t *testing.T) {

	pat := NewPatternRandom1(150, 150)
	dmap := LoadGrayFrom("Symbol.png")

	asg := NewASG(pat, dmap)
	SaveAs(asg, "testasg1")

}

func TestASG2(t *testing.T) {
	pat := NewPatternRandom2(150, 150)
	dmap := LoadGrayFrom("SmileyOK.png")

	asg := NewASG(pat, dmap)
	SaveAs(asg, "testasg2")
}

func ExampleNewASG() {
	// Create a random pattern
	pat := NewPatternRandom2(200, 100)
	// Load a depth map from an existing png
	dmap := LoadGrayFrom("Symbol.png")
	// comput ethe result image
	asg := NewASG(pat, dmap)
	// save result as png file
	SaveAs(asg, "example")
	// Output:
}
