package RenderEngine

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func LoadToVAO(positions []float32) RawModel {
	vaoID := createVAO()
	storeDataInAttributeList(0, positions)
	unbindVAO()
	return RawModel{
		vaoID:       vaoID,
		vertexCount: len(positions) / 3,
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

func unbindVAO() {
	gl.BindVertexArray(0)
}
