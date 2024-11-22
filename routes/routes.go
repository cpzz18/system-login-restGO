package routes

import (
	"myapi/controllers"
	"myapi/middleware"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

)

func Route() {
	r := mux.NewRouter()

	// Route tanpa autentikasi
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutWithBlacklist).Methods("POST")
	r.HandleFunc("/public", controllers.Public).Methods("GET")

	// Route dengan autentikasi
	private := r.PathPrefix("/users").Subrouter()
	private.Use(middleware.ValidateToken)
	private.HandleFunc("", controllers.GetUsers).Methods("GET")
	private.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	private.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")
	private.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")

	// Middleware logging untuk semua route
	r.Use(middleware.LoggingMiddleware)

	log.Println("server run on port 8000")
	http.ListenAndServe(":8000", r)
}
