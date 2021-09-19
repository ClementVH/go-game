package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Shaders"
	"go-game/src/State"
	"go-game/src/Window"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const FOV = 70
const NEAR_PLANE = 0.1
const FAR_PLANE = 1000

type MasterRenderer struct {
	StaticShader   *Shaders.StaticShader
	Entities[] Entities.IEntity
	entityRenderer *EntityRenderer
}

func NewMasterRenderer() *MasterRenderer {
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)

	shader := Shaders.NewStaticShader()
	projectionMatrix := createProjectionMatrix()
	return &MasterRenderer{
		shader,
		make([]Entities.IEntity, 0, 1024),
		NewEntityRenderer(shader, projectionMatrix),
	}
}

func (renderer *MasterRenderer) Render(light *Entities.Light, camera *Entities.Camera) {
	renderer.prepare()

	renderer.StaticShader.Start()
	renderer.StaticShader.LoadLight(light)
	renderer.StaticShader.LoadViewMatrix(camera)
	renderer.entityRenderer.Render(State.GetChunksToRender())
	renderer.entityRenderer.Render(renderer.Entities)
	renderer.StaticShader.Stop()
}

func (renderer *MasterRenderer) prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (renderer *MasterRenderer) CleanUp() {
	renderer.StaticShader.CleanUp()
}

func createProjectionMatrix() mgl32.Mat4 {
	width, height := Window.Window.GetSize()
	var top float32 = 16;
	var bottom = -top
	var right = top * float32(width) / float32(height)
	var left = -right
	return mgl32.Ortho(left, right, bottom, top, -1000, 1000)
}
