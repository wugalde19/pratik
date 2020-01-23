package v1

import (
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
)

func Routes(mux http_multiplexer.IMultiplexer) {

	// ------ Registration
	// Version: 1
	// HTTP Verb: POST
	// Handler Func Name: RegistrationHandler
	mux.Post("/v1/registration/", RegistrationHandler(mux))
}
