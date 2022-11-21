package chain_of_responsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(patient *Patient) {
	if patient.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(department Department) {
	c.next = department
}
