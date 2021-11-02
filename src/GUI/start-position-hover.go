package GUI

import (
	"fmt"
	"go-game/src/State"
	"go-game/src/Systems"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func (gui *GUI) runStartPositionHover(button glfw.MouseButton, action glfw.Action) {

	if State.Combat.Combat != nil && button == glfw.MouseButton1 && action == glfw.Release {
		var position, err = State.GUI.MousePicker.GetStartPosition()
		chunk := State.Combat.Combat.GetChunk()
		player := Systems.Player
		if err == nil {
			fmt.Println(position)
			positionWorldCoordinates := mgl32.Vec3{
				chunk.Position.X() + position.X(),
				0,
				chunk.Position.Z() + position.Y(),
			}

			player.MoveTo(positionWorldCoordinates)
		}
	}
}
