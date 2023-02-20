package client

import (
	"encoding/json"
	"fmt"
)

type GlobalRouterService service

type GlobalRouterPayload struct {
	UniqueID       string `json:"unique_id,omitempty"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	IsEnabled      bool   `json:"is_enabled,omitempty"`
	IntegrationKey string `json:"integration_key,omitempty"`
}

func (c *GlobalRouterService) CreateGlobalRouter(router *GlobalRouterPayload) (*GlobalRouterPayload, error) {
	path := "/api/v2/account/events/router/"

	body, err := c.client.newRequestDo("POST", path, router)
	if err != nil {
		return nil, err
	}
	var r GlobalRouterPayload
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *GlobalRouterService) UpdateGlobalRouter(id string, router *GlobalRouterPayload) (*GlobalRouterPayload, error) {
	path := fmt.Sprintf("/api/v2/account/events/router/%s/", id)

	body, err := c.client.newRequestDo("PATCH", path, router)
	if err != nil {
		return nil, err
	}
	var r GlobalRouterPayload
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *GlobalRouterService) GetGlobalRouters() ([]GlobalRouterPayload, error) {
	path := "/api/v2/account/events/router/"
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var i []GlobalRouterPayload
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (c *GlobalRouterService) GetGlobalRouter(id string) (*GlobalRouterPayload, error) {
	path := fmt.Sprintf("/api/v2/account/events/router/%s/", id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var r GlobalRouterPayload
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *GlobalRouterService) DeleteGlobalRouter(id string) error {
	path := fmt.Sprintf("/api/v2/account/events/router/%s/", id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
