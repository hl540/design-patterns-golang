package visitor

type Square struct {
	side int
}

func (s *Square) getType() string {
	return "Square"
}

func (s *Square) accept(visitor Visitor) {
	visitor.visitForSquare(s)
}
