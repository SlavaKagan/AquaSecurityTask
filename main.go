package main

import (
	"AquaSecurityChallenge/pkg/db"
	"AquaSecurityChallenge/pkg/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db.ConnectDB()

	router := gin.Default()                             // init a new router
	router.GET("/host", handlers.GetAllHosts)           // getAllHosts
	router.GET("/container", handlers.GetAllContainers) // getAllContainers

	router.GET("/host/:id", handlers.GetHostByID)           // getHostById
	router.GET("/container/:id", handlers.GetContainerByID) // getContainerById

	router.GET("/host/:id/container", handlers.GetContainerForHost) // getContainerByHostId

	router.POST("/container", handlers.PostContainer) // postContainer

	log.Println("API is running!")
	router.Run("localhost:9090") // port for application
}
