package utils

import (
	"encoding/json"
	"net/http"

)

func RespondJson(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJson(w, status, map[string]string{"error": message})
}
