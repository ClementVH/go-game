package Shaders

import (
	"go-game/src/Entities"
	"go-game/src/ToolBox"

	"github.com/go-gl/mathgl/mgl32"
)

type TerrainShader struct {
	transformationMatrix int32
	projectionMatrix     int32
	viewMatrix           int32
	lightPosition        int32
	lightColor           int32
	ShaderProgram
}

func NewTerrainShader() *TerrainShader {
	shaderProgram := NewShaderProgram(TerrainVertexShader, TerrainFragmentShader)
	shader := TerrainShader{0, 0, 0, 0, 0, shaderProgram}
	shader.bindAttributes()
	shader.setup()
	shader.getAllUniformLocations()
	return &shader
}

func (shader *TerrainShader) bindAttributes() {
	shader.bindAttribute(0, "position")
	shader.bindAttribute(1, "textureCoords")
}

func (shader *TerrainShader) getAllUniformLocations() {
	shader.transformationMatrix = shader.getUniformLocation("transformationMatrix")
	shader.projectionMatrix = shader.getUniformLocation("projectionMatrix")
	shader.viewMatrix = shader.getUniformLocation("viewMatrix")
	shader.lightPosition = shader.getUniformLocation("lightPosition")
	shader.lightColor = shader.getUniformLocation("lightColor")
}

func (shader *TerrainShader) LoadTransformationMatrix(transformation mgl32.Mat4) {
	loadMatrix(shader.transformationMatrix, transformation)
}

func (shader *TerrainShader) LoadProjectionMatrix(projection mgl32.Mat4) {
	loadMatrix(shader.projectionMatrix, projection)
}

func (shader *TerrainShader) LoadViewMatrix(camera *Entities.Camera) {
	matrix := ToolBox.CreateViewMatrix(camera)
	loadMatrix(shader.viewMatrix, matrix)
}

func (shader *TerrainShader) LoadLight(light *Entities.Light) {
	loadVector(shader.lightPosition, light.Position)
	loadVector(shader.lightColor, light.Color)
}
