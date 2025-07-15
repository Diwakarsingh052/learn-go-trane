package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// context -> cancellation signals
	//     		-> passing values in different layers of a http request

	// empty container
	// we need a container to hold the context object
	// Background returns a non-nil, empty Context.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	// clean up the resources taken up by the context
	// if we don't put cancel in defer timer would be immediately cancelled
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
		// done would receive signal in case the timer is over
		fmt.Println("context done")
		return 0, ctx.Err()
	default:
		// select would unblock immediately if done has not received the value
		fmt.Println("time is still there")
	}
	return i, nil
	//sql.DB.ExecContext()

}
