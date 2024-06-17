package routes

import (
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/controllers"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/gin-gonic/gin"
)

func LoginRoutes(userDB interfaces.UserDB, r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		controllers.LoginController(c, userDB)
	})
}
