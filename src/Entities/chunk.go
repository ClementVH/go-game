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
	StartPositions
	StartPositionsId uint32
}

func NewChunk(model []*Models.TexturedModel, X, Z int, startPositionId uint32) *Chunk {
	entity := Entity{
		model,
		mgl32.Vec3{float32(X) * 16, 0, float32(Z) * 16},
		0, 0, 0, 1,
	}

	startPositions := StartPositions{
		[][]mgl32.Vec2{
			{
				mgl32.Vec2{2, 4},
				mgl32.Vec2{2, 5},
				mgl32.Vec2{2, 6},

				mgl32.Vec2{3, 3},
				mgl32.Vec2{4, 3},
				mgl32.Vec2{5, 3},
			},
			{
				mgl32.Vec2{13, 9},
				mgl32.Vec2{14, 9},
				mgl32.Vec2{15, 9},

				mgl32.Vec2{16, 8},
				mgl32.Vec2{16, 7},
				mgl32.Vec2{16, 6},
			},
		},
	}

	chunk := &Chunk{
		entity,
		startPositions,
		startPositionId,
	}

	return chunk
}

type StartPositions struct {
	Teams [][]mgl32.Vec2
}
