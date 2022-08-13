package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	work(wg)
	wg.Wait()
	fmt.Println("finish")
}
func work(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("work")
}

//you can't pass wait group ass a value!!!!
//since you want to control a value of existing int you want access to the source and not to a copy of it
