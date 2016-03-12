package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/imageagogo/process"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
	m := getImage("./src/github.com/imageagogo/test/testImage1.jpg")

	bounds := m.Bounds()

	// graph.Histogram(m)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := m.At(x, y)
			pixel = process.BitwiseMask(pixel, 0, 150)
			m.Set(x, y, pixel)
		}
	}

	drawImage(m)
}

func getImage(path string) *image.RGBA {
	// Read image from file
	reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	bounds := m.Bounds()

	rgbaImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgbaImg.Set(x, y, m.At(x, y))
		}
	}

	return rgbaImg
}

func drawImage(image image.Image) {
	filename := "image.jpg"

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	err = jpeg.Encode(f, image, &jpeg.Options{Quality: 100})
	if err != nil {
		println(err)
	}
}
