package amo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Account an information on the authorized account.
type Account struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Subdomain    string      `json:"subdomain"`
	Currency     string      `json:"currency"`
	Timezone     string      `json:"timezone"`
	Language     string      `json:"language"`
	CurrentUser  int         `json:"current_user"`
	Version      int         `json:"version"`
	DatePattern  string      `json:"date_pattern"`
	DateFormat   interface{} `json:"date_format"`
	TimeFormat   interface{} `json:"time_format"`
	Country      interface{} `json:"country"`
	CustomFields struct {
		Contacts  []CustomField `json:"contacts"`
		Companies []CustomField `json:"companies"`
	} `json:"custom_fields"`
	TimezoneOffset string `json:"timezoneoffset"`
}

// GetAccount obtains all necessary information about the account: name, paid period, account’s users and their privileges,
// guides on custom contact and lead fields, lead status guide, event types guide, task types guide and other account parameters.
// https://developers.amocrm.com/rest_api/accounts_current.php
func (c *Client) GetAccount(freeUsers bool) (*Account, error) {
	type _AccountResponse struct {
		Response struct {
			Account    Account `json:"account"`
			ServerTime int     `json:"server_time"`
		} `json:"response"`
	}
	reqURL := *c.accountWebAddress
	reqURL.Path = "/private/api/v2/json/accounts/current"
	values := reqURL.Query()
	values.Set("USER_LOGIN", c.userLogin)
	values.Set("USER_HASH", c.userHash)
	if freeUsers {
		values.Set("free_users", "Y")
	}
	reqURL.RawQuery = values.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	c.rateLimiter.WaitForRequest()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var acc _AccountResponse
	err = json.Unmarshal(body, &acc)
	if err != nil {
		return nil, err
	}
	return &acc.Response.Account, nil
}
