package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndToEnd(t *testing.T) {
	// Start the server in a goroutine
	go main()

	// Wait for the server to start
	// You might need to add some delay to ensure the server has enough time to start
	// time.Sleep(time.Second)

	// Make a request to the server
	resp, err := http.Get("http://localhost:8081/")
	assert.NoError(t, err, "Error making request")
	defer resp.Body.Close()

	// Check if the response status code is OK
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response status should be OK")

	// You can add more assertions to check the response body or headers if needed
	// For example, you can check if the response body contains expected data
	// body, err := ioutil.ReadAll(resp.Body)
	// assert.NoError(t, err, "Error reading response body")
	// assert.Contains(t, string(body), "Expected content", "Response body should contain expected content")
}
