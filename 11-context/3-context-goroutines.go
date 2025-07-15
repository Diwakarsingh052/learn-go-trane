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
		select {
		case <-ctx.Done():
			fmt.Println("sender goroutine, doing cleanup", ctx.Err())
			return
		case ch <- x:
			fmt.Println("sent value to slow function")
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		// this goroutine would quit if we don't receive the value
		// in the time specified in context
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return

			// will receive the value if timer is still there
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
