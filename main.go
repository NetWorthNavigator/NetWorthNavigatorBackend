package main

import (
	"fmt"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/db"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := constants.DATABASE_URL
	DBClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database", err)
		return
	}
	InitDB(DBClient)

	accessTokenDB := db.NewAccessTokenDB(DBClient)
	userDB := db.NewUserDB(DBClient)
	PlaidAccountDB := db.NewPlaidAccountDB(DBClient)
	itemDB := db.NewItemDB(DBClient)

	router := gin.Default()

	SetupRouter(router, accessTokenDB, userDB, PlaidAccountDB, itemDB)

	router.Run(":8080")
}
