package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/db"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/utils"
	"golang.org/x/crypto/bcrypt"
)

func UserHandler(userDB *db.UserDB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetUser(userDB, w, r)
	case http.MethodPost:
		handlePostUser(userDB, w, r)
	case http.MethodPut:
		handlePutUser(userDB, w, r)
	case http.MethodDelete:
		handleDeleteUser(userDB, w, r)
	default:
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}

// handleGetUser handles GET requests for the /user route
func handleGetUser(userDB *db.UserDB, w http.ResponseWriter, r *http.Request) {
	// Assuming the user ID is passed as a query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	user, err := userDB.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// / handlePostUser handles POST requests for the /user route
func handlePostUser(userDB *db.UserDB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Normalize the email address
	normalizedEmail := utils.NormalizeEmail(user.Email)
	user.Email = normalizedEmail // Update the user's email to the normalized one

	// Check if the email already exists
	_, err := userDB.GetUser(user.Email)
	if err == nil {
		http.Error(w, "Email already in use", http.StatusBadRequest)
		return
	}
	// Assuming the error returned when no user is found is a specific error, you might want to check for that
	// and continue only if the error is about not finding the user.

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Now, user.Password contains the hashed password
	if err := userDB.CreateUser(user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// Consider not returning the password, even if it's hashed
	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

// handlePutUser handles PUT requests for the /user route
func handlePutUser(userDB *db.UserDB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if err := userDB.UpdateUser(user); err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// handleDeleteUser handles DELETE requests for the /user route
func handleDeleteUser(userDB *db.UserDB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	err := userDB.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
