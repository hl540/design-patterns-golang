package observer

import "fmt"

type Item struct {
	observers map[string]Observer
	name      string
	inStock   bool
}

func newItem(name string) *Item {
	return &Item{
		name:      name,
		observers: make(map[string]Observer),
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(observer Observer) {
	i.observers[observer.getId()] = observer
}

func (i *Item) deregister(observer Observer) {
	delete(i.observers, observer.getId())
}

func (i *Item) notifyAll() {
	for _, observer := range i.observers {
		observer.update(i.name)
	}
}
