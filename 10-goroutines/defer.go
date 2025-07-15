package main

import (
	"fmt"
)

func main() {
	// defer statement runs when the function is returning
	// defer statements gives guarantee of exec , once registered
	// defer maintains the stack
	// first in last out
	defer fmt.Println(1)
	defer fmt.Println(2)
	panic("")
	fmt.Println(3)
}
