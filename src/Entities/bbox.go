package Entities

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type IBBox interface {
	GetSignedDistance(position mgl32.Vec3) float32
}

type BBox struct {
	position mgl32.Vec3
	bounds mgl32.Vec3
}

func (bbox *BBox) GetSignedDistance(pos mgl32.Vec3) float32 {
	position := pos.Sub(bbox.position)

	q := mgl32.Vec3{
		float32(math.Abs(float64(position[0]))),
		float32(math.Abs(float64(position[1]))),
		float32(math.Abs(float64(position[2]))),
	}.Sub(bbox.bounds)

	return mgl32.Vec3{
		float32(math.Max(float64(q[0]), 0)),
		float32(math.Max(float64(q[1]), 0)),
		float32(math.Max(float64(q[2]), 0)),
	}.Len() + float32(math.Min(
		math.Max(
			float64(q.X()),
			math.Max(
				float64(q.Y()),
				float64(q.Z()),
			),
		),
		0,
	))
}
