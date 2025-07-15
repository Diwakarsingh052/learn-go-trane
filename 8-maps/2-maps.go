package main

import (
	"fmt"
	"runtime"
)

type user struct {
	name string
	age  int
}

type users map[int]user

func main() {
	runtime.GOMAXPROCS(1)
	u := users{1: user{"john", 20}, 2: user{"jane", 21}}

	usr, ok := u[1] // ok would be true if user is found
	if !ok {        // ok == false
		fmt.Println("user not found")
		return
	}
	fmt.Println(usr)
}
