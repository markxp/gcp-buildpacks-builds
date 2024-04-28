package main

import (
	"net/http"

	"example.com/hello"
)

func main() {
	http.HandleFunc("/", hello.Hello)
	http.ListenAndServe(":8080", nil)
}
