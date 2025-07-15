package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	for i := 1; i <= 5; i++ {

		x := <-ch
		fmt.Println(x)

	}
	wg.Wait()

}
