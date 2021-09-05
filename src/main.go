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
	RenderEngine.UpdateDisplay()
	RenderEngine.CloseDisplay()
}
