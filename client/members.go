package client

import (
	"encoding/json"
	"fmt"
)

type MemberService service

type Member struct {
	Unique_Id string `json:"unique_id",omitempty`
	Team      string `json:"team",omitempty`
	User      string `json:"user",omitempty`
	Role      int    `json:"role",omitempty`
}

type MemberResponse struct {
	Unique_Id    string `json:"unique_id"`
	Team         string `json:"team"`
	User         User   `json:"user"`
	Joining_Date string `json:"joining_date"`
	Role         int    `json:"role"`
}

func (c *MemberService) CreateTeamMember(team string, member *Member) (*Member, error) {
	path := fmt.Sprintf("/api/account/teams/%s/members/", team)

	body, err := c.client.newRequestDo("POST", path, member)
	if err != nil {
		return nil, err
	}
	var s Member
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *MemberService) UpdateTeamMember(member *Member) (*Member, error) {
	path := fmt.Sprintf("/api/account/teams/%s/members/%s/", member.Team, member.Unique_Id)
	body, err := c.client.newRequestDo("PATCH", path, member)
	if err != nil {
		return nil, err
	}
	var s Member
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *MemberService) DeleteTeamMember(team string, member string) error {
	path := fmt.Sprintf("/api/account/teams/%s/members/%s/", team, member)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *MemberService) GetTeamMembers(team string) ([]MemberResponse, error) {
	path := fmt.Sprintf("/api/account/teams/%s/members/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []MemberResponse
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *MemberService) GetTeamMembersByID(team, id string) (*MemberResponse, error) {
	path := fmt.Sprintf("/api/account/teams/%s/members/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s MemberResponse
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
