package main

import "fmt"

func main() {
	x := []int{10, 20, 30}
	x = updateSlice(x)
	fmt.Println(x)
}

func updateSlice(x []int) []int {
	// if you want to update the value ,
	//then passing the reference of existing backing array is fine
	x[0] = 999
	x = append(x, 40, 50, 60, 70)
	return x // make sure to return the updated slice if using append
}
