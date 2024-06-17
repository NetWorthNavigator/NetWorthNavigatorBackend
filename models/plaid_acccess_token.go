package models

import "gorm.io/gorm"

type PlaidAccessToken struct {
	gorm.Model
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	ItemID      string `json:"item_id"`
	RequestID   string `json:"request_id"`
}
