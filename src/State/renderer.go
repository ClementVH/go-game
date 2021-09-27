package State

import "github.com/go-gl/mathgl/mgl32"

var Renderer RendererState = RendererState{}

type RendererState struct {
	ProjectionMatrix mgl32.Mat4
}

func (state *RendererState) SetProjectionMatrix(matrix mgl32.Mat4) {
	state.ProjectionMatrix = matrix
}
