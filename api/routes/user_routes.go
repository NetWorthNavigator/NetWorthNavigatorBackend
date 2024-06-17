package routes

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/controllers"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, protected *gin.RouterGroup, userDB interfaces.UserDB) {
	protected.GET("/user", func(c *gin.Context) {
		controllers.GetUser(c, userDB, c.MustGet("email").(string))
	})
	router.POST("/user", func(c *gin.Context) {
		controllers.PostUser(c, userDB)
	})
	protected.PUT("/user", func(c *gin.Context) {
		controllers.PutUser(c, userDB, c.MustGet("email").(string))
	})
	protected.DELETE("/user", func(c *gin.Context) {
		controllers.DeleteUser(c, userDB, c.MustGet("email").(string))
	})
}
