package state

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemState) addItem(count int) error {
	n.vendingMachine.incrementItemCount(count)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

func (n *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (n *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (n *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
