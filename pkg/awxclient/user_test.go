package awxclient

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/users/", r.URL.Path)
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	client := NewClient(server.URL, "mockToken")
	user := User{
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@email.com",
	}
	resp, err := client.AddUser(user)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode())
}

func TestDeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/users/1/", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient(server.URL, "mockToken")
	resp, err := client.DeleteUser(1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode())
}
