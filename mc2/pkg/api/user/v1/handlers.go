package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdatePasswordHandler(routesMng RoutesManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := UpdatePasswordResponse{}

		token, err := routesMng.jwt.ParseToken(r)
		if err != nil || !token.Valid {
			response.Message = "Not Authorized"
			generateResponse(w, response, http.StatusForbidden)
			return
		}

		if r.Body == nil {
			fmt.Fprint(w, "unable to handle request. No body provided.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		model := UpdatePasswordRequestParams{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&model); err != nil {
			fmt.Fprint(w, "unable to decode request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		err = routesMng.service.updatePassword(
			model.UserId, model.NewPassword, model.OldPassword,
		)
		if err != nil {
			response.Message = fmt.Sprintf("unable to update password: %s", err.Error())
			generateResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Message = "Password updated successfully."
		generateResponse(w, response, http.StatusOK)
	}
}

func UserCountHandler(routesMng RoutesManager) func(http.ResponseWriter, *http.Request) {
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
	response interface{},
	statusCode int,
) {
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(fmt.Errorf("error while marshaling UserDetailsResponse response: %s", err.Error()))
	}

	w.WriteHeader(statusCode)
	w.Write(responseJson)
}
