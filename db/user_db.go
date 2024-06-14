package db

import (
	"log"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"gorm.io/gorm"
)

type UserDB struct {
	DBClient *gorm.DB
}

func NewUserDB(DBClient *gorm.DB) *UserDB {
	return &UserDB{DBClient: DBClient}
}

// CreateUser creates a new user in the database.
func (udb *UserDB) CreateUser(user models.User) error {
	result := udb.DBClient.Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return result.Error
	}
	return nil
}

// GetUser retrieves a user based on a given ID.
func (udb *UserDB) GetUser(email string) (*models.User, error) {
	var user models.User
	result := udb.DBClient.First(&user, "Email = ?", email)
	if result.Error != nil {
		log.Printf("Error retrieving user with Email %s: %v", email, result.Error)
		return nil, result.Error
	}
	return &user, nil
}

// DeleteUser deletes a user based on a given ID.
func (udb *UserDB) DeleteUser(id string) error {
	result := udb.DBClient.Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		log.Printf("Error deleting user with ID %s: %v", id, result.Error)
		return result.Error
	}
	return nil
}

// UpdateUser updates an existing user's details.
func (udb *UserDB) UpdateUser(user models.User) error {
	result := udb.DBClient.Save(&user)
	if result.Error != nil {
		log.Printf("Error updating user with Email %s: %v", user.Email, result.Error)
		return result.Error
	}
	return nil
}
