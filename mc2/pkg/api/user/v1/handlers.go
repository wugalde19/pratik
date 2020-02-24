package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UserHandler(routesMng RoutesManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := UsersCountResponse{}

		token, err := routesMng.jwt.ParseToken(r)
		if err != nil || !token.Valid {
			response.Error = "Not Authorized"
			generateResponse(w, response, http.StatusForbidden)
			return
		}

		count, err := routesMng.service.usersCount()
		if err != nil {
			response.Error = fmt.Sprintf("unable to get registered user count: %s", err.Error())
			generateResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Count = count
		generateResponse(w, response, http.StatusOK)
	}
}

func generateResponse(
	w http.ResponseWriter,
	response UsersCountResponse,
	statusCode int,
) {
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(fmt.Errorf("error while marshaling UserDetailsResponse response: %s", err.Error()))
	}

	w.WriteHeader(statusCode)
	w.Write(responseJson)
}
