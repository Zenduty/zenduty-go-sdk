package client

import (
	"encoding/json"
	"fmt"
)

type Service service

type Services struct {
	Name         string `json:"name"`
	CreationDate string `json:"creation_date"`
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	UniqueID     string `json:"unique_id,omitempty"`
	// 0 disables auto-resolution; sent unconditionally so it can be reset.
	AutoResolveTimeout int    `json:"auto_resolve_timeout"`
	CreatedBy          string `json:"created_by,omitempty"`
	TeamPriority       string `json:"team_priority"`
	TaskTemplate       string `json:"task_template"`
	// The backend field is acknowledgement_timeout (not acknowledge_timeout)
	// and only takes effect while AcknowledgementTimeoutEnabled is true.
	AcknowledgmentTimeout         int    `json:"acknowledgement_timeout"`
	AcknowledgementTimeoutEnabled bool   `json:"acknowledgement_timeout_enabled"`
	Status                        int    `json:"status,omitempty"`
	EscalationPolicy              string `json:"escalation_policy"`
	Team                          string `json:"team"`
	SLA                           string `json:"sla"`
	CollationTime                 int    `json:"collation_time"`
	Collation                     int    `json:"collation"`
	UnderMaintenance              bool   `json:"under_maintenance,omitempty"`
}

func (c *Service) CreateService(team string, service *Services) (*Services, error) {

	path := fmt.Sprintf("/api/account/teams/%s/services/", team)
	body, err := c.client.newRequestDo("POST", path, service)
	if err != nil {
		return nil, err
	}
	var i Services
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Service) GetServices(team string) ([]Services, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/", team)

	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var i []Services
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (c *Service) GetServicesByID(team, id string) (*Services, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var i Services
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Service) UpdateService(team, id string, service *Services) (*Services, error) {

	path := fmt.Sprintf("/api/account/teams/%s/services/%s/", team, id)
	body, err := c.client.newRequestDo("PATCH", path, service)
	if err != nil {
		return nil, err
	}
	var i Services
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Service) DeleteService(team string, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
