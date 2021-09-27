package MousePicker

import (
	"errors"
	"go-game/src/Entities"
	"go-game/src/Systems"

	"github.com/go-gl/mathgl/mgl32"
)

func InitEntityPicker(camera *Entities.Camera, projectionMatrix mgl32.Mat4) {
	getPicker(camera, projectionMatrix)
}

func (picker *MousePicker) GetEntity() (interface{}, error) {
	wildMonsterSystem := Systems.Systems["WILD_MONSTER_SYSTEM"]

	var startPos = picker.RayOrigin
	var currPos = startPos
	var distance float32 = 10000
	var m interface{}
	var currIt = 50

	entities := wildMonsterSystem.GetEntities()

	for distance >= 0.01 && currIt > 0 {
		currIt--
		for _, monster := range entities {
			newDistance := monster.GetSignedDistance(currPos)
			if newDistance < distance {
				distance = newDistance
				m = monster
			}
		}

		currPos = currPos.Add(picker.Ray.Mul(distance))
	}

	if distance < 1 {
		return m, nil
	}

	return nil, errors.New("no entity found")
}
