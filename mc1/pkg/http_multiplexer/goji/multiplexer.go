package goji

import (
	"fmt"
	"net/http"

	goji "goji.io"

	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
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

// Serve will begin servicing requests.
func (m *multiplexer) Serve() {
	m.initialize()
	err := m.server.ListenAndServe()
	if err != nil {
		fmt.Errorf("failed calling ListenAndServe() on http.Server. %s", err.Error())
	}
}

func (m *multiplexer) initialize() {
	m.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", m.host, m.port),
		Handler: m.mux,
	}

	m.initialized = true
}
