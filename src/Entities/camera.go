package Entities

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Entity
	target           *Entity
	Pitch, Yaw, Roll float32
}

func NewCamera(position mgl32.Vec3, target *Entity) *Camera {
	camera := Camera{
		*NewEntity(nil, position, 0, 0, 0, 0),
		target,
		35.264, 45, 0,
	}
	return &camera
}

func (camera *Camera) GetInfo() struct {
	Pitch, Yaw float32
	Position   mgl32.Vec3
} {
	return struct {
		Pitch    float32
		Yaw      float32
		Position mgl32.Vec3
	}{
		camera.Pitch,
		camera.Yaw,
		camera.Position,
	}
}

func (camera *Camera) GetViewMatrix() mgl32.Mat4 {
	info := camera.GetInfo()
	matrix := mgl32.Ident4()
	matrix = matrix.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(info.Pitch), mgl32.Vec3{1, 0, 0}))
	matrix = matrix.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(info.Yaw), mgl32.Vec3{0, 1, 0}))

	cameraPos := info.Position
	matrix = matrix.Mul4(mgl32.Translate3D(-cameraPos.X(), -cameraPos.Y(), -cameraPos.Z()))
	return matrix
}

func (camera *Camera) GetTargetPosition() mgl32.Vec3 {
	return camera.target.GetPosition()
}
