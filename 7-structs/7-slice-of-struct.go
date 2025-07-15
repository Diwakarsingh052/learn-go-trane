package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) print() {
	fmt.Println(u.name, u.age)
}
func main() {

	u := []user{
		{
			name: "abc",
			age:  20,
		},
		{
			name: "abc",
			age:  20,
		},
		{
			name: "abc",
			age:  20,
		},
	}

	for _, v := range u { // range returns index and value, if want to ignore value use _
		v.print()
	}

}
