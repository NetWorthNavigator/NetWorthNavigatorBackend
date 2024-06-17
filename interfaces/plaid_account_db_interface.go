package interfaces

import "github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"

// IPlaidAccountDB defines the interface for PlaidAccountDB operations
type PlaidAccountDB interface {
    CreatePlaidAccount(account models.PlaidAccount) error
    GetPlaidAccount(email string) (*models.PlaidAccount, error)
    GetPlaidAccounts(email string) ([]models.PlaidAccount, error)
    DeletePlaidAccount(id string) error
    UpdatePlaidAccount(account models.PlaidAccount) error
}