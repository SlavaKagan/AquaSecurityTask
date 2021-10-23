package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// container represents dat
type container struct {
	ID         string `json:"id"`
	Host_ID    string `json:"host_id"`
	Name       string `json:"name"`
	Image_Name string `json:"image_name"`
	Host_name  string `json:"host_name"`
}

type host struct {
	ID string `json:"id"`
}

// containers slice to seed record container data.
var containers = []container{
	{ID: "1", Host_ID: "2", Name: "06e461b1-2673-45fc-4ea6-2add1480c014", Image_Name: "nginx", Host_name: "Azure Host"},
}

var hosts = []host{
	{ID: "1"},
}

func main() {
	router := gin.Default()
	router.GET("/host", getHosts)                  // getAllHosts
	router.GET("/container", getContainers)        // getAllContainers
	router.GET("/host/:id", getHostByID)           // getHostById
	router.GET("/container/:id", getContainerByID) // getContainerById
	//router.GET("/container/:id", getContainersForHost) // getContainersForHost
	//router.POST("/container", postContainer)
	router.Run("localhost:9090")
}

// getHosts responds with the list of all hosts as JSON.
func getHosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, hosts)
}

// getContainers responds with the list of all containers as JSON.
func getContainers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, containers)
}

// getHostByID locates the host whose ID value matches the id
// parameter sent by the client, then returns that host as a response.
func getHostByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of hosts, looking for
	// a host whose ID value matches the parameter.
	for _, a := range hosts {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "host not found"})
}

// getContainerByID locates the container whose ID value matches the id
// parameter sent by the client, then returns that container as a response.
func getContainerByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of containers, looking for
	// a container whose ID value matches the parameter.
	for _, a := range containers {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "container not found"})
}

// postContainer adds an container from JSON received in the request body.
func postContainer(c *gin.Context) {
	var newContainer container

	// Call BindJSON to bind the received JSON to newContainer.
	if err := c.BindJSON(&newContainer); err != nil {
		return
	}

	// Add the new container to the slice.
	containers = append(containers, newContainer)
	c.IndentedJSON(http.StatusCreated, newContainer)
}
