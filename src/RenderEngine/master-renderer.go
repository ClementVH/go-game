package RenderEngine

import (
	"go-game/src/Entities"
	"go-game/src/Models"
	"go-game/src/Shaders"
)

type MasterRenderer struct {
	Shader   *Shaders.StaticShader
	entities map[*Models.TexturedModel][]*Entities.Entity
	renderer *Renderer
}

func NewMasterRenderer() *MasterRenderer {
	shader := Shaders.NewStaticShader()
	return &MasterRenderer{
		shader,
		make(map[*Models.TexturedModel][]*Entities.Entity),
		NewRenderer(shader),
	}
}

func (renderer *MasterRenderer) Render(light *Entities.Light, camera *Entities.Camera) {
	renderer.renderer.Prepare()
	renderer.Shader.Start()
	renderer.Shader.LoadLight(light)
	renderer.Shader.LoadViewMatrix(camera)
	renderer.renderer.Render(renderer.entities)
	Shaders.Stop()
	renderer.entities = make(map[*Models.TexturedModel][]*Entities.Entity)
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