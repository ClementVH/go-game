package GUI

import (
	"go-game/src/MousePicker"
	"go-game/src/State"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type GUI struct {
}

func NewGUI() *GUI {
	State.GUI.SetMousePicker(
		MousePicker.NewMousePicker(State.Camera.Camera, State.Renderer.ProjectionMatrix),
	)

	gui := &GUI{}

	glfw.GetCurrentContext().SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		gui.runMonsterGroupHover(button, action)
		gui.runStartPositionHover(button, action)
	})

	return gui
}

func (gui *GUI) Update() {
	State.GUI.MousePicker.Update()
}
