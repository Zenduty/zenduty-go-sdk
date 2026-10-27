package client

import (
	"encoding/json"
	"fmt"
)

type HeartbeatService service

// Heartbeat ping types (HeartbeatPingType on the backend).
const (
	HeartbeatPingWebhook = 0
	HeartbeatPingEmail   = 1
)

// Heartbeat interval limits enforced by the API, in seconds.
const (
	HeartbeatMinIntervalSeconds = 60
	HeartbeatMaxIntervalSeconds = 90 * 24 * 3600
)

// Heartbeat is an account-level heartbeat monitor. Status is one of
// "disabled", "pending", "healthy" or "expired".
type Heartbeat struct {
	UniqueID             string `json:"unique_id,omitempty"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Service              string `json:"service,omitempty"`
	ServiceName          string `json:"service_name,omitempty"`
	Team                 string `json:"team,omitempty"`
	TeamName             string `json:"team_name,omitempty"`
	IntervalSeconds      int    `json:"interval_seconds"`
	Status               string `json:"status,omitempty"`
	IsEnabled            bool   `json:"is_enabled"`
	LastPingAt           string `json:"last_ping_at,omitempty"`
	ExpiresAt            string `json:"expires_at,omitempty"`
	OpenIncidentNumber   *int   `json:"open_incident_number,omitempty"`
	AlertMessage         string `json:"alert_message"`
	AlertUrgency         *int   `json:"alert_urgency"`
	PingType             int    `json:"ping_type"`
	EmailSubjectContains string `json:"email_subject_contains"`
	PingURL              string `json:"ping_url,omitempty"`
	PingEmail            string `json:"ping_email,omitempty"`
	CreatedBy            string `json:"created_by,omitempty"`
	CreationDate         string `json:"creation_date,omitempty"`
}

const heartbeatsPath = "/api/account/heartbeats/"

func (c *HeartbeatService) CreateHeartbeat(heartbeat *Heartbeat) (*Heartbeat, error) {
	body, err := c.client.newRequestDo("POST", heartbeatsPath, heartbeat)
	if err != nil {
		return nil, err
	}
	var h Heartbeat
	if err := json.Unmarshal(body.BodyBytes, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

// GetHeartbeats lists every heartbeat of the account.
func (c *HeartbeatService) GetHeartbeats() ([]Heartbeat, error) {
	var all []Heartbeat
	err := c.client.getAll(heartbeatsPath, func(results []byte) error {
		var page []Heartbeat
		if err := json.Unmarshal(results, &page); err != nil {
			return err
		}
		all = append(all, page...)
		return nil
	})
	return all, err
}

func (c *HeartbeatService) GetHeartbeat(id string) (*Heartbeat, error) {
	path := fmt.Sprintf("%s%s/", heartbeatsPath, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var h Heartbeat
	if err := json.Unmarshal(body.BodyBytes, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

// UpdateHeartbeat patches a heartbeat. The service cannot be changed after
// creation; the API ignores it on update.
func (c *HeartbeatService) UpdateHeartbeat(id string, heartbeat *Heartbeat) (*Heartbeat, error) {
	path := fmt.Sprintf("%s%s/", heartbeatsPath, id)
	body, err := c.client.newRequestDo("PATCH", path, heartbeat)
	if err != nil {
		return nil, err
	}
	var h Heartbeat
	if err := json.Unmarshal(body.BodyBytes, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

func (c *HeartbeatService) DeleteHeartbeat(id string) error {
	path := fmt.Sprintf("%s%s/", heartbeatsPath, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
