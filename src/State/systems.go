package State

import "go-game/src/Entities"

var Systems SystemsState = SystemsState{}

type ISystem interface {
	GetEntities() []Entities.IEntity
	Tick()
}

type IPlayerSystem interface {
	ISystem
	GetPlayer() *Entities.Player
}

type IWildMonsterSystem interface {
	ISystem
	GetGroups() [][]*Entities.Monster
}

type IChunkSystem interface {
	ISystem
	GetChunk(x, y int) *Entities.Chunk
}

type SystemsState struct {
	PlayerSystem      IPlayerSystem
	ChunkSystem       IChunkSystem
	WildMonsterSystem IWildMonsterSystem
}

func (state *SystemsState) SetPlayerSystem(playerSystem IPlayerSystem) {
	state.PlayerSystem = playerSystem
}

func (state *SystemsState) SetChunkSystem(chunksSystem IChunkSystem) {
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
