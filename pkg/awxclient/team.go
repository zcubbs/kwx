package awxclient

import (
	"encoding/json"
	"fmt"
)

type Team struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization"`
}

// CreateTeam creates a new team within a specified organization.
func (c *Client) CreateTeam(team Team) (*Team, error) {
	resp, err := c.httpClient.R().SetBody(team).Post("/api/v2/teams/")
	if err != nil {
		return nil, err
	}
	var result Team
	err = json.Unmarshal(resp.Body(), &result)
	return &result, err
}

// GetTeam retrieves details of a team by its ID.
func (c *Client) GetTeam(id int) (*Team, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/teams/%d/", id))
	if err != nil {
		return nil, err
	}
	var team Team
	err = json.Unmarshal(resp.Body(), &team)
	return &team, err
}
