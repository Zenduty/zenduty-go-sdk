package client

import (
	"encoding/json"
	"fmt"
)

type OutgoingRulesService service

type OutgoingRule struct {
	UniqueID string `json:"unique_id"`
	Enabled  bool   `json:"is_enabled"`
	RuleJSON string `json:"rule_json"`
}

func (c *OutgoingRulesService) CreateOutgoingRule(teamID, serviceID, integrationID string, rule *OutgoingRule) (*OutgoingRule, error) {
	path := fmt.Sprintf("/api/v2/account/teams/%s/services/%s/integrations/%s/outgoingrules/", teamID, serviceID, integrationID)
	body, err := c.client.newRequestDo("POST", path, rule)
	if err != nil {
		return nil, err
	}
	var s OutgoingRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *OutgoingRulesService) GetOutgoingRules(teamID, serviceID, integrationID string) ([]OutgoingRule, error) {

	path := fmt.Sprintf("/api/v2/account/teams/%s/services/%s/integrations/%s/outgoingrules/", teamID, serviceID, integrationID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []OutgoingRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *OutgoingRulesService) GetOutgoingRule(teamID, serviceID, integrationID, id string) (*OutgoingRule, error) {
	path := fmt.Sprintf("/api/v2/account/teams/%s/services/%s/integrations/%s/outgoingrules/%s/", teamID, serviceID, integrationID, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s OutgoingRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *OutgoingRulesService) UpdateOutgoingRule(teamID, serviceID, integrationID, id string, rule *OutgoingRule) (*OutgoingRule, error) {
	path := fmt.Sprintf("/api/v2/account/teams/%s/services/%s/integrations/%s/outgoingrules/%s/", teamID, serviceID, integrationID, id)
	body, err := c.client.newRequestDo("PUT", path, rule)
	if err != nil {
		return nil, err
	}
	var s OutgoingRule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *OutgoingRulesService) DeleteOutgoingRule(teamID, serviceID, integrationID, id string) error {
	path := fmt.Sprintf("/api/v2/account/teams/%s/services/%s/integrations/%s/outgoingrules/%s/", teamID, serviceID, integrationID, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
