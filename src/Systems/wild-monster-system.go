package Systems

import (
	"go-game/src/Entities"
	"go-game/src/Loaders"
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

type WildMonsterSystem struct {
	System
}

var wildMonsterGroups [][]*Entities.Monster
var spawnZones [][]Entities.MonsterPosition
var currentSpawnZone int

func NewWildMonsterSystem() *WildMonsterSystem {
	spawnZones = Loaders.GetSpawnZones(zones)

	loadSpawnZone(getZoneIndex())

	system := &WildMonsterSystem{
		System: *NewSystem(),
	}

	return system
}

func (wildMonsterSystem *WildMonsterSystem) Tick() {
	zoneIndex := getZoneIndex()
	if zoneIndex != currentSpawnZone {
		loadSpawnZone(zoneIndex)
	}
}

func (wildMonsterSystem *WildMonsterSystem) GetEntities() []*Entities.Entity {
	entities := make([]*Entities.Entity, 0, 256)
	playerChunkX := math.Floor(float64(Player.Position.X()) / 16)
	playerChunkZ := math.Floor(float64(Player.Position.Z()) / 16)

	for _, group := range wildMonsterGroups {
		for _, wildMonster := range group {
			monsterChunkX := math.Floor(float64(wildMonster.Position.X()) / 16)
			monsterChunkZ := math.Floor(float64(wildMonster.Position.Z()) / 16)

			if monsterChunkX > playerChunkX-float64(DISPLAY_CHUNKS_SIZE/2) &&
				monsterChunkX < playerChunkX+float64(DISPLAY_CHUNKS_SIZE/2) &&

				monsterChunkZ > playerChunkZ-float64(DISPLAY_CHUNKS_SIZE/2) &&
				monsterChunkZ < playerChunkZ+float64(DISPLAY_CHUNKS_SIZE/2) {

				entities = append(entities, &wildMonster.Entity)
			}
		}
	}

	return entities
}

func loadSpawnZone(zoneIndex int) {
	currentSpawnZone = zoneIndex

	wildMonsterGroups = make([][]*Entities.Monster, 0, 256)

	monsterModel := Loaders.LoadGltf("../res/player", "player.gltf")

	for _, position := range spawnZones[zoneIndex] {
		random := rand.Intn(100)

		group := make([]*Entities.Monster, 0, 1)

		if random < 20 {

			wildMonster := Entities.NewMonster(
				monsterModel,
				mgl32.Vec3{
					float32(position.X*16 + 8),
					2,
					float32(position.Z*16 + 8),
				},
			)
			group = append(group, wildMonster)

		}
		wildMonsterGroups = append(wildMonsterGroups, group)
	}

}

func (wildMonsterSystem *WildMonsterSystem) GetGroups() [][]*Entities.Monster {
	groups := make([][]*Entities.Monster, 0, 256)
	playerChunkX := math.Floor(float64(Player.Position.X()) / 16)
	playerChunkZ := math.Floor(float64(Player.Position.Z()) / 16)

	for _, group := range wildMonsterGroups {
		for _, wildMonster := range group {
			monsterChunkX := math.Floor(float64(wildMonster.Position.X()) / 16)
			monsterChunkZ := math.Floor(float64(wildMonster.Position.Z()) / 16)

			if monsterChunkX > playerChunkX-float64(DISPLAY_CHUNKS_SIZE/2) &&
				monsterChunkX < playerChunkX+float64(DISPLAY_CHUNKS_SIZE/2) &&

				monsterChunkZ > playerChunkZ-float64(DISPLAY_CHUNKS_SIZE/2) &&
				monsterChunkZ < playerChunkZ+float64(DISPLAY_CHUNKS_SIZE/2) {

				groups = append(groups, group)
				break
			}
		}
	}

	return groups
}
