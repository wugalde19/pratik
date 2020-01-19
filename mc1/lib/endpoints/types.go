package endpoints

import (
	"net/http"
)

type definition struct {
	Name        string
	Method      string
	Route       string
	HandlerFunc http.HandlerFunc
}

type definitions []definition
