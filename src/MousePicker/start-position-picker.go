package MousePicker

import (
	"errors"
	"fmt"
	"go-game/src/State"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

func (picker *MousePicker) GetStartPosition() (mgl32.Vec2, error) {
	var startPos = picker.RayOrigin
	var currPos = startPos
	var distance float32 = 10000
	var res mgl32.Vec2
	var currIt = 50

	chunk := State.Combat.Combat.GetChunk()
	positions := chunk.StartPositions.Teams[0]

	for distance >= 0.01 && currIt > 0 {
		currIt--
		for _, position := range positions {
			// project currPos onto plane
			projectedCurrPos := mgl32.Vec2{currPos.X(), currPos.Z()}
			// get distance to position
			projectedDistance := sdAxisAlignedRect(
				projectedCurrPos,
				mgl32.Vec2{chunk.Position.X() + position.X(), chunk.Position.Z() + position.Y()},
				mgl32.Vec2{chunk.Position.X() + position.X() + 1, chunk.Position.Z() + position.Y() + 1},
			)
			// distance is sqrt ( z² + distance ²)
			newDistance := float32(math.Sqrt(math.Pow(float64(currPos.Y()), 2) + math.Pow(float64(projectedDistance), 2)))
			fmt.Println(projectedCurrPos, startPos, currPos, newDistance, distance)
			if newDistance < distance {
				fmt.Println("#############################")
				distance = newDistance
				res = position
			}
		}
		currPos = currPos.Add(picker.Ray.Mul(distance))
	}

	if distance < 1 {
		return res, nil
	}

	return mgl32.Vec2{}, errors.New("no entity found")
}

func sdAxisAlignedRect(uv, tl, br mgl32.Vec2) float32 {
	d := maxVec2(tl.Sub(uv), uv.Sub(br))
	return (maxVec2(mgl32.Vec2{0, 0}, d)).Len() + float32(math.Min(0.0, math.Max(float64(d.X()), float64(d.Y()))))
}

func maxVec2(a, b mgl32.Vec2) mgl32.Vec2 {
	return mgl32.Vec2{
		float32(math.Max(float64(a.X()), float64(b.X()))),
		float32(math.Max(float64(a.Y()), float64(b.Y()))),
	}
}

// float sdAxisAlignedRect(vec2 uv, vec2 tl, vec2 br)
// {
//     vec2 d = max(tl-uv, uv-br);
//     return length(max(vec2(0.0), d)) + min(0.0, max(d.x, d.y));
// }
