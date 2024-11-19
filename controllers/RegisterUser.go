package controllers

import (
	"encoding/json"
	"myapi/config"
	"myapi/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"

)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users models.Register

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
	}

	users.Password = string(hash)

	if err := config.DB.Create(&users).Error; err != nil {
		http.Error(w, "failed to create", http.StatusInternalServerError)
		return
	}

	users.Password = ""
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)

}
