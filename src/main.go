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

	for !RenderEngine.Window.ShouldClose() {
		RenderEngine.Prepare()
		RenderEngine.Render(model)
		RenderEngine.UpdateDisplay()
	}

	RenderEngine.UpdateDisplay()
	RenderEngine.CloseDisplay()
}
