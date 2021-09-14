package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Models"
	"go-game/src/Shaders"
	"go-game/src/ToolBox"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type EntityRenderer struct {
	shader           *Shaders.StaticShader
}

func NewEntityRenderer(shader *Shaders.StaticShader, matrix mgl32.Mat4) *EntityRenderer {
	renderer := &EntityRenderer{
		shader,
	}
	shader.Start()
	shader.LoadProjectionMatrix(matrix)
	shader.Stop()

	return renderer
}

func (renderer *EntityRenderer) Render(entities []*Entities.Entity) {
	for _, entity := range entities {
		for _, mesh := range entity.Meshes {
			renderer.prepareTexturedModel(mesh)
			renderer.prepareInstance(entity)
			gl.DrawElements(gl.TRIANGLES, int32(mesh.RawModel.VertexCount), gl.UNSIGNED_INT, nil)
		}
	}
}

func (renderer *EntityRenderer) prepareTexturedModel(model *Models.TexturedModel) {
	rawModel := model.RawModel
	gl.BindVertexArray(rawModel.VaoID)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, model.Texture.TextureID)
}

func (renderer *EntityRenderer) unbindTexturedModel(model *Models.TexturedModel) {
	rawModel := model.RawModel
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.BindVertexArray(0)
}

func (renderer *EntityRenderer) prepareInstance(entity *Entities.Entity) {
	transformationMatrix := ToolBox.CreateTransformationMatrix(
		entity.Position,
		entity.RotX, entity.RotY, entity.RotZ,
		entity.Scale,
	)

	renderer.shader.LoadTransformationMatrix(transformationMatrix)

}