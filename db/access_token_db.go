package db

import (
	"log"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"gorm.io/gorm"
)

type AccessTokenDB struct {
	DBClient *gorm.DB
}

func NewAccessTokenDB(DBClient *gorm.DB) *AccessTokenDB {
	return &AccessTokenDB{DBClient: DBClient}
}

func (atdb *AccessTokenDB) CreateAccessToken(token models.PlaidAccessToken) error {
	result := atdb.DBClient.Create(&token)
	if result.Error != nil {
		log.Printf("Error creating access token: %v", result.Error)
		return result.Error
	}
	return nil
}

// GetAccessToken retrieves an access token based on a given ID.
func (atdb *AccessTokenDB) GetAccessToken(email string) (*models.PlaidAccessToken, error) {
	var token models.PlaidAccessToken
	result := atdb.DBClient.First(&token, "email = ?", email)
	if result.Error != nil {
		log.Printf("Error retrieving access token with ID %s: %v", email, result.Error)
		return nil, result.Error
	}
	return &token, nil
}

// DeleteAccessToken deletes an access token based on a given ID.
func (atdb *AccessTokenDB) DeleteAccessToken(id string) error {
	result := atdb.DBClient.Delete(&models.PlaidAccessToken{}, "id = ?", id)
	if result.Error != nil {
		log.Printf("Error deleting access token with ID %s: %v", id, result.Error)
		return result.Error
	}
	return nil
}

// UpdateAccessToken updates an existing access token's details.
func (atdb *AccessTokenDB) UpdateAccessToken(token models.PlaidAccessToken) error {
	result := atdb.DBClient.Save(&token)
	if result.Error != nil {
		log.Printf("Error updating access token with ID %d: %v", token.ID, result.Error)
		return result.Error
	}
	return nil
}
