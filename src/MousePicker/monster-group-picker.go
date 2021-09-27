package MousePicker

import (
	"errors"
	"go-game/src/Entities"
	"go-game/src/Systems"

	"github.com/go-gl/mathgl/mgl32"
)

var MonsterGroupPicker MousePicker

func InitMonsterGroupPicker(camera *Entities.Camera, projectionMatrix mgl32.Mat4) {
	getPicker(camera, projectionMatrix)
}

func (picker *MousePicker) GetMonsterGroup() (interface{}, error) {
	wildMonsterSystem := Systems.Systems["WILD_MONSTER_SYSTEM"].(*Systems.WildMonsterSystem)

	var startPos = picker.RayOrigin
	var currPos = startPos
	var distance float32 = 10000
	var res interface{}
	var currIt = 50

	groups := wildMonsterSystem.GetGroups()

	for distance >= 0.01 && currIt > 0 {
		currIt--
		for _, group := range groups {
			for _, monster := range group {
				newDistance := monster.GetSignedDistance(currPos)
				if newDistance < distance {
					distance = newDistance
					res = group
				}
			}
		}
		currPos = currPos.Add(picker.Ray.Mul(distance))
	}

	if distance < 1 {
		return res, nil
	}

	return nil, errors.New("no entity found")
}
