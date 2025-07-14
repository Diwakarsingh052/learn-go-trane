package db

import "fmt"

// making first letter uppercase exports the function

func Insert() {
	fmt.Println("inserted")
}

//q2. Create a stringops package with three files. The package should export one function that utilizes internal (unexported) functions.
//Files in stringops Package:
//File 1: strings.go
//Exported Function: TrimSpaceAndUppercase(s1, s2 string) string
//File 2: upper.go
//Internal Function: toUpperCase(s string) string
//File 3: trim.go
//Internal Function: trimSpace(s string) string
