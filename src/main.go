package main

import (
	"runtime"

	"go-game/src/Entities"
	"go-game/src/GUI"
	"go-game/src/MousePicker"
	"go-game/src/RenderEngine"
	"go-game/src/Systems"
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

	playerSystem := Systems.NewPlayerSystem()
	chunkSystem := Systems.NewChunkSystem()
	wildMonsterSystem := Systems.NewWildMonsterSystem()

	Systems.Systems = map[string]Systems.ISystem{
		"CHUNK_SYSTEM":        chunkSystem,
		"PLAYER_SYSTEM":       playerSystem,
		"WILD_MONSTER_SYSTEM": wildMonsterSystem,
	}

	light := Entities.NewLight(
		mgl32.Vec3{3333, 10000, -3333},
		mgl32.Vec3{1, 1, 1},
	)

	gui := GUI.NewGUI()

	MousePicker.InitEntityPicker(playerSystem.GetPlayer().Camera, renderer.ProjectionMatrix)

	for !Window.Window.ShouldClose() {
		ToolBox.FpsCount()
		for _, system := range Systems.Systems {
			system.Tick()
		}
		gui.Update()
		MousePicker.Picker.Update()

		renderer.Render(light, playerSystem.GetPlayer().Camera)
		Window.UpdateDisplay()
	}

	renderer.CleanUp()
	Window.CloseDisplay()
}
