package pixbuilder

import (
	"image"
	// Loading various image engines
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

// SaveAs and image to file
func SaveAs(im image.Image, name string) {
	w, err := os.Create(name + ".png")
	check(err)
	defer w.Close()
	err = png.Encode(w, im)
	check(err)
}

// LoadFrom loads image from file name (with extension)
func LoadFrom(name string) image.Image {
	r, err := os.Open(name)
	check(err)
	defer r.Close()
	im, typ, err := image.Decode(r)
	check(err)
	log.Printf("Juste loaded image %s of type %s\n", name, typ)
	return im
}

// check and panic
func check(err error) {
	if err != nil {
		panic(err)
	}
}
