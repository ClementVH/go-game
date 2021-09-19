package Systems

import (
	"go-game/src/Entities"
	"go-game/src/State"
	"math"
	_ "math/bits"
)

type ChunkSystem struct {
	System
}

func NewChunkSystem() *ChunkSystem {
	chunks := make([][]Entities.ChunkPosition, State.DISPLAY_CHUNKS_SIZE)
	for i := range chunks {
		chunks[i] = make([]Entities.ChunkPosition, State.DISPLAY_CHUNKS_SIZE)
	}
	State.CurrentChunks = chunks

	return &ChunkSystem{
		System: *NewSystem(),
	}
}

func (chunkSystem *ChunkSystem) Tick() {
	character := State.Character
	posX := math.Floor(float64(character.Position[0] / 16))
	posZ := math.Floor(float64(character.Position[2] / 16)) + 1

	startX := int(posX) - (State.DISPLAY_CHUNKS_SIZE / 2)
	startZ := int(posZ) - (State.DISPLAY_CHUNKS_SIZE / 2)

	for x := 0; x < State.DISPLAY_CHUNKS_SIZE; x++ {
		for z := 0; z < State.DISPLAY_CHUNKS_SIZE; z++ {
			State.CurrentChunks[x][z] = Entities.ChunkPosition{
				X: startX + x,
				Z: startZ + z,
			}
		}
	}
}