package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single1 struct {
	
}

var singleInstance1 *single1

func getInstance1() *single1 {
	if singleInstance1 == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance1 == nil {
			fmt.Println("Creating single instance now.")
			singleInstance1 = &single1{}
		} else {
			fmt.Println("single instance already created.")
		}
	} else {
		fmt.Println("single instance already created.")
	}
	return singleInstance1
}