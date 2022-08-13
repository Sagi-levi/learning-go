package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

//the second time we add 1 to the wait group causing the wait to throw an error
//you cant call a wait if you will never get done to decrease the add number to 0
