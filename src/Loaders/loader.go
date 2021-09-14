package Loaders

import (
	"go-game/src/Models"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func LoadToVAO(positions, textureCoords, normals []float32, indices []uint32) *Models.RawModel {
	VaoID := createVAO()
	bindIndicesBuffer(indices)
	storeDataInAttributeList(0, 3, positions)
	storeDataInAttributeList(1, 2, textureCoords)
	storeDataInAttributeList(2, 3, normals)
	unbindVAO()
	return &Models.RawModel{
		VaoID:       VaoID,
		VertexCount: len(indices),
	}
}

func LoadTexture(folder string, file string) uint32 {
	imgFile, err := os.Open(folder + "/" + file)
	if err != nil {
		log.Fatalf("texture %q not found on disk: %v\n", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

func createVAO() uint32 {
	var VaoID uint32
	gl.GenVertexArrays(1, &VaoID)
	gl.BindVertexArray(VaoID)
	return VaoID
}

func storeDataInAttributeList(attributeNumber uint32, coordinateSize int32, data []float32) {
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	gl.VertexAttribPointer(attributeNumber, coordinateSize, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func bindIndicesBuffer(indices []uint32) {
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
}

func unbindVAO() {
	gl.BindVertexArray(0)
}
