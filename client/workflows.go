package client

import (
	"encoding/json"
	"fmt"
)

type WorkflowService service

// WorkflowTeamService is a service a workflow is restricted to.
type WorkflowTeamService struct {
	UniqueID string `json:"unique_id"`
	Name     string `json:"name,omitempty"`
}

// WorkflowTeam is the read shape of a team scope; no Services means every service.
type WorkflowTeam struct {
	UniqueID string                `json:"unique_id"`
	Name     string                `json:"name,omitempty"`
	Services []WorkflowTeamService `json:"services"`
}

// Workflow is the read shape of a workflow.
type Workflow struct {
	UniqueID     string         `json:"unique_id,omitempty"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	IsDraft      bool           `json:"is_draft"`
	IsEnabled    bool           `json:"is_enabled"`
	JSONSchema   string         `json:"json_schema"`
	Owner        interface{}    `json:"owner,omitempty"`
	Teams        []WorkflowTeam `json:"teams"`
	CreationDate string         `json:"creation_date,omitempty"`
	UpdatedDate  string         `json:"updated_date,omitempty"`
}

// WorkflowTeamInput is the write shape of one team in a workflow's scope.
type WorkflowTeamInput struct {
	UniqueID string   `json:"unique_id"`
	Services []string `json:"services"`
}

// WorkflowInput is the write shape; Teams replaces the scope, so nil is sent as [].
type WorkflowInput struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	IsDraft     bool                `json:"is_draft"`
	IsEnabled   bool                `json:"is_enabled"`
	JSONSchema  string              `json:"json_schema"`
	Teams       []WorkflowTeamInput `json:"teams"`
}

func (w WorkflowInput) MarshalJSON() ([]byte, error) {
	type alias WorkflowInput
	a := alias(w)
	if a.Teams == nil {
		a.Teams = []WorkflowTeamInput{}
	}
	for i := range a.Teams {
		if a.Teams[i].Services == nil {
			a.Teams[i].Services = []string{}
		}
	}
	return json.Marshal(a)
}

const workflowsPath = "/api/workflows/"

func (c *WorkflowService) CreateWorkflow(input *WorkflowInput) (*Workflow, error) {
	body, err := c.client.newRequestDo("POST", workflowsPath, input)
	if err != nil {
		return nil, err
	}
	var w Workflow
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

// GetWorkflows follows pagination; the list shape has no json_schema or services.
func (c *WorkflowService) GetWorkflows() ([]Workflow, error) {
	var all []Workflow
	err := c.client.getAll(workflowsPath, func(results []byte) error {
		var page []Workflow
		if err := json.Unmarshal(results, &page); err != nil {
			return err
		}
		all = append(all, page...)
		return nil
	})
	return all, err
}

func (c *WorkflowService) GetWorkflow(id string) (*Workflow, error) {
	path := fmt.Sprintf("%s%s/", workflowsPath, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var w Workflow
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *WorkflowService) UpdateWorkflow(id string, input *WorkflowInput) (*Workflow, error) {
	path := fmt.Sprintf("%s%s/", workflowsPath, id)
	body, err := c.client.newRequestDo("PATCH", path, input)
	if err != nil {
		return nil, err
	}
	var w Workflow
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *WorkflowService) DeleteWorkflow(id string) error {
	path := fmt.Sprintf("%s%s/", workflowsPath, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
