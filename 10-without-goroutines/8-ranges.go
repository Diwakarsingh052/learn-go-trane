package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int, 5)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		ch <- 10
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
