package db

import (
	"log"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"gorm.io/gorm"
)

type PlaidAccountDB struct {
	DBClient *gorm.DB
}

func NewPlaidAccountDB(DBClient *gorm.DB) *PlaidAccountDB {
	return &PlaidAccountDB{DBClient: DBClient}
}

func (padb *PlaidAccountDB) CreatePlaidAccount(account models.PlaidAccount) error {
    // Check if the account already exists with the same email
    var existingAccount models.PlaidAccount
    if err := padb.DBClient.Where("item_id = ? AND account_id = ? AND email = ?", account.ItemID, account.AccountID, account.Email).First(&existingAccount).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // Account does not exist, proceed with creation
            result := padb.DBClient.Create(&account)
            if result.Error != nil {
                log.Printf("Error creating Plaid account: %v", result.Error)
                return result.Error
            }
            return nil
        }
        // Other error occurred
        return err
    }
    // Account with the same item_id, account_id, and email already exists, return nil or a custom error
    log.Printf("Account with item_id %s, account_id %s, and email %s already exists", account.ItemID, account.AccountID, account.Email)
    return nil // or return an error indicating the account already exists
}


// GetPlaidAccount retrieves an access token based on a given ID.
func (padb *PlaidAccountDB) GetPlaidAccounts(email string) ([]models.PlaidAccount, error) {
	var accounts []models.PlaidAccount
    result := padb.DBClient.Where("email = ?", email).Find(&accounts)
	if result.Error != nil {
		log.Printf("Error retrieving access account with Email %s: %v", email, result.Error)
		return nil, result.Error
	}
	return accounts, nil
}



// GetPlaidAccount retrieves an access token based on a given ID.
func (padb *PlaidAccountDB) GetPlaidAccount(email string) (*models.PlaidAccount, error) {
	var account models.PlaidAccount
	result := padb.DBClient.First(&account, "email = ?", email)
	if result.Error != nil {
		log.Printf("Error retrieving access account with ID %s: %v", email, result.Error)
		return nil, result.Error
	}
	return &account, nil
}

// DeletePlaidAccount deletes an access token based on a given ID.
func (padb *PlaidAccountDB) DeletePlaidAccount(id string) error {
	result := padb.DBClient.Delete(&models.PlaidAccount{}, "id = ?", id)
	if result.Error != nil {
		log.Printf("Error deleting access account with ID %s: %v", id, result.Error)
		return result.Error
	}
	return nil
}

// UpdatePlaidAccount updates an existing access token's details.
func (padb *PlaidAccountDB) UpdatePlaidAccount(account models.PlaidAccount) error {
	result := padb.DBClient.Save(&account)
	if result.Error != nil {
		log.Printf("Error updating access account with ID %v: %v", account.AccountID, result.Error)
		return result.Error
	}
	return nil
}
