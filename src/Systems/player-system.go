package Systems

import (
	"go-game/src/Entities"
	"go-game/src/Loaders"
	"go-game/src/State"

	"github.com/go-gl/mathgl/mgl32"
)

type PlayerSystem struct {
	System
	Player *Entities.Player
}

func NewPlayerSystem() *PlayerSystem {

	player := Entities.NewPlayer(
		Loaders.LoadGltf("../res/player", "player.gltf"),
		mgl32.Vec3{43*16 + 8, 0, 46*16 + 8},
		0, 0, 0, 1,
	)

	State.Camera.SetCamera(player.Camera)

	return &PlayerSystem{
		System: *NewSystem(),
		Player: player,
	}
}

func (playerSystem *PlayerSystem) Tick() {
	playerSystem.Player.Move()
}

func (playerSystem *PlayerSystem) GetEntities() []Entities.IEntity {
	return []Entities.IEntity{playerSystem.Player}
}

func (playerSystem *PlayerSystem) GetPlayer() *Entities.Player {
	return playerSystem.Player
}
