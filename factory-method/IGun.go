package factory_method

// IGun 武器抽象接口
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
