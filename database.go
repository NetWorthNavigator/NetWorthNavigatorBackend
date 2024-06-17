package main

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"gorm.io/gorm"
)

func InitDB(DBClient *gorm.DB) {
	DBClient.AutoMigrate(&models.PlaidAccessToken{})
	DBClient.AutoMigrate(&models.User{})

	DBClient.AutoMigrate(&models.PlaidAccount{})
	//DBClient.AutoMigrate(&models.Institution{})
	DBClient.AutoMigrate(&models.Item{})
	//DBClient.AutoMigrate(&models.Transaction{})

}
