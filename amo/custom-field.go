package amo

// CustomField 	Contains custom fields for different entities
type CustomField struct {
	ID       string            `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Code     string            `json:"code,omitempty"`
	Multiple string            `json:"multiple,omitempty"`
	TypeID   string            `json:"type_id,omitempty"`
	Disabled string            `json:"disabled,omitempty"`
	Sort     int               `json:"sort,omitempty"`
	Enums    map[string]string `json:"enums,omitempty"`
}
