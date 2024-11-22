package controllers

import (
	"myapi/utils"
	"net/http"

)

func Public(w http.ResponseWriter, r *http.Request) {
	respon := map[string]string{
		"message": "Hello, this is public endpoint",
	}

	w.Header().Set("Content-Type", "application/json")
	utils.RespondJson(w , http.StatusOK, respon)

}
