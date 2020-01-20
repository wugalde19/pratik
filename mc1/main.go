package main

import (
	"net/http"

	"github.com/wugalde19/pratik/mc1/lib/endpoints"
	"github.com/wugalde19/pratik/mc1/lib/services"
	"github.com/wugalde19/pratik/mc1/middleware"
	"goji.io"
)

func main() {
	mux := goji.NewMux()

	authenticator := services.NewAuthenticator()

	mux.Use(middleware.AuthMiddleware(authenticator))

	endpoints.RegisterAll(mux)

	http.ListenAndServe("localhost:8000", mux)
}
