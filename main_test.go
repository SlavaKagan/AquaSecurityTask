// Golang REST API unit testing program
package main

import (
	"AquaSecurityChallenge/pkg/db"
	"AquaSecurityChallenge/pkg/handlers"
	"AquaSecurityChallenge/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

/* Connect to the DB */
func initTest() {
	db.ConnectDB()
}

/* Clear the containers table */
func ClearContainersTable() {
	db := db.GetDB()
	db.Exec("DELETE FROM containers")
}

/* Testing get all hosts */
func TestGetAllHosts(t *testing.T) {
	initTest()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/host", handlers.GetAllHosts)

	req, err := http.NewRequest(http.MethodGet, "/host", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	expected := "[\n    {\n        \"id\": 1,\n        \"uuid\": \"4e9edc48-2869-4172-903d-65008fd2895e\",\n        \"name\": \"AWS Host\",\n        \"ip_address\": \"1.2.3.4\"\n    },\n    {\n        \"id\": 2,\n        \"uuid\": \"f89cda2e-628a-4f6e-b1d8-1ecf389e2454\",\n        \"name\": \"Azure Host\",\n        \"ip_address\": \"4.5.6.7\"\n    },\n    {\n        \"id\": 3,\n        \"uuid\": \"863d9084-935e-4c71-990d-3a7dad113097\",\n        \"name\": \"GCP Host\",\n        \"ip_address\": \"7.8.9.0\"\n    },\n    {\n        \"id\": 4,\n        \"uuid\": \"86d33421-8945-4a50-bcf6-fd3750e51942\",\n        \"name\": \"IBM Host\",\n        \"ip_address\": \"2.4.6.8\"\n    }\n]"
	if w.Body.String() != expected {
		t.Fatalf("Expected to get body %s but instead got %s\n", expected, w.Body)
	}
}

/* Testing get all containers */
func TestGetAllContainers(t *testing.T) {
	initTest()
	ClearContainersTable()
	db.CreateContainer(2, "test_name")

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/container", handlers.GetAllContainers)

	req, err := http.NewRequest(http.MethodGet, "/container", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	expectedSize := 1
	var res []models.Container
	json.Unmarshal(w.Body.Bytes(), &res)

	if len(res) != expectedSize {
		t.Fatalf("Expected to get size %d but instead got %d\n", expectedSize, len(res))
	}
	ClearContainersTable()
}

/* Testing get host by Id */
func TestGetHostById(t *testing.T) {
	initTest()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/host/:id", handlers.GetHostByID)

	req, err := http.NewRequest(http.MethodGet, "/host/2", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	expected := "{\n    \"id\": 2,\n    \"uuid\": \"f89cda2e-628a-4f6e-b1d8-1ecf389e2454\",\n    \"name\": \"Azure Host\",\n    \"ip_address\": \"4.5.6.7\"\n}"
	if w.Body.String() != expected {
		t.Fatalf("Expected to get body %s but instead got %s\n", expected, w.Body)
	}
}

/* Testing get container by ID */
func TestGetContainerById(t *testing.T) {
	initTest()
	ClearContainersTable()
	db.CreateContainer(2, "test_name")
	containers := db.GetAllContainersFromDB()
	expected := containers[0].ID
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/container/:id", handlers.GetContainerByID)

	req, err := http.NewRequest(http.MethodGet, "/container/"+strconv.Itoa(expected), nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String())

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	var res models.Container
	json.Unmarshal(w.Body.Bytes(), &res)

	if res.ID != expected {
		t.Fatalf("Expected to get ID %d but instead got %d\n", expected, res.ID)
	}
	ClearContainersTable()
}
