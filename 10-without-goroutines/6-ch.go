package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		ch <- 10
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch
		fmt.Println(x)
	}()

	wg.Wait()
}
