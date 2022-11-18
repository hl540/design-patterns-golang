package flyweight

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (g *game) addTerrorist(dressType string) {
	player := NewPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *game) addCounterTerrorist(dressType string) {
	player := NewPlayer("CT", dressType)
	g.counterTerrorists = append(g.terrorists, player)
}
