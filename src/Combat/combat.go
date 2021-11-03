package Combat

import (
	"go-game/src/Entities"
	"go-game/src/Systems"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	PRE_COMBAT string = "PRE_COMBAT"
)

type Combat struct {
	Monsters []*Entities.Monster
	Chunk    *Entities.Chunk
	Status   string
}

func NewCombat(group []*Entities.Monster, chunk *Entities.Chunk) *Combat {
	player := Systems.Player
	startPosition := chunk.StartPositions.Teams[0][0]
	startPositionWorldCoordinates := mgl32.Vec3{
		chunk.Position.X() + startPosition.X(),
		0,
		chunk.Position.Z() + startPosition.Y(),
	}

	player.MoveTo(startPositionWorldCoordinates)

	monstersStartPositions := chunk.StartPositions.Teams[1]
	for index, monster := range group {
		monstersStartPosition := monstersStartPositions[index]
		monsterStartPositionWorldCoordinates := mgl32.Vec3{
			chunk.Position.X() + monstersStartPosition.X(),
			0,
			chunk.Position.Z() + monstersStartPosition.Y(),
		}

		monster.MoveTo(monsterStartPositionWorldCoordinates)
	}

	return &Combat{
		group,
		chunk,
		PRE_COMBAT,
	}
}

func (combat *Combat) GetMonsters() []*Entities.Monster {
	return combat.Monsters
}

func (combat *Combat) GetChunk() *Entities.Chunk {
	return combat.Chunk
}
