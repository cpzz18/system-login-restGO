package controllers

import (
	"myapi/config"
	"myapi/models"
	"myapi/utils"
	"net/http"

	"github.com/gorilla/mux"

)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var users models.Register
	if err := config.DB.First(&users, id).Error; err != nil {
		utils.RespondError(w,  http.StatusNotFound, "id not found")
		return
	}


	if err := config.DB.Delete(&users, id).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "delete failed")
		return
	}

	utils.RespondJson(w, http.StatusOK, map[string]interface{} {
		"message" : "user succesfuly deleted",
		"data" : users,
	})
}
