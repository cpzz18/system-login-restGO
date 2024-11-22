package controllers

import (
	"myapi/middleware"
	"myapi/utils"
	"net/http"
	"strings"

)

// Fungsi untuk logout dan menambahkan token ke dalam blacklist
func LogoutWithBlacklist(w http.ResponseWriter, r *http.Request) {
	// Ambil token dari header Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.RespondError(w, http.StatusUnauthorized, "Missing Authorization Header")
		return
	}

	// Format header sesuai dengan 'Bearer <token>'
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		utils.RespondError(w, http.StatusUnauthorized, "Invalid Authorization Header")
		return
	}

	// Tambahkan token ke blacklist
	middleware.BlacklistToken[tokenString] = true

	// Kirim response
	response := map[string]string{
		"message": "Logout successful. Token is now blacklisted.",
	}
	utils.RespondJson(w, http.StatusOK, response)
}
