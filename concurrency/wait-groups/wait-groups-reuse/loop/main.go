package main

import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(3)
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			wg.Wait()
		}()
	}
}

//since the for loop create 300 adds and 300 dones but not in the same order then the wait between causing an error
// if we are waiting for all the 300 to finish or even waiting for each 3 to finish before starting another 3 it can work
