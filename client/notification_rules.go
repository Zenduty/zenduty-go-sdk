package client

import (
	"encoding/json"
	"fmt"
)

type NotificationRulesService service

type NotificationRules struct {
	UniqueID   string `json:"unique_id"`
	StartDelay int    `json:"start_delay"`
	Type       string `json:"type"`
	Contact    string `json:"contact"`
	Urgency    int    `json:"urgency"`
}

type CreateNotificationRules struct {
	Contact    string `json:"contact"`
	StartDelay int    `json:"start_delay"`
	Urgency    int    `json:"urgency"`
}

func (c *NotificationRulesService) CreateNotificationRules(username string, notificationRule *CreateNotificationRules) (*NotificationRules, error) {
	path := fmt.Sprintf("/api/account/users/%s/notification_rules/", username)
	body, err := c.client.newRequestDo("POST", path, notificationRule)
	if err != nil {
		return nil, err
	}
	var s NotificationRules
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *NotificationRulesService) GetNotificationRules(username string) ([]NotificationRules, error) {
	path := fmt.Sprintf("/api/account/users/%s/notification_rules/", username)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []NotificationRules
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *NotificationRulesService) GetNotificationRulesByID(username, notificationRuleID string) (*NotificationRules, error) {
	path := fmt.Sprintf("/api/account/users/%s/notification_rules/%s/", username, notificationRuleID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s NotificationRules
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *NotificationRulesService) DeleteNotificationRules(username string, notificationRuleID string) error {
	path := fmt.Sprintf("/api/account/users/%s/notification_rules/%s/", username, notificationRuleID)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *NotificationRulesService) UpdateNotificationRules(username string, notificationRuleID string, notificationRule *NotificationRules) (*NotificationRules, error) {
	path := fmt.Sprintf("/api/account/users/%s/notification_rules/%s/", username, notificationRuleID)
	body, err := c.client.newRequestDo("PUT", path, notificationRule)
	if err != nil {
		return nil, err
	}
	var s NotificationRules
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil

}
