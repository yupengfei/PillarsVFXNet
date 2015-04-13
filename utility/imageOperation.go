package utility

import (
	"image"
	"image/png"
	"log"
	"os"
)

func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

func saveImage(path string, img image.Image) (err error) {
	imgfile, err := os.Create(path)
	defer imgfile.Close()
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
	return
}
