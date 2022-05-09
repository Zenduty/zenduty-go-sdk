package client

import (
	"encoding/json"
	"fmt"
)

type TagsService service

type Tag struct {
	Unique_Id     string `json:"unique_id,omitempty"`
	Name          string `json:"name"`
	Color         string `json:"color"`
	Team          string `json:"team,omitempty"`
	Creation_Date string `json:"creation_date,omitempty"`
}

func (c *TagsService) CreateTag(team string, tags *Tag) (*Tag, error) {
	path := fmt.Sprintf("/api/account/teams/%s/tags/", team)
	body, err := c.client.newRequestDo("POST", path, tags)
	if err != nil {
		return nil, err
	}
	var s Tag
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *TagsService) GetPriority(team string) ([]Tag, error) {
	path := fmt.Sprintf("/api/account/teams/%s/tags/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []Tag
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *TagsService) GetPriorityById(team, id string) (*Tag, error) {
	path := fmt.Sprintf("/api/account/teams/%s/tags/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s Tag
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *TagsService) DeletePriority(team, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/tags/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
