package client

import (
	"encoding/json"
)

type ApplicationsService service

// CatalogApplication is an entry in the account's integration application
// catalog. Its unique_id differs between Zenduty instances, so configs should
// look applications up by name rather than hardcode the id.
type CatalogApplication struct {
	UniqueID        string `json:"unique_id"`
	Name            string `json:"name"`
	Summary         string `json:"summary"`
	ApplicationType int    `json:"application_type"`
}

func (s *ApplicationsService) GetApplications() ([]CatalogApplication, error) {
	body, err := s.client.newRequestDo("GET", "/api/account/applications/", nil)
	if err != nil {
		return nil, err
	}
	var apps []CatalogApplication
	if err := json.Unmarshal(body.BodyBytes, &apps); err != nil {
		return nil, err
	}
	return apps, nil
}
