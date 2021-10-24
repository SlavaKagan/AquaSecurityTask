package handlers

import (
	"AquaSecurityChallenge/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getAllHosts responds with the list of all hosts as JSON.

func GetAllHosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.GetAllHostsFromDB())
}

// getHostByID locates the host whose ID value matches the id
//parameter sent by the client, then returns that host as a response.

func GetHostByID(c *gin.Context) {
	id := c.Param("id")
	searchId, _ := strconv.Atoi(id)

	var host = db.GetHostByIDFromDB(searchId)
	if host == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Host not found"})
	} else {
		c.IndentedJSON(http.StatusOK, host)
	}
}
