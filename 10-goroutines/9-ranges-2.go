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
			//we need to block our goroutine before closing the channel
			//because we want to make sure all the work
			// is done and finished
			// wgWorker waitgroup we are using to track number of worker goroutines
			wgWorker.Add(1)
			go func() {
				defer wgWorker.Done()
				ch <- i
			}()
		}

		// waiting until all the workers are not finished
		wgWorker.Wait()
		close(ch)

	}()

	for v := range ch {
		fmt.Println(v)
	}
	wg.Wait()

}
