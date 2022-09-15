package adapter

import "testing"

func Test1(t *testing.T) {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windows := &Windows{}
	windowsMacAdapter := &WindowsAdapter{windowsMacClient: windows}
	client.InsertLightningConnectorIntoComputer(windowsMacAdapter)
}
