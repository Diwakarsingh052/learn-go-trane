package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		b := abc()
		fmt.Println(b)
	}()

	wg.Wait()
}

func abc() int {
	return 10
}
