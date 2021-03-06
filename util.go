package pixbuilder

// Various utilities for save/load

import (
	"image"
	// Loading various image engines, add more if needed...
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

// SaveAs image to file. Do not provide extension, a ".png" is added.
func SaveAs(im image.Image, name string) {
	w, err := os.Create(name + ".png")
	check(err)
	defer w.Close()
	err = png.Encode(w, im)
	check(err)
}

// LoadFrom image from file name. File name extension is expected.
func LoadFrom(name string) image.Image {
	r, err := os.Open(name)
	check(err)
	defer r.Close()
	im, typ, err := image.Decode(r)
	check(err)
	log.Printf("Just loaded image %s of type %s\n", name, typ)
	return im
}

// LoadGrayFrom loads ANY image, and convert it to a Gray image.
// File name extension is expected.
func LoadGrayFrom(name string) *image.Gray {
	img := LoadFrom(name)

	// Converting image to grayscale
	grayImg := image.NewGray(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	return grayImg
}

// check and panic
func check(err error) {
	if err != nil {
		panic(err)
	}
}
