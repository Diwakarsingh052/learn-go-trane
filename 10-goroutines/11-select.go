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
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		get <- "get"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		post <- "post"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		put <- "put"
	}()

	//fmt.Println(<-get)
	//
	//fmt.Println(<-post)
	//fmt.Println(<-put)

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
	}
	wg.Wait()
}
