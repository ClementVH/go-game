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

	camera := Entities.NewCamera()
	entity := Entities.NewEntity(
		Loaders.LoadGltf("../res/duck", "Duck.gltf"),
		mgl32.Vec3{0, -5, -10},
		0, 0, 0, 0.05,
	)
	light := Entities.NewLight(
		mgl32.Vec3{100, 0, 0},
		mgl32.Vec3{1, 1, 1},
	)

	for !RenderEngine.Window.ShouldClose() {
		ToolBox.FpsCount()
		entity.IncreaseRotation(0, -0.01, 0)
		RenderEngine.Prepare()
		staticShader.Start()
		staticShader.LoadLight(light)
		staticShader.LoadViewMatrix(camera)
		RenderEngine.Render(entity, staticShader)
		Shaders.Stop()
		RenderEngine.UpdateDisplay()
	}

	staticShader.CleanUp()
	RenderEngine.CloseDisplay()
}
