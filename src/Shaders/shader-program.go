package Shaders

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type IShaderProgram interface {
	bindAttributes()
	getAllUniformLocations()
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
	shader.getAllUniformLocations()
}

func (shader *ShaderProgram) getUniformLocation(name string) int32 {
	nameArray := []uint8(name)
	return gl.GetUniformLocation(shader.programID, &nameArray[0])
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

func loadFloat(location int32, value float32) {
	gl.Uniform1f(location, value)
}

func loadVector(location int32, vector mgl32.Vec3) {
	gl.Uniform3f(location, vector.X(), vector.Y(), vector.Z())
}

func loadBoolean(location int32, value bool) {
	var toLoad float32 = 0
	if value {
		toLoad = 1
	}
	gl.Uniform1f(location, toLoad)
}

func loadMatrix(location int32, matrix mgl32.Mat4) {
	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
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
