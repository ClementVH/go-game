package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Models"
	"go-game/src/Shaders"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const FOV = 70
const NEAR_PLANE = 0.1
const FAR_PLANE = 1000

type MasterRenderer struct {
	Shader   *Shaders.StaticShader
	entities map[*Models.TexturedModel][]*Entities.Entity
	renderer *EntityRenderer
}

func NewMasterRenderer() *MasterRenderer {
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)

	shader := Shaders.NewStaticShader()
	projectionMatrix := createProjectionMatrix()
	return &MasterRenderer{
		shader,
		make(map[*Models.TexturedModel][]*Entities.Entity),
		NewRenderer(shader, projectionMatrix),
	}
}

func (renderer *MasterRenderer) Render(light *Entities.Light, camera *Entities.Camera) {
	renderer.prepare()
	renderer.Shader.Start()
	renderer.Shader.LoadLight(light)
	renderer.Shader.LoadViewMatrix(camera)
	renderer.renderer.Render(renderer.entities)
	renderer.Shader.Stop()
	renderer.entities = make(map[*Models.TexturedModel][]*Entities.Entity)
}

func (renderer *MasterRenderer) prepare() {
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (renderer *MasterRenderer) ProcessEntity(entity *Entities.Entity) {
	for _, mesh := range entity.Meshes {
		var batch = renderer.entities[mesh]
		if (batch != nil) {
			batch = append(batch, entity)
			renderer.entities[mesh] = batch
		} else {
			batch := make([]*Entities.Entity, 0)
			batch = append(batch, entity)
			renderer.entities[mesh] = batch
		}
	}
}

func createProjectionMatrix() mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(FOV), float32(640)/480, NEAR_PLANE, FAR_PLANE)
}
