package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()
	wg.Done()
}

//here we are calling to many times to done and that's a problem because our wait group can handle only the amount of Add
//or zero amount of waiting
