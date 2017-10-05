package amo

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
