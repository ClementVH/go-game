package RenderEngine

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var Window *glfw.Window

func CreateDisplay() {
	var err error
	err = glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Samples, 4)

	Window, err = glfw.CreateWindow(640, 480, "Go-Game", nil, nil)
	if err != nil {
		panic(err)
	}

	Window.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.MULTISAMPLE)
}

func UpdateDisplay() {
	Window.SwapBuffers()
	glfw.PollEvents()
}

func CloseDisplay() {
	glfw.Terminate()
}
