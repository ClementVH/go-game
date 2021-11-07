package State

import "go-game/src/Entities"

func GetPlayer() *Entities.Player {
	return Systems.PlayerSystem.GetPlayer()
}
