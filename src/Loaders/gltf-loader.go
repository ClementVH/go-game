package Loaders

import (
	"encoding/binary"
	"go-game/src/Models"
	"go-game/src/RenderEngine"
	"go-game/src/Textures"
	"math"

	"github.com/qmuntal/gltf"
)

func LoadGltf(filename string) Models.TexturedModel {
	doc, _ := gltf.Open("../res/" + filename)
	vertices := getFloats(doc, 1)
	indices := getIndices(doc, 0)
	textureCoords := getFloats(doc, 4)
	return Models.TexturedModel{
		RawModel: RenderEngine.LoadToVAO(vertices, textureCoords, indices),
		Texture: Textures.ModelTexture{
			TextureID: RenderEngine.LoadTexture("Cube_BaseColor.png"),
		},
	}
}

func getFloats(doc *gltf.Document, index int) []float32 {
	var floatArr []float32

	accessor := doc.Accessors[index]

	var buffViewInd = *accessor.BufferView
	var count = accessor.Count
	var accByteOffset = accessor.ByteOffset
	var accessorType = accessor.Type

	var bufferView = doc.BufferViews[buffViewInd]
	var byteOffset = bufferView.ByteOffset

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

	var beginningOfData = byteOffset + accByteOffset
	var lengthOfData = count * 4 * numPerVert
	for i := beginningOfData; i < beginningOfData+lengthOfData; i += 4 {
		var bytes = []byte{
			doc.Buffers[0].Data[i+0],
			doc.Buffers[0].Data[i+1],
			doc.Buffers[0].Data[i+2],
			doc.Buffers[0].Data[i+3],
		}
		var bits = binary.LittleEndian.Uint32(bytes)
		var float = math.Float32frombits(bits)

		floatArr = append(floatArr, float)
	}

	return floatArr
}

func getIndices(doc *gltf.Document, index int) []uint32 {
	var indicesArr []uint32

	accessor := doc.Accessors[index]

	var buffViewInd = *accessor.BufferView
	var count = accessor.Count
	var accByteOffset = accessor.ByteOffset
	var componentType = accessor.ComponentType

	var bufferView = doc.BufferViews[buffViewInd]
	var byteOffset = bufferView.ByteOffset

	var beginningOfData = byteOffset + accByteOffset
	if componentType == gltf.ComponentUint {
		lengthOfData := count * 4
		for i := beginningOfData; i < beginningOfData+lengthOfData; i += 4 {
			var bytes = []byte{
				doc.Buffers[0].Data[i+0],
				doc.Buffers[0].Data[i+1],
				doc.Buffers[0].Data[i+2],
				doc.Buffers[0].Data[i+3],
			}
			var index = binary.LittleEndian.Uint32(bytes)
			indicesArr = append(indicesArr, index)
		}
	}

	if componentType == gltf.ComponentUshort {
		lengthOfData := count * 2
		for i := beginningOfData; i < beginningOfData+lengthOfData; i += 2 {
			var bytes = []byte{
				doc.Buffers[0].Data[i+0],
				doc.Buffers[0].Data[i+1],
			}
			var index = binary.LittleEndian.Uint16(bytes)
			indicesArr = append(indicesArr, uint32(index))
		}
	}

	if componentType == gltf.ComponentShort {
		lengthOfData := count * 2
		for i := beginningOfData; i < beginningOfData+lengthOfData; i += 2 {
			var bytes = []byte{
				doc.Buffers[0].Data[i+0],
				doc.Buffers[0].Data[i+1],
			}
			var index = binary.LittleEndian.Uint16(bytes)
			indicesArr = append(indicesArr, uint32(index))
		}
	}

	return indicesArr
}
