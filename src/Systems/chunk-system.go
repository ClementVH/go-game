package Systems

import (
	"go-game/src/Entities"
	"go-game/src/Loaders"
	"math"
	_ "math/bits"
)

var DISPLAY_CHUNKS_SIZE = 8
var WORLD_CHUNKS_SIZE = 256

var currentChunks [][]Entities.ChunkPosition
var chunkEntities [][]*Entities.Chunk
var zones [][]Entities.ChunkPosition
var currentZone int

type ChunkSystem struct {
	System
}

func NewChunkSystem() *ChunkSystem {
	chunksPositions := make([][]Entities.ChunkPosition, DISPLAY_CHUNKS_SIZE)
	for i := range chunksPositions {
		chunksPositions[i] = make([]Entities.ChunkPosition, DISPLAY_CHUNKS_SIZE)
	}
	currentChunks = chunksPositions

	zones = Loaders.GetChunkZones()

	loadZone(getZoneIndex())

	return &ChunkSystem{
		System: *NewSystem(),
	}
}

func (chunkSystem *ChunkSystem) Tick() {
	zoneIndex := getZoneIndex()
	if zoneIndex != currentZone {
		loadZone(zoneIndex)
	}

	posX := math.Floor(float64(Player.Position[0] / 16))
	posZ := math.Floor(float64(Player.Position[2] / 16))

	startX := int(posX) - (DISPLAY_CHUNKS_SIZE / 2)
	startZ := int(posZ) - (DISPLAY_CHUNKS_SIZE / 2)

	for x := 0; x < DISPLAY_CHUNKS_SIZE; x++ {
		for z := 0; z < DISPLAY_CHUNKS_SIZE; z++ {
			currentChunks[x][z] = Entities.ChunkPosition{
				X: startX + x,
				Z: startZ + z,
			}
		}
	}
}

func (chunkSystem *ChunkSystem) GetEntities() []Entities.IEntity {
	entities := make([]Entities.IEntity, 0, DISPLAY_CHUNKS_SIZE*DISPLAY_CHUNKS_SIZE)

	for _, chunks := range currentChunks {
		for _, chunk := range chunks {
			if chunk.X >= 0 && chunk.Z >= 0 {
				entity := chunkEntities[chunk.X][chunk.Z]
				if entity != nil {
					entities = append(entities, entity)
				}
			}
		}
	}

	return entities
}

func loadZone(zoneIndex int) {
	currentZone = zoneIndex

	chunks := make([][]*Entities.Chunk, WORLD_CHUNKS_SIZE)
	for i := range chunks {
		chunks[i] = make([]*Entities.Chunk, WORLD_CHUNKS_SIZE)
	}
	chunkEntities = chunks

	model := Loaders.LoadGltf("../res/plane", "plane.gltf")
	startPositionTexture := Loaders.LoadTexture("../res/textures", "start-positions.png")

	for _, position := range zones[zoneIndex] {
		chunkEntities[position.X][position.Z] = Entities.NewChunk(
			model,
			position.X,
			position.Z,
			startPositionTexture,
		)
	}
}

func getZoneIndex() int {
	var zoneIndex = 0
	for i, zone := range zones {
		for _, position := range zone {
			diffX := Player.Position.X() - float32(position.X)*16
			diffZ := Player.Position.Z() - float32(position.Z)*16
			if diffX >= 0 && diffX < 16 && diffZ >= 0 && diffZ < 16 {
				zoneIndex = i
			}
		}
	}

	return zoneIndex
}

func (chunkSystem *ChunkSystem) GetChunk(x, y int) *Entities.Chunk {
	return chunkEntities[x][y]
}
