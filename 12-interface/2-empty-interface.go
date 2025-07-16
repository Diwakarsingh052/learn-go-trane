package main

import "fmt"

func main() {
	// empty interface can store value of any type
	// any is an alias of interface{}
	var i any = 10
	//var i interface{} = 10
	a, ok := i.(int) // type assertion, checking if any storing int kind of value
	if ok {
		fmt.Println(a)
	}
	fmt.Printf("%T\n", i)
	i = "hello"
	fmt.Printf("%T\n", i)

}
