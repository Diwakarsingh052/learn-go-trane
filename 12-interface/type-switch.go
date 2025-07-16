package main

import (
	"fmt"
)

// TypeDescriber is a simple function that takes any value and prints a description
// based on its type using a type switch
func TypeDescriber(v interface{}) {
	// Using type switch to handle different types
	switch x := v.(type) {
	case nil:
		fmt.Println("nil value")
	case int:
		fmt.Println("integer with value:", x)
	case float64:
		fmt.Println("float with value:", x)
	case bool:
		if x {
			fmt.Println("boolean with value: true")
		} else {
			fmt.Println("boolean with value: false")
		}
	case string:
		fmt.Println("string with value:", x)
	case []int:
		fmt.Println("integer slice with length:", len(x))
	case map[string]interface{}:
		fmt.Println("map with", len(x), "key-value pairs")
	case func():
		fmt.Println("function with no parameters")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	// Test with various types
	values := []interface{}{
		42,                                       // int
		3.14159,                                  // float64
		true,                                     // bool
		"Hello, Go!",                             // string
		nil,                                      // nil
		[]int{1, 2, 3, 4, 5},                     // slice
		map[string]interface{}{"key": "value"},   // map
		func() { fmt.Println("I'm a function") }, // function
		struct{ name string }{"custom struct"},   // custom type
	}

	fmt.Println("Type Switch Demonstration")
	fmt.Println("========================")

	// Process each value using the type switch
	for i, val := range values {
		fmt.Printf("Value %d is a ", i+1)
		TypeDescriber(val)

		// For demonstration, execute the function if it's a function type
		if fn, ok := val.(func()); ok {
			fmt.Print("Executing function: ")
			fn()
		}
	}
}
