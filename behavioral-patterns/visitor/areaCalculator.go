package visitor

import "fmt"

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(square *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(circle *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitForRectangle(rectangle *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
