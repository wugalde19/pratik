package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wugalde19/pratik/mc1/pkg/middleware/jwt"
)

func UserHandler(routesMng RoutesManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := UserDetailsResponse{}

		token, err := routesMng.jwt.ParseToken(r)
		if err != nil || !token.Valid {
			response.Error = "Not Authorized"
			generateResponse(w, response, http.StatusForbidden)
			return
		}

		claimsMap, ok := token.Claims.(*jwt.CustomMapClaims)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			response.Error = "unable to get info from token"
			generateResponse(w, response, http.StatusBadRequest)
			return
		}

		email := claimsMap.Email
		userModel, err := routesMng.service.findUserByEmail(email)
		if err != nil {
			response.Error = fmt.Sprintf("unable to find user details: %s", err.Error())
			generateResponse(w, response, http.StatusBadRequest)
			return
		}

		response.UserData = UserData {
			Name: userModel.Name,
			MobileNumber: userModel.MobileNumber,
		}

		generateResponse(w, response, http.StatusOK)
	}
}

func generateResponse(
	w http.ResponseWriter,
	response UserDetailsResponse,
	statusCode int,
) {
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(fmt.Errorf("error while marshaling UserDetailsResponse response: %s", err.Error()))
	}

	w.WriteHeader(statusCode)
	w.Write(responseJson)
}
