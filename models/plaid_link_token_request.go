package models

type PlaidLinkTokenRequest struct {
	ClientID     string    `json:"client_id"`
	Secret       string    `json:"secret"`
	ClientName   string    `json:"client_name"`
	CountryCodes []string  `json:"country_codes"`
	Language     string    `json:"language"`
	Products     []string  `json:"products"`
	User         PlaidUser `json:"user"`
}
