package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/db"
	"github.com/thegera4/cool-morning-lights-api/routes"
)

func TestMainFunction(t *testing.T) {
	// Configure the Gin router in test mode
	gin.SetMode(gin.TestMode)

	// Initialize the database
	db.InitDB()

	// Configure the server
	server := setupServer()

	// Register a test route
	server.GET("/testroute", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	// Register the routes
	routes.RegisterRoutes(server)

	// Create a request to send to to the test route
	req, err := http.NewRequest(http.MethodGet, "/testroute", nil)
	if err != nil { t.Fatalf("Couldn't create request: %v\n", err) }

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Create a context to pass to the handler
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Serve the request to the test route
	server.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
