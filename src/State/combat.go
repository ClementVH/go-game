package State

import "go-game/src/Entities"

var Combat *CombatState

type ICombat interface {
	GetMonsters() []*Entities.Monster
	GetChunk() *Entities.Chunk
	GetStatus() string
}

type CombatState struct {
	ICombat
}

func (state *CombatState) SetCombat(combat ICombat) {
	Combat = &CombatState{
		combat,
	}
}
