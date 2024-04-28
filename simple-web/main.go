package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello simple web")
}
