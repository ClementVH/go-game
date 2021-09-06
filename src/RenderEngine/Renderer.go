package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Shaders"
	"go-game/src/ToolBox"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func Prepare() {
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Render(entity Entities.Entity, shader Shaders.StaticShader) {
	var model = entity.Model
	var rawModel = model.RawModel
	gl.BindVertexArray(rawModel.VaoID)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 1)
	var transformationMatrix = ToolBox.CreateTransformationMatrix(
		entity.Position,
		entity.RotX, entity.RotY, entity.RotZ,
		entity.Scale,
	)

	shader.LoadTransformationMatrix(transformationMatrix)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, model.Texture.TextureID)
	gl.DrawElements(gl.TRIANGLES, int32(rawModel.VertexCount), gl.UNSIGNED_INT, nil)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.BindVertexArray(0)
}
