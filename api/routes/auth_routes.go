package routes

import (
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/handlers"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/db"
)

func AuthHandler(userDB *db.UserDB, w http.ResponseWriter, r *http.Request) {

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(userDB, w, r)
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(userDB, w, r)
	})
}
