package Shaders

import (
	"go-game/src/Entities"
	"go-game/src/ToolBox"

	"github.com/go-gl/mathgl/mgl32"
)

type StaticShader struct {
	ShaderProgram
	transformationMatrix int32
	projectionMatrix     int32
	viewMatrix           int32
	lightPosition        int32
	lightColor           int32
}

func NewStaticShader() *StaticShader {
	shader := StaticShader{ShaderProgram{}, 0, 0, 0, 0, 0}
	shader.ShaderProgram.IShaderProgram = &shader
	shader.create()
	return &shader
}

func (shader *StaticShader) bindAttributes() {
	shader.bindAttribute(0, "position")
	shader.bindAttribute(1, "textureCoords")
}

func (shader *StaticShader) getAllUniformLocations() {
	shader.transformationMatrix = shader.getUniformLocation("transformationMatrix")
	shader.projectionMatrix = shader.getUniformLocation("projectionMatrix")
	shader.viewMatrix = shader.getUniformLocation("viewMatrix")
	shader.lightPosition = shader.getUniformLocation("lightPosition")
	shader.lightColor = shader.getUniformLocation("lightColor")
}

func (shader *StaticShader) LoadTransformationMatrix(transformation mgl32.Mat4) {
	loadMatrix(shader.transformationMatrix, transformation)
}

func (shader *StaticShader) LoadProjectionMatrix(projection mgl32.Mat4) {
	loadMatrix(shader.projectionMatrix, projection)
}

func (shader *StaticShader) LoadViewMatrix(camera *Entities.Camera) {
	matrix := ToolBox.CreateViewMatrix(camera)
	loadMatrix(shader.viewMatrix, matrix)
}

func (shader *StaticShader) LoadLight(light *Entities.Light) {
	loadVector(shader.lightPosition, light.Position)
	loadVector(shader.lightColor, light.Color)
}
