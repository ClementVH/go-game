package Shaders

var FragmentShader = `
#version 330

in vec2 pass_textureCoords;

out vec4 out_Color;

uniform sampler2D textureSampler;

void main(void) {
	out_Color = texture(textureSampler, pass_textureCoords);
}
` + "\x00"