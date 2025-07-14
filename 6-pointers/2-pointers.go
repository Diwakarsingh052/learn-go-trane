package main

import "fmt"

var x int = 10

func main() {
	// address of p is x80 = and the value it is storing is nil
	var p *int // nil // default value of a pointer is nil

	// after calling the update value p is still nil, as we never updated the pointer
	updateValue(p)
	fmt.Println(*p) // panic // nil pointer dereference

}

func updateValue(p *int) {
	p = &x
	*p = 100

}
