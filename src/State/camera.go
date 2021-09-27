package State

import "go-game/src/Entities"

var Camera CameraState = CameraState{}

type CameraState struct {
	Camera *Entities.Camera
}

func (state *CameraState) SetCamera(camera *Entities.Camera) {
	state.Camera = camera
}
