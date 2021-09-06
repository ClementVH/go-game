package Shaders

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type IShaderProgram interface {
	bindAttributes()
}

type ShaderProgram struct {
	IShaderProgram
	programID        uint32
	vertexShaderID   uint32
	fragmentShaderID uint32
}

func (shader *ShaderProgram) create() {
	vertexShaderID := loadShader(gl.VERTEX_SHADER)
	fragmentShaderID := loadShader(gl.FRAGMENT_SHADER)
	programID := gl.CreateProgram()
	gl.AttachShader(programID, vertexShaderID)
	gl.AttachShader(programID, fragmentShaderID)
	shader.programID = programID
	shader.vertexShaderID = vertexShaderID
	shader.fragmentShaderID = fragmentShaderID
	shader.bindAttributes()
	gl.LinkProgram(programID)
	gl.ValidateProgram(programID)
}

func (shader *ShaderProgram) Start() {
	gl.UseProgram(shader.programID)
}

func Stop() {
	gl.UseProgram(0)
}

func (shader *ShaderProgram) CleanUp() {
	Stop()
	gl.DetachShader(shader.programID, shader.vertexShaderID)
	gl.DetachShader(shader.programID, shader.fragmentShaderID)
	gl.DeleteShader(shader.vertexShaderID)
	gl.DeleteShader(shader.fragmentShaderID)
	gl.DeleteProgram(shader.programID)
}

func (shader *ShaderProgram) bindAttribute(attribute uint32, name string) {
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
