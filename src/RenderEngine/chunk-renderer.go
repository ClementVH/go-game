package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Loaders"
	"go-game/src/Models"
	"go-game/src/Shaders"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type ChunkRenderer struct {
	shader      *Shaders.ChunkShader
	gridBlendId uint32
}

func NewChunkRenderer(shader *Shaders.ChunkShader, matrix mgl32.Mat4) *ChunkRenderer {
	renderer := &ChunkRenderer{
		shader,
		Loaders.LoadTexture("../res/textures", "grid-blend.png"),
	}
	shader.Start()
	shader.LoadProjectionMatrix(matrix)
	shader.LoadTextures()
	shader.Stop()

	return renderer
}

func (renderer *ChunkRenderer) Render(entities []Entities.IEntity) {
	for _, entity := range entities {
		for _, mesh := range entity.GetMeshes() {
			renderer.prepareTexturedModel(mesh, entity)
			renderer.prepareInstance(entity)
			gl.DrawElements(gl.TRIANGLES, int32(mesh.RawModel.VertexCount), gl.UNSIGNED_INT, nil)
			renderer.unbindTexturedModel(mesh)
		}
	}
}

func (renderer *ChunkRenderer) prepareTexturedModel(model *Models.TexturedModel, entity interface{}) {
	var chunk = entity.(*Entities.Chunk)
	rawModel := model.RawModel
	gl.BindVertexArray(rawModel.VaoID)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, model.Texture.TextureID)
	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, renderer.gridBlendId)
	gl.ActiveTexture(gl.TEXTURE2)
	gl.BindTexture(gl.TEXTURE_2D, chunk.StartPositionsId)
}

func (renderer *ChunkRenderer) unbindTexturedModel(model *Models.TexturedModel) {
	rawModel := model.RawModel
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.BindVertexArray(0)
}

func (renderer *ChunkRenderer) prepareInstance(entity Entities.IEntity) {
	transformationMatrix := entity.GetTransformationMatrix()
	renderer.shader.LoadTransformationMatrix(transformationMatrix)
}
