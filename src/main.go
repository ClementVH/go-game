package main

import (
	"runtime"

	"go-game/src/RenderEngine"
	"go-game/src/Shaders"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	RenderEngine.CreateDisplay()

	var vertices = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		0.5, 0.5, 0,
	}

	var indices = []uint32{
		0, 1, 3,
		3, 1, 2,
	}

	var model RenderEngine.RawModel = RenderEngine.LoadToVAO(vertices, indices)
	var staticShader = Shaders.CreateStaticShader()

	for !RenderEngine.Window.ShouldClose() {
		RenderEngine.Prepare()
		Shaders.Start(staticShader)
		RenderEngine.Render(model)
		Shaders.Stop()
		RenderEngine.UpdateDisplay()
	}

	Shaders.CleanUp(staticShader)
	RenderEngine.CloseDisplay()
}
