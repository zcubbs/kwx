package awxclient

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddCredential(t *testing.T) {
	mockCredential := Credential{
		Name: "TestCredential",
		// ... populate other fields
	}
	mockResponse := `{"id": 1, "name": "TestCredential"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/credentials/", r.URL.Path)
		var receivedCredential Credential
		_ = json.NewDecoder(r.Body).Decode(&receivedCredential)
		assert.Equal(t, mockCredential.Name, receivedCredential.Name)
		w.WriteHeader(http.StatusCreated) // Assuming AWX returns 201 (Created) for successful addition
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewClient(server.URL, "mockToken")
	resp, err := client.CreateCredential(mockCredential)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode())
}

func TestDeleteCredential(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/credentials/1/", r.URL.Path)
		w.WriteHeader(http.StatusNoContent) // Assuming AWX returns 204 (No Content) for successful delete
	}))
	defer server.Close()

	client := NewClient(server.URL, "mockToken")
	resp, err := client.DeleteCredential(1) // Using 1 as a mock credential ID
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode())
}

func TestUpdateCredential(t *testing.T) {
	mockUpdate := Credential{Name: "UpdatedTestCredential"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/credentials/1/", r.URL.Path)
		var receivedUpdate Credential
		_ = json.NewDecoder(r.Body).Decode(&receivedUpdate)
		assert.Equal(t, mockUpdate.Name, receivedUpdate.Name)
		w.WriteHeader(http.StatusOK) // Assuming AWX returns 200 (OK) for successful update
	}))
	defer server.Close()

	client := NewClient(server.URL, "mockToken")
	resp, err := client.UpdateCredential(1, mockUpdate)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
}
