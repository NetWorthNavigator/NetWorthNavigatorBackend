package routes

import (
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/controllers"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/db"
)

func SetupLinkRoutes(accessTokenDB *db.AccessTokenDB) {
	http.HandleFunc("/create_link_token", controllers.CreateLinkTokenHandler)
	http.HandleFunc("/create_access_token", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateAccessTokenHandler(accessTokenDB, w, r)
	})
}
