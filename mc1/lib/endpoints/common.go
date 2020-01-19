package endpoints

import (
	hellov1 "github.com/wugalde19/pratik/mc1/lib/handlers/hello/v1"
)

func endpointsDefinitions() definitions {
	return definitions{
		definition{
			Name:        "Hello",
			Method:      "GET",
			Route:       "/hello/v1/:name",
			HandlerFunc: hellov1.HelloHandler,
		},
	}
}
