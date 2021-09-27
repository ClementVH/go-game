package State

import "go-game/src/Entities"

var GUI GUIState = GUIState{}

type IMousePicker interface {
	Update()
	GetMonsterGroup() ([]*Entities.Monster, error)
}

type GUIState struct {
	MousePicker IMousePicker
}

func (gui *GUIState) SetMousePicker(mousePicker IMousePicker) {
	gui.MousePicker = mousePicker
}
