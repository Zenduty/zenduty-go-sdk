package client

import (
	"encoding/json"
	"fmt"
)

type ForwardingRuleService service

// ForwardingRule forwards a user's notifications for a period; no update exists.
// A user's list also contains rules forwarded to them (see User).
type ForwardingRule struct {
	UniqueID     string `json:"unique_id,omitempty"`
	User         string `json:"user,omitempty"`
	ToUser       string `json:"to_user"`
	FromDate     string `json:"from_date"`
	ToDate       string `json:"to_date"`
	TimeZone     string `json:"time_zone,omitempty"`
	CreatedBy    string `json:"created_by,omitempty"`
	CreationDate string `json:"creation_date,omitempty"`
}

func forwardingRulesPath(username string) string {
	return fmt.Sprintf("/api/account/users/%s/forwarding_rules/", username)
}

func (c *ForwardingRuleService) CreateForwardingRule(username string, rule *ForwardingRule) (*ForwardingRule, error) {
	body, err := c.client.newRequestDo("POST", forwardingRulesPath(username), rule)
	if err != nil {
		return nil, err
	}
	var r ForwardingRule
	if err := json.Unmarshal(body.BodyBytes, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *ForwardingRuleService) GetForwardingRules(username string) ([]ForwardingRule, error) {
	var all []ForwardingRule
	err := c.client.getAll(forwardingRulesPath(username), func(results []byte) error {
		var page []ForwardingRule
		if err := json.Unmarshal(results, &page); err != nil {
			return err
		}
		all = append(all, page...)
		return nil
	})
	return all, err
}

func (c *ForwardingRuleService) GetForwardingRule(username, id string) (*ForwardingRule, error) {
	path := forwardingRulesPath(username) + id + "/"
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var r ForwardingRule
	if err := json.Unmarshal(body.BodyBytes, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *ForwardingRuleService) DeleteForwardingRule(username, id string) error {
	path := forwardingRulesPath(username) + id + "/"
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
