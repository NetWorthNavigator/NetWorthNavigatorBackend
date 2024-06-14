package routes

import (
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/controllers"
)

func SetupUserRoutes() {
	http.HandleFunc("/user", controllers.UserHandler)
}
