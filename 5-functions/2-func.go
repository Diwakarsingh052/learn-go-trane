package main

import (
	"fmt"
	"strconv"
)

type money int
type operation func(int, int)

// datatype of func -> func(args)returnType
func main() {

	//var m money = 100
	fmt.Println(strconv.Atoi("10"))
	operate(add(), 10, 20)
	operateV2(sub, 10, 20)
	operate(sub, 10, 20)
}

func addV2(a, b int) {
	fmt.Println(a + b)
}

func add() operation {
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

func operateV2(op operation, x, y int) {
	op(x, y)
}
