package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			go func() {
				ch <- i
			}()
		}

	}()

	for v := range ch {
		fmt.Println(v)
	}
	wg.Wait()

}
