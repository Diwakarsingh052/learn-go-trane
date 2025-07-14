package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// default value of error is nil
// which means no error
func main() {
	greet()

}

func greet() {
	args := os.Args[1:]
	if len(args) != 3 {
		log.Println("invalid args")
		//log.Fatalln("usage: greet name age gender")
		return
	}

	name := args[0]
	fmt.Println("hello", name)
	ageString := args[1]
	marksString := args[2]
	age, err := strconv.Atoi(ageString)
	if err != nil {
		// error happened
		log.Println("invalid age")
		return
	}

	marks, err := strconv.Atoi(marksString)
	if err != nil {
		log.Println("invalid marks")
		return
	}

	fmt.Println(name, age, marks)

}
