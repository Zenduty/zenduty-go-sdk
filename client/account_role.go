package client

import (
	"encoding/json"
	"fmt"
)

type AccountRoleService service

type AccountRole struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	UniqueID    string   `json:"unique_id,omitempty"`
}

type AddRoleToUser struct {
	AccountRole *string `json:"custom_role"`
	Username    string  `json:"user,omitempty"`
}

func (c *AccountRoleService) CreateAccountRole(role *AccountRole) (*AccountRole, error) {
	path := "/api/account/customroles/"
	body, err := c.client.newRequestDo("POST", path, role)
	if err != nil {
		return nil, err
	}
	var r AccountRole
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *AccountRoleService) UpdateAccountRole(UniqueID string, role *AccountRole) (*AccountRole, error) {
	path := fmt.Sprintf("/api/account/customroles/%s/", UniqueID)
	body, err := c.client.newRequestDo("PUT", path, role)
	if err != nil {
		return nil, err
	}
	var r AccountRole
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
func (c *AccountRoleService) DeleteAccountRole(id string) error {
	path := fmt.Sprintf("/api/account/customroles/%s/", id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *AccountRoleService) GetAccountRoleByID(id string) (*AccountRole, error) {
	path := fmt.Sprintf("/api/account/customroles/%s/", id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var r AccountRole
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *AccountRoleService) GetAccountRoles() ([]AccountRole, error) {
	path := "/api/account/customroles/"
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var r []AccountRole
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *AccountRoleService) AssignRoleToUser(username string, role *AddRoleToUser) (*AddRoleToUser, error) {
	path := fmt.Sprintf("/api/account/users/%s/customroles/", username)
	body, err := c.client.newRequestDo("POST", path, role)
	if err != nil {
		return nil, err
	}
	var a AddRoleToUser
	err = json.Unmarshal(body.BodyBytes, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
