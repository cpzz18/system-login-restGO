package controllers

import (
	"encoding/json"
	"myapi/config"
	"myapi/models"
	"myapi/utils"
	"net/http"

	"github.com/gorilla/mux"

)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var users models.Register
	//mencari id ada apa tidak
	if err := config.DB.First(&users, "id = ? ", id).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "User not found")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	//meng update
	if err := config.DB.Model(&users).Updates(users).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to update user")
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to encode user")
	}
	
	utils.RespondJson(w, http.StatusOK, map[string]interface{} {
		"message": "User updated successfully",
		"data" : users,
	})

}
