package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wugalde19/pratik/mc1/pkg/middleware/jwt"
)

// LoginHandler handles request made to /v1/login
func LoginHandler(
	routesMng RoutesManager,
	jwt jwt.JWTService,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Body == nil {
			response := LoginResponse{Message: "unable to handle request. No body provided."}
			generateResponse(w, response, http.StatusBadRequest)
			return
		}

		loginData := LoginData{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&loginData); err != nil {
			response := LoginResponse{Message: "unable to decode request body"}
			generateResponse(w, response, http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		email, err := routesMng.service.executeLogin(loginData)
		if err != nil {
			response := LoginResponse{Message: err.Error()}
			generateResponse(w, response, http.StatusBadRequest)
			return
		}

		token, _, _ := jwt.GenerateToken(email)
		msg := fmt.Sprintf("valid credentials for %s", email)
		response := LoginResponse{Message: msg, Token: token}
		generateResponse(w, response, http.StatusOK)
	}
}

func generateResponse(w http.ResponseWriter, response LoginResponse, statusCode int) {
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(fmt.Errorf("error while marshaling login response: %s", err.Error()))
	}

	w.WriteHeader(statusCode)
	w.Write(responseJson)
}
