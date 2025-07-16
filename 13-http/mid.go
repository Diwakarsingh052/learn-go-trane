package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", Mid(Mid(hello)))
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	w.Write([]byte("hello"))
}

func Mid(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("preprocessing logic")
		next(w, r)
		fmt.Println("postprocessing logic")
	}
}

//func abc() int {
//	return int
//	return 0
//}
