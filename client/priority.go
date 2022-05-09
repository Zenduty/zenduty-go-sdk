package client

import (
	"encoding/json"
	"fmt"
)

type PriorityService service

type Priority struct {
	Unique_Id     string `json:"unique_id,omitempty"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Color         string `json:"color"`
	Team          string `json:"team,omitempty"`
	Creation_Date string `json:"creation_date,omitempty"`
}

func (c *PriorityService) CreatePriority(team string, priority *Priority) (*Priority, error) {
	path := fmt.Sprintf("/api/account/teams/%s/priority/", team)
	body, err := c.client.newRequestDo("POST", path, priority)
	if err != nil {
		return nil, err
	}
	var s Priority
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *PriorityService) GetPriority(team string) ([]Priority, error) {
	path := fmt.Sprintf("/api/account/teams/%s/priority/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []Priority
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *PriorityService) GetPriorityById(team, id string) (*Priority, error) {
	path := fmt.Sprintf("/api/account/teams/%s/priority/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s Priority
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *PriorityService) UpdatePriority(team, id string, priority *Priority) (*Priority, error) {
	path := fmt.Sprintf("/api/account/teams/%s/priority/%s/", team, id)
	body, err := c.client.newRequestDo("PUT", path, priority)
	if err != nil {
		return nil, err
	}
	var s Priority
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *PriorityService) DeletePriority(team, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/priority/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
