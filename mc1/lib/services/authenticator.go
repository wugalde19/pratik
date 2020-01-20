package services

import (
	"net/http"
	"os"
)

type Authenticator struct{}

func NewAuthenticator() Authenticator {
	return Authenticator{}
}

func (a Authenticator) Authenticate(token string) int {
	if token == "" {
		return http.StatusForbidden
	}

	signingKey := os.Getenv("SIGNING_KEY_ENV")
	if signingKey == "" {
		return http.StatusInternalServerError
	}

	if token != signingKey {
		return http.StatusForbidden
	}

	return http.StatusOK
}
