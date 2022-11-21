package chain_of_responsibility

type Patient struct {
	name             string
	registrationDone bool
	doctorCheckDone  bool
	medicineDone     bool
	paymentDone      bool
}
