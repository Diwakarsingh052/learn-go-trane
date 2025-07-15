package main

import (
	"fmt"
	"sync"
)

// A send on an unbuffered channel can proceed if a receiver is ready.
// A send on a buffered channel can proceed if there is room in the buffer.
// A send on a closed channel
func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int) // unbuffered channel has size of 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch // recv // this would block if no sender is present,
		fmt.Println(x)

		//and another goroutine from the queue would be picked up
		//which is sender goroutine in this case
		//communication completes, and we get 10 on the screen
	}()
	ch <- 10

	wg.Wait()
}
