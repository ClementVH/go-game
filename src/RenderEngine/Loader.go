package RenderEngine

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func LoadToVAO(positions []float32, indices []uint32) RawModel {
	vaoID := createVAO()
	bindIndicesBuffer(indices)
	storeDataInAttributeList(0, positions)
	unbindVAO()
	return RawModel{
		vaoID:       vaoID,
		vertexCount: len(indices),
	}
}

func createVAO() uint32 {
	var vaoID uint32
	gl.GenVertexArrays(1, &vaoID)
	gl.BindVertexArray(vaoID)
	return vaoID
}

func storeDataInAttributeList(attributeNumber uint32, data []float32) {
	var vboID uint32
	gl.GenBuffers(1, &vboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vboID)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	gl.VertexAttribPointer(attributeNumber, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
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
