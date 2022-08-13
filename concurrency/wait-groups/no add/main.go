package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("go routine done")
	}()
	wg.Wait()
	fmt.Println("execute immediately")
}

//here we can see what happiness when we don't add nothing to the wait group and only telling done
//the function in the go routine will never be executed
