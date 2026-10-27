package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeListHandlesArrayAndEnvelope(t *testing.T) {
	var fromArray []AccountStakeholderTemplate
	next, err := decodeList([]byte(`[{"title":"a"},{"title":"b"}]`), &fromArray)
	if err != nil || next != "" || len(fromArray) != 2 {
		t.Fatalf("array: next=%q err=%v len=%d", next, err, len(fromArray))
	}

	var fromPage []AccountStakeholderTemplate
	body := `{"next":"http://localhost:8000/api/x/?page=2","previous":null,"results":[{"title":"c"}]}`
	next, err = decodeList([]byte(body), &fromPage)
	if err != nil || next != "http://localhost:8000/api/x/?page=2" || len(fromPage) != 1 {
		t.Fatalf("envelope: next=%q err=%v len=%d", next, err, len(fromPage))
	}

	var fromEmpty []AccountStakeholderTemplate
	next, err = decodeList([]byte(`{"next":null,"results":[]}`), &fromEmpty)
	if err != nil || next != "" || len(fromEmpty) != 0 {
		t.Fatalf("empty envelope: next=%q err=%v len=%d", next, err, len(fromEmpty))
	}
}

func TestNextPagePathKeepsPathAndQuery(t *testing.T) {
	got, err := nextPagePath("http://internal-host:8000/api/workflows/?page=3&page_size=50")
	if err != nil || got != "/api/workflows/?page=3&page_size=50" {
		t.Fatalf("got %q, %v", got, err)
	}
}

// getAll must follow "next" across pages and stop on the last one.
func TestGetAllFollowsPagination(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "token t" {
			t.Errorf("missing auth header")
		}
		switch r.URL.RawQuery {
		case "":
			w.Write([]byte(`{"next":"` + server.URL + `/api/account/account-stakeholder-template/?page=2","results":[{"title":"one"}]}`))
		case "page=2":
			w.Write([]byte(`{"next":null,"results":[{"title":"two"}]}`))
		default:
			t.Errorf("unexpected query %q", r.URL.RawQuery)
		}
	}))
	defer server.Close()

	c, err := NewClient(&Config{BaseURL: server.URL, Token: "t"})
	if err != nil {
		t.Fatal(err)
	}
	templates, err := c.AccountStakeholderTemplate.GetAccountStakeholderTemplates()
	if err != nil {
		t.Fatal(err)
	}
	if len(templates) != 2 || templates[0].Title != "one" || templates[1].Title != "two" {
		t.Fatalf("got %+v", templates)
	}
}

// alert_urgency is nullable on the API: nil must serialise as null, not be dropped.
func TestHeartbeatMarshalsNullUrgency(t *testing.T) {
	body, err := json.Marshal(Heartbeat{Name: "hb", IntervalSeconds: 60})
	if err != nil {
		t.Fatal(err)
	}
	s := string(body)
	for _, want := range []string{`"alert_urgency":null`, `"is_enabled":false`, `"ping_type":0`, `"interval_seconds":60`} {
		if !strings.Contains(s, want) {
			t.Errorf("body %s lacks %s", s, want)
		}
	}
	for _, reject := range []string{"ping_url", "status", "unique_id", "service"} {
		if strings.Contains(s, `"`+reject+`"`) {
			t.Errorf("body %s must omit empty read-only %s", s, reject)
		}
	}
}

// The API rejects null for form_toggles and custom_fields; nil slices go out as [].
func TestIncidentTypeInputSendsEmptyArrays(t *testing.T) {
	body, err := json.Marshal(&IncidentTypeInput{Name: "Outage"})
	if err != nil {
		t.Fatal(err)
	}
	s := string(body)
	if !strings.Contains(s, `"form_toggles":[]`) || !strings.Contains(s, `"custom_fields":[]`) {
		t.Errorf("got %s", s)
	}
	if strings.Contains(s, "is_active") {
		t.Errorf("unset is_active must be omitted, got %s", s)
	}
}

func TestWorkflowInputSendsEmptyArrays(t *testing.T) {
	body, err := json.Marshal(&WorkflowInput{Name: "w", Teams: []WorkflowTeamInput{{UniqueID: "team"}}})
	if err != nil {
		t.Fatal(err)
	}
	s := string(body)
	if !strings.Contains(s, `"teams":[{"unique_id":"team","services":[]}]`) {
		t.Errorf("got %s", s)
	}
	body, _ = json.Marshal(&WorkflowInput{Name: "w"})
	if !strings.Contains(string(body), `"teams":[]`) {
		t.Errorf("nil teams must be sent as [], got %s", body)
	}
}

func TestAccountSettingOmitsUnsetFields(t *testing.T) {
	f := false
	body, err := json.Marshal(AccountSetting{EnableRegenKey: &f})
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"enable_regen_key":false}` {
		t.Errorf("got %s", body)
	}
}

func TestIncidentTypeReadShapeDecodes(t *testing.T) {
	body := `{"unique_id":"u","name":"Outage","is_default":true,"is_active":true,
	  "form_toggles":[{"impacted_services":{"enabled":true,"ui:widget":"x"}},{"responders":false}],
	  "custom_fields":{"type":"object","properties":{"custom_field_7":{"type":"string","title":"Region","default":"eu","position":1}},"required":["custom_field_7"]}}`
	var it IncidentType
	if err := json.Unmarshal([]byte(body), &it); err != nil {
		t.Fatal(err)
	}
	f := it.CustomFields.Properties["custom_field_7"]
	if f.Position != 1 || string(f.Default) != `"eu"` || it.CustomFields.Required[0] != "custom_field_7" {
		t.Errorf("got %+v", it.CustomFields)
	}
	if len(it.FormToggles) != 2 {
		t.Errorf("got %+v", it.FormToggles)
	}
}
