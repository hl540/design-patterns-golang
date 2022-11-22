package mediator

type Train interface {
	// 列车进入站台
	arrive()
	// 列车离开站台
	depart()
	// 允许列车进入站台
	permitArrive()
}
