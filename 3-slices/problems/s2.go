package main

import "learn-go/slice"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	b := x[1:5] // 20,30,40,50

	slice.Inspect("x", x)
	slice.Inspect("b", b)
	//below line would change x, b refers to the same backing array,
	//there is room to add two more elements to the existing backing array
	//so it would add the value to the backing array refer  by x

	b = append(b, 777)
	//b[0] = 1000
	//b = append(b, 777, 888, 999) // this would create a new backing array for b,
	//it will not change the x slice

	slice.Inspect("b", b)
	slice.Inspect("x", x)

}
