package abstract_factory

// Adidas 实现ISportsFactory的具体工厂
type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
