package iterator

type Collection interface {
	createCollection() Iterator
}
