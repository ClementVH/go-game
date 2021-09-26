package ToolBox

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type ICamera interface {
	GetViewMatrix() mgl32.Mat4
}

type Raycast struct {
	Ray mgl32.Vec3
	projectionMatrix mgl32.Mat4
	viewMatrix mgl32.Mat4
	camera ICamera
}

func NewRaycast(camera ICamera, projectionMatrix mgl32.Mat4) *Raycast {
	return &Raycast{
		mgl32.Vec3{},
		projectionMatrix,
		camera.GetViewMatrix(),
		camera,
	}
}

func (raycast *Raycast) Update() {
	raycast.viewMatrix = raycast.camera.GetViewMatrix()
	raycast.Ray = raycast.calculateMouseRay()
}

func (raycast *Raycast) calculateMouseRay() mgl32.Vec3 {
	mouseX, mouseY := glfw.GetCurrentContext().GetCursorPos()
	normalizedCoords := raycast.getNormalizedDeviceCoords(float32(mouseX), float32(mouseY))
	clipCoords := mgl32.Vec4{normalizedCoords.X(), normalizedCoords.Y(), -1, 1}
	eyeCoords := raycast.toEyeCoords(clipCoords)
	worldRay := raycast.toWorldCoords(eyeCoords)
	return worldRay
}

func (raycast *Raycast) getNormalizedDeviceCoords(mouseX, mouseY float32) mgl32.Vec2 {
	width, height := glfw.GetCurrentContext().GetSize()
	x := (2 * mouseX) / float32(width) - 1
	y := (2 * mouseY) / float32(height) - 1
	return mgl32.Vec2{x, -y}
}

func (raycast *Raycast) toEyeCoords(clipCoords mgl32.Vec4) mgl32.Vec4{
	invertedProjection := raycast.projectionMatrix.Inv()
	eyeCoords := invertedProjection.Mul4x1(clipCoords)
	return mgl32.Vec4{eyeCoords.X(), eyeCoords.Y(), -1, 0}
}

func (raycast *Raycast) toWorldCoords(eyeCoords mgl32.Vec4) mgl32.Vec3 {
	invertedView := raycast.viewMatrix.Inv()
	worldCoords := invertedView.Mul4x1(eyeCoords)
	mouseRay := mgl32.Vec3{worldCoords.X(), worldCoords.Y(), worldCoords.Z()}
	return mouseRay.Normalize()
}
