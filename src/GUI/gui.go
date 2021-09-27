package GUI

type GUI struct {
	MonsterGroupHover
}

func NewGUI() *GUI {
	return &GUI{
		newMonsterGroupHover(),
	}
}

func (gui *GUI) Update() {

}
