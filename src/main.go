package main

import (
	"runtime"

	"go-game/src/Entities"
	"go-game/src/Loaders"
	"go-game/src/RenderEngine"
	"go-game/src/ToolBox"
	"go-game/src/Window"

	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {

	Window.CreateDisplay()

	renderer := RenderEngine.NewMasterRenderer()

	model := Loaders.LoadGltf("../res/plane", "plane.gltf")

	entity := Entities.NewEntity(
		model,
		mgl32.Vec3{0, 0, 0},
		0, 0, 0, 1,
	)
	renderer.Entities = append(renderer.Entities, entity)

	character := Entities.NewCharacter(
		Loaders.LoadGltf("../res/character", "character.gltf"),
		mgl32.Vec3{8, 2, -8},
		0, 0, 0, 1,
	)
	renderer.Entities = append(renderer.Entities, character)

	light := Entities.NewLight(
		mgl32.Vec3{50, 100, 0},
		mgl32.Vec3{1, 1, 1},
	)

	for !Window.Window.ShouldClose() {
		ToolBox.FpsCount()
		character.Move()
		renderer.Render(light, character.Camera)
		Window.UpdateDisplay()
	}

	renderer.CleanUp()
	Window.CloseDisplay()
}
