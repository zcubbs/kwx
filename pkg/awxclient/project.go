package awxclient

import "fmt"

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Client) ListAllProjects() ([]Project, error) {
	var projects struct {
		Results []Project `json:"results"`
	}

	resp, err := c.httpClient.R().SetResult(&projects).Get("/api/v2/projects/")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}

	return projects.Results, nil
}

func (c *Client) GetProjectByID(projectID int) (*Project, error) {
	var project Project

	resp, err := c.httpClient.R().SetResult(&project).Get(fmt.Sprintf("/api/v2/projects/%d/", projectID))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}

	return &project, nil
}

func (c *Client) AddProject(newProject Project) (*Project, error) {
	var createdProject Project

	resp, err := c.httpClient.R().SetBody(newProject).SetResult(&createdProject).Post("/api/v2/projects/")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}

	return &createdProject, nil
}

func (c *Client) UpdateProject(projectID int, updatedProject Project) (*Project, error) {
	var project Project

	resp, err := c.httpClient.R().SetBody(updatedProject).SetResult(&project).Put(fmt.Sprintf("/api/v2/projects/%d/", projectID))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}

	return &project, nil
}

func (c *Client) DeleteProject(projectID int) error {
	resp, err := c.httpClient.R().Delete(fmt.Sprintf("/api/v2/projects/%d/", projectID))
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}

	return nil
}
