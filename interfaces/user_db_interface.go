package interfaces

import "github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"

type UserDB interface {
	CreateUser(user models.User) error
	GetUser(email string) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(user models.User) error
}
