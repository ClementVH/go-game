package MousePicker

import (
	"go-game/src/Entities"
	"go-game/src/ToolBox"

	"github.com/go-gl/mathgl/mgl32"
)

type MousePicker struct {
	ToolBox.Raycast
}

func NewMousePicker(camera *Entities.Camera, projectionMatrix mgl32.Mat4) *MousePicker {
	raycast := ToolBox.NewRaycast(camera, projectionMatrix)
	return &MousePicker{
		raycast,
	}
}
