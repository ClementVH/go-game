package Loaders

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

func GetChunkPositions() []struct{X, Z int} {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("../res/zones/zone.png")

	if err != nil {
			fmt.Println("Error: File could not be opened")
			os.Exit(1)
	}

	defer file.Close()

	pixels, width, height, err := getPixels(file)

	if err != nil {
			fmt.Println("Error: Image could not be decoded")
			os.Exit(1)
	}

	positions := make([]struct{X, Z int}, 0, width * height)

	for z, column := range pixels {
		for x, pixel := range column {
			if pixel.A > 0 {
				positions = append(positions, struct{X, Z int}{x, z})
			}
		}
	}

	return positions
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, int, int, error) {
	img, _, err := image.Decode(file)

	if err != nil {
			return nil, 0, 0, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
			var row []Pixel
			for x := 0; x < width; x++ {
					row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
			}
			pixels = append(pixels, row)
	}

	return pixels, width, height, nil
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