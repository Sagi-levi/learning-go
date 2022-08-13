package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go task1(&wg)
	wg.Wait()

	wg.Add(2)
	go task2(&wg)
	go task3(&wg)
	wg.Wait()

	wg.Add(1)
	go task4(&wg)
	wg.Wait()

	wg.Add(1)
	go task5(&wg)
	wg.Wait()

}
func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("task1")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("task2")
}
func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("task3")
}
func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("task4")
}
func task5(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("task5")
}

//since we're only adding 1 to the wg then telling it we are done and waiting before starting another then the orders stays
