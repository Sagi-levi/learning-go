package main

import (
	"fmt"
	"sync"
)

type request func()

func main() {
	requests := map[int]request{}
	for i := 1; i <= 100; i++ {
		f := func(n int) request {
			return func() {
				fmt.Println("request", n)
			}
		}
		requests[i] = f(i)
	}
	var wg sync.WaitGroup
	max := 10
	for i := 0; i < len(requests); i += max {
		for j := i; j < i+max; j++ {
			wg.Add(1)
			go func(r request) {
				defer wg.Done()
				r()
			}(requests[j+1])
		}
		wg.Wait()
		fmt.Println(max, "requests processed")
	}
}

//way to limit request simultaneously to 10 by wait group
//since we wait for requests to give as a "Done" and we are use "Add"(1)*10 and "Wait" between them, we assure
//nest iteration will happen only after 10 requests
