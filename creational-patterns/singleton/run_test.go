package singleton

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstance1()
	}
	fmt.Scanln()
}

func Test2(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
