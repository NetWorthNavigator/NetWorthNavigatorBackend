package routes

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/controllers"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/gin-gonic/gin"
)

func PlaidAccessTokenRoutes(router *gin.Engine, protected *gin.RouterGroup, accessTokenDB interfaces.AccessTokenDB, plaidAccountDB interfaces.PlaidAccountDB, itemDB interfaces.ItemDB) {
	protected.POST("/plaid_access_token", func(c *gin.Context) {
		controllers.CreateAccessToken(c, accessTokenDB, plaidAccountDB, itemDB, c.MustGet("email").(string))
	})
}
