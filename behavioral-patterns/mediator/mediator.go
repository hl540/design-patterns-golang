package mediator

type Mediator interface {
	// 询问是否可以停靠
	canArrive(train Train) bool
	// 列车离开通知
	notifyAboutDeparture()
}
