package main

import (
	"learn-go/db"
	"learn-go/stringops"
	"log"
)

func main() {
	c, err := db.NewConn("postgres")
	if err != nil {
		log.Fatal(err)
	}
	c.Insert()
	stringops.TrimAndUpper("   hello    ")

	//log.New
	//os.OpenFile()

}
