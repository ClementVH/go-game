package RenderEngine

import (
	"go-game/src/Shaders"
	"go-game/src/Terrains"
	"go-game/src/ToolBox"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type TerrainRenderer struct {
	shader *Shaders.TerrainShader
}

func NewTerrainRenderer(shader *Shaders.TerrainShader, matrix mgl32.Mat4) *TerrainRenderer {
	renderer := &TerrainRenderer{
		shader,
	}
	shader.Start()
	shader.LoadProjectionMatrix(matrix)
	shader.Stop()

	return renderer
}

func (renderer *TerrainRenderer) Render(terrains []*Terrains.Terrain) {
	for _, terrain := range terrains {
		renderer.prepareTerrain(terrain)
		renderer.loadModelMatrix(terrain)
		gl.DrawElements(gl.TRIANGLES, int32(terrain.Model.VertexCount), gl.UNSIGNED_INT, nil)
		renderer.unbindTexturedModel(terrain)
	}
}

func (renderer *TerrainRenderer) prepareTerrain(terrain *Terrains.Terrain) {
	rawModel := terrain.Model
	gl.BindVertexArray(rawModel.VaoID)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, terrain.Texture.TextureID)
}

func (renderer *TerrainRenderer) unbindTexturedModel(terrain *Terrains.Terrain) {
	rawModel := terrain.Model
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.BindVertexArray(0)
}

func (renderer *TerrainRenderer) loadModelMatrix(terrain *Terrains.Terrain) {
	transformationMatrix := ToolBox.CreateTransformationMatrix(
		mgl32.Vec3{terrain.X, 0, terrain.Z},
		0, 0, 0, 1,
	)

	renderer.shader.LoadTransformationMatrix(transformationMatrix)

}