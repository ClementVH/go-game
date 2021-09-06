package Entities

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position         mgl32.Vec3
	Pitch, Yaw, Roll float32
}

func NewCamera() Camera {
	return Camera{}
}

func (camera *Camera) Move(keyName string) {
	if keyName == "z" {
		camera.Position[1] -= 0.02
	}
	if keyName == "d" {
		camera.Position[0] -= 0.02
	}
	if keyName == "q" {
		camera.Position[0] += 0.02
	}
	if keyName == "s" {
		camera.Position[1] += 0.02
	}

}
