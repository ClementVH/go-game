package GUI

import (
	"go-game/src/MousePicker"
	"go-game/src/State"
)

type GUI struct {
	MonsterGroupHover
}

func NewGUI() *GUI {
	State.GUI.SetMousePicker(
		MousePicker.NewMousePicker(State.Camera.Camera, State.Renderer.ProjectionMatrix),
	)

	return &GUI{
		newMonsterGroupHover(),
	}
}

func (gui *GUI) Update() {
	State.GUI.MousePicker.Update()
}
