package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Models"
	"go-game/src/Shaders"
	"go-game/src/Terrains"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const FOV = 70
const NEAR_PLANE = 0.1
const FAR_PLANE = 1000

type MasterRenderer struct {
	StaticShader   *Shaders.StaticShader
	TerrainShader   *Shaders.TerrainShader
	entities map[*Models.TexturedModel][]*Entities.Entity
	terrains []*Terrains.Terrain
	entityRenderer *EntityRenderer
	terrainRenderer *TerrainRenderer
}

func NewMasterRenderer() *MasterRenderer {
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)

	shader := Shaders.NewStaticShader()
	terrainShader := Shaders.NewTerrainShader()
	projectionMatrix := createProjectionMatrix()
	return &MasterRenderer{
		shader,
		terrainShader,
		make(map[*Models.TexturedModel][]*Entities.Entity),
		make([]*Terrains.Terrain, 0),
		NewEntityRenderer(shader, projectionMatrix),
		NewTerrainRenderer(terrainShader, projectionMatrix),
	}
}

func (renderer *MasterRenderer) Render(light *Entities.Light, camera *Entities.Camera) {
	renderer.prepare()

	renderer.StaticShader.Start()
	renderer.StaticShader.LoadLight(light)
	renderer.StaticShader.LoadViewMatrix(camera)
	renderer.entityRenderer.Render(renderer.entities)
	renderer.StaticShader.Stop()

	renderer.TerrainShader.Start()
	renderer.TerrainShader.LoadLight(light)
	renderer.TerrainShader.LoadViewMatrix(camera)
	renderer.terrainRenderer.Render(renderer.terrains)
	renderer.TerrainShader.Stop()

	renderer.entities = make(map[*Models.TexturedModel][]*Entities.Entity)
	renderer.terrains = make([]*Terrains.Terrain, 0)
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

func (renderer *MasterRenderer) ProcessTerrain(terrain *Terrains.Terrain) {
	renderer.terrains = append(renderer.terrains, terrain)
}

func (renderer *MasterRenderer) CleanUp() {
	renderer.StaticShader.CleanUp()
	renderer.TerrainShader.CleanUp()
}

func createProjectionMatrix() mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(FOV), float32(640)/480, NEAR_PLANE, FAR_PLANE)
}
