package client

import (
	"encoding/json"
	"fmt"
)

type WebhookConfigurationService service

// WebhookConfigurationEvents are the events an account webhook can subscribe to.
var WebhookConfigurationEvents = []string{"oncall_change", "incident_create", "incident_acknowledge", "incident_resolve"}

// WebhookConfiguration is an account-level outgoing webhook. Reads return
// masked secrets; an empty Filter or secret clears it on update.
type WebhookConfiguration struct {
	UniqueID      string   `json:"unique_id,omitempty"`
	URL           string   `json:"url"`
	Event         []string `json:"event"`
	AuthSecretKey string   `json:"auth_secret_key"`
	APIKey        string   `json:"api_key"`
	Filter        string   `json:"filter"`
}

const webhookConfigurationsPath = "/api/v2/account/webhookconfiguration/"

func (c *WebhookConfigurationService) CreateWebhookConfiguration(config *WebhookConfiguration) (*WebhookConfiguration, error) {
	body, err := c.client.newRequestDo("POST", webhookConfigurationsPath, config)
	if err != nil {
		return nil, err
	}
	var w WebhookConfiguration
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *WebhookConfigurationService) GetWebhookConfigurations() ([]WebhookConfiguration, error) {
	var all []WebhookConfiguration
	err := c.client.getAll(webhookConfigurationsPath, func(results []byte) error {
		var page []WebhookConfiguration
		if err := json.Unmarshal(results, &page); err != nil {
			return err
		}
		all = append(all, page...)
		return nil
	})
	return all, err
}

func (c *WebhookConfigurationService) GetWebhookConfiguration(id string) (*WebhookConfiguration, error) {
	path := fmt.Sprintf("%s%s/", webhookConfigurationsPath, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var w WebhookConfiguration
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *WebhookConfigurationService) UpdateWebhookConfiguration(id string, config *WebhookConfiguration) (*WebhookConfiguration, error) {
	path := fmt.Sprintf("%s%s/", webhookConfigurationsPath, id)
	body, err := c.client.newRequestDo("PATCH", path, config)
	if err != nil {
		return nil, err
	}
	var w WebhookConfiguration
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *WebhookConfigurationService) DeleteWebhookConfiguration(id string) error {
	path := fmt.Sprintf("%s%s/", webhookConfigurationsPath, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
