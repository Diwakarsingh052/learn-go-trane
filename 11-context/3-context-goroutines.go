package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()

	go func() {
		x := slowFunc()
		ch <- x
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case x := <-ch:
		fmt.Println("received value from slow function", x)
	}

	fmt.Println("main function is done")
}

func slowFunc() int {
	time.Sleep(time.Second * 1)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
