package api

import (
	"github.com/wugalde19/pratik/mc1/config"
	gojimultiplexer "github.com/wugalde19/pratik/mc1/pkg/http_multiplexer/goji"
	"github.com/wugalde19/pratik/mc1/pkg/server"
)

func Start(cfg *config.Config) {
	mux := gojimultiplexer.New(cfg.Server.Host, cfg.Server.Port)
	srv := server.New(mux)
	srv.Serve()
}
