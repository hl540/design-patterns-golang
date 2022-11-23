package state

import "fmt"

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	notItem       State

	currentState State

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount int, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	v.hasItem = &HasItemState{
		vendingMachine: v,
	}
	v.itemRequested = &ItemRequestedState{
		vendingMachine: v,
	}
	v.hasMoney = &HasMoneyState{
		vendingMachine: v,
	}
	v.notItem = &NoItemState{
		vendingMachine: v,
	}
	v.setState(v.hasItem)
	return v
}

func (m *VendingMachine) addItem(i int) error {
	return m.currentState.addItem(i)
}

func (m *VendingMachine) requestItem() error {
	return m.currentState.requestItem()
}

func (m *VendingMachine) insertMoney(money int) error {
	return m.currentState.insertMoney(money)
}

func (m *VendingMachine) dispenseItem() error {
	return m.currentState.dispenseItem()
}

func (m *VendingMachine) setState(state State) {
	m.currentState = state
}

func (m *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	m.itemCount = m.itemCount + count
}
