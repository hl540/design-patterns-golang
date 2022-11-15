package factory_method

// Ak47 武器接口的Ak47实现
type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun{
			name:  "AK47",
			power: 4,
		},
	}
}
