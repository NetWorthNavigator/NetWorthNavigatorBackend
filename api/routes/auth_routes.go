package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"

	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	//InMemoryUserDB[user.Username] = user

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s successfully registered", user.Email)
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//var user User
	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// storedUser, ok := InMemoryUserDB[user.Username]
	// if !ok {
	// 	http.Error(w, "User not found", http.StatusUnauthorized)
	// 	return
	// }

	// err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	// if err != nil {
	// 	http.Error(w, "Invalid password", http.StatusUnauthorized)
	// 	return
	// }

	// expirationTime := time.Now().Add(5 * time.Minute)
	// claims := &Claims{
	// 	Username: user.Username,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: expirationTime.Unix(),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString(constants.JWT_SECRET)
	// if err != nil {
	// 	http.Error(w, "Failed to generate token", http.StatusInternalServerError)
	// 	return
	// }

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })
}
