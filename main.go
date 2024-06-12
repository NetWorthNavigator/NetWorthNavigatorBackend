package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/api/routes"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
	"github.com/supabase-community/supabase-go"
)

func main() {

	client, err := supabase.NewClient(constants.DATABASE_URL, constants.DATABASE_API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("test").Select("*", "exact", false).Execute()
	log.Print(data, count)

	http.HandleFunc("/create_link_token", routes.CreateLinkTokenHandler)
	http.HandleFunc("/create_access_token", routes.CreateAccessTokenHandler)
	http.HandleFunc("/test", routes.Test)
	http.ListenAndServe(":8080", nil)
}
