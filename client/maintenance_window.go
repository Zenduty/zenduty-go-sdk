package client

import (
	"encoding/json"
	"fmt"
)

type MaintenanceWindowService service

type ServiceMaintenance struct {
	UniqueID string `json:"unique_id,omitempty"`
	Service  string `json:"service"`
}

type MaintenanceWindow struct {
	UniqueID       string               `json:"unique_id,omitempty"`
	Name           string               `json:"name"`
	RepeatInterval int                  `json:"repeat_interval,omitempty"`
	RepeatUntil    string               `json:"repeat_until,omitempty"`
	StartTime      string               `json:"start_time"`
	EndTime        string               `json:"end_time"`
	TimeZone       string               `json:"time_zone"`
	Services       []ServiceMaintenance `json:"services"`
	CreationDate   string               `json:"creation_date"`
}

func (c *MaintenanceWindowService) CreateMaintenanceWindow(team string, maintenance *MaintenanceWindow) (*MaintenanceWindow, error) {
	path := fmt.Sprintf("/api/account/teams/%s/maintenance/", team)
	body, err := c.client.newRequestDo("POST", path, maintenance)
	if err != nil {
		return nil, err
	}
	var s MaintenanceWindow
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *MaintenanceWindowService) UpdateMaintenanceWindow(team string, maintenanceID string, maintenance *MaintenanceWindow) (*MaintenanceWindow, error) {
	path := fmt.Sprintf("/api/account/teams/%s/maintenance/%s/", team, maintenanceID)
	body, err := c.client.newRequestDo("PUT", path, maintenance)
	if err != nil {
		return nil, err
	}
	var s MaintenanceWindow
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *MaintenanceWindowService) GetMaintenanceWindows(team string) ([]MaintenanceWindow, error) {
	path := fmt.Sprintf("/api/account/teams/%s/maintenance/", team)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s []MaintenanceWindow
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *MaintenanceWindowService) DeleteMaintenanceWindow(team string, maintenanceID string) error {
	path := fmt.Sprintf("/api/account/teams/%s/maintenance/%s/", team, maintenanceID)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *MaintenanceWindowService) GetMaintenanceWindowByID(team string, maintenanceID string) (*MaintenanceWindow, error) {
	path := fmt.Sprintf("/api/account/teams/%s/maintenance/%s/", team, maintenanceID)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var s MaintenanceWindow
	err = json.Unmarshal(body.BodyBytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
