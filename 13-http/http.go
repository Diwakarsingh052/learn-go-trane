package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", Home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("server started")
}
func Home(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter, is used to write resp to the client
	// http.Request// anything user send us would be in the request struct
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//panic("") // http can handle panic automatically
	// panic inside goroutine kills the server
	//go panic("error") // this panic below would quit our service
	//time.Sleep(time.Second * 5)
	ctx := r.Context()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	default:
		fmt.Println("default")

	}
	w.Write([]byte("Hello World"))
}
