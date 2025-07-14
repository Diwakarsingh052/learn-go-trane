package main

import (
	"learn-go/db"
	"learn-go/stringops"
)

func main() {
	db.Insert()
	stringops.TrimAndUpper("   hello    ")
}
