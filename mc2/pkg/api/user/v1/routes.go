package v1

import (
	"github.com/wugalde19/pratik/mc2/pkg/http_multiplexer"
	"github.com/wugalde19/pratik/mc2/pkg/middleware/jwt"
)

type RoutesManager struct {
	jwt     jwt.JWTService
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

	// ------ Update Password
	// Version: 1
	// HTTP Verb: POST
	// Handler Func Name: UpdatePasswordHandler
	mux.Post("/v1/user/update-password", UpdatePasswordHandler(routesMng))

	// ------ Get Total Registered User Count
	// Version: 1
	// HTTP Verb: GET
	// Handler Func Name: UserCountHandler
	mux.Get("/v1/user/count", UserCountHandler(routesMng))
}
