package chain_of_responsibility

type Department interface {
	execute(patient *Patient)
	setNext(department Department)
}
