package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("3")
	}()
	wg.Wait()
}

//here you can see another way to make sure what you wanted to be happened, before start running go routines
//we are preparing the app that it should wait for some sort of call 3 times before she is ending
//WaitGroup we are adding the number 3 wg.add and on the other side(the fo routine side we call wg.done)
//at the end of the main or wherever we want to make sure things happened we put wg.wait

//things to know to understand:
//defer is happening at the end of the function call he belongs to -->https://go.dev/tour/flowcontrol/12
//goroutine is a lightweight thread of execution--->https://www.geeksforgeeks.org/golang-goroutine-vs-thread/
