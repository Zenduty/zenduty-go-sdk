package client

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type PostIncidentTaskService service

type AssignedToObj struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

type PostIncidentTaskObj struct {
	UniqueID      string        `json:"unique_id,omitempty"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Team          string        `json:"team,omitempty"`
	Status        int           `json:"status"`
	DueInTime     *string       `json:"due_in_time"`
	CreationDate  string        `json:"creation_date,omitempty"`
	AssignedTo    string        `json:"assigned_to,omitempty"`
	AssignedToObj AssignedToObj `json:"assigned_to_obj"`
}

type PostIncidentTaskPagination struct {
	Results  []PostIncidentTaskObj `json:"results"`
	Next     string                `json:"next"`
	Previous string                `json:"previous"`
	Count    int                   `json:"count"`
}

func (c *PostIncidentTaskService) CreatePostIncidentTask(team string, postincidentask *PostIncidentTaskObj) (*PostIncidentTaskObj, error) {

	path := fmt.Sprintf("/api/account/teams/%s/post_incident_tasks/", team)
	body, err := c.client.newRequestDo("POST", path, postincidentask)
	if err != nil {
		return nil, err
	}
	var s PostIncidentTaskObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetPostIncidentTasks fetches the first page of a team's post-incident tasks.
func (c *PostIncidentTaskService) GetPostIncidentTasks(team string) (*PostIncidentTaskPagination, error) {
	return c.GetPostIncidentTasksPage(team, 1)
}

// GetPostIncidentTasksPage fetches one page (10 tasks) of a team's
// post-incident tasks. page is 1-based; 0 and 1 both fetch the first page.
func (c *PostIncidentTaskService) GetPostIncidentTasksPage(team string, page int) (*PostIncidentTaskPagination, error) {
	path := fmt.Sprintf("/api/account/teams/%s/post_incident_tasks/", team)
	if page > 1 {
		path += "?page=" + strconv.Itoa(page)
	}
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s PostIncidentTaskPagination
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetAllPostIncidentTasks follows pagination and returns every post-incident
// task on the team.
func (c *PostIncidentTaskService) GetAllPostIncidentTasks(team string) ([]PostIncidentTaskObj, error) {
	var tasks []PostIncidentTaskObj
	for page := 1; ; page++ {
		p, err := c.GetPostIncidentTasksPage(team, page)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, p.Results...)
		if p.Next == "" || len(p.Results) == 0 {
			return tasks, nil
		}
	}
}

func (c *PostIncidentTaskService) GetPostIncidentTaskByID(team, id string) (*PostIncidentTaskObj, error) {
	path := fmt.Sprintf("/api/account/teams/%s/post_incident_tasks/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s PostIncidentTaskObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *PostIncidentTaskService) DeletePostIncidentTaskByID(team, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/post_incident_tasks/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *PostIncidentTaskService) UpdatePostIncidentTaskByID(team, id string, PostIncidentTask *PostIncidentTaskObj) (*PostIncidentTaskObj, error) {

	path := fmt.Sprintf("/api/account/teams/%s/post_incident_tasks/%s/", team, id)
	body, err := c.client.newRequestDo("PATCH", path, PostIncidentTask)
	if err != nil {
		return nil, err
	}
	var s PostIncidentTaskObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
