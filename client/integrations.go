package client

import (
	"encoding/json"
	"fmt"
)

type IntegrationServerice service

type ApplicationReference struct {
	Name              string `json:"name"`
	IconURL           string `json:"icon_url"`
	Summary           string `json:"summary"`
	Description       string `json:"description"`
	UniqueID          string `json:"unique_id"`
	AvalabilityPlanID int    `json:"availability_plan_id"`
	SetupInstructions string `json:"setup_instructions"`
	Extension         string `json:"extension"`
	ApplicationType   int    `json:"application_type"`
	Categories        string `json:"categories"`
	DocumentationLink string `json:"documentation_link"`
}

type IntegrationCreate struct {
	Name              string `json:"name"`
	Summary           string `json:"summary"`
	Application       string `json:"application"`
	IsEnabled         bool   `json:"is_enabled"`
	CreateIncidentFor int    `json:"create_incidents_for"`
	DefaultUrgency    int    `json:"default_urgency"`
}

type Integration struct {
	Name                 string               `json:"name"`
	CreationDate         string               `json:"creation_date"`
	Summary              string               `json:"summary"`
	Description          string               `json:"description"`
	UniqueID             string               `json:"unique_id"`
	Service              string               `json:"service"`
	Application          string               `json:"application"`
	ApplicationReference ApplicationReference `json:"application_reference"`
	IntegrationKey       string               `json:"integration_key"`
	WebhookURL           string               `json:"webhook_url"`
	CreatedBy            string               `json:"created_by"`
	IsEnabled            bool                 `json:"is_enabled"`
	CreateIncidentFor    int                  `json:"create_incidents_for"`
	IntegrationType      int                  `json:"integration_type"`
	DefaultUrgency       int                  `json:"default_urgency"`
}

func (c *IntegrationServerice) CreateIntegration(team string, serviceID string, integration *IntegrationCreate) (*Integration, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/", team, serviceID)

	body, err := c.client.newRequestDo("POST", path, integration)
	if err != nil {
		return nil, err
	}
	var i Integration
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IntegrationServerice) UpdateIntegration(team string, serviceID string, integrationID string, integration *IntegrationCreate) (*Integration, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/", team, serviceID, integrationID)

	body, err := c.client.newRequestDo("PATCH", path, integration)
	if err != nil {
		return nil, err
	}
	var i Integration
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IntegrationServerice) GetIntegrations(team, serviceID string) ([]Integration, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/", team, serviceID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var i []Integration
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (c *IntegrationServerice) GetIntegrationByID(team, serviceID, id string) (*Integration, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/", team, serviceID, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var i Integration
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IntegrationServerice) DeleteIntegration(team, serviceID, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/", team, serviceID, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
