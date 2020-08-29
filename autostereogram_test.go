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
	dmap := LoadGrayFrom("Symbol.png")

	asg := NewASG(pat, dmap)
	SaveAs(asg, "testasg2")

}
