package adapter

import "fmt"

// WindowsAdapter 适配器
type WindowsAdapter struct {
	windowsMacClient *Windows
}

// InsertIntoLightningPort 实现Computer接口
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("适配器将 Lightning 信号转换为 USB。")
	w.windowsMacClient.insertIntoUSBPort()
}
