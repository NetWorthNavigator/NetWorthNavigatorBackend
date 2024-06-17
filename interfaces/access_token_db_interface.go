package interfaces

import "github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"

// AccessTokenService defines the interface for access token operations
type AccessTokenDB interface {
	CreateAccessToken(token models.PlaidAccessToken) error
	GetAccessToken(email string) (*models.PlaidAccessToken, error)
	DeleteAccessToken(id string) error
	UpdateAccessToken(token models.PlaidAccessToken) error
}
