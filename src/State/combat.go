package State

import "go-game/src/Entities"

var Combat CombatState = CombatState{}

type ICombat interface {
	GetMonsters() []*Entities.Monster
	GetChunk() *Entities.Chunk
}

type CombatState struct {
	Combat ICombat
}

func (state *CombatState) SetCombat(combat ICombat) {
	state.Combat = combat
}
