package pixbuilder

import "testing"

func TestASG(t *testing.T) {

	pat := NewPatternRandom(150, 150)
	dmap := LoadGrayFrom("Symbol.png")

	asg := NewASG(pat, dmap)
	SaveAs(asg, "testasg")

}
