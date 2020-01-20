package middleware

import (
	"fmt"
	"net/http"

	"github.com/wugalde19/pratik/mc1/lib/services"
)

const (
	tokenParamName = "token"
)

func AuthMiddleware(authService services.Authenticator) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.URL.Query().Get(tokenParamName)

			if statusCode := authService.Authenticate(token); statusCode != http.StatusOK {
				w.WriteHeader(statusCode)
				handleNotOkStatus(statusCode, w)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}

func handleNotOkStatus(statusCode int, w http.ResponseWriter) {
	switch statusCode {
	case http.StatusForbidden:
		fmt.Fprintf(w, "Not Authorized\n")
	case http.StatusInternalServerError:
		fmt.Fprintf(w, "Error occurred during authorization process\n")
	}
}
