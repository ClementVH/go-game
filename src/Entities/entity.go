package Entities

import (
	"go-game/src/Models"

	"github.com/go-gl/mathgl/mgl32"
)

type Entity struct {
	Model                   *Models.TexturedModel
	Position                mgl32.Vec3
	RotX, RotY, RotZ, Scale float32
}

func NewEntity(model *Models.TexturedModel, position mgl32.Vec3, RotX, RotY, RotZ, scale float32) *Entity {
	return &Entity{
		model,
		position,
		RotX, RotY, RotZ,
		scale,
	}
}

func (entity *Entity) IncreasePostion(dx, dy, dz float32) {
	entity.Position = entity.Position.Add(mgl32.Vec3{dx, dy, dz})
}

func (entity *Entity) IncreaseRotation(dx, dy, dz float32) {
	entity.RotX += dx
	entity.RotY += dy
	entity.RotZ += dz
}
