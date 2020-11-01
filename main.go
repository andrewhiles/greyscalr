package main

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"path"
	"strings"

	"github.com/kyokomi/emoji"
)

func main() {
	fileName := os.Args[1]
	infile, err := os.Open(fileName)

	if err != nil {
		emoji.Println(":angry:")
		log.Printf("Failed to open file (%s). Error: %s", fileName, err)
		panic(err.Error())
	}
	fileExtension := path.Ext(fileName)
	if fileExtension != ".png" {
		log.Println("got in here...")
		panic(errors.New("Image format is not supported! Must be a .png file!"))
	}
	defer infile.Close()

	imgSrc, _, err := image.Decode(infile)
	if err != nil {
		panic(err.Error())
	}

	imgBounds := imgSrc.Bounds()
	w, h := imgBounds.Max.X, imgBounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := imgSrc.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	filenameWithoutExtension := strings.Split(fileName, fileExtension)[0]
	newFileName := filenameWithoutExtension + "-grayscale.png"

	newFile, err := os.Create(newFileName)
	if err != nil {
		log.Printf("Failed creating file %s. Error: %s", newFile, err)
		panic(err.Error())
	}
	defer newFile.Close()
	png.Encode(newFile, grayScale)
}
