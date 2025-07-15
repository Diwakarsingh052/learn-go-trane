package main

import "fmt"

// Function that can only send to the channel
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	
	close(ch)
}

// Function that can only receive from the channel
func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
