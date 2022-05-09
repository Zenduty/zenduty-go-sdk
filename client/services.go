package client

import (
	"encoding/json"
	"fmt"
)

type Service service

type Services struct {
	Name                   string `json:"name"`
	Creation_Date          string `json:"creation_date"`
	Summary                string `json:"summary"`
	Description            string `json:"description"`
	Unique_Id              string `json:"unique_id",omitempty`
	Auto_Resolve_Timeout   int    `json:"auto_resolve_timeout",omitempty`
	Created_By             string `json:"created_by",omitempty`
	Team_Priority          string `json:"team_priority"`
	Task_Template          string `json:"task_template"`
	Acknowledgment_Timeout int    `json:"acknowledge_timeout",omitempty`
	Status                 int    `json:"status",omitempty`
	Escalation_Policy      string `json:"escalation_policy"`
	Team                   string `json:"team"`
	Sla                    string `json:"sla"`
	Collation_Time         int    `json:"collation_time"`
	Collation              int    `json:"collation"`
	Under_Maintenance      bool   `json:"under_maintenance",omitempty`
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

func (c *Service) GetServicesById(team, id string) (*Services, error) {
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
