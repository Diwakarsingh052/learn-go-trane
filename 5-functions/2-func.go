package main

import (
	"fmt"
	"strconv"
)

// datatype of func -> func(args)returnType
func main() {

	fmt.Println(strconv.Atoi("10"))
	operate(add(), 10, 20)
	operate(sub, 10, 20)
}

func addV2(a, b int) {
	fmt.Println(a + b)
}

func add() func(int, int) {
	i := func(a, b int) {
		fmt.Println(a + b)
	}
	return i
}
func sub(a, b int) {
	fmt.Println(a - b)
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op func(int, int), x, y int) {
	op(x, y)
}
