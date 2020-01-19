package main

import (
	"net/http"

	"github.com/wugalde19/pratik/mc1/lib/endpoints"
	"goji.io"
)

func main() {
	mux := goji.NewMux()

	endpoints.RegisterAll(mux)

	http.ListenAndServe("localhost:8000", mux)
}
