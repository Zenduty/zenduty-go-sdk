package client

import (
	"encoding/json"
	"strings"
	"testing"
)

// Default values must still be serialised when settings are present.
func TestEscalationPolicyMarshalsAssignmentSettings(t *testing.T) {
	policy := EscalationPolicy{
		Name:               "esp",
		AssignmentSettings: &EscalationPolicyAssignmentSettings{AssigneeStrategy: AssigneeStrategyAny},
	}
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	var payload struct {
		AssignmentSettings map[string]interface{} `json:"assignment_settings"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatal(err)
	}
	got := payload.AssignmentSettings
	if got["assignee_strategy"] != float64(AssigneeStrategyAny) || got["call_rr_assignee"] != false {
		t.Errorf("assignment_settings = %v, want assignee_strategy 1 and call_rr_assignee false", got)
	}
	// read-only fields are omitted when empty
	if len(got) != 2 {
		t.Errorf("assignment_settings = %v, want only assignee_strategy and call_rr_assignee", got)
	}

	body, err = json.Marshal(EscalationPolicy{Name: "esp"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "assignment_settings") {
		t.Errorf("nil settings must be omitted, got %s", body)
	}
}

func TestEscalationPolicyUnmarshalsAssignmentSettings(t *testing.T) {
	var withSettings EscalationPolicy
	err := json.Unmarshal([]byte(`{"name":"esp","assignment_settings":{"unique_id":"abc","assignee_strategy":2,"assignment_index":3,"call_rr_assignee":true}}`), &withSettings)
	if err != nil {
		t.Fatal(err)
	}
	got := withSettings.AssignmentSettings
	if got == nil || got.AssigneeStrategy != AssigneeStrategyRoundRobin || !got.CallRRAssignee || got.AssignmentIndex != 3 || got.UniqueID != "abc" {
		t.Errorf("unexpected settings %+v", got)
	}

	// response for a policy without saved settings
	var defaults EscalationPolicy
	err = json.Unmarshal([]byte(`{"name":"esp","assignment_settings":{"unique_id":null,"assignee_strategy":1,"assignment_index":0,"call_rr_assignee":false}}`), &defaults)
	if err != nil {
		t.Fatal(err)
	}
	if defaults.AssignmentSettings == nil || defaults.AssignmentSettings.AssigneeStrategy != AssigneeStrategyAny || defaults.AssignmentSettings.CallRRAssignee {
		t.Errorf("unexpected default settings %+v", defaults.AssignmentSettings)
	}
}
