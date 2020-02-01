package v1

import (
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
)

type RoutesManager struct {
	service *Service
	mux     http_multiplexer.IMultiplexer
}

// Registers all v1 routes related with registration
func Routes(
	mux http_multiplexer.IMultiplexer,
	service *Service,
) {

	routesMng := RoutesManager{service, mux}

	// ------ Registration
	// Version: 1
	// HTTP Verb: POST
	// Handler Func Name: RegistrationHandler
	mux.Post("/v1/registration/", RegistrationHandler(routesMng))
}
