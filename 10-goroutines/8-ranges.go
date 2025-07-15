package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int)
	ch <- 10
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i

		}
		fmt.Println("send finished")
		close(ch) // sends a signal to stop the range
		// close: signal range that no more values be sent and it can stop after receiving remaining values

		// once the channel is closed, we can't send more values to it
		//ch <- 10
	}()

	// it would run infinitely, channel needs to be closed to stop this range
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			time.Sleep(time.Second * 1)
			fmt.Println("range 1", v)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			time.Sleep(time.Second * 1)
			fmt.Println("range 2", v)
		}
	}()

	wg.Wait()

}
