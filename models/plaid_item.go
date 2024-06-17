package models

import (
	"time"

	"gorm.io/gorm"
)



type PlaidItemResponse struct {
    Item Item `json:"item"`
    // Include other fields from the response if necessary
}


type PlaidItemRequest struct {
	ClientID    string `json:"client_id"`
	Secret 	string `json:"secret"`
	AccessToken string `json:"access_token"`

}



type Item struct {
    gorm.Model
    //AvailableProducts       []string `gorm:"type:text" json:"available_products"`
    //BilledProducts          []string `gorm:"type:text" json:"billed_products"`
    Error                   *string  `gorm:"type:text" json:"error"`
    InstitutionID           string   `gorm:"type:varchar(255)" json:"institution_id"`
    ItemID                  string   `gorm:"uniqueIndex;type:varchar(255)" json:"item_id"`
    UpdateType              string   `gorm:"type:varchar(255)" json:"update_type"`
    Webhook                 string   `gorm:"type:text" json:"webhook"`
    ConsentExpirationTime   *time.Time `gorm:"type:timestamp" json:"consent_expiration_time"`
}

type Status struct {
    ItemID                  string `gorm:"uniqueIndex;type:varchar(255);not null" json:"item_id"`
    LastSuccessfulUpdate    time.Time `gorm:"type:timestamp" json:"last_successful_update"`
    LastFailedUpdate        time.Time `gorm:"type:timestamp" json:"last_failed_update"`
    LastWebhookSentAt       time.Time `gorm:"type:timestamp" json:"last_webhook_sent_at"`
    LastWebhookCodeSent     string    `gorm:"type:varchar(255)" json:"last_webhook_code_sent"`
}

