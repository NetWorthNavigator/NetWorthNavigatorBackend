package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: fmt.Sprintf("URL: %s", DATABASE_URL)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}
