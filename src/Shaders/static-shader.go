package Shaders

type StaticShader struct {
	ShaderProgram
}

func NewStaticShader() StaticShader {
	shader := StaticShader{ShaderProgram{}}
	shader.ShaderProgram.IShaderProgram = &shader
	shader.create()
	return shader
}

func (shader *StaticShader) bindAttributes() {
	shader.bindAttribute(0, "position")
	shader.bindAttribute(1, "textureCoords")
}
