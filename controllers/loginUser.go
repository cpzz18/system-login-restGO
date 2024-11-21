package controllers

import (
	"encoding/json"
	"myapi/config"
	"myapi/models"
	"myapi/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"

)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var login models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	var user models.Register
	if err := config.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to generate token")
	}

	utils.RespondJson(w, http.StatusOK, map[string]interface{}{
		"message": "login successfully",
		"token": token,
		"data": map[string]interface{}{
			"name":  user.Username,
			"email": user.Email,
		},
	})
}
