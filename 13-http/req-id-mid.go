package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type key string

const reqKey key = "reqId"

// go get moduleName
// get all deps for the project
// remove all the unused deps for a project
func main() {
	http.HandleFunc("/home", RequestId(Hello))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqId, ok := ctx.Value(reqKey).(string)
	fmt.Println(reqId)
	if !ok {
		reqId = "unknown"
	}

	fmt.Println("req processed, sent a response ", reqId)
	w.Write([]byte("Hello World"))
}

// The provided key must be comparable and should not be of
// type string or any other built-in type to avoid collisions
// between packages using context.
// Users of WithValue should define their own types for keys.

func RequestId(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := uuid.New()
		fmt.Println("req started with ", reqId)
		ctx := r.Context()
		ctx = context.WithValue(ctx, reqKey, reqId)
		//withContext would update the internal context of the request object
		//with the updated context with values or timeouts
		next(w, r.WithContext(ctx))
	}
}
