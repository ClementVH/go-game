package main

import (
	"math/rand"
	"runtime"

	"go-game/src/Entities"
	"go-game/src/Loaders"
	"go-game/src/RenderEngine"
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

	camera := Entities.NewCamera()

	model := Loaders.LoadGltf("../res/duck", "Duck.gltf")

	entities := make([]*Entities.Entity, 0)
	for i := 0; i < 100; i++ {
		entity := Entities.NewEntity(
			model,
			mgl32.Vec3{randomFloat(-50, 50), randomFloat(-20, 5), randomFloat(-100, -50)},
			0, 0, 0, 0.1,
		)
		entities = append(entities, entity)
	}
	light := Entities.NewLight(
		mgl32.Vec3{0, 100, 0},
		mgl32.Vec3{1, 1, 1},
	)


	renderer := RenderEngine.NewMasterRenderer()
	for !RenderEngine.Window.ShouldClose() {
		ToolBox.FpsCount()
		for _, entity := range entities {
			entity.IncreaseRotation(0, -0.01, 0)
			renderer.ProcessEntity(entity)
		}
		renderer.Render(light, camera)
		RenderEngine.UpdateDisplay()
	}

	renderer.Shader.CleanUp()
	RenderEngine.CloseDisplay()
}

func randomFloat(min, max int) float32 {
	return float32(min + rand.Intn(max-min+1))
}
