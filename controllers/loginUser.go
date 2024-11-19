package controllers

import (
	"encoding/json"
	"myapi/config"
	"myapi/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"

)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var login models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	var user models.Register
	if err := config.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		http.Error(w, "User not Found", http.StatusNotFound)
		return
	}
	

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login Success"})
}
