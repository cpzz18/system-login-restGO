package controllers

import (
	"encoding/json"
	"myapi/config"
	"myapi/models"
	"net/http"

	"github.com/gorilla/mux"

)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	id := mux.Vars(r)["id"]
	
	var users models.Register
	//mencari id ada apa tidak
	if err := config.DB.First(&users, "id = ? ", id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, "Invalid rquest body", http.StatusBadRequest)
		return
	}

	//meng update 
	if err := config.DB.Model(&users).Updates(users).Error; err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "failed to encode user data", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	
	
}
