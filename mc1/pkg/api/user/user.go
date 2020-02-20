package user

import (
	"database/sql"

	v1 "github.com/wugalde19/pratik/mc1/pkg/api/user/v1"
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
	"github.com/wugalde19/pratik/mc1/pkg/middleware/jwt"
)

func AllRoutes(
	multiplexer http_multiplexer.IMultiplexer,
	dbConnection *sql.DB,
	jwt jwt.JWTService,
) {
	userService := v1.NewService(dbConnection)
	v1.Routes(multiplexer, userService, jwt)
}
