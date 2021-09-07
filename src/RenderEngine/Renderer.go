package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Shaders"
	"go-game/src/ToolBox"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const FOV = 70
const NEAR_PLANE = 0.1
const FAR_PLANE = 1000

func Setup(shader *Shaders.StaticShader) {
	matrix := createProjectionMatrix()
	shader.Start()
	shader.LoadProjectionMatrix(matrix)
	Shaders.Stop()
}

func Prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func Render(entity *Entities.Entity, shader *Shaders.StaticShader) {
	meshes := entity.Meshes
	for _, mesh := range meshes {
		rawModel := mesh.RawModel
		gl.BindVertexArray(rawModel.VaoID)
		gl.EnableVertexArrayAttrib(rawModel.VaoID, 0)
		gl.EnableVertexArrayAttrib(rawModel.VaoID, 1)
		transformationMatrix := ToolBox.CreateTransformationMatrix(
			entity.Position,
			entity.RotX, entity.RotY, entity.RotZ,
			entity.Scale,
		)

		shader.LoadTransformationMatrix(transformationMatrix)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, mesh.Texture.TextureID)
		gl.DrawElements(gl.TRIANGLES, int32(rawModel.VertexCount), gl.UNSIGNED_INT, nil)
		gl.DisableVertexArrayAttrib(rawModel.VaoID, 0)
		gl.DisableVertexArrayAttrib(rawModel.VaoID, 1)
		gl.BindVertexArray(0)
	}
}

func createProjectionMatrix() mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(FOV), float32(640)/480, NEAR_PLANE, FAR_PLANE)
}
