package models

type PlaidSyncRequest struct {
	ClientID     string    `json:"client_id"`
	Secret       string    `json:"secret"`
	AccessToken   string    `json:"access_token"`
}
