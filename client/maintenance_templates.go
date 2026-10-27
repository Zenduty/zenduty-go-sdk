package client

import (
	"encoding/json"
	"fmt"
)

type MaintenanceTemplateService service

// MaintenanceTemplate is a team-level rule (react-query-builder tree in
// RuleJSON, the same format alert rules use) that suppresses matching alerts
// while a maintenance window referencing it is active.
type MaintenanceTemplate struct {
	UniqueID  string `json:"unique_id,omitempty"`
	Name      string `json:"name"`
	IsEnabled bool   `json:"is_enabled"`
	RuleJSON  string `json:"rule_json"`
}

func maintenanceTemplatesPath(team string) string {
	return fmt.Sprintf("/api/v2/account/teams/%s/maintenance_template/", team)
}

func (c *MaintenanceTemplateService) CreateMaintenanceTemplate(team string, template *MaintenanceTemplate) (*MaintenanceTemplate, error) {
	body, err := c.client.newRequestDo("POST", maintenanceTemplatesPath(team), template)
	if err != nil {
		return nil, err
	}
	var t MaintenanceTemplate
	if err := json.Unmarshal(body.BodyBytes, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *MaintenanceTemplateService) GetMaintenanceTemplates(team string) ([]MaintenanceTemplate, error) {
	var all []MaintenanceTemplate
	err := c.client.getAll(maintenanceTemplatesPath(team), func(results []byte) error {
		var page []MaintenanceTemplate
		if err := json.Unmarshal(results, &page); err != nil {
			return err
		}
		all = append(all, page...)
		return nil
	})
	return all, err
}

func (c *MaintenanceTemplateService) GetMaintenanceTemplate(team, id string) (*MaintenanceTemplate, error) {
	path := maintenanceTemplatesPath(team) + id + "/"
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var t MaintenanceTemplate
	if err := json.Unmarshal(body.BodyBytes, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *MaintenanceTemplateService) UpdateMaintenanceTemplate(team, id string, template *MaintenanceTemplate) (*MaintenanceTemplate, error) {
	path := maintenanceTemplatesPath(team) + id + "/"
	body, err := c.client.newRequestDo("PATCH", path, template)
	if err != nil {
		return nil, err
	}
	var t MaintenanceTemplate
	if err := json.Unmarshal(body.BodyBytes, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *MaintenanceTemplateService) DeleteMaintenanceTemplate(team, id string) error {
	path := maintenanceTemplatesPath(team) + id + "/"
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
