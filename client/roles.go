package client

import (
	"encoding/json"
	"fmt"
)

type RoleService service

type Roles struct {
	Team          string `json:"team",omitempty`
	Unique_Id     string `json:"unique_id",omitempty`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Creation_Date string `json:"creation_date",omitempty`
	Rank          int    `json:"rank",omitempty`
}

func (c *RoleService) CreateRole(team string, role *Roles) (*Roles, error) {
	path := fmt.Sprintf("/api/account/teams/%s/roles/", team)

	body, err := c.client.newRequestDo("POST", path, role)
	if err != nil {
		return nil, err
	}
	var r Roles
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *RoleService) GetRoles(team string) ([]Roles, error) {
	path := fmt.Sprintf("/api/account/teams/%s/roles/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var r []Roles
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *RoleService) UpdateRoles(team string, role *Roles) (*Roles, error) {
	path := fmt.Sprintf("/api/account/teams/%s/roles/%s/", team, role.Unique_Id)

	body, err := c.client.newRequestDo("PATCH", path, role)
	if err != nil {
		return nil, err
	}
	var r Roles
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *RoleService) DeleteRole(team string, role string) error {
	path := fmt.Sprintf("/api/account/teams/%s/roles/%s/", team, role)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
