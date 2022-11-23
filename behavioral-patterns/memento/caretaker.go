package memento

type Caretaker struct {
	memento []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.memento = append(c.memento, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.memento[index]
}
