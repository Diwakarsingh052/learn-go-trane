package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {

	// empty container
	// we need a container to hold the context object
	// Background returns a non-nil, empty Context.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	fmt.Println(SlowFunc(ctx, "10"))

}

func SlowFunc(ctx context.Context, input string) (int, error) {
	i, err := strconv.Atoi(input)
	time.Sleep(time.Millisecond * 2)
	if err != nil {
		return 0, err
	}
	select {
	case <-ctx.Done():
		fmt.Println("context done")
		return 0, ctx.Err()
	default:
		fmt.Println("time is still there")
	}
	return i, nil
	//sql.DB.ExecContext()

}
