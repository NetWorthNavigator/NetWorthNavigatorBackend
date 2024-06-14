package main

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/routes"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/db"
)

func SetupRouter(accessTokenDB *db.AccessTokenDB, userDB *db.UserDB) {

	routes.SetupLinkRoutes(accessTokenDB)

	routes.SetupUserRoutes(userDB)

	//http.HandleFunc("/test", routes.Test)
}

//5T4DVT
