package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	Applications      *ApplicationsService
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
	c.Applications = &ApplicationsService{c}
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
	// url.Parse (rather than url.URL{Path: path}) keeps query strings like
	// "?page=2" intact instead of escaping them into the path.
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
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
	var res *Response
	var err error
	backoff := 2 * time.Second
	for attempt := 0; attempt < 6; attempt++ {
		// build a fresh request each attempt: the previous one's body reader
		// is already consumed
		var req *http.Request
		req, err = c.newRequest(method, path, body)
		if err != nil {
			return nil, err
		}
		res, err = c.doRequest(req)
		// a 429 means the server rejected the request before processing it,
		// so retrying is safe for every method; anything else returns as-is
		if err == nil || res == nil || res.Response == nil || res.Response.StatusCode != http.StatusTooManyRequests {
			return res, err
		}
		wait := backoff + time.Duration(rand.Int63n(int64(backoff/2)))
		if retryAfter := res.Response.Header.Get("Retry-After"); retryAfter != "" {
			if secs, parseErr := strconv.Atoi(retryAfter); parseErr == nil && secs > 0 {
				wait = time.Duration(secs)*time.Second + time.Duration(rand.Int63n(int64(time.Second)))
			}
		}
		time.Sleep(wait)
		if backoff < 30*time.Second {
			backoff *= 2
		}
	}
	return res, err
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
	// Fall back to a bare *Error when the body is not JSON, is not an object
	// (e.g. {"error": "text"}), or explicitly nulls the error key — every
	// non-2xx response must yield a typed *Error carrying the status code.
	if err := c.DecodeJSON(res, v); err != nil || v.Error == nil {
		return &Error{
			ErrorResponse: res,
			Code:          res.Response.StatusCode,
			Message:       string(res.BodyBytes),
		}
	}

	return v.Error
}

func CheckError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
