package main

import (
	"fmt"
	"sync"
)

func main() {
	// don't use this pattern with buffered channels
	// if you want to use a buffered channel, then use range to receive the remaining value
	// select is used when we want to listen or send values to over a multiple channel
	get := make(chan string, 1)
	post := make(chan string, 1)
	put := make(chan string, 1)
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wgWorker := &sync.WaitGroup{}
	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()
		//time.Sleep(3 * time.Second)
		get <- "get"
		get <- "get 2"
	}()

	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()
		//time.Sleep(1 * time.Second)
		post <- "post"
	}()

	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()
		put <- "put"
	}()

	wg.Add(1)
	go func() {
		// tracking when all the work is finished,
		// once work is finished, we will send signal to done chan
		defer wg.Done()
		wgWorker.Wait()

		// close is a send signal
		close(done)
	}()

	//fmt.Println(<-get)
	//
	//fmt.Println(<-post)
	//fmt.Println(<-put)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			select {
			//whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default
			case x := <-get:
				fmt.Println(x)
			case x := <-post:
				fmt.Println(x)
			case x := <-put:
				fmt.Println(x)

				// we will receive a close signal when all
				// go routines finished sending
			case <-done:
				fmt.Println("all goroutines are done")
				return

			}
		}
	}()

	wg.Wait()
}
