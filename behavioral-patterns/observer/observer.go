package observer

type Observer interface {
	update(string string)
	getId() string
}
