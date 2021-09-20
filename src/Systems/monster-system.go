package Systems

import (
	"go-game/src/Entities"
	"go-game/src/Loaders"

	"github.com/go-gl/mathgl/mgl32"
)

type MonsterSystem struct {
	System
}

var monsterGroups [][]*Entities.Monster

func NewMonsterSystem() *MonsterSystem {
	monsterGroups = make([][]*Entities.Monster, 0, 256)

	monster := Entities.NewMonster(
		Loaders.LoadGltf("../res/player", "player.gltf"),
		mgl32.Vec3{-8, 2, -8},
	)

	group := make([]*Entities.Monster, 0, 1)
	monsterGroups = append(monsterGroups, append(group, monster))

	return &MonsterSystem{
		System: *NewSystem(),
	}
}

func (monsterSystem *MonsterSystem) Tick() {

}

func (monsterSystem *MonsterSystem) GetEntities() []Entities.IEntity {
	entities := make([]Entities.IEntity, 0, 256)

	for _, group := range monsterGroups {
		for _, monster := range group {
			entities = append(entities, monster)
		}
	}

	return entities
}
