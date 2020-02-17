package registration

import (
	"database/sql"

	registration_v1 "github.com/wugalde19/pratik/mc1/pkg/api/registration/v1"
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
)

func AllRoutes(
	multiplexer http_multiplexer.IMultiplexer,
	dbConnection *sql.DB,
) {
	registrationService := registration_v1.NewService(dbConnection)
	registration_v1.Routes(multiplexer, registrationService)
}
