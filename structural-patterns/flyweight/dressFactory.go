package flyweight

import "errors"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (f *DressFactory) getDressByType(dressType string) (Dress, error) {
	if f.dressMap[dressType] != nil {
		return f.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		f.dressMap[dressType] = newTerroristDress()
		return f.dressMap[dressType], nil
	}

	if dressType == CounterTerroristDressType {
		f.dressMap[dressType] = newCounterTerroristDress()
		return f.dressMap[dressType], nil
	}
	return nil, errors.New("Wrong dress tytpe passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
