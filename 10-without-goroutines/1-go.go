package main

import (
	"fmt"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(1) // set number of processors to use
	go hello() // this call would be concurrent,
	// which means hello is a new goroutine
	time.Sleep(time.Second * 2)
}

func hello() {
	//time.Sleep(time.Second * 3)
	fmt.Println("Hello World")
}
