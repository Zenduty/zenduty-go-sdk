package client

import (
	"encoding/json"
)

type UserService service

type UserObj struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

type GetUsers struct {
	Unique_Id string  `json:"unique_id"`
	User      UserObj `json:"user"`
}

func (c *UserService) GetUsers(email string) ([]GetUsers, error) {
	path := "/api/account/users"

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
