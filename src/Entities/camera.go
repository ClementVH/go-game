package Entities

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position         mgl32.Vec3
	Pitch, Yaw, Roll float32
}

func NewCamera() *Camera {
	camera := Camera{}
	camera.initCameraMovements()
	return &camera
}

func (camera *Camera) Move(keyName string) {
	if keyName == "z" {
		camera.Position[1] += 0.04
	}
	if keyName == "d" {
		camera.Position[0] += 0.04
	}
	if keyName == "q" {
		camera.Position[0] -= 0.04
	}
	if keyName == "s" {
		camera.Position[1] -= 0.04
	}
}

func (camera *Camera) initCameraMovements() {
	glfw.GetCurrentContext().SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		keyName := glfw.GetKeyName(key, scancode)
		switch keyName {
		case "z":
			camera.Move(keyName)
		case "q":
			camera.Move(keyName)
		case "s":
			camera.Move(keyName)
		case "d":
			camera.Move(keyName)
		}
	})
}
