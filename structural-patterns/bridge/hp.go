package bridge

import "fmt"

// Hp 惠普打印机
type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
