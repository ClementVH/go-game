package Terrains

import (
	"go-game/src/Loaders"
	"go-game/src/Models"
	"go-game/src/Textures"
)

var SIZE = float32(800)
var VERTEX_COUNT = uint32(120)

type Terrain struct {
	X       float32
	Z       float32
	Model   *Models.RawModel
	Texture *Textures.ModelTexture
}

func NewTerrain(x, z float32, texture *Textures.ModelTexture) *Terrain {
	return &Terrain{
		x * SIZE,
		z * SIZE,
		generateTerrain(),
		texture,
	}
}

func generateTerrain() *Models.RawModel {
	count := VERTEX_COUNT * VERTEX_COUNT
	vertices := make([]float32, 0, count*3)
	normals := make([]float32, 0, count*3)
	textureCoords := make([]float32, 0, count*2)
	indices := make([]uint32, 0, 6*(VERTEX_COUNT-1)*(VERTEX_COUNT-1))
	// vertexIndex := 0
	for i := uint32(0); i < VERTEX_COUNT; i++ {
		for j := uint32(0); j < VERTEX_COUNT; j++ {
			vertices = append(vertices, float32(j)/(float32(VERTEX_COUNT) - 1)*SIZE)
			vertices = append(vertices, 0)
			vertices = append(vertices, float32(i)/(float32(VERTEX_COUNT) - 1)*SIZE)

			normals = append(normals, 0)
			normals = append(normals, 1)
			normals = append(normals, 0)

			textureCoords = append(textureCoords, float32(j)/(float32(VERTEX_COUNT) - 1))
			textureCoords = append(textureCoords, float32(i)/(float32(VERTEX_COUNT) - 1))

			// vertices[vertexIndex*3] = float32(j)/float32(VERTEX_COUNT) - 1*SIZE
			// vertices[vertexIndex*3+1] = 0
			// vertices[vertexIndex*3+2] = float32(i)/float32(VERTEX_COUNT) - 1*SIZE
			// normals[vertexIndex*3] = 0
			// normals[vertexIndex*3+1] = 1
			// normals[vertexIndex*3+2] = 0
			// textureCoords[vertexIndex*2] = float32(j)/float32(VERTEX_COUNT) - 1
			// textureCoords[vertexIndex*2+1] = float32(i)/float32(VERTEX_COUNT) - 1
			// vertexIndex++
		}
	}

	// index := 0
	for gz := uint32(0); gz < VERTEX_COUNT; gz++ {
		for gx := uint32(0); gx < VERTEX_COUNT; gx++ {
			topLeft := (gz * VERTEX_COUNT) + gx
			topRight := topLeft + 1
			bottomLeft := ((gz + 1) * VERTEX_COUNT) + gx
			bottomRight := bottomLeft + 1
			indices = append(indices, topLeft)
			indices = append(indices, bottomLeft)
			indices = append(indices, topRight)
			indices = append(indices, topRight)
			indices = append(indices, bottomLeft)
			indices = append(indices, bottomRight)
			// indices[index] = topLeft
			// index++
			// indices[index] = bottomLeft
			// index++
			// indices[index] = topRight
			// index++
			// indices[index] = topRight
			// index++
			// indices[index] = bottomLeft
			// index++
			// indices[index] = bottomRight
			// index++
		}
	}
	return Loaders.LoadToVAO(vertices, textureCoords, normals, indices)
}
