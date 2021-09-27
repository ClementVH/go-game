package MousePicker

import (
	"go-game/src/Entities"
	"go-game/src/Physics"

	"github.com/go-gl/mathgl/mgl32"
)

type MousePicker struct {
	Physics.Raycast
}

func NewMousePicker(camera *Entities.Camera, projectionMatrix mgl32.Mat4) *MousePicker {
	raycast := Physics.NewRaycast(camera, projectionMatrix)
	return &MousePicker{
		raycast,
	}
}
