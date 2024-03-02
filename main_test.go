package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProductsHandler(t *testing.T) {
	// Set up a test router
	router := gin.Default()
	router.GET("/products", productsHandler)

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the HTTP request to the test router
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "text/html; charset=utf-8"
	if ct := rr.Header().Get("Content-Type"); ct != expected {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			ct, expected)
	}
}

func productsHandler(c *gin.Context) {
	// This is the handler function for /products route
	// You can place your existing handler logic here
}
