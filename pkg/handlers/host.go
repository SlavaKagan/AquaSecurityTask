package handlers

import (
	"AquaSecurityChallenge/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/* GetAllHosts responds with the list of all hosts as JSON after taking it from DB*/

func GetAllHosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.GetAllHostsFromDB())
}

/* GetHostByID locates the host whose ID value matches the id
parameter sent by the client, then returns that host as a json response after taking it from DB.
*/

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
