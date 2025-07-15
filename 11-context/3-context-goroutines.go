package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := slowFunc()
		ch <- x
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return

		case x := <-ch:
			fmt.Println("received value from slow function", x)
		}
	}()

	fmt.Println("main: done")
	wg.Wait()

}

func slowFunc() int {
	time.Sleep(time.Second * 1)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
