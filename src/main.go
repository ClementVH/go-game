package main

import (
	"runtime"

	"go-game/src/Entities"
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

	Systems.Systems = map[string]Systems.ISystem{
		"CHUNK_SYSTEM": Systems.NewChunkSystem(),
		"PLAYER_SYSTEM": playerSystem,
		"MONSTER_SYSTEM": Systems.NewMonsterSystem(),
	}

	light := Entities.NewLight(
		mgl32.Vec3{3333, 10000, -3333},
		mgl32.Vec3{1, 1, 1},
	)

	for !Window.Window.ShouldClose() {
		ToolBox.FpsCount()
		for _, system := range Systems.Systems {
			system.Tick()
		}
		renderer.Render(light, playerSystem.GetPlayer().Camera)
		Window.UpdateDisplay()
	}

	renderer.CleanUp()
	Window.CloseDisplay()
}
