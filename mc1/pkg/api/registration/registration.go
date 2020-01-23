package registration

import (
	v1 "github.com/wugalde19/pratik/mc1/pkg/api/registration/v1"
	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
)

func AllRoutes(multiplexer http_multiplexer.IMultiplexer) {
	v1.Routes(multiplexer)
}
