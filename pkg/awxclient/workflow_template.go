package awxclient

import (
	"encoding/json"
	"fmt"
)

type WorkflowTemplate struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateWorkflowTemplate creates a new workflow template.
func (c *Client) CreateWorkflowTemplate(wf WorkflowTemplate) (*WorkflowTemplate, error) {
	resp, err := c.httpClient.R().SetBody(wf).Post("/api/v2/workflow_job_templates/")
	if err != nil {
		return nil, err
	}
	var result WorkflowTemplate
	err = json.Unmarshal(resp.Body(), &result)
	return &result, err
}

// GetWorkflowTemplate retrieves a workflow template by its ID.
func (c *Client) GetWorkflowTemplate(id int) (*WorkflowTemplate, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/workflow_job_templates/%d/", id))
	if err != nil {
		return nil, err
	}
	var wf WorkflowTemplate
	err = json.Unmarshal(resp.Body(), &wf)
	return &wf, err
}

// DeleteWorkflowTemplate deletes a workflow template.
func (c *Client) DeleteWorkflowTemplate(id int) error {
	_, err := c.httpClient.R().Delete(fmt.Sprintf("/api/v2/workflow_job_templates/%d/", id))
	return err
}

// LaunchWorkflow launches a workflow based on a given template ID.
func (c *Client) LaunchWorkflow(templateID int, extraVars map[string]interface{}) (*Job, error) {
	resp, err := c.httpClient.R().SetBody(map[string]interface{}{
		"extra_vars": extraVars,
	}).Post(fmt.Sprintf("/api/v2/workflow_job_templates/%d/launch/", templateID))
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(resp.Body(), &job)
	return &job, err
}
