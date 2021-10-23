package handlers

import (
	"AquaSecurityChallenge/pkg/mocks"
	"encoding/json"
	"net/http"
)

func GetAllHosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Hosts)
}
