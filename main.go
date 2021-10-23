package main

import (
	"AquaSecurityChallenge/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/host", handlers.GetAllHosts).Methods(http.MethodGet)           // getAllHosts
	router.HandleFunc("/container", handlers.GetAllContainers).Methods(http.MethodGet) // getAllContainers

	router.HandleFunc("/host/{id}", handlers.GetHostById).Methods(http.MethodGet)           // getHostById
	router.HandleFunc("/container/{id}", handlers.GetContainerById).Methods(http.MethodGet) //getContainerById

	router.HandleFunc("/container/{hostId}", handlers.GetContainerForHost).Methods(http.MethodGet) //getContainersForHost

	router.HandleFunc("/container", handlers.AddContainer).Methods(http.MethodPost) //postContainer

	log.Println("API is running!")
	http.ListenAndServe(":9090", router)
}
