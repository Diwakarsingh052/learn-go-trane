package main

import "fmt"

// when to use a pointer
// https://go.dev/doc/faq#methods_on_values_or_pointers:~:text=Should%20I%20define%20methods%20on%20values%20or%20pointers%3F%C2%B6
type author struct {
	name string
	age  int
}

func (a *author) updateName(name string) {
	a.name = name
}
func main() {
	var author author
	//a1 := &author{name: "abc", age: 18}
	author.updateName("Bob")
	fmt.Println(author.name)

}
