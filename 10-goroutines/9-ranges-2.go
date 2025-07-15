package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wgWorker := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			wgWorker.Add(1)
			go func() {
				defer wgWorker.Done()
				ch <- i
			}()
		}

		wgWorker.Wait()
		close(ch)

	}()

	for v := range ch {
		fmt.Println(v)
	}
	wg.Wait()

}
