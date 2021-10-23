package handlers

import (
	"AquaSecurityChallenge/pkg/mocks"
	"AquaSecurityChallenge/pkg/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddContainer(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var container models.Container
	json.Unmarshal(body, &container)

	// Append to the Book mocks
	container.ID = rand.Intn(100)
	mocks.Containers = append(mocks.Containers, container)

	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
