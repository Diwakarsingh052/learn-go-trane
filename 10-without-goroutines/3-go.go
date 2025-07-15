package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 1; i <= 5; i++ {
		// adding one to the counter in each iteration
		wg.Add(1)
		go work(i, wg)
	}
	// waiting for goroutines
	wg.Wait()

}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("work is going on", id)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("sub work is going on", id)
	}()
}
