package ToolBox

import (
	"go-game/src/Entities"

	"github.com/go-gl/mathgl/mgl32"
)

func CreateTransformationMatrix(translation mgl32.Vec3, rx, ry, rz, scale float32) mgl32.Mat4 {
	var matrix = mgl32.Ident4()
	matrix = matrix.Mul4(mgl32.Translate3D(translation.X(), translation.Y(), translation.Z()))
	matrix = matrix.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(rx), mgl32.Vec3{1, 0, 0}))
	matrix = matrix.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(ry), mgl32.Vec3{0, 1, 0}))
	matrix = matrix.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(rz), mgl32.Vec3{0, 0, 1}))
	matrix = matrix.Mul4(mgl32.Scale3D(scale, scale, scale))
	return matrix
}

func CreateViewMatrix(camera Entities.Camera) mgl32.Mat4 {
	var matrix = mgl32.Ident4()
	matrix = matrix.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(camera.Pitch), mgl32.Vec3{1, 0, 0}))
	matrix = matrix.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(camera.Yaw), mgl32.Vec3{0, 1, 0}))

	var cameraPos = camera.Position
	matrix = matrix.Mul4(mgl32.Translate3D(-cameraPos.X(), -cameraPos.Y(), -cameraPos.Z()))
	return matrix
}
