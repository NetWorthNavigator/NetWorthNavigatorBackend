package main

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/routes"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, accessTokenDB interfaces.AccessTokenDB, userDB interfaces.UserDB, plaidAccountDB interfaces.PlaidAccountDB, itemDB interfaces.ItemDB) {
	// Middleware that applies to all routes
	router.Use(middleware.CORSMiddleware())

	// Group for protected routes with JWT authentication
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// Public routes
	routes.LoginRoutes(userDB, router)

	// Protected routes
	routes.UserRoutes(router, protected, userDB)
	routes.PlaidAccessTokenRoutes(router, protected, accessTokenDB, plaidAccountDB, itemDB)
	routes.PlaidLinkTokenRoutes(router, protected, accessTokenDB)
	routes.AccountRoutes(router, protected, userDB, accessTokenDB, plaidAccountDB)
}
