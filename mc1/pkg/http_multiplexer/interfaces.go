package http_multiplexer

// IMultiplexer is a HTTP multiplexer/router.
type IMultiplexer interface {
	// Serve listens on the port number and calls ServeHTTP to handle incoming requests.
	Serve()
}
