package mediator

import "fmt"

type PassengerTrain struct {
	mediator Mediator
	name     string
}

func (p *PassengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Printf("%s:到达站台, 等待其他列出离开\n", p.name)
		return
	}
	fmt.Printf("%s进入站台\n", p.name)
}

func (p *PassengerTrain) depart() {
	fmt.Printf("%s离开站台\n", p.name)
	p.mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) permitArrive() {
	fmt.Printf("允许%s进入站台\n", p.name)
	p.arrive()
}
