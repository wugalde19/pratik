package v1

import (
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
	"github.com/wugalde19/pratik/mc1/pkg/middleware/jwt"
)

type RoutesManager struct {
	jwt 		jwt.JWTService
	service *Service
	mux     http_multiplexer.IMultiplexer
}

// Registers all v1 routes related with user
func Routes(
	mux http_multiplexer.IMultiplexer,
	service *Service,
	jwt jwt.JWTService,
) {

	routesMng := RoutesManager{jwt, service, mux}

	// ------ User
	// Version: 1
	// HTTP Verb: GET
	// Handler Func Name: UserHandler
	mux.Get("/v1/user/", UserHandler(routesMng))
}
