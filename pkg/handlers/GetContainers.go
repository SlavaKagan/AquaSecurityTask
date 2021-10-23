package handlers

import (
	"AquaSecurityChallenge/pkg/mocks"
	"encoding/json"
	"net/http"
)

func GetAllContainers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Containers)
}