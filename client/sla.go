package client

import (
	"encoding/json"
	"fmt"
)

type SLAService service

type ResponderUser struct {
	User     string `json:"user"`
	UniqueID string `json:"unique_id,omitempty"`
}

type SLAEscalations struct {
	UniqueID   string          `json:"unique_id,omitempty"`
	Responders []ResponderUser `json:"responders"`
	Time       int             `json:"time"`
	Type       int             `json:"type"`
}

type SLAObj struct {
	UniqueID        string           `json:"unique_id,omitempty"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	AcknowledgeTime int              `json:"acknowledge_time"`
	IsActive        bool             `json:"is_active"`
	ResolveTime     int              `json:"resolve_time,omitempty"`
	Escalations     []SLAEscalations `json:"escalations"`
}

type SLAs struct {
	UniqueID        string `json:"unique_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	AcknowledgeTime int    `json:"acknowledge_time"`
	IsActive        bool   `json:"is_active"`
	ResolveTime     int    `json:"resolve_time"`
}

func (c *SLAService) CreateSLA(team string, sla *SLAObj) (*SLAObj, error) {

	path := fmt.Sprintf("/api/account/teams/%s/sla/", team)
	body, err := c.client.newRequestDo("POST", path, sla)
	if err != nil {
		return nil, err
	}
	var s SLAObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *SLAService) GetSLAs(team string) ([]SLAs, error) {
	path := fmt.Sprintf("/api/account/teams/%s/sla/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []SLAs
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *SLAService) GetSLAByID(team, id string) (*SLAObj, error) {
	path := fmt.Sprintf("/api/account/teams/%s/sla/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s SLAObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *SLAService) DeleteSLAByID(team, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/sla/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *SLAService) UpdateSLAByID(team, id string, SLA *SLAObj) (*SLAObj, error) {

	path := fmt.Sprintf("/api/account/teams/%s/sla/%s/", team, id)
	body, err := c.client.newRequestDo("PATCH", path, SLA)
	if err != nil {
		return nil, err
	}
	var s SLAObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
