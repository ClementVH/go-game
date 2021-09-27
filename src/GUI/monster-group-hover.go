package GUI

import (
	"fmt"
	"go-game/src/Entities"
	"go-game/src/MousePicker"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type MonsterGroupHover struct {
}

func newMonsterGroupHover() MonsterGroupHover {

	glfw.GetCurrentContext().SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

		if button == glfw.MouseButton1 && action == glfw.Release {
			var _group, _ = MousePicker.Picker.GetMonsterGroup()

			if _group != nil {
				group := _group.([]*Entities.Monster)
				fmt.Println(group[0].Position)
			}
		}
	})

	return MonsterGroupHover{}
}
