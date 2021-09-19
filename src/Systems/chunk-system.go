package Systems

import (
	"go-game/src/Entities"
	"math"
	_ "math/bits"
)

var DISPLAY_CHUNKS_SIZE = 7
var WORLD_CHUNKS_SIZE = 256

var CurrentChunks [][]Entities.ChunkPosition
var ChunkEntities []*Entities.Chunk = make([]*Entities.Chunk, WORLD_CHUNKS_SIZE * WORLD_CHUNKS_SIZE)

type ChunkSystem struct {
	System
}

func NewChunkSystem() *ChunkSystem {
	chunks := make([][]Entities.ChunkPosition, DISPLAY_CHUNKS_SIZE)
	for i := range chunks {
		chunks[i] = make([]Entities.ChunkPosition, DISPLAY_CHUNKS_SIZE)
	}
	CurrentChunks = chunks

	return &ChunkSystem{
		System: *NewSystem(),
	}
}

func (chunkSystem *ChunkSystem) Tick() {
	posX := math.Floor(float64(Player.Position[0] / 16))
	posZ := math.Floor(float64(Player.Position[2] / 16)) + 1

	startX := int(posX) - (DISPLAY_CHUNKS_SIZE / 2)
	startZ := int(posZ) - (DISPLAY_CHUNKS_SIZE / 2)

	for x := 0; x < DISPLAY_CHUNKS_SIZE; x++ {
		for z := 0; z < DISPLAY_CHUNKS_SIZE; z++ {
			CurrentChunks[x][z] = Entities.ChunkPosition{
				X: startX + x,
				Z: startZ + z,
			}
		}
	}
}

func GetChunksToRender() []Entities.IEntity {
	entities := make([]Entities.IEntity, 0, DISPLAY_CHUNKS_SIZE * DISPLAY_CHUNKS_SIZE)

	for _, chunks := range CurrentChunks {
		for _, chunk := range chunks {
			entity := ChunkEntities[(chunk.X + WORLD_CHUNKS_SIZE / 2) * WORLD_CHUNKS_SIZE + chunk.Z + WORLD_CHUNKS_SIZE / 2]
			entities = append(entities, entity)
		}
	}

	return entities
}
