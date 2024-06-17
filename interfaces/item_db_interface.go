package interfaces

import "github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"

type ItemDB interface {
    CreateItem(item models.Item) error
    GetItems() ([]models.Item, error)
    GetItem(itemID string) (*models.Item, error)
    DeleteItem(itemID string) error
    UpdateItem(item models.Item) error
}