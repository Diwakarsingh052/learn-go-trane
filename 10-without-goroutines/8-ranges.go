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
		close(ch) // sends a signal to stop the range
		// close: signal range that no more values be sent and it can stop after receiving remaining values

		// once the channel is closed, we can't send more values to it
		//ch <- 10
	}()

	// it would run infinitely, channel needs to be closed to stop this range
	for v := range ch {
		fmt.Println(v)
	}
}
