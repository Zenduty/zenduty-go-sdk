package client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type IncidentService service

type Incident struct {
	Service          string `json:"service"`
	EscalationPolicy string `json:"escalation_policy"`
	User             string `json:"user"`
	Title            string `json:"title"`
	Summary          string `json:"summary"`
}

type serviceObject struct {
	Name                  string `json:"name"`
	CreationDate          string `json:"creation_date"`
	Summary               string `json:"summary"`
	Description           string `json:"description"`
	UniqueID              string `json:"unique_id"`
	AutoResolveTimeouts   int    `json:"auto_resolve_timeout"`
	CreatedBy             string `json:"created_by"`
	TeamPriority          string `json:"team_priority"`
	TaskTemplate          string `json:"task_template"`
	AcknowledgmentTimeout int    `json:"acknowledge_timeout"`
	Status                int    `json:"status"`
	EscalationPolicy      string `json:"escalation_policy"`
	Team                  string `json:"team"`
	SLA                   string `json:"sla"`
	CollationTime         int    `json:"collation_time"`
	Collation             int    `json:"collation"`
}

type escalationPolicyObject struct {
	UniqueID string `json:"unique_id"`
	Name     string `json:"name"`
}

type Incidents struct {
	Summary                string `json:"summary"`
	IncidentNumber         int    `json:"incident_number"`
	CreationDate           string `json:"creation_date"`
	Status                 int    `json:"status"`
	UniqueID               string `json:"unique_id"`
	ServiceObject          serviceObject
	Title                  string                 `json:"title"`
	IncidentKey            string                 `json:"incident_key"`
	Service                string                 `json:"service"`
	Urgency                int                    `json:"urgency"`
	MergedWith             string                 `json:"merged_with"`
	AssignedTo             string                 `json:"assigned_to"`
	EscalationPolicy       string                 `json:"escalation_policy"`
	EscalationPolicyObject escalationPolicyObject `json:"escalationPolicyObject"`
	AssignedToName         string                 `json:"assigned_to_name"`
	ResolvedDate           string                 `json:"resolved_date"`
	AcknowledgedDate       string                 `json:"acknowledged_date"`
	ContextWindowStart     string                 `json:"context_window_start"`
	ContextWindowEnd       string                 `json:"context_window_end"`
	Tags                   []IncidentTag          `json:"tags"`
	SLA                    string                 `json:"sla"`
	TeamPriority           string                 `json:"team_priority"`
	TeamPriorityObject     string                 `json:"team_priority_object"`
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
type AddIncidentNote struct {
	Note string `json:"note"`
}
type IncidentNote struct {
	UniqueID     string `json:"unique_id"`
	Incident     int    `json:"incident"`
	Username     string `json:"user"`
	Note         string `json:"note"`
	Name         string `json:"user_name"`
	CreationDate string `json:"creation_date"`
}

type IncidentNotes struct {
	Results  []IncidentNote `json:"results"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Count    int            `json:"count"`
}

type IncidentTag struct {
	UniqueID     string `json:"unique_id"`
	Name         string `json:"name"`
	Incident     int    `json:"incident"`
	CreationDate string `json:"creation_date"`
	TagID        string `json:"tagID"`
	Color        string `json:"color"`
}

type AddIncidentTag struct {
	TagID string `json:"team_tag"` // uniqueID of team tags
}

type IncidentAlerts struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Alerts   []AlertResponse `json:"results"`
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

func (c *IncidentService) UpdateIncident(incidentNumber string, incident *IncidentStatus) (*Incidents, error) {
	path := fmt.Sprintf("/api/incidents/%s/", incidentNumber)
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

func (c *IncidentService) GetIncidentByNumber(incidentNumber string) (*Incidents, error) {
	path := fmt.Sprintf("/api/incidents/%s/", incidentNumber)
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

func (c *IncidentService) CreateIncidentNote(incidentNumber string, note *AddIncidentNote) (*IncidentNote, error) {
	path := fmt.Sprintf("/api/incidents/%s/note/", incidentNumber)

	body, err := c.client.newRequestDo("POST", path, note)
	if err != nil {
		return nil, err
	}
	var i IncidentNote
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IncidentService) UpdateIncidentNote(incidentNumber, noteID string, note *AddIncidentNote) (*IncidentNote, error) {
	path := fmt.Sprintf("/api/incidents/%s/note/%s/", incidentNumber, noteID)

	body, err := c.client.newRequestDo("PUT", path, note)
	if err != nil {
		return nil, err
	}
	var i IncidentNote
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IncidentService) GetIncidentNotes(incidentNumber string) (*IncidentNotes, error) {
	path := &url.URL{
		Path: fmt.Sprintf("/api/incidents/%s/note/", incidentNumber),
	}
	body, err := c.client.newRequestDo("GET", path.String(), nil)
	if err != nil {
		return nil, err
	}
	var i IncidentNotes
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil

}

func (c *IncidentService) DeleteIncidentNote(incidentNumber, noteID string) error {
	path := fmt.Sprintf("/api/incidents/%s/note/%s/", incidentNumber, noteID)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *IncidentService) CreateIncidentTags(incidentNumber string, note *AddIncidentTag) (*IncidentTag, error) {
	path := fmt.Sprintf("/api/incidents/%s/tags/", incidentNumber)

	body, err := c.client.newRequestDo("POST", path, note)
	if err != nil {
		return nil, err
	}
	var i IncidentTag
	err = json.Unmarshal(body.BodyBytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *IncidentService) GetIncidentTags(incidentNumber string) ([]IncidentTag, error) {
	path := fmt.Sprintf("/api/incidents/%s/tags/", incidentNumber)

	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var r []IncidentTag
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *IncidentService) GetIncidentTag(incidentNumber, tagID string) (*IncidentTag, error) {
	path := fmt.Sprintf("/api/incidents/%s/tags/%s/", incidentNumber, tagID)

	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var r IncidentTag
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *IncidentService) DeleteIncidentTag(incidentNumber, tagID string) error {
	path := fmt.Sprintf("/api/incidents/%s/tags/%s/", incidentNumber, tagID)
	_, err := c.client.newRequestDo("DELETE", path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *IncidentService) GetIncientAlerts(incidentNumber string) (*IncidentAlerts, error) {

	path := fmt.Sprintf("/api/incidents/%s/alerts/", incidentNumber)
	body, err := c.client.newRequestDo("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var r IncidentAlerts
	err = json.Unmarshal(body.BodyBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
