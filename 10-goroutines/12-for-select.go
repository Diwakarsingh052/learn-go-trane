package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wgWorker := &sync.WaitGroup{}
	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()
		time.Sleep(3 * time.Second)
		get <- "get"
		get <- "get 2"
	}()

	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()
		time.Sleep(1 * time.Second)
		post <- "post"
	}()

	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()
		put <- "put"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
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
			case <-done:
				fmt.Println("all goroutines are done")
				return

			}
		}
	}()

	wg.Wait()
}
