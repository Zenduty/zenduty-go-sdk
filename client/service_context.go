package client

import (
	"encoding/json"
	"fmt"
)

type ServiceContextService service

// ServiceContext links a context integration to a service. A service can
// link each integration once; there is no update, only create and delete.
type ServiceContext struct {
	UniqueID     string `json:"unique_id,omitempty"`
	Service      string `json:"service,omitempty"`
	Integration  string `json:"integration"`
	CreatedBy    string `json:"created_by,omitempty"`
	CreationDate string `json:"creation_date,omitempty"`
}

func serviceContextsPath(team, service string) string {
	return fmt.Sprintf("/api/account/teams/%s/services/%s/context_integrations/", team, service)
}

func (c *ServiceContextService) CreateServiceContext(team, service string, context *ServiceContext) (*ServiceContext, error) {
	body, err := c.client.newRequestDo("POST", serviceContextsPath(team, service), context)
	if err != nil {
		return nil, err
	}
	var s ServiceContext
	if err := json.Unmarshal(body.BodyBytes, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *ServiceContextService) GetServiceContexts(team, service string) ([]ServiceContext, error) {
	var all []ServiceContext
	err := c.client.getAll(serviceContextsPath(team, service), func(results []byte) error {
		var page []ServiceContext
		if err := json.Unmarshal(results, &page); err != nil {
			return err
		}
		all = append(all, page...)
		return nil
	})
	return all, err
}

func (c *ServiceContextService) GetServiceContext(team, service, id string) (*ServiceContext, error) {
	path := serviceContextsPath(team, service) + id + "/"
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s ServiceContext
	if err := json.Unmarshal(body.BodyBytes, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *ServiceContextService) DeleteServiceContext(team, service, id string) error {
	path := serviceContextsPath(team, service) + id + "/"
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
