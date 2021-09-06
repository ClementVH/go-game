package Shaders

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderProgram struct {
	programID        uint32
	vertexShaderID   uint32
	fragmentShaderID uint32
}

func CreateShaderProgram() ShaderProgram {
	vertexShaderID := loadShader(gl.VERTEX_SHADER)
	fragmentShaderID := loadShader(gl.FRAGMENT_SHADER)
	programID := gl.CreateProgram()
	gl.AttachShader(programID, vertexShaderID)
	gl.AttachShader(programID, fragmentShaderID)
	gl.LinkProgram(programID)
	gl.ValidateProgram(programID)
	return ShaderProgram{
		programID:        programID,
		vertexShaderID:   vertexShaderID,
		fragmentShaderID: fragmentShaderID,
	}
}

func Start(shader ShaderProgram) {
	gl.UseProgram(shader.programID)
}

func Stop() {
	gl.UseProgram(0)
}

func CleanUp(shader ShaderProgram) {
	Stop()
	gl.DetachShader(shader.programID, shader.vertexShaderID)
	gl.DetachShader(shader.programID, shader.fragmentShaderID)
	gl.DeleteShader(shader.vertexShaderID)
	gl.DeleteShader(shader.fragmentShaderID)
	gl.DeleteProgram(shader.programID)
}

func bindAttribute(shader ShaderProgram, attribute uint32, name string) {
	nameArray := []uint8(name)
	gl.BindAttribLocation(shader.programID, attribute, &nameArray[0])
}

func loadShader(shaderType uint32) uint32 {
	var shaderCode string
	if shaderType == gl.FRAGMENT_SHADER {
		shaderCode = FragmentShader
	} else {
		shaderCode = VertexShader
	}
	csources, free := gl.Strs(shaderCode)
	shaderID := gl.CreateShader(shaderType)
	gl.ShaderSource(shaderID, 1, csources, nil)
	gl.CompileShader(shaderID)
	var status int32
	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderID, logLength, nil, gl.Str(log))

		panic(fmt.Errorf("failed to compile %v: %v", shaderCode, log))
	}
	free()
	return shaderID
}
