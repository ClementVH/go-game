package main

import (
	"runtime"

	"go-game/src/RenderEngine"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	RenderEngine.CreateDisplay()

	var vertices = []float32{
		// Left Bottom triangle
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		// Right Top triangle
		0.5, -0.5, 0,
		0.5, 0.5, 0,
		-0.5, 0.5, 0,
	}

	var model RenderEngine.RawModel = RenderEngine.LoadToVAO(vertices)

	for !RenderEngine.Window.ShouldClose() {
		RenderEngine.Prepare()
		RenderEngine.Render(model)
		RenderEngine.UpdateDisplay()
	}

	RenderEngine.UpdateDisplay()
	RenderEngine.CloseDisplay()
}
