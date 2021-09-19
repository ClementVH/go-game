package Systems

import "go-game/src/Entities"

var Character *Entities.Character

type CharacterSystem struct {
	System
}

func NewCharacterSystem() *ChunkSystem {

	return &ChunkSystem{
		System: *NewSystem(),
	}
}