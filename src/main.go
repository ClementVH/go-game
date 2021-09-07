package main

import (
	"runtime"

	"go-game/src/Entities"
	"go-game/src/Loaders"
	"go-game/src/RenderEngine"
	"go-game/src/Shaders"

	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {

	RenderEngine.CreateDisplay()

	var staticShader = Shaders.NewStaticShader()
	RenderEngine.Setup(staticShader)
	texturedModel := Loaders.LoadGltf("cube.gltf")

	camera := Entities.NewCamera()

	var entity = Entities.NewEntity(
		texturedModel,
		mgl32.Vec3{0, 0, -5},
		0, 0, 0, 1,
	)

	for !RenderEngine.Window.ShouldClose() {
		entity.IncreaseRotation(-0.01, -0.01, 0)
		RenderEngine.Prepare()
		staticShader.Start()
		staticShader.LoadViewMatrix(camera)
		RenderEngine.Render(entity, staticShader)
		Shaders.Stop()
		RenderEngine.UpdateDisplay()
	}

	staticShader.CleanUp()
	RenderEngine.CloseDisplay()
}
