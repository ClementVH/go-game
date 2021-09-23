package Entities

import (
	"go-game/src/Models"

	"github.com/go-gl/mathgl/mgl32"
)

type MonsterPosition struct {
	X int
	Z int
}

type Monster struct {
	Entity
}

func NewMonster(model []*Models.TexturedModel, position mgl32.Vec3) *Monster {
	entity := Entity{
		model,
		position,
		0, 0, 0, 1,
	}

	monster := &Monster{
		entity,
	}

	return monster
}