package Entities

import "github.com/go-gl/mathgl/mgl32"

type Light struct {
	Position, Color mgl32.Vec3
}

func NewLight(position, color mgl32.Vec3) *Light {
	return &Light{
		position,
		color,
	}
}
