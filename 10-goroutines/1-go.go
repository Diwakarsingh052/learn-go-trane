package main

import "fmt"

func main() {
	// runtime.GOMAXPROCS(1) // set number of processors to use
	go hello()
}

func hello() {
	//time.Sleep(time.Second * 2)
	fmt.Println("Hello World")
}
