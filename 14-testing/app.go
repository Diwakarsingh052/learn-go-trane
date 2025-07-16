package main

import (
	"net/http"
	_ "net/http/pprof"
) // here be dragons

func main() {
	
	go func() { http.ListenAndServe(":3000", nil) }()
	//err := startApp()

}

//http://localhost:3000/debug/pprof // go tool pprof http://localhost:3000/debug/pprof/goroutine
