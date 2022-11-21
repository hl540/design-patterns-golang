package chain_of_responsibility

import "testing"

func Test1(t *testing.T) {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)
	
	patient := &Patient{name: "abc"}
	reception.execute(patient)
}
