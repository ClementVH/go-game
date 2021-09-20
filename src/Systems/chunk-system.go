package Systems

import (
	"go-game/src/Entities"
	"go-game/src/Loaders"
	"math"
	_ "math/bits"
)

var DISPLAY_CHUNKS_SIZE = 7
var WORLD_CHUNKS_SIZE = 256

var currentChunks [][]Entities.ChunkPosition
var chunkEntities []*Entities.Chunk = make([]*Entities.Chunk, WORLD_CHUNKS_SIZE * WORLD_CHUNKS_SIZE)

type ChunkSystem struct {
	System
}

func NewChunkSystem() *ChunkSystem {
	chunks := make([][]Entities.ChunkPosition, DISPLAY_CHUNKS_SIZE)
	for i := range chunks {
		chunks[i] = make([]Entities.ChunkPosition, DISPLAY_CHUNKS_SIZE)
	}
	currentChunks = chunks

	model := Loaders.LoadGltf("../res/plane", "plane.gltf")

	for x := 0; x < WORLD_CHUNKS_SIZE; x++ {
		for z := 0; z < WORLD_CHUNKS_SIZE; z++ {
			chunkEntities[x * WORLD_CHUNKS_SIZE + z] = Entities.NewChunk(
				model,
				x - WORLD_CHUNKS_SIZE / 2,
				z - WORLD_CHUNKS_SIZE / 2,
			)
		}
	}

	return &ChunkSystem{
		System: *NewSystem(),
	}
}

func (chunkSystem *ChunkSystem) Tick() {
	posX := math.Floor(float64(player.Position[0] / 16))
	posZ := math.Floor(float64(player.Position[2] / 16)) + 1

	startX := int(posX) - (DISPLAY_CHUNKS_SIZE / 2)
	startZ := int(posZ) - (DISPLAY_CHUNKS_SIZE / 2)

	for x := 0; x < DISPLAY_CHUNKS_SIZE; x++ {
		for z := 0; z < DISPLAY_CHUNKS_SIZE; z++ {
			currentChunks[x][z] = Entities.ChunkPosition{
				X: startX + x,
				Z: startZ + z,
			}
		}
	}
}

func (chunkSystem *ChunkSystem) GetEntities() []Entities.IEntity {
	entities := make([]Entities.IEntity, 0, DISPLAY_CHUNKS_SIZE * DISPLAY_CHUNKS_SIZE)

	for _, chunks := range currentChunks {
		for _, chunk := range chunks {
			entity := chunkEntities[(chunk.X + WORLD_CHUNKS_SIZE / 2) * WORLD_CHUNKS_SIZE + chunk.Z + WORLD_CHUNKS_SIZE / 2]
			entities = append(entities, entity)
		}
	}

	return entities
}
