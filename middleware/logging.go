package middleware

import (
	"log"
	"myapi/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log setiap permintaan yang masuk
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Melanjutkan ke handler berikutnya
	})
}

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondError(w, http.StatusUnauthorized, "Missing Authorization Header")
			return
		}

		// format header sesuai dengan 'Bearer <token>'
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid Authorization Header")
			return
		}

		// Parse dan verifikasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler 
			}
			// Kembalikan secret key untuk verifikasi token
			return []byte(utils.SecretKey), nil
		})

		if err != nil || !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
