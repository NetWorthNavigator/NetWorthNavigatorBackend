package controllers

import (
	"net/http"
	"time"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWT(email string) (string, error) {
	claims := &jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(constants.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func LoginController(c *gin.Context, userDB interfaces.UserDB) {
	var test models.User

	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := userDB.GetUser(test.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Login"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(test.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Login"})
		return
	}

	token, err := GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
