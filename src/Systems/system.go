package Systems

type ISystem interface {
	Tick()
}

type System struct {
}

func NewSystem() *System {
	return &System{}
}