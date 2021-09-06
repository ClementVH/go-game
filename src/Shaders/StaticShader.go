package Shaders

func CreateStaticShader() ShaderProgram {
	var shader = CreateShaderProgram()
	bindAttribute(shader, 0, "position")
	return shader
}
