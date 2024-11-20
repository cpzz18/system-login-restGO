package controllers

import (
	"myapi/config"
	"myapi/models"
	"myapi/utils"
	"net/http"

	"github.com/gorilla/mux"

)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []models.Register

	if err := config.DB.Find(&users).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "user not found")
	}
	utils.RespondJson(w, http.StatusOK, map[string]interface{} {
		"data": users,
	})

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//mengambil id dari parameter url
	id := mux.Vars(r)["id"]

	//inisiasi user dan lakukan pencarian bedasarkan url
	var users models.Register
	if err := config.DB.First(&users, "id = ?", id).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "user not found{id}")
		return
	}

	utils.RespondJson(w, http.StatusOK, map[string]interface{} {
		"data": users,
	})

}
