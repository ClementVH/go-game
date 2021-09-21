package Systems

import (
	"go-game/src/Entities"
	"go-game/src/Loaders"

	"github.com/go-gl/mathgl/mgl32"
)

type WildMonsterSystem struct {
	System
}

var wildMonsterGroups [][]*Entities.Monster

func NewWildMonsterSystem() *WildMonsterSystem {
	wildMonsterGroups = make([][]*Entities.Monster, 0, 256)

	wildMonster := Entities.NewMonster(
		Loaders.LoadGltf("../res/player", "player.gltf"),
		mgl32.Vec3{-8, 2, -8},
	)

	group := make([]*Entities.Monster, 0, 1)
	wildMonsterGroups = append(wildMonsterGroups, append(group, wildMonster))

	return &WildMonsterSystem{
		System: *NewSystem(),
	}
}

func (wildMonsterSystem *WildMonsterSystem) Tick() {

}

func (wildMonsterSystem *WildMonsterSystem) GetEntities() []Entities.IEntity {
	entities := make([]Entities.IEntity, 0, 256)

	for _, group := range wildMonsterGroups {
		for _, wildMonster := range group {
			entities = append(entities, wildMonster)
		}
	}

	return entities
}
