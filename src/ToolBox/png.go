package ToolBox

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

func GetPixels(path string) [][]Pixel {

	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(path)

	if err != nil {
			fmt.Println("Error: File could not be opened")
			os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixelsFromFile(file)

	if err != nil {
			fmt.Println("Error: Image could not be decoded")
			os.Exit(1)
	}

	return pixels
}

// Get the bi-dimensional pixel array
func getPixelsFromFile(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
			return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for x := 0; x < width; x++ {
			var row []Pixel
			for z := 0; z < height; z++ {
					row = append(row, rgbaToPixel(img.At(x, z).RGBA()))
			}
			pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}