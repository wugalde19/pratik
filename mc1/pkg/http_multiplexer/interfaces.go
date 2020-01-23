package http_multiplexer

import (
	"net/http"
)

// IMultiplexer is a HTTP multiplexer/router.
type IMultiplexer interface {
	// Post will register an endpoint to handle the "POST" HTTP verb.
	Post(string, func(http.ResponseWriter, *http.Request))

	// Serve listens on the port number and calls ServeHTTP to handle incoming requests.
	Serve()

	// Use appends a middleware to the middleware stack.
	Use(middleware func(http.Handler) http.Handler)
}
