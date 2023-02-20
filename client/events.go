package client

import (
	"encoding/json"
	"fmt"
)

type EventsService service

type AlertPayload struct {
	Message    string      `json:"message"`
	Summary    string      `json:"summary"`
	AlertType  string      `json:"alert_type"`
	Suppressed bool        `json:"suppressed,omitempty"`
	EntityID   string      `json:"entity_id"`
	Payload    interface{} `json:"payload"` // payload map[string]interface{}{"severity": "SEV1"}
}

type IntegrationObject struct {
	Name           string `json:"name"`
	Summary        string `json:"summary"`
	Service        string `json:"service"`
	Team           string `json:"team"`
	IntegrationKey string `json:"integration_key"`
	IsEnabled      bool   `json:"is_enabled"`
}

type AlertResponse struct {
	UniqueID          string            `json:"unique_id"`
	Message           string            `json:"message"`
	Summary           string            `json:"summary"`
	AlertType         string            `json:"alert_type"`
	Suppressed        bool              `json:"suppressed"`
	EntityID          string            `json:"entity_id"`
	Payload           string            `json:"payload"`
	IntegrationObject IntegrationObject `json:"integration_object"`
	Integration       string            `json:"integration"`
	CreationDate      string            `json:"creation_date"`
}

func (c *EventsService) SendAlert(integrationKey string, payload *AlertPayload) (*AlertResponse, error) {
	path := fmt.Sprintf("/api/events/%s/", integrationKey)

	body, err := c.client.newRequestDo("POST", path, payload)
	if err != nil {
		return nil, err
	}
	var i AlertResponse
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
