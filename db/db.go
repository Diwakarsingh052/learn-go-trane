package db

import (
	"errors"
	"fmt"
)

// making first letter uppercase exports the function

// don't export global variables ,
//until you have no problem if someone changes that

//var DB string = "mysql"

type Conn struct {
	db string // no one can change this outside the current package
}

// create a NewFunc to initialize the struct fields which are not
// exported

func NewConn(db string) (Conn, error) {
	if db == "" {
		return Conn{}, errors.New("db name is empty")
	}
	return Conn{db: db}, nil
}

func (c Conn) Insert() {
	fmt.Println("inserted in ", c.db)
}
