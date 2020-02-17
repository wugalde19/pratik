package login

import (
	"database/sql"

	v1 "github.com/wugalde19/pratik/mc1/pkg/api/login/v1"
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
	"github.com/wugalde19/pratik/mc1/pkg/middleware/jwt"
)

func AllRoutes(
	multiplexer http_multiplexer.IMultiplexer,
	dbConnection *sql.DB,
	jwt jwt.JWTService,
) {
	loginService := v1.NewService(dbConnection)
	v1.Routes(multiplexer, loginService, jwt)
}
