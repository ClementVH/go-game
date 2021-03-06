package Shaders

import (
	"go-game/src/Entities"

	"github.com/go-gl/mathgl/mgl32"
)

type StaticShader struct {
	transformationMatrix int32
	projectionMatrix     int32
	viewMatrix           int32
	lightPosition        int32
	lightColor           int32
	ShaderProgram
}

func NewStaticShader() *StaticShader {
	shaderProgram := NewShaderProgram(VertexShader, FragmentShader)
	shader := StaticShader{0, 0, 0, 0, 0, shaderProgram}
	shader.bindAttributes()
	shader.setup()
	shader.getAllUniformLocations()
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
	matrix := camera.GetViewMatrix()
	loadMatrix(shader.viewMatrix, matrix)
}

func (shader *StaticShader) LoadLight(light *Entities.Light) {
	loadVector(shader.lightPosition, light.Position)
	loadVector(shader.lightColor, light.Color)
}

var VertexShader = `
#version 330

in vec3 position;
in vec2 textureCoords;
in vec3 normal;

out vec2 pass_textureCoords;
out vec3 surfaceNormal;
out vec3 toLightVector;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;
uniform mat4 viewMatrix;
uniform vec3 lightPosition;

void main(void) {
	vec4 worldPosition = transformationMatrix * vec4(position, 1.0);
	gl_Position = projectionMatrix * viewMatrix * worldPosition;
	pass_textureCoords = textureCoords;

	surfaceNormal = (transformationMatrix * vec4(normal, 0.0)).xyz;
	toLightVector = lightPosition - worldPosition.xyz;
}
` + "\x00"

var FragmentShader = `
#version 330

in vec2 pass_textureCoords;
in vec3 surfaceNormal;
in vec3 toLightVector;

out vec4 out_Color;

uniform sampler2D textureSampler;
uniform vec3 lightColor;

void main(void) {
	vec3 unitNormal = normalize(surfaceNormal);
	vec3 unitLightVector = normalize(toLightVector);

	float nDot1 = dot(unitNormal, unitLightVector);
	float brightness = max(nDot1, 0.4);
	vec3 diffuse = brightness * lightColor;

	out_Color = vec4(diffuse, 1.0) * texture(textureSampler, pass_textureCoords);
}
` + "\x00"
