package routes

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/controllers"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/gin-gonic/gin"
)

func AccountRoutes(router *gin.Engine, protected *gin.RouterGroup, userDB interfaces.UserDB, accessTokenDB interfaces.AccessTokenDB, plaidAccountDB interfaces.PlaidAccountDB) {
	protected.GET("/accounts", func(c *gin.Context) {
		controllers.GetAccounts(c, userDB, plaidAccountDB, accessTokenDB,c.MustGet("email").(string))
	})
	protected.POST("/account", func(c *gin.Context) {
		controllers.PostAccount(c, userDB,accessTokenDB, c.MustGet("email").(string))
	})
	protected.PUT("/account", func(c *gin.Context) {
		controllers.PutAccount(c, userDB,accessTokenDB, c.MustGet("email").(string))
	})
}
