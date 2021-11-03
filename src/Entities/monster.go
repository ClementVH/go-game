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
		position.Add(mgl32.Vec3{0, 1, 0}),
		mgl32.Vec3{0.5, 1, 0.5},
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

func (monster *Monster) MoveTo(target mgl32.Vec3) {
	var diff = target.Sub(monster.Position)
	diff = diff.Add(mgl32.Vec3{0.5, 0, 0.5})
	monster.IncreasePostion(diff.Elem())
}
