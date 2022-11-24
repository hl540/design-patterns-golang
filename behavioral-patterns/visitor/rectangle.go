package visitor

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

func (r *Rectangle) accept(visitor Visitor) {
	visitor.visitForRectangle(r)
}
