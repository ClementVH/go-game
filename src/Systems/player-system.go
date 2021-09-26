package Systems

import (
	"go-game/src/Entities"
	"go-game/src/Loaders"

	"github.com/go-gl/mathgl/mgl32"
)

var Player *Entities.Player

type IPlayerSystem interface {
	GetPlayer() *Entities.Player
}

type PlayerSystem struct {
	System
}

func NewPlayerSystem() *PlayerSystem {

	Player = Entities.NewPlayer(
		Loaders.LoadGltf("../res/player", "player.gltf"),
		mgl32.Vec3{43 * 16 + 8, 2, 46 * 16 + 8},
		0, 0, 0, 1,
	)

	return &PlayerSystem{
		System: *NewSystem(),
	}
}

func (playerSystem *PlayerSystem) Tick() {
	Player.Move()
}

func (playerSystem *PlayerSystem) GetEntities() []Entities.IEntity {
	return []Entities.IEntity{Player}
}

func (playerSystem *PlayerSystem) GetPlayer() *Entities.Player {
	return Player
}