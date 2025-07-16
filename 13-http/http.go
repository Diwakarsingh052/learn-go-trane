package main

import (
	"encoding/json"
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

func sendJsonResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user struct {
		FirstName string `json:"first_name"` // field tag // we are giving what the name of field should be in json output
		Password  string `json:"-"`          // ignoring the field in JSON output
		Email     string `json:"email"`
	}
	user.FirstName = "abc"
	user.Password = "123"
	user.Email = "abc@gmail.com"

	// json.Marshal converts a type to json
	jsonData, err := json.Marshal(user)
	if err != nil {
		// sending text based error resp
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// setting header and then sending a response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
