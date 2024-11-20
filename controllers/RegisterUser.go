package controllers

import (
	"encoding/json"
	"myapi/config"
	"myapi/models"
	"myapi/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"

)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users models.Register

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	//cek apakah users sudah ada
	if err := config.DB.Where("email = ? ", users.Email).First(&users).Error; err == nil {
		utils.RespondError(w, http.StatusBadRequest, "Email already exists")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to hash password")
	}

	users.Password = string(hash)

	if err := config.DB.Create(&users).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	users.Password = ""
	utils.RespondJson(w, http.StatusCreated, map[string]interface{} {
		"message": "User created successfully",
		"data": users,
	})

}
