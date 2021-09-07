package ToolBox

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var lastTime = 0.0
var currentTime = 0.0
var numberFrames = 0

func FpsCount() {
	currentTime = glfw.GetTime()
	if currentTime-lastTime >= 1.0 {
		fmt.Println("FPS: " + fmt.Sprint(numberFrames))
		lastTime = currentTime
		numberFrames = 0
	}
	numberFrames++
}
