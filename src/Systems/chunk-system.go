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
var chunkEntities [][]*Entities.Chunk

type ChunkSystem struct {
	System
}

func NewChunkSystem() *ChunkSystem {
	chunksPositions := make([][]Entities.ChunkPosition, DISPLAY_CHUNKS_SIZE)
	for i := range chunksPositions {
		chunksPositions[i] = make([]Entities.ChunkPosition, DISPLAY_CHUNKS_SIZE)
	}
	currentChunks = chunksPositions

	chunks := make([][]*Entities.Chunk, WORLD_CHUNKS_SIZE)
	for i := range chunks {
		chunks[i] = make([]*Entities.Chunk, WORLD_CHUNKS_SIZE)
	}
	chunkEntities = chunks

	model := Loaders.LoadGltf("../res/plane", "plane.gltf")

	positions := Loaders.GetChunkPositions()

	for _, position := range positions {
		chunkEntities[position.X][position.Z] = Entities.NewChunk(
			model,
			position.X,
			position.Z,
		)
	}

	return &ChunkSystem{
		System: *NewSystem(),
	}
}

func (chunkSystem *ChunkSystem) Tick() {
	posX := math.Floor(float64(player.Position[0] / 16))
	posZ := math.Floor(float64(player.Position[2] / 16))

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
			if (chunk.X > 0 && chunk.Z > 0) {
				entity := chunkEntities[chunk.X][chunk.Z]
				if (entity != nil) {
					entities = append(entities, entity)
				}
			}
		}
	}

	return entities
}
