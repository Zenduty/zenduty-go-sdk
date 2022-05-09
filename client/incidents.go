package client

import (
	"encoding/json"
)

type IncidentService service

type Incident struct {
	Service          string `json:"service"`
	EscalationPolicy string `json:"escalation_policy"`
	User             string `json:"user"`
	Title            string `json:"title"`
	Summary          string `json:"summary"`
}

type service_object struct {
	Name                   string `json:"name"`
	Creation_Date          string `json:"creation_date"`
	Summary                string `json:"summary"`
	Description            string `json:"description"`
	Unique_Id              string `json:"unique_id"`
	Auto_Resolve_Timeouts  int    `json:"auto_resolve_timeout"`
	Created_By             string `json:"created_by"`
	Team_Priority          string `json:"team_priority"`
	Task_Template          string `json:"task_template"`
	Acknowledgment_Timeout int    `json:"acknowledge_timeout"`
	Status                 int    `json:"status"`
	EscalationPolicy       string `json:"escalation_policy"`
	Team                   string `json:"team"`
	Sla                    string `json:"sla"`
	Collation_Time         int    `json:"collation_time"`
	Collation              int    `json:"collation"`
}

type escalation_policy_object struct {
	Unique_Id string `json:"unique_id"`
	Name      string `json:"name"`
}

type Incidents struct {
	Summary                  string `json:"summary"`
	Incident_Number          int    `json:"incident_number"`
	Creation_Date            string `json:"creation_date"`
	Status                   int    `json:"status"`
	Unique_Id                string `json:"unique_id"`
	Service_Object           service_object
	Title                    string                   `json:"title"`
	Incident_Key             string                   `json:"incident_key"`
	Service                  string                   `json:"service"`
	Urgency                  int                      `json:"urgency"`
	Merged_With              string                   `json:"merged_with"`
	Assigned_To              string                   `json:"assigned_to"`
	Escalation_Policy        string                   `json:"escalation_policy"`
	Escalation_Policy_Object escalation_policy_object `json:"escalation_policy_object"`
	Assigned_to_name         string                   `json:"assigned_to_name"`
	Resolved_Date            string                   `json:"resolved_date"`
	Acknowledged_Date        string                   `json:"acknowledged_date"`
	Context_Window_start     string                   `json:"context_window_start"`
	Context_Window_end       string                   `json:"context_window_end"`
	Tags                     []string                 `json:"tags"`
	Sla                      string                   `json:"sla"`
	Team_Priority            string                   `json:"team_priority"`
	Team_Priority_Object     string                   `json:"team_priority_object"`
}

type IncidentPagination struct {
	Results  []Incidents `json:"results"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Count    int         `json:"count"`
}

type IncidentStatus struct {
	Status int `json:"status"`
}

func (c *IncidentService) CreateIncident(incident *Incident) (*Incidents, error) {
	path := "/api/incidents/"

	body, err := c.client.newRequestDo("POST", path, incident)
	if err != nil {
		return nil, err
	}
	var i Incidents
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IncidentService) UpdateIncident(id string, incident *IncidentStatus) (*Incidents, error) {
	path := "/api/incidents/" + id + "/"

	body, err := c.client.newRequestDo("PATCH", path, incident)
	if err != nil {
		return nil, err
	}
	var i Incidents
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IncidentService) GetIncidents() (*IncidentPagination, error) {
	path := "/api/incidents/"
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var i IncidentPagination
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IncidentService) GetIncidentByNumber(id string) (*Incidents, error) {
	path := "/api/incidents/" + id + "/"
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var i Incidents
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
