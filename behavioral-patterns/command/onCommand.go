package command

type OnCommand struct {
	device Device
}

func (o *OnCommand) execute() {
	o.device.on()
}
