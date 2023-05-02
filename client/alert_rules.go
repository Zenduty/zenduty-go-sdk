package client

import (
	"encoding/json"
	"fmt"
)

type AlertRuleService service

type AlertAction struct {
	UniqueID         string `json:"unique_id,omitempty"`
	ActionType       int    `json:"action_type"`
	Key              string `json:"key"`
	Value            string `json:"value"`
	AssignedTo       string `json:"assign_to"`
	EscalationPolicy string `json:"escalation_policy"`
	Schedule         string `json:"schedule"`
	TeamPriority     string `json:"team_priority"`
	SLA              string `json:"sla"`
}

type AlertRule struct {
	UniqueID    string        `json:"unique_id"`
	Description string        `json:"description"`
	Position    int           `json:"position,omitempty"`
	Stop        bool          `json:"stop,omitempty"`
	RuleType    int           `json:"ruleType,omitempty"`
	RuleJSON    string        `json:"rule_json"`
	Conditions  []string      `json:"conditions,omitempty"`
	Actions     []AlertAction `json:"actions,omitempty"`
}

func (c *AlertRuleService) CreateAlertRule(teamID, serviceID, integrationID string, rule *AlertRule) (*AlertRule, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/transformers/", teamID, serviceID, integrationID)
	body, err := c.client.newRequestDo("POST", path, rule)
	if err != nil {
		return nil, err
	}
	var s AlertRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *AlertRuleService) GetAlertRules(teamID, serviceID, integrationID string) ([]AlertRule, error) {

	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/transformers/", teamID, serviceID, integrationID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []AlertRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *AlertRuleService) GetAlertRule(teamID, serviceID, integrationID, id string) (*AlertRule, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/transformers/%s/", teamID, serviceID, integrationID, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s AlertRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *AlertRuleService) UpdateAlertRule(teamID, serviceID, integrationID, id string, rule *AlertRule) (*AlertRule, error) {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/transformers/%s/", teamID, serviceID, integrationID, id)
	body, err := c.client.newRequestDo("PUT", path, rule)
	if err != nil {
		return nil, err
	}
	var s AlertRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *AlertRuleService) DeleteAlertRule(teamID, serviceID, integrationID, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/services/%s/integrations/%s/transformers/%s/", teamID, serviceID, integrationID, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
