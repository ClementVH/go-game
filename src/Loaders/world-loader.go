package Loaders

import (
	"fmt"
	"go-game/src/Entities"
	"image"
	"image/png"
	"io"
	"os"
	"reflect"
)

func GetZones() [][]Entities.ChunkPosition {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("../res/zones/zone.png")

	if err != nil {
			fmt.Println("Error: File could not be opened")
			os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixels(file)

	if err != nil {
			fmt.Println("Error: Image could not be decoded")
			os.Exit(1)
	}

	black := Pixel{0,0,0,255}
	red := Pixel{255,0,0,255}
	green := Pixel{0,255,0,255}
	blue := Pixel{0,0,255,255}

	positions := make([][]Entities.ChunkPosition, 4)
	for i := range positions {
		positions[i] = make([]Entities.ChunkPosition, 0)
	}

	for z, column := range pixels {
		for x, pixel := range column {
			if reflect.DeepEqual(pixel, black) {
				positions[0] = append(positions[0], Entities.ChunkPosition{X: x, Z: z})
			}
			if reflect.DeepEqual(pixel, red) {
				positions[1] = append(positions[1], Entities.ChunkPosition{X: x, Z: z})
			}
			if reflect.DeepEqual(pixel, green) {
				positions[2] = append(positions[2], Entities.ChunkPosition{X: x, Z: z})
			}
			if reflect.DeepEqual(pixel, blue) {
				positions[3] = append(positions[3], Entities.ChunkPosition{X: x, Z: z})
			}
		}
	}

	return positions
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
			return nil, err
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