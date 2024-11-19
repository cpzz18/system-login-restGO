package controllers

import (
	"encoding/json"
	"myapi/config"
	"myapi/models"
	"net/http"

	"github.com/gorilla/mux"

)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var users models.Register
	if err := config.DB.First(&users, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
	}
	
	if err := config.DB.Delete(&users, id).Error; err != nil {
		http.Error(w, "failed to delete", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusNoContent)
}
