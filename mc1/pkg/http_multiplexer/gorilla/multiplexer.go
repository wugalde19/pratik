package gorilla

import (
	"fmt"
	"net/http"

	gorilla_mux "github.com/gorilla/mux"

	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
)

type multiplexer struct {
	initialized bool
	host        string
	mux         *gorilla_mux.Router
	port        int
	server      *http.Server
}

// New returns a new IMultiplexer
// This implementation uses Gorilla mux (http://www.gorillatoolkit.org/pkg/mux)
func New(host string, port int) http_multiplexer.IMultiplexer {
	return &multiplexer{
		host: host,
		mux:  gorilla_mux.NewRouter(),
		port: port,
	}
}

// Get will register an endpoint to handle the "GET" HTTP verb.
func (m *multiplexer) Get(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	m.mux.HandleFunc(endpoint, handler).Methods("GET")
}

// Post will register an endpoint to handle the "POST" HTTP verb.
func (m *multiplexer) Post(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	m.mux.HandleFunc(endpoint, handler).Methods("POST")
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
