package visitor

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (m *MiddleCoordinates) visitForSquare(square *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (m *MiddleCoordinates) visitForCircle(circle *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (m *MiddleCoordinates) visitForRectangle(rectangle *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
