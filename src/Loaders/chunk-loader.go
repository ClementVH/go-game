package Loaders

import (
	"go-game/src/Entities"
	"go-game/src/ToolBox"
	"reflect"
)

func GetChunkZones() [][]Entities.ChunkPosition {

	pixels := ToolBox.GetPixels("../res/zones/zone.png")

	black := ToolBox.Pixel{R: 0, G: 0, B: 0, A: 255}
	red := ToolBox.Pixel{R: 255, G: 0, B: 0, A: 255}
	green := ToolBox.Pixel{R: 0, G: 255, B: 0, A: 255}
	blue := ToolBox.Pixel{R: 0, G: 0, B: 255, A: 255}

	positions := make([][]Entities.ChunkPosition, 4)
	for i := range positions {
		positions[i] = make([]Entities.ChunkPosition, 0)
	}

	for x, row := range pixels {
		for z, pixel := range row {
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