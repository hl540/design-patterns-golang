package mediator

import "testing"

func Test1(t *testing.T) {
	stationManager := newStationManager()
	t1 := &PassengerTrain{mediator: stationManager, name: "列车1"}
	t2 := &PassengerTrain{mediator: stationManager, name: "列车2"}
	t3 := &PassengerTrain{mediator: stationManager, name: "列车3"}

	t1.arrive()
	t2.arrive()
	t1.depart()
	t3.arrive()
	t2.depart()
	t3.depart()

}
