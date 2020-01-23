package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wugalde19/pratik/mc1/pkg/http_multiplexer"
)

func RegistrationHandler(mux http_multiplexer.IMultiplexer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

		if r.Body == nil {
			fmt.Fprint(w, "unable to handle request. No body provided.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body := RequestBody{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&body); err != nil {
			fmt.Fprint(w, "unable to decode request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		fmt.Fprintf(w, "%s has been successfully registered!", body.Name)
	}
}
