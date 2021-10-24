package Window

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var Window *glfw.Window
var LastFrameTime float32
var Delta float32

func CreateDisplay() {
	var err error
	err = glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Samples, 4)

	Window, err = glfw.CreateWindow(800, 600, "Go-Game", nil, nil)
	if err != nil {
		panic(err)
	}

	Window.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.MULTISAMPLE)

	LastFrameTime = getCurrentTime()
}

func UpdateDisplay() {
	Window.SwapBuffers()
	glfw.PollEvents()
	currentFrameTime := getCurrentTime()
	Delta = currentFrameTime - LastFrameTime
	LastFrameTime = getCurrentTime()
}

func CloseDisplay() {
	glfw.Terminate()
}

func getCurrentTime() float32 {
	return float32(glfw.GetTime())
}
