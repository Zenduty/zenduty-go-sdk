package client

import (
	"encoding/json"
	"fmt"
)

type EspService service
type Targets struct {
	TargetType int    `json:"target_type"`
	TargetID   string `json:"target_id"`
}
type Rules struct {
	Delay    int       `json:"delay"`
	Targets  []Targets `json:"targets"`
	Position int       `json:"position"`
	UniqueID string    `json:"unique_id"`
}

type EscalationPolicy struct {
	Name         string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	Summary      string  `json:"summary"`
	Team         string  `json:"team"`
	UniqueID     string  `json:"unique_id"`
	RepeatPolicy int     `json:"repeat_policy"`
	MoveToNext   bool    `json:"move_to_next"`
	GlobalEp     bool    `json:"global_ep"`
	Rules        []Rules `json:"rules"`
}

func (c *EspService) CreateEscalationPolicy(team string, policy *EscalationPolicy) (*EscalationPolicy, error) {
	path := fmt.Sprintf("/api/account/teams/%s/escalation_policies/", team)
	body, err := c.client.newRequestDo("POST", path, policy)
	if err != nil {
		return nil, err
	}
	var s EscalationPolicy
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *EspService) GetEscalationPolicy(team string) ([]EscalationPolicy, error) {
	path := fmt.Sprintf("/api/account/teams/%s/escalation_policies/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []EscalationPolicy
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *EspService) GetEscalationPolicyByID(team, id string) (*EscalationPolicy, error) {
	path := fmt.Sprintf("/api/account/teams/%s/escalation_policies/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s EscalationPolicy
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *EspService) DeleteEscalationPolicy(team, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/escalation_policies/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *EspService) UpdateEscalationPolicy(team, id string, policy *EscalationPolicy) (*EscalationPolicy, error) {
	path := fmt.Sprintf("/api/account/teams/%s/escalation_policies/%s/", team, id)
	body, err := c.client.newRequestDo("PUT", path, policy)
	if err != nil {
		return nil, err
	}
	var s EscalationPolicy
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
