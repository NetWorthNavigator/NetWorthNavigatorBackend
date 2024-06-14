package main

import (
	"fmt"
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/db"

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

	SetupRouter(accessTokenDB)
	http.ListenAndServe(":8080", nil)
}
