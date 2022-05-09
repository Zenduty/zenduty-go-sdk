package client

import (
	"encoding/json"
	"fmt"
)

type TeamService service

type User struct {
	Username   string `json:"username"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
}

type members struct {
	Unique_Id    string `json:"unique_id"`
	Team         string `json:"team"`
	User         User   `json:"user"`
	Joining_Date string `json:"joining_date"`
	Role         int    `json:"role"`
}

type CreateTeams struct {
	Name      string `json:"name"`
	Unique_Id string `json:"unique_id",omitempty`
}
type Team struct {
	Unique_Id     string    `json:"unique_id"`
	Name          string    `json:"name"`
	Account       string    `json:"account"`
	Creation_Date string    `json:"creation_date"`
	Owner         string    `json:"owner"`
	Roles         []Roles   `json:"roles"`
	Members       []members `json:"members"`
}

func (c *TeamService) CreateTeam(team *CreateTeams) (*Team, error) {
	path := "/api/account/teams/"
	body, err := c.client.newRequestDo("POST", path, team)
	if err != nil {
		return nil, err
	}
	var t Team
	err = json.Unmarshal(body.BodyBytes, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *TeamService) UpdateTeam(id string, team *CreateTeams) (*Team, error) {

	path := fmt.Sprintf("/api/account/teams/%s/", id)
	res, err := c.client.newRequestDo("PATCH", path, team)
	if err != nil {
		return nil, err
	}
	var t Team
	err = json.Unmarshal(res.BodyBytes, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *TeamService) GetTeamById(uniqie_id string) (*Team, error) {
	path := fmt.Sprintf("/api/account/teams/%s/", uniqie_id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var t Team
	err = json.Unmarshal(body.BodyBytes, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *TeamService) GetTeams() ([]Team, error) {
	path := "/api/account/teams/"

	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var t []Team
	err = json.Unmarshal(body.BodyBytes, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *TeamService) DeleteTeam(id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/", id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}
