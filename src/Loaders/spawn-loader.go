package Loaders

import (
	"go-game/src/Entities"
	"go-game/src/ToolBox"
	"reflect"
)

func GetSpawnZones(zones [][]Entities.ChunkPosition) [][]Entities.MonsterPosition {

	positions := make([][]Entities.MonsterPosition, 4)
	for i := range positions {
		positions[i] = make([]Entities.MonsterPosition, 0)
	}

	pixels := ToolBox.GetPixels("../res/zones/spawn.png")

	groupType1 := ToolBox.Pixel{R: 25, G: 25, B: 25, A: 255}
	groupType2 := ToolBox.Pixel{R: 50, G: 50, B: 50, A: 255}
	groupType3 := ToolBox.Pixel{R: 75, G: 75, B: 75, A: 255}

	for zoneIndex, zone := range zones {
		for _, ChunkPosition := range zone {
			if reflect.DeepEqual(pixels[ChunkPosition.X][ChunkPosition.Z], groupType1) {
				positions[zoneIndex] = append(
					positions[zoneIndex],
					Entities.MonsterPosition(ChunkPosition),
				)
			}
			if reflect.DeepEqual(pixels[ChunkPosition.X][ChunkPosition.Z], groupType2) {
				positions[zoneIndex] = append(
					positions[zoneIndex],
					Entities.MonsterPosition(ChunkPosition),
				)
			}
			if reflect.DeepEqual(pixels[ChunkPosition.X][ChunkPosition.Z], groupType3) {
				positions[zoneIndex] = append(
					positions[zoneIndex],
					Entities.MonsterPosition(ChunkPosition),
				)
			}
		}
	}

	return positions
}
