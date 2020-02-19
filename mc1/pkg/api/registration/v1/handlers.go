package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RegistrationHandler(routesMng RoutesManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

		if r.Body == nil {
			fmt.Fprint(w, "unable to handle request. No body provided.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		model := RegistrationModel{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&model); err != nil {
			fmt.Fprint(w, "unable to decode request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		if err := routesMng.service.createRegistration(model); err != nil {
			fmt.Fprint(w, fmt.Sprintf("unable to register user: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fmt.Sprintf("User %s succesfully created.", model.Name))
	}
}
