package observer

import "testing"

func Test1(t *testing.T) {
	shirtItem := newItem("Nike Shirt")

	observer1 := &Customer{id: "abc@gmail.com"}
	observer2 := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observer1)
	shirtItem.register(observer2)

	shirtItem.updateAvailability()
}
