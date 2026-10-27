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

// ServiceMaintenanceWindow is a maintenance window attached to one service
// (teams/<team>/services/<service>/maintenance/), as opposed to the
// team-level MaintenanceWindow that lists its services. RepeatInterval is
// always sent so a window can be set back to non-repeating (0).
type ServiceMaintenanceWindow struct {
	UniqueID       string `json:"unique_id,omitempty"`
	Service        string `json:"service,omitempty"`
	Name           string `json:"name"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	TimeZone       string `json:"time_zone"`
	RepeatInterval int    `json:"repeat_interval"`
	RepeatUntil    string `json:"repeat_until,omitempty"`
	CreatedBy      string `json:"created_by,omitempty"`
	CreationDate   string `json:"creation_date,omitempty"`
}

func serviceMaintenancePath(team, service string) string {
	return fmt.Sprintf("/api/account/teams/%s/services/%s/maintenance/", team, service)
}

func (c *MaintenanceWindowService) CreateServiceMaintenanceWindow(team, service string, window *ServiceMaintenanceWindow) (*ServiceMaintenanceWindow, error) {
	body, err := c.client.newRequestDo("POST", serviceMaintenancePath(team, service), window)
	if err != nil {
		return nil, err
	}
	var w ServiceMaintenanceWindow
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *MaintenanceWindowService) GetServiceMaintenanceWindows(team, service string) ([]ServiceMaintenanceWindow, error) {
	var all []ServiceMaintenanceWindow
	err := c.client.getAll(serviceMaintenancePath(team, service), func(results []byte) error {
		var page []ServiceMaintenanceWindow
		if err := json.Unmarshal(results, &page); err != nil {
			return err
		}
		all = append(all, page...)
		return nil
	})
	return all, err
}

func (c *MaintenanceWindowService) GetServiceMaintenanceWindowByID(team, service, id string) (*ServiceMaintenanceWindow, error) {
	path := serviceMaintenancePath(team, service) + id + "/"
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var w ServiceMaintenanceWindow
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *MaintenanceWindowService) UpdateServiceMaintenanceWindow(team, service, id string, window *ServiceMaintenanceWindow) (*ServiceMaintenanceWindow, error) {
	path := serviceMaintenancePath(team, service) + id + "/"
	body, err := c.client.newRequestDo("PATCH", path, window)
	if err != nil {
		return nil, err
	}
	var w ServiceMaintenanceWindow
	if err := json.Unmarshal(body.BodyBytes, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

func (c *MaintenanceWindowService) DeleteServiceMaintenanceWindow(team, service, id string) error {
	path := serviceMaintenancePath(team, service) + id + "/"
	_, err := c.client.newRequestDo("DELETE", path, nil)
	return err
}
