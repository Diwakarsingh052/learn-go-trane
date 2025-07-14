package main

import "learn-go/slice"

func main() {
	//If the capacity of s is not large enough to fit the additional values,
	//append allocates a new, sufficiently large underlying array that
	//fits both the existing slice elements and the additional values.
	//Otherwise, append re-uses the underlying array.
	a := []int{10, 20, 30, 40, 50}
	//a := make([]int, 0, 10)
	slice.Inspect("a", a)

	a = append(a, 60, 70)
	slice.Inspect("a", a)
}
