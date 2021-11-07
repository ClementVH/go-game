package Controllers

import (
	"go-game/src/State"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var RUN_SPEED float32 = 20

var Z_PRESSED = false
var Q_PRESSED = false
var S_PRESSED = false
var D_PRESSED = false

func InitCameraMovements() {
	player := State.GetPlayer()
	glfw.GetCurrentContext().SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		keyName := glfw.GetKeyName(key, scancode)

		press := action == glfw.Press
		release := action == glfw.Release

		if press {
			switch keyName {
			case "z":
				Z_PRESSED = true
			case "q":
				Q_PRESSED = true
			case "s":
				S_PRESSED = true
			case "d":
				D_PRESSED = true
			}
		} else if release {
			switch keyName {
			case "z":
				Z_PRESSED = false
			case "q":
				Q_PRESSED = false
			case "s":
				S_PRESSED = false
			case "d":
				D_PRESSED = false
			}
		}

		if D_PRESSED {
			player.Velocity[0] = 1
		} else if Q_PRESSED {
			player.Velocity[0] = -1
		} else {
			player.Velocity[0] = 0
		}

		if Z_PRESSED {
			player.Velocity[1] = -1
		} else if S_PRESSED {
			player.Velocity[1] = 1
		} else {
			player.Velocity[1] = 0
		}

		if player.Velocity.Len() > 0 {
			player.Velocity = player.Velocity.Normalize()
		}

		if State.Combat != nil && State.Combat.GetStatus() == "PRE_COMBAT" {
			player.Velocity = player.Velocity.Mul(0)
		} else {
			player.Velocity = player.Velocity.Mul(RUN_SPEED)
		}
	})
}
