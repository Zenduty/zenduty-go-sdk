package client

import "encoding/json"

type AccountSettingService service

// AccountSetting is the account's singleton settings object. Fields are
// pointers so an update only sends the settings that were set.
type AccountSetting struct {
	Account                    string `json:"account,omitempty"`
	EnableDeleteIncident       *bool  `json:"enable_delete_incident,omitempty"`
	EnableTeamManagerInvite    *bool  `json:"enable_team_manager_invite,omitempty"`
	EnableRegenKey             *bool  `json:"enable_regen_key,omitempty"`
	EnableFirstMTTACalculation *bool  `json:"enable_first_mtta_calculation,omitempty"`
}

const accountSettingPath = "/api/account/account-setting/"

// GetAccountSetting reads the account settings. The backend creates the
// settings row with its defaults on first read.
func (c *AccountSettingService) GetAccountSetting() (*AccountSetting, error) {
	body, err := c.client.newRequestDo("GET", accountSettingPath, nil)
	if err != nil {
		return nil, err
	}
	var s AccountSetting
	if err := json.Unmarshal(body.BodyBytes, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// UpdateAccountSetting patches the settings that are set on setting.
func (c *AccountSettingService) UpdateAccountSetting(setting *AccountSetting) (*AccountSetting, error) {
	body, err := c.client.newRequestDo("PATCH", accountSettingPath, setting)
	if err != nil {
		return nil, err
	}
	var s AccountSetting
	if err := json.Unmarshal(body.BodyBytes, &s); err != nil {
		return nil, err
	}
	return &s, nil
}
