package main

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"gorm.io/gorm"
)

func InitDB(DBClient *gorm.DB) {
	DBClient.AutoMigrate(&models.PlaidAccessToken{})

}
