package Entities

import (
	"go-game/src/Models"
	"go-game/src/Window"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var RUN_SPEED float32 = 20

var Z_PRESSED = false
var Q_PRESSED = false
var S_PRESSED = false
var D_PRESSED = false

type Player struct {
	Entity
	Camera   *Camera
	velocity mgl32.Vec2
}

func NewPlayer(model []*Models.TexturedModel, position mgl32.Vec3, RotX, RotY, RotZ, scale float32) *Player {
	entity := Entity{
		model,
		position,
		RotX, RotY, RotZ,
		scale,
	}

	player := &Player{
		entity,
		nil,
		mgl32.Vec2{0, 0},
	}

	player.Camera = NewCamera(mgl32.Vec3{position.X() - 25, position.Y() + 25, position.Z() + 25}, player)

	player.initCameraMovements()

	return player
}

func (player *Player) Move() {
	rotationMatrix := mgl32.Rotate2D(mgl32.DegToRad(player.Camera.Yaw))
	velocity := player.velocity.Mul(Window.Delta)
	x, z := rotationMatrix.Mul2x1(velocity).Elem()
	player.IncreasePostion(x, 0, z)
	player.Camera.IncreasePostion(x, 0, z)
}

func (player *Player) initCameraMovements() {
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
			player.velocity[0] = 1
		} else if Q_PRESSED {
			player.velocity[0] = -1
		} else {
			player.velocity[0] = 0
		}

		if Z_PRESSED {
			player.velocity[1] = -1
		} else if S_PRESSED {
			player.velocity[1] = 1
		} else {
			player.velocity[1] = 0
		}

		if player.velocity.Len() > 0 {
			player.velocity = player.velocity.Normalize()
		}

		player.velocity = player.velocity.Mul(RUN_SPEED)
	})
}
