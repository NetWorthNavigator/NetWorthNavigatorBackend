package routes

import (
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/controllers"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/db"
)

func SetupUserRoutes(userDB *db.UserDB) {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		controllers.UserHandler(userDB, w, r)
	})
}
