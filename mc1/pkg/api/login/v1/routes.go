package v1

import (
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
	"github.com/wugalde19/pratik/mc1/pkg/middleware/jwt"
)

type RoutesManager struct {
	service *Service
	mux     http_multiplexer.IMultiplexer
}

// Registers all v1 routes related with login
func Routes(
	mux http_multiplexer.IMultiplexer,
	service *Service,
	jwt jwt.JWTService,
) {

	routesMng := RoutesManager{service, mux}

	// ------ Login
	// Version: 1
	// HTTP Verb: POST
	// Handler Func Name: LoginHandler
	mux.Post("/v1/login/", LoginHandler(routesMng, jwt))
}
