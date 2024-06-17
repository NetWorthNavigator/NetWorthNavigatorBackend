package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"github.com/gin-gonic/gin"
)

func fetchItem(accessToken string) (models.Item, error) {

	url := "https://sandbox.plaid.com/item/get"
	data := models.PlaidItemRequest{
		ClientID:    constants.PLAID_CLIENT_ID,
		Secret:      constants.PLAID_SECRET,
		AccessToken: accessToken,
	}
	payload, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return models.Item{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Item{}, err
	}
	log.Print(string(body))

	var plaidResponse models.PlaidItemResponse
	if err := json.Unmarshal(body, &plaidResponse); err != nil {
		return models.Item{}, err
	}

	return plaidResponse.Item, nil

}

// Define a struct for the request payload
type RequestPayload struct {
	PublicToken string   `json:"public_token"`
	Metadata    Metadata `json:"metadata"` // Adjust this according to the actual structure of your metadata
}

// Metadata represents the metadata information in the request payload.
type Metadata struct {
	Status             *string     `json:"status"`
	LinkSessionID      string      `json:"link_session_id"`
	Account            Account     `json:"account"`
	ClassType          *string     `json:"class_type"`
	ID                 string      `json:"id"`
	Mask               string      `json:"mask"`
	Name               string      `json:"name"`
	Subtype            string      `json:"subtype"`
	Type               string      `json:"type"`
	VerificationStatus *string     `json:"verification_status"`
	AccountID          string      `json:"account_id"`
	Accounts           []Account   `json:"accounts"`
	Institution        Institution `json:"institution"`
	InstitutionID      string      `json:"institution_id"`
	PublicToken        string      `json:"public_token"`
	TransferStatus     *string     `json:"transfer_status"`
	Wallet             *string     `json:"wallet"`
}

// Account represents the account information.
type Account struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Mask               string  `json:"mask"`
	Type               string  `json:"type"`
	Subtype            string  `json:"subtype"`
	VerificationStatus *string `json:"verification_status"`
}

// Institution represents the institution information.
type Institution struct {
	Name          string `json:"name"`
	InstitutionID string `json:"institution_id"`
}

func CreateAccessToken(c *gin.Context, accessTokenDB interfaces.AccessTokenDB, plaidAccountDB interfaces.PlaidAccountDB, itemDB interfaces.ItemDB, email string) {
	var payload RequestPayload

	// Parse the JSON body into the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing request body"})
		return
	}

	curAccounts, err := plaidAccountDB.GetPlaidAccounts(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving accounts"})
		return
	}

	for _, account := range curAccounts {
		if account.InstitutionID == payload.Metadata.Institution.InstitutionID {
			c.JSON(http.StatusConflict, gin.H{"error": "Account already exists"})
			return
		}
	}

	exchangeRequest := models.PlaidExchangeRequest{
		ClientID:    constants.PLAID_CLIENT_ID,
		Secret:      constants.PLAID_SECRET,
		PublicToken: payload.PublicToken,
	}
	log.Print("hi", payload)
	exchangeRequestBody, err := json.Marshal(exchangeRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating exchange request"})
		return
	}

	resp, err := http.Post("https://sandbox.plaid.com/item/public_token/exchange", "application/json", bytes.NewBuffer(exchangeRequestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error exchanging public token"})
		return
	}
	defer resp.Body.Close()

	var accessToken models.PlaidAccessToken

	err = json.NewDecoder(resp.Body).Decode(&accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading exchange response"})
		return
	}

	// Assuming accessTokenDB is accessible, e.g., via c.MustGet or similar
	accessToken.Email = email
	accessTokenDB.CreateAccessToken(accessToken)

	// Create Accounts based on the access token
	apiAccounts, item, err := FetchAndSaveAccounts(accessToken.AccessToken)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	// Save the item
	itemDB.CreateItem(*item)

	accounts := ConvertToPlaidAccounts(apiAccounts)
	for _, account := range accounts {
		account.Email = email
		account.AccessToken = accessToken.AccessToken
		account.ItemID = accessToken.ItemID
		account.InstitutionID = payload.Metadata.Institution.InstitutionID
		err := plaidAccountDB.CreatePlaidAccount(account)
		if err != nil {
			log.Fatalf("Failed to save account: %v", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access token created successfully"})
}
