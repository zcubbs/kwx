package awxclient

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
)

type JobTemplate struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateJobTemplate struct {
	Description string `json:"description"`
}

func (c *Client) GetJobTemplates() ([]JobTemplate, error) {
	var jobTemplates []JobTemplate

	resp, err := c.httpClient.R().SetResult(&jobTemplates).Get("/api/v2/job_templates/")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}

	return jobTemplates, nil
}

func (c *Client) UpdateJobTemplate(jobTemplateID int, update UpdateJobTemplate) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/api/v2/job_templates/%d/", jobTemplateID)
	resp, err := c.httpClient.R().SetBody(update).Patch(endpoint)
	return resp, err
}

type JobTemplates struct {
	Count   int           `json:"count"`
	Next    string        `json:"next"`
	Results []JobTemplate `json:"results"`
}

func (c *Client) GetAllJobTemplates() ([]JobTemplate, error) {
	var allTemplates []JobTemplate
	page := 1

	for {
		endpoint := fmt.Sprintf("/api/v2/job_templates/?page=%d", page)
		var templates JobTemplates
		resp, err := c.httpClient.R().SetResult(&templates).Get(endpoint)
		if err != nil {
			return nil, err
		}
		if resp.IsError() {
			return nil, fmt.Errorf("API responded with status code %d", resp.StatusCode())
		}
		allTemplates = append(allTemplates, templates.Results...)

		// If there's no next page, break out of loop
		if templates.Next == "" {
			break
		}
		page++
	}

	return allTemplates, nil
}

func (c *Client) SearchJobTemplates(query string) ([]JobTemplate, error) {
	endpoint := fmt.Sprintf("/api/v2/job_templates/?search=%s", url.QueryEscape(query))
	var templates JobTemplates
	resp, err := c.httpClient.R().SetResult(&templates).Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}
	return templates.Results, nil
}

// CreateJobTemplate creates a new job template.
func (c *Client) CreateJobTemplate(template JobTemplate) (*JobTemplate, error) {
	resp, err := c.httpClient.R().SetBody(template).Post("/api/v2/job_templates/")
	if err != nil {
		return nil, err
	}
	var result JobTemplate
	err = json.Unmarshal(resp.Body(), &result)
	return &result, err
}

// GetJobTemplate retrieves a job template by its ID.
func (c *Client) GetJobTemplate(id int) (*JobTemplate, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/job_templates/%d/", id))
	if err != nil {
		return nil, err
	}
	var template JobTemplate
	err = json.Unmarshal(resp.Body(), &template)
	return &template, err
}
