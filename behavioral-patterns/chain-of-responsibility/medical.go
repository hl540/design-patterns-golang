package chain_of_responsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) execute(patient *Patient) {
	if patient.medicineDone {
		fmt.Println("Medical already given to patient")
		m.next.execute(patient)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	patient.medicineDone = true
	m.next.execute(patient)
}

func (m *Medical) setNext(department Department) {
	m.next = department
}
