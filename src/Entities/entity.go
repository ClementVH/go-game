package Entities

import (
	"go-game/src/Models"
	"go-game/src/ToolBox"

	"github.com/go-gl/mathgl/mgl32"
)

type Entity struct {
	Meshes                  []*Models.TexturedModel
	Position                mgl32.Vec3
	RotX, RotY, RotZ, Scale float32
}

func NewEntity(model []*Models.TexturedModel, position mgl32.Vec3, RotX, RotY, RotZ, scale float32) *Entity {
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

func (entity *Entity) GetMeshes() []*Models.TexturedModel {
	return entity.Meshes
}

func (entity *Entity) GetTransformationMatrix() mgl32.Mat4 {
	return ToolBox.CreateTransformationMatrix(
		entity.Position,
		entity.RotX, entity.RotY, entity.RotZ,
		entity.Scale,
	)
}

func (entity *Entity) GetPosition() mgl32.Vec3 {
	return entity.Position
}

func (entity *Entity) GetSignedDistance(position mgl32.Vec3) float32 {
	return 0
}
