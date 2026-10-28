package client

import "encoding/json"

type AccountSettingService service

// AccountSetting is the account's singleton settings; unset pointers are not sent.
type AccountSetting struct {
	Account                    string `json:"account,omitempty"`
	EnableDeleteIncident       *bool  `json:"enable_delete_incident,omitempty"`
	EnableTeamManagerInvite    *bool  `json:"enable_team_manager_invite,omitempty"`
	EnableRegenKey             *bool  `json:"enable_regen_key,omitempty"`
	EnableFirstMTTACalculation *bool  `json:"enable_first_mtta_calculation,omitempty"`
}

const accountSettingPath = "/api/account/account-setting/"

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
