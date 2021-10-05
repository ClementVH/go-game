package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Models"
	"go-game/src/Shaders"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type ChunkRenderer struct {
	shader *Shaders.ChunkShader
}

func NewChunkRenderer(shader *Shaders.ChunkShader, matrix mgl32.Mat4) *ChunkRenderer {
	renderer := &ChunkRenderer{
		shader,
	}
	shader.Start()
	shader.LoadProjectionMatrix(matrix)
	shader.Stop()

	return renderer
}

func (renderer *ChunkRenderer) Render(entities []*Entities.Entity) {
	for _, entity := range entities {
		for _, mesh := range entity.GetMeshes() {
			renderer.prepareTexturedModel(mesh)
			renderer.prepareInstance(entity)
			gl.DrawElements(gl.TRIANGLES, int32(mesh.RawModel.VertexCount), gl.UNSIGNED_INT, nil)
			renderer.unbindTexturedModel(mesh)
		}
	}
}

func (renderer *ChunkRenderer) prepareTexturedModel(model *Models.TexturedModel) {
	rawModel := model.RawModel
	gl.BindVertexArray(rawModel.VaoID)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, model.Texture.TextureID)
}

func (renderer *ChunkRenderer) unbindTexturedModel(model *Models.TexturedModel) {
	rawModel := model.RawModel
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.BindVertexArray(0)
}

func (renderer *ChunkRenderer) prepareInstance(entity *Entities.Entity) {
	transformationMatrix := entity.GetTransformationMatrix()
	renderer.shader.LoadTransformationMatrix(transformationMatrix)
}