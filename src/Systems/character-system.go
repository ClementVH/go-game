package Systems

import "go-game/src/Entities"

var Player *Entities.Player

type PlayerSystem struct {
	System
}

func NewPlayerSystem() *ChunkSystem {

	return &ChunkSystem{
		System: *NewSystem(),
	}
}