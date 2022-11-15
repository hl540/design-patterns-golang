package factory_method

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun IGun) {
	fmt.Printf("Gun: %s\n", gun.getName())
	fmt.Printf("Power: %d\n", gun.getPower())
}
