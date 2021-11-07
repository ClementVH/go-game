package Entities

import (
	"go-game/src/Models"
	"go-game/src/Window"

	"github.com/go-gl/mathgl/mgl32"
)

type Player struct {
	Entity
	Camera   *Camera
	Velocity mgl32.Vec2
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

	player.Camera = NewCamera(
		mgl32.Vec3{position.X() - 25, position.Y() + 25, position.Z() + 25},
		&player.Entity,
	)

	return player
}

func (player *Player) MoveTo(target mgl32.Vec3) {
	var diff = target.Sub(player.Position)
	diff = diff.Add(mgl32.Vec3{0.5, 0, 0.5})
	player.IncreasePostion(diff.Elem())
	player.Camera.IncreasePostion(diff.Elem())
}

func (player *Player) Move() {
	rotationMatrix := mgl32.Rotate2D(mgl32.DegToRad(player.Camera.Yaw))
	velocity := player.Velocity.Mul(Window.Delta)
	x, z := rotationMatrix.Mul2x1(velocity).Elem()
	player.IncreasePostion(x, 0, z)
	player.Camera.IncreasePostion(x, 0, z)
}
