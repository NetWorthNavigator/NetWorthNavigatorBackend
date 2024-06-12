package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
)

type User struct {
	ClientUserID string `json:"client_user_id"`
}
type PlaidLinkTokenRequest struct {
	ClientID     string   `json:"client_id"`
	Secret       string   `json:"secret"`
	ClientName   string   `json:"client_name"`
	CountryCodes []string `json:"country_codes"`
	Language     string   `json:"language"`
	Products     []string `json:"products"`
	User         User     `json:"user"`
}

func CreateLinkTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	url := "https://sandbox.plaid.com/link/token/create"

	data := PlaidLinkTokenRequest{
		ClientID:     constants.PLAID_CLIENT_ID,
		Secret:       constants.PLAID_SECRET,
		ClientName:   constants.CLIENT_NAME,            // replace with your client name
		CountryCodes: []string{"US"},                   // replace with your country codes
		Language:     "en",                             // replace with your language
		Products:     []string{"auth", "transactions"}, // replace with your products
		User:         User{ClientUserID: "user-id"},    // replace with your user id
	}
	payload, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

type PublicToken struct {
	PublicToken string `json:"public_token"`
}

type ExchangeRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	PublicToken string `json:"public_token"`
}

func CreateAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Preflight request. Reply successfully:
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var token PublicToken
	err = json.Unmarshal(body, &token)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	exchangeRequest := ExchangeRequest{
		ClientID:    constants.PLAID_CLIENT_ID,
		Secret:      constants.PLAID_SECRET,
		PublicToken: token.PublicToken,
	}
	exchangeRequestBody, err := json.Marshal(exchangeRequest)
	if err != nil {
		http.Error(w, "Error creating exchange request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("https://sandbox.plaid.com/item/public_token/exchange", "application/json", bytes.NewBuffer(exchangeRequestBody))
	if err != nil {
		http.Error(w, "Error exchanging public token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading exchange response", http.StatusInternalServerError)
		return
	}

	res := responseBody
	log.Print(string(res))
}
