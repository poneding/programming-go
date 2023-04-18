package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(time.Second * 10):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println("server:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

/*
$ go run context.go
server: hello handler started
server: context canceled
server: hello handler ended


$ curl localhost:8080/hello
^C
*/
