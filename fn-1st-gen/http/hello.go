// Package http provides a 1st-gen HTTP function.
package http

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		e := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(e), e)
	}
	io.WriteString(w, "hello, cloud function 1st gen HTTP.")
}
