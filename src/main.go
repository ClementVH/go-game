package main

import (
	"runtime"

	"go-game/src/Models"
	"go-game/src/RenderEngine"
	"go-game/src/Shaders"
	"go-game/src/Textures"
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

	var textureCoords = []float32{
		0, 0,
		0, 1,
		1, 1,
		1, 0,
	}

	var staticShader = Shaders.NewStaticShader()
	var texturedModel = Models.TexturedModel{
		RawModel: RenderEngine.LoadToVAO(vertices, textureCoords, indices),
		Texture: Textures.ModelTexture{
			TextureID: RenderEngine.LoadTexture("texture.png"),
		},
	}

	for !RenderEngine.Window.ShouldClose() {
		RenderEngine.Prepare()
		staticShader.Start()
		RenderEngine.Render(texturedModel)
		Shaders.Stop()
		RenderEngine.UpdateDisplay()
	}

	staticShader.CleanUp()
	RenderEngine.CloseDisplay()
}
