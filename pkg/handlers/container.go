package handlers

import (
	"AquaSecurityChallenge/pkg/db"
	"AquaSecurityChallenge/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/* GetAllContainers responds with the list of all hosts as JSON after taking it from DB */

func GetAllContainers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.GetAllContainersFromDB())
}

/* GeContainerByID locates the container whose ID value matches the id
parameter sent by the client, then returns that host as a json response after taking it from DB */

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

/* GetContainerForHost locates the container whose HostID value matches the id
parameter sent by the client, then returns that host as a json response after taking it from DB */

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

/* PostContainer creates a new container and saving it to the DB */

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
