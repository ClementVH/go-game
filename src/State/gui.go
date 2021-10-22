package State

import (
	"go-game/src/Entities"

	"github.com/go-gl/mathgl/mgl32"
)

var GUI GUIState = GUIState{}

type IMousePicker interface {
	Update()
	GetMonsterGroup() ([]*Entities.Monster, error)
	GetStartPosition() (mgl32.Vec2, error)
}

type GUIState struct {
	MousePicker IMousePicker
}

func (gui *GUIState) SetMousePicker(mousePicker IMousePicker) {
	gui.MousePicker = mousePicker
}
