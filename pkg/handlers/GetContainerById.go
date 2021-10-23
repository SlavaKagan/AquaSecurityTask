package handlers

import (
	"AquaSecurityChallenge/pkg/mocks"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetContainerById(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock hosts
	for _, container := range mocks.Containers {
		if container.ID == id {
			// If ids are equal send book as a response
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(container)
			break
		}
	}
}
