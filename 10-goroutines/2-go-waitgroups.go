package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//func(a, b int) int { // a func without name
	//anonymous function
	//	return a + b
	//}(10, 20) // () is calling the func
	wg := new(sync.WaitGroup)
	//wg := &sync.WaitGroup{}

	// waitgroup counter represents number of goroutine we are running
	wg.Add(1) // add 1 to the counter
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("hello world")
		wg.Done() // decrement the counter by one
	}()
	fmt.Println("doing more things in main")
	wg.Wait() // it will block until counter is not 0

	fmt.Println("end of main")
}
