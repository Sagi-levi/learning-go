package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Go Scheduler will create 8 threads
	// which in theory will be able to execute
	// 8 blocking go routines in parallel
	// totalling ~100ms
	now := time.Now()
	fmt.Println("number of cores:", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go work(i + 1)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main is done ", time.Since(now))
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task", id, "is done")
}

//what really happened is that the application starts running and creating 8 different go routines but every one of
//them takes 100ms ms and the last thing in our main is ending in 100ms as well, so duo that no one is waiting enough time for
//the go routines will finish then not all of them print what they should

//here we can understand why we need wait groups
// since we want to make sure all tasks are done and control what is happening before what in our app, we want to use wait groups
