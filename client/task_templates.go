package client

import (
	"encoding/json"
	"fmt"
)

type TaskTemplateService service

type TaskTemplateObj struct {
	UniqueID     string `json:"unique_id,omitempty"`
	Team         string `json:"team,omitempty"`
	Name         string `json:"name"`
	Summary      string `json:"summary"`
	CreationDate string `json:"creation_date,omitempty"`
}

type TaskTemplateTaskObj struct {
	UniqueID     string `json:"unique_id,omitempty"`
	TaskTemplate string `json:"task_template"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	DueIn        int    `json:"due_in"`
	CreationDate string `json:"creation_date,omitempty"`
	Role         string `json:"role"`
	Positon      int    `json:"position,omitempty"`
}

func (c *TaskTemplateService) CreateTaskTemplate(team string, tasktemplare *TaskTemplateObj) (*TaskTemplateObj, error) {

	path := fmt.Sprintf("/api/account/teams/%s/task_templates/", team)
	body, err := c.client.newRequestDo("POST", path, tasktemplare)
	if err != nil {
		return nil, err
	}
	var s TaskTemplateObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *TaskTemplateService) GetTaskTemplates(team string) ([]TaskTemplateObj, error) {
	path := fmt.Sprintf("/api/account/teams/%s/task_templates/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []TaskTemplateObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *TaskTemplateService) GetTaskTemplateByID(team, id string) (*TaskTemplateObj, error) {
	path := fmt.Sprintf("/api/account/teams/%s/task_templates/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s TaskTemplateObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *TaskTemplateService) DeleteTaskTemplateByID(team, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/task_templates/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *TaskTemplateService) UpdateTaskTemplateByID(team, id string, TaskTemplate *TaskTemplateObj) (*TaskTemplateObj, error) {

	path := fmt.Sprintf("/api/account/teams/%s/task_templates/%s/", team, id)
	body, err := c.client.newRequestDo("PATCH", path, TaskTemplate)
	if err != nil {
		return nil, err
	}
	var s TaskTemplateObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *TaskTemplateService) CreateTaskTemplateTask(team, task_id string, tasktemplatetask *TaskTemplateTaskObj) (*TaskTemplateTaskObj, error) {

	path := fmt.Sprintf("/api/account/teams/%s/task_templates/%s/tasks/", team, task_id)
	body, err := c.client.newRequestDo("POST", path, tasktemplatetask)
	if err != nil {
		return nil, err
	}
	var s TaskTemplateTaskObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *TaskTemplateService) GetTaskTemplateTasks(team, task_id string) ([]TaskTemplateTaskObj, error) {
	path := fmt.Sprintf("/api/account/teams/%s/task_templates/%s/tasks/", team, task_id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []TaskTemplateTaskObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *TaskTemplateService) GetTaskTemplateTaskByID(team, task_id, id string) (*TaskTemplateTaskObj, error) {
	path := fmt.Sprintf("/api/account/teams/%s/task_templates/%s/tasks/%s/", team, task_id, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s TaskTemplateTaskObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *TaskTemplateService) DeleteTaskTemplateTaskByID(team, task_id, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/task_templates/%s/tasks/%s/", team, task_id, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *TaskTemplateService) UpdateTaskTemplateTaskByID(team, task_id, id string, TaskTemplateTask *TaskTemplateTaskObj) (*TaskTemplateTaskObj, error) {

	path := fmt.Sprintf("/api/account/teams/%s/task_templates/%s/tasks/%s/", team, task_id, id)
	body, err := c.client.newRequestDo("PATCH", path, TaskTemplateTask)
	if err != nil {
		return nil, err
	}
	var s TaskTemplateTaskObj
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
