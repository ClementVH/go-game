package MousePicker

import (
	"errors"
	"go-game/src/Entities"
	"go-game/src/State"
)

var MonsterGroupPicker MousePicker

func (picker *MousePicker) GetMonsterGroup() ([]*Entities.Monster, error) {
	wildMonsterSystem := State.Systems.WildMonsterSystem

	var startPos = picker.RayOrigin
	var currPos = startPos
	var distance float32 = 10000
	var res []*Entities.Monster
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
