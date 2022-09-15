package bridge

// Computer 计算机接口，抽象了打印功能
type Computer interface {
	Print()
	SetPrinter(printer Printer)
}
