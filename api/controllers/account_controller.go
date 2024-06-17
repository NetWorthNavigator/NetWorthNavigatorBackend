package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
)


func FetchAndSaveAccounts(accessToken string) ([]models.Account, *models.Item, error) {
    url := "https://sandbox.plaid.com/accounts/get"
    data := models.PlaidSyncRequest{
        ClientID:    constants.PLAID_CLIENT_ID,
        Secret:      constants.PLAID_SECRET,
        AccessToken: accessToken,
    }
    payload, _ := json.Marshal(data)

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        return nil, nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, nil, err
    }
    log.Print(string(body))
    var apiResponse models.ApiResponse
    if err := json.Unmarshal(body, &apiResponse); err != nil {
        return nil, nil, err
    }
    log.Print(apiResponse.Item)


    return apiResponse.Accounts, &apiResponse.Item, nil
}


// Modified method to convert a slice of Account objects to a slice of PlaidAccount objects.
func ConvertToPlaidAccounts(accounts []models.Account) []models.PlaidAccount {
    var plaidAccounts []models.PlaidAccount
    for _, a := range accounts {
        plaidAccount := models.PlaidAccount{
            AccountID:                  a.AccountID,
            Mask:                       a.Mask,
            Name:                       a.Name,
            OfficialName:               a.OfficialName,
            Subtype:                    a.Subtype,
            Type:                       a.Type,
            BalanceAvailable:           strconv.FormatFloat(a.Balances.Available, 'f', 2, 64),
            BalanceCurrent:             strconv.FormatFloat(a.Balances.Current, 'f', 2, 64),
            BalanceISOCurrencyCode:     a.Balances.IsoCurrencyCode,
            BalanceLimit:               float64ToString(a.Balances.Limit),
            BalanceUnofficialCurrencyCode: stringPointerToString(a.Balances.UnofficialCurrencyCode),
        }
        plaidAccounts = append(plaidAccounts, plaidAccount)
    }
    return plaidAccounts
}

// Helper function to convert *float64 to string safely.
func float64ToString(f *float64) string {
    if f != nil {
        return strconv.FormatFloat(*f, 'f', 2, 64)
    }
    return ""
}

// Helper function to convert *string to string safely.
func stringPointerToString(s *string) string {
    if s != nil {
        return *s
    }
    return ""
}


// GetAccount retrieves a user's account information
func GetAccounts(c *gin.Context, userDB interfaces.UserDB, plaidAccountDB interfaces.PlaidAccountDB, accessTokenDB interfaces.AccessTokenDB, email string) {
    // Extract account ID from request parameters
    if email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
        return
    }

    // Retrieve the account using the account ID
    accounts, err := plaidAccountDB.GetPlaidAccounts(email)
    if err != nil {
        // Assuming err means the account was not found; adjust as necessary
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve account"})
        return
    }

    // Return the account in the response
    c.JSON(http.StatusOK, accounts)
}

// PostAccount creates a new user account
func PostAccount(c *gin.Context, userDB interfaces.UserDB, accessTokenDB interfaces.AccessTokenDB, email string) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := userDB.CreateUser(newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// PutAccount updates an existing user's account information
func PutAccount(c *gin.Context, userDB interfaces.UserDB, accessTokenDB interfaces.AccessTokenDB, email string) {
    var updatedUser models.User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := userDB.UpdateUser(updatedUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}