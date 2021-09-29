package Combat

import "go-game/src/Entities"

const (
	PRE_COMBAT string = "PRE_COMBAT"
)

type Combat struct {
	Monsters []*Entities.Monster
	Chunk    *Entities.Chunk
	Status   string
}

func NewCombat(group []*Entities.Monster, chunk *Entities.Chunk) *Combat {
	return &Combat{
		group,
		chunk,
		PRE_COMBAT,
	}
}

func (combat *Combat) GetMonsters() []*Entities.Monster {
	return combat.Monsters
}
