package db

import (
	"log"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"gorm.io/gorm"
)

type ItemDB struct {
    DBClient *gorm.DB
}

func NewItemDB(DBClient *gorm.DB) *ItemDB {
    return &ItemDB{DBClient: DBClient}
}

// CreateItem creates a new Item record.
func (idb *ItemDB) CreateItem(item models.Item) error {
    // Check if the item already exists with the same ItemID
    var existingItem models.Item
    if err := idb.DBClient.Where("item_id = ?", item.ItemID).First(&existingItem).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // Item does not exist, proceed with creation
            result := idb.DBClient.Create(&item)
            if result.Error != nil {
                log.Printf("Error creating Item: %v", result.Error)
                return result.Error
            }
            return nil
        }
        // Other error occurred
        return err
    }
    // Item with the same ItemID already exists, return nil or a custom error
    log.Printf("Item with ItemID %s already exists", item.ItemID)
    return nil // or return an error indicating the item already exists
}

// GetItems retrieves all items.
func (idb *ItemDB) GetItems() ([]models.Item, error) {
    var items []models.Item
    result := idb.DBClient.Find(&items)
    if result.Error != nil {
        log.Printf("Error retrieving items: %v", result.Error)
        return nil, result.Error
    }
    return items, nil
}

// GetItem retrieves a single item based on ItemID.
func (idb *ItemDB) GetItem(itemID string) (*models.Item, error) {
    var item models.Item
    result := idb.DBClient.Where("item_id = ?", itemID).First(&item)
    if result.Error != nil {
        log.Printf("Error retrieving item with ItemID %s: %v", itemID, result.Error)
        return nil, result.Error
    }
    return &item, nil
}

// DeleteItem deletes an item based on ItemID.
func (idb *ItemDB) DeleteItem(itemID string) error {
    result := idb.DBClient.Delete(&models.Item{}, "item_id = ?", itemID)
    if result.Error != nil {
        log.Printf("Error deleting item with ItemID %s: %v", itemID, result.Error)
        return result.Error
    }
    return nil
}

// UpdateItem updates an existing item's details.
func (idb *ItemDB) UpdateItem(item models.Item) error {
    result := idb.DBClient.Save(&item)
    if result.Error != nil {
        log.Printf("Error updating item with ItemID %s: %v", item.ItemID, result.Error)
        return result.Error
    }
    return nil
}