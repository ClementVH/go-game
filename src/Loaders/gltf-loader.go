package Loaders

import (
	"encoding/binary"
	"go-game/src/Models"
	"go-game/src/RenderEngine"
	"go-game/src/Textures"
	"math"

	"github.com/qmuntal/gltf"
)

func LoadGltf(filename string) *Models.TexturedModel {
	doc, _ := gltf.Open("../res/" + filename)
	vertices := getFloats(doc, 1)
	indices := getIndices(doc, 0)
	textureCoords := getFloats(doc, 4)
	return &Models.TexturedModel{
		RawModel: RenderEngine.LoadToVAO(vertices, textureCoords, indices),
		Texture: Textures.ModelTexture{
			TextureID: RenderEngine.LoadTexture("Cube_BaseColor.png"),
		},
	}
}

func getFloats(doc *gltf.Document, index int) []float32 {
	floatArr := make([]float32, 0)

	accessor := doc.Accessors[index]

	buffViewInd := *accessor.BufferView
	count := accessor.Count
	accByteOffset := accessor.ByteOffset
	accessorType := accessor.Type

	bufferView := doc.BufferViews[buffViewInd]
	byteOffset := bufferView.ByteOffset

	var numPerVert uint32
	switch accessorType {
	case gltf.AccessorScalar:
		numPerVert = 1
	case gltf.AccessorVec2:
		numPerVert = 2
	case gltf.AccessorVec3:
		numPerVert = 3
	case gltf.AccessorVec4:
		numPerVert = 4
	}

	beginningOfData := byteOffset + accByteOffset
	lengthOfData := count * 4 * numPerVert
	for i := beginningOfData; i < beginningOfData+lengthOfData; i += 4 {
		bytes := []byte{
			doc.Buffers[0].Data[i+0],
			doc.Buffers[0].Data[i+1],
			doc.Buffers[0].Data[i+2],
			doc.Buffers[0].Data[i+3],
		}
		bits := binary.LittleEndian.Uint32(bytes)
		float := math.Float32frombits(bits)

		floatArr = append(floatArr, float)
	}

	return floatArr
}

func getIndices(doc *gltf.Document, index int) []uint32 {
	accessor := doc.Accessors[index]

	buffViewInd := *accessor.BufferView
	count := accessor.Count
	accByteOffset := accessor.ByteOffset
	componentType := accessor.ComponentType

	bufferView := doc.BufferViews[buffViewInd]
	byteOffset := bufferView.ByteOffset

	indicesArr := make([]uint32, 0)
	beginningOfData := byteOffset + accByteOffset
	if componentType == gltf.ComponentUint {
		lengthOfData := count * 4
		for i := beginningOfData; i < beginningOfData+lengthOfData; i += 4 {
			bytes := []byte{
				doc.Buffers[0].Data[i+0],
				doc.Buffers[0].Data[i+1],
				doc.Buffers[0].Data[i+2],
				doc.Buffers[0].Data[i+3],
			}
			index := binary.LittleEndian.Uint32(bytes)
			indicesArr = append(indicesArr, index)
		}
	}

	if componentType == gltf.ComponentUshort {
		lengthOfData := count * 2
		for i := beginningOfData; i < beginningOfData+lengthOfData; i += 2 {
			bytes := []byte{
				doc.Buffers[0].Data[i+0],
				doc.Buffers[0].Data[i+1],
			}
			index := binary.LittleEndian.Uint16(bytes)
			indicesArr = append(indicesArr, uint32(index))
		}
	}

	if componentType == gltf.ComponentShort {
		lengthOfData := count * 2
		for i := beginningOfData; i < beginningOfData+lengthOfData; i += 2 {
			bytes := []byte{
				doc.Buffers[0].Data[i+0],
				doc.Buffers[0].Data[i+1],
			}
			index := binary.LittleEndian.Uint16(bytes)
			indicesArr = append(indicesArr, uint32(index))
		}
	}

	return indicesArr
}