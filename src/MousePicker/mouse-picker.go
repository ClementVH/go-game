package MousePicker

import (
	"go-game/src/ToolBox"
	"sync"

	"github.com/go-gl/mathgl/mgl32"
)

var once sync.Once

type ICamera interface {
	GetViewMatrix() mgl32.Mat4
	GetTargetPosition() mgl32.Vec3
	GetPosition() mgl32.Vec3
}

type MousePicker struct {
	ToolBox.Raycast
}

var Picker *MousePicker

func getPicker(camera ICamera, projectionMatrix mgl32.Mat4) *MousePicker {
	once.Do(func() {
		raycast := ToolBox.NewRaycast(camera, projectionMatrix)
		Picker = &MousePicker{
			raycast,
		}
	})

	return Picker
}
