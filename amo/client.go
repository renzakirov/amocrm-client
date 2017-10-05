package amo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client AmoCRM API Client
type Client struct {
	userLogin         string
	userHash          string
	Timezone          string
	accountWebAddress *url.URL
	rateLimiter       RateLimiter
}

// NewClient creates and initializes AmoCRM API.
// accountWebAddress is your AmoCRM account address like https://your-account.amocrm.com
func NewClient(accountWebAddress string, rateLimiter RateLimiter) (*Client, error) {
	c := &Client{rateLimiter: rateLimiter}
	if c.rateLimiter == nil {
		c.rateLimiter = defaultRTLimiter
	}
	var err error
	c.accountWebAddress, err = url.Parse(accountWebAddress)
	return c, err
}

// AuthResponse information about Amo' user and account.
type AuthResponse struct {
	Response struct {
		Auth     bool `json:"auth"`
		Accounts []struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Subdomain string `json:"subdomain"`
			Language  string `json:"language"`
			Timezone  string `json:"timezone"`
		} `json:"accounts"`
		ServerTime int `json:"server_time"`
	} `json:"response"`
}

// Authorize In case of successful authorization,
// this method returns the user’s session ID,
// which must be used while accessing to other API methods.
// Saves user's login and hash into the client for next usage.
func (c *Client) Authorize(userLogin, userHash string) (*AuthResponse, error) {
	c.userHash = userHash
	c.userLogin = userLogin
	reqURL := *c.accountWebAddress
	reqURL.Path = "/private/api/auth.php"
	values := reqURL.Query()
	values.Set("type", "json")
	reqURL.RawQuery = values.Encode()
	c.rateLimiter.WaitForRequest()
	resp, err := http.PostForm(reqURL.String(),
		url.Values{"USER_LOGIN": {c.userLogin}, "USER_HASH": {c.userHash}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var authResponse AuthResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, err
	}
	if len(authResponse.Response.Accounts) > 0 {
		c.Timezone = authResponse.Response.Accounts[0].Timezone
	}
	return &authResponse, nil
}
