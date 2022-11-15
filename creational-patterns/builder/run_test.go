package builder

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := NewDirector(normalBuilder)
	normalHouse := director.buildHouse()
	printHouseDetails(normalHouse)

	fmt.Println()

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	printHouseDetails(iglooHouse)
}

func printHouseDetails(h House) {
	fmt.Printf("House Door TYpe: %s\n", h.doorType)
	fmt.Printf("House Window TYpe: %s\n", h.windowType)
	fmt.Printf("House Num Floor: %d\n", h.floor)

}
