package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Shaders"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const FOV = 70
const NEAR_PLANE = 0.1
const FAR_PLANE = 1000

type MasterRenderer struct {
	StaticShader   *Shaders.StaticShader
	Entities[] *Entities.Entity
	entityRenderer *EntityRenderer
}

func NewMasterRenderer() *MasterRenderer {
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)

	shader := Shaders.NewStaticShader()
	projectionMatrix := createProjectionMatrix()
	return &MasterRenderer{
		shader,
		make([]*Entities.Entity, 0, 1024),
		NewEntityRenderer(shader, projectionMatrix),
	}
}

func (renderer *MasterRenderer) Render(light *Entities.Light, camera *Entities.Camera) {
	renderer.prepare()

	renderer.StaticShader.Start()
	renderer.StaticShader.LoadLight(light)
	renderer.StaticShader.LoadViewMatrix(camera)
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
	return mgl32.Perspective(mgl32.DegToRad(FOV), float32(640)/480, NEAR_PLANE, FAR_PLANE)
}
