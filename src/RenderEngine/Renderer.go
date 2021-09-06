package RenderEngine

import "github.com/go-gl/gl/v3.3-core/gl"

func Prepare() {
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Render(model RawModel) {
	gl.BindVertexArray(model.vaoID)
	gl.EnableVertexArrayAttrib(model.vaoID, 0)
	gl.DrawElements(gl.TRIANGLES, int32(model.vertexCount), gl.UNSIGNED_INT, nil)
	gl.DisableVertexArrayAttrib(model.vaoID, 0)
	gl.BindVertexArray(0)
}
