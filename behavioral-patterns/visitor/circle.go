package visitor

type Circle struct {
	radius int
}

func (c *Circle) getType() string {
	return "Circle"
}

func (c *Circle) accept(visitor Visitor) {
	visitor.visitForCircle(c)
}
