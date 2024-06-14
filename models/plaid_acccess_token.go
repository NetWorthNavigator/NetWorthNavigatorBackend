package models

import "gorm.io/gorm"

type PlaidAccessToken struct {
	gorm.Model
	UserID      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	ItemID      string `json:"item_id"`
	RequestID   string `json:"request_id"`
}
