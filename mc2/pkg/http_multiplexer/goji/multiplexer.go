package goji

import (
	"fmt"
	"net/http"

	goji "goji.io"
	"goji.io/pat"

	"github.com/wugalde19/pratik/mc2/pkg/http_multiplexer"
)

type multiplexer struct {
	initialized bool
	host        string
	mux         *goji.Mux
	port        int
	server      *http.Server
}

// New returns a new IMultiplexer
// This implementation uses Goji framework (https://goji.io/)
func New(host string, port int) http_multiplexer.IMultiplexer {
	return &multiplexer{
		host: host,
		mux:  goji.NewMux(),
		port: port,
	}
}

// Get will register an endpoint to handle the "GET" HTTP verb.
func (m *multiplexer) Get(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	m.mux.HandleFunc(pat.Get(endpoint), handler)
}

// Post will register an endpoint to handle the "POST" HTTP verb.
func (m *multiplexer) Post(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	m.mux.HandleFunc(pat.Post(endpoint), handler)
}

// Serve will begin servicing requests.
func (m *multiplexer) Serve() {
	m.initialize()
	err := m.server.ListenAndServe()
	if err != nil {
		fmt.Errorf("failed calling ListenAndServe() on http.Server. %s", err.Error())
	}
}

// Use appends a middleware to the middleware stack.
func (m *multiplexer) Use(middleware func(http.Handler) http.Handler) {
	m.mux.Use(middleware)
}

func (m *multiplexer) initialize() {
	m.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", m.host, m.port),
		Handler: m.mux,
	}

	m.initialized = true
}
