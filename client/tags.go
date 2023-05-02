package client

import (
	"encoding/json"
	"fmt"
)

type TagsService service

type Tag struct {
	UniqueID     string `json:"unique_id,omitempty"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Team         int    `json:"team,omitempty"`
	CreationDate string `json:"creation_date,omitempty"`
}

type ReadTag struct {
	UniqueID     string `json:"unique_id,omitempty"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Team         string `json:"team,omitempty"`
	CreationDate string `json:"creation_date,omitempty"`
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

func (c *TagsService) UpdateTag(team string, tagID string, tags *Tag) (*Tag, error) {
	path := fmt.Sprintf("/api/account/teams/%s/tags/%s/", team, tagID)
	body, err := c.client.newRequestDo("PUT", path, tags)
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

func (c *TagsService) GetTags(team string) ([]ReadTag, error) {
	path := fmt.Sprintf("/api/account/teams/%s/tags/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []ReadTag
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *TagsService) GetTagID(team, id string) (*ReadTag, error) {
	path := fmt.Sprintf("/api/account/teams/%s/tags/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s ReadTag
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *TagsService) DeleteTag(team, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/tags/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
