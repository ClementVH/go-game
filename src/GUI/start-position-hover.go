package GUI

import (
	"fmt"
	"go-game/src/State"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func (gui *GUI) runStartPositionHover(button glfw.MouseButton, action glfw.Action) {

	if State.Combat.Combat != nil && button == glfw.MouseButton1 && action == glfw.Release {
		var position, _ = State.GUI.MousePicker.GetStartPosition()

		fmt.Println(position)
	}
}
