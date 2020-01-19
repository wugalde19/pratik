package endpoints

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func RegisterAll(mux *goji.Mux) {
	endpointsDefinitions := endpointsDefinitions()

	for _, endpoint := range endpointsDefinitions {
		registerEndpoint(mux, endpoint)
	}
}

func registerEndpoint(mux *goji.Mux, endpoint definition) {
	switch endpoint.Method {
	case http.MethodGet:
		mux.HandleFunc(pat.Get(endpoint.Route), endpoint.HandlerFunc)
	}
}
