package client

import (
	"encoding/json"
	"fmt"
)

type UserService service

type UserObj struct {
	Email      string `json:"email"`
	Username   string `json:"username,omitempty"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Role         int  `json:"role,omitempty"`
}

type GetUsers struct {
	Unique_Id string  `json:"unique_id"`
	User      UserObj `json:"user"`
}
type CreateUser struct {
	Team string `json:"team"`
	UniqueId string `json:"unique_id,omitempty"`
	User      UserObj `json:"user_detail"`
}
type GetUser struct {
	Team      string `json:"team,omitempty"`
	UniqueId  string `json:"unique_id,omitempty"`
	User      UserObj `json:"user"`
	Role      int     `json:"role"`
}

func (c *UserService) CreateUser(user *CreateUser) (*GetUser, error) {
	path := "/api/account/api_invite/"
	body, err := c.client.newRequestDo("POST", path, user)
	if err != nil {
		return nil, err
	}
	var s GetUser
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *UserService) UpdateUser(username string,user *UserObj) (*GetUser, error) {

	path := fmt.Sprintf("/api/account/users/%s/", username)
	body, err := c.client.newRequestDo("PATCH", path, user)
	if err != nil {
		return nil, err
	}
	var s GetUser
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c * UserService) GetUser(username string) (*GetUser, error) {
	path := fmt.Sprintf("/api/account/users/%s/", username)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s GetUser
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *UserService) GetUsers(email string) ([]GetUsers, error) {
	path := "/api/account/users/"

	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var i []GetUsers

	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}

	var j []GetUsers

	for _, v := range i {
		if v.User.Email == email {
			j = append(j, v)
		}
	}
	return j, nil
}
