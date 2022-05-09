package client

import (
	"encoding/json"
)

type InviteService service

type EmailAccounts struct {
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Role       int    `json:"role"`
}

type Invite struct {
	EmailAccounts []EmailAccounts `json:"email_accounts"`
	Team          string          `json:"team"`
}

type InviteResponse struct {
	Unique_Id    string `json:"unique_id"`
	Team         string `json:"team"`
	User         User   `json:"user"`
	Joining_Date string `json:"joining_date"`
	Role         int    `json:"role"`
}

func (c *InviteService) CreateInvite(invite *Invite) ([]InviteResponse, error) {
	path := "/api/account/invite/"
	body, err := c.client.newRequestDo("POST", path, invite)
	if err != nil {
		return nil, err
	}
	var s []InviteResponse
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
