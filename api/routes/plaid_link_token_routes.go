package routes

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/controllers"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/gin-gonic/gin"
)

func PlaidLinkTokenRoutes(router *gin.Engine, protected *gin.RouterGroup, accessTokenDB interfaces.AccessTokenDB) {
	protected.GET("/plaid_link_token", controllers.CreateLinkToken)
}
