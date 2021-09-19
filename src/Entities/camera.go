package Entities

import (
	"github.com/go-gl/mathgl/mgl32"
)

type ICamera interface {
	GetInfo() struct{
		Pitch, Yaw float32
		Position mgl32.Vec3
	}
}

type Camera struct {
	Entity
	Pitch, Yaw, Roll float32
}

func NewCamera(position mgl32.Vec3) *Camera {
	camera := Camera{
		*NewEntity(nil, position, 0, 0, 0, 0),
		35.264, 45, 0,
	}
	return &camera
}

func (camera *Camera) GetInfo() struct{
	Pitch, Yaw float32
	Position mgl32.Vec3
}{
	return struct{Pitch float32; Yaw float32; Position mgl32.Vec3}{
		camera.Pitch,
		camera.Yaw,
		camera.Position,
	}
}