package main

import (
	"learn-go/db"
	"learn-go/stringops"
)

func main() {
	db.DB = "postgres"
	db.Insert()
	stringops.TrimAndUpper("   hello    ")
}
