package factory_method

// Musket 武器接口的musket实现
type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
