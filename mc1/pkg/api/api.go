package api

import (
	"fmt"
	"github.com/wugalde19/pratik/mc1/config"
	"github.com/wugalde19/pratik/mc1/pkg/middleware/jwt"

	"github.com/wugalde19/pratik/mc1/pkg/api/registration"
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
	gojimultiplexer "github.com/wugalde19/pratik/mc1/pkg/http_multiplexer/goji"
	"github.com/wugalde19/pratik/mc1/pkg/server"
)

type registerRoutesFn func(http_multiplexer.IMultiplexer)

func Start(cfg *config.Config) {
	mux := gojimultiplexer.New(cfg.Server.Host, cfg.Server.Port)

	registerRoutes(mux, registration.AllRoutes)

	jwt, err := jwt.New(cfg.JWT.SigningKeyEnv, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	if err != nil {
		panic(fmt.Errorf("problem occured while creating JWT middleware. %s", err.Error()))
	}

	mux.Use(jwt.MWFunc)

	srv := server.New(mux)
	srv.Serve()
}

func registerRoutes(mux http_multiplexer.IMultiplexer, funcs ...registerRoutesFn) {
	for _, function := range funcs {
		function(mux)
	}
}
