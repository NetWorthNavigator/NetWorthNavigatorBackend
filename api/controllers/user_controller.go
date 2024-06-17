package controllers

import (
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/interfaces"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUser handles GET requests for the /user route
func GetUser(c *gin.Context, userDB interfaces.UserDB, email string) {
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Email is required"})
		return
	}

	user, err := userDB.GetUser(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PostUser handles POST requests for the /user route
func PostUser(c *gin.Context, userDB interfaces.UserDB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	normalizedEmail := utils.NormalizeEmail(user.Email)
	user.Email = normalizedEmail

	_, err := userDB.GetUser(user.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := userDB.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// PutUser handles PUT requests for the /user route
func PutUser(c *gin.Context, userDB interfaces.UserDB, email string) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if err := userDB.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// DeleteUser handles DELETE requests for the /user route
func DeleteUser(c *gin.Context, userDB interfaces.UserDB, email string) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	err := userDB.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
