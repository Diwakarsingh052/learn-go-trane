package main

import (
	"log"
	"os"
)

func main() {
	// defer statement runs when the function is returning
	// defer statements gives guarantee of exec , once registered
	// defer maintains the stack
	// first in last out
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//panic("")
	//fmt.Println(3)

	f, err := os.Open("test.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

}
