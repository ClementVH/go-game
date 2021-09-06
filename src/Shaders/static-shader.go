package Shaders

func CreateStaticShader() ShaderProgram {
	var shader = CreateShaderProgram()
	bindAttribute(shader, 0, "position")
	bindAttribute(shader, 1, "textureCoords")
	return shader
}
