package client

import (
	"encoding/json"
	"fmt"
)

type IntegrationServerice service

type ApplicationReference struct {
	Name                string `json:"name"`
	Icon_Url            string `json:"icon_url"`
	Summary             string `json:"summary"`
	Description         string `json:"description"`
	Unique_Id           string `json:"unique_id"`
	Avalability_Plan_id int    `json:"availability_plan_id"`
	Setup_Instructions  string `json:"setup_instructions"`
	Extension           string `json:"extension"`
	Application_Type    int    `json:"application_type"`
	Categories          string `json:"categories"`
	Documentation_Link  string `json:"documentation_link"`
}

type IntegrationCreate struct {
	Name                string `json:"name"`
	Summary             string `json:"summary"`
	Application         string `json:"application"`
	Is_Enabled          bool   `json:"is_enabled"`
	Create_Incident_For int    `json:"create_incidents_for"`
	Default_Urgency     int    `json:"default_urgency"`
}

type Integration struct {
	Name                  string               `json:"name"`
	Creation_Date         string               `json:"creation_date"`
	Summary               string               `json:"summary"`
	Description           string               `json:"description"`
	Unique_Id             string               `json:"unique_id"`
	Service               string               `json:"service"`
	Application           string               `json:"application"`
	Application_Reference ApplicationReference `json:"application_reference"`
	Integration_key       string               `json:"integration_key"`
	Webhook_url           string               `json:"webhook_url"`
	Created_By            string               `json:"created_by"`
	Is_Enabled            bool                 `json:"is_enabled"`
	Create_Incident_For   int                  `json:"create_incidents_for"`
	Integration_Type      int                  `json:"integration_type"`
	Default_Urgency       int                  `json:"default_urgency"`
}

func (c *IntegrationServerice) CreateIntegration(team string, service_id string, integration *IntegrationCreate) (*Integration, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/", team, service_id)

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

func (c *IntegrationServerice) UpdateIntegration(team string, service_id string, integration_id string, integration *IntegrationCreate) (*Integration, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/", team, service_id, integration_id)

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

func (c *IntegrationServerice) GetIntegrations(team, service_id string) ([]Integration, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/", team, service_id)
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

func (c *IntegrationServerice) GetIntegrationByID(team, service_id, id string) (*Integration, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/", team, service_id, id)
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

func (c *IntegrationServerice) DeleteIntegration(team, service_id, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/", team, service_id, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
