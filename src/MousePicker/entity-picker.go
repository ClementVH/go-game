package MousePicker

import (
	"errors"
	"go-game/src/Entities"
	"go-game/src/State"
)

func (picker *MousePicker) GetEntity() (*Entities.Entity, error) {
	wildMonsterSystem := State.Systems.WildMonsterSystem

	var startPos = picker.RayOrigin
	var currPos = startPos
	var distance float32 = 10000
	var m *Entities.Entity
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
