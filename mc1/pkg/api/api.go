package api

import (
	"github.com/wugalde19/pratik/mc1/config"

	"github.com/wugalde19/pratik/mc1/pkg/api/registration"
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
	gojimultiplexer "github.com/wugalde19/pratik/mc1/pkg/http_multiplexer/goji"
	"github.com/wugalde19/pratik/mc1/pkg/server"
)

type registerRoutesFn func(http_multiplexer.IMultiplexer)

func Start(cfg *config.Config) {
	mux := gojimultiplexer.New(cfg.Server.Host, cfg.Server.Port)

	registerRoutes(mux, registration.AllRoutes)

	srv := server.New(mux)
	srv.Serve()
}

func registerRoutes(mux http_multiplexer.IMultiplexer, funcs ...registerRoutesFn) {
	for _, function := range funcs {
		function(mux)
	}
}
