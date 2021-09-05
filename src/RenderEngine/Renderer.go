package RenderEngine

import "github.com/go-gl/gl/v3.3-core/gl"

func Prepare() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(1, 0, 0, 1)
}

func Render(model RawModel) {
	gl.BindVertexArray(model.vaoID)
	gl.EnableVertexArrayAttrib(model.vaoID, 0)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(model.vertexCount))
	gl.DisableVertexArrayAttrib(model.vaoID, 0)
	gl.BindVertexArray(0)
}
