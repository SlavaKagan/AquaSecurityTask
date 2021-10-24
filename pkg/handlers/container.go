package handlers

import (
	"AquaSecurityChallenge/pkg/db"
	"AquaSecurityChallenge/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getAllContainers responds with the list of all containers as JSON.

func GetAllContainers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.GetAllContainersFromDB())
}

// getContainerByID locates the host whose ID value matches the id
//parameter sent by the client, then returns that container as a response.

func GetContainerByID(c *gin.Context) {
	id := c.Param("id")
	searchId, _ := strconv.Atoi(id)

	var container = db.GetContainerByIDFromDB(searchId)
	if container == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Container not found"})
	} else {
		c.IndentedJSON(http.StatusOK, container)
	}
}

// Get container for specific host

func GetContainerForHost(c *gin.Context) {
	id := c.Param("id")
	hostId, _ := strconv.Atoi(id)

	var host = db.GetHostByIDFromDB(hostId)
	if host == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Host not found"})
	} else {
		c.IndentedJSON(http.StatusOK, db.GetContainersForHost(hostId))
	}
}

// Create a new container with specific request format

func PostContainer(c *gin.Context) {
	reqBody := new(models.RequestBody)
	c.Bind(reqBody)
	created := db.CreateContainer(reqBody.Host_id, reqBody.Image_name)
	if created {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Container created successfully"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}
}
