package State

import "go-game/src/Entities"

var Systems SystemsState = SystemsState{}

type ISystem interface {
	GetEntities() []Entities.IEntity
	Tick()
}

type IWildMonsterSystem interface {
	ISystem
	GetGroups() [][]*Entities.Monster
}

type SystemsState struct {
	PlayerSystem      ISystem
	ChunkSystem       ISystem
	WildMonsterSystem IWildMonsterSystem
}

func (state *SystemsState) SetPlayerSystem(playerSystem ISystem) {
	state.PlayerSystem = playerSystem
}

func (state *SystemsState) SetChunkSystem(chunksSystem ISystem) {
	state.ChunkSystem = chunksSystem
}

func (state *SystemsState) SetWildMonsterSystem(wildMonstersSystem IWildMonsterSystem) {
	state.WildMonsterSystem = wildMonstersSystem
}

func (state *SystemsState) GetAll() []ISystem {
	return []ISystem{
		state.PlayerSystem,
		state.ChunkSystem,
		state.WildMonsterSystem,
	}
}
