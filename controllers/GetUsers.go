package controllers

import (
	"encoding/json"
	"myapi/config"
	"myapi/models"
	"net/http"

	"github.com/gorilla/mux"

)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []models.Register

	err := config.DB.Find(&users).Error
	if err != nil {
		w.WriteHeader(500)
		resp, _ := json.Marshal(map[string]string{"status": "failed"})
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(&users)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//mengambil id dari parameter url
	id := mux.Vars(r)["id"]

	//inisiasi user dan lakukan pencarian bedasarkan url
	var users models.Register
	if err := config.DB.First(&users, "id = ?", id).Error; err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, `{"status": "failed to encode users"}`, http.StatusInternalServerError)
	}

}
