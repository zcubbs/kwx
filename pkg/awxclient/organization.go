package awxclient

import (
	"encoding/json"
	"fmt"
)

type Organization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OrganizationDetail struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateOrganization creates a new organization.
func (c *Client) CreateOrganization(org Organization) (*Organization, error) {
	resp, err := c.httpClient.R().SetBody(org).Post("/api/v2/organizations/")
	if err != nil {
		return nil, err
	}
	var result Organization
	err = json.Unmarshal(resp.Body(), &result)
	return &result, err
}

// GetOrganization retrieves an organization by its ID.
func (c *Client) GetOrganization(id int) (*Organization, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/organizations/%d/", id))
	if err != nil {
		return nil, err
	}
	var org Organization
	err = json.Unmarshal(resp.Body(), &org)
	return &org, err
}

func (c *Client) GetOrganizationDetail(orgID int) (OrganizationDetail, error) {
	var org OrganizationDetail
	endpoint := fmt.Sprintf("/api/v2/organizations/%d/", orgID)
	resp, err := c.httpClient.R().SetResult(&org).Get(endpoint)
	if err != nil {
		return OrganizationDetail{}, err
	}
	if resp.IsError() {
		return OrganizationDetail{}, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}
	return org, nil
}
