package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Models"
	"go-game/src/Shaders"
	"go-game/src/ToolBox"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const FOV = 70
const NEAR_PLANE = 0.1
const FAR_PLANE = 1000

type Renderer struct {
	projectionMatrix mgl32.Mat4
	shader           *Shaders.StaticShader
}

func NewRenderer(shader *Shaders.StaticShader) *Renderer {
	renderer := &Renderer{
		createProjectionMatrix(),
		shader,
	}
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	shader.Start()
	shader.LoadProjectionMatrix(renderer.projectionMatrix)
	Shaders.Stop()

	return renderer
}

func (renderer *Renderer) Prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (renderer *Renderer) Render(entities map[*Models.TexturedModel][]*Entities.Entity) {
	for model, batch := range entities {
		renderer.prepareTexturedModel(model)
		for _, entity := range batch {
			renderer.prepareInstance(entity)
			gl.DrawElements(gl.TRIANGLES, int32(model.RawModel.VertexCount), gl.UNSIGNED_INT, nil)
		}
		renderer.unbindTexturedModel(model)
	}
}

func (renderer *Renderer) prepareTexturedModel(model *Models.TexturedModel) {
	rawModel := model.RawModel
	gl.BindVertexArray(rawModel.VaoID)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.EnableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, model.Texture.TextureID)
}

func (renderer *Renderer) unbindTexturedModel(model *Models.TexturedModel) {
	rawModel := model.RawModel
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 0)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 1)
	gl.DisableVertexArrayAttrib(rawModel.VaoID, 2)
	gl.BindVertexArray(0)
}

func (renderer *Renderer) prepareInstance(entity *Entities.Entity) {
	transformationMatrix := ToolBox.CreateTransformationMatrix(
		entity.Position,
		entity.RotX, entity.RotY, entity.RotZ,
		entity.Scale,
	)

	renderer.shader.LoadTransformationMatrix(transformationMatrix)

}

func createProjectionMatrix() mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(FOV), float32(640)/480, NEAR_PLANE, FAR_PLANE)
}
