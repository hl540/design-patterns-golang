package command

import "testing"

func Test1(t *testing.T) {
	tv := &Tv{}

	onCommand := &OnCommand{device: tv}
	onButton := &Button{command: onCommand}
	onButton.press()

	offCommand := &OffCommand{device: tv}
	offButton := &Button{command: offCommand}
	offButton.press()
}
