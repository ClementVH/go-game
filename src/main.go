package main

import (
	"runtime"

	"go-game/src/Entities"
	"go-game/src/Loaders"
	"go-game/src/RenderEngine"
	"go-game/src/Shaders"
	"go-game/src/ToolBox"

	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {

	RenderEngine.CreateDisplay()

	staticShader := Shaders.NewStaticShader()
	RenderEngine.Setup(staticShader)
	texturedModels := Loaders.LoadGltf("../res/zelda", "scene.gltf")

	camera := Entities.NewCamera()
	var entities []*Entities.Entity
	for i := 0; i < len(texturedModels); i++ {
		entity := Entities.NewEntity(
			texturedModels[i],
			mgl32.Vec3{0, -5, -10},
			0, 0, 0, 0.05,
		)
		entities = append(entities, entity)
	}

	for !RenderEngine.Window.ShouldClose() {
		ToolBox.FpsCount()
		RenderEngine.Prepare()
		staticShader.Start()
		staticShader.LoadViewMatrix(camera)

		for i := 0; i < len(entities); i++ {
			entities[i].IncreaseRotation(0, -0.015, 0)
			RenderEngine.Render(entities[i], staticShader)
		}

		Shaders.Stop()
		RenderEngine.UpdateDisplay()
	}

	staticShader.CleanUp()
	RenderEngine.CloseDisplay()
}
