package zenduty

import (
	"encoding/json"
	"fmt"
)

type ScheduleService service

type Restrictions struct {
	Duration       int    `json:"duration"`
	StartDayOfWeek int    `json:"start_day_of_week"`
	StartTimeOfDay string `json:"start_time_of_day"`
	Unique_Id      string `json:"unique_id",omitempty`
}
type Users struct {
	User      string `json:"user"`
	Position  int    `json:"position"`
	Unique_Id string `json:"unique_id",omitempty`
}

type Overrides struct {
	Name      string `json:"name"`
	User      string `json:"user"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Unique_Id string `json:"unique_id",omitempty`
}
type Layers struct {
	ShiftLength       int            `json:"shift_length"`
	Name              string         `json:"name"`
	RotationStartTime string         `json:"rotation_start_time"`
	RotationEndTime   string         `json:"rotation_end_time"`
	UniqueId          string         `json:"unique_id"`
	LastEdited        string         `json:"last_edited"`
	RestrictionType   int            `json:"restriction_type",omitempty`
	IsActive          bool           `json:"is_active",omitempty`
	Restrictions      []Restrictions `json:"restrictions"`
	Users             []Users        `json:"users"`
}

type CreateUserLayer struct {
	User string `json:"user"`
}
type CreateLayers struct {
	ShiftLength       int               `json:"shift_length"`
	Name              string            `json:"name"`
	RotationStartTime string            `json:"rotation_start_time"`
	RotationEndTime   string            `json:"rotation_end_time"`
	RestrictionType   int               `json:"restriction_type",omitempty`
	Users             []CreateUserLayer `json:"users"`
}

type CreateSchedule struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Summary     string         `json:"summary"`
	Time_zone   string         `json:"time_zone"`
	Team        string         `json:"team"`
	Layers      []CreateLayers `json:"layers"`
	Overrides   []Overrides    `json:"overrides"`
	Unique_Id   string         `json:"unique_id",omitempty`
}
type Schedules struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Summary     string      `json:"summary"`
	Time_zone   string      `json:"time_zone"`
	Team        string      `json:"team"`
	Unique_Id   string      `json:"unique_id",omitempty`
	Layers      []Layers    `json:"layers"`
	Overrides   []Overrides `json:"overrides"`
}

func (c *ScheduleService) CreateSchedule(team string, schedule *CreateSchedule) (*CreateSchedule, error) {

	path := fmt.Sprintf("/api/account/teams/%s/schedules/", team)
	body, err := c.client.newRequestDo("POST", path, schedule)
	if err != nil {
		return nil, err
	}
	var s CreateSchedule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *ScheduleService) GetSchedules(team string) ([]Schedules, error) {
	path := fmt.Sprintf("/api/account/teams/%s/schedules/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []Schedules
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *ScheduleService) GetScheduleByID(team, id string) (*Schedules, error) {
	path := fmt.Sprintf("/api/account/teams/%s/schedules/%s/", team, id)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s Schedules
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *ScheduleService) DeleteScheduleByID(team, id string) error {
	path := fmt.Sprintf("/api/account/teams/%s/schedules/%s/", team, id)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *ScheduleService) UpdateScheduleByID(team, id string, schedule *CreateSchedule) (*CreateSchedule, error) {

	path := fmt.Sprintf("/api/account/teams/%s/schedules/%s/", team, id)
	body, err := c.client.newRequestDo("PATCH", path, schedule)
	if err != nil {
		return nil, err
	}
	var s CreateSchedule
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
