package main

import (
	"math/rand"
	"runtime"
	"time"

	"go-game/src/Controllers"
	"go-game/src/Entities"
	"go-game/src/GUI"
	"go-game/src/RenderEngine"
	"go-game/src/State"
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
	rand.Seed(time.Now().UTC().UnixNano())

	Window.CreateDisplay()

	renderer := RenderEngine.NewMasterRenderer()

	State.Systems.SetPlayerSystem(Systems.NewPlayerSystem())
	State.Systems.SetChunkSystem(Systems.NewChunkSystem())
	State.Systems.SetWildMonsterSystem(Systems.NewWildMonsterSystem())
	gui := GUI.NewGUI()

	Controllers.InitCameraMovements()

	light := Entities.NewLight(
		mgl32.Vec3{3333, 10000, -3333},
		mgl32.Vec3{1, 1, 1},
	)

	for !Window.Window.ShouldClose() {
		ToolBox.FpsCount()
		for _, system := range State.Systems.GetAll() {
			system.Tick()
		}
		gui.Update()

		renderer.Render(light, State.Camera.Camera)
		Window.UpdateDisplay()
	}

	renderer.CleanUp()
	Window.CloseDisplay()
}
