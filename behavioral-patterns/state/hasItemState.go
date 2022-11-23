package state

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *HasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.notItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (h *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}
