package awxclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (c *Client) AddUser(user User) (*resty.Response, error) {
	resp, err := c.httpClient.R().SetBody(user).Post("/api/v2/users/")
	return resp, err
}

func (c *Client) DeleteUser(userID int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/api/v2/users/%d/", userID)
	resp, err := c.httpClient.R().Delete(endpoint)
	return resp, err
}
