package client

import (
	"encoding/json"
	"fmt"
)

type AlertGroupingPolicyService service

// AlertGroupingMatchFields lists the payload fields a content-based
// collation policy compares across alerts. StaticFields holds the built-in
// options (incident_title, incident_summary, integration); CustomFields
// holds arbitrary JSON keys from the integration payload.
type AlertGroupingMatchFields struct {
	StaticFields []string `json:"static_fields"`
	CustomFields []string `json:"custom_fields"`
}

// AlertGroupingPolicy configures content-based alert collation for a
// service (backend endpoint: alert-grouping-policies). MatchMode selects
// whether any or all of MatchFields must match for alerts to be grouped.
type AlertGroupingPolicy struct {
	UniqueID  string `json:"unique_id,omitempty"`
	Service   string `json:"service,omitempty"`
	MatchMode int    `json:"match_mode"`
	// TimeWindow only takes effect while IsActive is true.
	MatchFields AlertGroupingMatchFields `json:"match_fields"`
	TimeWindow  int                      `json:"time_window"`
	IsActive    bool                     `json:"is_active"`
	CreatedAt   string                   `json:"created_at,omitempty"`
	UpdatedAt   string                   `json:"updated_at,omitempty"`
}

func (c *AlertGroupingPolicyService) CreateAlertGroupingPolicy(teamID, serviceID string, policy *AlertGroupingPolicy) (*AlertGroupingPolicy, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/alert-grouping-policies/", teamID, serviceID)
	body, err := c.client.newRequestDo("POST", path, policy)
	if err != nil {
		return nil, err
	}
	var p AlertGroupingPolicy
	if err := json.Unmarshal(body.BodyBytes, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *AlertGroupingPolicyService) GetAlertGroupingPolicies(teamID, serviceID string) ([]AlertGroupingPolicy, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/alert-grouping-policies/", teamID, serviceID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var p []AlertGroupingPolicy
	if err := json.Unmarshal(body.BodyBytes, &p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *AlertGroupingPolicyService) GetAlertGroupingPolicy(teamID, serviceID, id string) (*AlertGroupingPolicy, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/alert-grouping-policies/%s/", teamID, serviceID, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var p AlertGroupingPolicy
	if err := json.Unmarshal(body.BodyBytes, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *AlertGroupingPolicyService) UpdateAlertGroupingPolicy(teamID, serviceID, id string, policy *AlertGroupingPolicy) (*AlertGroupingPolicy, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/alert-grouping-policies/%s/", teamID, serviceID, id)
	body, err := c.client.newRequestDo("PATCH", path, policy)
	if err != nil {
		return nil, err
	}
	var p AlertGroupingPolicy
	if err := json.Unmarshal(body.BodyBytes, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *AlertGroupingPolicyService) DeleteAlertGroupingPolicy(teamID, serviceID, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/alert-grouping-policies/%s/", teamID, serviceID, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
