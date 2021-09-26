package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Shaders"
	"go-game/src/Systems"

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
	ProjectionMatrix mgl32.Mat4
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
		projectionMatrix,
	}
}

func (renderer *MasterRenderer) Render(light *Entities.Light, camera *Entities.Camera) {
	renderer.prepare()

	renderer.StaticShader.Start()
	renderer.StaticShader.LoadLight(light)
	renderer.StaticShader.LoadViewMatrix(camera)
	for _, system := range Systems.Systems {
		renderer.entityRenderer.Render(system.GetEntities())
	}
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
	// width, height := Window.Window.GetSize()
	// var top float32 = 16;
	// var bottom = -top
	// var right = top * float32(width) / float32(height)
	// var left = -right
	// return mgl32.Ortho(left, right, bottom, top, -1000, 1000)
}
