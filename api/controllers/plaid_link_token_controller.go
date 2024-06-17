package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"github.com/gin-gonic/gin"
)

func CreateLinkToken(c *gin.Context) {
	url := "https://sandbox.plaid.com/link/token/create"

	data := models.PlaidLinkTokenRequest{
		ClientID:     constants.PLAID_CLIENT_ID,
		Secret:       constants.PLAID_SECRET,
		ClientName:   constants.CLIENT_NAME,
		CountryCodes: []string{"US"},
		Language:     "en",
		Products:     []string{"auth", "transactions"},
		User:         models.PlaidUser{ClientUserID: "user-id"},
	}
	payload, _ := json.Marshal(data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	c.JSON(http.StatusOK, result)
}
