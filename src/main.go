package main

import (
	"runtime"

	"go-game/src/Entities"
	"go-game/src/Loaders"
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

	chunkManager := Systems.NewChunkSystem()

	renderer := RenderEngine.NewMasterRenderer()

	model := Loaders.LoadGltf("../res/plane", "plane.gltf")

	for x := 0; x < Systems.WORLD_CHUNKS_SIZE; x++ {
		for z := 0; z < Systems.WORLD_CHUNKS_SIZE; z++ {
			Systems.ChunkEntities[x * Systems.WORLD_CHUNKS_SIZE + z] = Entities.NewChunk(
				model,
				x - Systems.WORLD_CHUNKS_SIZE / 2,
				z - Systems.WORLD_CHUNKS_SIZE / 2,
			)
		}
	}

	player := Entities.NewPlayer(
		Loaders.LoadGltf("../res/player", "player.gltf"),
		mgl32.Vec3{8, 2, -8},
		0, 0, 0, 1,
	)
	Systems.Player = player
	renderer.Entities = append(renderer.Entities, player)

	light := Entities.NewLight(
		mgl32.Vec3{50, 100, 0},
		mgl32.Vec3{1, 1, 1},
	)

	for !Window.Window.ShouldClose() {
		ToolBox.FpsCount()
		chunkManager.Tick()
		player.Move()
		renderer.Render(light, player.Camera)
		Window.UpdateDisplay()
	}

	renderer.CleanUp()
	Window.CloseDisplay()
}
