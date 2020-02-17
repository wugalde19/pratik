package api

import (
	"database/sql"
	"fmt"

	"github.com/wugalde19/pratik/mc1/config"
	"github.com/wugalde19/pratik/mc1/pkg/api/login"
	"github.com/wugalde19/pratik/mc1/pkg/api/registration"
	"github.com/wugalde19/pratik/mc1/pkg/db/postgres"
	"github.com/wugalde19/pratik/mc1/pkg/middleware/jwt"

	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
	gojimultiplexer "github.com/wugalde19/pratik/mc1/pkg/http_multiplexer/goji"
	"github.com/wugalde19/pratik/mc1/pkg/server"
)

type registerRoutesFn func(http_multiplexer.IMultiplexer, *sql.DB)

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

	registerPublicRoutes(mux, dbConnection, *jwt)

	mux.Use(jwt.MWFunc)

	srv := server.New(mux)
	srv.Serve()

}

func registerPublicRoutes(
	mux http_multiplexer.IMultiplexer,
	dbConnection *sql.DB,
	jwt jwt.JWTService,
) {
	login.AllRoutes(mux, dbConnection, jwt)
	registration.AllRoutes(mux, dbConnection)
}
