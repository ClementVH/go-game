package Shaders

import "github.com/go-gl/mathgl/mgl32"

type StaticShader struct {
	ShaderProgram
	transformationMatrix int32
}

func NewStaticShader() StaticShader {
	shader := StaticShader{ShaderProgram{}, 0}
	shader.ShaderProgram.IShaderProgram = &shader
	shader.create()
	return shader
}

func (shader *StaticShader) bindAttributes() {
	shader.bindAttribute(0, "position")
	shader.bindAttribute(1, "textureCoords")
}

func (shader *StaticShader) getAllUniformLocations() {
	shader.transformationMatrix = shader.getUniformLocation("transformationMatrix")
}

func (shader *StaticShader) LoadTransformationMatrix(matrix mgl32.Mat4) {
	loadMatrix(shader.transformationMatrix, matrix)
}
