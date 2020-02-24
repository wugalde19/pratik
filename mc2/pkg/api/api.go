package api

import (
	"fmt"

	"github.com/wugalde19/pratik/mc2/config"
	"github.com/wugalde19/pratik/mc2/pkg/db/postgres"
	"github.com/wugalde19/pratik/mc2/pkg/middleware/jwt"

	gojimultiplexer "github.com/wugalde19/pratik/mc2/pkg/http_multiplexer/goji"
	"github.com/wugalde19/pratik/mc2/pkg/server"
)

func Start(cfg *config.Config) {
	mux := gojimultiplexer.New(cfg.Server.Host, cfg.Server.Port)

	database, err := postgres.NewDatabase(cfg.DB)
	if err != nil {
		panic(fmt.Errorf("problem occured while creating database. %s", err.Error()))
	}

	dbConnection, err := database.Connect()
	if err != nil {
		panic(fmt.Errorf("problem occured while connecting to database. %s", err.Error()))
	}

	jwt, err := jwt.New(cfg.JWT.SigningKeyEnv, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	if err != nil {
		panic(fmt.Errorf("problem occured while creating JWT middleware. %s", err.Error()))
	}

	mux.Use(jwt.MWFunc)

	srv := server.New(mux)
	srv.Serve()

}
