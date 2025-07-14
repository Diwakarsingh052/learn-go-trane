package db

import (
	"fmt"
)

// making first letter uppercase exports the function

// don't export global variables ,
//until you have no problem if someone changes that

var DB string = "mysql"

func Insert() {
	fmt.Println("inserted in ", DB)
}
