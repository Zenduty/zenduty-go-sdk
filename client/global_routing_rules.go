package client

import (
	"encoding/json"
	"fmt"
)

type GlobalRoutingRuleIntegrationObject struct {
	Name     string `json:"name"`
	UniqueID string `json:"unique_id"`
}
type GlobalRoutingRuleAction struct {
	UniqueID          string                             `json:"unique_id,omitempty"`
	ActionType        int                                `json:"action_type"`
	Integration       string                             `json:"integration,omitempty"`
	IntegrationObject GlobalRoutingRuleIntegrationObject `json:"integration_obj,omitempty"`
}

// action := client.GlobalRoutingRuleAction{ ActionType:  1}

type GlobalRoutingRule struct {
	UniqueID string                    `json:"unique_id,omitempty"`
	Name     string                    `json:"name"`
	Position int                       `json:"position,omitempty"`
	RuleJSON string                    `json:"rule_json"`
	Actions  []GlobalRoutingRuleAction `json:"actions"` // []client.GlobalRoutingRuleAction{action}
}

func (c *GlobalRouterService) CreateGlobalRoutingRule(routerID string, rule *GlobalRoutingRule) (*GlobalRoutingRule, error) {
	path := fmt.Sprintf("/api/v2/account/events/router/%s/rulesets/", routerID)
	body, err := c.client.newRequestDo("POST", path, rule)
	if err != nil {
		return nil, err
	}
	var s GlobalRoutingRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *GlobalRouterService) GetGlobalRoutingRules(routerID string) ([]GlobalRoutingRule, error) {

	path := fmt.Sprintf("/api/v2/account/events/router/%s/rulesets/", routerID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []GlobalRoutingRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *GlobalRouterService) GetGlobalRoutingRule(routerID, rulesetID string) (*GlobalRoutingRule, error) {
	path := fmt.Sprintf("/api/v2/account/events/router/%s/rulesets/%s/", routerID, rulesetID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s GlobalRoutingRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *GlobalRouterService) UpdateGlobalRoutingRule(routerID, rulesetID string, rule *GlobalRoutingRule) (*GlobalRoutingRule, error) {
	path := fmt.Sprintf("/api/v2/account/events/router/%s/rulesets/%s/", routerID, rulesetID)
	body, err := c.client.newRequestDo("PUT", path, rule)
	if err != nil {
		return nil, err
	}
	var s GlobalRoutingRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *GlobalRouterService) DeleteGlobalRoutingRule(routerID, rulesetID string) error {
	path := fmt.Sprintf("/api/v2/account/events/router/%s/rulesets/%s/", routerID, rulesetID)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
