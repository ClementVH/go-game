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
	BBox BBox
}

func NewMonster(model []*Models.TexturedModel, position mgl32.Vec3) *Monster {
	entity := Entity{
		model,
		position,
		0, 0, 0, 1,
	}

	bbox := BBox{
		position,
		mgl32.Vec3{1, 1, 1},
	}

	monster := &Monster{
		entity,
		bbox,
	}

	return monster
}

func (monster *Monster) GetSignedDistance(position mgl32.Vec3) float32 {
	return monster.BBox.GetSignedDistance(position)
}
