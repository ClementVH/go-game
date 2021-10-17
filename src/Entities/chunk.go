package Entities

import (
	"go-game/src/Models"

	"github.com/go-gl/mathgl/mgl32"
)

type ChunkPosition struct {
	X int
	Z int
}

type Chunk struct {
	Entity
	StartPositionsId uint32
}

func NewChunk(model []*Models.TexturedModel, X, Z int, startPositionId uint32) *Chunk {
	entity := Entity{
		model,
		mgl32.Vec3{float32(X) * 16, 0, float32(Z) * 16},
		0, 0, 0, 1,
	}

	chunk := &Chunk{
		entity,
		startPositionId,
	}

	return chunk
}
