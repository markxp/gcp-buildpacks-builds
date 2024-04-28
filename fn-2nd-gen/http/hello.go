package hello

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("hello", get)
}

func get(w http.ResponseWriter, r *http.Request) {
	paths := os.Getenv("PATH")
	io.WriteString(w, fmt.Sprintf("hello, cloud functions 2nd generation HTTP.\nPATH=%s",paths))
}
