package State

import (
	"go-game/src/Entities"
)

var DISPLAY_CHUNKS_SIZE = 7

var WORLD_CHUNKS_SIZE = 256

var Character *Entities.Character
var CurrentChunks [][]Entities.ChunkPosition
var ChunkEntities []*Entities.Chunk = make([]*Entities.Chunk, WORLD_CHUNKS_SIZE * WORLD_CHUNKS_SIZE)

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
