package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://www.zenduty.com"
)

type service struct {
	client *Client
}

type Config struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

type Client struct {
	baseURL           *url.URL
	client            *http.Client
	Config            *Config
	Teams             *TeamService
	Services          *Service
	Schedules         *ScheduleService
	Roles             *RoleService
	Integrations      *IntegrationServerice
	Incidents         *IncidentService
	Esp               *EspService
	Members           *MemberService
	Invite            *InviteService
	Users             *UserService
	AlertRules        *AlertRuleService
	Priority          *PriorityService
	Tags              *TagsService
	MaintenanceWindow *MaintenanceWindowService
	NotificationRules *NotificationRulesService
	ContactMethod     *ContactMethodService
	AccountRole       *AccountRoleService
	GlobalRouter      *GlobalRouterService
	Events            *EventsService
	Sla               *SLAService
	PostIncidentTask  *PostIncidentTaskService
	TaskTemplate      *TaskTemplateService
	OutgoingRules     *OutgoingRulesService
}

type Response struct {
	Response  *http.Response
	BodyBytes []byte
}

func NewClient(config *Config) (*Client, error) {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}

	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL: baseURL,
		client:  config.HTTPClient,
		Config:  config,
	}
	c.Teams = &TeamService{c}
	c.Services = &Service{c}
	c.Schedules = &ScheduleService{c}
	c.Roles = &RoleService{c}
	c.Integrations = &IntegrationServerice{c}
	c.Incidents = &IncidentService{c}
	c.Esp = &EspService{c}
	c.Members = &MemberService{c}
	c.Invite = &InviteService{c}
	c.Users = &UserService{c}
	c.AlertRules = &AlertRuleService{c}
	c.Priority = &PriorityService{c}
	c.Tags = &TagsService{c}
	c.MaintenanceWindow = &MaintenanceWindowService{c}
	c.NotificationRules = &NotificationRulesService{c}
	c.ContactMethod = &ContactMethodService{c}
	c.AccountRole = &AccountRoleService{c}
	c.GlobalRouter = &GlobalRouterService{c}
	c.Events = &EventsService{c}
	c.Sla = &SLAService{c}
	c.PostIncidentTask = &PostIncidentTaskService{c}
	c.TaskTemplate = &TaskTemplateService{c}
	c.OutgoingRules = &OutgoingRulesService{c}
	return c, nil

}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)

	var buf []byte
	if body != nil {
		buf, _ = json.Marshal(body)
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.Config.Token))

	return req, nil
}

func (c *Client) newRequestDo(method, path string, body interface{}) (*Response, error) {
	req, err := c.newRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req)
}

func (c *Client) doRequest(req *http.Request) (*Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	response := &Response{
		Response:  res,
		BodyBytes: body,
	}
	if err := c.checkResponse(response); err != nil {
		return response, err
	}

	// if v != nil {
	// 	if err := c.DecodeJSON(response, v); err != nil {
	// 		return response, err
	// 	}
	// }

	return response, nil

}

func (c *Client) DecodeJSON(res *Response, v interface{}) error {
	return json.Unmarshal(res.BodyBytes, v)
}

func (c *Client) checkResponse(res *Response) error {
	if res.Response.StatusCode >= 200 && res.Response.StatusCode <= 299 {
		return nil
	}

	return c.decodeErrorResponse(res)
}

func (c *Client) decodeErrorResponse(res *Response) error {

	v := &errorResponse{Error: &Error{ErrorResponse: res, Code: res.Response.StatusCode}}
	if err := c.DecodeJSON(res, v); err != nil {

		return fmt.Errorf("%s APIs call to %s failed: %v error: %s", res.Response.Request.Method, res.Response.Request.URL.String(), res.Response.Status, string(res.BodyBytes))
	}

	return v.Error
}

func CheckError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
