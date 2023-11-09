package awxclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/ping/", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a new AWX client with the mock server's URL
	client := NewClient(server.URL, "mockToken")

	// Test the Ping function
	resp, err := client.Ping()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}
