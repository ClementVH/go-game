package RenderEngine

import (
	"go-game/src/Models"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func Prepare() {
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Render(model Models.TexturedModel) {
	var rawModel = model.RawModel
	gl.BindVertexArray(rawModel.VaoID)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, model.Texture.TextureID)
	gl.DrawElements(gl.TRIANGLES, int32(rawModel.VertexCount), gl.UNSIGNED_INT, nil)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.BindVertexArray(0)
}
