package main

import (
	"fmt"
	"math"
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
	chunkSystem := Systems.NewChunkSystem()
	wildMonsterSystem := Systems.NewWildMonsterSystem()

	Systems.Systems = map[string]Systems.ISystem{
		"CHUNK_SYSTEM": chunkSystem,
		"PLAYER_SYSTEM": playerSystem,
		"WILD_MONSTER_SYSTEM": wildMonsterSystem,
	}

	light := Entities.NewLight(
		mgl32.Vec3{3333, 10000, -3333},
		mgl32.Vec3{1, 1, 1},
	)

	raycast := ToolBox.NewRaycast(playerSystem.GetPlayer().Camera, renderer.ProjectionMatrix)

	for !Window.Window.ShouldClose() {
		ToolBox.FpsCount()
		for _, system := range Systems.Systems {
			system.Tick()
		}
		raycast.Update()

		var startPos = raycast.RayOrigin
		var minMul float32 = 0
		var middleMul float32 = 50
		var maxMul float32 = 100
		var endPos = startPos.Add(raycast.Ray.Mul(middleMul))

		for i := 0; i < 200; i++ {
			if endPos.Y() < 0 {
				maxMul = middleMul
			} else if endPos.Y() >= 0 {
				minMul = middleMul
			}
			middleMul = minMul + (maxMul - minMul) / 2
			endPos = startPos.Add(raycast.Ray.Mul(middleMul))
		}

		fmt.Println(math.Floor(float64(endPos.X()) / 16), math.Floor(float64(endPos.Z()) / 16))

		renderer.Render(light, playerSystem.GetPlayer().Camera)
		Window.UpdateDisplay()
	}

	renderer.CleanUp()
	Window.CloseDisplay()
}
