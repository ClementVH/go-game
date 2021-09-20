package Systems

import "go-game/src/Entities"

type ISystem interface {
	Tick()
	GetEntities() []Entities.IEntity
}

type System struct {
}

func NewSystem() *System {
	return &System{}
}

var Systems map[string]ISystem
