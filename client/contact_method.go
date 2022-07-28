package client

import (
	"encoding/json"
	"fmt"
)

type ContactMethodService service

type ContactMethod struct {
	UniqueID     string `json:"unique_id"`
	Name         string `json:"name"`
	ContactType  int    `json:"contact_type"`
	Value        string `json:"value"`
	CreationDate string `json:"creation_date"`
}

func (c *ContactMethodService) GetContactMethods(username string) ([]ContactMethod, error) {
	path := fmt.Sprintf("/api/account/users/%s/contacts/", username)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []ContactMethod
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *ContactMethodService) GetContactMethodByID(username string, contactMethodID string) (*ContactMethod, error) {
	path := fmt.Sprintf("/api/account/users/%s/contacts/%s/", username, contactMethodID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s1 ContactMethod
	err = json.Unmarshal(body.BodyBytes, &s1)
	if err != nil {
		return nil, err
	}
	return &s1, nil
}
