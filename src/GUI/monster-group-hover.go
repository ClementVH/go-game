package GUI

import (
	"go-game/src/Combat"
	"go-game/src/State"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type MonsterGroupHover struct {
}

func (gui *GUI) runMonsterGroupHover(button glfw.MouseButton, action glfw.Action) {

	if button == glfw.MouseButton1 && action == glfw.Release {
		var group, _ = State.GUI.MousePicker.GetMonsterGroup()

		if group != nil {
			x, y := group[0].GetChunkCoords()
			chunk := State.Systems.ChunkSystem.GetChunk(x, y)
			State.Combat.SetCombat(Combat.NewCombat(group, chunk))
		}
	}
}
