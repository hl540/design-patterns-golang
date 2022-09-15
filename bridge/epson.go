package bridge

import "fmt"

// Epson 爱普生打印机
type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
