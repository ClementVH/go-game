package RenderEngine

import (
	"go-game/src/Constants"
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
	StaticShader     *Shaders.StaticShader
	ChunkShader      *Shaders.ChunkShader
	entityRenderer   *EntityRenderer
	chunkRenderer    *ChunkRenderer
	ProjectionMatrix mgl32.Mat4
}

func NewMasterRenderer() *MasterRenderer {
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)

	shader := Shaders.NewStaticShader()
	chunkShader := Shaders.NewChunkShader()
	projectionMatrix := createProjectionMatrix()
	return &MasterRenderer{
		shader,
		chunkShader,
		NewEntityRenderer(shader, projectionMatrix),
		NewChunkRenderer(chunkShader, projectionMatrix),
		projectionMatrix,
	}
}

func (renderer *MasterRenderer) Render(light *Entities.Light, camera *Entities.Camera) {
	renderer.prepare()

	renderer.ChunkShader.Start()
	renderer.ChunkShader.LoadLight(light)
	renderer.ChunkShader.LoadViewMatrix(camera)
	renderer.ChunkShader.LoadCombatChunk()
	renderer.chunkRenderer.Render(State.Systems.ChunkSystem.GetEntities())
	renderer.ChunkShader.Stop()

	renderer.StaticShader.Start()
	renderer.StaticShader.LoadLight(light)
	renderer.StaticShader.LoadViewMatrix(camera)
	renderer.entityRenderer.Render(State.Systems.PlayerSystem.GetEntities())
	renderer.entityRenderer.Render(State.Systems.WildMonsterSystem.GetEntities())
	renderer.StaticShader.Stop()
}

func (renderer *MasterRenderer) prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
}

func (renderer *MasterRenderer) CleanUp() {
	renderer.StaticShader.CleanUp()
}

func createProjectionMatrix() mgl32.Mat4 {
	if Constants.PROJECTION == "ORTHO" {
		width, height := Window.Window.GetSize()
		var top float32 = 16
		var bottom = -top
		var right = top * float32(width) / float32(height)
		var left = -right
		return mgl32.Ortho(left, right, bottom, top, 0.001, 1000)
	} else {
		return mgl32.Perspective(mgl32.DegToRad(FOV), float32(640)/480, NEAR_PLANE, FAR_PLANE)
	}
}
