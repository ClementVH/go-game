package GUI

import (
	"fmt"
	"go-game/src/State"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type MonsterGroupHover struct {
}

func newMonsterGroupHover() MonsterGroupHover {

	glfw.GetCurrentContext().SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

		if button == glfw.MouseButton1 && action == glfw.Release {
			var group, _ = State.GUI.MousePicker.GetMonsterGroup()

			if group != nil {
				fmt.Println(group[0].Position)
			}
		}
	})

	return MonsterGroupHover{}
}
