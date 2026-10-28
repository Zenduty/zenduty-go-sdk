package client

import (
	"encoding/json"
	"fmt"
)

type AccountStakeholderTemplateService service

// AccountStakeholderTemplate is an account-level stakeholder update template.
// The API's account field on this object is a numeric primary key, unlike the
// unique_id slugs elsewhere, so it is not modelled.
type AccountStakeholderTemplate struct {
	UniqueID     string `json:"unique_id,omitempty"`
	Title        string `json:"title"`
	Subject      string `json:"subject"`
	Template     string `json:"template"`
	CreationDate string `json:"creation_date,omitempty"`
}

const accountStakeholderTemplatesPath = "/api/account/account-stakeholder-template/"

func (c *AccountStakeholderTemplateService) CreateAccountStakeholderTemplate(template *AccountStakeholderTemplate) (*AccountStakeholderTemplate, error) {
	body, err := c.client.newRequestDo("POST", accountStakeholderTemplatesPath, template)
	if err != nil {
		return nil, err
	}
	var t AccountStakeholderTemplate
	if err := json.Unmarshal(body.BodyBytes, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *AccountStakeholderTemplateService) GetAccountStakeholderTemplates() ([]AccountStakeholderTemplate, error) {
	var all []AccountStakeholderTemplate
	err := c.client.getAll(accountStakeholderTemplatesPath, func(results []byte) error {
		var page []AccountStakeholderTemplate
		if err := json.Unmarshal(results, &page); err != nil {
			return err
		}
		all = append(all, page...)
		return nil
	})
	return all, err
}

func (c *AccountStakeholderTemplateService) GetAccountStakeholderTemplate(id string) (*AccountStakeholderTemplate, error) {
	path := fmt.Sprintf("%s%s/", accountStakeholderTemplatesPath, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var t AccountStakeholderTemplate
	if err := json.Unmarshal(body.BodyBytes, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *AccountStakeholderTemplateService) UpdateAccountStakeholderTemplate(id string, template *AccountStakeholderTemplate) (*AccountStakeholderTemplate, error) {
	path := fmt.Sprintf("%s%s/", accountStakeholderTemplatesPath, id)
	body, err := c.client.newRequestDo("PATCH", path, template)
	if err != nil {
		return nil, err
	}
	var t AccountStakeholderTemplate
	if err := json.Unmarshal(body.BodyBytes, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *AccountStakeholderTemplateService) DeleteAccountStakeholderTemplate(id string) error {
	path := fmt.Sprintf("%s%s/", accountStakeholderTemplatesPath, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
