package abstract_factory

import "fmt"

// ISportsFactory 抽象工厂接口
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

// GetShortsFactory 根据brand返回不太的工厂实现
func GetShortsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
