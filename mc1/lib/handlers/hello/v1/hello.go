package hello

import (
	"fmt"
	"net/http"

	"goji.io/pat"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}
