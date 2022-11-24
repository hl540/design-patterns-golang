package visitor

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	square := &Square{}
	circle := &Circle{}
	rectangle := &Rectangle{}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()

	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
