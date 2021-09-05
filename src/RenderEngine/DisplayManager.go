package RenderEngine

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var window *glfw.Window

func CreateDisplay() {
	var err error
	err = glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Samples, 4)

	window, err = glfw.CreateWindow(640, 480, "Go-Game", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.MULTISAMPLE)
}

func UpdateDisplay() {
	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.ClearColor(1.0, 0, 0, 0.0)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func CloseDisplay() {
	glfw.Terminate()
}
