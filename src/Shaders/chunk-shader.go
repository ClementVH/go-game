package Shaders

import (
	"go-game/src/Entities"
	"go-game/src/State"

	"github.com/go-gl/mathgl/mgl32"
)

type ChunkShader struct {
	transformationMatrix  int32
	projectionMatrix      int32
	viewMatrix            int32
	lightPosition         int32
	lightColor            int32
	combatChunk           int32
	chunkTexture          int32
	blendMapTexture       int32
	startPositionsTexture int32
	ShaderProgram
}

func NewChunkShader() *ChunkShader {
	shaderProgram := NewShaderProgram(ChunkVertexShader, ChunkFragmentShader)
	var shader ChunkShader
	shader.ShaderProgram = shaderProgram
	shader.bindAttributes()
	shader.setup()
	shader.getAllUniformLocations()
	return &shader
}

func (shader *ChunkShader) bindAttributes() {
	shader.bindAttribute(0, "position")
	shader.bindAttribute(1, "textureCoords")
}

func (shader *ChunkShader) getAllUniformLocations() {
	shader.transformationMatrix = shader.getUniformLocation("transformationMatrix")
	shader.projectionMatrix = shader.getUniformLocation("projectionMatrix")
	shader.viewMatrix = shader.getUniformLocation("viewMatrix")
	shader.lightPosition = shader.getUniformLocation("lightPosition")
	shader.lightColor = shader.getUniformLocation("lightColor")
	shader.combatChunk = shader.getUniformLocation("combatChunk")
	shader.chunkTexture = shader.getUniformLocation("chunkTexture")
	shader.blendMapTexture = shader.getUniformLocation("blendMapTexture")
	shader.startPositionsTexture = shader.getUniformLocation("startPositionsTexture")
}

func (shader *ChunkShader) LoadTransformationMatrix(transformation mgl32.Mat4) {
	loadMatrix(shader.transformationMatrix, transformation)
}

func (shader *ChunkShader) LoadProjectionMatrix(projection mgl32.Mat4) {
	loadMatrix(shader.projectionMatrix, projection)
}

func (shader *ChunkShader) LoadViewMatrix(camera *Entities.Camera) {
	matrix := camera.GetViewMatrix()
	loadMatrix(shader.viewMatrix, matrix)
}

func (shader *ChunkShader) LoadLight(light *Entities.Light) {
	loadVector(shader.lightPosition, light.Position)
	loadVector(shader.lightColor, light.Color)
}

func (shader *ChunkShader) LoadTextures() {
	loadInt(shader.chunkTexture, 0)
	loadInt(shader.blendMapTexture, 1)
	loadInt(shader.startPositionsTexture, 2)
}

func (shader *ChunkShader) LoadCombatChunk() {
	if State.Combat.Combat != nil && State.Combat.Combat.GetChunk() != nil {
		loadVector(shader.combatChunk, State.Combat.Combat.GetChunk().Position)
	} else {
		loadVector(shader.combatChunk, mgl32.Vec3{-1, 0, 1})
	}
}

var ChunkVertexShader = `
#version 330

in vec3 position;
in vec2 textureCoords;
in vec3 normal;

out vec2 pass_textureCoords;
out vec3 surfaceNormal;
out vec3 toLightVector;
out vec4 worldPosition;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;
uniform mat4 viewMatrix;
uniform vec3 lightPosition;

void main(void) {
	worldPosition = transformationMatrix * vec4(position, 1.0);
	gl_Position = projectionMatrix * viewMatrix * worldPosition;
	pass_textureCoords = textureCoords;

	surfaceNormal = (transformationMatrix * vec4(normal, 0.0)).xyz;
	toLightVector = lightPosition - worldPosition.xyz;
}
` + "\x00"

var ChunkFragmentShader = `
#version 330

in vec2 pass_textureCoords;
in vec3 surfaceNormal;
in vec3 toLightVector;
in vec4 worldPosition;

out vec4 out_Color;

uniform sampler2D chunkTexture;
uniform sampler2D blendMapTexture;
uniform sampler2D startPositionsTexture;
uniform vec3 lightColor;
uniform vec3 combatChunk;

void main(void) {
	vec3 unitNormal = normalize(surfaceNormal);
	vec3 unitLightVector = normalize(toLightVector);

	float nDot1 = dot(unitNormal, unitLightVector);
	float brightness = max(nDot1, 0.4);
	vec3 diffuse = brightness * lightColor;

	bool inCombatChunkX = worldPosition.x >= combatChunk.x && worldPosition.x <= combatChunk.x + 16;
	bool inCombatChunkZ = worldPosition.z >= combatChunk.z && worldPosition.z <= combatChunk.z + 16;
	bool inCombatChunk = inCombatChunkX && inCombatChunkZ;

	vec4 textureColor = texture(chunkTexture, pass_textureCoords);

	if (inCombatChunk) {
		vec4 gridColor = vec4(0.7, 0.7, 0.7, 1.0);
		vec4 blueColor = vec4(0.0, 0.0, 1.0, 1.0);
		vec4 redColor = vec4(1.0, 0.0, 0.0, 1.0);

		vec4 blendMapColor = texture(blendMapTexture, pass_textureCoords);
		vec4 startPositionColor = texture(startPositionsTexture, pass_textureCoords);

		if (blendMapColor.r > 0.5) {
			textureColor = gridColor;
		}

		if (startPositionColor.r > 0.5) {
			textureColor = blueColor;
		} else if (startPositionColor.g > 0.5) {
			textureColor = redColor;
		}
	}

	out_Color = vec4(diffuse, 1.0) * textureColor;
}
` + "\x00"
