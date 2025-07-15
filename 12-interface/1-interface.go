package main

import (
	"bufio"
	"fmt"
)

// Polymorphism means that a piece of code changes
// its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike

// Bigger the interface weaker the abstraction // Rob Pike

type Reader interface {
	Read(b []byte) (int, error)
	//abc() // // all methods must be implemented to implement the interface
}
type File struct {
	Name string
}

func (f File) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing them", f.Name)
	return 0, nil
}

type IO struct {
	name string
}

func (i IO) Read(b []byte) (int, error) {
	fmt.Println("reading and processing io ", i.name)
	return 0, nil
}

func ReadSomething(r Reader) {
	n, err := r.Read(nil)

	_, _ = n, err
}

func main() {
	f := File{Name: "file1"}
	i := IO{name: "io1"}
	ReadSomething(f)
	ReadSomething(i)

	bufio.NewReader(f)
}
