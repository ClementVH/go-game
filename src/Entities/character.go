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

type ICharacter interface {
	GetCamera()
}

type Character struct {
	Entity
	Camera *Camera
	velocity mgl32.Vec2
}

func NewCharacter(model []*Models.TexturedModel, position mgl32.Vec3, RotX, RotY, RotZ, scale float32) *Character {
	entity := Entity{
		model,
		position,
		RotX, RotY, RotZ,
		scale,
	}

	character := &Character{
		entity,
		NewCamera(position),
		mgl32.Vec2{0, 0},
	}

	character.initCameraMovements();

	return character;
}

func (character *Character) Move() {
	rotationMatrix := mgl32.Rotate2D(mgl32.DegToRad(character.Camera.Yaw))
	velocity := character.velocity.Mul(Window.Delta)
	x, z := rotationMatrix.Mul2x1(velocity).Elem()
	character.IncreasePostion(x, 0, z)
	character.Camera.IncreasePostion(x, 0, z)
}

func (character *Character) initCameraMovements() {
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
			character.velocity[0] = 1
		} else if Q_PRESSED {
			character.velocity[0] = -1
		} else {
			character.velocity[0] = 0
		}

		if Z_PRESSED {
			character.velocity[1] = -1
		} else if S_PRESSED {
			character.velocity[1] = 1
		} else {
			character.velocity[1] = 0
		}

		if (character.velocity.Len() > 0) {
			character.velocity = character.velocity.Normalize()
		}

		character.velocity = character.velocity.Mul(RUN_SPEED)
	})
}
