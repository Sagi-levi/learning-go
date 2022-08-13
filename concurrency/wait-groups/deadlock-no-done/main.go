package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

//a good example of an app get stuck without DONE to decrease the wait group counter to zero
