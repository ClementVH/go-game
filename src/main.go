package main

import (
	"runtime"

	"go-game/src/Entities"
	"go-game/src/Models"
	"go-game/src/RenderEngine"
	"go-game/src/Shaders"
	"go-game/src/Textures"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	RenderEngine.CreateDisplay()

	var vertices = []float32{
		-0.5, 0.5, -0.5,
		-0.5, -0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, 0.5, -0.5,

		-0.5, 0.5, 0.5,
		-0.5, -0.5, 0.5,
		0.5, -0.5, 0.5,
		0.5, 0.5, 0.5,

		0.5, 0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, -0.5, 0.5,
		0.5, 0.5, 0.5,

		-0.5, 0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, -0.5, 0.5,
		-0.5, 0.5, 0.5,

		-0.5, 0.5, 0.5,
		-0.5, 0.5, -0.5,
		0.5, 0.5, -0.5,
		0.5, 0.5, 0.5,

		-0.5, -0.5, 0.5,
		-0.5, -0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, -0.5, 0.5,
	}

	var indices = []uint32{
		0, 1, 3,
		3, 1, 2,
		4, 5, 7,
		7, 5, 6,
		8, 9, 11,
		11, 9, 10,
		12, 13, 15,
		15, 13, 14,
		16, 17, 19,
		19, 17, 18,
		20, 21, 23,
		23, 21, 22,
	}

	var textureCoords = []float32{
		0, 0,
		0, 1,
		1, 1,
		1, 0,
		0, 0,
		0, 1,
		1, 1,
		1, 0,
		0, 0,
		0, 1,
		1, 1,
		1, 0,
		0, 0,
		0, 1,
		1, 1,
		1, 0,
		0, 0,
		0, 1,
		1, 1,
		1, 0,
		0, 0,
		0, 1,
		1, 1,
		1, 0,
	}

	var staticShader = Shaders.NewStaticShader()
	RenderEngine.Setup(staticShader)
	var texturedModel = Models.TexturedModel{
		RawModel: RenderEngine.LoadToVAO(vertices, textureCoords, indices),
		Texture: Textures.ModelTexture{
			TextureID: RenderEngine.LoadTexture("texture2.png"),
		},
	}

	var camera = Entities.NewCamera()
	glfw.GetCurrentContext().SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		var keyName = glfw.GetKeyName(key, scancode)
		switch keyName {
		case "z":
			camera.Move(keyName)
		case "q":
			camera.Move(keyName)
		case "s":
			camera.Move(keyName)
		case "d":
			camera.Move(keyName)
		}
	})

	var entity = Entities.NewEntity(
		texturedModel,
		mgl32.Vec3{0, 0, -2},
		0, 0, 0, 1,
	)

	for !RenderEngine.Window.ShouldClose() {
		entity.IncreaseRotation(-0.002, -0.002, -0.002)
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
